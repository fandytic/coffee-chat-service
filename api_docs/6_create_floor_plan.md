# 6. Create Floor Plan

Endpoint ini digunakan untuk membuat data denah lantai baru, termasuk posisi meja-mejanya. Endpoint ini menggunakan URL gambar yang didapat dari API Upload Image (`/upload-image`). **Endpoint ini terproteksi**.

- **Endpoint**: `POST /floor-plans`
- **Content-Type**: `application/json`
- **Authentication**: `Bearer Token`

---

### Request Body

```json
{
    "floor_number": 1,
    "image_url": "/public/uploads/1728243900_floor_plan.jpg",
    "tables": [
        {
            "table_number": "01",
            "table_name": "Dekat Jendela",
            "x": 120,
            "y": 250
        },
        {
            "table_number": "02",
            "table_name": "Tengah",
            "x": 350,
            "y": 250
        }
    ]
}
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/floor-plans' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data '{
    "floor_number": 1,
    "image_url": "/public/uploads/1728243900_floor_plan.jpg",
    "tables": [
        {
            "table_number": "01",
            "table_name": "Dekat Jendela",
            "x": 120,
            "y": 250
        }
    ]
}'
```
---

### Contoh Success Response (Code: 201)
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