package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	fmt.Println("initializing...")

	mongoConnURL, dbName, _ := LoadConfig()
	wordCollection := Connect(mongoConnURL, dbName)

	Init(wordCollection)
}

func Init(cc *mongo.Collection) {

	gin.SetMode(gin.ReleaseMode)
	ginn := gin.New()

	ginn.Use(func(c *gin.Context) {
		fmt.Println("middleware")
		c.Next()
	})

	group := ginn.Group("/word")

	group.POST("/create", CreateWord(cc))
	group.POST("/correct", CorrectWord(cc))
	group.POST("/edit", EditWord(cc))
	group.POST("/delete", DeleteWord(cc))
	log.Fatal(ginn.Run(":3000"))
}

func LoadConfig() (string, string, string) {

	fmt.Println("loading config...")

	configFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal("error reading config file", err)
	}

	var config ConfigType

	json.Unmarshal(configFile, &config)

	return config.MongoConnURL, config.DbName, config.Version
}

type ConfigType struct {
	Version      string `json:"version"`
	MongoConnURL string `json:"mongoConnURL"`
	DbName       string `json:"dbName"`
}
