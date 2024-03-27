package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
	"strings"
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
			_, err := edgeStmt.Exec(nodeID, nodeClassID, "", edge.DestinationNodeID, edge.DestinationNodeClassID, "", edge.Relationship)
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

func SelectNodeGraph(db *sql.DB) (graph definition.Graph, err error) {
	nodeRows, err := db.Query(selectNodeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node definition.GraphNode
		var nodeID, classID string
		if err = nodeRows.Scan(&nodeID, &classID); err != nil {
			return
		}
		node.ID = definition.GraphNodeID(classID, nodeID)
		node.Class = classID
		node.Description = node.ID
		graph.Nodes = append(graph.Nodes, node)
	}

	linkRows, err := db.Query(selectNodeEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeEdgeSelectStmt)
		return
	}
	defer linkRows.Close()

	for linkRows.Next() {
		var link definition.GraphLink
		var sourceNodeID, sourceNodeClassID, destinationNodeID, destinationNodeClassID string
		if err = linkRows.Scan(&sourceNodeID, &sourceNodeClassID, &destinationNodeID, &destinationNodeClassID, &link.Relationship); err != nil {
			return
		}
		link.Source = definition.GraphNodeID(sourceNodeClassID, sourceNodeID)
		link.Target = definition.GraphNodeID(destinationNodeClassID, destinationNodeID)

		graph.Links = append(graph.Links, link)
	}

	return
}
