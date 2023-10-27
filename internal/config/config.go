package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const projectDirName = "yakshop" // change to relevant project name

var Cfg config

func init() {
	loadConfig()
}

type config struct {
	PORT string `envconfig:"PORT" required:"true"`
}

func loadConfig() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalln(err)
	}
	err = envconfig.Process("", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
