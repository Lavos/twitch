package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/Lavos/twitch"
	"flag"
	"github.com/kelseyhightower/envconfig"
)

var (
	c twitch.ClientConfiguration
	t *twitch.TwitchClient

	newlines = flag.Bool("newlines", false, "Separate entries with newlines instead of spaces.")
)

func main () {
	flag.Parse()
	var err error
	envconfig.MustProcess("TWITCH", &c)

	if c.UserID == 0 || c.ClientID == "" {
		fmt.Fprintf(os.Stderr, "Missing required ENV variables: %#v", c)
		os.Exit(1)
	}

	t = twitch.New(c)

	channels, err := t.Follows()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting follows: %#v", err)
		os.Exit(1)
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
