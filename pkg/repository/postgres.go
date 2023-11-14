package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	usersTable     = "users"
	todoItemsTable = "todo_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	connString := fmt.Sprintf(`host=%s 
										port=%s 
										user=%s 
										dbname=%s 
										password=%s
										sslmode=%s`,
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to database: ", err.Error())
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Couldn't get generic database object sql.DB to use its functions: ", err.Error())
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		log.Println("Couldn't send ping to database: ", err.Error())
		return nil, err
	}

	return db, nil
}

func PostgresCloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
