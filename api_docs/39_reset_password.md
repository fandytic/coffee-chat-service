# 39. Reset Password

Endpoint ini digunakan untuk mereset password admin menggunakan username. Memerlukan token admin.

- **Endpoint**: `POST /admin/reset-password`
- **Content-Type**: `application/json`
- **Authorization**: `Bearer <token>`

---

### Request Body

```json
{
    "username": "admin",
    "new_password": "newpassword123"
}
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/admin/reset-password' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "new_password": "newpassword123"
}'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Password reset successful",
    "data": null
}
```
