package graphtweets

import (
	"github.com/dghubble/go-twitter/twitter"
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

// GetMentionsGraph Get a mentions graph
func GetMentionsGraph(search twitter.Search) Graph {
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
	return (graph)
}
