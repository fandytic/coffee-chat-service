# 2. Send Message

Endpoint ini digunakan untuk mengirim pesan baru ke dalam chat. Pesan akan disimpan di database dan disiarkan secara real-time ke semua klien WebSocket yang terhubung.

> **Catatan**: Idealnya, endpoint ini harus diproteksi dan memerlukan `Authorization: Bearer <token>` di header.

- **Endpoint**: `POST /admin/send`
- **Content-Type**: `application/json`

---

### Request Body

```json
{
    "user": "Budi",
    "text": "Halo, selamat pagi semua!"
}
```

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/admin/send' \
--header 'Content-Type: application/json' \
--data '{
    "user": "Budi",
    "text": "Halo, selamat pagi semua!"
}'
```

---

### Contoh Success Response (Code: 201)

```json
{
    "success": true,
    "code": 201,
    "message": "Message sent successfully",
    "data": {
        "id": 1,
        "user": "Budi",
        "text": "Halo, selamat pagi semua!",
        "timestamp": "2025-10-07T13:45:00.123Z"
    }
}
```