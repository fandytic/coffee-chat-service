# 20. Admin: Get All Customers

Endpoint ini digunakan oleh admin untuk mendapatkan daftar semua pelanggan yang pernah atau sedang melakukan *check-in*. Data bisa dicari berdasarkan nama pelanggan. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /admin/customers`
- **Authentication**: `Bearer Token` (Admin)

---

### Query Parameter (Opsional)

- **`search`** (string): Filter daftar pelanggan yang namanya mengandung teks ini.

### Contoh cURL

**Mengambil semua pelanggan:**
```sh
curl --location 'http://localhost:8080/admin/customers' \
--header 'Authorization: Bearer <ADMIN_TOKEN>'
```

**Mencari pelanggan dengan nama "Mary":**
```sh
curl --location 'http://localhost:8080/admin/customers?search=Mary' \
--header 'Authorization: Bearer <ADMIN_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Customers retrieved successfully",
    "data": [
        {
            "id": 1,
            "name": "Mary Holmes",
            "photo_url": "/public/uploads/...",
            "table_number": "01",
            "status": "active",
            "last_login": "2025-09-16T14:22:00Z"
        },
        {
            "id": 4,
            "name": "Johnny Mendez",
            "photo_url": "/public/uploads/...",
            "table_number": "02",
            "status": "inactive",
            "last_login": "2025-09-16T13:10:00Z"
        }
    ]
}
```