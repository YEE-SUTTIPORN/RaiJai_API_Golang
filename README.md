# RaiJai_API_Golang

This project provides a simple REST API built with Gin and GORM.

## Authentication

Send a POST request to `/api/login` with JSON:

```json
{
  "name": "your-username",
  "password": "your-password"
}
```

The response returns a token. Include it in subsequent requests:

```
Authorization: Bearer <token>
```

User registration remains open at `POST /api/users`.
