# crudLight

## Start server
```
docker build -t crudlight . && docker run --rm -p 8080:8080 crudlight
```

* ### Create user
```
curl --location --request POST 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Nemo",
    "birth_date": "2000-01-01"
}'
```

* ### Get user
```
curl --location --request GET 'http://localhost:8080/user/1'
```

* ### Update user
```
curl --location --request PATCH 'http://localhost:8080/user/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Max"
}'
```

* ### Delete user
```
curl --location --request DELETE 'http://localhost:8080/user/1'
```
