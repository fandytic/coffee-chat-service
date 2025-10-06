# 3. Get All Messages

Endpoint ini digunakan untuk mengambil seluruh riwayat percakapan yang tersimpan di database.

> **Catatan**: Idealnya, endpoint ini harus diproteksi dan memerlukan `Authorization: Bearer <token>` di header.

- **Endpoint**: `GET /messages`

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/messages'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Messages retrieved successfully",
    "data": [
        {
            "id": 1,
            "user": "Budi",
            "text": "Halo, selamat pagi semua!",
            "timestamp": "2025-10-07T13:45:00.123Z"
        },
        {
            "id": 2,
            "user": "Citra",
            "text": "Pagi juga, Budi!",
            "timestamp": "2025-10-07T13:45:15.456Z"
        }
    ]
}
```