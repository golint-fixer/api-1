/**
 * Just a build API.
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thedodd/goapi/elasticsearch"
)

func main() {
	port := 3000
	mux := mux.NewRouter()

	mux.Handle("/elasticsearch/builds/", &elasticsearch.BuildsListResource{})
	mux.Handle("/elasticsearch/builds/{buildID}", &elasticsearch.BuildsDetailResource{})

	log.Printf("API listening at 0.0.0.0:%d.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
