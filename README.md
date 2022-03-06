# very small account app

#### I showed you how to use mongodb easily in this repo

### Dependencies
* Golang
* MongoDb
* Docker
* Gin web framework


#### First we launch mongodb with docker
```bash
  docker run -d -p 2717:27017 -v d:/mongodata:/data/db --name mymongo mongo
```

### create account

```bash
  curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "baran can", "status": true, "balance":100}' \
    http://localhost:8080/account
```
#### output

```json
{
    "message": "ok"
}
```

### Get account info

```bash
  curl http://localhost:8080/account/62250425f537d1628ef6fb6f
```
#### output

```json
{
    "data": [
        {
            "_id": "62250425f537d1628ef6fb6f",
            "balance": 50,
            "name": "baran",
            "status": true
        }
    ],
    "message": "ok"
}
```

### Add money

```bash
    curl -X PUT -H "Content-Type: application/json" \
    -d '{"id":"62250425f537d1628ef6fb6f", "balance":100}' \
    http://localhost:8080/add/money
```
#### output

```json
{
    "message": "ok"
}
```

### Reduce money

```bash
    curl -X PUT -H "Content-Type: application/json" \
    -d '{"id":"62250425f537d1628ef6fb6f", "balance":100}' \
    http://localhost:8080/reduce/money
```
#### output

```json
{
    "message": "ok"
}
```
