package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	slackURL := flag.String("u", "", "-u delete target slack comment url")
	channel := flag.String("c", "", "-c slack channel name")
	token := flag.String("tk", "", "slack access channel")
	ts := flag.String("ts", "0", "timestamp of remove target comment")
	insecure := flag.Bool("insecure", false, "HTTP requests with InsecureSkipVerify transport")
	flag.StringVar(slackURL, "url", "", "url is delete target slack comment url")
	flag.StringVar(channel, "channel", "", "slack channel name")
	flag.StringVar(token, "token", os.Getenv("SLACK_API_TOKEN"), "slack api access token")
	flag.StringVar(ts, "timestamp", "0", "timestamp of remove target comment")
	flag.BoolVar(insecure, "InsecureSkipVerify", false, "HTTP requests with InsecureSkipVerify transport")
	flag.Parse()

	if *token == "" {
		panic("token is required")
	}

	if *slackURL != "" {
		split := strings.Split(*slackURL, "/")
		if len(split) != 6 {
			log.Fatal("url is must follow this format: https://<your slack domain>/archives/<channel name>/<comment id>")
		}
		channel = &split[4]
		ts = &split[5]

		// If thread comment then format like below
		// https://<domain>.slack.com/archives/<channel>/p1566545900001800?thread_ts=1566545900.001700&cid=GJJ9NJ1B8
		ts = &strings.Split(*ts, "?")[0]
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
		proxyURL, err := url.Parse(getEnvAny("HTTP_PROXY", "http_proxy"))
		if err != nil {
			log.Fatalf("HTTP_PROXY is invalid url")
		}

		// Under particular http proxy environment in enterprise then need to insecure option
		hc := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
			Proxy: http.ProxyURL(proxyURL),
		}}
		sc = slack.New(*token, slack.OptionHTTPClient(hc))
	}

	if _, _, err := sc.DeleteMessage(*channel, timestamp); err != nil {
		log.Fatal("failed to delete: " + err.Error())
	}
	fmt.Println("success")
}

func getEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}
