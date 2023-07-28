package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}
var awsOnce = sync.Once{}
var tableOnce = sync.Once{}
var saltOnce = sync.Once{}

type Table struct {
	UserTableName string `json:"USER_TABLE_NAME"`
}

type Application struct {
	Host string `json:"HOST"`
	Port string `json:"PORT"`
}

type Aws struct {
	AccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	Region          string `mapstructure:"AWS_REGION"`
}

type Salt struct {
	SecretKey int `mapstructure:"SECRET_SALT_KEY"`
}

var appConfig *Application
var awsConfig *Aws
var tableConfig *Table
var saltConfig *Salt

func loadApp() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	appConfig = &Application{
		Host: viper.GetString("HOST"),
		Port: viper.GetString("PORT"),
	}
}

func loadAws() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	awsConfig = &Aws{
		AccessKeyId:     viper.GetString("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		Region:          viper.GetString("AWS_REGION"),
	}
}

func loadTable() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	tableConfig = &Table{
		UserTableName: viper.GetString("USER_TABLE_NAME"),
	}
}

func loadSalt() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	saltConfig = &Salt{
		SecretKey: viper.GetInt("SECRET_SALT_KEY"),
	}
}

func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})
	return appConfig
}

func GetAws() *Aws {
	awsOnce.Do(func() {
		loadAws()
	})
	return awsConfig
}

func GetTable() *Table {
	tableOnce.Do(func() {
		loadTable()
	})
	return tableConfig
}

func GetSalt() *Salt {
	saltOnce.Do(func() {
		loadSalt()
	})
	return saltConfig
}
