package main

import (
	"context"
	"os"
	"strings"

	"SocialNetworkBackend/awsgo"
	"SocialNetworkBackend/db"
	"SocialNetworkBackend/handlers"
	"SocialNetworkBackend/models"
	"SocialNetworkBackend/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)
func main() {
	lambda.Start(ExecLambda)
}

func ExecLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	
	var res *events.APIGatewayProxyResponse

	awsgo.InitAWS()
	
	if !ValidParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Missing parameters",
			Headers:   map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error getting secret",
			Headers:   map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["socialnetworkgo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))	

	// Check Connection to the database

	err = db.ConnectDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connecting to the database",
			Headers:   map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.Handlers(awsgo.Ctx, request)
	if respAPI.CustomResp != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers:    map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil

	} else {
			return respAPI.CustomResp, nil
	}

}

func ValidParams() bool {
	_, getParameter := os.LookupEnv("SecretName")

	if !getParameter {
		return false
	}

	_, getParameter = os.LookupEnv("BucketName")

	if !getParameter {
		return false
	}

	_, getParameter = os.LookupEnv("UrlPrefix")

	if !getParameter {
		return false
	}

	return true

}


