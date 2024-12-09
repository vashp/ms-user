curl -X POST http://localhost:8181/users \
-H "Content-Type: application/json" \
-d '{
"email": "example@example.com"
}'

curl -X GET http://localhost:8181/users/1

curl -X PUT http://localhost:8181/users/1 \
-H "Content-Type: application/json" \
-d '{
"email": "new_email@example.com"
}'

curl -X DELETE http://localhost:8181/users/1