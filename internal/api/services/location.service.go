package services

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/api"
	"github.com/emoss08/trenova/internal/ent"
	"github.com/emoss08/trenova/internal/ent/location"
	"github.com/emoss08/trenova/internal/ent/locationcomment"
	"github.com/emoss08/trenova/internal/ent/locationcontact"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/util"
	"github.com/rs/zerolog"

	"github.com/google/uuid"
)

type LocationRequest struct {
	BusinessUnitID     uuid.UUID             `json:"businessUnitId"`
	OrganizationID     uuid.UUID             `json:"organizationId"`
	CreatedAt          time.Time             `json:"createdAt"`
	UpdatedAt          time.Time             `json:"updatedAt"`
	Version            int                   `json:"version" validate:"omitempty"`
	Status             location.Status       `json:"status" validate:"required,oneof=A I"`
	Code               string                `json:"code" validate:"required,max=10"`
	LocationCategoryID *uuid.UUID            `json:"locationCategoryId" validate:"omitempty"`
	Name               string                `json:"name" validate:"required"`
	Description        string                `json:"description" validate:"omitempty"`
	AddressLine1       string                `json:"addressLine1" validate:"required,max=150"`
	AddressLine2       string                `json:"addressLine2" validate:"omitempty,max=150"`
	City               string                `json:"city" validate:"required,max=150"`
	StateID            uuid.UUID             `json:"stateId" validate:"omitempty,uuid"`
	PostalCode         string                `json:"postalCode" validate:"required,max=10"`
	Longitude          float64               `json:"longitude" validate:"omitempty"`
	Latitude           float64               `json:"latitude" validate:"omitempty"`
	PlaceID            string                `json:"placeId" validate:"omitempty,max=255"`
	IsGeocoded         bool                  `json:"isGeocoded"`
	Comments           []ent.LocationComment `json:"comments" validate:"omitempty,dive"`
	Contacts           []ent.LocationContact `json:"contacts" validate:"omitempty,dive"`
}

type LocationUpdateRequest struct {
	ID uuid.UUID `json:"id,omitempty"`
	LocationRequest
}

type LocationService struct {
	Client *ent.Client
	Logger *zerolog.Logger
}

// NewLocationService creates a new location service.
func NewLocationService(s *api.Server) *LocationService {
	return &LocationService{
		Client: s.Client,
		Logger: s.Logger,
	}
}

// GetLocations gets the locations for an organization.
func (r *LocationService) GetLocations(
	ctx context.Context, limit, offset int, orgID, buID uuid.UUID,
) ([]*ent.Location, int, error) {
	entityCount, countErr := r.Client.Location.Query().Where(
		location.HasOrganizationWith(
			organization.IDEQ(orgID),
			organization.BusinessUnitIDEQ(buID),
		),
	).Count(ctx)

	if countErr != nil {
		return nil, 0, countErr
	}

	entities, err := r.Client.Location.Query().
		Limit(limit).
		WithLocationCategory().
		WithComments().
		WithContacts().
		WithState().
		Offset(offset).
		Order(
			location.ByName(
				sql.OrderDesc(),
			),
		).
		Where(
			location.HasOrganizationWith(
				organization.IDEQ(orgID),
				organization.BusinessUnitIDEQ(buID),
			),
		).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return entities, entityCount, nil
}

// CreateLocation creates a new location.
func (r *LocationService) CreateLocation(
	ctx context.Context, newEntity *LocationRequest,
) (*ent.Location, error) {
	var createdEntity *ent.Location

	err := util.WithTx(ctx, r.Client, func(tx *ent.Tx) error {
		var err error
		createdEntity, err = r.createLocationEntity(ctx, tx, newEntity)
		if err != nil {
			return err
		}

		// If comments are provided, create them and associate them with the location
		if len(newEntity.Comments) > 0 {
			if err = r.createLocationComments(ctx, tx, createdEntity.ID, newEntity); err != nil {
				return err
			}
		}

		// If locations are provided, create them and associate them with the location
		if len(newEntity.Contacts) > 0 {
			if err = r.createLocationContacts(ctx, tx, createdEntity.ID, newEntity); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return createdEntity, nil
}

func (r *LocationService) createLocationEntity(
	ctx context.Context, tx *ent.Tx, newEntity *LocationRequest,
) (*ent.Location, error) {
	return tx.Location.Create().
		SetOrganizationID(newEntity.OrganizationID).
		SetBusinessUnitID(newEntity.BusinessUnitID).
		SetCode(newEntity.Code).
		SetStatus(newEntity.Status).
		SetNillableLocationCategoryID(newEntity.LocationCategoryID).
		SetName(newEntity.Name).
		SetAddressLine1(newEntity.AddressLine1).
		SetAddressLine2(newEntity.AddressLine2).
		SetCity(newEntity.City).
		SetStateID(newEntity.StateID).
		SetPostalCode(newEntity.PostalCode).
		Save(ctx)
}

func (r *LocationService) createLocationComments(
	ctx context.Context, tx *ent.Tx, locationID uuid.UUID, newEntity *LocationRequest,
) error {
	for _, comment := range newEntity.Comments {
		_, err := tx.LocationComment.Create().
			SetLocationID(locationID).
			SetComment(comment.Comment).
			SetBusinessUnitID(newEntity.BusinessUnitID).
			SetOrganizationID(newEntity.OrganizationID).
			SetUserID(comment.UserID).
			SetCommentTypeID(comment.CommentTypeID).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateLocation updates a location and its associated comments.
func (r *LocationService) UpdateLocation(ctx context.Context, entity *LocationUpdateRequest) (*ent.Location, error) {
	var updatedEntity *ent.Location

	err := util.WithTx(ctx, r.Client, func(tx *ent.Tx) error {
		var err error
		updatedEntity, err = r.updateLocationEntity(ctx, tx, entity)
		if err != nil {
			return err
		}

		if err = r.syncContacts(ctx, tx, entity, updatedEntity); err != nil {
			return err
		}

		return r.syncComments(ctx, tx, entity, updatedEntity)
	})
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (r *LocationService) updateLocationEntity(
	ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest,
) (*ent.Location, error) {
	current, err := tx.Location.Get(ctx, entity.ID)
	if err != nil {
		return nil, err
	}

	if current.Version != entity.Version {
		return nil, util.NewValidationError("This record has been updated by another user. Please refresh and try again",
			"syncError",
			"code")
	}

	return tx.Location.UpdateOneID(entity.ID).
		SetCode(entity.Code).
		SetStatus(entity.Status).
		SetNillableLocationCategoryID(entity.LocationCategoryID).
		SetName(entity.Name).
		SetAddressLine1(entity.AddressLine1).
		SetAddressLine2(entity.AddressLine2).
		SetCity(entity.City).
		SetStateID(entity.StateID).
		SetPostalCode(entity.PostalCode).
		SetVersion(entity.Version + 1).
		Save(ctx)
}

func (r *LocationService) syncComments(
	ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest, updatedEntity *ent.Location,
) error {
	existingComments, err := tx.Location.QueryComments(updatedEntity).Where(
		locationcomment.HasLocationWith(location.IDEQ(entity.ID)),
	).All(ctx)
	if err != nil {
		return err
	}

	// Delete unmatched comments
	if err = r.deleteUnmatchedComments(ctx, tx, entity, existingComments); err != nil {
		return err
	}

	// Update or create new comments
	return r.updateOrCreateComments(ctx, tx, entity)
}

func (r *LocationService) deleteUnmatchedComments(
	ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest, existingComments []*ent.LocationComment,
) error {
	commentPresent := make(map[uuid.UUID]bool)
	for _, comment := range entity.Comments {
		if comment.ID != uuid.Nil {
			commentPresent[comment.ID] = true
		}
	}

	for _, existingComment := range existingComments {
		if !commentPresent[existingComment.ID] {
			if err := tx.LocationComment.DeleteOneID(existingComment.ID).Exec(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *LocationService) updateOrCreateComments(ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest) error {
	for _, comment := range entity.Comments {
		if comment.ID != uuid.Nil {
			if _, err := tx.LocationComment.UpdateOneID(comment.ID).
				SetComment(comment.Comment).
				SetUserID(comment.UserID).
				SetCommentTypeID(comment.CommentTypeID).
				Save(ctx); err != nil {
				return err
			}
		} else {
			if _, err := tx.LocationComment.Create().
				SetLocationID(entity.ID).
				SetComment(comment.Comment).
				SetBusinessUnitID(entity.BusinessUnitID).
				SetOrganizationID(entity.OrganizationID).
				SetUserID(comment.UserID).
				SetCommentTypeID(comment.CommentTypeID).
				Save(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

// createLocationContact create a location contact entity.
func (r *LocationService) createLocationContacts(
	ctx context.Context, tx *ent.Tx, locationID uuid.UUID, newEntity *LocationRequest,
) error {
	for _, contact := range newEntity.Contacts {
		_, err := tx.LocationContact.Create().
			SetLocationID(locationID).
			SetBusinessUnitID(newEntity.BusinessUnitID).
			SetOrganizationID(newEntity.OrganizationID).
			SetName(contact.Name).
			SetEmailAddress(contact.EmailAddress).
			SetPhoneNumber(contact.PhoneNumber).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LocationService) syncContacts(
	ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest, updatedEntity *ent.Location,
) error {
	existingContacts, err := tx.Location.QueryContacts(updatedEntity).Where(
		locationcontact.HasLocationWith(location.IDEQ(entity.ID)),
	).All(ctx)
	if err != nil {
		return err
	}

	// Delete unmatched contacts
	if err = r.deleteUnmatchedContacts(ctx, tx, entity, existingContacts); err != nil {
		return err
	}

	// Update or create new contacts
	return r.updateOrCreateContacts(ctx, tx, entity)
}

func (r *LocationService) deleteUnmatchedContacts(
	ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest, existingContacts []*ent.LocationContact,
) error {
	contactPresent := make(map[uuid.UUID]bool)
	for _, contact := range entity.Contacts {
		if contact.ID != uuid.Nil {
			contactPresent[contact.ID] = true
		}
	}

	for _, existingContact := range existingContacts {
		if !contactPresent[existingContact.ID] {
			if err := tx.LocationComment.DeleteOneID(existingContact.ID).Exec(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *LocationService) updateOrCreateContacts(ctx context.Context, tx *ent.Tx, entity *LocationUpdateRequest) error {
	for _, contact := range entity.Contacts {
		if contact.ID != uuid.Nil {
			if _, err := tx.LocationContact.UpdateOneID(contact.ID).
				SetName(contact.Name).
				SetEmailAddress(contact.EmailAddress).
				SetPhoneNumber(contact.PhoneNumber).
				Save(ctx); err != nil {
				return err
			}
		} else {
			if _, err := tx.LocationContact.Create().
				SetLocationID(entity.ID).
				SetBusinessUnitID(entity.BusinessUnitID).
				SetOrganizationID(entity.OrganizationID).
				SetName(contact.Name).
				SetEmailAddress(contact.EmailAddress).
				SetPhoneNumber(contact.PhoneNumber).
				Save(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
