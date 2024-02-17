package definition

type (
	// NodeClassAttributes TODO
	NodeClassAttributes map[string]NodeClassAttribute

	// NodeClassAttribute TODO
	NodeClassAttribute struct {
		Description string `yaml:"Description,omitempty"`
		Type        string `yaml:"Type"`
		IsRequired  bool   `yaml:"IsRequired"`
	}

	// NodeClassEdges TODO
	NodeClassEdges []NodeClassEdge

	// NodeClassEdge TODO
	NodeClassEdge struct {
		DestinationNodeClassID string `yaml:"DestinationNodeClass"`
		Relationship           string `yaml:"Relationship"`
		IsBidirectional        bool   `yaml:"IsBidirectional,omitempty"`
	}

	// NodeClassDefinition TODO
	NodeClassDefinition struct {
		// Description TODO
		Description string `yaml:"Description,omitempty"`

		// Attributes TODO
		Attributes NodeClassAttributes `yaml:"Attributes"`

		// Edges TODO
		Edges NodeClassEdges `yaml:"Edges"`
	}

	// NodeClassSpecification TODO
	NodeClassSpecification struct {
		// Definitions TODO
		Definitions map[string]NodeClassDefinition `yaml:"Definitions,omitempty"`
	}
)
