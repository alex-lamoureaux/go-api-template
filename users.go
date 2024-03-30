package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City  string `json:"city"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type name struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName1 string `json:"middle_name_1"`
}

type user struct {
	Name    name    `json:"name"`
	Address address `json:"address"`
}

func getUsers(c *gin.Context) {
	content, err := os.ReadFile("./users.json")
	if err != nil {
		log.Print("Error when opening file: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var users []user
	err = json.Unmarshal(content, &users)
	if err != nil {
		log.Print("Error unmarshalling json data: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	log.Printf("unmarshalled users: %v", users)
	c.IndentedJSON(http.StatusOK, users)
}
