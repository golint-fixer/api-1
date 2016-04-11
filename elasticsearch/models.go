package elasticsearch

import (
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/thedodd/api/common"
)

var indexOnce sync.Once

func init() {
	(&BuildModel{}).EnsureIndices()
}

// BuildModel - the Elasticsearch build model.
type BuildModel struct {
	ID             bson.ObjectId `json:"id" bson:"_id" validate:"-"`
	User           string        `json:"user" bson:"user" validate:"-"`
	NumClientNodes int           `json:"numClientNodes" bson:"numClientNodes" validate:"required"`
	NumDataNodes   int           `json:"numDataNodes" bson:"numDataNodes" validate:"required"`
	NumMasterNodes int           `json:"numMasterNodes" bson:"numMasterNodes" validate:"required,max=5"`
}

// Collection - get the collection for this data model.
func (model *BuildModel) Collection() *mgo.Collection {
	db := common.GetDatabase()
	return db.C("elasticsearchBuilds")
}

// EnsureIndices - ensure any indices needed for this model's collection are in place.
func (model *BuildModel) EnsureIndices() {
	indexOnce.Do(func() {
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numClientNodes"}})
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numDataNodes"}})
		model.Collection().EnsureIndex(mgo.Index{Background: true, Key: []string{"numMasterNodes"}})
	})
}
