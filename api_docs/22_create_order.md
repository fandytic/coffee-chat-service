# 22. Customer: Create Order

Endpoint ini digunakan oleh pelanggan untuk membuat pesanan baru. **Endpoint ini terproteksi** dan memerlukan token otentikasi pelanggan.

Setiap pesanan secara otomatis dikaitkan dengan meja tempat pelanggan duduk sehingga perhitungan pembayaran dilakukan per-meja.

-   **Endpoint**: `POST /customer/orders`
-   **Authentication**: `Bearer Token` (Customer)

---

### Request Body

-   **`need_type`**: Jenis kebutuhan pesanan. Nilai yang tersedia:
    -   `self_order` — pelanggan memesan untuk dirinya sendiri.
    -   `order_for_other` — pelanggan memesankan untuk pelanggan lain (misal traktir teman).
    -   `request_treat` — pelanggan meminta pelanggan lain mentraktir pesanannya.
-   **`recipient_customer_id`**: ID pelanggan lain yang terlibat.
    -   Wajib diisi ketika `need_type` adalah `order_for_other` atau `request_treat`.
    -   Opsional (dan harus dikosongkan) ketika `need_type` adalah `self_order`.
-   **`items`**: Array dari objek yang berisi `menu_id` dan `quantity`.
-   **`notes`** (opsional): Catatan tambahan untuk pesanan.

```json
{
    "need_type": "order_for_other",
    "recipient_customer_id": 9,
    "notes": "French fries jangan terlalu asin ya.",
    "items": [
        {
            "menu_id": 1,
            "quantity": 1
        },
        {
            "menu_id": 7,
            "quantity": 2
        }
    ]
}
```

---
### Contoh Success Response (Code: 201)

API akan secara otomatis menghitung `sub_total`, `tax`, dan `total` berdasarkan harga menu saat itu dan mengembalikan ringkasan pesanan siap tampil di halaman checkout.

```json
{
    "success": true,
    "code": 201,
    "message": "Order created successfully",
    "data": {
        "order_id": 42,
        "customer_id": 5,
        "customer_name": "Christine Stanley",
        "table_id": 7,
        "table_number": "01",
        "table_name": "Christine Stanley",
        "need_type": "order_for_other",
        "recipient": {
            "customer_id": 9,
            "name": "Adi Wijaya",
            "table_id": 7,
            "table_number": "01"
        },
        "notes": "French fries jangan terlalu asin ya.",
        "sub_total": 60000,
        "tax": 6600,
        "total": 66600,
        "created_at": "2025-10-11T13:30:00Z",
        "items": [
            {
                "menu_id": 1,
                "menu_name": "French Fries",
                "quantity": 1,
                "unit_price": 25000,
                "total_price": 25000
            },
            {
                "menu_id": 7,
                "menu_name": "Roti Bakar",
                "quantity": 2,
                "unit_price": 17500,
                "total_price": 35000
            }
        ]
    }
}
```

> **Catatan**:
> -   Ketika `need_type` bukan `self_order`, sistem akan mengirim ringkasan pesanan melalui WebSocket chat kepada pelanggan yang dipilih.
> -   Notifikasi real-time ke admin hanya dikirim untuk pesanan `self_order` dan `order_for_other`.
> -   Untuk `request_treat`, API hanya membuat permintaan traktir dan mengirimkannya ke pelanggan yang diminta melalui chat. `order_id` akan bernilai `0` dan admin baru menerima data pesanan ketika pelanggan yang diminta menekan tombol **Bayar** (front end akan memanggil kembali endpoint ini sebagai `order_for_other`).

### Contoh Respons `request_treat`

```json
{
    "success": true,
    "code": 201,
    "message": "Order created successfully",
    "data": {
        "order_id": 0,
        "customer_id": 11,
        "customer_name": "Unknown User",
        "table_id": 15,
        "table_number": "07",
        "table_name": "Meja 07",
        "need_type": "request_treat",
        "recipient": {
            "customer_id": 5,
            "name": "Christine Stanley",
            "table_id": 7,
            "table_number": "01"
        },
        "notes": "Traktir please...",
        "sub_total": 25000,
        "tax": 2750,
        "total": 27750,
        "created_at": "2025-10-11T13:45:00Z",
        "items": [
            {
                "menu_id": 1,
                "menu_name": "French Fries",
                "quantity": 1,
                "unit_price": 25000,
                "total_price": 25000
            }
        ]
    }
}
```
