
### Run mongoDB
```
docker run -d \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=dw2nl3 \
    --name medical_record \
    --rm \
    mongo:latest
```

