package database

import (
	"log"
	"motionserver/app/database/schema"
	"motionserver/app/database/seeds"
	"motionserver/internal/bootstrap/seeder"
	"motionserver/utils/config"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	Log zerolog.Logger
	Cfg *config.Config
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (_db *Database) ConnectDatabase() {
	conn, err := gorm.Open(postgres.Open(_db.Cfg.DB.Postgres.DSN), &gorm.Config{})
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database!")
	} else {
		_db.Log.Info().Msg("Connected the database succesfully!")
	}

	_db.DB = conn
}

func (_db *Database) ShutdownDatabase() {
	sqlDB, err := _db.DB.DB()
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	} else {
		_db.Log.Info().Msg("Shutdown the database succesfully!")
	}
	sqlDB.Close()
}

func (_db *Database) MigrateModels() {
	err := _db.DB.Exec(`
	DO $$ BEGIN 
		CREATE TYPE role as ENUM ('admin', 'superadmin', 'user')	;
	EXCEPTION 
		WHEN duplicate_object THEN null;
	END $$;
	`).Error
	if err != nil {
		log.Println(err)
	}
	if err := _db.DB.AutoMigrate(
		Models()...,
	); err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to migrate the database!")
	}
}

// reset models
func (_db *Database) ResetModels() {
	if err := _db.DB.Migrator().DropTable(
		Models()...,
	); err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to reset the database!")
	}
	_db.DB.Exec("drop table product_categories")
}

// list of models for migration
func Models() []interface{} {
	return []interface{}{
		schema.User{},
		schema.Account{},
		schema.Category{},
		schema.Product{},
		schema.Gallery{},
		schema.Cart{},
		schema.News{},
		schema.Compro{},
		schema.Banner{},
		schema.Youtube{},
		schema.PromotionProduct{},
	}
}

// seed data
func (_db *Database) SeedModels() {
	user := seeds.NewUserSeeder()
	category := seeds.NewCategorySeeder()
	seeder := []seeder.Seeder{
		user,
		category,
	}
	for _, seed := range seeder {

		count, err := seed.Count(_db.DB)
		if err != nil {
			_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
		}

		if count == 0 {
			if err := seed.Seed(_db.DB); err != nil {
				_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
			}

			_db.Log.Info().Msg("Seeded the database succesfully!")
		} else {
			_db.Log.Info().Msg("Database is already seeded!")
		}
	}

	_db.Log.Info().Msg("Seeded the database succesfully!")
}
