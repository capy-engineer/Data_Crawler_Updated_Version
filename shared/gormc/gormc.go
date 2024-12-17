package gormc

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type GormDBType int

const (
	GormDBTypeMySQL GormDBType = iota + 1
	GormDBTypePostgres
	GormDBTypeSQLite
	GormDBTypeMSSQL
	GormDBTypeNotSupported
)

type GormParam struct {
	dsn                   string
	dbType                string
	maxOpenConnections    int
	maxIdleConnections    int
	maxConnectionIdleTime int
}

type GormDB struct {
	id string
	*GormParam
}

func NewGormDB(id string) *GormDB {
	return &GormDB{
		id:        id,
		GormParam: new(GormParam),
	}
}

// Init initializes the database connection and returns the instance
func Init() *gorm.DB {
	dbInstance, err := NewGormDB("gorm").connectionDB()

	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %w", err))
	}

	// Call migration function after initializing DB
	migration(dbInstance)

	return dbInstance
}

// migration handles database migrations
func migration(db *gorm.DB) {
	// Example migration logic
	fmt.Println("Running migrations...")
	// Uncomment and add migration models as needed:
	// db.AutoMigrate(&YourModel{})
}

// connectionDB connects to the database based on the GormParam configuration
func (gormDB *GormDB) connectionDB() (*gorm.DB, error) {
	// Retrieve database configuration from environment variables
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	// Fallback to default values if environment variables are not set

	// Build the DSN for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Get the database type
	dbType := getDBType("mysql") // Hardcoded for now, update based on your configuration
	if dbType == GormDBTypeNotSupported {
		return nil, errors.New("database type not supported")
	}

	switch dbType {
	case GormDBTypeMySQL:
		dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
		}

		// Configure the connection pool
		if sqlDB, err := dbConnection.DB(); err == nil {
			sqlDB.SetConnMaxLifetime(time.Second * time.Duration(gormDB.maxConnectionIdleTime))
			sqlDB.SetMaxIdleConns(gormDB.maxIdleConnections)
			sqlDB.SetMaxOpenConns(gormDB.maxOpenConnections)
		}
		return dbConnection, nil
	default:
		return nil, errors.New("database type not supported")
	}
}

// getDBType returns the GormDBType based on the string identifier
func getDBType(dbType string) GormDBType {
	switch dbType {
	case "mysql":
		return GormDBTypeMySQL
	case "postgres":
		return GormDBTypePostgres
	case "sqlite":
		return GormDBTypeSQLite
	case "mssql":
		return GormDBTypeMSSQL
	default:
		return GormDBTypeNotSupported
	}
}
