package userhandler

import (
	"context"
	"strings"

	"github.com/bradford-hamilton/go-micro-gorm/internal/db"
	pb "github.com/bradford-hamilton/go-micro-gorm/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

// User represents our user service and it's available methods
type User struct {
	DB *gorm.DB
}

// Create creates a user
func (u *User) Create(ctx context.Context, req *pb.Request, res *pb.Response) error {
	log.Log("Received User.Create request")

	// Create a new user with data from the request
	newUser := db.User{
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		MailingList: req.GetMailingList(),
	}

	// Insert the new user record into the database. If there are any errors,
	// attach them to the respond and return.
	if errs := u.DB.Create(&newUser).GetErrors(); len(errs) > 0 {
		errStrings := db.ErrSliceToString(errs)
		log.Errorf("Error creating user: %v", strings.Join(errStrings, ", "))
		res.Errors = errStrings
		return nil
	}

	res.Message = "Succesfully created user!"

	return nil
}
