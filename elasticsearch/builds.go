package elasticsearch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedodd/buildAPI/common"
	"gopkg.in/mgo.v2/bson"
)

// GetElasticsearchBuilds get a list of all Elasticsearch builds.
func GetElasticsearchBuilds(context *gin.Context) {
	var builds []BuildModel
	(&BuildModel{}).Collection().Find(bson.M{}).All(&builds)

	// If slice is nil, return empty slice.
	if builds == nil {
		builds = make([]BuildModel, 0)
	}
	context.JSON(http.StatusOK, gin.H{"data": builds})
}

// CreateElasticsearchBuild create a new Elasticsearch build.
func CreateElasticsearchBuild(context *gin.Context) {
	// TODO(TheDodd): these params need to come from POST body.
	build := &BuildModel{
		ID:             bson.NewObjectId(),
		NumClientNodes: 5,
		NumDataNodes:   5,
		NumMasterNodes: 3,
	}
	build.Collection().Insert(build)
	context.JSON(http.StatusOK, gin.H{"data": build})
}

// GetElasticsearchBuildByID get an Elasticsearch build by ID.
func GetElasticsearchBuildByID(context *gin.Context) {
	rawID := context.Param("id")
	id, err := common.GetObjectID(rawID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid build ID specified."})
	}

	build := &BuildModel{}
	err = build.Collection().FindId(id).One(build)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Build not found."})
	}
	context.JSON(http.StatusOK, gin.H{"data": build})
}
