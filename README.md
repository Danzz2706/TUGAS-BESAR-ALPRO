# Pendataan Hewan

Sebuah aplikasi baris perintah (command-line) yang ditulis dalam bahasa Go untuk mengelola data hewan. Aplikasi ini memungkinkan pengguna untuk menambah, menampilkan, mencari, mengurutkan, mengedit, menghapus, dan melihat statistik data hewan. Program diinisialisasi dengan 20 data hewan sampel.

## Pembuat

Proyek ini dibuat sebagai kolaborasi dengan kontribusi yang setara oleh:

* Zaidan Kamil Munadi
* Roby Ariga Siagian

## Fitur

Program ini menyediakan antarmuka berbasis menu dengan fungsionalitas sebagai berikut:

1.  **Tambah Data Hewan:** Memungkinkan penambahan data hewan baru meliputi ID, Jenis, Nama, Umur, dan Nama Pemilik.
2.  **Tampilkan Semua Data:** Menampilkan semua data hewan yang tersimpan dalam format tabel.
3.  **Hitung Rata-rata Umur:** Menghitung dan menampilkan rata-rata umur dari semua hewan.
4.  **Cari Hewan Tertua dan Termuda:** Mengidentifikasi dan menampilkan data hewan tertua dan termuda.
5.  **Sorting Berdasarkan Nama Pemilik (A-Z):** Mengurutkan dan menampilkan data hewan berdasarkan nama pemilik secara alfabetis (A-Z) menggunakan algoritma Insertion Sort.
6.  **Cari Hewan Berdasarkan Nama:** Mencari hewan berdasarkan namanya (sensitif terhadap huruf besar/kecil) dan menampilkan data yang cocok.
7.  **Cari Hewan Berdasarkan ID (Binary Search):** Mencari hewan berdasarkan ID-nya. Data akan diurutkan terlebih dahulu berdasarkan ID (menggunakan Insertion Sort) untuk memungkinkan pencarian dengan Binary Search.
8.  **Edit Data Hewan:** Memungkinkan modifikasi detail hewan yang sudah ada (Jenis, Nama, Umur, Pemilik) berdasarkan ID-nya.
9.  **Hapus Data Hewan:** Menghapus data hewan dari sistem berdasarkan ID-nya.
10. **Urutkan Berdasarkan Umur Dari Tertua:** Mengurutkan dan menampilkan data hewan berdasarkan umur, dari yang tertua hingga termuda, menggunakan algoritma Selection Sort.
11. **Urutkan Berdasarkan Umur Dari Termuda (Insertion Sort):** Mengurutkan dan menampilkan data hewan berdasarkan umur, dari yang termuda hingga tertua, menggunakan algoritma Insertion Sort.
12. **Statistik Berdasarkan Jenis:** Menampilkan ringkasan jumlah hewan untuk setiap jenis yang ada dalam data.
0.  **Keluar:** Untuk keluar dari program.

## Prasyarat

* Go compiler terinstal di sistem Anda. Anda dapat mengunduhnya dari [golang.org](https://golang.org/dl/).

## Cara Menjalankan

1.  **Clone repositori atau unduh berkas `pendataanHewan.go`.**
    ```bash
    # Jika Anda memiliki git terinstal
    # git clone <url_repositori>
    # cd <direktori_repositori>
    ```
    Jika Anda hanya memiliki berkas `.go`, simpan di sebuah direktori pada komputer Anda.

2.  **Buka terminal atau command prompt Anda.**

3.  **Navigasi ke direktori tempat berkas `pendataanHewan.go` disimpan.**
    ```bash
    cd path/to/your/directory
    ```

4.  **Kompilasi program Go:**
    ```bash
    go build pendataanHewan.go
    ```
    Ini akan membuat berkas eksekusi bernama `pendataanHewan` (atau `pendataanHewan.exe` di Windows).

5.  **Jalankan berkas eksekusi:**
    * Di macOS/Linux:
        ```bash
        ./pendataanHewan
        ```
    * Di Windows:
        ```bash
        .\pendataanHewan.exe
        ```

    Sebagai alternatif, Anda dapat menjalankan program secara langsung tanpa mengkompilasi berkas eksekusi terpisah menggunakan:
    ```bash
    go run pendataanHewan.go
    ```

## Penggunaan

Setelah program berjalan, Anda akan disajikan sebuah menu. Masukkan nomor yang sesuai dengan tindakan yang ingin Anda lakukan dan tekan Enter. Ikuti petunjuk di layar untuk setiap opsi.

## Struktur Kode

* **`pendataanHewan.go`**: Berkas utama yang berisi semua kode sumber.
* **`Hewan` struct**: Mendefinisikan struktur untuk menyimpan data hewan:
    * `ID` (string): Pengidentifikasi unik untuk hewan.
    * `Jenis` (string): Jenis atau spesies hewan.
    * `Nama` (string): Nama hewan.
    * `Umur` (int): Usia hewan dalam tahun.
    * `Pemilik` (string): Nama pemilik hewan.
* **`dataHewan [maxHewan]Hewan`**: Sebuah array yang menyimpan data hewan. `maxHewan` adalah konstanta yang diatur ke 100.
* **`jumlahData int`**: Variabel yang mencatat jumlah data hewan saat ini yang tersimpan dalam array `dataHewan`.
* **`main()` function**: Menginisialisasi beberapa data sampel dan memanggil fungsi `menu()`.
* **`menu()` function**: Menampilkan menu utama dan menangani input pengguna untuk menavigasi melalui berbagai fungsionalitas.
* **Berbagai fungsi** (`tambahData`, `tampilkanData`, `rataRataUmur`, `binarySearchByID`, `selectionSortUmur`, `insertionSortUmur`, dll.): Mengimplementasikan fitur-fitur spesifik yang tercantum dalam menu.

## Validasi Input

Program ini menyertakan validasi input dasar:
* Kolom ID, Jenis, Nama, dan Pemilik tidak boleh kosong saat menambahkan data baru.
* Umur tidak boleh negatif.
* Program memeriksa apakah kapasitas data maksimum (`maxHewan`) telah tercapai sebelum menambahkan data baru.
