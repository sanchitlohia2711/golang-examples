package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/app/status", getStatus)

	// Listen and Server in 0.0.0.0:8080
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
	}
	go s.ListenAndServe()
	//Wait for server to get started
	time.Sleep(time.Second * 2)
	go startLoadTest()

	fmt.Scanln()
}

func getStatus(c *gin.Context) {
	ctx := c.Request.Context()
	err := callGoogle(ctx)
	if err != nil {
		c.Writer.WriteHeader(400)
		return
	}

	c.Writer.WriteHeader(200)
}

func callGoogle(ctx context.Context) error {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		return fmt.Errorf("Some error occuerd %s", err.Error())
	}
	defer resp.Body.Close()
	return nil
}
