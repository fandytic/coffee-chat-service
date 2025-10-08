# 19. Menu Management (CRUD)

Kumpulan endpoint ini digunakan oleh admin untuk mengelola menu makanan dan minuman. Semua endpoint di bawah ini **terproteksi** dan memerlukan token otentikasi admin.

---

### 1. Create Menu

Membuat item menu baru. Gunakan API `/upload-image` terlebih dahulu untuk mendapatkan `image_url`.

- **Endpoint**: `POST /admin/menus`
- **Request Body**:
  ```json
  {
      "name": "Kopi Susu Gula Aren",
      "price": 22000,
      "image_url": "/public/uploads/12345_kopi.jpg"
  }
  ```
- **Success Response (201)**:
  ```json
  {
      "success": true,
      "code": 201,
      "message": "Menu created successfully",
      "data": {
          "ID": 1,
          "CreatedAt": "...",
          "UpdatedAt": "...",
          "DeletedAt": null,
          "name": "Kopi Susu Gula Aren",
          "price": 22000,
          "image_url": "/public/uploads/12345_kopi.jpg"
      }
  }
  ```

---

### 2. Get All Menus (with Search)

Mengambil daftar semua item menu. Bisa difilter berdasarkan nama.

- **Endpoint**: `GET /admin/menus`
- **Query Parameter (Opsional)**:
  - `search` (string): Filter menu yang namanya mengandung teks ini.
- **Contoh cURL (Pencarian)**: `curl ... '/admin/menus?search=kopi'`
- **Success Response (200)**:
  ```json
  {
      "success": true,
      "code": 200,
      "message": "Menus retrieved successfully",
      "data": [
          {
              "ID": 1,
              "CreatedAt": "...",
              "UpdatedAt": "...",
              "DeletedAt": null,
              "name": "Kopi Susu Gula Aren",
              "price": 22000,
              "image_url": "/public/uploads/12345_kopi.jpg"
          }
      ]
  }
  ```

---

### 3. Get Menu by ID

Mengambil detail satu item menu spesifik.

- **Endpoint**: `GET /admin/menus/:id`
- **Success Response (200)**: Respons berisi satu objek menu, sama seperti pada *Create Menu*.

---

### 4. Update Menu

Memperbarui detail item menu yang sudah ada.

- **Endpoint**: `PUT /admin/menus/:id`
- **Request Body**: Sama seperti *Create Menu*.
- **Success Response (200)**: Respons berisi objek menu yang sudah diperbarui.

---

### 5. Delete Menu

Menghapus item menu.

- **Endpoint**: `DELETE /admin/menus/:id`
- **Success Response (200)**:
  ```json
  {
      "success": true,
      "code": 200,
      "message": "Menu deleted successfully"
  }
  ```