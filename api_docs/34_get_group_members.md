# 34. Customer: Get Group Members

Endpoint ini digunakan oleh anggota grup untuk melihat daftar semua pelanggan yang ada di dalam grup. **Endpoint ini terproteksi**.

-   **Endpoint**: `GET /customer/groups/:group_id/members`
-   **Authentication**: `Bearer Token` (Customer)
-   **Path Parameter**: `:group_id` - ID dari grup.

---
### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Group members retrieved successfully",
    "data": [
        {
            "id": 2,
            "name": "Christine",
            "photo_url": "/public/uploads/...",
            "table_number": "01",
            "floor_number": 1
        },
        {
            "id": 5,
            "name": "Edward",
            "photo_url": "/public/uploads/...",
            "table_number": "03",
            "floor_number": 1
        },
        {
            "id": 8,
            "name": "Mary Holmes",
            "photo_url": "/public/uploads/...",
            "table_number": "05",
            "floor_number": 2
        }
    ]
}