# 21. Customer: Get All Menus

Endpoint ini digunakan oleh pelanggan untuk mendapatkan daftar semua menu makanan yang tersedia. Data bisa dicari berdasarkan nama menu. **Endpoint ini terproteksi** dan memerlukan token otentikasi pelanggan.

-   **Endpoint**: `GET /customer/menus`
-   **Authentication**: `Bearer Token` (Customer)

***

### Query Parameter (Opsional)

-   **`search`** (string): Filter daftar menu yang namanya mengandung teks ini.

### Contoh cURL

**Mengambil semua menu:**
```sh
curl --location 'http://localhost:8080/customer/menus' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

**Mencari menu dengan nama "chicken":**
```sh
curl --location 'http://localhost:8080/customer/menus?search=chicken' \
--header 'Authorization: Bearer <CUSTOMER_TOKEN>'
```

***

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Menus retrieved successfully",
    "data": [
        {
            "ID": 2,
            "CreatedAt": "...",
            "UpdatedAt": "...",
            "DeletedAt": null,
            "name": "Chicken Wings",
            "price": 35000,
            "image_url": "/public/uploads/12345_wings.jpg"
        },
        {
            "ID": 1,
            "CreatedAt": "...",
            "UpdatedAt": "...",
            "DeletedAt": null,
            "name": "French Fries",
            "price": 25000,
            "image_url": "/public/uploads/12345_fries.jpg"
        }
    ]
}
```