package common

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB access concurrency control symbols.
var apiSession *mgo.Session
var dbSync sync.Once

// GetDatabase get a handle to this API's MongoDB database.
func GetDatabase() *mgo.Database {
	config := GetConfig()
	session := GetSession()
	return session.DB(config.BackendDBName)
}

// GetObjectID get an ObjectId from the given string, or error.
func GetObjectID(id string) (oid bson.ObjectId, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Invalid ObjectId.")
		}
	}()
	oid = bson.ObjectIdHex(id)
	return oid, err
}

// GetSession get this API's current MongoDB session.
func GetSession() *mgo.Session {
	dbSync.Do(func() {
		apiSession = dialSession()
	})
	return apiSession
}

// SerializeDBErrors serialize the given *mgo.LastError as a reasonable API response.
func SerializeDBErrors(err *mgo.LastError) (abortCode int, dbError map[string]interface{}) {
	dbError = make(map[string]interface{})
	switch err.Code {
	case 11000:
		abortCode = 400
		dbError["error"] = "Resource already exists."

	default:
		abortCode = 500
		dbError["error"] = "Database error."
		dbError["dbError"] = err.Code
	}
	return abortCode, dbError
}

//////////////////////
// Private symbols. //
//////////////////////
func dialSession() *mgo.Session {
	config := GetConfig()
	var session *mgo.Session
	var err error

	// Establish a session according to API mode.
	if config.Mode == "prod" {
		dialInfo := &mgo.DialInfo{
			Addrs:    []string{config.BackendURL},
			Database: config.BackendDBName,
			Username: config.BackendUsername,
			Password: config.BackendPassword,
			DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
				return tls.Dial("tcp", addr.String(), &tls.Config{})
			},
			Timeout: time.Second * 10,
		}
		session, err = mgo.DialWithInfo(dialInfo)
	} else {
		connectionString := fmt.Sprintf("%s/%s", config.BackendURL, config.BackendDBName)
		session, err = mgo.Dial(connectionString)
	}

	if err != nil {
		panic(err)
	}

	// Ensure we are operating in strong mode.
	session.SetMode(mgo.Strong, true)
	return session
}
