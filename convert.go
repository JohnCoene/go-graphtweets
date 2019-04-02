package graphtweets

// Edge A single edge
type Edge struct {
	Source string
	Target string
	Number int
}

// Node A single node
type Node struct {
	Name string
}

// WideGraph Constructs a wide graph
type WideGraph struct {
	Edges []Edge
	Nodes []Node
}

// ToWide Convert Graph to wide
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
			Name: v,
		}
		resultingGraph.Nodes[i] = n
	}

	return resultingGraph

}
