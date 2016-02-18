package elasticsearch

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/thedodd/buildAPI/common"
)

func init() {
	(&BuildModel{}).init()
}

// BuildModel the Elasticsearch build model.
type BuildModel struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	NumClientNodes int           `json:"num_client_nodes" bson:"num_client_nodes"`
	NumDataNodes   int           `json:"num_data_nodes" bson:"num_data_nodes"`
	NumMasterNodes int           `json:"num_master_nodes" bson:"num_master_nodes"`
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
