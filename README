# Introduction

This project is created to fulfill Bareksa's backend engineer take home test requirements by Michael Koh

## Prerequisite

Simply run docker compose up command. Do note if you're using M1 macbook, don't forget to change the `platform` on the docker configuration file from `linux/x86_64` into `linux/amd64`.

Also you need to run the database migration manually. The migration files is located on `bareksa-take-home-test-michael-koh/database/migration/mysql`

```bash
docker-compose up
go build main.go
./main
```

## Get News API
```JSON
EXAMPLE ONLY
scheme : http
host: localhost
port: 8000
base URL: http://localhost:8000
``` 

```json
Method : GET
path : /news
query: status, topicSerials
example : http://localhost:8000/news?status=publish&topicSerials=TPC-S47HF0,TPC-FIT905
Response body:
{
    "data": [
        {
            "serial": "NWS-F64920",
            "topic": {
                "serial": "TPC-FIT905",
                "title": "Everything about crypto"
            },
            "status": "publish",
            "title": "Cardano in a nutshell",
            "authorName": "Michael Koh",
            "description": "desc : Isi berita cardano",
            "tag": null
        }
    ],
    "message": "success"
}
```

## Create News API
```json
Method: POST
path : /news
example : http://localhost:8000/news
Request body: 
{
    "serial": "NWS-TESTING1238",
    "topicSerial": "TPC-S47HF0",
    "status": "draft",
    "title": "Create news title",
    "authorName": "Bryan testing",
    "description": "penjelasan testing",
    "tags": [
        "investasi",
        "crypto"
    ]
}

Response: status code 201
```

## Update News API
```json
Method: PATCH
path : /news
Request body: 
{
    "serial": "NWS-TESTING1234", // to be updated news serial
    "topicSerial": "TPC-S47HF0",
    "status": "publish",
    "title": "ini hasil update",
    "authorName": "Ferdy Update",
    "description": "penjelasan setelah update",
    "tags": [
        "investasi",
        "crypto"
    ]
}
example : http://localhost:8000/news
Response: status code 200
```

## Delete News API
```json
Method: DELETE
path : /news/{newsSerial}
example : http://localhost:8000/news/NWS-TESTING1234
Response: status code 200
```

## Get Tags API
```json
Method : GET
path : /tags
query: page, size
example : http://localhost:8000/tags?page=1&size=6
Response body:
{
    "data": [
        {
            "name": "crypto"
        },
        {
            "name": "saham"
        },
        {
            "name": "reksadana"
        },
        {
            "name": "investasi"
        },
        {
            "name": "obligasi"
        },
        {
            "name": "equity"
        }
    ],
    "message": "success"
}
```

## Create Tags API
```json
Method : POST
path : /tags
example : http://localhost:8000/tags
Request body:
{
    "name": "surat-berharga-negara"
}
Response body: status code 201
```

## Update Tags API
```json
Method : PUT
path : /tags/{tagName}
example : http://localhost:8000/tags/investasi
Request body:
{
    "name": "reksadana"
}
Response body: status code 200
```

## Delete Tags API
```json
Method : DELETE
path : /tags/{tagName}
example : http://localhost:8000/tags/obligasi
Response body: status code 200
```