package users

import (
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/thedodd/api/common"
)

var indexOnce sync.Once

func init() {
	(&User{}).EnsureIndices()
}

// BaseUser - the base User model. Should only contain fields acceptable for HTTP responses.
type BaseUser struct {
	ID         bson.ObjectId `json:"id" bson:"_id" validate:"-"`
	Username   string        `json:"username" bson:"username" validate:"required,alphanum,min=2,max=255"`
	Email      string        `json:"email" bson:"email" validate:"required,email"`
	IsVerified bool          `json:"-" bson:"isVerified" validate:"-"`
}

// User - the User model. Use User.BaseUser for HTTP responses.
type User struct {
	BaseUser          `json:",squash" bson:",inline"`
	Password          string `json:"password" bson:"passwordHash" validate:"required"`
	VerificationToken string `json:"-" bson:"verificationToken" validate:"-"`
}

// Collection - get the collection for this data model.
func (model *User) Collection() *mgo.Collection {
	db := common.GetDatabase()
	return db.C("users")
}

// EnsureIndices - ensure any indices needed for this model's collection are in place.
func (model *User) EnsureIndices() {
	indexOnce.Do(func() {
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"username"}, Unique: true})
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"verificationToken"}})
	})
}

// VerificationBody - the struct modeling a user verification body.
type VerificationBody struct {
	Token string
}
