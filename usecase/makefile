# Variables
SERVICE_URL := http://localhost:8080
USERNAME := testuser
PASSWORD := testpassword
TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzMwNDYwNTQsInVzZXJfaWQiOjF9.ue2zG69qnmpYuTVKe4sUhFTDn9bPrCqzJUwrwuURXZ0"

# Tasks
.PHONY: register login validate-token

register:
	@echo "Registering user..."
	@curl -X POST $(SERVICE_URL)/register \
		-H "Content-Type: application/json" \
		-d '{"username": "$(USERNAME)", "password": "$(PASSWORD)"}'

login:
	@echo "Logging in and retrieving token..."
	@TOKEN=$$(curl -s -X POST $(SERVICE_URL)/login \
		-H "Content-Type: application/json" \
		-d '{"username": "$(USERNAME)", "password": "$(PASSWORD)"}' | jq -r '.token'); \
		echo "TOKEN=$${TOKEN}" && echo $${TOKEN} > token.txt

validate-token:
	@echo "Validating token..."
	@TOKEN=$$(cat token.txt); \
		curl -X GET $(SERVICE_URL)/validate-token \
		-H "Authorization: Bearer $${TOKEN}"

