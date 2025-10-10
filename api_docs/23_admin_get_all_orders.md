# 23. Admin: Get All Orders

Endpoint ini digunakan oleh admin untuk mendapatkan daftar semua pesanan pelanggan, diurutkan dari yang terbaru. **Endpoint ini terproteksi**.

-   **Endpoint**: `GET /admin/orders`
-   **Authentication**: `Bearer Token` (Admin)

---
### Contoh Success Response (Code: 200)

Responsnya berisi array dari semua pesanan, lengkap dengan detail pelanggan dan item menu yang dipesan.

```json
{
    "success": true,
    "code": 200,
    "message": "Orders retrieved successfully",
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2025-10-11T13:30:00Z",
            "UpdatedAt": "...",
            "CustomerID": 5,
            "Total": 66600,
            "Status": "pending",
            "Notes": "French fries jangan terlalu asin ya.",
            "Customer": {
                "ID": 5,
                "Name": "Christine Stanley",
                "PhotoURL": "..."
            },
            "OrderItems": [
                {
                    "ID": 1,
                    "MenuID": 1,
                    "Quantity": 1,
                    "Price": 25000,
                    "Menu": {
                        "ID": 1,
                        "Name": "French Fries",
                        "Price": 25000,
                        "ImageURL": "..."
                    }
                }
            ]
        }
    ]
}
```