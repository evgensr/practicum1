#!/bin/bash

export SERVER_ADDRESS=0.0.0.0:8080
export BASE_URL=http://localhost:8080/
export FILE_STORAGE_PATH=123.txt
export DATABASE_DSN="host=localhost user=postgres password=postgres dbname=restapi sslmode=disable"



curl -i -d "{\"url\": \"https://habr8.ru\"}" -X POST http://localhost:8080/api/shorten


curl -i -d "https://habr3.ru" -X POST http://localhost:8080/

curl -i http://localhost:8080/e2828d79249d0279acbb78e0ce8072ce

curl -i http://localhost:8080/5792bb5d149a507c1ac10358fd67cccf



http -v  http://localhost:8080/e8c31fc8dfc6ab9feab44bdfe8b9e145

http -v --session=prac11 POST http://localhost:8080/api/shorten url='http://habr.ru/'
http -v --session=prac11  http://localhost:8080/api/user/urls

curl -i -d "http://ya.ru" -X POST http://localhost:8080/


curl -v -X POST http://localhost:8080/api/shorten/batch \
   -H 'Content-Type: application/json' \
   -d '[{"correlation_id":"my_login","original_url":"http://ya.ru"},{"correlation_id":"my_login2","original_url":"http://ya2.ru"}]'



curl -v -i -X POST http://localhost:8080/api/shorten/batch \
-H 'Content-Type: application/json' \
-d '[{"correlation_id":"6e11db6c-a119-4e8d-99c4-1bb5f11304c5","original_url":"http://mqwidmfgdvt7xj.net/efsgtpmzlvk2tg"},{"correlation_id":"27ccbed1-97b3-4d4e-ab3f-b1599d06c073","original_url":"http://ivcetm.com"}]'




http -v --session=prac11 POST http://localhost:8080/api/shorten url='http://habr.ru/'
http -v  http://localhost:8080/e8c31fc8dfc6ab9feab44bdfe8b9e145
http -v  http://localhost:8080/ping


http -v  http://localhost:8080/114f277c99eac452f0b44b552a154b4d

http -v  http://localhost:8080/e8c31fc8dfc6ab9feab44bdfe8b9e145


http -v http://localhost:8080/e8c31fc8dfc6ab9feab44bdfe8b9e145 \
     Accept: \
     Accept-Encoding: \
     Connection:

curl -i -v --compressed -d "{\"url\": \"https://habr8.ru\"}" -X POST http://localhost:8080/api/shorten

curl -i -v --compressed http://localhost:8080/ffefb7c03d4639c3314318571baf6b2a



curl -v -i -X   DELETE http://localhost:8080/api/user/urls \
--cookie "practicum=b5dbcaf5e822cff4cdd6b21d410c8314afb3a222cdb38b3a5caf9b2fb48a3c90fff895be71849581b5a41675c413e91558523a06b6c62467d81511be6b8eea99" \
-H 'Content-Type: application/json' \
-d '["e8c31fc8dfc6ab9feab44bdfe8b9e145","98981d87735b7f871c516eaf9b6bf461"]'



curl -v -i   http://localhost:8080/98981d87735b7f871c516eaf9b6bf461 \
--cookie "practicum=b5dbcaf5e822cff4cdd6b21d410c8314afb3a222cdb38b3a5caf9b2fb48a3c90fff895be71849581b5a41675c413e91558523a06b6c62467d81511be6b8eea99"



curl -v -i  -d "{\"url\": \"https://habr.ru\"}" -X POST http://localhost:8080/api/shorten \
--cookie "practicum=b5dbcaf5e822cff4cdd6b21d410c8314afb3a222cdb38b3a5caf9b2fb48a3c90fff895be71849581b5a41675c413e91558523a06b6c62467d81511be6b8eea99"


