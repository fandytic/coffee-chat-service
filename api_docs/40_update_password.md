# 40. Update Password Admin

Endpoint ini digunakan oleh admin yang sudah login untuk memperbarui password mereka sendiri.

- **Endpoint**: `PUT /admin/update-password`
- **Content-Type**: `application/json`
- **Authorization**: `Bearer <token>`

---

### Request Body

```json
{
    "old_password": "password123",
    "new_password": "newpassword123"
}
```

---

### Contoh cURL

```sh
curl --location --request PUT 'http://localhost:8080/admin/update-password' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "old_password": "password123",
    "new_password": "newpassword123"
}'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Password updated successful",
    "data": null
}
```
