# 7. Get Floor Plan by Floor Number

Endpoint ini digunakan untuk mengambil detail denah lantai, termasuk URL gambar dan data semua meja di lantai tersebut. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /admin/floor-plans/:floor_number`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/admin/floor-plans/1' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Floor plan retrieved successfully",
    "data": {
        "id": 1,
        "floor_number": 1,
        "image_url": "/public/uploads/1728243900_floor_plan.jpg",
        "tables": [
            {
                "table_id": 1,
                "table_number": "01",
                "table_name": "Dekat Jendela",
                "x": 120,
                "y": 250,
                "active_users_count": 3,
                "wishlist_id": 43
            },
            {
                "table_id": 2,
                "table_number": "02",
                "table_name": "Tengah",
                "x": 350,
                "y": 250,
                "active_users_count": 2,
                "wishlist_id": 43
            },
            {
                "table_id": 6,
                "table_number": "06",
                "table_name": "Payung Outdoor",
                "x": 600,
                "y": 300,
                "active_users_count": 0,
                "wishlist_id": 43
            }
        ]
    }
}
```