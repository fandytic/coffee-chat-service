# 30. Customer: Block Customer
Endpoint ini digunakan oleh pelanggan yang login untuk memblokir pelanggan lain. Setelah diblokir, Anda tidak akan lagi menerima pesan chat dari pelanggan tersebut. **Endpoint ini terproteksi**.

-   **Endpoint**: `POST /customer/block/:id`
-   **Authentication**: `Bearer Token` (Customer)
-   **Path Parameter**: `:id` - ID pelanggan yang ingin Anda blokir.

### Contoh cURL
```sh
curl --location --request POST 'http://localhost:8080/customer/block/5' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Customer blocked successfully"
}
```