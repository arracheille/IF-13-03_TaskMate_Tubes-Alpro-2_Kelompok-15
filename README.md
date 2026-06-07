# 🏠 TaskMate - Aplikasi Manajemen Tugas Rumah Tangga

TaskMate adalah aplikasi berbasis terminal yang ditulis dalam bahasa **Go** untuk membantu mengelola tugas-tugas rumah tangga. Program ini mendukung operasi CRUD lengkap, pencarian, pengurutan, dan statistik tugas.

## ✨ Fitur

| No | Fitur | Keterangan |
|----|-------|------------|
| 1 | Tambah Tugas | Menambahkan tugas baru dengan detail lengkap |
| 2 | Tampilkan Tugas | Menampilkan seluruh daftar tugas |
| 3 | Ubah Tugas | Mengubah data tugas berdasarkan ID |
| 4 | Hapus Tugas | Menghapus tugas berdasarkan ID |
| 5 | Tandai Selesai | Menandai tugas sebagai sudah selesai |
| 6 | Sequential Search | Mencari tugas berdasarkan nama atau kategori |
| 7 | Binary Search | Mencari tugas nama (data diurutkan terlebih dahulu) |
| 8 | Insertion Sort | Mengurutkan tugas berdasarkan durasi (ascending) |
| 9 | Selection Sort | Mengurutkan tugas berdasarkan tingkat kesulitan (descending) |
| 10 | Statistik | Menampilkan ringkasan dan rata-rata waktu pengerjaan |

## 🗂 Struktur Data

Setiap tugas disimpan dalam struct `Tugas` dengan field berikut:

```go
type Tugas struct {
    ID        string  // Identitas unik tugas
    NamaTugas string  // Nama pekerjaan
    Kategori  string  // Kategori ruangan (misal: Dapur, Kamar)
    Deskripsi string  // Deskripsi pekerjaan
    Kesulitan int     // Skala kesulitan: 1=Mudah, 2=Sedang, 3=Sulit
    Durasi    int     // Estimasi durasi pengerjaan (dalam menit)
    Selesai   bool    // Status penyelesaian tugas
}
```

## 🗒 Menu Program

Setelah program berjalan, akan muncul menu interaktif berikut:

```
===== MENU TASKMATE =====
1.  Tambah Tugas
2.  Tampilkan Tugas
3.  Ubah Tugas
4.  Hapus Tugas
5.  Tandai Tugas Selesai
6.  Sequential Search
7.  Binary Search
8.  Insertion Sort (Durasi)
9.  Selection Sort (Kesulitan)
10. Statistik
0.  Keluar
```

Masukkan angka sesuai menu yang diinginkan, lalu tekan **Enter**.

## 💡 Contoh Penggunaan

**Menambah tugas baru:**
```
Masukkan ID Tugas            : T001
Masukkan Nama Tugas          : Menyapu
Masukkan Kategori Ruangan    : Ruang Tamu
Masukkan Deskripsi Pekerjaan : Menyapu lantai ruang tamu
Pilih skala kesulitan (1-3)  : 1
Estimasi Durasi Pengerjaan   : 15
Data tugas berhasil ditambahkan!
```

**Melihat statistik:**
```
===== STATISTIK TASKMATE =====
Total Tugas           : 5
Tugas Selesai         : 3
Tugas Belum Selesai   : 2
Rata-rata Waktu Kerja : 20.00 menit
```

---

## 📝 Catatan

- Input nama, kategori, dan deskripsi yang mengandung **spasi** tidak didukung karena menggunakan `fmt.Scan()`. Gunakan kata tunggal atau garis bawah (contoh: `Ruang_Tamu`).