# 24. Customer: Get Chat History

Endpoint ini digunakan untuk mengambil riwayat percakapan antara pengguna yang sedang login dengan pengguna lain. **Endpoint ini terproteksi**.

-   **Endpoint**: `GET /customer/chats/:id`
-   **Authentication**: `Bearer Token` (Customer)

### Path Parameter
- **`:id`**: ID dari pelanggan lain yang riwayat chatnya ingin Anda lihat.

---
### Contoh cURL

Misalnya, Anda (yang sedang login) ingin melihat riwayat chat dengan pelanggan lain yang memiliki ID `5`.
```sh
curl --location 'http://localhost:8080/customer/chats/5' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

---
### Contoh Success Response (Code: 200)

Responsnya adalah array dari semua objek pesan antara kedua pengguna, diurutkan dari yang terlama.

```json
{
    "success": true,
    "code": 200,
    "message": "Chat history retrieved successfully",
    "data": [
        {
            "ID": 123,
            "CreatedAt": "2025-10-07T19:59:00Z",
            "UpdatedAt": "...",
            "DeletedAt": null,
            "SenderID": 5,
            "RecipientID": 2,
            "Text": "Edward Cullen? haha",
            "ReplyToMessageID": null,
            "IsRead": true
        },
        {
            "ID": 125,
            "CreatedAt": "2025-10-07T20:02:00Z",
            "UpdatedAt": "...",
            "DeletedAt": null,
            "SenderID": 2,
            "RecipientID": 5,
            "Text": "Haha mirip dikit",
            "ReplyToMessageID": 123,
            "IsRead": true
        }
    ]
}
```