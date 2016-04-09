package users

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/thedodd/api/common"
)

var indexOnce sync.Once

func init() {
	(&User{}).EnsureIndices()
}

// User - the User model.
type User struct {
	ID         bson.ObjectId `json:"id" bson:"_id" validate:"-"`
	Username   string        `json:"username" bson:"username" validate:"required,min=2,max=255"`
	Email      string        `json:"email" bson:"email" validate:"required,email"`
	IsVerified bool          `json:"-" bson:"isVerified" validate:"-"`
}

// EnsureIndices - ensure any indices needed for this model's collection are in place.
func (model *User) EnsureIndices() {
	indexOnce.Do(func() {
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"username"}, Unique: true})
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"verificationToken"}})
	})
}

// Collection - get the collection for this data model.
func (model *User) Collection() *mgo.Collection {
	db := common.GetDatabase()
	return db.C("users")
}

// HandleValidationErrors - handle validation errors related to this model.
func (model *User) HandleValidationErrors(context *gin.Context, errors validator.ValidationErrors) {
	errCollector := common.SerializeValidationErrors(model, errors)
	context.JSON(http.StatusBadRequest, gin.H{"errors": errCollector, "numErrors": len(errors)})
}

// UserInbound - model for user creation.
type UserInbound struct {
	User              `json:",squash"`
	Password          string `json:"password" bson:"passwordHash" validate:"required"`
	VerificationToken string `json:"-" bson:"verificationToken" validate:"-"`
}
