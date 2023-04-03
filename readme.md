# Intro

# Tech Stack

# Functionality

# Running Locally

## Postgres Docker Image

With Docker installed, run:

`docker pull postgres`

To run a local instance of postgres, run but configure to your own desired credentials and DB name:

```

docker run ^
    --name myPostgresDb ^
    -p 5455:5432 ^
    -e POSTGRES_USER=postgresUser ^
    -e POSTGRES_PASSWORD=postgresPW ^
    -e POSTGRES_DB=postgresDB ^
    -d ^
    postgres
```

Once the instance is running, connect to it with any DB connection GUI (e.g. pgAdmin) or command line
