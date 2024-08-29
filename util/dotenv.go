package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadDotEnv(fpath string) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "=")
		os.Setenv(res[0], res[1])
	}

}
