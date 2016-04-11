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

// VerifyUser verifies that a user is legitimate by checking a given verification token.
func VerifyUser(context *gin.Context) {
	// Get username.
	username := context.Param("username")
	verificationToken := context.MustGet("data").(*VerificationBody).Token

	// Query for user.
	user := &User{}
	if err := user.Collection().Find(bson.M{"username": username}).One(user); err != nil {
		// A unique index exists on this field, so no user with matching username was found.
		abortCode, dbError := common.SerializeDBErrors(err.(*mgo.LastError))
		context.JSON(abortCode, gin.H{"errors": dbError})
		return
	}

	// Ensure user has not already been verified.
	if user.IsVerified != false {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}

	// Ensure verification tokens match up.
	if common.CheckHash(user.VerificationToken, verificationToken) != true {
		context.JSON(http.StatusBadRequest, gin.H{"errors": []map[string]string{
			{"error": "Error encountered when verifying token.", "message": "Invalid token provided."},
		}})
		return
	}

	// Verification complete. Update model.
	update := bson.M{"$set": map[string]bool{"isVerified": true}}
	if err := user.Collection().UpdateId(user.ID, update); err != nil {
		abortCode, dbError := common.SerializeDBErrors(err.(*mgo.LastError))
		context.JSON(abortCode, gin.H{"errors": dbError})
	}
	context.AbortWithStatus(http.StatusNoContent)
}
