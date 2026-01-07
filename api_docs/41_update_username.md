# 41. Update Username Admin

Endpoint ini digunakan oleh admin yang sudah login untuk memperbarui username mereka sendiri.

- **Endpoint**: `PUT /admin/update-username`
- **Content-Type**: `application/json`
- **Authorization**: `Bearer <token>`

---

### Request Body

```json
{
    "new_username": "newadmin"
}
```

---

### Contoh cURL

```sh
curl --location --request PUT 'http://localhost:8080/admin/update-username' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data '{
    "new_username": "newadmin"
}'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Username updated successful",
    "data": null
}
```
