# 6. Create or Update Floor Plan

Endpoint ini digunakan untuk mengunggah gambar denah lantai beserta data posisi meja. Jika denah untuk lantai tersebut sudah ada, endpoint ini akan menimpanya. **Endpoint ini terproteksi**.

- **Endpoint**: `POST /floor-plans`
- **Content-Type**: `multipart/form-data`
- **Authentication**: `Bearer Token`

---

### Request Body (form-data)

- **`floor_plan_image`** (file): File gambar denah (e.g., `.jpg`, `.png`).
- **`floor_number`** (text): Nomor lantai, misalnya `1`.
- **`tables`** (text): Sebuah string JSON yang berisi array data meja.

**Contoh isi `tables`:**
```json
[
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
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/floor-plans' \
--header 'Authorization: Bearer <TOKEN>' \
--form 'floor_plan_image=@"/path/to/your/floor_plan.jpg"' \
--form 'floor_number="1"' \
--form 'tables="[{\"table_number\":\"01\",\"table_name\":\"Dekat Jendela\",\"x\":120,\"y\":250},{\"table_number\":\"02\",\"table_name\":\"Tengah\",\"x\":350,\"y\":250}]"'
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
}
```