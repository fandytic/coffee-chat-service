# 36. Customer: Get My Groups

Endpoint ini digunakan oleh pelanggan untuk mengambil daftar semua grup chat di mana mereka terdaftar sebagai anggota.

-   **Endpoint**: `GET /customer/groups`
-   **Authentication**: `Bearer Token` (Customer)

---
### Contoh cURL
```sh
curl --location 'http://localhost:8080/customer/groups' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

---
### Contoh Success Response (Code: 200)

Responsnya adalah sebuah array yang berisi daftar grup yang diikuti oleh pelanggan.

```json
{
    "success": true,
    "code": 200,
    "message": "Groups retrieved successfully",
    "data": [
        {
            "id": 1,
            "name": "Tim Santai",
            "creator_id": 2,
            "members": [
                "Edward",
                "Christine",
                "Megan"
            ],
            "last_message": {
                    "text": "Bro, join sini lah...",
                    "timestamp": "2025-10-07T20:02:00Z"
                },
            "unread_count": 2
        },
        {
            "id": 2,
            "name": "Anak IPB",
            "creator_id": 5,
            "members": [
                "Christine",
                "Jeremy"
            ],
            "last_message": {
                    "text": "Bro, join sini lah...",
                    "timestamp": "2025-10-07T20:02:00Z"
                },
            "unread_count": 0
        }
    ]
}
```