# 26. Penjelasan: Manajemen Status Pelanggan Otomatis

Dokumen ini menjelaskan bagaimana sistem secara otomatis mengelola status `online` (`active`) dan `offline` (`inactive`) pelanggan berdasarkan masa berlaku sesi mereka.

### Kapan Status Menjadi `active`?
-   Seorang pelanggan mendapatkan status `active` ketika mereka berhasil melakukan **check-in** melalui API `POST /check-in`. Status ini akan tetap `active` selama sesi (token) mereka masih valid.

### Kapan Status Menjadi `inactive`?
-   **Pembersihan Sesi Kedaluwarsa**: Server memiliki tugas otomatis yang berjalan setiap **10 menit**. Tugas ini akan mencari semua pelanggan yang statusnya masih `active` tetapi sesi (token) mereka seharusnya sudah kedaluwarsa (lebih dari 8 jam sejak aktivitas terakhir). Pelanggan yang ditemukan akan diubah statusnya menjadi `inactive`.

### Dampak
-   Seorang pelanggan dapat menutup browser dan membukanya kembali kapan saja selama token mereka masih valid (kurang dari 8 jam) dan mereka akan tetap dianggap `active`.
-   API `GET /customer/active-list` dan API dasbor akan selalu menampilkan data pelanggan yang sesinya masih valid.
-   Pelanggan yang sesinya sudah benar-benar berakhir tidak akan muncul lagi di daftar, mencegah "ghost users" dan potensi *error*.