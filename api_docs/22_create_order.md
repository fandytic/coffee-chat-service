# 22. Customer: Create Order

Endpoint ini digunakan oleh pelanggan untuk membuat pesanan baru. **Endpoint ini terproteksi** dan memerlukan token otentikasi pelanggan.

-   **Endpoint**: `POST /customer/orders`
-   **Authentication**: `Bearer Token` (Customer)

---

### Request Body

-   **`items`**: Array dari objek yang berisi `menu_id` dan `quantity`.
-   **`notes`** (opsional): Catatan tambahan untuk pesanan.

```json
{
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

API akan secara otomatis menghitung `sub_total`, `tax`, dan `total` berdasarkan harga menu saat itu.

```json
{
    "success": true,
    "code": 201,
    "message": "Order created successfully",
    "data": {
        "ID": 1,
        "CreatedAt": "2025-10-11T13:30:00Z",
        "UpdatedAt": "2025-10-11T13:30:00Z",
        "DeletedAt": null,
        "CustomerID": 5,
        "Total": 66600,
        "Tax": 6600,
        "SubTotal": 60000,
        "Status": "pending",
        "Notes": "French fries jangan terlalu asin ya.",
        "OrderItems": [
            {
                "ID": 1,
                "CreatedAt": "...",
                "OrderID": 1,
                "MenuID": 1,
                "Quantity": 1,
                "Price": 25000
            },
            {
                "ID": 2,
                "CreatedAt": "...",
                "OrderID": 1,
                "MenuID": 7,
                "Quantity": 2,
                "Price": 17500
            }
        ],
        "Customer": null
    }
}
```