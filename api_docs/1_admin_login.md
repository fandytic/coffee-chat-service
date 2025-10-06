# 1. Admin Login

Endpoint ini digunakan untuk otentikasi admin dan mendapatkan Bearer Token untuk mengakses endpoint lain yang terproteksi.

- **Endpoint**: `POST /login`
- **Content-Type**: `application/json`

---

### Request Body

```json
{
    "username": "admin",
    "password": "password123"
}
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "password": "password123"
}'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Login successful",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjIyNTA2MjcsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.some_generated_signature_string"
    }
}
```