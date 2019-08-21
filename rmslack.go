package main

import (
	"flag"
	"github.com/nlopes/slack"
	"log"
)

func main() {

	channel := flag.String("c", "", "-c slack channel name")
	token := flag.String("tk", "", "slack access channel")
	ts := flag.String("ts", "0", "timestamp of remove target comment")
	flag.StringVar(channel, "channel", "", "slack channel name")
	flag.StringVar(token, "token", "", "slack api access token")
	flag.StringVar(ts, "timestamp", "0", "timestamp of remove target comment")
	flag.Parse()

	if channel == nil {
		log.Fatal("channel is required")
	}

	if token == nil {
		log.Fatal("token is required")
	}

	if ts == nil {
		log.Fatal("timestamp is required")
	}

	_, _, err := slack.New(*token).DeleteMessage(*channel, *ts)
	if err != nil {
		log.Fatalf("faile to delete: %v", err)
	}
	log.Println("success")
}
