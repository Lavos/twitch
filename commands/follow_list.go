package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/Lavos/twitch"
	"flag"
)

var (
	newlines = flag.Bool("newlines", false, "Separate entries with newlines instead of spaces.")
)

func main () {
	flag.Parse()
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

	if len(names) > 0 {
		var spacer = " "

		if (*newlines) {
			spacer = "\n"
		}

		fmt.Print(strings.Join(names, spacer))
	}
}
