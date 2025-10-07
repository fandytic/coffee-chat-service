# 14. Get Active Customers List

Endpoint ini mengambil daftar semua pelanggan yang saat ini berstatus aktif di kafe. **Endpoint ini terproteksi** dan memerlukan token otentikasi yang didapat dari API Check-in.

- **Endpoint**: `GET /customer/active-list`
- **Authentication**: `Bearer Token`

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
    "data": [
        {
            "id": 1,
            "name": "Mary Holmes",
            "photo_url": "/public/uploads/...",
            "table_number": "01",
            "unread_messages_count": 3
        },
        {
            "id": 1,
            "name": "Mary Holmes",
            "photo_url": "/public/uploads/...",
            "table_number": "01",
            "unread_messages_count": 3
        }
    ]
}
```