# 35. Customer: Get Group Chat History

Endpoint ini digunakan untuk mengambil riwayat percakapan dari sebuah grup chat. **Endpoint ini terproteksi**.

-   **Endpoint**: `GET /customer/groups/:id/history`
-   **Authentication**: `Bearer Token` (Customer)
-   **Path Parameter**: `:id` - ID dari grup chat.

---

### Contoh cURL
```sh
curl --location 'http://localhost:8080/customer/groups/1/history' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

---
### Contoh Success Response (Code: 200)

Responsnya adalah array dari semua objek pesan untuk grup tersebut. Formatnya identik dengan riwayat chat personal (`ChatMessage`), sehingga front-end dapat menggunakan komponen render yang sama.

```json
{
    "success": true,
    "code": 200,
    "message": "Chat history retrieved successfully",
    "data": [
        {
            "message_id": 130,
            "sender_id": 2,
            "sender_name": "Edward",
            "sender_photo_url": "/public/uploads/...",
            "sender_table_number": "05",
            "sender_floor_number": 1,
            "chat_group_id": 1,
            "text": "Halo semua, ini grup baru kita!",
            "timestamp": "2025-10-07T21:00:00Z"
        },
        {
            "message_id": 131,
            "sender_id": 5,
            "sender_name": "Christine",
            "sender_photo_url": "/public/uploads/...",
            "sender_table_number": "01",
            "sender_floor_number": 1,
            "chat_group_id": 1,
            "text": "Asiiik!",
            "timestamp": "2025-10-07T21:01:00Z"
        }
    ]
}
```

#### `api_docs/4_websocket_connection.md` (Diperbarui)
Pastikan payload pengiriman untuk WebSocket sudah benar.

````markdown
# 4. WebSocket Real-time Chat
(...Deskripsi dan URL Koneksi tidak berubah...)

---
### 2. Mengirim Pesan (Client -> Server)

- **Struktur Pesan:**
  - `recipient_id` (integer, opsional): ID pelanggan tujuan (untuk chat 1-lawan-1).
  - `chat_group_id` (integer, opsional): ID grup tujuan (untuk chat grup).
  - `text` (string): Isi pesan.
  - `reply_to_message_id` (integer, opsional): ID dari pesan yang ingin dibalas.
  - `menu_id` (integer, opsional): ID menu untuk pesan traktiran.
  - `order_id` (integer, opsional): ID pesanan untuk pesan notifikasi order.

**PENTING**: Isi salah satu antara `recipient_id` (untuk privat) atau `chat_group_id` (untuk grup).

**Contoh Mengirim Pesan Grup:**
```json
{
    "chat_group_id": 1,
    "text": "Halo semua!"
}
```
(...Sisa dokumentasi tidak berubah...)