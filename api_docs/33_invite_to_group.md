# 33. Customer: Invite to Group

Endpoint ini digunakan oleh anggota grup untuk mengundang pelanggan lain ke grup chat yang sudah ada. **Endpoint ini terproteksi**.

-   **Endpoint**: `POST /customer/groups/:group_id/members`
-   **Authentication**: `Bearer Token` (Customer)
-   **Path Parameter**: `:group_id` - ID dari grup.

### Request Body
-   **`customer_ids`** (array of int): Daftar ID pelanggan yang ingin diundang.

```json
{
    "customer_ids": [10, 12]
}

```succsess
{
    "success": true,
    "code": 200,
    "message": "Members invited successfully"
}

```error
{
    "success": false,
    "code": 403,
    "message": "forbidden: you are not a member of this group"
}