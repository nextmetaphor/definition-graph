package db

import (
	"database/sql"
	"github.com/nextmetaphor/definition-graph/definition"
	"github.com/rs/zerolog/log"
)

func StoreNodeClassSpecification(db *sql.DB, ncs *definition.NodeClassSpecification) error {
	stmt, err := db.Prepare(createNodeClassSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassStmt)
		return err
	}

	attributeStmt, err := db.Prepare(insertNodeClassAttributeSQL)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	edgeStmt, err := db.Prepare(insertNodeClassEdgeSQL)
	if err != nil {
		log.Error().Err(err).Msg(logCannotPrepareNodeClassEdgeStmt)
		return err
	}

	for classID, classDefinition := range ncs.Definitions {
		// create NodeClass record
		_, err := stmt.Exec(classID, "default", classDefinition.Description)
		if err != nil {
			log.Warn().Err(err).Msgf(logCannotExecuteNodeClassStmt, classID, classDefinition)
		}

		// create NodeClassAttribute records
		for attributeID, attribute := range classDefinition.Attributes {
			_, err := attributeStmt.Exec(attributeID, classID, "default", attribute.Description, attribute.Type, boolToInt(attribute.IsRequired))
			if err != nil {
				log.Warn().Err(err)
			}
		}

		// create NodeClassEdge records
		for _, edge := range classDefinition.Edges {
			// TODO - needs namespaces in definition files
			_, err := edgeStmt.Exec(classID, "default", edge.DestinationNodeClassID, "default", edge.Relationship)
			if err != nil {
				log.Warn().Err(err).Msgf(logCannotExecuteNodeClassEdgeStmt, classID, edge)
			}
			if edge.IsBidirectional {
				_, err := edgeStmt.Exec(edge.DestinationNodeClassID, classID, edge.Relationship)
				if err != nil {
					log.Warn().Err(err).Msgf(logCannotExecuteNodeClassEdgeStmt, classID, edge)
				}
			}
		}
	}

	return nil
}

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
