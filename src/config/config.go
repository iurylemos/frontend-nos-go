package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// HashKey is used to authenticate cookie
// BlockKey is used to crypto the data of cookie

var (
	API_URL  = ""
	Port     = 0
	HashKey  []byte
	BlockKey []byte
)

func LoadEnviroment() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("APP_PORT"))

	if erro != nil {
		log.Fatal(erro)
	}

	API_URL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
