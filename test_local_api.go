package main

import (
	"./packngo"
	"github.com/hashicorp/go-cleanhttp"
	"log"
	"fmt"
	"reflect"
	"os"
)

const (
	// Taken from packet terraform project
	consumerToken = "aZ9GmqHTPtxevvFq9SK3Pi2yr9YCbRzduCSXF2SNem5sjB91mDq7Th3ZwTtRqMWZ"
)

func main() {
	httpClient := cleanhttp.DefaultClient()
	apiKey, set := os.LookupEnv("PACKET_API_KEY")
	if !set {
		fmt.Println("PACKET_API_KEY environment variable not set, exiting")
		return
	}
	client := packngo.NewClient(consumerToken, apiKey, httpClient)

	projectID, set := os.LookupEnv("PACKET_PROJECT_ID")
	if !set {
		fmt.Println("PACKET_PROJECT_ID environment variable not set, exiting")
		return
	}
	networks, _, err := client.ProjectVirtualNetworks.List(projectID)
	if err != nil {
		log.Fatal("Nah bad stuff happened")
	}

	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(networks), networks)
}