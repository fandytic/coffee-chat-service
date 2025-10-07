# 13. Admin Logout

Endpoint ini digunakan untuk proses logout admin. **Endpoint ini terproteksi**.

**Alur Kerja Penting**: Setelah memanggil API ini dan mendapatkan respons sukses, **klien (front-end) wajib menghapus Bearer Token** yang tersimpan di *local storage* atau *cookies*.

- **Endpoint**: `POST /logout`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

```sh
curl --location --request POST 'http://localhost:8080/logout' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Logout successful"
}
```