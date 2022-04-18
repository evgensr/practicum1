#!/bin/bash

export SERVER_ADDRESS=0.0.0.0:8080
export BASE_URL=http://localhost:8080/
export FILE_STORAGE_PATH=123.txt



curl -i -d "{\"url\": \"https://habr8.ru\"}" -X POST http://localhost:8080/api/shorten


curl -i -d "https://habr3.ru" -X POST http://localhost:8080/

curl -i http://localhost:8080/e2828d79249d0279acbb78e0ce8072ce

curl -i http://localhost:8080/5792bb5d149a507c1ac10358fd67cccf
