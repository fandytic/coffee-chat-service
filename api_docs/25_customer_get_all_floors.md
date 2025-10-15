# 25. Customer: Get All Floors

Endpoint ini digunakan oleh pelanggan untuk mendapatkan daftar semua lantai yang telah dibuat denahnya, agar mereka dapat memilih lantai mana yang ingin dilihat. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /customer/floor-plans`
- **Authentication**: `Bearer Token` (Customer)

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/customer/floor-plans' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

---

### Contoh Success Response (Code: 200)

Responsnya adalah sebuah array yang berisi objek dengan ID dan nomor lantai.

```json
{
    "success": true,
    "code": 200,
    "message": "Floors retrieved successfully",
    "data": [
        {
            "id": 1,
            "floor_number": 1
        },
        {
            "id": 2,
            "floor_number": 2
        }
    ]
}
```