# 7. Get Floor Plan by Floor Number

Endpoint ini digunakan untuk mengambil detail denah lantai, termasuk URL gambar dan data semua meja di lantai tersebut. **Endpoint ini terproteksi**.

- **Endpoint**: `GET /floor-plans/:floor_number`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

```sh
curl --location 'http://localhost:8080/floor-plans/1' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)

Sama seperti respons saat membuat denah baru.