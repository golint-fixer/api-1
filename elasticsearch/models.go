package elasticsearch

import (
	"net/http"

	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/thedodd/buildAPI/common"
)

func init() {
	(&BuildModel{}).init()
}

// BuildModel the Elasticsearch build model.
type BuildModel struct {
	ID             bson.ObjectId `json:"id" bson:"_id" validate:"-"`
	User           string        `json:"user" bson:"user" validate:"-"`
	NumClientNodes int           `json:"numClientNodes" bson:"numClientNodes" validate:"required"`
	NumDataNodes   int           `json:"numDataNodes" bson:"numDataNodes" validate:"required"`
	NumMasterNodes int           `json:"numMasterNodes" bson:"numMasterNodes" validate:"required,max=5"`
}

func (model *BuildModel) init() {
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numClientNodes"}})
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numDataNodes"}})
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numMasterNodes"}})
}

// Collection get the collection for this data model.
func (model *BuildModel) Collection() *mgo.Collection {
	db := common.GetDatabase()
	return db.C("elasticsearchBuilds")
}

// HandleValidationErrors handle validation errors related to this model.
func (model *BuildModel) HandleValidationErrors(context *gin.Context, errors validator.ValidationErrors) {
	errCollector := common.SerializeValidationErrors(model, errors)
	context.JSON(http.StatusBadRequest, gin.H{"errors": errCollector, "numErrors": len(errors)})
}
