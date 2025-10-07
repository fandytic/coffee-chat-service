# 7. Get Floor Plan by Floor Number

Endpoint ini digunakan untuk mengambil detail denah lantai, termasuk URL gambar dan data semua meja di lantai tersebut. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /floor-plans/:floor_number`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/floor-plans/1' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 201,
    "message": "Floor plan created successfully",
    "data": {
        "id": 1,
        "floor_number": 1,
        "image_url": "/public/uploads/1759802274_156380C0-A84B-4846-9FDF-9B531BDEBA95.JPG",
        "tables": [
            {
                "table_number": "01",
                "table_name": "Dekat Jendela",
                "x": 120,
                "y": 250
            }
        ]
    }
}
```