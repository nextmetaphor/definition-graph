CREATE TABLE NodeClass (
    ID          TEXT not null primary key,
    Description TEXT
);

create table NodeClassAttribute (
    ID          TEXT    not null,
    ClassID     TEXT    not null references NodeClass,
    Type        TEXT    not null,
    IsRequired  INTEGER not null,
    Description TEXT,
    primary key (ID, ClassID)
);

create table NodeClassEdge (
    SourceClassID      TEXT    not null references NodeClass,
    DestinationClassID TEXT    not null references NodeClass,
    Relationship       TEXT    not null,
    IsFromSource       INTEGER not null,
    IsFromDestination  INTEGER not null,
    primary key (SourceClassID, DestinationClassID, Relationship)
);

create table Node (
    ID      TEXT not null,
    ClassID TEXT references NodeClass,
    primary key (ID, ClassID)
);

create table NodeAttribute (
    NodeID               TEXT not null,
    ClassID              TEXT not null,
    NodeClassAttributeID TEXT not null,
    Value                TEXT not null,
    primary key (Nodeid, ClassID, NodeClassAttributeID),
    foreign key (Nodeid, ClassID) references node,
    foreign key (NodeClassAttributeID, ClassID) references NodeClassAttribute
);

