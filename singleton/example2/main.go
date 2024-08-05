package main

import (
	"coding_challenges/singleton/example2/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", getUsers)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method Not Allowed", "source": ctx.ClientIP()})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func getUsers(c *gin.Context) {
	conn := db.GetInstance()
	rows, err := conn.Query("SELECT id, name FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	for rows.Next() {
		var user struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
