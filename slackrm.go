package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	channel := flag.String("c", "", "-c slack channel name")
	token := flag.String("tk", "", "slack access channel")
	ts := flag.String("ts", "0", "timestamp of remove target comment")
	flag.StringVar(channel, "channel", "", "slack channel name")
	flag.StringVar(token, "token", "", "slack api access token")
	flag.StringVar(ts, "timestamp", "0", "timestamp of remove target comment")
	flag.Parse()

	if *channel == "" {
		panic("channel is required")
	}

	if *token == "" {
		panic("token is required")
	}

	if *ts == "" {
		panic("timestamp is required")
	}

	if _, _, err := slack.New(*token).DeleteMessage(*channel, *ts); err != nil {
		panic("failed to delete: " + err.Error())
	}
	fmt.Println("success")
}
