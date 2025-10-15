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

Responsnya adalah array dari semua objek pesan antara kedua pengguna, diurutkan dari yang terlama. Struktur datanya sama dengan payload yang dikirimkan lewat websocket sehingga front end dapat menampilkan balasan maupun traktiran dengan konsisten.

Ketika sebuah pesan merupakan balasan terhadap traktiran, properti `reply_to.menu` akan berisi detail menu yang sama seperti kartu traktiran asli.

```json
{
    "success": true,
    "code": 200,
    "message": "Chat history retrieved successfully",
    "data": [
        {
            "message_id": 123,
            "sender_id": 5,
            "sender_name": "Christine",
            "sender_photo_url": "https://cdn.example.com/avatars/christine.png",
            "sender_table_number": "B12",
            "sender_floor_number": 2,
            "text": "Edward Cullen? haha",
            "timestamp": "2025-10-07T19:59:00Z"
        },
        {
            "message_id": 125,
            "sender_id": 2,
            "sender_name": "Edward",
            "sender_photo_url": "https://cdn.example.com/avatars/edward.png",
            "sender_table_number": "A04",
            "sender_floor_number": 1,
            "text": "Haha mirip dikit",
            "timestamp": "2025-10-07T20:02:00Z",
            "reply_to": {
                "id": 123,
                "text": "Edward Cullen? haha",
                "sender_name": "Christine",
                "menu": {
                    "id": 42,
                    "name": "Cappuccino",
                    "price": 35000,
                    "image_url": "https://cdn.example.com/menu/cappuccino.png"
                }
            }
        },
        {
            "message_id": 127,
            "sender_id": 5,
            "sender_name": "Christine",
            "sender_photo_url": "https://cdn.example.com/avatars/christine.png",
            "sender_table_number": "B12",
            "sender_floor_number": 2,
            "text": "Ini traktir cappuccino favoritmu ya!",
            "timestamp": "2025-10-07T20:05:00Z",
            "menu": {
                "id": 42,
                "name": "Cappuccino",
                "price": 35000,
                "image_url": "https://cdn.example.com/menu/cappuccino.png"
            }
        }
    ]
}
```