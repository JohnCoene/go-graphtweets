package graphtweets

import (
	"github.com/dghubble/go-twitter/twitter"
)

// GetMentionsGraph constructs a graph of mentions
func GetMentionsGraph(search twitter.Search) Graph {
	edges := GetMentionEdges(search) // get edges
	nodes := GetNodes(edges)         // get nodes based on edges

	graph := Graph{nodes, edges}
	return (graph)
}

// GetRetweetGraph constructs a graph of retweets
func GetRetweetGraph(search twitter.Search) Graph {
	edges := GetRetweetEdges(search) // get edges
	nodes := GetNodes(edges)         // get nodes based on edges

	graph := Graph{nodes, edges}
	return (graph)
}
