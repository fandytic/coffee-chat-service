# 16. Get Customer Dashboard Statistics

Endpoint ini digunakan untuk mengambil data statistik kafe yang akan ditampilkan kepada pelanggan setelah mereka berhasil melakukan check-in. **Endpoint ini terproteksi** dan memerlukan token otentikasi pelanggan.

- **Endpoint**: `GET /customer/stats`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/stats' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
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