package search

import (
	"os"
	"github.com/gin-gonic/gin/json"
)

const dataFile = "/root/github/go/src/easyCrawler/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}