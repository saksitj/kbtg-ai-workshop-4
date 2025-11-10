# API Testing with SQLite Database

## Database Schema

The `users` table has the following fields:
- `id` - Auto-increment primary key
- `first_name` - User's first name (required)
- `last_name` - User's last name (required)
- `email` - User's email address (required, unique)
- `phone` - User's phone number (optional)
- `address` - User's address (optional)
- `avatar` - URL to user's avatar image (optional)
- `member_level` - Membership level: Bronze, Silver, Gold, Platinum (default: Bronze)
- `point_balance` - User's loyalty points balance (default: 0)
- `created_at` - Timestamp when user was created
- `updated_at` - Timestamp when user was last updated

## API Endpoints

### 1. Get All Users
```bash
curl http://localhost:3000/users
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "first_name": "สมชาย",
      "last_name": "ใจดี",
      "email": "somchai@example.com",
      "phone": "0812345678",
      "address": "123 Main St, Bangkok",
      "avatar": "https://example.com/avatar.jpg",
      "member_level": "Gold",
      "point_balance": 1500,
      "created_at": "2025-11-10T13:44:00Z",
      "updated_at": "2025-11-10T13:44:00Z"
    }
  ]
}
```

### 2. Get User by ID
```bash
curl http://localhost:3000/users/1
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "first_name": "สมชาย",
    "last_name": "ใจดี",
    "email": "somchai@example.com",
    "phone": "0812345678",
    "address": "123 Main St, Bangkok",
    "avatar": "https://example.com/avatar.jpg",
    "member_level": "Gold",
    "point_balance": 1500,
    "created_at": "2025-11-10T13:44:00Z",
    "updated_at": "2025-11-10T13:44:00Z"
  }
}
```

### 3. Create New User
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "สมหญิง",
    "last_name": "รักดี",
    "email": "somying@example.com",
    "phone": "0898765432",
    "address": "456 Sukhumvit Rd, Bangkok",
    "avatar": "https://example.com/somying.jpg",
    "member_level": "Silver",
    "point_balance": 750
  }'
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "first_name": "สมหญิง",
    "last_name": "รักดี",
    "email": "somying@example.com",
    "phone": "0898765432",
    "address": "456 Sukhumvit Rd, Bangkok",
    "avatar": "https://example.com/somying.jpg",
    "member_level": "Silver",
    "point_balance": 750,
    "created_at": "2025-11-10T13:45:00Z",
    "updated_at": "2025-11-10T13:45:00Z"
  }
}
```

### 4. Update User
```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "สมชาย",
    "last_name": "ใจดี (อัพเดท)",
    "email": "somchai.updated@example.com",
    "phone": "0811111111",
    "address": "789 New Address, Bangkok",
    "avatar": "https://example.com/somchai-new.jpg",
    "member_level": "Platinum",
    "point_balance": 3000
  }'
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "first_name": "สมชาย",
    "last_name": "ใจดี (อัพเดท)",
    "email": "somchai.updated@example.com",
    "phone": "0811111111",
    "address": "789 New Address, Bangkok",
    "avatar": "https://example.com/somchai-new.jpg",
    "member_level": "Platinum",
    "point_balance": 3000,
    "created_at": "2025-11-10T13:44:00Z",
    "updated_at": "2025-11-10T13:46:00Z"
  }
}
```

### 5. Delete User
```bash
curl -X DELETE http://localhost:3000/users/1
```

**Response:**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

## Error Responses

### User Not Found (404)
```json
{
  "success": false,
  "error": "User not found"
}
```

### Invalid User ID (400)
```json
{
  "success": false,
  "error": "Invalid user ID"
}
```

### Missing Required Fields (400)
```json
{
  "success": false,
  "error": "First name, last name, and email are required"
}
```

### Invalid Request Body (400)
```json
{
  "success": false,
  "error": "Invalid request body"
}
```

## Testing Workflow

1. **Create a user:**
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "ทดสอบ",
    "last_name": "ระบบ",
    "email": "test@example.com",
    "phone": "0801234567",
    "address": "Test Address",
    "avatar": "https://example.com/test.jpg",
    "member_level": "Bronze",
    "point_balance": 100
  }'
```

2. **Get all users:**
```bash
curl http://localhost:3000/users
```

3. **Get specific user (use ID from step 1):**
```bash
curl http://localhost:3000/users/1
```

4. **Update user:**
```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "ทดสอบ",
    "last_name": "ระบบ (อัพเดท)",
    "email": "test.updated@example.com",
    "phone": "0809999999",
    "address": "Updated Address",
    "avatar": "https://example.com/test-updated.jpg",
    "member_level": "Silver",
    "point_balance": 500
  }'
```

5. **Delete user:**
```bash
curl -X DELETE http://localhost:3000/users/1
```

## Database Location

The SQLite database is stored at: `./users.db`

To view the database directly:
```bash
sqlite3 users.db "SELECT * FROM users;"
```

## Notes

- `first_name`, `last_name`, and `email` are required fields
- Email must be unique
- `member_level` defaults to "Bronze" if not specified (options: Bronze, Silver, Gold, Platinum)
- `point_balance` defaults to 0 if not specified
- Timestamps are automatically managed by the database
- The database file is automatically created on first run
