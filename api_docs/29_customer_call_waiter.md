# 29. Customer: Call Waiter

Endpoint ini digunakan oleh pelanggan untuk memanggil waiter ke meja mereka. Panggilan ini akan dikirim secara *real-time* ke dasbor admin. **Endpoint ini terproteksi**.

-   **Endpoint**: `POST /customer/call-waiter`
-   **Authentication**: `Bearer Token` (Customer)

### Request Body
Tidak ada *body* yang diperlukan. Server akan mengambil detail pelanggan dari token.

---
### Contoh cURL

```sh
curl --location --request POST 'http://localhost:8080/customer/call-waiter' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

---
### Contoh Success Response (Code: 200)
```json
{
    "success": true,
    "code": 200,
    "message": "Waiter has been called"
}
```