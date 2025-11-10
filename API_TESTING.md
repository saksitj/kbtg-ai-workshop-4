# API Testing Examples

## Using curl

### Health Check
```bash
curl http://localhost:3000/health
```

### Welcome
```bash
curl http://localhost:3000/
```

### Get All Users
```bash
curl http://localhost:3000/api/v1/users
```

### Get User by ID
```bash
curl http://localhost:3000/api/v1/users/1
```

### Create User
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Update User
```bash
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:3000/api/v1/users/1
```

## Using HTTPie (if installed)

### Get All Users
```bash
http GET localhost:3000/api/v1/users
```

### Create User
```bash
http POST localhost:3000/api/v1/users \
  name="Jane Doe" \
  email="jane@example.com"
```

### Update User
```bash
http PUT localhost:3000/api/v1/users/1 \
  name="Jane Updated" \
  email="jane.updated@example.com"
```

### Delete User
```bash
http DELETE localhost:3000/api/v1/users/1
```

## Expected Response Examples

### Success Response
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

### Error Response
```json
{
  "success": false,
  "error": "Invalid request body"
}
```

### List Response
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "email": "jane@example.com"
    }
  ]
}
```
