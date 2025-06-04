# 🐾 Manajemen Data Hewan 🐾

Program konsol sederhana yang ditulis dalam bahasa Go untuk mengelola data hewan peliharaan. Aplikasi ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data hewan, serta menyediakan fungsionalitas pencarian dan pengurutan data menggunakan berbagai algoritma dasar.

Program ini dibuat sebagai contoh implementasi struktur data dan algoritma dasar dalam Go, dengan fokus pada penggunaan pustaka standar `fmt` untuk interaksi pengguna.

---

## 📋 Fitur Utama

* **Manajemen Data**: Tambah, tampilkan, edit, dan hapus data hewan.
* **Kalkulasi**: Menghitung rata-rata umur hewan.
* **Analisis Data**: Menemukan hewan tertua dan termuda.
* **Pencarian**:
    * Berdasarkan nama panggilan (menggunakan Sequential Search).
    * Berdasarkan ID hewan (menggunakan Binary Search, setelah data diurutkan).
* **Pengurutan**:
    * Berdasarkan nama pemilik (A-Z, menggunakan Insertion Sort).
    * Berdasarkan umur dari yang tertua (menggunakan Selection Sort).
    * Berdasarkan umur dari yang termuda (menggunakan Insertion Sort).
* **Statistik**: Menampilkan statistik jumlah hewan berdasarkan jenisnya.

---

## 🛠️ Teknologi yang Digunakan

* **Go (Golang)**: Bahasa pemrograman utama.
* **Pustaka Standar Go**: Terutama pustaka `fmt` untuk semua operasi input/output.

---

## Pembuat
###Zaidan Kamil Munadi
###Roby Ariga Siagian

---

## 🧠 Algoritma yang Digunakan

Berikut adalah penjelasan dan implementasi kode untuk algoritma pencarian dan pengurutan utama yang digunakan dalam proyek ini. Kode diambil dari file `main.go`.

### 1. Sequential Search (Pencarian Berurutan)

Sequential Search adalah algoritma pencarian sederhana yang bekerja dengan cara memeriksa setiap elemen dalam kumpulan data secara berurutan, satu per satu, hingga elemen yang dicari ditemukan atau seluruh elemen telah diperiksa. Algoritma ini cocok untuk data yang tidak terurut atau kumpulan data kecil.

**Implementasi Kode (Pencarian Hewan berdasarkan Nama Panggilan):**

Fungsi `searchHewan` melakukan iterasi pada array `dataHewan` untuk menemukan kecocokan berdasarkan nama. Perbandingan nama bersifat case-sensitive.

```go
// Fungsi searchHewan (pencarian berdasarkan nama)
func searchHewan(dataHewan *ArrHewan, jumlahData *int) {
	var nama string
	fmt.Print("Masukkan nama panggilan hewan yang dicari: (contoh: Oyen): ")
	fmt.Scanln(&nama)
	fmt.Println()

	ditemukan := false
	// Bagian tampilan header tabel
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                       🔎 HASIL PENCARIAN 🔎                                       ║")
	fmt.Println("╠════╦═════════════╦═══════════════════╦══════════════════════╦═══════════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan    │ Jenis Hewan       │ Nama Hewan           │ Umur      │ Nama Pemilik          ║")
	fmt.Println("╠════╬═════════════╬═══════════════════╬══════════════════════╬═══════════╬═══════════════════════╣")

	count := 0
	// Loop untuk Sequential Search
	for i := 0; i < *jumlahData; i++ {
		if (*dataHewan)[i].Nama == nama { // Perbandingan langsung (case-sensitive)
			h := (*dataHewan)[i]
			count++
			fmt.Printf("║ %-2d │ %-11s │ %-17s │ %-20s │ %-9d │ %-21s ║\n",
				count, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("║                                    Data tidak ditemukan.                                        ║")
	}
	fmt.Println("╚════╩═════════════╩═══════════════════╩══════════════════════╩═══════════╩═══════════════════════╝")
}
```

##2. Binary Search (Pencarian Biner)
Binary Search adalah algoritma pencarian yang efisien dan hanya dapat digunakan pada data yang sudah terurut. Algoritma ini bekerja dengan membandingkan elemen target dengan elemen tengah array. Jika tidak cocok, separuh array di mana target tidak mungkin berada akan dieliminasi, dan pencarian dilanjutkan pada separuh sisanya hingga target ditemukan atau interval pencarian kosong.

Implementasi Kode (Pencarian Hewan berdasarkan ID):

Pertama, data diurutkan berdasarkan ID menggunakan fungsi sortByIDLogic (yang merupakan implementasi dari Insertion Sort). Kemudian, binarySearchByID melakukan pencarian.
```go
// Fungsi binarySearchByID
func binarySearchByID(dataHewan *ArrHewan, jumlahData *int) {
    var id string
    fmt.Print("Masukkan ID hewan yang dicari: (contoh K001): ")
    fmt.Scanln(&id)
    fmt.Println()

    if *jumlahData == 0 {
        fmt.Println("Tidak ada data untuk dicari.")
        return
    }
    sortByIDLogic(dataHewan, jumlahData) // Data harus terurut berdasarkan ID

    low := 0
    high := *jumlahData - 1
    found := false
    var hewanDitemukan Hewan

    // Loop untuk Binary Search
    for low <= high {
        mid := (low + high) / 2
        if (*dataHewan)[mid].ID == id {
            hewanDitemukan = (*dataHewan)[mid]
            found = true
            break
        } else if (*dataHewan)[mid].ID < id { // Perbandingan string (case-sensitive)
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    if !found {
        fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
        fmt.Println("║                             🔍 HASIL PENCARIAN BINARY SEARCH 🔍                               ║")
        fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════╣")
        fmt.Printf("║ Data dengan ID '%-76s' tidak ditemukan.                                ║\n", id)
        fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════╝")
        return
    }

    // Tampilkan data yang ditemukan
    fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
    fmt.Println("║                             🔍 HASIL PENCARIAN BINARY SEARCH 🔍                               ║")
    fmt.Println("╠═════════════╦═══════════════════╦══════════════════════╦═══════════╦═══════════════════════╣")
    fmt.Println("║ ID Hewan    │ Jenis Hewan       │ Nama Hewan           │ Umur      │ Nama Pemilik          ║")
    fmt.Println("╠═════════════╬═══════════════════╬══════════════════════╬═══════════╬═══════════════════════╣")
    fmt.Printf("║ %-11s │ %-17s │ %-20s │ %-9d │ %-21s ║\n",
        hewanDitemukan.ID, hewanDitemukan.Jenis, hewanDitemukan.Nama, hewanDitemukan.Umur, hewanDitemukan.Pemilik)
    fmt.Println("╚═════════════╩═══════════════════╩══════════════════════╩═══════════╩═══════════════════════╝")
}

// Fungsi sortByIDLogic (menggunakan Insertion Sort untuk mengurutkan berdasarkan ID)
func sortByIDLogic(dataHewan *ArrHewan, jumlahData *int) {
    for i := 1; i < *jumlahData; i++ {
        temp := (*dataHewan)[i]
        j := i - 1
        // Algoritma Insertion Sort untuk ID
        for j >= 0 && (*dataHewan)[j].ID > temp.ID { // Perbandingan string (case-sensitive)
            (*dataHewan)[j+1] = (*dataHewan)[j]
            j--
        }
        (*dataHewan)[j+1] = temp
    }
}
```
##Selection Sort (Pengurutan Seleksi)
Selection Sort adalah algoritma pengurutan berbasis perbandingan. Algoritma ini bekerja dengan berulang kali menemukan elemen minimum (atau maksimum, tergantung pada urutan yang diinginkan) dari bagian array yang belum terurut dan menukarnya dengan elemen pertama dari bagian yang belum terurut. Proses ini diulang untuk sisa array.

Implementasi Kode (Pengurutan Umur dari Tertua ke Termuda):

Fungsi selectionSortUmurLogic mengurutkan data hewan berdasarkan umur secara menurun (dari yang tertua). Fungsi selectionSortUmurAndDisplay adalah wrapper yang memanggil logika sorting kemudian menampilkan hasilnya.
```go
// Fungsi selectionSortUmurLogic (pengurutan umur dari tertua ke termuda)
func selectionSortUmurLogic(dataHewan *ArrHewan, jumlahData *int) { //descending tua -> muda
	for pass := 0; pass < *jumlahData-1; pass++ {
		idx := pass // Asumsikan elemen saat ini adalah yang terbesar/tertua
		for i := pass + 1; i < *jumlahData; i++ {
			// Cari elemen yang lebih besar/tua dari elemen di 'idx'
			if (*dataHewan)[i].Umur > (*dataHewan)[idx].Umur {
				idx = i // Update indeks elemen terbesar/tertua
			}
		}
		// Tukar elemen di 'pass' dengan elemen terbesar/tertua yang ditemukan
		(*dataHewan)[pass], (*dataHewan)[idx] = (*dataHewan)[idx], (*dataHewan)[pass]
	}
}

// Fungsi wrapper untuk sorting dan menampilkan
func selectionSortUmurAndDisplay(dataHewan *ArrHewan, jumlahData *int) {
	selectionSortUmurLogic(dataHewan, jumlahData)
	fmt.Println("\nData Diurutkan Berdasarkan Umur (Tertua -> Termuda):")
	tampilkanData(dataHewan, jumlahData) // Memanggil fungsi display umum
}
```
##Insertion Sort (Pengurutan Sisip)
Insertion Sort adalah algoritma pengurutan sederhana yang membangun array terurut akhir satu elemen pada satu waktu. Algoritma ini mengambil setiap elemen dari input dan memasukkannya ke posisi yang benar dalam bagian array yang sudah terurut. Efisien untuk data kecil atau data yang hampir terurut.

Implementasi Kode:

Pengurutan berdasarkan Nama Pemilik (A-Z):
```go
// Fungsi sortByPemilikLogic (menggunakan Insertion Sort untuk nama pemilik)
func sortByPemilikLogic(dataHewan *ArrHewan, jumlahData *int) {
    for i := 1; i < *jumlahData; i++ {
        temp := (*dataHewan)[i] // Ambil elemen untuk disisipkan
        j := i - 1
        // Geser elemen yang lebih besar dari 'temp' ke kanan
        for j >= 0 && (*dataHewan)[j].Pemilik > temp.Pemilik { // Perbandingan string (case-sensitive)
            (*dataHewan)[j+1] = (*dataHewan)[j]
            j--
        }
        // Sisipkan 'temp' ke posisi yang benar
        (*dataHewan)[j+1] = temp
    }
}

// Fungsi wrapper untuk sorting dan menampilkan
func sortByPemilikAndDisplay(dataHewan *ArrHewan, jumlahData *int) {
    sortByPemilikLogic(dataHewan, jumlahData)
    fmt.Println("\nData Diurutkan Berdasarkan Nama Pemilik (A-Z):")
    tampilkanData(dataHewan, jumlahData)
}
```
Pengurutan berdasarkan Umur (Termuda ke Tertua):
```
// Fungsi insertionSortUmurLogic (pengurutan umur dari termuda ke tertua)
func insertionSortUmurLogic(dataHewan *ArrHewan, jumlahData *int) { //ascending muda -> tua
    for i := 1; i < *jumlahData; i++ {
        temp := (*dataHewan)[i] // Ambil elemen untuk disisipkan
        j := i - 1
        // Geser elemen yang lebih besar dari 'temp.Umur' ke kanan
        for j >= 0 && (*dataHewan)[j].Umur > temp.Umur {
            (*dataHewan)[j+1] = (*dataHewan)[j]
            j--
        }
        // Sisipkan 'temp' ke posisi yang benar
        (*dataHewan)[j+1] = temp
    }
}

// Fungsi wrapper untuk sorting dan menampilkan
func insertionSortUmurAndDisplay(dataHewan *ArrHewan, jumlahData *int) {
    insertionSortUmurLogic(dataHewan, jumlahData)
    fmt.Println("\nData Diurutkan Berdasarkan Umur (Termuda -> Tertua):")
    tampilkanData(dataHewan, jumlahData)
}
```
