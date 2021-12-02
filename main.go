package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")

	// cluster, err := gocb.Connect(
	// 	"localhost",
	// 	gocb.ClusterOptions{
	// 		Username: "Administrator",
	// 		Password: "Password",
	// 	})
	// if err != nil {
	// 	panic(err)
	// }

	// bucket := cluster.Bucket("speedtest")

	// // We wait until the bucket is definitely connected and setup.
	// err = bucket.WaitUntilReady(5*time.Second, nil)
	// if err != nil {
	// 	panic(err)
	// }

	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hi how are you")
	})

	err := router.Run(":8081")
	if err != nil {
		panic(err.Error())
	}
}
