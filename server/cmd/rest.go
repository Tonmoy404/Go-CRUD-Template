package cmd

import (
	"github.com/Tonmoy404/project/config"
	"github.com/Tonmoy404/project/repo"
	"github.com/Tonmoy404/project/rest"
	"github.com/Tonmoy404/project/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func serveRest() {
	appConfig := config.GetApp()
	awsConfig := config.GetAws()
	tableConfig := config.GetTable()
	saltConfig := config.GetSalt()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsConfig.Region),
	})
	if err != nil {
		panic(err)
	}

	ddbClient := dynamodb.New(sess)

	userRepo := repo.NewUserRepo(ddbClient, tableConfig.UserTableName)
	svc := service.NewService(userRepo)
	server, err := rest.NewServer(appConfig, svc, saltConfig)
	if err != nil {
		panic("Server can not start")
	}

	server.Start()
}
