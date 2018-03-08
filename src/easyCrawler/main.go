package main

import (
	"log"
	"os"
	_ "easyCrawler/matchers"
	"easyCrawler/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
