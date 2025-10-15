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
            "message_id": 127,
            "sender_id": 5,
            "sender_name": "Christine",
            "text": "Ini traktir cappuccino favoritmu ya!",
            "timestamp": "2025-10-07T20:05:00Z",
            "order": {
                "id": 88,
                "need_type": "order_for_other",
                "total": 38850,
                // ... detail pesanan lainnya
            }
        },
        {
            "message_id": 128,
            "sender_id": 2,
            "sender_name": "Edward",
            "text": "Wah, makasih banyak ya!",
            "timestamp": "2025-10-07T20:06:00Z",
            "reply_to": {
                "id": 127,
                "text": "Ini traktir cappuccino favoritmu ya!",
                "sender_name": "Christine",
                "order": {
                    "id": 88,
                    "need_type": "order_for_other",
                    "total": 38850,
                    // ... detail pesanan yang dibalas
                }
            }
        }
    ]
}
```