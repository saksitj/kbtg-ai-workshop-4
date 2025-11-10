# ✅ CRUD Users with SQLite - Setup Complete!

## 🎯 สิ่งที่สร้างเสร็จแล้ว

### 1. Database Setup ✅
- ✅ ติดตั้ง SQLite driver (`github.com/mattn/go-sqlite3`)
- ✅ สร้าง database package สำหรับจัดการ connection
- ✅ สร้างตาราง `users` พร้อม schema ที่ครบถ้วน
- ✅ Auto-create database file (`users.db`) เมื่อรันครั้งแรก

### 2. User Model ✅
Fields ที่สอดคล้องกับ UI Profile Form:
- `id` - Primary key (auto-increment)
- `name` - ชื่อผู้ใช้ (required)
- `email` - อีเมล (required, unique)
- `phone` - เบอร์โทรศัพท์
- `address` - ที่อยู่
- `avatar` - URL รูปโปรไฟล์
- `created_at` - วันที่สร้าง (auto)
- `updated_at` - วันที่แก้ไข (auto)

### 3. Repository Pattern ✅
สร้าง `UserRepository` พร้อมฟังก์ชัน:
- `GetAll()` - ดึงข้อมูลผู้ใช้ทั้งหมด
- `GetByID(id)` - ดึงข้อมูลผู้ใช้รายบุคคล
- `Create(req)` - สร้างผู้ใช้ใหม่
- `Update(id, req)` - แก้ไขข้อมูลผู้ใช้
- `Delete(id)` - ลบผู้ใช้

### 4. API Endpoints ✅
ตรงตามที่ระบุใน requirements:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | ดึงรายการผู้ใช้ทั้งหมด |
| GET | `/users/{id}` | ดึงข้อมูลรายละเอียดผู้ใช้รายบุคคล |
| POST | `/users` | สร้างผู้ใช้ใหม่ |
| PUT | `/users/{id}` | แก้ไขข้อมูลผู้ใช้ |
| DELETE | `/users/{id}` | ลบผู้ใช้ |

### 5. Error Handling ✅
- ✅ Validation สำหรับ required fields
- ✅ Check user not found (404)
- ✅ Invalid user ID handling (400)
- ✅ Database error handling
- ✅ Proper HTTP status codes

## 🚀 วิธีใช้งาน

### เริ่มต้น Server
```bash
cd /Users/saksit.ja/Desktop/workshop_4
go run main.go
```

Server จะรันที่: **http://localhost:3000**

### ทดสอบ API

#### 1. สร้างผู้ใช้ใหม่
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "สมชาย ใจดี",
    "email": "somchai@example.com",
    "phone": "0812345678",
    "address": "123 ถนนสุขุมวิท กรุงเทพฯ",
    "avatar": "https://example.com/somchai.jpg"
  }'
```

#### 2. ดึงข้อมูลผู้ใช้ทั้งหมด
```bash
curl http://localhost:3000/users
```

#### 3. ดึงข้อมูลผู้ใช้รายบุคคล
```bash
curl http://localhost:3000/users/1
```

#### 4. แก้ไขข้อมูลผู้ใช้
```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "สมชาย ใจดี (แก้ไข)",
    "email": "somchai.updated@example.com",
    "phone": "0899999999",
    "address": "456 ถนนใหม่ กรุงเทพฯ",
    "avatar": "https://example.com/somchai-new.jpg"
  }'
```

#### 5. ลบผู้ใช้
```bash
curl -X DELETE http://localhost:3000/users/1
```

## 📁 โครงสร้างโปรเจค

```
workshop_4/
├── database/              # Database connection
│   └── db.go
├── repository/            # Data access layer
│   └── user_repository.go
├── models/                # Data models
│   └── user.go
├── handlers/              # Request handlers
│   └── user.go
├── routes/                # Route definitions
│   └── routes.go
├── config/                # Configuration
│   └── config.go
├── middleware/            # Middleware
│   └── auth.go
├── main.go                # Entry point
├── users.db               # SQLite database (auto-created)
└── API_DATABASE_TESTING.md # API documentation
```

## 🎨 Response Format

### Success Response
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "สมชาย ใจดี",
    "email": "somchai@example.com",
    "phone": "0812345678",
    "address": "123 ถนนสุขุมวิท กรุงเทพฯ",
    "avatar": "https://example.com/somchai.jpg",
    "created_at": "2025-11-10T13:44:00Z",
    "updated_at": "2025-11-10T13:44:00Z"
  }
}
```

### Error Response
```json
{
  "success": false,
  "error": "User not found"
}
```

## 💾 Database

### ดู Database โดยตรง
```bash
sqlite3 users.db
```

### SQL Commands
```sql
-- ดูข้อมูลทั้งหมด
SELECT * FROM users;

-- ดูจำนวนผู้ใช้
SELECT COUNT(*) FROM users;

-- ลบข้อมูลทั้งหมด
DELETE FROM users;
```

## 📝 Features

✅ Full CRUD operations  
✅ SQLite database integration  
✅ Repository pattern for clean architecture  
✅ Proper error handling  
✅ Request validation  
✅ Auto-managed timestamps  
✅ Unique email constraint  
✅ RESTful API design  
✅ JSON responses  
✅ CORS enabled  

## 🎯 ทดสอบแล้ว

✅ Server starts successfully  
✅ Database initialized  
✅ All endpoints working  
✅ Error handling works correctly  
✅ Data persists in SQLite  

---

**โปรเจคพร้อมใช้งานแล้ว! 🎉**

สำหรับรายละเอียดเพิ่มเติม ดูได้ที่: `API_DATABASE_TESTING.md`
