package common

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

// GetDatabase get a handle to this API's MongoDB database.
func GetDatabase() *mgo.Database {
	config := GetConfig()
	session := getSession()
	return session.DB(config.BackendDBName)
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
var apiSession *mgo.Session
var sessionSync sync.Once

func getSession() *mgo.Session {
	sessionSync.Do(func() {
		apiSession = dialSession()
	})
	return apiSession
}

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
