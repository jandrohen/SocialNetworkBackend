package handlers

import (
	"context"
	"fmt"

	"WebstormProjects/UDEMY/GO/SocialNetworkBackend/jwt"
	"WebstormProjects/UDEMY/GO/SocialNetworkBackend/models"

	"github.com/aws/aws-lambda-go/events"
)
func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Hello from the handler" + ctx.Value(models.Key("path")).(string) + " " + ctx.Value(models.Key("method")).(string))
	
	var res models.RespAPI
	res.Status = 400

	isOk, statusCode, msg, claim := validAuth(ctx, request)
	if !isOk {
		res.Status = statusCode
		res.Message = msg
		return res
	}

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

func validAuth(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "OK", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token)==0 {
		return false, 400, "Token not found", models.Claim{}
	}

	claim, isOk, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !isOk {
		if err != nil {
			fmt.Println("Error in ProcessToken: " + err.Error())
			return false, 401, "Error in ProcessToken", models.Claim{}
		} else {
			fmt.Println("Token not valid: " + msg)
			return false, 401, "Token not valid", models.Claim{}
		}

	}

	fmt.Println("Token valid")
	return true, 200, msg, *claim
}
