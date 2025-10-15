# 4. WebSocket Real-time Chat

Endpoint ini digunakan untuk membangun koneksi *real-time* antar pelanggan yang sudah melakukan *check-in*. Sistem ini memungkinkan pengiriman pesan pribadi dari satu pelanggan ke pelanggan lainnya.

---

### 1. URL Koneksi & Otentikasi

Untuk terhubung, klien harus menyediakan token otentikasi pelanggan (didapat dari API `POST /check-in`) sebagai *query parameter*.

- **URL**: `ws://localhost:8080/ws?token=<CUSTOMER_AUTH_TOKEN>`

**Penting**: Jika `token` tidak disediakan atau tidak valid, koneksi akan ditolak oleh server.

---

### 2. Mengirim Pesan (Client -> Server)

Untuk mengirim pesan, klien mengirimkan format JSON berikut. Untuk membalas pesan, sertakan `reply_to_message_id`.

- **Struktur Pesan:**
  - `recipient_id` (integer): ID pelanggan tujuan.
  - `text` (string): Isi pesan.
  - `reply_to_message_id` (integer, opsional): ID dari pesan yang ingin dibalas.
  - `menu_id` (integer, opsional): Sertakan ID menu jika ini adalah permintaan "traktir".

**Contoh Mengirim Pesan Biasa:**
```json
{
    "recipient_id": 5,
    "text": "Halo, salam kenal juga!",
    "menu_id": 1
}
```

**Contoh Membalas Pesan (yang memiliki ID 123):**
```json
{
    "recipient_id": 5,
    "text": "Haha mirip dikit",
    "reply_to_message_id": 123
}
```

---
### 3. Menerima Pesan (Server -> Client)

Klien akan menerima pesan dalam format JSON yang kaya akan informasi.

- **Struktur Pesan:**
  - `message_id` (integer): ID unik dari pesan ini.
  - `sender_id` (integer): ID pengirim.
  - `sender_name` (string): Nama pengirim.
  - `sender_photo_url` (string): URL foto pengirim.
  - `sender_table_number` (string): Nomor meja pengirim.
  - `sender_floor_number` (integer): **Nomor lantai pengirim.**
  - `text` (string): Isi pesan.
  - `timestamp` (string): Waktu pesan dibuat (format ISO 8601).
  - `reply_to` (objek, opsional): Berisi detail pesan yang dibalas.
    - `id` (integer): ID pesan asli.
    - `text` (string): Teks pesan asli.
    - `sender_name` (string): Nama pengirim pesan asli.
    - `menu` (objek, opsional): Detail menu jika pesan asli adalah traktiran.
      - `id` (integer): ID menu.
      - `name` (string): Nama menu.
      - `price` (float): Harga menu.
      - `image_url` (string): URL gambar menu.
  - `menu` (objek, opsional): Berisi detail menu jika ini adalah pesan traktir.
    - `id` (integer): ID menu.
    - `name` (string): Nama menu.
    - `price` (float): Harga menu.
    - `image_url` (string): URL gambar menu.
- `order` (objek, opsional): Ringkasan pesanan ketika pesan tercipta dari proses "order untuk orang lain" atau "request treat".
    - `id` (integer): ID pesanan.
    - `customer_id` (integer): ID pelanggan pembuat pesanan.
    - `recipient_id` (integer, opsional): ID pelanggan penerima traktiran.
    - `table_id` (integer): ID meja yang digunakan untuk pesanan.
    - `table_number` (string): Nomor meja.
    - `table_name` (string): Nama meja (biasanya sama dengan nama pelanggan pada meja individual).
    - `table_floor_number` (integer): Nomor lantai meja.
    - `need_type` (string): Jenis kebutuhan pesanan (`order_for_other`, `request_treat`, dll.).
    - `notes` (string, opsional): Catatan khusus dari pembuat pesanan.
    - `sub_total` (float): Total harga sebelum pajak.
    - `tax` (float): Nilai pajak yang diterapkan (11%).
    - `total` (float): Total akhir setelah pajak.
    - `order_items` (array): Detail item pesanan.
      - `id` (integer): ID item pesanan.
      - `menu_id` (integer): ID menu.
      - `quantity` (integer): Jumlah menu yang dipesan.
      - `price` (float): Harga satuan yang dikenakan pada saat pemesanan.
      - `menu` (objek, opsional): Informasi menu lengkap, termasuk `image_url` untuk ditampilkan di UI.

**Contoh Menerima Pesan Biasa:**
```json
{
    "message_id": 124,
    "sender_id": 5,
    "sender_name": "Christine",
    "sender_photo_url": "/public/uploads/...",
    "sender_table_number": "01",
    "text": "gw Christine",
    "timestamp": "2025-10-07T20:00:00Z",
    "menu": {
        "id": 1,
        "name": "French Fries",
        "price": 25000,
        "image_url": "/public/uploads/12345_fries.jpg"
    }
}
```

**Contoh Menerima Pesan Balasan:**
```json
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
```
Front-end kemudian dapat menggunakan `sender_id` untuk menampilkan nama dan foto pengirim yang sesuai dari daftar pelanggan aktif.

---

### 4. Alur Kerja Lengkap (Contoh Kasus)

1.  **Andi Check-in**: Andi (ID 10) melakukan check-in di meja 3 dan mendapatkan token `token_andi`.
2.  **Budi Check-in**: Budi (ID 12) melakukan check-in di meja 5 dan mendapatkan token `token_budi`.
3.  **Koneksi**:
    -   Aplikasi Andi terhubung ke `ws://.../ws?token=token_andi`.
    -   Aplikasi Budi terhubung ke `ws://.../ws?token=token_budi`.
4.  **Andi Mengirim Pesan**: Aplikasi Andi mengirim pesan JSON ke WebSocket:
    ```json
    {
      "recipient_id": 12,
      "text": "Bro, join sini lah..."
    }
    ```
5.  **Budi Menerima Pesan**: Server merutekan pesan tersebut, dan aplikasi Budi menerima JSON:
    ```json
    {
      "sender_id": 10,
      "text": "Bro, join sini lah..."
    }
    ```

---
### 5. Notifikasi Real-time untuk Admin

Admin yang terhubung ke WebSocket akan menerima notifikasi *real-time* untuk event-event tertentu.

#### Pesanan Baru

Ketika seorang pelanggan membuat pesanan baru, admin akan menerima pesan dengan format:

```json
{
    "type": "NEW_ORDER",
    "data": {
        // ... (Objek pesanan lengkap seperti pada response GET /admin/orders)
    }
}
```