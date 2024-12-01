curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "testuser",
        "password": "testpassword"
      }'


curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
        "username": "testuser",
        "password": "testpassword"
      }'

curl -X GET http://localhost:8080/validate-token \
  -H "Authorization: Bearer <ваш JWT токен>"
