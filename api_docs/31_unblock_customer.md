# 31. Customer: Unblock Customer
Endpoint ini digunakan oleh pelanggan yang login untuk membuka blokir pelanggan lain. **Endpoint ini terproteksi**.

-   **Endpoint**: `POST /customer/unblock/:id`
-   **Authentication**: `Bearer Token` (Customer)
-   **Path Parameter**: `:id` - ID pelanggan yang ingin Anda buka blokirnya.

### Contoh cURL
```sh
curl --location --request POST 'http://localhost:8080/customer/unblock/5' \
--header 'Authorization: Bearer <YOUR_CUSTOMER_TOKEN>'
```

### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Customer unblocked successfully"
}
```