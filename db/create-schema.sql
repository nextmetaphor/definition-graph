CREATE TABLE NodeClass
(
    ID          TEXT not null primary key,
    Description TEXT
);

CREATE TABLE NodeClassAttribute
(
    ID          TEXT    not null,
    NodeClassID TEXT    not null references NodeClass,
    Type        TEXT    not null,
    IsRequired  INTEGER not null,
    Description TEXT,
    primary key (ID, NodeClassID)
);

CREATE TABLE NodeClassEdge
(
    SourceNodeClassID      TEXT    not null references NodeClass,
    DestinationNodeClassID TEXT    not null references NodeClass,
    Relationship           TEXT    not null,
    IsFromSource           INTEGER not null,
    IsFromDestination      INTEGER not null,
    primary key (SourceNodeClassID, DestinationNodeClassID, Relationship, IsFromSource, IsFromDestination)
);

CREATE TABLE Node
(
    ID          TEXT not null,
    NodeClassID TEXT references NodeClass,
    primary key (ID, NodeClassID)
);

CREATE TABLE NodeAttribute
(
    NodeID               TEXT not null,
    NodeClassID          TEXT not null,
    NodeClassAttributeID TEXT not null,
    Value                TEXT not null,
    primary key (NodeID, NodeClassID, NodeClassAttributeID),
    foreign key (NodeID, NodeClassID) references Node (ID, NodeClassID),
    foreign key (NodeClassAttributeID, NodeClassID) references NodeClassAttribute (ID, NodeClassID)
);

CREATE TABLE NodeEdge
(
    SourceNodeID           TEXT    not null,
    SourceNodeClassID      TEXT    not null references NodeClass,
    DestinationNodeID      TEXT    not null,
    DestinationNodeClassID TEXT    not null references NodeClass,
    Relationship           TEXT    not null,
    IsFromSource           INTEGER not null,
    IsFromDestination      INTEGER not null,
    primary key (SourceNodeID, SourceNodeClassID, DestinationNodeID, DestinationNodeClassID, Relationship, IsFromSource,
                 IsFromDestination),
    foreign key (SourceNodeID, SourceNodeClassID) references Node (ID, NodeClassID),
    foreign key (DestinationNodeID, DestinationNodeClassID) references Node (ID, NodeClassID),
    foreign key (SourceNodeClassID, DestinationNodeClassID, Relationship, IsFromSource,
                 IsFromDestination) references NodeClassEdge (SourceNodeClassID, DestinationNodeClassID, Relationship,
                                                              IsFromSource, IsFromDestination)
)