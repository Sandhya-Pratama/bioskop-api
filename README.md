
# Bioskop API

Bioskop API adalah proyek backend berbasis Golang yang menyediakan layanan untuk sistem pemesanan tiket bioskop. API ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada jadwal tayang film di bioskop, serta menangani proses autentikasi dan manajemen pengguna.

## Daftar Isi

- [Fitur](#fitur)
- [Teknologi yang Digunakan](#teknologi-yang-digunakan)
- [Persyaratan](#persyaratan)
- [Instalasi](#instalasi)
- [Penggunaan API](#penggunaan-api)
- [Mengimpor Koleksi Postman](#mengimpor-koleksi-postman)
- [Lisensi](#lisensi)

## Fitur

- **CRUD Jadwal Tayang**: Menambah, melihat, memperbarui, dan menghapus jadwal tayang film di berbagai cabang bioskop.
- **Autentikasi JWT**: Mengamankan endpoint dengan otorisasi token berbasis JWT.
- **Manajemen Pengguna**: API untuk mengelola pengguna, dengan peran khusus untuk administrator.

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman utama.
- **Echo**: Framework web untuk membangun API.
- **GORM**: ORM untuk mengelola database.
- **PostgreSQL**: Database yang digunakan untuk menyimpan data.
- **Goose**: Untuk migrasi database.
- **Postman**: Untuk menguji API.

## Persyaratan

- Golang versi 1.16 atau lebih baru
- PostgreSQL versi 12 atau lebih baru
- Postman (untuk pengujian API)
- Goose (untuk migrasi database)

## Instalasi

1. **Clone repository ini**:
   ```bash
   git clone https://github.com/username/bioskop-api.git
   cd bioskop-api
   ```

2. **Konfigurasi Environment**:
   Buat file `.env` di direktori utama dengan konfigurasi berikut:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=bioskop_db
   JWT_SECRET=your_jwt_secret_key
   ```

3. **Jalankan Migrasi Database**:
   Pastikan Goose sudah terinstall:
   ```bash
   go get -u github.com/pressly/goose/v3/cmd/goose
   ```
   Lalu, jalankan migrasi:
   ```bash
   goose up
   ```

4. **Jalankan Server**:
   ```bash
   go run main.go
   ```
   Server akan berjalan di `http://localhost:8080`.


## Penggunaan API

### Autentikasi

Gunakan endpoint berikut untuk autentikasi pengguna:

- **Login**: `POST /v1/login`
- **Register**: `POST /v1/register`

### Endpoint Jadwal Tayang

| Method | Endpoint              | Deskripsi                           |
|--------|------------------------|-------------------------------------|
| GET    | `/api/v1/jadwals`     | Mendapatkan semua jadwal tayang     |
| GET    | `/api/v1/jadwals/:id` | Mendapatkan jadwal tayang berdasarkan ID |
| POST   | `/api/v1/jadwals`     | Menambah jadwal tayang baru         |
| PUT    | `/api/v1/jadwals/:id` | Memperbarui jadwal tayang           |
| DELETE | `/api/v1/jadwals/:id` | Menghapus jadwal tayang             |

**Contoh JSON untuk Menambah Jadwal Tayang**:
```json
{
  "film_id": 1,
  "bioskop_id": 2,
  "tanggal_tayang": "2024-11-10",
  "waktu_tayang": "15:30:00"
}
```

## Mengimpor Koleksi Postman

1. Buka Postman dan pilih **Import**.
2. Pilih file JSON koleksi Postman yang sudah diekspor dari proyek ini.
3. Setelah impor selesai, Anda akan melihat semua endpoint API di dalam Postman, siap untuk diuji.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
