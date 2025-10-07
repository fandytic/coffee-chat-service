# 17. Get Customer Floor Plan

Endpoint ini digunakan oleh pelanggan untuk mengambil detail denah lantai, termasuk URL gambar dan data semua meja beserta jumlah pelanggan aktif di setiap meja. **Endpoint ini terproteksi** dan memerlukan token otentikasi pelanggan.

- **Endpoint**: `GET /customer/floor-plans/:floor_number`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/floor-plans/1' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

Responsnya sama persis dengan endpoint admin, memberikan semua data yang dibutuhkan untuk merender denah interaktif di sisi front-end.

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
                "active_users_count": 3
            },
            {
                "table_id": 2,
                "table_number": "02",
                "table_name": "Tengah",
                "x": 350,
                "y": 250,
                "active_users_count": 2
            }
        ]
    }
}
```