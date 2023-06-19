package main

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	// Set up Twitter API credentials
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	// Authenticate with Twitter API
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	// Send a tweet
	tweet, err := api.PostTweet("Hello, Twitter from my Go bot!", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tweet sent successfully:", tweet.Text)
}
