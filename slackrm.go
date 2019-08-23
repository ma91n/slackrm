package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := flag.String("u", "", "-u delete target slack comment url")
	channel := flag.String("c", "", "-c slack channel name")
	token := flag.String("tk", "", "slack access channel")
	ts := flag.String("ts", "0", "timestamp of remove target comment")
	insecure := flag.Bool("insecure", false, "HTTP requests with InsecureSkipVerify transport")
	flag.StringVar(url, "url", "", "url is delete target slack comment url")
	flag.StringVar(channel, "channel", "", "slack channel name")
	flag.StringVar(token, "token", os.Getenv("SLACK_API_TOKEN"), "slack api access token")
	flag.StringVar(ts, "timestamp", "0", "timestamp of remove target comment")
	flag.BoolVar(insecure, "InsecureSkipVerify", false, "HTTP requests with InsecureSkipVerify transport")
	flag.Parse()

	if *token == "" {
		panic("token is required")
	}

	if *url != "" {
		split := strings.Split(*url, "/")
		if len(split) != 6 {
			log.Fatal("url is must follow this format: https://<your slack domain>/archives/<channel name>/<comment id>")
		}
		channel = &split[4]
		ts = &split[5]
	}

	if *channel == "" {
		log.Fatal("channel is required")
	}
	if *ts == "" {
		log.Fatal("timestamp is required")
	}

	timestamp := *ts
	if strings.HasPrefix(*ts, "p") && len(*ts) == 17 {
		timestamp = timestamp[1:11] + "." + timestamp[11:]
	}

	fmt.Println("timestamp", timestamp)

	sc := slack.New(*token)
	if *insecure {
		// Under particular http proxy environment in enterprise then need to insecure option
		hc := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
		sc = slack.New(*token, slack.OptionHTTPClient(hc))
	}

	if _, _, err := sc.DeleteMessage(*channel, timestamp); err != nil {
		log.Fatal("failed to delete: " + err.Error())
	}
	fmt.Println("success")
}
