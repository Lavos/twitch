package main

import (
	"os"
	"fmt"
	"log"
	"github.com/Lavos/twitch"
	"github.com/daviddengcn/go-colortext"
)

func main () {
	username := os.Getenv("TWITCH_USERNAME")

	if username == "" {
		log.Fatalf("No username found in TWITCH_USERNAME.")
	}

	tc := &twitch.TwitchClient{username}

	channels, err := tc.Online()

	if err != nil {
		log.Fatal(err)
	}

	ct.ChangeColor(ct.Yellow, true, ct.Black, false)
	fmt.Printf("Twitch Channels Online: %d\n", len(channels))

	for _, c := range channels {
		ct.ChangeColor(ct.White, true, ct.Black, false)
		fmt.Printf("%s", c.Name)
		ct.ChangeColor(ct.White, false, ct.Black, false)
		fmt.Printf(" - %s\n", c.Game)
		ct.ChangeColor(ct.Cyan, false, ct.Black, false)
		fmt.Printf("  %s\n", c.Status)
	}
}
