package elasticsearch

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
}

func (resource *BuildsListResource) post(response http.ResponseWriter, request *http.Request) {
	log.Println("POST /elasticsearch/builds/ 200")
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
