/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("loading .env file error")
	}

	return os.Getenv(key)
}
