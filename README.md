# CRUD Category API

This project exposes a simple REST API for managing products and categories using Go's standard `net/http` package. The router is mounted in `routes/routes.go`, while the request handling logic resides inside the controller packages. All data is stored in-memory via slices of DTO structs.

## Base URL

```
http://localhost:8080
```

## Health Check

- **GET** `/health`
  - Returns a simple JSON confirming the server is running.

### Sample Response

```json
{
  "status": "OK",
  "message": "Server is up and running"
}
```

## Products

### GET /api/products

- Returns the full list of products.
- Response body is a JSON array of product objects.

### GET /api/products/{id}

- Returns a single product by `id`.
- Responds with `404` if no product matches.

### POST /api/products

- Creates a new product.
- Expects JSON payload:

```json
{
  "name": "Product 1",
  "price": 3000,
  "stock": 100
}
```

- Response: `201 Created` with the created product (includes assigned `id`).

### PUT /api/products/{id}

- Updates an existing product by `id`.
- Request body follows the same schema as the create payload.
- Responds with `404` if the product does not exist.

### DELETE /api/products/{id}

- Removes the product with the given `id`.
- Returns `204 No Content` on success or `404` if not found.

## Categories

### GET /api/categories

- Returns a JSON array of category objects.

### GET /api/categories/{id}

- Returns a single category by `id`.

### POST /api/categories

- Creates a new category.
- Expects JSON payload:

```json
{
  "name": "Category 1",
  "description": "This is the first category"
}
```

- Responds with `201 Created` and the saved category object.

### PUT /api/categories/{id}

- Updates a category by `id`.
- Request body uses the same schema as create.

### DELETE /api/categories/{id}

- Deletes the specified category.
- Returns `204` on success or `404` when not found.

## Checkout

### POST /api/checkout

- Initiates a checkout process.
- Expects JSON payload:

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

- Responds with `201 Created` and the order object.

## Reports

### GET /api/reports/transaction

- Returns a JSON of order objects.

```json
{
  "total_revenue": 60000,
  "total_transactions": 4,
  "best_seller": {
    "id": 9,
    "name": "Product 5",
    "sold_quantity": 5
  }
}
```

## DTO Schemas

### Product

```json
{
  "id": 1,
  "name": "Product Name",
  "price": 1234,
  "stock": 10
}
```

### Category

```json
{
  "id": 1,
  "name": "Category Name",
  "description": "Some description"
}
```

## Notes

- All controllers decode requests via `encoding/json` and respond with JSON.
- IDs are assigned sequentially based on the slice length.
- Since data is kept in memory, restarting the server resets all products and categories.
- The routes rely on Go 1.25's built-in routing, so same path/method combinations must match exactly as registered in `routes/routes.go`.
