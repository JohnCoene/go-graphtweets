package graphtweets

import "github.com/dghubble/go-twitter/twitter"

// Edges a list of edges and their occurrences
type Edges struct {
	Source []string `json:"source"`
	Target []string `json:"target"`
	Number []int    `json:"weight"`
}

// Nodes a list of nodes and their occurrences
type Nodes struct {
	Name   []string `json:"name"`
	Number []int    `json:"size"`
}

// Graph a list containing nodes and edges
type Graph struct {
	Nodes Nodes
	Edges Edges
}

// GetNodes Get the nodes from the edges
func GetNodes(edges *Edges) Nodes {
	encountered := map[string]int{}

	// Create a map of all unique elements.
	for v := range edges.Source {
		encountered[edges.Source[v]] = encountered[edges.Source[v]] + 1
	}

	for v := range edges.Target {
		encountered[edges.Target[v]] = encountered[edges.Target[v]] + 1
	}

	// Place all keys from the map into a slice.
	result := []string{}
	count := []int{}
	for key, index := range encountered {
		result = append(result, key)
		count = append(count, index)
	}
	nodes := Nodes{result, count}

	return (nodes)
}

// GetMentionEdges construct edges of mentions
// where the source of the edge is the user tweeting
// and the target is the user(s) tagged in the tweet(s).
func GetMentionEdges(search twitter.Search) Edges {
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

	cn := make([]int, len(to))

	edges := Edges{fr, to, cn}
	countEdges(&edges)

	return (edges)
}

// GetRetweetEdges construct edges of retweets
// where the source of the edge is the user tweeting
// and the target is the user retweeted.
func GetRetweetEdges(search twitter.Search) Edges {
	fr := make([]string, 0)
	to := make([]string, 0)

	for _, v := range search.Statuses {
		to = append(to, v.InReplyToScreenName)
		fr = append(fr, v.User.ScreenName)
	}

	cn := make([]int, len(to))

	edges := Edges{fr, to, cn}
	countEdges(&edges)

	return (edges)
}

// GetHashtagEdges
// where the source of the edge is the user tweeting
// and the target is the hashtag used in the tweet.
func GetHashtagEdges(search twitter.Search) Edges {
	fr := make([]string, 0)
	to := make([]string, 0)

	for _, v := range search.Statuses {
		hashtag := v.Entities.Hashtags

		for _, v := range hashtag {
			fr = append(fr, v.Text)
		}

		for i := 0; i < len(hashtag); i++ {
			to = append(to, v.User.ScreenName)
		}

	}

	cn := make([]int, len(to))

	edges := Edges{fr, to, cn}
	countEdges(&edges)

	return (edges)
}

func countEdges(edges *Edges) {
	ed := make(map[string]string)
	occ := edges.Number

	for i := range edges.Source {
		if _, ok := ed[edges.Source[i]]; !ok {
			ed[edges.Source[i]] = edges.Target[i]
			occ[i] = occ[i] + 1
		} else {
			occ[i] = occ[i] + 1
		}
	}

	*edges = Edges{edges.Source, edges.Target, occ}
}
