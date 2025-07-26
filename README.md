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

- `GET /api/books` – list all books.
- `POST /api/books` – create a new book by providing a `title`.
- `GET /api/books/{id}` – retrieve a book with its members.
- `PUT /api/books/{id}` – update the book title.
- `DELETE /api/books/{id}` – delete a book.
- `POST /api/books/{id}/users/{userId}` – add a user to a book.

### Transactions

Transactions must specify which book they belong to:

- `POST /api/transactions` – create a transaction using `amount`, `note`, `date`, `user_id`, `book_id` and `category_id`.
- `GET /api/transactions` – list all transactions.
- `GET /api/transactions/{id}` – retrieve a transaction.
- `PUT /api/transactions/{id}` – update a transaction.
- `DELETE /api/transactions/{id}` – remove a transaction.
