package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ArticleRetrieveResult struct {
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Items []*Article `json:"items"`
}

type Article struct {
	Title	string
	CreatedAt string `json:"created_at"`
}

func main() {

	filename, err := os.Open("example_entries.json")

	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var result ArticleRetrieveResult

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(result)
}
