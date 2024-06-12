# Server App

### Prerequisites
- Database MySQL
- Golang 1.20
- Import Database from file attached in root directory
- When running the client, use the username and password that already exists in database (all passwords are 123456)
- You can also add new users by API Call POST /register with following reques json format
  `{"username":"username",
	"full_name":"full_name",
	"password":"password"}`
- Run the server: go run main.go
