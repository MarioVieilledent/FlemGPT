package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startApi() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get/all", getAll)

	r.GET("/set/:variable/:value", set)

	r.GET("/prompt/:prompt", prompt)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func prompt(c *gin.Context) {
	prompt := c.Param("prompt")
	sentence := strings.Split(prompt, " ")
	result := continueText(sentence, nbWords)
	c.JSON(200, gin.H{
		"result": result,
	})
}

func getAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"depth":            depth,
		"nbWords":          nbWords,
		"randPickOverBest": randPickOverBest,
		"coefOccur":        coefOccur,
		"coefDistance":     coefDistance,
		"coefWordLength":   coefWordLength,
		"minOccurToAccept": minOccurToAccept,
	})
}

func set(c *gin.Context) {
	variable := c.Param("variable")
	valueStr := c.Param("value")
	var intVal int
	var floatVal float64
	if strings.Contains(valueStr, ".") {
		if s, err := strconv.ParseFloat(valueStr, 64); err == nil {
			floatVal = s
		} else {
			c.JSON(500, gin.H{
				"message": "Impossible to parse value given in query param.",
			})
		}
	} else {
		if v, err := strconv.Atoi(valueStr); err == nil {
			intVal = v
		} else {
			c.JSON(500, gin.H{
				"message": "Impossible to parse value given in query param.",
			})
		}
	}

	switch variable {
	case "depth":
		depth = intVal
		break
	case "nbWords":
		nbWords = intVal
		break
	case "randPickOverBest":
		randPickOverBest = intVal
		break
	case "coefOccur":
		coefOccur = floatVal
		break
	case "coefDistance":
		coefDistance = floatVal
		break
	case "coefWordLength":
		coefWordLength = floatVal
		break
	case "minOccurToAccept":
		minOccurToAccept = intVal
		break
	default:
		c.JSON(500, gin.H{
			"message": "Variable not valid.",
		})
	}

	c.JSON(200, gin.H{
		"message": "Ok",
	})
}
