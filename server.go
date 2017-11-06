package main

import "github.com/gin-gonic/gin"
import "strconv"

type Sighting struct {
	Id          string `json:"id"`
	Species     string `json:"species"`
	Description string `json:"desc"`
	DateTime    string `json:"date"`
	Count       int    `json:"count"`
}

func main() {
	species := []string{
		"mallard",
		"redhead",
		"gadwall",
		"canvasback",
		"lesser scaup",
	}

	sightings := []Sighting{
		{
			Id:          "1",
			Species:     "gadwall",
			Description: "All your ducks are belong to us",
			DateTime:    "2016-10-01T01:01:00Z",
			Count:       1,
		},
		{
			Id:          "2",
			Species:     "lesser scaup",
			Description: "This is awesome",
			DateTime:    "2016-12-13T12:05:00Z",
			Count:       5,
		},
		{
			Id:          "3",
			Species:     "canvasback",
			Description: "...",
			DateTime:    "2016-11-30T23:59:00Z",
			Count:       2,
		},
		{
			Id:          "4",
			Species:     "mallard",
			Description: "Getting tired",
			DateTime:    "2016-11-29T00:00:00Z",
			Count:       18,
		},
		{
			Id:          "5",
			Species:     "redhead",
			Description: "I think this one is called Alfred J.",
			DateTime:    "2016-11-29T10:00:01Z",
			Count:       1,
		},
		{
			Id:          "6",
			Species:     "redhead",
			Description: "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.",
			DateTime:    "2016-12-01T13:59:00Z",
			Count:       1,
		},
		{
			Id:          "7",
			Species:     "mallard",
			Description: "Too many ducks to be counted",
			DateTime:    "2016-12-12T12:12:12Z",
			Count:       100,
		},
		{
			Id:          "8",
			Species:     "canvasback",
			Description: "KWAAK!!!1",
			DateTime:    "2016-12-11T01:01:00Z",
			Count:       5,
		},
	}

	router := gin.Default()
	router.GET("/sightings", func(c *gin.Context) {
		c.JSON(200, sightings)
	})

	router.POST("/sightings", func(c *gin.Context) {
		var sighting Sighting
		c.BindJSON(&sighting)
		sighting.Id = strconv.Itoa(len(sightings) + 1)
		sightings = append(sightings, sighting)
		c.JSON(200, sighting)
	})

	router.GET("/species", func(c *gin.Context) {
		c.JSON(200, species)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
