# Manajemen Data Hewan (Go)

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg) ![License](https://img.shields.io/badge/License-MIT-green.svg) Aplikasi konsol sederhana yang ditulis dalam bahasa Go untuk melakukan manajemen data hewan. Aplikasi ini memungkinkan pengguna untuk menambah, menampilkan, mengedit, menghapus, dan melakukan berbagai operasi lain terkait data hewan.

## Daftar Isi

* [Fitur](#fitur)
* [Prasyarat](#prasyarat)
* [Instalasi & Menjalankan](#instalasi--menjalankan)
    * [Mendapatkan Kode](#mendapatkan-kode)
    * [Menjalankan Aplikasi](#menjalankan-aplikasi)
* [Penggunaan](#penggunaan)
* [Struktur Proyek](#struktur-proyek-opsional)
* [Kontribusi](#kontribusi-opsional)
* [Lisensi](#lisensi)

## Fitur

Aplikasi ini menyediakan berbagai fungsionalitas, antara lain:

* **Tambah Data Hewan:** Memasukkan data hewan baru (ID, Jenis, Nama, Umur, Pemilik).
* **Tampilkan Semua Data:** Menampilkan seluruh data hewan yang tersimpan.
* **Hitung Rata-rata Umur:** Menghitung dan menampilkan rata-rata umur semua hewan.
* **Cari Hewan Tertua dan Termuda:** Menemukan dan menampilkan hewan dengan umur tertua dan termuda.
* **Urutkan Berdasarkan Umur (Bubble Sort):** Mengurutkan data hewan berdasarkan umur secara menaik.
* **Cari Hewan Berdasarkan Nama:** Melakukan pencarian data hewan berdasarkan nama.
* **Edit Data Hewan:** Mengubah detail data hewan yang sudah ada berdasarkan ID.
* **Hapus Data Hewan:** Menghapus data hewan berdasarkan ID.
* **Statistik Per Jenis:** Menampilkan jumlah hewan untuk setiap jenisnya.
* **Sorting Berdasarkan Nama Pemilik (A-Z - Insertion Sort):** Mengurutkan data hewan berdasarkan nama pemilik secara alfabetis.
* **Urutkan Berdasarkan Umur (Insertion Sort):** Mengurutkan data hewan berdasarkan umur secara menaik menggunakan Insertion Sort.

## Prasyarat

Untuk menjalankan aplikasi ini, Anda memerlukan:

* **Go (Golang):** Versi 1.18 atau yang lebih baru. Anda bisa mengunduhnya dari [golang.org](https://golang.org/dl/).
* **Git (Opsional):** Untuk melakukan clone repositori ini.

## Instalasi & Menjalankan

### Mendapatkan Kode

1.  **Clone repositori (jika sudah diunggah ke GitHub):**
    ```bash
    git clone [https://github.com/Danzz2706/TUGAS BESAR ALPRO.git](https://github.com/NAMA_PENGGUNA_ANDA/NAMA_REPOSITORI_ANDA.git)
    cd TUGAS BESAR ALPRO
    ```
2.  **Atau, jika Anda memiliki file `main.go` secara lokal:**
    Pastikan Anda berada di direktori yang berisi file `main.go`.

### Menjalankan Aplikasi

1.  **Buka terminal atau command prompt** di direktori proyek Anda (yang berisi file `main.go`).
2.  **Jalankan program Go:**
    ```bash
    go run main.go
    ```
    Ini akan mengkompilasi dan langsung menjalankan program.

3.  **Atau, Anda bisa mengkompilasi menjadi file executable terlebih dahulu:**
    ```bash
    go build main.go
    ```
    Ini akan menghasilkan file executable (misalnya `main` atau `main.exe` tergantung sistem operasi Anda). Kemudian jalankan executable tersebut:
    * Untuk Linux/macOS:
        ```bash
        ./main
        ```
    * Untuk Windows:
        ```bash
        main.exe
        ```

## Penggunaan

Setelah aplikasi berjalan, Anda akan melihat menu utama dengan berbagai pilihan. Masukkan nomor pilihan yang sesuai untuk melakukan operasi yang diinginkan. Ikuti instruksi yang muncul di layar untuk setiap operasi.

Contoh:
