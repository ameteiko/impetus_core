# MindMap

MindMap stands for creating a model for a particular field of study.

**Technical stack:**
* `Neo4J` database as a persistence layer.
* `Golang` as a backend engine.
* `MermaidJS` as a diagram-building tool. 

## Launching the application

### Create a docker volume to store the data: 
```bash
export BASE_PATH=$(pwd)
export PRJ_NAME=sandbox
mkdir -p ${BASE_PATH}/.data/${PRJ_NAME}/volume
docker volume create                                  \
--name ${PRJ_NAME}_mindnet                            \
--opt type=none                                       \
--opt device=${BASE_PATH}/.data/${PRJ_NAME}/volume    \
--opt o=bind
```

### Preconfigure the database container
```bash
export BASE_PATH=$(pwd)
export PRJ_NAME=sandbox
export BOLT_CERTS=${BASE_PATH}/.data/${PRJ_NAME}/neo4j_certs/bolt
mkdir -p ${BOLT_CERTS}/trusted
mkdir -p ${BOLT_CERTS}/revoked
cd ${BOLT_CERTS}
openssl req -x509 -newkey rsa:4096 -keyout private.key \
  -out public.crt -days 36500 -subj '/CN=localhost' -nodes
chmod 400 private.key
```

### Launch the database container
```bash
export BASE_PATH=$(pwd)
export PRJ_NAME=sandbox
export BOLT_CERTS=${BASE_PATH}/.data/${PRJ_NAME}/neo4j_certs/bolt
docker run                                                   \
  --publish=7474:7474                                        \
  --publish=7687:7687                                        \
  --volume=${PRJ_NAME}_mindnet:/data                         \
  --volume=${BOLT_CERTS}:/var/lib/neo4j/certificates/bolt    \
  --env=NEO4J_dbms_connector_bolt_tls__level=OPTIONAL        \
  --env NEO4J_AUTH=none                                      \
  --env=NEO4J_dbms_ssl_policy_bolt_enabled=TRUE              \
  neo4j
```

### Login to the database

Visit http://localhost:7474