# task-management-be

To run demo with docker

### Export env
```
export ENV=local
```

### Set secret environment variable
Go to the file `/secrets/local.env`, define `ADMIN_ACCESS_TOKEN` to configure administration access before running.

### Build docker
```
docker buildx bake
```

### Run docker compose
```
docker compose up -d
```

### Use swagger
Go to [api-docs](http://localhost:3001) to try the APIs.
