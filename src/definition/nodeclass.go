package definition

type (
	// NodeClassAttributes TODO
	NodeClassAttributes map[string]NodeClassAttribute

	// NodeClassAttribute TODO
	NodeClassAttribute struct {
		Description string `yaml:"Description,omitempty"`
		Type        string `yaml:"Type"`
		IsRequired  int    `yaml:"IsRequired"`
	}

	// NodeClassEdges TODO
	NodeClassEdges []NodeClassEdge

	// NodeClassEdge TODO
	NodeClassEdge struct {
		DestinationNodeClassID string `yaml:"DestinationNodeClass"`
		Relationship           string `yaml:"Relationship"`
		IsToDestination        int    `yaml:"IsToDestination"`
		IsFromDestination      int    `yaml:"IsFromDestination"`
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
