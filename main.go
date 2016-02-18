/**
 * Just a build API.
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/thedodd/buildAPI/elasticsearch"
)

func main() {
	mux := mux.NewRouter()

	mux.Handle("/elasticsearch/builds/", &elasticsearch.BuildsListResource{})
	mux.Handle("/elasticsearch/builds/{buildID}", &elasticsearch.BuildsDetailResource{})

	port := 3000
	log.Printf("API listening at 0.0.0.0:%d.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
