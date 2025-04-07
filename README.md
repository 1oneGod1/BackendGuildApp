# 🏰 Guild Management API

Guild Management API adalah aplikasi backend berbasis Golang menggunakan **Gin** dan **GORM** untuk mengelola data guild dan anggotanya. Aplikasi ini mendukung operasi CRUD pada guild dan anggota dalam sebuah sistem manajemen guild.

---

## 🚀 Fitur Utama
- Tambah, lihat, ubah, dan hapus guild.
- Tambah, lihat, ubah, dan hapus anggota dalam sebuah guild.
- Terintegrasi dengan **MySQL** menggunakan GORM.
- Menggunakan **Gin** sebagai web framework.
- API berjalan pada port **8080** secara default.

---

## 🛠️ Teknologi yang Digunakan
- **Golang**: Bahasa pemrograman utama.
- **Gin**: Web framework untuk REST API.
- **GORM**: ORM untuk akses database.
- **MySQL**: Database yang digunakan.

---

## ⚙️ Instalasi dan Konfigurasi

### 1. Clone Repository
```bash
git clone https://github.com/username/guild-management.git
cd guild-management
```

### 2. Buat file konfigurasi database
Pastikan MySQL sudah terpasang dan jalankan servernya.

### 3. Buat Database
```sql
CREATE DATABASE guild_management;
```

### 4. Sesuaikan Koneksi Database (config/database.go)
```go
dsn := "root:password@tcp(127.0.0.1:3306)/guild_management?charset=utf8mb4&parseTime=True&loc=Local"
```

### 5. Jalankan Aplikasi
```bash
go run main.go
```
Aplikasi akan berjalan di:
```
http://localhost:8080
```

---

## 🌐 Dokumentasi API

### 🔍 Cek API Utama
**GET /**  
- **URL:** `http://localhost:8080/`  
- **Response:**
  ```json
  {
    "message": "Welcome to Guild Management API"
  }
  ```

### 🏰 Guild Endpoints

#### 📥 Tambah Guild
**POST /guilds/**  
- **Body:**  
  ```json
  {
    "name": "Warriors",
    "description": "Guild of brave warriors"
  }
  ```
- **Response:**  
  ```json
  {
    "id": 1,
    "name": "Warriors",
    "description": "Guild of brave warriors"
  }
  ```

#### 📜 Lihat Semua Guild
**GET /guilds/**  
- **Response:**  
  ```json
  [
    {
      "id": 1,
      "name": "Warriors",
      "description": "Guild of brave warriors"
    }
  ]
  ```

#### 🔍 Lihat Guild Berdasarkan ID
**GET /guilds/:id**  
- **URL:** `http://localhost:8080/guilds/1`  
- **Response:**
  ```json
  {
    "id": 1,
    "name": "Warriors",
    "description": "Guild of brave warriors"
  }
  ```

#### 📝 Ubah Guild
**PUT /guilds/:id**  
- **Body:**  
  ```json
  {
    "name": "Warriors Updated",
    "description": "Guild of legendary warriors"
  }
  ```
- **Response:**  
  ```json
  {
    "id": 1,
    "name": "Warriors Updated",
    "description": "Guild of legendary warriors"
  }
  ```

#### 🗑️ Hapus Guild
**DELETE /guilds/:id**  
- **Response:**  
  ```json
  {
    "message": "Guild deleted successfully"
  }
  ```

---

### 👥 Member Endpoints

#### 📥 Tambah Member ke Guild
**POST /guilds/:id/members/**  
- **Body:**  
  ```json
  {
    "name": "Arthur",
    "role": "Knight"
  }
  ```
- **Response:**  
  ```json
  {
    "id": 1,
    "name": "Arthur",
    "role": "Knight",
    "guild_id": 1
  }
  ```

#### 📜 Lihat Semua Anggota dalam Guild
**GET /guilds/:id/members/**  
- **Response:**  
  ```json
  [
    {
      "id": 1,
      "name": "Arthur",
      "role": "Knight"
    }
  ]
  ```

#### 🔍 Lihat Detail Anggota
**GET /guilds/:id/members/:member_id**  
- **Response:**  
  ```json
  {
    "id": 1,
    "name": "Arthur",
    "role": "Knight"
  }
  ```

#### 🔄 Ubah Anggota
**PUT /guilds/:id/members/:member_id**  
- **Body:**  
  ```json
  {
    "name": "Arthur Pendragon",
    "role": "King"
  }
  ```
- **Response:**  
  ```json
  {
    "id": 1,
    "name": "Arthur Pendragon",
    "role": "King"
  }
  ```

#### ❌ Hapus Anggota
**DELETE /guilds/:id/members/:member_id**  
- **Response:**  
  ```json
  {
    "message": "Member deleted successfully"
  }
  ```

---

Dokumentasi Gambar yang dijalankan menggunakan postman

#### 1. Create Guild
![image](https://github.com/user-attachments/assets/bdabee20-e470-4d77-a8ca-b7976983ebf7)

#### 2. Cek API Utama
![image](https://github.com/user-attachments/assets/3d2f4c49-7f9f-4829-aa8d-2b99a5d3157f)

#### 3. Lihat Semua Guild
![image](https://github.com/user-attachments/assets/185fe460-fed4-4eea-bb17-b2938e2a7da7)

#### 4. Update Guild 
![image](https://github.com/user-attachments/assets/92aca943-c516-493e-8501-b82350bd1a15)

#### 5. Lihat Detail Guild Berdasarkan ID
![image](https://github.com/user-attachments/assets/1f535c1f-6064-4731-b4e1-4cf6358cb4bd)

#### 6. Hapus Guild
![image](https://github.com/user-attachments/assets/5b48f2ec-02b1-4890-8a26-187c8a8b0bb8)

#### 7. Tambah Member ke Guild
![image](https://github.com/user-attachments/assets/b6bb1101-743c-4165-b3d7-7274d949faad)

#### 8. Lihat Semua memberi dalam Guild
![image](https://github.com/user-attachments/assets/61ad736a-9d46-4d90-adbd-70bd41c64c64)

#### 9. Update Member
![image](https://github.com/user-attachments/assets/5a8a1444-ce34-4dea-9fb2-4ea3d028e74a)

#### 10. Hapus Member
![image](https://github.com/user-attachments/assets/c073af76-dd8e-4e52-a401-67ce7b16a35e)




