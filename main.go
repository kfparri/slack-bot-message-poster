package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"flag"
)

func main() {
	textPtr := flag.String("text", "Howdy!", "The message to be sent")

	channelPtr := flag.String("channel", "bot-testing", "The name of the channel to post to")

	flag.Parse()

	url := "https://slack.com/api/chat.postMessage"

	payload := strings.NewReader("token=" + os.Getenv("SLACK_TOKEN") +"&text=" + *textPtr + "&as_user=true&channel=" + *channelPtr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "d39b0862-0aa5-4546-910a-ab396017a0cc")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("An error occurred getting response")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}