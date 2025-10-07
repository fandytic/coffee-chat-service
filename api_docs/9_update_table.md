# 9. Update Table Information

Endpoint ini digunakan untuk memperbarui nama dan posisi (koordinat x,y) dari sebuah meja. Anda perlu mengetahui ID unik dari meja yang ingin diubah. **Endpoint ini terproteksi**.

- **Endpoint**: `PUT /admin/tables/:table_id`
- **Content-Type**: `application/json`
- **Authentication**: `Bearer Token`

---

### Request Body

```json
{
    "table_name": "Meja Pojok Baru",
    "x": 150,
    "y": 300
}
```

---

### Contoh cURL

Ganti `<TABLE_ID>` dengan ID meja yang ingin diubah (misalnya `1`).

```sh
curl --location --request PUT 'http://localhost:8080/admin/tables/<TABLE_ID>' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data '{
    "table_name": "Meja Pojok Baru",
    "x": 150,
    "y": 300
}'
```

---

### Contoh Success Response (Code: 200)

```json
{
    "success": true,
    "code": 200,
    "message": "Table updated successfully",
    "data": {
        "ID": 1,
        "CreatedAt": "2025-10-07T10:00:00Z",
        "UpdatedAt": "2025-10-07T11:30:00Z",
        "DeletedAt": null,
        "TableNumber": "01",
        "TableName": "Meja Pojok Baru",
        "XCoordinate": 150,
        "YCoordinate": 300,
        "FloorID": 1
    }
}
```