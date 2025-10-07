# 5. Generate QR Code

Endpoint ini digunakan untuk membuat gambar QR code dari teks atau URL yang diberikan. **Endpoint ini terproteksi** dan memerlukan otentikasi Bearer Token.

- **Endpoint**: `POST /admin/generate-qr`
- **Content-Type**: `application/json`
- **Authentication**: `Bearer Token`

---

### Request Header

```
Authorization: Bearer <token_hasil_login>
```

---

### Request Body

```json
{
    "content": "[https://www.google.com](https://www.google.com)"
}
```

---

### Contoh cURL

Ganti `<TOKEN>` dengan token yang Anda dapatkan dari endpoint `/login`.

```sh
curl --location 'http://localhost:8080/admin/generate-qr' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data '{
    "content": "Teks ini akan menjadi QR Code"
}' \
--output qrcode.png
```
**Tips**: Perintah `curl` di atas akan menyimpan output gambar langsung ke file bernama `qrcode.png`.

---

### Contoh Success Response (Code: 200)

Respons dari API ini bukanlah JSON, melainkan **data gambar mentah** (`image/png`). Jika Anda menggunakan Postman atau browser, gambar QR code akan langsung ditampilkan.



[Image of a black and white QR code]