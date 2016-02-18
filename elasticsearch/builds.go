package elasticsearch

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// BuildsListResource a request handler Elasticsearch builds.
type BuildsListResource struct{}

func (resource *BuildsListResource) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resource.get(response, request)
		break

	case "POST":
		resource.post(response, request)
		break

	default:
		http.NotFound(response, request)
	}
}

func (resource *BuildsListResource) get(response http.ResponseWriter, request *http.Request) {
	log.Println("GET /elasticsearch/builds/ 200")
	var builds []BuildModel
	(&BuildModel{}).Collection().Find(bson.M{}).All(&builds)

	if builds == nil {
		response.Write([]byte(`{"data": []}`))
	} else {
		data, _ := json.Marshal(bson.M{"data": builds})
		response.Write(data)
	}
}

func (resource *BuildsListResource) post(response http.ResponseWriter, request *http.Request) {
	log.Println("POST /elasticsearch/builds/ 200")
	build := &BuildModel{
		ID:             bson.NewObjectId(),
		NumClientNodes: 5,
		NumDataNodes:   5,
		NumMasterNodes: 3,
	}
	build.Collection().Insert(build)
	data, _ := json.Marshal(build)
	response.Write(data)
}

// BuildsDetailResource a request handler Elasticsearch build details.
type BuildsDetailResource struct{}

func (resource *BuildsDetailResource) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		resource.get(response, request)
		break

	default:
		http.NotFound(response, request)
	}
}

func (resource *BuildsDetailResource) get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	log.Printf("GET /elasticsearch/builds/%s 200", vars["buildID"])
}
