# 4. WebSocket Real-time Chat

Endpoint ini digunakan untuk membangun koneksi *real-time* antar pelanggan yang sudah melakukan *check-in*. Sistem ini memungkinkan pengiriman pesan pribadi dari satu pelanggan ke pelanggan lainnya.

---

### 1. URL Koneksi & Otentikasi

Untuk terhubung, klien harus menyediakan token otentikasi pelanggan (didapat dari API `POST /check-in`) sebagai *query parameter*.

- **URL**: `ws://localhost:8080/ws?token=<CUSTOMER_AUTH_TOKEN>`

**Penting**: Jika `token` tidak disediakan atau tidak valid, koneksi akan ditolak oleh server.

---

### 2. Mengirim Pesan (Client -> Server)

Untuk mengirim pesan ke pelanggan lain, klien harus mengirim pesan dalam format JSON dengan struktur sebagai berikut:

- **Struktur Pesan:**
  - `recipient_id` (integer): ID unik dari pelanggan yang akan menerima pesan.
  - `text` (string): Isi pesan yang ingin dikirim.

**Contoh JSON yang Dikirim:**
Misalnya, Anda (dengan ID 1) ingin mengirim pesan ke "Christine" (dengan ID 5):

```json
{
    "recipient_id": 5,
    "text": "Halo, salam kenal juga!"
}
```

---

### 3. Menerima Pesan (Server -> Client)

Ketika seseorang mengirimi Anda pesan, server akan meneruskan pesan tersebut dalam format JSON dengan struktur yang kaya akan informasi pengirim:

- **Struktur Pesan:**
  - `sender_id` (integer): ID unik dari pelanggan yang mengirim pesan.
  - `sender_name` (string): Nama pengirim.
  - `sender_photo_url` (string): URL foto profil pengirim.
  - `sender_table_number` (string): Nomor meja tempat pengirim berada.
  - `text` (string): Isi pesan yang diterima.

**Contoh JSON yang Diterima:**
Jika "Christine" (ID 5, di meja "01") membalas pesan Anda:

```json
{
    "sender_id": 5,
    "sender_name": "Christine Stanley",
    "sender_photo_url": "/public/uploads/12345_christine.jpg",
    "sender_table_number": "01",
    "text": "Iya, salam kenal :)"
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