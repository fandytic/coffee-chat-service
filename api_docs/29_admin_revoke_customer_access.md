# 29. Admin: Revoke Customer Access

Endpoint ini digunakan oleh admin untuk mencabut akses seorang pelanggan. Operasi ini tidak menghapus data pelanggan, melainkan mengubah statusnya menjadi `revoked`. Setelah akses dicabut, token pelanggan tersebut akan menjadi tidak valid dan mereka harus melakukan check-in ulang untuk mendapatkan sesi baru. **Endpoint ini terproteksi**.

- **Endpoint**: `DELETE /admin/customers/:id`
- **Authentication**: `Bearer Token` (Admin)

### Path Parameter
- **`:id`**: ID dari pelanggan yang aksesnya ingin dicabut.

---
### Contoh cURL
```sh
curl --location --request DELETE 'http://localhost:8080/admin/customers/5' \
--header 'Authorization: Bearer <ADMIN_TOKEN>'
```

---
### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Customer access revoked successfully"
}
```

---
### Contoh Error Response (Code: 404 Not Found)
```json
{
    "success": false,
    "code": 404,
    "message": "customer with ID 999 not found"
}
```