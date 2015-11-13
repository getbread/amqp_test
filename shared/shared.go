package shared

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envFilePath string

func WorkingDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func LoadEnvironment() {
	// load .env
	flag.StringVar(&envFilePath, "env_path", WorkingDirectory()+"/.env", "Sets path for .env config")
	flag.Parse()
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}
