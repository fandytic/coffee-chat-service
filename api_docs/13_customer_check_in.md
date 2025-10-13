# 13. Customer Check-in

Endpoint ini digunakan oleh pelanggan setelah memindai QR code meja untuk mendaftarkan diri. `table_id` didapatkan dari QR code. Jika pelanggan ingin mengunggah foto, mereka harus menggunakan endpoint `/upload-image` terlebih dahulu dan menyertakan URL hasilnya di `photo_url`.

- **Endpoint**: `POST /check-in`
- **Content-Type**: `application/json`

---

### Request Body

```json
{
    "table_id": 1,
    "name": "Christine Stanley",
    "photo_url": "/public/uploads/1728345678_profile.jpg"
}
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/check-in' \
--header 'Content-Type: application/json' \
--data '{
    "table_id": 1,
    "name": "Christine Stanley",
    "photo_url": ""
}'
```

---

### Contoh Success Response (Code: 201)

Responsnya berisi detail customer dan `auth_token` yang harus disimpan oleh front-end untuk mengakses API lain dan WebSocket.

```json
{
    "success": true,
    "code": 201,
    "message": "Check-in successful",
    "data": {
        "id": 1,
        "name": "Christine Stanley",
        "photo_url": "",
        "table_id": 1,
        "table_number": "01",
        "floor_number": 1,
        "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lcl9pZCI6MSwiZXhwIjoxNzI4MjgwNzgwLCJuYW1lIjoiQ2hyaXN0aW5lIFN0YW5sZXkiLCJ0YWJsZV9pZCI6MX0.some_signature"
    }
}
```