# Ports service

This is a simple service to manage ports using clean architecture.

## How to run using postgres

```bash
ports server -p 8080 -r postgres -d postgres://postgres:postgres@localhost:5432/ports?sslmode=disable
```

## How to run using memory

```bash
ports -p 8080 -r memory 
```


## How to run in docker

```bash
docker-compose up
```

## Endpoints

### PUT /ports

```bash
curl -i -X PUT -d @./testdata/ports.json http://localhost:8080/ports
```