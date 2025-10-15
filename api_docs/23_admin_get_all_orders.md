# 23. Admin: Get All Orders

Endpoint ini digunakan oleh admin untuk mendapatkan daftar semua pesanan pelanggan, diurutkan dari yang terbaru. **Endpoint ini terproteksi**.

-   **Endpoint**: `GET /admin/orders`
-   **Authentication**: `Bearer Token` (Admin)

---
### Contoh Success Response (Code: 200)

Responsnya berisi array dari semua pesanan, lengkap dengan detail pelanggan, meja, tujuan pesanan (`need_type`), dan (jika ada) pelanggan penerima pesanan/traktiran.

```json
{
    "success": true,
    "code": 200,
    "message": "Orders retrieved successfully",
    "data": [
        {
            "ID": 42,
            "CreatedAt": "2025-10-11T13:30:00Z",
            "UpdatedAt": "...",
            "CustomerID": 5,
            "Total": 66600,
            "Tax": 6600,
            "SubTotal": 60000,
            "Status": "pending",
            "Notes": "French fries jangan terlalu asin ya.",
            "NeedType": "order_for_other",
            "RecipientID": 9,
            "TableID": 7,
            "Table": {
                "ID": 7,
                "TableNumber": "01",
                "TableName": "Christine Stanley"
            },
            "Customer": {
                "ID": 5,
                "Name": "Christine Stanley",
                "PhotoURL": "...",
                "TableID": 3,
                "Table": {
                    "ID": 3,
                    "TableNumber": "01",
                    "TableName": "Christine Stanley"
                }
            },
            "Recipient": {
                "ID": 9,
                "Name": "Adi Wijaya",
                "PhotoURL": "...",
                "TableID": 7,
                "Table": {
                    "ID": 7,
                    "TableNumber": "01",
                    "TableName": "Christine Stanley"
                }
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
