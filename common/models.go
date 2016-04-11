package common

import "gopkg.in/mgo.v2"

// ModelInterface - the interface definition for data models.
type ModelInterface interface {
	// The collection of the model.
	Collection() *mgo.Collection

	// Ensure indices needed by the model are in place.
	EnsureIndices()
}
