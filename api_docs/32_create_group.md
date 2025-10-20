# 32. Customer: Create Group Chat

Endpoint ini digunakan oleh pelanggan untuk membuat grup chat baru. Pembuat grup secara otomatis ditambahkan sebagai anggota. **Endpoint ini terproteksi**.

-   **Endpoint**: `POST /customer/groups`
-   **Authentication**: `Bearer Token` (Customer)

### Request Body
-   **`name`** (string): Nama grup chat.
-   **`member_ids`** (array of int): Daftar ID pelanggan lain yang ingin langsung diundang ke grup.

```json
{
    "name": "Tim Kopi Pagi",
    "member_ids": [5, 8]
}

``success
{
    "success": true,
    "code": 201,
    "message": "Group created successfully",
    "data": {
        "id": 1,
        "name": "Tim Kopi Pagi",
        "creator_id": 2
    }
}