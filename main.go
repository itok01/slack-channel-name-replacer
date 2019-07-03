package main

import (
	"log"
	"os"
	"regexp"

	"github.com/nlopes/slack"
)

func main() {
	args := os.Args
	slackToken := args[1]
	replaceChannelRegexp := args[2]
	placeChannelString := args[3]

	api := slack.New(slackToken)

	channels, err := api.GetChannels(false)
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range channels {
		r := regexp.MustCompile(replaceChannelRegexp)
		if r.MatchString(channel.Name) {
			_, err := api.RenameChannel(channel.ID, r.ReplaceAllString(channel.Name, placeChannelString))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
