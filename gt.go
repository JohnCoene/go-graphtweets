package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Edges A struct of edges
type Edges struct {
	Source []string
	Target []string
}

// Nodes list of nodes
type Nodes struct {
	Name []string
}

// Graph an object containing nodes and edges
type Graph struct {
	Nodes Nodes
	Edges Edges
}

// GetNodes Get the nodes from the edges
func GetNodes(edges Edges) Nodes {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range edges.Source {
		encountered[edges.Source[v]] = true
	}

	for v := range edges.Target {
		encountered[edges.Target[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	nodes := Nodes{result}
	return (nodes)
}

func main() {
	config := oauth1.NewConfig("niLRMTSNzOw7Fqp1WNbdkdln8", "yFFYvgW8kJguFBtXmaR602Hyo1RDDRBeetZo7nIo5JecAAC4aA")
	token := oauth1.NewToken("1571967648-u5ezlrsWplrHaXT4VMuVLpES2ZB8kx6UIRH4Obj", "admb6ncPJJOaCwh7HR4YASQdSJ2H9B9dIAL50yzmOHsvL")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search Tweets
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "rstats",
	})

	if err != nil {
		log.Fatal(err)
	}

	fr := make([]string, 0)
	to := make([]string, 0)

	for _, v := range search.Statuses {
		userMention := v.Entities.UserMentions

		for _, v := range userMention {
			fr = append(fr, v.ScreenName)
		}

		for i := 0; i < len(userMention); i++ {
			to = append(to, v.User.ScreenName)
		}

	}

	edges := Edges{fr, to}
	nodes := GetNodes(edges)

	graph := Graph{nodes, edges}

	fmt.Println(graph)

}
