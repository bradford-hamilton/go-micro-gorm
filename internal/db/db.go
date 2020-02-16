package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	// Needed for side effects (pg driver)
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Emedding `gorm.Model` adds the following conventional db fields:
//		ID        uint `gorm:"primary_key"`
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 		DeletedAt *time.Time `sql:"index"`
//
// By keeping the `DeletedAt` field, gorm will mark records with a
// timestamp instead to destroying them

// Message defines the model for a Message in the system and embeds a gorm model
type Message struct {
	gorm.Model
	MsgType string
	Msg     string
}

// User defines the model for a User in the system and embeds a gorm model
type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	MailingList bool `gorm:"default:false"`
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

	// Double check that our connection is good
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&Message{}, &User{})

	// Seed the db when in development
	if env := os.Getenv("GO_MICRO_GORM_ENV"); env == "development" {
		// Seed database with some records in development
		seedDatabase(db)
	}

	return db, nil
}

// Seed the db with a couple of each message type
func seedDatabase(db *gorm.DB) {
	db.Create(&Message{MsgType: "Public", Msg: "Hey, I'm a public message!"})
	db.Create(&Message{MsgType: "Public", Msg: "I'm another public message."})
	db.Create(&Message{MsgType: "Private", Msg: "Hey, I'm a private message!"})
	db.Create(&Message{MsgType: "Private", Msg: "I'm another private message."})
	db.Create(&Message{MsgType: "Protected", Msg: "Hey, I'm a protected message!"})
	db.Create(&Message{MsgType: "Protected", Msg: "I'm another protected message."})
}

// ErrSliceToString is a small helper function for now to join a slice of errors, which
// is the type that is returned from a gorm GetErrors() call.
func ErrSliceToString(errs []error) []string {
	var errStrings []string
	for _, e := range errs {
		errStrings = append(errStrings, e.Error())
	}
	return errStrings
}
