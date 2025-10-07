# 8. Get All Floors

Endpoint ini digunakan untuk mendapatkan daftar semua lantai yang telah dibuat denahnya. Responsnya sederhana, hanya berisi ID dan nomor lantai. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /admin/floor-plans`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/admin/floor-plans' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Floors retrieved successfully",
    "data": [
        {
            "id": 1,
            "floor_number": 1
        },
        {
            "id": 2,
            "floor_number": 2
        }
    ]
}
```