# 12. Delete Floor

Endpoint ini digunakan untuk menghapus data denah lantai, semua data meja yang ada di lantai tersebut, dan file gambar denahnya dari server. Operasi ini bersifat permanen dan tidak dapat diurungkan. **Endpoint ini terproteksi**.

- **Endpoint**: `DELETE /floor-plans/:floor_id`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

Ganti `<FLOOR_ID>` dengan ID lantai yang ingin dihapus (misalnya `1`).

```sh
curl --location --request DELETE 'http://localhost:8080/floor-plans/<FLOOR_ID>' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Floor and its associated data deleted successfully"
}
```

---

### Contoh Error Response (Code: 404 Not Found)

Jika lantai dengan ID yang diberikan tidak ditemukan.

```json
{
    "success": false,
    "code": 404,
    "message": "floor with ID 999 not found"
}
```