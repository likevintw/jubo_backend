
### Run mongoDB
```
docker run -d \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=$MONGO_INITDB_ROOT_USERNAME \
    -e MONGO_INITDB_ROOT_PASSWORD=$MONGO_INITDB_ROOT_PASSWORD \
    --name medical_record \
    --rm \
    mongo:latest
```
### Prepare Data
```
use patients_dialog
db.patients.insertMany(
    [{
        "id":"1",
        "name":"Wang one",
        "orderId":"ZZpxtMUp",
        "gender":"M",
        "Illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":"2",
        "name":"lin two",
        "orderId":"3utEFDDr",
        "gender":"M",
        "Illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":"3",
        "name":"shu three",
        "orderId":"a4QQ8Adn",
        "gender":"F",
        "Illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":"4",
        "name":"Wang four",
        "orderId":"QyaShY6G",
        "gender":"M",
        "Illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    },
    {
        "id":"5",
        "name":"yung five",
        "orderId":"Th6mxx4X",
        "gender":"F",
        "Illness":"feel bad",
        "history":"feel very bad on 2023-03-13",
        "dialog":"",
    }]
    )
db.patients.find().pretty()
```