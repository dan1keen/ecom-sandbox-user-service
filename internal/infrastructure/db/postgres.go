package db

import (
	"fmt"
	"log"
	"sync"
	"user-service/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// GetPostgresDB returns singleton *sql.DB
func GetPostgresDB() *gorm.DB {
	dbOnce.Do(func() {
		c := config.LoadDBConfig()
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.DBName,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}
	})

	return db
}
