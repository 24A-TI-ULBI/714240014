# Product Requirements Document (PRD)
## Portal Informasi Akademik Kampus

---

## 1. Ringkasan Produk

**Portal Informasi Akademik Kampus** adalah aplikasi web fullstack berbasis REST API yang memudahkan pengelolaan data akademik kampus. Aplikasi ini dibangun menggunakan **Go Fiber** sebagai backend dan **MongoDB** sebagai database, di-deploy di **Alwaysdata** dengan CI/CD otomatis via GitHub Actions.

---

## 2. Tujuan

- Menyediakan sistem pengelolaan data akademik yang terpusat
- Menjadi media latihan implementasi REST API + Frontend secara end-to-end
- Setiap mahasiswa berkontribusi membangun modul secara mandiri namun tetap terintegrasi dalam satu aplikasi

---

## 3. Tech Stack

| Layer | Teknologi |
|---|---|
| Backend | Go + Go Fiber v2 |
| Database | MongoDB Atlas |
| Frontend | HTML + CSS + JS (vanilla) |
| Hosting | Alwaysdata (free for life) |
| CI/CD | GitHub Actions |
| Boilerplate | github.com/gocroot/alwaysdata |

---

## 4. Arsitektur Sistem

```mermaid
graph TD
    Dev["Mahasiswa (Local Dev)"]
    GH["GitHub Repository\n(nama repo = NPM)"]
    GA["GitHub Actions\nCI/CD Pipeline"]
    AD["Alwaysdata Hosting"]

    subgraph App["Go Fiber Application"]
        URL["url/\n(Router)"]
        CTRL["controller/\n(Handler)"]
        HLP["helper/\n(Business Logic)"]
        MDL["model/\n(Struct)"]
        CFG["config/\n(DB, CORS, Env)"]
    end

    DB["MongoDB Atlas"]
    FE["Frontend\n(HTML/CSS/JS)"]
    USER["User / Browser"]

    Dev -->|git push| GH
    GH -->|trigger| GA
    GA -->|SCP binary + SSH restart| AD
    AD --> App
    URL --> CTRL
    CTRL --> HLP
    CTRL --> MDL
    HLP --> CFG
    CFG --> DB
    USER --> FE
    FE -->|HTTP Request| URL
    CTRL -->|JSON Response| FE
```

---

## 5. Fitur Aplikasi

### 5.1 Fitur Global (semua mahasiswa wajib implementasi)

| Fitur | Endpoint | Deskripsi |
|-------|----------|-----------|
| Homepage | `GET /` | Menampilkan nama aplikasi dan status server |
| IP Server | `GET /ip` | Menampilkan IP address server yang sedang berjalan |
| CORS | — | Semua endpoint bisa diakses dari frontend |
| Auto-deploy | — | Setiap `git push` ke `main` otomatis deploy ke Alwaysdata |

---

### 5.2 Fitur Per Modul

Setiap mahasiswa membangun **2 modul Backend + 1 menu Frontend** yang saling berkaitan.

---

#### Modul 1 — Data Mahasiswa & Autentikasi
**Deskripsi:** Mengelola data profil mahasiswa dan sistem autentikasi berbasis nomor telepon.

Backend 1 — **Mahasiswa**
- `GET /mahasiswa` → ambil semua data mahasiswa
- `GET /mahasiswa/:npm` → ambil data mahasiswa berdasarkan NPM
- `POST /mahasiswa` → tambah data mahasiswa baru
- `PUT /mahasiswa/:npm` → update data mahasiswa
- `DELETE /mahasiswa/:npm` → hapus data mahasiswa

Backend 2 — **Auth**
- `POST /auth/login` → login menggunakan nomor telepon, return token
- `GET /auth/profile/:phone` → ambil profil berdasarkan nomor telepon

Frontend — **Halaman Data Mahasiswa**
- Tabel list semua mahasiswa
- Form tambah mahasiswa baru
- Tombol hapus data

---

#### Modul 2 — Data Dosen & Jabatan
**Deskripsi:** Mengelola data dosen beserta jabatan fungsional dan struktural.

Backend 1 — **Dosen**
- `GET /dosen` → ambil semua data dosen
- `GET /dosen/:nidn` → ambil data dosen berdasarkan NIDN
- `POST /dosen` → tambah data dosen baru
- `PUT /dosen/:nidn` → update data dosen
- `DELETE /dosen/:nidn` → hapus data dosen

Backend 2 — **Jabatan**
- `GET /jabatan` → ambil semua jabatan
- `POST /jabatan` → tambah jabatan baru
- `GET /jabatan/:id` → detail jabatan

Frontend — **Halaman Data Dosen**
- Tabel list semua dosen beserta jabatannya
- Form tambah dosen baru
- Filter berdasarkan jabatan

---

#### Modul 3 — Mata Kuliah & KRS
**Deskripsi:** Mengelola data mata kuliah dan pengambilan KRS oleh mahasiswa.

Backend 1 — **Mata Kuliah**
- `GET /matkul` → ambil semua mata kuliah
- `GET /matkul/:kode` → ambil detail mata kuliah berdasarkan kode
- `POST /matkul` → tambah mata kuliah baru
- `PUT /matkul/:kode` → update mata kuliah
- `DELETE /matkul/:kode` → hapus mata kuliah

Backend 2 — **KRS**
- `GET /krs/:npm` → ambil KRS mahasiswa berdasarkan NPM
- `POST /krs` → daftarkan mata kuliah ke KRS
- `DELETE /krs/:id` → batalkan KRS

Frontend — **Halaman Mata Kuliah**
- Tabel list mata kuliah (kode, nama, SKS, semester)
- Form tambah mata kuliah
- Form input KRS mahasiswa

---

#### Modul 4 — Jadwal & Ruangan
**Deskripsi:** Mengelola jadwal perkuliahan dan data ruangan yang tersedia.

Backend 1 — **Jadwal**
- `GET /jadwal` → ambil semua jadwal kuliah
- `GET /jadwal/:id` → detail jadwal
- `POST /jadwal` → tambah jadwal baru
- `PUT /jadwal/:id` → update jadwal
- `DELETE /jadwal/:id` → hapus jadwal

Backend 2 — **Ruangan**
- `GET /ruangan` → ambil semua ruangan
- `POST /ruangan` → tambah ruangan baru
- `GET /ruangan/:kode` → cek ketersediaan ruangan
- `PUT /ruangan/:kode` → update data ruangan

Frontend — **Halaman Jadwal Kuliah**
- Tampilan jadwal per hari/minggu
- Filter jadwal berdasarkan prodi atau dosen
- Info ruangan yang digunakan

---

#### Modul 5 — Nilai & Transkrip
**Deskripsi:** Mengelola input nilai mahasiswa dan rekap transkrip akademik.

Backend 1 — **Nilai**
- `GET /nilai/:npm` → ambil semua nilai mahasiswa berdasarkan NPM
- `POST /nilai` → input nilai mahasiswa
- `PUT /nilai/:id` → update nilai
- `DELETE /nilai/:id` → hapus nilai

Backend 2 — **Transkrip**
- `GET /transkrip/:npm` → ambil rekap seluruh nilai dan total SKS mahasiswa
- `GET /transkrip/:npm/ipk` → hitung dan return nilai IPK

Frontend — **Halaman Input Nilai**
- Form input nilai per mahasiswa per mata kuliah
- Tabel rekap nilai dengan IPK
- Filter berdasarkan semester

---

#### Modul 6 — Absensi & Rekap
**Deskripsi:** Mengelola data kehadiran mahasiswa dan rekap absensi per periode.

Backend 1 — **Absensi**
- `GET /absensi/:npm` → ambil absensi mahasiswa
- `POST /absensi` → input absensi mahasiswa
- `PUT /absensi/:id` → update status absensi
- `GET /absensi/hari-ini` → absensi yang masuk hari ini

Backend 2 — **Rekap Absensi**
- `GET /rekap-absensi/:npm` → rekap persentase kehadiran per matkul
- `GET /rekap-absensi/matkul/:kode` → rekap semua mahasiswa dalam satu matkul

Frontend — **Halaman Form Absensi**
- Form input absensi dengan tanggal dan status (hadir/izin/alpha)
- Tabel rekap persentase kehadiran per mahasiswa

---

#### Modul 7 — Pengumuman & Kategori
**Deskripsi:** Papan pengumuman digital untuk informasi kampus dengan sistem kategori.

Backend 1 — **Pengumuman**
- `GET /pengumuman` → ambil semua pengumuman (terbaru di atas)
- `GET /pengumuman/:id` → detail pengumuman
- `POST /pengumuman` → tambah pengumuman baru
- `PUT /pengumuman/:id` → update pengumuman
- `DELETE /pengumuman/:id` → hapus pengumuman

Backend 2 — **Kategori**
- `GET /kategori` → ambil semua kategori
- `POST /kategori` → tambah kategori baru
- `GET /pengumuman/kategori/:nama` → filter pengumuman berdasarkan kategori

Frontend — **Halaman Board Pengumuman**
- Tampilan card pengumuman terbaru
- Filter berdasarkan kategori
- Form tambah pengumuman

---

#### Modul 8 — Beasiswa & Pendaftaran
**Deskripsi:** Informasi beasiswa yang tersedia dan sistem pendaftaran beasiswa.

Backend 1 — **Beasiswa**
- `GET /beasiswa` → ambil semua beasiswa
- `GET /beasiswa/:id` → detail beasiswa
- `POST /beasiswa` → tambah data beasiswa
- `PUT /beasiswa/:id` → update beasiswa
- `DELETE /beasiswa/:id` → hapus beasiswa

Backend 2 — **Pendaftaran Beasiswa**
- `POST /beasiswa/daftar` → daftarkan mahasiswa ke beasiswa
- `GET /beasiswa/pendaftar/:id` → lihat daftar pendaftar beasiswa
- `GET /beasiswa/status/:npm` → cek status pendaftaran beasiswa mahasiswa

Frontend — **Halaman List Beasiswa**
- Tabel list beasiswa (nama, syarat, deadline)
- Form pendaftaran beasiswa
- Status pendaftaran per mahasiswa

---

#### Modul 9 — Perpustakaan & Peminjaman
**Deskripsi:** Katalog buku perpustakaan dan sistem peminjaman buku oleh mahasiswa.

Backend 1 — **Buku**
- `GET /buku` → ambil semua buku
- `GET /buku/:id` → detail buku
- `GET /buku/cari?judul=` → cari buku berdasarkan judul
- `POST /buku` → tambah buku baru
- `PUT /buku/:id` → update data buku

Backend 2 — **Peminjaman**
- `POST /peminjaman` → pinjam buku
- `PUT /peminjaman/:id/kembali` → kembalikan buku
- `GET /peminjaman/:npm` → riwayat peminjaman mahasiswa
- `GET /peminjaman/aktif` → daftar buku yang sedang dipinjam

Frontend — **Halaman Cari Buku**
- Search bar pencarian buku
- Tabel hasil pencarian dengan status ketersediaan
- Form peminjaman buku

---

#### Modul 10 — Prestasi & Kategori
**Deskripsi:** Pencatatan prestasi mahasiswa beserta kategori jenis prestasi.

Backend 1 — **Prestasi**
- `GET /prestasi` → ambil semua prestasi
- `GET /prestasi/:npm` → prestasi mahasiswa tertentu
- `POST /prestasi` → input prestasi baru
- `PUT /prestasi/:id` → update prestasi
- `DELETE /prestasi/:id` → hapus prestasi

Backend 2 — **Kategori Prestasi**
- `GET /kategori-prestasi` → ambil semua kategori (akademik, non-akademik, dll)
- `POST /kategori-prestasi` → tambah kategori baru
- `GET /prestasi/kategori/:nama` → filter prestasi berdasarkan kategori

Frontend — **Halaman Input Prestasi**
- Form input prestasi (nama event, tingkat, juara, tanggal)
- Tabel list prestasi mahasiswa dengan filter kategori

---

#### Modul 11 — Alumni & Lowongan Kerja
**Deskripsi:** Data alumni kampus dan informasi lowongan kerja yang relevan.

Backend 1 — **Alumni**
- `GET /alumni` → ambil semua data alumni
- `GET /alumni/:npm` → detail alumni
- `POST /alumni` → tambah data alumni baru
- `PUT /alumni/:npm` → update data alumni
- `GET /alumni/angkatan/:tahun` → filter alumni berdasarkan angkatan

Backend 2 — **Lowongan Kerja**
- `GET /lowongan` → ambil semua lowongan
- `GET /lowongan/:id` → detail lowongan
- `POST /lowongan` → tambah lowongan baru
- `PUT /lowongan/:id` → update lowongan
- `DELETE /lowongan/:id` → hapus lowongan

Frontend — **Halaman Data Alumni & Lowongan**
- Tabel data alumni dengan info pekerjaan
- List lowongan kerja terbaru
- Filter lowongan berdasarkan bidang

---

#### Modul 12 — Ormawa & Kegiatan
**Deskripsi:** Mengelola data organisasi mahasiswa (ormawa) dan kegiatan yang diselenggarakan.

Backend 1 — **Ormawa**
- `GET /ormawa` → ambil semua data ormawa
- `GET /ormawa/:id` → detail ormawa
- `POST /ormawa` → tambah ormawa baru
- `PUT /ormawa/:id` → update data ormawa
- `DELETE /ormawa/:id` → hapus ormawa

Backend 2 — **Kegiatan Ormawa**
- `GET /kegiatan` → ambil semua kegiatan
- `GET /kegiatan/:id` → detail kegiatan
- `POST /kegiatan` → tambah kegiatan baru
- `PUT /kegiatan/:id` → update kegiatan
- `GET /kegiatan/ormawa/:id` → ambil kegiatan berdasarkan ormawa

Frontend — **Halaman Data Ormawa**
- Tabel list ormawa beserta deskripsi dan pengurus
- List kegiatan yang diselenggarakan per ormawa
- Form tambah ormawa dan kegiatan baru

---

## 6. Pembagian Tugas (12 Mahasiswa)

| No | Modul BE 1 | Modul BE 2 | Menu FE | Nama
|----|-----------|-----------|---------|---------|
| 1 | Mahasiswa (CRUD) | Auth (login phone) | Halaman Data Mahasiswa |
| 2 | Dosen (CRUD) | Jabatan (CRUD) | Halaman Data Dosen |
| 3 | Mata Kuliah (CRUD) | KRS (daftar matkul) | Halaman Mata Kuliah |
| 4 | Jadwal (CRUD) | Ruangan (CRUD) | Halaman Jadwal Kuliah |
| 5 | Nilai (CRUD) | Transkrip (rekap + IPK) | Halaman Input Nilai |
| 6 | Absensi (CRUD) | Rekap Absensi | Halaman Form Absensi |
| 7 | Pengumuman (CRUD) | Kategori (CRUD) | Halaman Board Pengumuman |
| 8 | Beasiswa (CRUD) | Pendaftaran Beasiswa | Halaman List Beasiswa |
| 9 | Buku/Perpustakaan (CRUD) | Peminjaman Buku | Halaman Cari Buku |
| 10 | Prestasi (CRUD) | Kategori Prestasi | Halaman Input Prestasi |
| 11 | Alumni (CRUD) | Lowongan Kerja (CRUD) | Halaman Alumni & Lowongan |
| 12 | Ormawa (CRUD) | Kegiatan Ormawa (CRUD) | Halaman Data Ormawa |

---

## 7. Flow Pengerjaan

```mermaid
flowchart TD
    A([Mulai]) --> B[Clone boilerplate\ngocroot/alwaysdata]
    B --> C[Push ke repo NPM\ndi organisasi GitHub]
    C --> D[Daftar Alwaysdata\nDapat subdomain .alwaysdata.net]
    D --> E[Setup GitHub Secrets\nsshhost, sshpassword, apikey, dll]
    E --> F[Test CI/CD\ngit push ke main]
    F --> G{Auto-deploy\nberjalan?}
    G -- Tidak --> H[Cek GitHub Actions log\ndan perbaiki config]
    H --> F
    G -- Ya --> I[Buat model struct\nmodel/modul.go]
    I --> J[Buat controller\ncontroller/modul.go]
    J --> K[Daftarkan route\nurl/url.go]
    K --> L[Test endpoint\nPostman / curl]
    L --> M{Endpoint\nberjalan?}
    M -- Tidak --> N[Debug dan perbaiki kode]
    N --> L
    M -- Ya --> O[Buat Frontend\nfrontend/index.html]
    O --> P[Koneksikan FE ke BE\nvia fetch API]
    P --> Q{FE dan BE\njalan bersama?}
    Q -- Tidak --> R[Debug CORS\ndan endpoint]
    R --> P
    Q -- Ya --> S[Tulis dokumentasi API\ndi README.md]
    S --> T[Push final ke main\nPastikan live di Alwaysdata]
    T --> U[Kumpulkan link repo\ndan link live app]
    U --> V([Selesai])
```

---

## 8. Flow Aplikasi End-to-End

Menggambarkan perjalanan data dari user di browser hingga kembali sebagai respons JSON.

```mermaid
sequenceDiagram
    actor User as User (Browser)
    participant FE as Frontend (HTML/JS)
    participant Fiber as Go Fiber (Router)
    participant CTRL as Controller
    participant HLP as Helper
    participant DB as MongoDB Atlas

    User->>FE: Buka halaman / klik tombol
    FE->>Fiber: HTTP Request (GET/POST/PUT/DELETE)
    Note over FE,Fiber: Header: Content-Type, Token (jika perlu)

    Fiber->>Fiber: Validasi CORS
    Fiber->>CTRL: Route cocok, panggil handler

    alt Request GET (ambil data)
        CTRL->>HLP: Panggil fungsi GetAllDoc / GetOneDoc
        HLP->>DB: db.Collection.Find()
        DB-->>HLP: Cursor / Document
        HLP-->>CTRL: Slice of struct
        CTRL-->>Fiber: c.JSON(data)
    else Request POST (tambah data)
        CTRL->>CTRL: c.BodyParser(&struct)
        CTRL->>HLP: Panggil fungsi InsertOneDoc
        HLP->>DB: db.Collection.InsertOne()
        DB-->>HLP: InsertedID
        HLP-->>CTRL: Result
        CTRL-->>Fiber: c.JSON(response)
    else Request PUT (update data)
        CTRL->>CTRL: c.BodyParser(&struct)
        CTRL->>HLP: Panggil fungsi UpdateDoc
        HLP->>DB: db.Collection.UpdateOne()
        DB-->>HLP: UpdateResult
        HLP-->>CTRL: Result
        CTRL-->>Fiber: c.JSON(response)
    else Request DELETE (hapus data)
        CTRL->>HLP: Panggil fungsi DeleteDoc
        HLP->>DB: db.Collection.DeleteOne()
        DB-->>HLP: DeleteResult
        HLP-->>CTRL: Result
        CTRL-->>Fiber: c.JSON(response)
    end

    Fiber-->>FE: HTTP Response (JSON)
    FE-->>User: Tampilkan data / notifikasi
```

---

## 9. Konsep CI/CD

### Penjelasan

CI/CD (Continuous Integration / Continuous Deployment) adalah praktik otomatisasi proses build dan deploy aplikasi. Setiap kali mahasiswa melakukan `git push` ke branch `main`, GitHub Actions secara otomatis akan:

1. **Build** — mengkompilasi kode Go menjadi binary executable
2. **Transfer** — mengirim binary ke server Alwaysdata via SCP (SSH Copy)
3. **Restart** — memanggil API Alwaysdata untuk merestart aplikasi agar perubahan aktif

### Diagram CI/CD Pipeline

```mermaid
flowchart LR
    subgraph Local["Lokal (Mahasiswa)"]
        CODE["Tulis Kode\n(Go + HTML)"]
        PUSH["git push\norigin main"]
    end

    subgraph GitHub["GitHub"]
        REPO["Repository\n(nama = NPM)"]
        GA["GitHub Actions\nRunner (ubuntu-latest)"]

        subgraph Pipeline["CI/CD Pipeline"]
            S1["Step 1\nCheckout code"]
            S2["Step 2\nSetup Go 1.22.x"]
            S3["Step 3\nSet CGO_ENABLED=0"]
            S4["Step 4\ngo get . (install deps)"]
            S5["Step 5\ngo build -o gocroot\nchmod a+x gocroot"]
            S6["Step 6\nSCP: kirim binary\nke Alwaysdata"]
            S7["Step 7\nSSH: ls + curl IP\n+ restart via API"]
        end
    end

    subgraph Alwaysdata["Alwaysdata Server"]
        BIN["Binary: gocroot\n(file executable)"]
        APP["Go Fiber App\n(berjalan di port)"]
        URL2["yourname.alwaysdata.net"]
    end

    CODE --> PUSH
    PUSH --> REPO
    REPO -->|trigger on push to main| GA
    GA --> S1 --> S2 --> S3 --> S4 --> S5 --> S6 --> S7
    S6 -->|SCP via SSH| BIN
    S7 -->|POST API restart| APP
    BIN --> APP
    APP --> URL2
```

### GitHub Secrets yang Dipakai di Pipeline

| Secret | Dipakai di Step | Fungsi |
|--------|----------------|--------|
| `sshhost` | Step 6 & 7 | Alamat server SSH Alwaysdata |
| `sshusername` | Step 6 & 7 | Username akun Alwaysdata |
| `sshpassword` | Step 6 & 7 | Password SSH |
| `sshport` | Step 6 & 7 | Port SSH (biasanya 22) |
| `folder` | Step 6 | Path folder tujuan binary di server |
| `apikey` | Step 7 | Token API untuk restart app |
| `appid` | Step 7 | ID site di dashboard Alwaysdata |

---

## 10. Struktur Folder Per Mahasiswa

Mengikuti struktur boilerplate `gocroot/alwaysdata`:

```
[NPM]/
├── .github/
│   └── workflows/
│       └── alwaysdata.yml       <- CI/CD config
├── config/
│   ├── api.go
│   ├── config.go
│   ├── cors.go
│   ├── db.go
│   └── token.go
├── controller/
│   ├── controller.go            <- homepage, ip
│   ├── [modul1].go              <- modul BE 1 (buat sendiri)
│   └── [modul2].go              <- modul BE 2 (buat sendiri)
├── model/
│   ├── model.go
│   └── [modul].go               <- struct model (buat sendiri)
├── helper/
│   ├── helper.go
│   └── mongodb.go
├── url/
│   └── url.go                   <- daftarkan semua route di sini
├── frontend/
│   └── index.html               <- 1 menu frontend
├── main.go
├── go.mod
└── .env
```

---

## 11. Standar REST API

Setiap modul wajib mengikuti standar berikut:

### Request & Response

```
GET    /[resource]          -> ambil semua data
GET    /[resource]/:id      -> ambil satu data
POST   /[resource]          -> tambah data baru
PUT    /[resource]/:id      -> update data
DELETE /[resource]/:id      -> hapus data
```

### Format Response Sukses
```json
{
  "status": "success",
  "data": { ... }
}
```

### Format Response Error
```json
{
  "status": "error",
  "message": "deskripsi error"
}
```

---

## 12. GitHub Secrets yang Dibutuhkan

| Secret | Keterangan |
|--------|-----------|
| `sshhost` | SSH host Alwaysdata |
| `sshusername` | Username SSH |
| `sshpassword` | Password SSH (strong password) |
| `sshport` | Port SSH (default 22) |
| `apikey` | API key dari Alwaysdata dashboard |
| `appid` | App ID dari menu Web > Sites |
| `folder` | Path folder home di Alwaysdata |

---

## 13. Environment Variables

```sh
MONGOSTRING=mongodb+srv://user:pass@cluster.mongodb.net/
PORT=porthost
IP=iphost
```

---

## 14. Kriteria Selesai (Definition of Done)

- [ ] App berhasil di-deploy di Alwaysdata dan bisa diakses via URL
- [ ] CI/CD berjalan: setiap push ke `main` otomatis deploy
- [ ] Minimal 2 endpoint BE per mahasiswa berjalan (GET + POST)
- [ ] 1 menu FE bisa menampilkan dan/atau mengirim data ke BE
- [ ] Dokumentasi API tersedia (minimal dalam README.md)
- [ ] Link live app + repo dikumpulkan di repo NPM dalam organisasi GitHub

---

## 15. Referensi

- Boilerplate: https://github.com/gocroot/alwaysdata
- Go Fiber Docs: https://gofiber.io/
- MongoDB Atlas: https://www.mongodb.com/atlas
- Alwaysdata: https://www.alwaysdata.com
- GitHub Actions Docs: https://docs.github.com/actions
