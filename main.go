package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/ping", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	type EnergyPair struct {
		Energy uint   `json:"energy"`
		Profit uint   `json:"profit"`
	}

	r.GET("/ping2", func(context *gin.Context) {

		energyPair := EnergyPair{nil, nil}

		context.JSON(200, gin.H{
			"username": nil,
			"password": nil,
			"energy": energyPair,
		})
	})

	r.GET("/token", func(context *gin.Context) {

		resp, err := http.Get("http://localhost:8080/token")
		if err != nil {
			panic(err)

		}
		defer resp.Body.Close()
		s, err:=ioutil.ReadAll(resp.Body)
		fmt.Printf(string(s))

		context.JSON(200, gin.H{
			"token": s,
		})
	})

	_ = r.Run(":8081")
}
