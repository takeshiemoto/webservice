# Webservice

## Application

```
go build
```

```
docker-compose up
```

Run

```
go run *.go
```

## Setup DB

Login docker container
```
docker exec -i -t webservice_postgres_1 bash
```

Login postgres and create role
```
su - postgress
psql
CREATE DATABASE gwp OWNER gwp;
```

Login user

```
psql --username=gwp --password --dbname=gwp
```