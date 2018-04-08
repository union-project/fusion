# Fusion
Fusion is the application responsible for synchronizing data from remote APIs into the Union Data Store.

# DB
Fusion uses [Cockroach DB](https://github.com/cockroachdb/cockroach) for the datastore.  For development,
it's easiest to use Docker for the local datastore:

```
$> docker run -ti -d --net=host --restart=always --name fusion-db cockroachdb/cockroach start --insecure
```

Once running, exec into the container and initialize the database:

```
$> docker exec -ti fusion-db /cockroach/cockroach sql --insecure -e 'CREATE DATABASE IF NOT EXISTS fusion;'
$> docker exec -ti fusion-db /cockroach/cockroach sql --insecure -e 'CREATE USER fusion;'
$> docker exec -ti fusion-db /cockroach/cockroach sql --insecure -e 'GRANT ALL ON DATABASE fusion to fusion;'
```

# Data
Fusion uses data sources for information.

## ProPublica
[ProPublica](https://www.propublica.org/datastore/apis) provides congressional data on members, bills etc.
You will need to sign up and get an API key to import data.  Once you have an API key you can sync the data:

```
$> export FUSION_PROPUBLICA_API_KEY=<api-key>
$> ./cmd/fusion/fusion sync
```

# Bulk Import
Fusion can bulk import data.  For example, to import the current year of ProPublic data, download the latest archive from [here](https://www.propublica.org/datastore/dataset/congressional-data-bulk-legislation-bills).  Then import the zip:

```
$>
```
