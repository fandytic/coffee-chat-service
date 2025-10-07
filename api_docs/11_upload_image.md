# 11. Upload Image

Endpoint ini digunakan untuk mengunggah file gambar (seperti denah lantai). Responsnya akan berisi URL publik dari gambar yang dapat digunakan di API lain. **Endpoint ini terproteksi**.

- **Endpoint**: `POST /upload-image`
- **Content-Type**: `multipart/form-data`
- **Authentication**: `Bearer Token`

---

### Request Body (form-data)

- **`image`** (file): File gambar yang ingin diunggah (e.g., `.jpg`, `.png`).

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/upload-image' \
--header 'Authorization: Bearer <TOKEN>' \
--form 'image=@"/path/to/your/floor_plan.jpg"'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Image uploaded successfully",
    "data": {
        "image_url": "/public/uploads/1728243900_floor_plan.jpg"
    }
}
```