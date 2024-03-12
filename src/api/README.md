
### namespace API
```shell
# SELECT/SEARCH
GET /namespace

# READ
GET /namespace/{namespace-id}
```

### nodeclass API
```shell
# SELECT/SEARCH
GET /nodeclass
  # HEADERS
  namespace [OPTIONAL]

# READ
GET /nodeclass/{nodeclass-ID}
  # HEADERS
  namespace [MANDATORY]
```

### node API
```shell
# SELECT/SEARCH
GET /node
  # HEADERS
  namespace [OPTIONAL]
  nodeclass [OPTIONAL]

# READ
GET /node/{node-id}
  # HEADERS
  namespace [MANDATORY]
  nodeclass [MANDATORY]
```

### graph API
```shell
### global nodeclass graph
GET /graph/nodeclass

### global node graph
GET /graph/node

```