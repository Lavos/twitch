package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/Lavos/twitch"
)

func main () {
	username := os.Getenv("TWITCH_USERNAME")

	if username == "" {
		log.Fatalf("No username found in TWITCH_USERNAME.")
	}

	tc := &twitch.TwitchClient{username}

	channels, err := tc.Follows()

	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, len(channels))
	for i, c := range channels {
		names[i] = c.Name
	}

	fmt.Print(strings.Join(names, " "))
}
