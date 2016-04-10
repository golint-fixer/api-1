package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/thedodd/api/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateUser persists a new user to the database unverified.
func CreateUser(context *gin.Context) {
	user := context.MustGet("data").(*User)
	verificationToken := uuid.NewV4().String()
	user.ID = bson.NewObjectId()
	user.IsVerified = false
	user.Password = common.Hash(user.Password)
	user.VerificationToken = common.Hash(verificationToken)

	if err := user.Collection().Insert(user); err != nil {
		abortCode, dbError := common.SerializeDBErrors(err.(*mgo.LastError))
		context.JSON(abortCode, gin.H{"errors": dbError})
		return
	}

	// TODO(TheDodd): wire up email verification system. For now, just print to stdout.
	fmt.Println("verification token:", verificationToken)

	// README: we only return user.BaseUser here for acceptable HTTP response.
	context.JSON(http.StatusOK, gin.H{"data": user.BaseUser})
}
