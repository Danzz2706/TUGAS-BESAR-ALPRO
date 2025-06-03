package main

import "fmt"

// Definisikan struct Hewan
type Hewan struct {
	ID      string
	Jenis   string
	Nama    string
	Umur    int
	Pemilik string
}

// Definisikan konstanta untuk ukuran maksimum array
const maxHewan = 20

// Definisikan tipe array kustom untuk data hewan
type ArrHewan [maxHewan]Hewan

// Fungsi main: inisialisasi data dan panggil menu
func main() {
	// Inisialisasi array dataHewan dengan tipe ArrHewan dan variabel jumlahData di dalam fungsi main
	var dataHewan ArrHewan // Menggunakan tipe ArrHewan
	var jumlahData int

	// Data dummy untuk diisi ke dalam array
	// Inisialisasi bisa langsung menggunakan tipe ArrHewan atau [maxHewan]Hewan karena kompatibel
	dataHewan = ArrHewan{
		{"K001", "Kucing", "Oyen", 3, "Aldi"},
		{"A002", "Anjing", "Bruno", 5, "Citra"},
		{"K003", "Kelinci", "Lulu", 2, "Dimas"},
		{"H004", "Burung", "Tweety", 1, "Erika"},
		{"H005", "Ikan", "Nemo", 1, "Fajar"},
		{"H006", "Kucing", "Milo", 4, "Gina"},
		{"H007", "Anjing", "Rocky", 6, "Hadi"},
		{"H008", "Kelinci", "Snowy", 2, "Indah"},
		{"H009", "Burung", "Rio", 3, "Joko"},
		{"H010", "Ikan", "Dory", 2, "Kiki"},
		{"H011", "Kucing", "Tom", 5, "Lina"},
		{"H012", "Anjing", "Spike", 7, "Mira"},
		{"H013", "Kelinci", "Bunny", 3, "Nina"},
		{"H014", "Burung", "Cici", 2, "Omar"},
		{"H015", "Ikan", "Goldy", 4, "Putri"},
		{"H016", "Kucing", "Garfield", 6, "Qory"},
		{"H017", "Anjing", "Snoopy", 5, "Rama"},
		{"H018", "Kelinci", "Fluffy", 1, "Salsa"},
		{"H019", "Burung", "Jacko", 3, "Tina"},
		{"H020", "Ikan", "Sharky", 2, "Umar"},
	}
	jumlahData = 20 // Set jumlah data awal sesuai data dummy

	// Panggil fungsi menu dengan mengoper pointer ke dataHewan dan jumlahData
	menu(&dataHewan, &jumlahData)
}

// Fungsi menu: menampilkan pilihan menu dan mengarahkan ke fungsi terkait
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func menu(dataHewan *ArrHewan, jumlahData *int) {
	for {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                       🐾 MENU UTAMA 🐾                                	     ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║                                       Created by                                                   ║")
		fmt.Println("║                                   Zaidan Kamil Munadi                                              ║")
		fmt.Println("║                                   Roby Ariga Siagian                                               ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ No │ Menu                                                                                          ║")
		fmt.Println("╠════╪═══════════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  1 │ Tambah Data Hewan                                                                             ║")
		fmt.Println("║  2 │ Tampilkan Semua Data                                                                          ║")
		fmt.Println("║  3 │ Hitung Rata-rata Umur                                                                         ║")
		fmt.Println("║  4 │ Cari Hewan Tertua dan Termuda                                                                 ║")
		fmt.Println("║  5 │ Sorting Berdasarkan Nama Pemilik (A-Z)                                                        ║")
		fmt.Println("║  6 │ Cari Hewan Berdasarkan Nama Panggilan                                                         ║")
		fmt.Println("║  7 │ Cari Hewan Berdasarkan ID                                                                     ║")
		fmt.Println("║  8 │ Edit Data Hewan                                                                               ║")
		fmt.Println("║  9 │ Hapus Data Hewan                                                                              ║")
		fmt.Println("║ 10 │ Urutkan Berdasarkan Umur Dari Tertua                                                          ║")
		fmt.Println("║ 11 │ Urutkan Berdasarkan Umur Dari Termuda                                                         ║")
		fmt.Println("║ 12 │ Statistik Berdasarkan Jenis                                                                   ║")
		fmt.Println("║  0 │ Keluar                                                                                        ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)
		// Proses pilihan menu
		if pilihan == 0 {
			fmt.Println("╔══════════════════════════════════════════════════════════════════════════════════════════════════==╗")
			fmt.Println("║                             🙏 TERIMA KASIH TELAH MENGGUNAKAN LAYANAN KAMI 🙏             	     ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════╝")
			return // Keluar dari loop dan fungsi menu
		} else if pilihan == 1 {
			tambahData(dataHewan, jumlahData)
		} else if pilihan == 2 {
			tampilkanData(dataHewan, jumlahData)
		} else if pilihan == 3 {
			rataRataUmur(dataHewan, jumlahData)
		} else if pilihan == 4 {
			temukanTuaMuda(dataHewan, jumlahData)
		} else if pilihan == 5 {
			sortByPemilik(dataHewan, jumlahData)
		} else if pilihan == 6 {
			searchHewan(dataHewan, jumlahData)
		} else if pilihan == 7 {
			binarySearchByID(dataHewan, jumlahData)
		} else if pilihan == 8 {
			editData(dataHewan, jumlahData)
		} else if pilihan == 9 {
			hapusData(dataHewan, jumlahData)
		} else if pilihan == 10 {
			selectionSortUmur(dataHewan, jumlahData)
		} else if pilihan == 11 {
			insertionSortUmur(dataHewan, jumlahData)
		} else if pilihan == 12 {
			statistikJenis(dataHewan, jumlahData)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk memeriksa apakah string mengandung angka
func mengandungAngka(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return false
}

// Fungsi untuk menambah data hewan baru
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func tambahData(dataHewan *ArrHewan, jumlahData *int) {
	var h Hewan

	if *jumlahData >= maxHewan {
		fmt.Println("Data penuh, tidak bisa menambahkan lagi.")
		return
	}

	fmt.Print("Masukkan ID Hewan: ")
	fmt.Scanln(&h.ID)
	if h.ID == "" {
		fmt.Println("ID tidak boleh kosong!")
		return
	}

	fmt.Print("Masukkan Jenis Hewan: ")
	fmt.Scanln(&h.Jenis)
	if h.Jenis == "" || mengandungAngka(h.Jenis) {
		fmt.Println("Jenis tidak boleh kosong atau mengandung angka!")
		return
	}

	fmt.Print("Masukkan Nama Panggilan Hewan: ")
	fmt.Scanln(&h.Nama)
	if h.Nama == "" || mengandungAngka(h.Nama) {
		fmt.Println("Nama tidak boleh kosong atau mengandung angka!")
		return
	}

	fmt.Print("Masukkan Umur Hewan: ")
	fmt.Scanln(&h.Umur)
	if h.Umur < 0 {
		fmt.Println("Umur tidak boleh negatif!")
		return
	}

	fmt.Print("Masukkan Nama Pemilik: ")
	fmt.Scanln(&h.Pemilik)
	if h.Pemilik == "" || mengandungAngka(h.Pemilik) {
		fmt.Println("Nama Pemilik tidak boleh kosong atau mengandung angka!")
		return
	}

	(*dataHewan)[*jumlahData] = h
	*jumlahData++
	fmt.Println("Data berhasil ditambahkan!")
}

// Fungsi untuk menampilkan semua data hewan
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func tampilkanData(dataHewan *ArrHewan, jumlahData *int) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════=╗")
	fmt.Println("║                                       📋 DATA HEWAN 📋                                     	  ║")
	fmt.Println("╠════╦═════════════╦═══════════════════╦══════════════════════╦═══════════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan    │ Jenis Hewan       │ Nama Hewan           │ Umur      │ Nama Pemilik          ║")
	fmt.Println("╠════╬═════════════╬═══════════════════╬══════════════════════╬═══════════╬═══════════════════════╣")

	for i := 0; i < *jumlahData; i++ {
		h := (*dataHewan)[i]
		fmt.Printf("║ %-2d │ %-11s │ %-17s │ %-20s │ %-9d │ %-21s ║\n",
			i+1, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
	}

	fmt.Println("╚════╩═════════════╩═══════════════════╩══════════════════════╩═══════════╩═══════════════════════╝")
}

// Fungsi untuk menghitung dan menampilkan rata-rata umur hewan
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func rataRataUmur(dataHewan *ArrHewan, jumlahData *int) {
	if *jumlahData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}
	total := 0
	for i := 0; i < *jumlahData; i++ {
		total += (*dataHewan)[i].Umur
	}
	rata := float64(total) / float64(*jumlahData)

	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                     📊 RATA-RATA UMUR HEWAN 📊                            	     ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-98s ║\n", fmt.Sprintf("Jumlah Data: %d", *jumlahData))
	fmt.Printf("║ %-98s ║\n", fmt.Sprintf("Rata-rata Umur: %.2f tahun", rata))
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

// Fungsi untuk menemukan dan menampilkan hewan tertua dan termuda
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func temukanTuaMuda(dataHewan *ArrHewan, jumlahData *int) {
	if *jumlahData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}
	tertua := (*dataHewan)[0]
	termuda := (*dataHewan)[0]
	for i := 1; i < *jumlahData; i++ {
		if (*dataHewan)[i].Umur > tertua.Umur {
			tertua = (*dataHewan)[i]
		}
		if (*dataHewan)[i].Umur < termuda.Umur {
			termuda = (*dataHewan)[i]
		}
	}

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                   🏆 HEWAN TERTUA DAN TERMUDA 🏆                      	║")
	fmt.Println("╠══════════╦═══════════════╦══════════════════╦══════════════════╦══════╦═══════════════════════╣")
	fmt.Println("║  Tipe    │ ID Hewan      │ Jenis Hewan      │ Nama Hewan       │ Umur │ Nama Pemilik          ║")
	fmt.Println("╠══════════╬═══════════════╬══════════════════╬══════════════════╬══════╬═══════════════════════╣")
	fmt.Printf("║ %-8s │ %-13s │ %-16s │ %-16s │ %-4d │ %-21s ║\n",
		"Tertua", tertua.ID, tertua.Jenis, tertua.Nama, tertua.Umur, tertua.Pemilik)
	fmt.Printf("║ %-8s │ %-13s │ %-16s │ %-16s │ %-4d │ %-21s ║\n",
		"Termuda", termuda.ID, termuda.Jenis, termuda.Nama, termuda.Umur, termuda.Pemilik)
	fmt.Println("╚══════════╩═══════════════╩══════════════════╩══════════════════╩══════╩═══════════════════════╝")
}

// Fungsi untuk mengurutkan data hewan berdasarkan nama pemilik (A-Z) menggunakan insertion sort
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func sortByPemilik(dataHewan *ArrHewan, jumlahData *int) {
	for i := 1; i < *jumlahData; i++ {
		temp := (*dataHewan)[i]
		j := i - 1
		for j >= 0 && (*dataHewan)[j].Pemilik > temp.Pemilik {
			(*dataHewan)[j+1] = (*dataHewan)[j]
			j--
		}
		(*dataHewan)[j+1] = temp
	}
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                             📊 Data Diurutkan Berdasarkan Nama Pemilik (A-Z) 📊            	     ║")
	fmt.Println("╠════╦══════════════╦══════════════════════╦══════════════════════════╦══════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan     │ Jenis Hewan          │ Nama Hewan               │ Umur │ Nama Pemilik          ║")
	fmt.Println("╠════╬══════════════╬══════════════════════╬══════════════════════════╬══════╬═══════════════════════╣")

	for i := 0; i < *jumlahData; i++ {
		h := (*dataHewan)[i]
		fmt.Printf("║ %-2d │ %-12s │ %-20s │ %-24s │ %-4d │ %-21s ║\n",
			i+1, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
	}
	fmt.Println("╚════╩══════════════╩══════════════════════╩══════════════════════════╩══════╩═══════════════════════╝")
}

// Fungsi untuk mencari hewan berdasarkan nama panggilan menggunakan sequential search
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func searchHewan(dataHewan *ArrHewan, jumlahData *int) {
	var nama string
	fmt.Print("Masukkan nama panggilan hewan yang dicari: (contoh: Oyen): ")
	fmt.Scanln(&nama)

	ditemukan := false

	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                       🔎 HASIL PENCARIAN 🔎                                        ║")
	fmt.Println("╠════╦══════════════╦══════════════════════╦══════════════════════════╦══════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan     │ Jenis Hewan          │ Nama Hewan               │ Umur │ Nama Pemilik          ║")
	fmt.Println("╠════╬══════════════╬══════════════════════╬══════════════════════════╬══════╬═══════════════════════╣")

	count := 0
	for i := 0; i < *jumlahData; i++ {
		if (*dataHewan)[i].Nama == nama {
			h := (*dataHewan)[i]
			count++
			fmt.Printf("║ %-2d │ %-12s │ %-20s │ %-24s │ %-4d │ %-21s ║\n",
				count, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("║                                      Data tidak ditemukan.                                        ║")
	}

	fmt.Println("╚════╩══════════════╩══════════════════════╩══════════════════════════╩══════╩═══════════════════════╝")
}

// Fungsi untuk mengurutkan data hewan berdasarkan ID menggunakan insertion sort (dibutuhkan untuk binary search)
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func sortByID(dataHewan *ArrHewan, jumlahData *int) {
	for i := 1; i < *jumlahData; i++ {
		temp := (*dataHewan)[i]
		j := i - 1
		for j >= 0 && (*dataHewan)[j].ID > temp.ID {
			(*dataHewan)[j+1] = (*dataHewan)[j]
			j--
		}
		(*dataHewan)[j+1] = temp
	}
}

// Fungsi untuk mencari hewan berdasarkan ID menggunakan binary search
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func binarySearchByID(dataHewan *ArrHewan, jumlahData *int) {
	var id string
	fmt.Print("Masukkan ID hewan yang dicari: (contoh K001): ")
	fmt.Scanln(&id)

	if *jumlahData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}

	sortByID(dataHewan, jumlahData)

	low := 0
	high := *jumlahData - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if (*dataHewan)[mid].ID == id {
			h := (*dataHewan)[mid] // Perbaikan: (*dataHewan) bukan (*dataHewAN)
			fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                                  🔍 HASIL PENCARIAN BINARY SEARCH 🔍                       ║")
			fmt.Println("╠═════════════╦═══════════════════╦══════════════════════╦═══════════╦═══════════════════════╣")
			fmt.Println("║ ID Hewan    │ Jenis Hewan       │ Nama Hewan           │ Umur      │ Nama Pemilik          ║")
			fmt.Println("╠═════════════╬═══════════════════╬══════════════════════╬═══════════╬═══════════════════════╣")
			fmt.Printf("║ %-11s │ %-17s │ %-20s │ %-9d │ %-21s ║\n", h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
			fmt.Println("╚═════════════╩═══════════════════╩══════════════════════╩═══════════╩═══════════════════════╝")
			found = true
			break
		} else if (*dataHewan)[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
	}
}

// Fungsi untuk mengedit data hewan berdasarkan ID
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func editData(dataHewan *ArrHewan, jumlahData *int) {
	var id string
	fmt.Print("Masukkan ID hewan yang ingin diedit (contoh K001): ")
	fmt.Scanln(&id)
	ditemukan := false

	for i := 0; i < *jumlahData; i++ {
		if (*dataHewan)[i].ID == id {
			ditemukan = true
			fmt.Println("Data ditemukan. Silakan masukkan data baru.")

			fmt.Print("Nama Baru: ")
			fmt.Scanln(&(*dataHewan)[i].Nama)
			if (*dataHewan)[i].Nama == "" || mengandungAngka((*dataHewan)[i].Nama) {
				fmt.Println("Nama tidak boleh kosong atau mengandung angka!")
				return
			}

			fmt.Print("Umur Baru: ")
			fmt.Scanln(&(*dataHewan)[i].Umur)
			if (*dataHewan)[i].Umur < 0 {
				fmt.Println("Umur tidak boleh negatif!")
				return
			}

			fmt.Print("Pemilik Baru: ")
			fmt.Scanln(&(*dataHewan)[i].Pemilik)
			if (*dataHewan)[i].Pemilik == "" || mengandungAngka((*dataHewan)[i].Pemilik) {
				fmt.Println("Nama Pemilik tidak boleh kosong atau mengandung angka!")
				return
			}

			fmt.Println("Data berhasil diperbarui!")
			return
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Fungsi untuk menghapus data hewan berdasarkan ID
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func hapusData(dataHewan *ArrHewan, jumlahData *int) {
	var id string
	fmt.Print("Masukkan ID hewan yang ingin dihapus: contoh K001: ")
	fmt.Scanln(&id)
	dihapus := false
	indeksDihapus := -1

	for i := 0; i < *jumlahData; i++ {
		if (*dataHewan)[i].ID == id {
			indeksDihapus = i
			dihapus = true
			break
		}
	}

	if dihapus {
		for j := indeksDihapus; j < *jumlahData-1; j++ {
			(*dataHewan)[j] = (*dataHewan)[j+1]
		}
		(*dataHewan)[*jumlahData-1] = Hewan{}
		*jumlahData--
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Fungsi untuk mengurutkan data hewan berdasarkan umur (Tertua -> Termuda) menggunakan selection sort
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func selectionSortUmur(dataHewan *ArrHewan, jumlahData *int) {
	var i, idx, pass int
	var temp Hewan

	for pass = 0; pass < *jumlahData-1; pass++ {
		idx = pass
		for i = pass + 1; i < *jumlahData; i++ {
			if (*dataHewan)[i].Umur > (*dataHewan)[idx].Umur {
				idx = i
			}
		}
		temp = (*dataHewan)[pass]
		(*dataHewan)[pass] = (*dataHewan)[idx]
		(*dataHewan)[idx] = temp
	}
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                             📉 Data Diurutkan Berdasarkan Umur (Tertua -> Termuda) 📉              ║")
	fmt.Println("╠════╦══════════════╦══════════════════════╦══════════════════════════╦══════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan     │ Jenis Hewan          │ Nama Hewan               │ Umur │ Nama Pemilik          ║")
	fmt.Println("╠════╬══════════════╬══════════════════════╬══════════════════════════╬══════╬═══════════════════════╣")

	for i = 0; i < *jumlahData; i++ {
		h := (*dataHewan)[i]
		fmt.Printf("║ %-2d │ %-12s │ %-20s │ %-24s │ %-4d │ %-21s ║\n",
			i+1, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
	}
	fmt.Println("╚════╩══════════════╩══════════════════════╩══════════════════════════╩══════╩═══════════════════════╝")
}

// Fungsi untuk mengurutkan data hewan berdasarkan umur (Termuda -> Tertua) menggunakan insertion sort
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func insertionSortUmur(dataHewan *ArrHewan, jumlahData *int) {
	for i := 1; i < *jumlahData; i++ {
		temp := (*dataHewan)[i]
		j := i - 1
		for j >= 0 && (*dataHewan)[j].Umur > temp.Umur {
			(*dataHewan)[j+1] = (*dataHewan)[j]
			j--
		}
		(*dataHewan)[j+1] = temp
	}
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                             📈 Data Diurutkan Berdasarkan Umur (Termuda -> Tertua) 📈              ║")
	fmt.Println("╠════╦══════════════╦══════════════════════╦══════════════════════════╦══════╦═══════════════════════╣")
	fmt.Println("║ No │ ID Hewan     │ Jenis Hewan          │ Nama Hewan               │ Umur │ Nama Pemilik          ║")
	fmt.Println("╠════╬══════════════╬══════════════════════╬══════════════════════════╬══════╬═══════════════════════╣")

	for i := 0; i < *jumlahData; i++ {
		h := (*dataHewan)[i]
		fmt.Printf("║ %-2d │ %-12s │ %-20s │ %-24s │ %-4d │ %-21s ║\n",
			i+1, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
	}
	fmt.Println("╚════╩══════════════╩══════════════════════╩══════════════════════════╩══════╩═══════════════════════╝")
}

// Fungsi untuk menampilkan statistik hewan berdasarkan jenisnya
// Parameter dataHewan sekarang menggunakan pointer ke tipe ArrHewan
func statistikJenis(dataHewan *ArrHewan, jumlahData *int) {
	const maxUniqueJenis = 10
	var jenisCounter [maxUniqueJenis]int
	var jenisNama [maxUniqueJenis]string
	var uniqueJenisSekarang int

	if *jumlahData == 0 {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                                     📈 STATISTIK BERDASARKAN JENIS 📈                           ║")
		fmt.Println("╠════╦═════════════════════════════════════════════════════════════╦═════════════════════════════╣")
		fmt.Println("║ No │                               Jenis Hewan                   │           Jumlah            ║")
		fmt.Println("╠════╬═════════════════════════════════════════════════════════════╬═════════════════════════════╣")
		fmt.Println("║           Tidak ada data hewan untuk ditampilkan statistik.                                     ║")
		fmt.Println("╚════╩═════════════════════════════════════════════════════════════╩═════════════════════════════╝")
		return
	}

	for i := 0; i < *jumlahData; i++ {
		jenisHewanSaatIni := (*dataHewan)[i].Jenis
		ditemukan := false
		indeksDitemukan := -1

		for j := 0; j < uniqueJenisSekarang; j++ {
			if jenisNama[j] == jenisHewanSaatIni {
				ditemukan = true
				indeksDitemukan = j
				break
			}
		}

		if ditemukan {
			jenisCounter[indeksDitemukan]++
		} else {
			if uniqueJenisSekarang < maxUniqueJenis {
				jenisNama[uniqueJenisSekarang] = jenisHewanSaatIni
				jenisCounter[uniqueJenisSekarang] = 1
				uniqueJenisSekarang++
			}
		}
	}

	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                     📈 STATISTIK BERDASARKAN JENIS 📈                          ║")
	fmt.Println("╠════╦═════════════════════════════════════════════════════════════╦═════════════════════════════╣")
	fmt.Println("║ No │                               Jenis Hewan                   │           Jumlah            ║")
	fmt.Println("╠════╬═════════════════════════════════════════════════════════════╬═════════════════════════════╣")

	if uniqueJenisSekarang == 0 && *jumlahData > 0 {
		fmt.Println("║              Tidak ada jenis hewan yang terdefinisi dalam data.                                 ║")
	} else {
		for i := 0; i < uniqueJenisSekarang; i++ {
			fmt.Printf("║ %-2d │ %-59s │ %27d ║\n", i+1, jenisNama[i], jenisCounter[i])
		}
	}
	fmt.Println("╚════╩═════════════════════════════════════════════════════════════╩═════════════════════════════╝")
}
