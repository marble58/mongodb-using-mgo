package main

import (
	"labix.org/v2/mgo"
	"testing"
	"strconv"
)

func TestMain(t *testing.T) {
	session, _ := mgo.Dial(dbHost)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("testdb").C("users")
	c.RemoveAll(nil)
	main()
	results := []User{}
	c.Find(nil).All(&results)

	testLen := 1
	if testLen != len(results) {
		t.Error("results should be " + strconv.Itoa(testLen) + ".")
	}
	result := results[0]

	testName := "Hanako Tanaka"
	if testName != result.Name {
		t.Error("Name should be " + testName + ".")
	}

	testEmail := "hanako.tanaka@updated.com"
	if testEmail != result.Email {
		t.Error("Email should be " + testEmail + ".")
	}
}
