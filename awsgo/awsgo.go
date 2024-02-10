package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var Err error


func InitAWS() (aws.Config, error) {
	Ctx = context.TODO()
	Cfg, Err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("eu-south-2"))
	if Err != nil {
		panic("Error loading AWS configuration" + Err.Error())
	}

}