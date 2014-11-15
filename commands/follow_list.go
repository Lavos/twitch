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

	names, err := tc.Follows()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(strings.Join(names, " "))
}
