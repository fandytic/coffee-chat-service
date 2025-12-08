# 37. Customer: Get My Order History

Endpoint ini digunakan oleh pelanggan yang sedang login untuk mengambil daftar riwayat semua pesanan yang mereka buat (`customer_id`) atau bayar (`payer_customer_id`, dalam kasus menerima wishlist). Data ini cocok untuk ditampilkan di menu "My Order" di aplikasi.

Endpoint ini mengembalikan ringkasan pesanan, termasuk status, total harga, dan **preview** (cuplikan) dari item menu yang dipesan (maksimal 3 item pertama) untuk kebutuhan tampilan kartu.

- **Endpoint**: `GET /customer/orders`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/orders' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'