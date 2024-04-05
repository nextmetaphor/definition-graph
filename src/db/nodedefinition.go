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
			// TODO needs namespace
			_, err := attributeStmt.Exec(nodeID, nodeClassID, "default", attributeID, attribute)
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
			//TODO needs namespace
			_, err := edgeStmt.Exec(nodeID, nodeClassID, "default", edge.DestinationNodeID, edge.DestinationNodeClassID, "default", edge.Relationship)
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
