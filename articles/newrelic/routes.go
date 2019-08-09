package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	nr "github.com/newrelic/go-agent"
)

//setUpRoutes set all the routes
func setUpRoutes(r *gin.Engine) {
	r.GET("/app/status", getStatus)
}

func getStatus(c *gin.Context) {
	ctx := c.Request.Context()
	err := callGoogle(ctx)
	resp := make(map[string]string)
	if err != nil {
		resp["status"] = "ok"
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(400)
		jsonResp, _ := json.Marshal(resp)
		c.Writer.Write(jsonResp)
		return
	}
	doSomeThing(ctx)

	resp["status"] = "ok"
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	jsonResp, _ := json.Marshal(resp)
	c.Writer.Write(jsonResp)
}

func callGoogle(ctx context.Context) error {
	if t := ctx.Value(keyNrID); t != nil {
		txn := t.(nr.Transaction)
		defer nr.StartSegment(txn, "callGoogle").End()
	}
	resp, err := http.Get("http://google.com/")
	if err != nil {
		return fmt.Errorf("Some error occuerd %s", err.Error())
	}
	defer resp.Body.Close()
	return nil
}

func doSomeThing(ctx context.Context) {
	if t := ctx.Value(keyNrID); t != nil {
		txn := t.(nr.Transaction)
		defer nr.StartSegment(txn, "doSomeThing").End()
	}
}
