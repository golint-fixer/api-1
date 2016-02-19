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
	ID             bson.ObjectId `json:"id" bson:"_id" binding:"-"`
	User           string        `json:"user" bson:"user" binding:"-"`
	NumClientNodes int           `json:"num_client_nodes" bson:"num_client_nodes" binding:"required"`
	NumDataNodes   int           `json:"num_data_nodes" bson:"num_data_nodes" binding:"required"`
	NumMasterNodes int           `json:"num_master_nodes" bson:"num_master_nodes" binding:"required,max=5"`
}

func (model *BuildModel) init() {
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"num_client_nodes"}})
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"num_data_nodes"}})
	model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"num_master_nodes"}})
}

// Collection get the collection for this data model.
func (model *BuildModel) Collection() *mgo.Collection {
	db := common.GetDatabase()
	return db.C("elasticsearch_builds")
}

// HandleValidationErrors handle validation errors related to this model.
func (model *BuildModel) HandleValidationErrors(context *gin.Context, errors validator.ValidationErrors) {
	errCollector := common.SerializeValidationErrors(model, errors)
	context.JSON(http.StatusBadRequest, gin.H{"errors": errCollector, "numErrors": len(errors)})
}
