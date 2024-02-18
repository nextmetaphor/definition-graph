package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	insertNodeSQL = `INSERT INTO GraphNode (ID, NodeClassID) values (?, ?);`

	insertNodeAttributeSQL = `INSERT INTO NodeAttribute (NodeID, NodeClassID, NodeClassAttributeID, Value) values (?, ?, ?, ?);`

	insertNodeEdgeSQL = `INSERT INTO NodeEdge (SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship) values (?, ?, ?, ?, ?);`

	logCannotPrepareNodeStmt          = "cannot prepare GraphNode insert statement"
	logCannotPrepareNodeAttributeStmt = "cannot prepare NodeAttribute insert statement"
	logCannotPrepareNodeEdgeStmt      = "cannot prepare NodeEdge insert statement"
	logCannotExecuteNodeStmt          = "cannot execute GraphNode insert statement, id=[%s], [%#v]"
	logCannotExecuteNodeAttributeStmt = "cannot execute NodeAttribute insert statement, classid=[%s], id=[%s], [%#v]"
	logCannotExecuteNodeEdgeStmt      = "cannot execute NodeEdge insert statement, classid=[%s], [%#v]"
	logAboutToCreateNode              = "about to create GraphNode, id=[%s], [%#v]"
	logAboutToCreateNodeAttribute     = "about to create NodeAttribute, classid=[%s], id=[%s], [%#v]"
	logAboutToCreateNodeEdge          = "about to create NodeEdge, classid=[%s], nodeid=[%s], [%#v]"
)

func StoreNodeSpecificationWithoutEdges(db *sql.DB, ns *definition.NodeSpecification) error {
	nodeStmt, err := db.Prepare(insertNodeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeStmt)
		return err
	}

	attributeStmt, err := db.Prepare(insertNodeAttributeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeAttributeStmt)
		return err
	}

	specClassID := strings.TrimSpace(ns.ClassID)
	for nodeID, node := range ns.Definitions {
		nodeClassID := strings.TrimSpace(node.ClassID)
		if nodeClassID == "" {
			nodeClassID = specClassID
		}
		log.Debug().Msgf(logAboutToCreateNode, nodeID, node)
		_, err := nodeStmt.Exec(nodeID, nodeClassID)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeStmt, nodeID, node)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range node.Attributes {
			log.Debug().Msgf(logAboutToCreateNodeAttribute, nodeClassID, nodeID, attribute)
			_, err := attributeStmt.Exec(nodeID, nodeClassID, attributeID, attribute)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeAttributeStmt, attributeID, nodeID, attribute)
			}
		}
	}

	return nil
}

func StoreNodeSpecificationOnlyEdges(db *sql.DB, ns *definition.NodeSpecification) error {
	edgeStmt, err := db.Prepare(insertNodeEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeEdgeStmt)
		return err
	}

	specClassID := strings.TrimSpace(ns.ClassID)
	for nodeID, node := range ns.Definitions {
		nodeClassID := strings.TrimSpace(node.ClassID)
		if nodeClassID == "" {
			nodeClassID = specClassID
		}

		// create NodeClassEdge records
		for _, edge := range node.Edges {
			log.Debug().Msgf(logAboutToCreateNodeEdge, nodeClassID, nodeID, edge)
			_, err := edgeStmt.Exec(nodeID, nodeClassID, edge.DestinationNodeID, edge.DestinationNodeClassID, edge.Relationship)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeEdgeStmt, nodeClassID, edge)
			}
			if edge.IsBidirectional {
				_, err := edgeStmt.Exec(edge.DestinationNodeID, edge.DestinationNodeClassID, nodeID, nodeClassID, edge.Relationship)
				if err != nil {
					log.Warn().Err(err).Msgf(logCannotExecuteNodeEdgeStmt, nodeClassID, edge)
				}
			}
		}
	}

	return nil
}
