package graphtweets

import (
	"github.com/dghubble/go-twitter/twitter"
)

// GetMentionsGraph Get a mentions graph
func GetMentionsGraph(search twitter.Search) Graph {
	edges := GetEdges(search) // get edges
	nodes := GetNodes(edges)  // get nodes based on edges

	graph := Graph{nodes, edges}
	return (graph)
}
