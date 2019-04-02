package graphtweets

// Edge A single edge
type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Number int    `json:"weight"`
}

// Node A single node
type Node struct {
	Name   string `json:"name"`
	Number int    `json:"size"`
}

// WideGraph constructs a wide graph
type WideGraph struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

// ToWide converts the default long Graph to a wide format
// more convenient for JSON marshalling
func ToWide(g *Graph) WideGraph {

	// instatiate
	var resultingGraph WideGraph
	resultingGraph.Edges = make([]Edge, len(g.Edges.Source))
	resultingGraph.Nodes = make([]Node, len(g.Nodes.Name))

	// add edges
	for i := range g.Edges.Source {

		e := Edge{
			Source: g.Edges.Source[i],
			Target: g.Edges.Target[i],
			Number: g.Edges.Number[i],
		}

		resultingGraph.Edges[i] = e

	}

	// add nodes
	for i, v := range g.Nodes.Name {
		n := Node{
			Name:   v,
			Number: g.Nodes.Number[i],
		}
		resultingGraph.Nodes[i] = n
	}

	return resultingGraph

}
