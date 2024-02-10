package handlers

import (
	"context"
	"fmt"

	"WebstormProjects/UDEMY/GO/SocialNetworkBackend/models"

	"github.com/aws/aws-lambda-go/events"
)
func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Hello from the handler" + ctx.Value(models.Key("path")).(string) + " " + ctx.Value(models.Key("method")).(string))
	
	var res models.RespAPI
	res.Status = 400

	switch ctx.Value(models.Key("path")).(string) {
		case "POST"	:
			switch ctx.Value(models.Key("path")).(string) {
			
			}
			//
		case "GET"	:
			switch ctx.Value(models.Key("path")).(string) {
			
			}
			//
		case "PUT"	:
			switch ctx.Value(models.Key("path")).(string) {
			
			}
			//
		case "DELETE"	:
			switch ctx.Value(models.Key("path")).(string) {
			
			}
			//
		}

	res.Message = "Path not found"
	return res

}