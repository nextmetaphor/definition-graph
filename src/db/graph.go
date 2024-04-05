package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

const (
	selectGraphNodeEdgeSQL = `SELECT SourceNodeID, SourceNodeClassID, SourceNodeClassNamespace, DestinationNodeID, DestinationNodeClassID, DestinationNodeClassNamespace, Relationship from NodeEdge;`
)

func SelectNodeClassGraph(db *sql.DB) (graph definition.Graph, err error) {
	nodeRows, err := db.Query(selectNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassSelectStmt)
		return
	}
	defer nodeRows.Close()

	for nodeRows.Next() {
		var node definition.GraphNode
		if err = nodeRows.Scan(&node.ID, &node.Namespace, &node.Description); err != nil {
			return
		}
		node.Class = node.ID
		graph.Nodes = append(graph.Nodes, node)
	}

	linkRows, err := db.Query(selectNodeClassEdgeSQL)

	if err != nil {
		log.Error().Err(err).Msg(logCannotQueryNodeClassEdgeSelectStmt)
		return
	}
	defer linkRows.Close()

	for linkRows.Next() {
		var link definition.GraphLink
		if err = linkRows.Scan(&link.Source, &link.Target, &link.Relationship); err != nil {
			return
		}
		graph.Links = append(graph.Links, link)
	}

	return
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

	linkRows, err := db.Query(selectGraphNodeEdgeSQL)
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
