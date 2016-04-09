package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/thedodd/api/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateUser persists a new user to the database unverified.
func CreateUser(context *gin.Context) {
	user := context.MustGet("data").(*UserInbound)
	user.ID = bson.NewObjectId()
	user.IsVerified = false
	user.Password = common.HashPassword(user.Password)
	user.VerificationToken = uuid.NewV4().String()

	if err := user.Collection().Insert(user); err != nil {
		abortCode, dbError := common.SerializeDBErrors(err.(*mgo.LastError))
		context.JSON(abortCode, gin.H{"errors": dbError})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": user})
	}
}
