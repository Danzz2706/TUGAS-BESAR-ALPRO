package main

import "fmt"

type Hewan struct {
	ID     string
	Jenis  string
	Nama   string
	Umur   int
	Pemilik string
}

const maxHewan = 30
var dataHewan [maxHewan]Hewan
var jumlahData int

func main() {
	menu()
}

func menu() {
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Data Hewan")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Hitung Rata-rata Umur")
		fmt.Println("4. Cari Hewan Tertua dan Termuda")
		fmt.Println("5. Urutkan Berdasarkan Umur")
		fmt.Println("6. Cari Hewan Berdasarkan Nama")
		fmt.Println("7. Edit Data Hewan")
		fmt.Println("8. Hapus Data Hewan")
		fmt.Println("9. Statistik Per Jenis")
		fmt.Println("10. Sorting Berdasarkan Nama Pemilik (A-Z)")
		fmt.Println("11. Urutkan Berdasarkan Umur (Insertion Sort)")
		fmt.Println("0. Keluar")
		fmt.Println("Pilih menu: ")

		var pilihan int
		fmt.Print("SELECT > ")
		fmt.Scanln(&pilihan)
		if pilihan == 0 {
			fmt.Println("Keluar dari program.")
			return
		} else if pilihan == 1 {
			tambahData()
		} else if pilihan == 2 {
			tampilkanData()
		} else if pilihan == 3 {
			rataRataUmur()
		} else if pilihan == 4 {
			temukanTuaMuda()
		} else if pilihan == 5 {
			urutkanUmur()
		} else if pilihan == 6 {
			searchHewan()
		} else if pilihan == 7 {
			editData()
		} else if pilihan == 8 {
			hapusData()
		} else if pilihan == 9 {
			statistikJenis()
		} else if pilihan == 10 {
			sortByPemilik()
		} else if pilihan == 11 {
			insertionSortUmur()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahData() {
	var h Hewan
	if jumlahData >= maxHewan {
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
	if h.Jenis == "" {
		fmt.Println("Jenis tidak boleh kosong!")
		return
	}
	fmt.Print("Masukkan Nama Hewan: ")
	fmt.Scanln(&h.Nama)
	if h.Nama == "" {
		fmt.Println("Nama tidak boleh kosong!")
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
	if h.Pemilik == "" {
		fmt.Println("Nama Pemilik tidak boleh kosong!")
		return
	}
	dataHewan[jumlahData] = h
	jumlahData++
	fmt.Println("Data berhasil ditambahkan!")
}

func tampilkanData() {
	fmt.Println("\n=== DATA HEWAN ===")
	for i := 0; i < jumlahData; i++ {
		h := dataHewan[i]
		fmt.Printf("%d. ID: %s, Jenis: %s, Nama: %s, Umur: %d, Pemilik: %s\n", i+1, h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
	}
}

func rataRataUmur() {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}
	total := 0
	for i := 0; i < jumlahData; i++ {
		total += dataHewan[i].Umur
	}
	rata := float64(total) / float64(jumlahData)
	fmt.Printf("Rata-rata umur hewan: %.2f\n", rata)
}

func temukanTuaMuda() {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data.")
		return
	}
	tertua := dataHewan[0]
	termuda := dataHewan[0]
	for i := 1; i < jumlahData; i++ {
		if dataHewan[i].Umur > tertua.Umur {
			tertua = dataHewan[i]
		}
		if dataHewan[i].Umur < termuda.Umur {
			termuda = dataHewan[i]
		}
	}
	fmt.Printf("Hewan tertua: %s (%d tahun)\n", tertua.Nama, tertua.Umur)
	fmt.Printf("Hewan termuda: %s (%d tahun)\n", termuda.Nama, termuda.Umur)
}

func urutkanUmur() {
	for i := 0; i < jumlahData-1; i++ {
		for j := 0; j < jumlahData-i-1; j++ {
			if dataHewan[j].Umur > dataHewan[j+1].Umur {
				temp := dataHewan[j]
				dataHewan[j] = dataHewan[j+1]
				dataHewan[j+1] = temp
			}
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan umur.")
}

func searchHewan() {
	var nama string
	fmt.Print("Masukkan nama hewan yang dicari: ")
	fmt.Scanln(&nama)
	ditemukan := false
	for i := 0; i < jumlahData; i++ {
		if dataHewan[i].Nama == nama {
			h := dataHewan[i]
			fmt.Printf("Ditemukan: ID: %s, Jenis: %s, Nama: %s, Umur: %d, Pemilik: %s\n", h.ID, h.Jenis, h.Nama, h.Umur, h.Pemilik)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func editData() {
	var id string
	fmt.Print("Masukkan ID hewan yang ingin diedit: ")
	fmt.Scanln(&id)
	for i := 0; i < jumlahData; i++ {
		if dataHewan[i].ID == id {
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scanln(&dataHewan[i].Nama)
			fmt.Print("Masukkan Umur Baru: ")
			fmt.Scanln(&dataHewan[i].Umur)
			fmt.Print("Masukkan Pemilik Baru: ")
			fmt.Scanln(&dataHewan[i].Pemilik)
			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}

func hapusData() {
	var id string
	fmt.Print("Masukkan ID hewan yang ingin dihapus: ")
	fmt.Scanln(&id)
	index := -1
	for i := 0; i < jumlahData; i++ {
		if dataHewan[i].ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for i := index; i < jumlahData-1; i++ {
		dataHewan[i] = dataHewan[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus!")
}

func statistikJenis() {
	jenisMap := make(map[string]int)
	for i := 0; i < jumlahData; i++ {
		jenisMap[dataHewan[i].Jenis]++
	}
	fmt.Println("\nStatistik jumlah berdasarkan jenis hewan:")
	for jenis, jumlah := range jenisMap {
		fmt.Printf("%s: %d\n", jenis, jumlah)
	}
}

func sortByPemilik() {
	for i := 1; i < jumlahData; i++ {
		temp := dataHewan[i]
		j := i - 1
		for j >= 0 && dataHewan[j].Pemilik > temp.Pemilik {
			dataHewan[j+1] = dataHewan[j]
			j--
		}
		dataHewan[j+1] = temp
	}
	fmt.Println("Data diurutkan berdasarkan nama pemilik.")
}

func insertionSortUmur() {
	for i := 1; i < jumlahData; i++ {
		temp := dataHewan[i]
		j := i - 1
		for j >= 0 && dataHewan[j].Umur > temp.Umur {
			dataHewan[j+1] = dataHewan[j]
			j--	
		}
		dataHewan[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan umur (insertion sort).")
}
