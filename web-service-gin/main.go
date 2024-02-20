package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// chapeau represents a hat in Dofus.
type chapeau struct {
	ID           string `json:"id"`           // Chapeau ID
	ItemLevel    uint8  `json:"itemLevel"`    // Level of the hat (1-200)
	Name         string `json:"name"`         // Name of the hat
	Critical     uint8  `json:"critical"`     // Critical hit rate of the hat
	Intelligence uint   `json:"intelligence"` // Intelligence of the hat
	Agility      uint   `json:"agility"`      // Agility of the hat
	Strength     uint   `json:"strength"`     // Strength of the hat
	Chance       uint   `json:"chance"`       // Chance of the hat
	Range        uint8  `json:"range"`        // Range of the hat
	AP           byte   `json:"ap"`           // Action points of the hat
	MP           byte   `json:"mp"`           // Movement points of the hat
	Summons      uint8  `json:"summons"`      // Summons of the hat
	Vitalite     uint   `json:"vitalite"`     // Vitality of the hat
}

var chapeaux = []chapeau{
	{
		ID:           "1",
		ItemLevel:    200,
		Name:         "Máscara trithona",
		Critical:     4,
		Intelligence: 0,
		Agility:      0,
		Strength:     0,
		Chance:       0,
		Range:        0,
		AP:           0,
		MP:           0,
		Summons:      1,
		Vitalite:     400,
	},
	{
		ID:           "2",
		ItemLevel:    172,
		Name:         "Sombrero de buen aporte",
		Critical:     6,
		Intelligence: 0,
		Agility:      50,
		Strength:     50,
		Chance:       0,
		Range:        0,
		AP:           0,
		MP:           0,
		Summons:      1,
		Vitalite:     300,
	},
	{
		ID:           "3",
		ItemLevel:    196,
		Name:         "El Kim",
		Critical:     7,
		Intelligence: 0,
		Agility:      0,
		Strength:     80,
		Chance:       0,
		Range:        0,
		AP:           0,
		MP:           0,
		Summons:      1,
		Vitalite:     400,
	},
}

func main() {
	router := gin.Default()
	router.GET("/chapeaux", getChapeaux)
	router.GET("/chapeaux/:id", getChapeauByID)
	router.POST("/chapeaux", postChapeaux)
	router.Run("localhost:8080")
}

// get chapeaux in JSON format
func getChapeaux(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, chapeaux)
}

// postChapeaux adds a chapeau from JSON received in the request body.
func postChapeaux(c *gin.Context) {
	var newChapeau chapeau

	// Call BindJSON to bind the received JSON to
	// newChapeau.
	if err := c.BindJSON(&newChapeau); err != nil {
		return
	}

	// Add the new chapeau to the slice.
	chapeaux = append(chapeaux, newChapeau)
	c.IndentedJSON(http.StatusCreated, newChapeau)
}

// getChapeauByID locates the chapeau whose ID value matches the id
// parameter sent by the client, then returns that chapeau as a response.
func getChapeauByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of chapeaux, looking for
	// a chapeau whose ID value matches the parameter.
	for _, a := range chapeaux {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "chapeau not found"})
}
