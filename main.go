package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const dbHost = "localhost"

type User struct {
	Id bson.ObjectId
	Name string
	Email string
}

func main() {
	// Connect to the database.
	session, err := mgo.Dial(dbHost)
	handleError(err)

	// Close the db session after the end of the `main` function.
	defer session.Close()

	// Set the mode of monotonic.
	session.SetMode(mgo.Monotonic, true)

	// Get the "users" collection on the "testdb" database.
	c := session.DB("testdb").C("users")

	// Remove all the "users" documents beforehand.
	_, err = c.RemoveAll(nil)

	// Insert the documents into the "users" collection.
	err = c.Insert(&User{bson.NewObjectId(), "Taro Yamada", "taro.yamada@example.com"},
	&User{bson.NewObjectId(), "Hanako Tanaka", "hanako.tanaka@example.com"})
	handleError(err)

	// Find the all documents
	findAll(c)

	// Find the document of "Taro Yamada".
	result := User{}
	err = c.Find(bson.M{"name": "Taro Yamada"}).One(&result)
	handleError(err)
	fmt.Printf("User: %+v\n", result)

	// Update the document of "Hanako Tanaka".
	err = c.Update(bson.M{"name": "Hanako Tanaka"}, bson.M{"$set": bson.M{"email": "hanako.tanaka@updated.com"}})
	handleError(err)
	
	// Find the all documents
	findAll(c)

	// Remove the document of "Taro Yamada".
	err = c.Remove(bson.M{"name": "Taro Yamada"})
	handleError(err)

	// Find the all documents
	findAll(c)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func findAll(c *mgo.Collection) {
	results := []User{}
	err := c.Find(nil).All(&results)
	handleError(err)
	fmt.Printf("Users: %+v\n", results)
}
