package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Message defines the model for a Message in the system and embeds a gorm model
type Message struct {
	gorm.Model
	MsgType string
	Msg     string
}

// New creates new connection to the database and returns it to caller
func New() (*gorm.DB, error) {
	// Don't feel like setting a password on my local db
	p := os.Getenv("GO_MICRO_GORM_DB_PASSWORD")
	if p != "" {
		p = fmt.Sprintf("password=%s", p)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s %s dbname=%s sslmode=%s",
		os.Getenv("GO_MICRO_GORM_DB_HOST"),
		os.Getenv("GO_MICRO_GORM_DB_PORT"),
		os.Getenv("GO_MICRO_GORM_DB_USER"),
		p,
		os.Getenv("GO_MICRO_GORM_DB_NAME"),
		os.Getenv("GO_MICRO_GORM_SSL_MODE"),
	)

	// Open connection to the database
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&Message{})

	// Seed some records
	db.Create(&Message{MsgType: "Default", Msg: "Hey, I'm a default message"})
	db.Create(&Message{MsgType: "Default", Msg: "Hey, I'm another default message"})
	db.Create(&Message{MsgType: "Short", Msg: "Hey, I'm a short message"})
	db.Create(&Message{MsgType: "Short", Msg: "Hey, I'm another short message"})
	db.Create(&Message{MsgType: "Long", Msg: "Hey, I'm a long message"})
	db.Create(&Message{MsgType: "Long", Msg: "Hey, I'm another long message"})

	return db, nil
}
