package graphtweets

import (
	"github.com/dghubble/go-twitter/twitter"
)

// GetMentionsGraph Get a mentions graph
func GetMentionsGraph(search twitter.Search) Graph {
	edges := GetMentionEdges(search) // get edges
	nodes := GetNodes(edges)         // get nodes based on edges

	graph := Graph{nodes, edges}
	return (graph)
}

// GetRetweetGraph Get a retweet graph
func GetRetweetGraph(search twitter.Search) Graph {
	edges := GetRetweetEdges(search) // get edges
	nodes := GetNodes(edges)         // get nodes based on edges

	graph := Graph{nodes, edges}
	return (graph)
}
