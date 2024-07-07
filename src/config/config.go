package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is url from social network API
	APIURL = ""
	// APIPort is a port with the WEBAPP go listening
	APIPort = 0
	//HashKey is a hash
	HashKey []byte
	//BlockKey is a hash
	BlockKey []byte
)

// Load loading enviroment variables for .env file
func Load() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("WARNING: Port cannot be loaded from env file, port set to 3000")
		APIPort = 3000
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
