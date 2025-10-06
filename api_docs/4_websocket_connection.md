# 4. WebSocket Connection

Endpoint ini digunakan untuk membuat koneksi WebSocket secara real-time. Setelah terhubung, klien akan menerima setiap pesan baru yang dikirim melalui endpoint `POST /send`.

- **Endpoint**: `GET /ws` (akan di-upgrade ke protokol WebSocket)

---

### Alamat Koneksi

Gunakan klien WebSocket untuk terhubung ke URL berikut:

```
ws://localhost:8080/ws
```

---

### Pesan yang Diterima

Setiap kali ada pesan baru, klien akan menerima data JSON seperti ini:

```json
{
    "id": 3,
    "user": "Dani",
    "text": "Ada yang lihat dokumen A?",
    "timestamp": "2025-10-07T13:46:05.789Z"
}
```