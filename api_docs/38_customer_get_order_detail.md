# 38. Customer: Get Order Detail

Endpoint ini digunakan oleh pelanggan yang sedang login untuk mengambil detail lengkap dari sebuah pesanan. Endpoint ini dirancang untuk menampilkan layar rincian penuh pesanan (seperti yang terlihat pada gambar UI Anda).

Untuk alasan keamanan, pelanggan hanya dapat melihat detail pesanan jika mereka adalah **Pemesan (`CustomerID`)** atau **Penerima (`RecipientID`)** atau **Pembayar (`PayerCustomerID`)** pesanan tersebut.

- **Endpoint**: `GET /customer/orders/:id`
- **Authentication**: `Bearer Token` (Customer)
- **Path Parameter**: `:id` - ID dari pesanan yang ingin dilihat detailnya.

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/orders/42' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

### Contoh Success Response 1 (Code: 200 - Self Order)
```json
{
    "success": true,
    "code": 200,
    "message": "Order detail retrieved successfully",
    "data": {
        "ID": 42,
        "CreatedAt": "2025-12-04T10:30:00Z",
        "UpdatedAt": "2025-12-04T10:30:00Z",
        "CustomerID": 5,
        "PayerCustomerID": null,
        "TableID": 7,
        "NeedType": "self_order",
        "RecipientID": null,
        "Total": 154000.0,
        "Tax": 14000.0,
        "SubTotal": 140000.0,
        "Status": "processing",
        "Notes": "Jangan pake sambel dan bawang goreng sama minta mangkok kosong 2",
        "OrderItems": [
            {
                "ID": 1,
                "MenuID": 1,
                "Quantity": 1,
                "Price": 25000,
                "Menu": { "ID": 1, "Name": "French Fries", "Price": 25000, "ImageURL": "/public/uploads/fries.jpg" }
            },
            {
                "ID": 2,
                "MenuID": 7,
                "Quantity": 2,
                "Price": 15000,
                "Menu": { "ID": 7, "Name": "Kebab", "Price": 15000, "ImageURL": "/public/uploads/kebab.jpg" }
            },
            {
                "ID": 3,
                "MenuID": 10,
                "Quantity": 2,
                "Price": 15000,
                "Menu": { "ID": 10, "Name": "Roti Bakar", "Price": 15000, "ImageURL": "/public/uploads/roti.jpg" }
            }
            // ... item lainnya
        ],
        "Customer": {
            "ID": 5,
            "Name": "Pamela Schneider",
            "PhotoURL": "/public/uploads/pamela.jpg",
            "TableID": 3
        },
        "Table": {
            "ID": 7,
            "TableNumber": "01",
            "TableName": "Christine Stanley"
        },
        "Recipient": null,
        "Payer": null
    }
}
```

### Contoh Success Response 2 (Code: 200 - Order for Other)
```json
{
    "success": true,
    "code": 200,
    "message": "Order detail retrieved successfully",
    "data": {
        "ID": 45,
        "CreatedAt": "2025-12-05T14:00:00Z",
        "UpdatedAt": "2025-12-05T14:00:00Z",
        "CustomerID": 9,
        "PayerCustomerID": null,
        "TableID": 5,
        "NeedType": "order_for_other",
        "RecipientID": 5,
        "Total": 38850.0,
        "Tax": 3850.0,
        "SubTotal": 35000.0,
        "Status": "pending",
        "Notes": "Tolong cepat diantar ke meja 3",
        "OrderItems": [
            {
                "ID": 6,
                "MenuID": 12,
                "Quantity": 1,
                "Price": 35000,
                "Menu": { "ID": 12, "Name": "Cappuccino", "Price": 35000, "ImageURL": "/public/uploads/capp.jpg" }
            }
        ],
        "Customer": {
            "ID": 9,
            "Name": "Adi Wijaya",
            "PhotoURL": "/public/uploads/adi.jpg",
            "TableID": 5
        },
        "Table": {
            "ID": 5,
            "TableNumber": "03",
            "TableName": "Adi Wijaya"
        },
        "Recipient": {
            "ID": 5,
            "Name": "Christine Stanley",
            "PhotoURL": "/public/uploads/christine.jpg",
            "TableID": 3
        },
        "Payer": null
    }
} 
```

### Contoh Error Response (Code: 403 Forbidden)
```json
{
    "success": false,
    "code": 403,
    "message": "forbidden: you do not own this order"
}
```