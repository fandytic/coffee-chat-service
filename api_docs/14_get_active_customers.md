# 14. Get Active Customers List

Endpoint ini mengambil daftar semua pelanggan yang saat ini berstatus aktif di kafe, beserta metadata, hitungan pesan belum dibaca, dan pesan terakhir dari setiap percakapan. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /customer/active-list`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/active-list' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Active customers retrieved successfully",
    "data": {
        "total": 5,
        "customers": [
            {
                "id": 1,
                "name": "Mary Holmes",
                "photo_url": "/public/uploads/...",
                "table_number": "01",
                "floor_number": 1,
                "unread_messages_count": 3,
                "last_message": {
                    "text": "Kok ga dibales?",
                    "timestamp": "2025-10-07T20:03:00Z"
                }
            },
            {
                "id": 2,
                "name": "Jeremy Gibson",
                "photo_url": "/public/uploads/...",
                "table_number": "01",
                "floor_number": 1,
                "unread_messages_count": 1,
                "last_message": {
                    "text": "Bro, join sini lah...",
                    "timestamp": "2025-10-07T20:02:00Z"
                }
            },
            {
                "id": 4,
                "name": "Johnny Mendez",
                "photo_url": "/public/uploads/...",
                "table_number": "03",
                "floor_number": 1,
                "unread_messages_count": 0,
                "last_message": null
            }
        ]
    }
}
```