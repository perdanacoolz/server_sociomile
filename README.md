Cara Menjalankan Aplikasi

Backend:
Set env: export DB_DSN="user:pass@tcp(localhost:3306)/ticketing_db" JWT_SECRET="secret" REDIS_ADDR="localhost:6379"
Run migrations: mysql < migrations/schema.sql
go run cmd/main.go
Swagger: http://localhost:8080/swagger/index.html

Dengan Docker: docker-compose up
Frontend: cd ticketing-frontend && npm run dev (runs on localhost:3000)

Environment Variables

DB_DSN: MySQL connection string
JWT_SECRET: JWT signing key
REDIS_ADDR: Redis address

Daftar Endpoint API

POST /login: Login
POST /tickets: Create ticket (admin/agent)
PUT /tickets/:id/assign: Assign agent
PUT /tickets/:id/status: Update status
GET /tickets: List tickets (?status=open&agent_id=1)
POST /conversations/:ticket_id: Send message
GET /conversations/:ticket_id: Get conversation
