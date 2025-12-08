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
```

### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Order history retrieved successfully",
    "data": [
        {
            "id": 42,
            "order_number": "ORD-42",
            "created_at": "2025-12-04T10:30:00Z",
            "status": "pending",
            "need_type": "self_order",
            "total": 60000,
            "item_count": 2,
            "preview_items": [
                {
                    "menu_id": 1,
                    "menu_name": "French Fries",
                    "quantity": 1,
                    "unit_price": 25000,
                    "total_price": 25000
                },
                {
                    "menu_id": 2,
                    "menu_name": "Iced Tea",
                    "quantity": 1,
                    "unit_price": 35000,
                    "total_price": 35000
                }
            ],
            "recipient": null
        },
        {
            "id": 40,
            "order_number": "ORD-40",
            "created_at": "2025-12-03T18:15:00Z",
            "status": "completed",
            "need_type": "order_for_other",
            "total": 154000,
            "item_count": 5,
            "preview_items": [
                {
                    "menu_id": 5,
                    "menu_name": "Beef Burger",
                    "quantity": 2,
                    "unit_price": 45000,
                    "total_price": 90000
                },
                {
                    "menu_id": 1,
                    "menu_name": "French Fries",
                    "quantity": 2,
                    "unit_price": 25000,
                    "total_price": 50000
                },
                {
                    "menu_id": 9,
                    "menu_name": "Vanilla Latte",
                    "quantity": 1,
                    "unit_price": 14000,
                    "total_price": 14000
                }
            ],
            "recipient": {
                "customer_id": 9,
                "name": "Adi Wijaya",
                "table_number": "05"
            }
        }
    ]
}
```
