package elasticsearch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedodd/api/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetElasticsearchBuilds get a list of all Elasticsearch builds.
func GetElasticsearchBuilds(context *gin.Context) {
	var builds []BuildModel
	user := context.MustGet("id").(string)
	(&BuildModel{}).Collection().Find(bson.M{"user": user}).All(&builds)

	// If slice is nil, return empty slice.
	if builds == nil {
		builds = make([]BuildModel, 0)
	}
	context.JSON(http.StatusOK, gin.H{"data": builds})
}

// CreateElasticsearchBuild create a new Elasticsearch build.
func CreateElasticsearchBuild(context *gin.Context) {
	build := context.MustGet("data").(*BuildModel)
	build.ID = bson.NewObjectId()
	build.User = context.MustGet("id").(string)

	if err := build.Collection().Insert(build); err != nil {
		abortCode, dbError := common.SerializeDBErrors(err.(*mgo.LastError))
		context.JSON(abortCode, gin.H{"errors": dbError})
	} else {
		context.JSON(http.StatusOK, gin.H{"data": build})
	}
}

// GetElasticsearchBuildByID get an Elasticsearch build by ID.
func GetElasticsearchBuildByID(context *gin.Context) {
	rawID := context.Param("id")
	id, err := common.GetObjectID(rawID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid build ID specified."})
		context.Abort()
		return
	}

	build := &BuildModel{}
	user := context.MustGet("id").(string)
	err = build.Collection().Find(bson.M{"_id": id, "user": user}).One(build)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Build not found."})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": build})
}
