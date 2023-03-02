package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people = []Person{
    {ID: "1", Firstname: "Tom", Lastname: "Jones", Address: &Address{City: "City X", State: "State X"}},
    {ID: "2", Firstname: "Bob", Lastname: "Mulligan", Address: &Address{City: "City X", State: "State X"}},
    {ID: "3", Firstname: "Marie", Lastname: "Oliver", Address: &Address{City: "City X", State: "State X"}},
}

func main() {
    router := gin.Default()

	router.GET("/people", getPeople)
	router.GET("/people/:id", getPersonByID)
	router.POST("/people", postPerson)

    router.Run("localhost:8080")
}

func getPeople(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, people)
}

func getPersonByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range people {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func postPerson(c *gin.Context) {
    var newPerson Person

    if err := c.BindJSON(&newPerson); err != nil {
        return
    }

    people = append(people, newPerson)
    c.IndentedJSON(http.StatusCreated, newPerson)
}