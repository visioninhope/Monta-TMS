package handlers

import (
	"github.com/emoss08/trenova/internal/api"
	"github.com/emoss08/trenova/internal/api/services"
	"github.com/emoss08/trenova/internal/util"
	"github.com/emoss08/trenova/internal/util/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WorkerHandler struct {
	Server  *api.Server
	Service *services.WorkerService
}

func NewWorkerHandler(s *api.Server) *WorkerHandler {
	return &WorkerHandler{
		Server:  s,
		Service: services.NewWorkerService(s),
	}
}

// GetWorkers is a handler that returns a list of workers.
//
// GET /workers
func (h *WorkerHandler) GetWorkers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		offset, limit, err := util.PaginationParams(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(types.ValidationErrorResponse{
				Type: "invalidRequest",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "invalidRequest",
						Detail: err.Error(),
						Attr:   "offset, limit",
					},
				},
			})
		}

		orgID, ok := c.Locals(util.CTXOrganizationID).(uuid.UUID)
		buID, buOK := c.Locals(util.CTXBusinessUnitID).(uuid.UUID)

		if !ok || !buOK {
			return c.Status(fiber.StatusInternalServerError).JSON(types.ValidationErrorResponse{
				Type: "internalError",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "internalError",
						Detail: "Organization ID or Business Unit ID not found in the request context",
						Attr:   "orgID, buID",
					},
				},
			})
		}

		entities, count, err := h.Service.GetWorkers(c.UserContext(), limit, offset, orgID, buID)
		if err != nil {
			errorResponse := util.CreateDBErrorResponse(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
		}

		nextURL := util.GetNextPageURL(c, limit, offset, count)
		prevURL := util.GetPrevPageURL(c, limit, offset)

		return c.Status(fiber.StatusOK).JSON(types.HTTPResponse{
			Results:  entities,
			Count:    count,
			Next:     nextURL,
			Previous: prevURL,
		})
	}
}

// CreateWorker is a handler that creates a worker.
//
// POST /workers
func (h *WorkerHandler) CreateWorker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		newEntity := new(services.WorkerRequest)

		orgID, ok := c.Locals(util.CTXOrganizationID).(uuid.UUID)
		buID, buOK := c.Locals(util.CTXBusinessUnitID).(uuid.UUID)

		if !ok || !buOK {
			return c.Status(fiber.StatusInternalServerError).JSON(types.ValidationErrorResponse{
				Type: "internalError",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "internalError",
						Detail: "Organization ID or Business Unit ID not found in the request context",
						Attr:   "orgID, buID",
					},
				},
			})
		}

		newEntity.BusinessUnitID = buID
		newEntity.OrganizationID = orgID

		if err := util.ParseBodyAndValidate(c, newEntity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(types.ValidationErrorResponse{
				Type: "invalidRequest",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "invalidRequest",
						Detail: err.Error(),
						Attr:   "body",
					},
				},
			})
		}

		entity, err := h.Service.CreateWorker(c.UserContext(), newEntity)
		if err != nil {
			errorResponse := util.CreateDBErrorResponse(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
		}

		return c.Status(fiber.StatusCreated).JSON(entity)
	}
}

// UpdateWorker is a handler that updates a worker.
//
// PUT /workers/:workerID
func (h *WorkerHandler) UpdateWorker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		workerID := c.Params("workerID")
		if workerID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(types.ValidationErrorResponse{
				Type: "invalidRequest",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "invalidRequest",
						Detail: "Email Profile ID is required",
						Attr:   "workerID",
					},
				},
			})
		}

		updatedEntity := new(services.WorkerUpdateRequest)

		if err := util.ParseBodyAndValidate(c, updatedEntity); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(types.ValidationErrorResponse{
				Type: "invalidRequest",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "invalidRequest",
						Detail: err.Error(),
						Attr:   "request body",
					},
				},
			})
		}

		updatedEntity.ID = uuid.MustParse(workerID)

		entity, err := h.Service.UpdateWorker(c.UserContext(), updatedEntity)
		if err != nil {
			errorResponse := util.CreateDBErrorResponse(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
		}

		return c.Status(fiber.StatusOK).JSON(entity)
	}
}
