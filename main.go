package main

import (
	"context"
	"os"

	"WebstormProjects/UDEMY/GO/SocialNetworkBackend/awsgo"

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


