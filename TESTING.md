# Unit Testing Documentation

## Overview

โปรเจคนี้มี unit tests ครอบคลุม handlers และ repository layers โดยใช้:
- **testify** - สำหรับ assertions และ mocking
- **SQLite in-memory** - สำหรับ integration tests

## Test Coverage

```
handlers     - 75.0% coverage (10 tests)
repository   - 87.5% coverage (8 tests)
```

## Running Tests

### รัน tests ทั้งหมด
```bash
go test ./... -v
```

หรือใช้ Makefile:
```bash
make test
```

### รัน tests แบบเฉพาะ package

#### Handler tests
```bash
go test ./handlers -v
```

#### Repository tests
```bash
go test ./repository -v
```

### รัน tests พร้อม coverage
```bash
make test-coverage
```

### สร้าง HTML coverage report
```bash
make test-coverage-html
```

จะสร้างไฟล์ `coverage.html` ที่สามารถเปิดดูใน browser

## Test Structure

### Handler Tests (`handlers/user_test.go`)

Tests สำหรับ HTTP handlers โดยใช้ mock repository:

**Tests ที่มี:**
1. `TestGetUsers_Success` - ทดสอบดึงข้อมูล users ทั้งหมด
2. `TestGetUser_Success` - ทดสอบดึงข้อมูล user ตาม ID
3. `TestGetUser_NotFound` - ทดสอบกรณี user ไม่พบ
4. `TestGetUser_InvalidID` - ทดสอบกรณี ID ไม่ถูกต้อง
5. `TestCreateUser_Success` - ทดสอบสร้าง user ใหม่
6. `TestCreateUser_MissingFields` - ทดสอบกรณีข้อมูลไม่ครบ
7. `TestUpdateUser_Success` - ทดสอบอัพเดท user
8. `TestUpdateUser_NotFound` - ทดสอบกรณี user ที่จะอัพเดทไม่พบ
9. `TestDeleteUser_Success` - ทดสอบลบ user
10. `TestDeleteUser_NotFound` - ทดสอบกรณี user ที่จะลบไม่พบ

**ตัวอย่าง:**
```go
func TestGetUsers_Success(t *testing.T) {
    // Setup mock repository
    mockRepo := new(MockUserRepository)
    userRepo = mockRepo
    
    users := []models.User{...}
    mockRepo.On("GetAll").Return(users, nil)
    
    // Test HTTP request
    app := setupTestApp()
    app.Get("/users", GetUsers)
    
    req := httptest.NewRequest("GET", "/users", nil)
    resp, err := app.Test(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 200, resp.StatusCode)
    mockRepo.AssertExpectations(t)
}
```

### Repository Tests (`repository/user_repository_test.go`)

Integration tests สำหรับ database operations โดยใช้ SQLite in-memory:

**Tests ที่มี:**
1. `TestUserRepository_Create` - ทดสอบสร้าง user ใหม่
2. `TestUserRepository_GetByID` - ทดสอบดึงข้อมูล user ตาม ID
3. `TestUserRepository_GetByID_NotFound` - ทดสอบกรณี user ไม่พบ
4. `TestUserRepository_GetAll` - ทดสอบดึงข้อมูล users ทั้งหมด
5. `TestUserRepository_Update` - ทดสอบอัพเดท user
6. `TestUserRepository_Delete` - ทดสอบลบ user
7. `TestUserRepository_Create_DuplicateEmail` - ทดสอบ unique constraint
8. `TestUserRepository_GetAll_Empty` - ทดสอบกรณีไม่มีข้อมูล

**ตัวอย่าง:**
```go
func TestUserRepository_Create(t *testing.T) {
    // Setup in-memory database
    db, cleanup := setupTestDB(t)
    defer cleanup()
    
    repo := NewUserRepository(db)
    
    req := models.CreateUserRequest{
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john@example.com",
        // ...
    }
    
    user, err := repo.Create(req)
    
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "John", user.FirstName)
}
```

## Mocking Strategy

### Repository Interface

สร้าง interface เพื่อให้ mock ได้ง่าย:

```go
type UserRepositoryInterface interface {
    GetAll() ([]models.User, error)
    GetByID(id int) (*models.User, error)
    Create(req models.CreateUserRequest) (*models.User, error)
    Update(id int, req models.UpdateUserRequest) (*models.User, error)
    Delete(id int) error
}
```

### Mock Implementation

```go
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
    args := m.Called()
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.User), args.Error(1)
}
```

## Test Database Setup

ใช้ SQLite in-memory database สำหรับแต่ละ test:

```go
func setupTestDB(t *testing.T) (*sql.DB, func()) {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open test database: %v", err)
    }
    
    // Create tables...
    
    cleanup := func() {
        db.Close()
    }
    
    return db, cleanup
}
```

## Best Practices

1. **แยก unit tests และ integration tests**
   - Handler tests ใช้ mocks
   - Repository tests ใช้ real database (in-memory)

2. **ใช้ testify สำหรับ assertions**
   ```go
   assert.NoError(t, err)
   assert.Equal(t, expected, actual)
   assert.NotNil(t, value)
   ```

3. **Cleanup resources**
   ```go
   defer cleanup()
   defer mockRepo.AssertExpectations(t)
   ```

4. **Test error cases**
   - Invalid inputs
   - Not found scenarios
   - Database constraints

5. **Use descriptive test names**
   - Format: `Test<Function>_<Scenario>`
   - Example: `TestGetUser_NotFound`

## Continuous Integration

เพิ่ม test commands ใน CI/CD pipeline:

```yaml
# Example GitHub Actions
- name: Run tests
  run: go test ./... -v

- name: Check coverage
  run: go test ./... -cover
```

## Future Improvements

- [ ] เพิ่ม integration tests สำหรับ API endpoints
- [ ] เพิ่ม benchmark tests
- [ ] เพิ่ม test coverage เป็น 90%+
- [ ] เพิ่ม tests สำหรับ middleware
- [ ] เพิ่ม tests สำหรับ config package

## Dependencies

```go
require (
    github.com/stretchr/testify v1.11.1
    github.com/mattn/go-sqlite3 v1.14.19
)
```

## Running Specific Tests

### รัน test เดียว
```bash
go test ./handlers -run TestGetUsers_Success -v
```

### รัน tests ที่ match pattern
```bash
go test ./... -run ".*User.*" -v
```

### รัน tests พร้อม timeout
```bash
go test ./... -timeout 30s
```

## Troubleshooting

### Tests ล้มเหลว
1. ตรวจสอบว่า dependencies ติดตั้งครบ: `go mod tidy`
2. ตรวจสอบ database schema ใน test setup
3. ตรวจสอบ mock expectations

### Coverage ต่ำ
1. เพิ่ม test cases สำหรับ edge cases
2. ทดสอบ error scenarios
3. ครอบคลุม validation logic

---

**สรุป:** โปรเจคมี unit tests ที่ครอบคลุมทั้ง handlers และ repository layers พร้อม coverage 75-87%
