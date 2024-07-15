package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type fruit struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Family   string `json:"family"`
}

var fruits = []fruit{
	{ID: "1", Name: "Banana", Quantity: 1, Family: "Musaceae"},
	{ID: "2", Name: "Dragon Fruit", Quantity: 26, Family: "Cactaceae"},
	{ID: "3", Name: "Peach", Quantity: 47, Family: "Rose"},
}

func sellFruit(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id params"})
		return
	}

	fruit, err := getFruitByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Fruit Not Found"})
		return
	}

	if fruit.Quantity <= 0 {

		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("no more %v in stock", fruit.Name)})
		return
	}

	fruit.Quantity -= 1

	c.IndentedJSON(http.StatusOK, fruit)
}


func getFruits(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, fruits)

}

func fruitById(c *gin.Context) {
	id := c.Param("id")

	fruit, err := getFruitByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "fruit does not exist"})
	}

	c.IndentedJSON(http.StatusOK, fruit)

}

func getFruitByID(id string) (*fruit, error) {
	for i, f := range fruits {
		if f.ID == id {
			return &fruits[i], nil
		}
	}

	return nil, errors.New("fruit not in system")

}

func addFruit(c *gin.Context) {

	var newFruit fruit

	if err := c.BindJSON(&newFruit); err != nil {
		return
	}

	fruits = append(fruits, newFruit)
	c.IndentedJSON(http.StatusCreated, newFruit)
}

func main() {

	router := gin.Default()

	router.GET("/fruits", getFruits)
	router.POST("/fruits/add", addFruit)
	router.GET("/fruits/:id", fruitById)
	router.PATCH("/sell", sellFruit)
	router.Run("localhost:8080")

}
