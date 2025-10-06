# 10. Delete Table

Endpoint ini digunakan untuk menghapus data meja dari denah. Operasi ini bersifat *soft delete*, artinya data tidak benar-benar hilang dari database tetapi ditandai sebagai telah dihapus. **Endpoint ini terproteksi**.

- **Endpoint**: `DELETE /tables/:table_id`
- **Authentication**: `Bearer Token`

---

### Contoh cURL

Ganti `<TABLE_ID>` dengan ID meja yang ingin dihapus (misalnya `1`).

```sh
curl --location --request DELETE 'http://localhost:8080/tables/<TABLE_ID>' \
--header 'Authorization: Bearer <TOKEN>'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Table deleted successfully"
}
```
*Catatan: Respons sukses untuk operasi `DELETE` biasanya tidak menyertakan `data`.*

---

### Contoh Error Response (Code: 404 Not Found)

Jika meja dengan ID yang diberikan tidak ditemukan.

```json
{
    "success": false,
    "code": 404,
    "message": "table with ID 999 not found"
}
```