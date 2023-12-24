### Introduce
* Using golang echo as backend service.
* Using MongoDB as database.
* Using docker to setup MongoDB


### Run mongoDB
```bash
docker run -d \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=$MONGO_INITDB_ROOT_USERNAME \
    -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_INITDB_ROOT_PASSWORD \
    --name medical_record \
    --rm \
    mongo:latest
```
### Prepare Data
```bash
use db
db.patients.createIndex( { "id": 1 }, { unique: true } )
db.patients.insertMany(
    [{
        "id":1,
        "name":"Wang one",
        "orderId":"ZZpxtMUp",
        "gender":"M",
        "illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":2,
        "name":"lin two",
        "orderId":"3utEFDDr",
        "gender":"M",
        "illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":3,
        "name":"shu three",
        "orderId":"a4QQ8Adn",
        "gender":"F",
        "illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":4,
        "name":"Wang four",
        "orderId":"QyaShY6G",
        "gender":"M",
        "illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":5,
        "name":"yung five",
        "orderId":"Th6mxx4X",
        "gender":"F",
        "illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    }])

db.patients.find().pretty()
```
### Prepare Data
```bash
go run main.go
```
