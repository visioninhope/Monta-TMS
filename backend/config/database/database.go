package database

import (
	"context"
	"gorm.io/gorm/logger"
	"log"
	"time"
	"trenova/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

type DBConfig struct {
	DSN             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// ConnectDb Connect to the database
func ConnectDb(config DBConfig) (*gorm.DB, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, cancel, err
	}

	// Migrating Database
	var mods = []interface{}{
		&models.BusinessUnit{},
		&models.Organization{},
		&models.EmailProfile{},
		&models.AccountingControl{},
		&models.BillingControl{},
		&models.DispatchControl{},
		&models.InvoiceControl{},
		&models.ShipmentControl{},
		&models.FeasibilityToolControl{},
		&models.RouteControl{},
		&models.Token{},
		&models.User{},
		&models.UserNotifications{},
		&models.UserFavorite{},
		&models.JobTitle{},
		&models.Tag{},
		&models.GeneralLedgerAccount{},
		&models.DivisionCode{},
		&models.RevenueCode{},
		&models.TableChangeAlert{},
	}

	if err := db.AutoMigrate(mods...); err != nil {
		return nil, cancel, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, cancel, err
	}

	select {
	case <-ctx.Done():
		return nil, cancel, ctx.Err()
	default:
		// Set connection pool settings
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
		sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	}

	log.Println("Connected to the database")

	DB = Dbinstance{Db: db}

	return db, cancel, nil
}