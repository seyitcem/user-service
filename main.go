package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var address string = "localhost:8080"

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Money   int    `json:"money"`
}

func main() {
	var users []User = []User{}
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})
	router.POST("/users/add", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}
		for i := 0; i < len(users); i++ {
			if users[i].ID == newUser.ID {
				c.IndentedJSON(http.StatusConflict, gin.H{"message": "Each user must have unique id."})
				return
			}
		}
		users = append(users, newUser)
		c.IndentedJSON(http.StatusCreated, newUser)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			for _, user := range users {
				if user.ID == id {
					c.IndentedJSON(http.StatusOK, user)
					return
				}
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
	})
	router.Run(address)
}
