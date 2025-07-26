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

### Books

Authenticated users can manage books:

- `POST /api/books` – create a new book by providing a `title`.
- `POST /api/books/{id}/users/{userId}` – add a user to a book.
- `GET /api/books/{id}` – retrieve a book with its members.
