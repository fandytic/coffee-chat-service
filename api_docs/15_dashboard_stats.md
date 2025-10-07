# 15. Get Dashboard Statistics

Endpoint ini digunakan untuk mengambil data statistik agregat untuk ditampilkan di dasbor admin. **Endpoint ini terproteksi** dan hanya bisa diakses oleh admin.

- **Endpoint**: `GET /admin/dashboard/stats`
- **Authentication**: `Bearer Token` (Admin)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/admin/dashboard/stats' \
--header 'Authorization: Bearer <ADMIN_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Dashboard statistics retrieved successfully",
    "data": {
        "total_tables": 12,
        "occupied_tables": 5,
        "empty_tables": 7,
        "active_users": 14
    }
}
```