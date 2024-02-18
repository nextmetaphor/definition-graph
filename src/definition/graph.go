package definition

type (
	Graph struct {
		Nodes []GraphNode `json:"nodes"`
		Links []GraphLink `json:"links"`
	}

	GraphNode struct {
		ID          string `json:"id"`
		Class       string `json:"class"`
		Description string `json:"description"`
	}

	GraphLink struct {
		Source       string `json:"source"`
		Target       string `json:"target"`
		Relationship string `json:"relationship"`
	}
)
