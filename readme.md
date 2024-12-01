### **Auth Service**

Auth Service is a microservice for user authentication using JWT tokens, PostgreSQL, and Memcached.

---

### **Features**

1. **User Registration** (`POST /register`):
   Registers a new user in the system.

2. **Authentication and Token Generation** (`POST /login`):
   Authenticates the user and returns a JWT token.

3. **Token Validation** (`GET /validate-token`):
   Validates the token using Memcached and the database as a fallback.

---

### **Technologies**

- **Go**: Programming language.
- **PostgreSQL**: Database for storing user information.
- **Memcached**: Caching layer for token validation.
- **Goose**: Database migration tool.
- **Gin**: Web framework for building REST APIs.

---

### **Installation and Setup**

#### **Steps:**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/auth-service.git
   cd auth-service
   ```

2. Ensure you have Go, Docker, and Docker Compose installed.

3. Start the services (PostgreSQL, Memcached, and the application):
   ```bash
   docker-compose up --build
   ```

4. Run database migrations:
   ```bash
   make migrate-up
   ```

5. The service will be available at:
   ```
   http://localhost:8080
   ```

---

### **API Examples**

#### **Register a User**
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "testuser",
        "password": "testpassword"
      }'
```

#### **Authenticate and Get a Token**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
        "username": "testuser",
        "password": "testpassword"
      }'
```

#### **Validate a Token**
```bash
curl -X GET http://localhost:8080/validate-token \
  -H "Authorization: Bearer <your JWT token>"
```

---

### **Makefile Commands**

- **Build the application**: `make build`
- **Run the application**: `make run`
- **Format the code**: `make fmt`
- **Lint the code**: `make lint`
- **Run migrations**: `make migrate-up`
- **Rollback migrations**: `make migrate-down`

---

### **Environment Variables**

The following environment variables can be used for configuration:

| Variable           | Description                     | Default                  |
|--------------------|---------------------------------|--------------------------|
| `DATABASE_URL`     | PostgreSQL connection string   | `postgres://user:password@localhost:5432/auth_db?sslmode=disable` |
| `MEMCACHED_HOST`   | Memcached host                 | `localhost`              |
| `MEMCACHED_PORT`   | Memcached port                 | `11211`                  |

---

### **Directory Structure**

```
auth-service/
├── cmd/
│   └── auth-service/      # Application entry point
├── internal/
│   ├── controller/        # API controllers
│   ├── model/             # Database models
│   ├── repository/        # Database access layer
│   └── service/           # Business logic
├── config/                # Configuration files
├── db/
│   └── migrations/        # Database migrations
├── Dockerfile             # Docker build configuration
├── docker-compose.yml     # Docker Compose configuration
├── Makefile               # Automation tasks
└── README.md              # Project documentation
```

---