# SIAKAD — Modul 8: Beasiswa & Pendaftaran

Nama: Al Yasmin Assa'diyah  
NPM: 714240014  
Mata Kuliah: Pemrograman Web Service  
Modul: 8 — Beasiswa & Pendaftaran  

---

# Deskripsi

Aplikasi web fullstack berbasis REST API untuk pengelolaan data beasiswa dan pendaftaran mahasiswa.  
Dibangun menggunakan Go Fiber v2 sebagai backend, MongoDB Atlas sebagai database, dan Vanilla HTML/CSS/JavaScript sebagai frontend.

---

# Tech Stack

| Layer | Teknologi |
|---|---|
| Backend | Go + Go Fiber v2 |
| Database | MongoDB Atlas |
| Frontend | HTML + CSS + JavaScript (Vanilla) |
| Hosting | Alwaysdata |
| API Testing | Postman |
| Version Control | Git & GitHub |

---

# Struktur Folder

```bash
714240014/
├── config/
│   ├── config.go
│   ├── cors.go
│   └── db.go
│
├── controller/
│   ├── controller.go
│   ├── beasiswaController.go
│   └── pendaftaranController.go
│
├── helper/
│   ├── helper.go
│   └── mongodb.go
│
├── model/
│   ├── model.go
│   ├── beasiswa.go
│   └── pendaftaran.go
│
├── url/
│   ├── beasiswaRoute.go
│   ├── pendaftaranRoute.go
│   └── url.go
│
├── frontend/
│   └── beasiswa/
│       ├── index.html
│       ├── admin.html
│       ├── style.css
│       └── script.js
│
├── main.go
├── go.mod
├── go.sum
└── .gitignore
```

---

# Alur Pengerjaan

```mermaid
flowchart TD

A([Mulai]) --> B[Setup MongoDB Atlas]

B --> C[Konfigurasi .env]

C --> D[Membuat Model]

D --> E[Membuat Controller]

E --> F[Membuat Route]

F --> G[Membuat Frontend]

G --> H[Testing CRUD]

H --> I[Push GitHub]

I --> J[Create Pull Request]

J --> K([Selesai])
```

---

# Environment Variables

Buat file `.env` di root project:

```env
MONGOSTRING=mongodb+srv://username:password@cluster.mongodb.net/
MONGODB_NAME=kampus
PORT=8080
IP=127.0.0.1
JWT_SECRET=rahasia_jwt_super_aman
```

---

# Cara Menjalankan

## Install Dependency

```bash
go mod tidy
```

## Jalankan Server

```bash
go run main.go
```

## Buka Browser

```bash
http://127.0.0.1:8080/beasiswa/index.html
```

---

# Dokumentasi API

## Base URL

```bash
http://127.0.0.1:8080
```

---

# Endpoint Beasiswa

## GET /beasiswa

Mengambil semua data beasiswa.

### Response

```json
[
  {
    "_id": "6649abc",
    "nama": "Beasiswa Prestasi",
    "syarat": "IPK minimal 3.5",
    "deadline": "2026-07-20"
  }
]
```

---

## GET /beasiswa/:id

Mengambil detail beasiswa berdasarkan ID.

### Response

```json
{
  "_id": "6649abc",
  "nama": "Beasiswa Prestasi",
  "syarat": "IPK minimal 3.5",
  "deadline": "2026-07-20"
}
```

---

## POST /beasiswa

Menambahkan data beasiswa baru.

### Request Body

```json
{
  "nama": "Beasiswa Prestasi",
  "syarat": "IPK minimal 3.5",
  "deadline": "2026-07-20"
}
```

---

## PUT /beasiswa/:id

Mengupdate data beasiswa.

### Request Body

```json
{
  "nama": "Beasiswa Akademik Updated"
}
```

---

## DELETE /beasiswa/:id

Menghapus data beasiswa berdasarkan ID.

---

# Endpoint Pendaftaran

## POST /beasiswa/daftar

Mahasiswa melakukan pendaftaran beasiswa.

### Request Body

```json
{
  "nama": "Al Yasmin",
  "npm": "714240014",
  "email": "yasmin@email.com",
  "semester": "4",
  "prodi": "Teknik Informatika",
  "ipk": "3.75",
  "beasiswa": "Beasiswa Prestasi/Akademik"
}
```

---

## GET /beasiswa/status/:npm

Cek status pendaftaran berdasarkan NPM.

### Response

```json
{
  "nama": "Al Yasmin",
  "npm": "714240014",
  "beasiswa": "Beasiswa Prestasi/Akademik",
  "status": "Pending"
}
```

---

## PUT /beasiswa/status/:npm

Update status pendaftaran mahasiswa.

### Request Body

```json
{
  "status": "Accepted"
}
```

---

## DELETE /beasiswa/:npm

Menghapus data pendaftaran mahasiswa berdasarkan NPM.

---

# Frontend

| Halaman | URL | Fungsi |
|---|---|---|
| Mahasiswa | `/beasiswa/index.html` | Form pendaftaran & cek status |
| Admin | `/beasiswa/admin.html` | Update & delete data |
| CSS | `/beasiswa/style.css` | Styling halaman |
| Script | `/beasiswa/script.js` | Logic frontend |

---

# Fitur :

1. CRUD Data Beasiswa  
2. Pendaftaran Mahasiswa  
3. Update Status Pendaftaran  
4. Delete Pendaftaran  
5. Validasi Input Form  
6. Responsive UI  
7. MongoDB Atlas Integration  
8. REST API Fiber v2  

---
