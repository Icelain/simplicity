package env

import (
	"log"
	"os"
	"strconv"
)

func GetServerPort() uint {

	u, err := strconv.ParseUint(os.Getenv("PORT"), 10, 64)
	if err != nil {

		log.Fatal(err)

	}

	return uint(u)

}

func GetModel() string {

	return os.Getenv("MODEL")

}
