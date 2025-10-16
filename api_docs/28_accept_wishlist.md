# 28. Customer: Accept a Wishlist
- **Endpoint**: `POST /customer/wishlists/:id/accept`
- **Deskripsi**: Pelanggan yang login "mengambil" dan setuju untuk membayar sebuah wishlist. Ini akan mengubah status wishlist menjadi pesanan `pending` dan mengirim notifikasi ke admin.
- **Respons**: Objek Order yang sudah diperbarui dengan status `pending` dan `PayerCustomerID` yang terisi.