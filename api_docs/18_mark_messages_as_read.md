# 18. Mark Messages as Read

Endpoint ini digunakan untuk menandai semua pesan yang belum dibaca dari satu pengirim spesifik sebagai "telah dibaca". Panggil endpoint ini ketika seorang pengguna membuka jendela chat dengan pengguna lain. **Endpoint ini terproteksi**.

- **Endpoint**: `POST /customer/chats/:sender_id/mark-as-read`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

Misalnya, Anda (penerima) ingin menandai semua pesan dari Mary (dengan ID 1) sebagai telah dibaca.

```sh
curl --location --request POST 'http://localhost:8080/customer/chats/1/mark-as-read' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Messages marked as read"
}
```