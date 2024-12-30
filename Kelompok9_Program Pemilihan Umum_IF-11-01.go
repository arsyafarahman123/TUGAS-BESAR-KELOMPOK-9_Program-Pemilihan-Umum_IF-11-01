package main
import (
	"fmt"
	"strings"
)
const NMAX = 1000 //maksimal data calon dan pemilih

type Calon struct { // Struktur data calon
	namacalon  string
	namapartai string
	Suara      int
}
type Pemilih struct { //Struktur data pemilih
	namapemilih  string
	SudahMemilih bool
	waktu        int
	Pilihan      string
}
var ( //variabel global
	daftarCalon   [NMAX]Calon
	daftarPemilih [NMAX]Pemilih
	jumlahCalon   int
	jumlahPemilih int
	threshold     int
	durasi int
)
// Menu petugas
func menuPetugas() {
	for {
		fmt.Println("\n================== PEMILU APP - MENU PAGE (ADMIN) ==================:")
		fmt.Println("1. Tampilkan semua daftar calon (berurutan berdasarkan abjad)")
		fmt.Println("2. Cari data calon berdasarkan")
		fmt.Println("3. Tampilkan perolehan suara calon (berurutan berdasarkan perolehan suara)")
		fmt.Println("4. Pengoperasian data calon")
		fmt.Println("5. Pengoperasian data pemilih")
		fmt.Println("6. Masukkan threshold calon")                
		fmt.Println("7. Set waktu")                               
		fmt.Println("8. Tampilkan calon yang memenuhi threshold") 
		fmt.Println("9. Logout")
		var pilihan int
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&pilihan)
		// Switch case pilihan admin
		switch pilihan {
		case 1:
			tampilkanSemuaCalon(false) //Menampilkan calon berdasarkan abjad(false = descending, true= ascending)
		case 2:
			cariDataCalon()
		case 3:
			tampilkanPerolehanSuara(false) // Menampilkan perolehan suara calon berdasarkan suara(false = descending, true= ascending)
		case 4:
			pengoperasianDataCalon()
		case 5:
			pengoperasianDataPemilih()
		case 6:
			masukkanThreshold()
		case 7:
			setWaktu()
		case 8:
			tampilkanCalonLolosThreshold()
		case 9:
			fmt.Println("Berhasil Logout")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
// menu Pemilih
func menuPemilih() {
	for {
		fmt.Println("\n================== PEMILU APP - MENU PAGE (PEMILIH) ==================:")
		fmt.Println("1. Tampilkan semua daftar calon (berurutan berdasarkan abjad)")
		fmt.Println("2. Cari data calon berdasarkan")
		fmt.Println("3. Tampilkan perolehan suara calon (berurutan berdasarkan perolehan suara)")
		fmt.Println("4. Pemilihan calon")
		fmt.Println("5. Tampilkan calon yang memenuhi threshold")
		fmt.Println("9. Logout")
		var pilihan int
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanSemuaCalon(true)
		case 2:
			cariDataCalon()
		case 3:
			tampilkanPerolehanSuara(true)
		case 4:
			pemilihanCalon()
		case 5:
			tampilkanCalonLolosThreshold()
		case 9:
			fmt.Println("Berhasil Logout")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
// Fungsi menampilkan semua calon
func tampilkanSemuaCalon(ascending bool) {
	// Cek apakah jumlah calon 0
	if jumlahCalon == 0 {
		// Jika tidak ada calon, tampilkan pesan dan keluar dari fungsi
		fmt.Println("Tidak ada calon.")
		return
	}
	for i := 0; i < jumlahCalon-1; i++ {
		for j := i + 1; j < jumlahCalon; j++ {
			//jika ascending true, urutkan kecil ke besar
			//jjika ascending false, urutkan besar ke kecil
			if (ascending && strings.Compare(daftarCalon[i].namacalon, daftarCalon[j].namacalon) > 0) ||
				(!ascending && strings.Compare(daftarCalon[i].namacalon, daftarCalon[j].namacalon) < 0) {
				daftarCalon[i], daftarCalon[j] = daftarCalon[j], daftarCalon[i]
			}
		}
	}
	fmt.Println("Daftar Calon:")
	for i := 0; i < jumlahCalon; i++ {
		// tampil calon dan partainya
		fmt.Printf("%d. %s (%s)\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai)
	}
}
// Cari Data Calon berdasarkan
func cariDataCalon() {
	fmt.Println("\n1. Cari berdasarkan nama calon")
	fmt.Println("2. Cari berdasarkan nama partai")
	fmt.Println("3. Cari berdasarkan nama pemilih ")
	var pilihan int
	fmt.Print("Pilih pencarian: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&nama)
		cariCalon(nama)
	case 2:
		var partai string
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&partai)
		cariPartai(partai)
	case 3:
		var nama string
		fmt.Print("Masukkan nama pemilih: ")
		fmt.Scan(&nama)
		cariPeserta(nama)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariCalon(nama string) {
	// Loop untuk mencari calon dalam daftar
	for i := 0; i < jumlahCalon; i++ {
		// Jika nama calon ditemukan, tampilkan informasi calon
		if strings.EqualFold(daftarCalon[i].namacalon, nama) {
			fmt.Printf("Nama: %s, Partai: %s, Suara: %d\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			return
		}
	} // Jika calon tidak ditemukan, tampilkan pesan
	fmt.Println("Calon tidak ditemukan.")
}

func cariPartai(partai string) {
	// Variabel untuk menandai apakah partai ditemukan
	for i := 0; i < jumlahCalon; i++ {
		// Jika nama partai ditemukan, tampilkan informasi calon
		if strings.EqualFold(daftarCalon[i].namapartai, partai) {
			fmt.Printf("Nama: %s, Partai: %s, Suara: %d\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			//keluar fungsi setelah menampilkan informasi calon
			return
		}
	} // Jika partai tidak ditemukan, tampilkan pesan
	fmt.Println("Partai tidak ditemukan.")
}

func cariPeserta(namaPemilih string) {
	//loop mencari pemilih dalam daftar
	for i := 0; i < jumlahPemilih; i++ {
		if strings.EqualFold(daftarPemilih[i].namapemilih, namaPemilih) && daftarPemilih[i].SudahMemilih {
			//loop mencari calon berdasarkan pilihan pemilih
			for j := 0; j < jumlahCalon; j++ {
				if strings.EqualFold(daftarPemilih[i].Pilihan, daftarCalon[j].namacalon) {
					fmt.Printf("Pemilih %s memilih = calon: %s, Partai: %s, Suara: %d\n",namaPemilih, daftarCalon[j].namacalon, daftarCalon[j].namapartai, daftarCalon[j].Suara)
					//keluar fungsi setelah menampilkan informasi calon
					return
				}
			}
		}
	}//jika pemilih tidak ditemukan atau belum memilih
	fmt.Println("Pemilih tidak ditemukan atau belum memilih.")
}

// Fungsi untuk menampilkan perolehan suara calon berdasarkan urutan suara
func tampilkanPerolehanSuara(descending bool) {
	if jumlahCalon == 0 {
		// Jika tidak ada calon, tampilkan pesan dan keluar dari fungsi
		fmt.Println("Tidak ada calon.")
		return
	}
	//Bubble Sort untuk mengurutkan calon berdasarkan  suara
	for i := 0; i < jumlahCalon-1; i++ { // Loop untuk iterasi calon
		for j := i + 1; j < jumlahCalon; j++ { // Loop untuk iterasi calon selanjutnya
			if (descending && daftarCalon[i].Suara < daftarCalon[j].Suara) ||
				(!descending && daftarCalon[i].Suara > daftarCalon[j].Suara) { // Jika descending true, urutkan secara menurun
				daftarCalon[i], daftarCalon[j] = daftarCalon[j], daftarCalon[i] // Jika descending false, urutkan secara naik
			}
		}
	}
	fmt.Println("Perolehan Suara:")
	for i := 0; i < jumlahCalon; i++ { // Loop untuk menampilkan perolehan suara calon
		fmt.Printf("%d. %s (%s): %d suara\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
	}
}

//fungsi tampil pemilih
func tampilkanpemilih(){
	var status string
	fmt.Println("Daftar pemilih:")
	for i:=0; i<jumlahPemilih; i++ {
		//menentukan status hak pilih
		if daftarPemilih[i].SudahMemilih==true{
			status = "sudah memilih"
		}else{
			status = "belum memilih"
		}
		fmt.Printf("%d. nama = %s (%s)\n", i+1, daftarPemilih[i].namapemilih, status)
	}
}
//Pemilihan Calon
func pemilihanCalon() {
	var nama string
	if durasi > 0 {
		fmt.Printf("Durasi Pemilu: %d menit\n", durasi)
		tampilkanpemilih()
		fmt.Print("Masukkan nama pemilih: ")
		fmt.Scan(&nama)
		for i := 0; i < jumlahPemilih; i++ {
			if strings.EqualFold(daftarPemilih[i].namapemilih, nama) {
				//jika pemilih sudah menggunakan hak pilihnya
				if daftarPemilih[i].SudahMemilih {
					fmt.Println("Anda sudah melakukan pemilihan.")
					return
				}
				//input durasi pemilihan calon
				fmt.Print("Masukkan lama waktu pemilihan: ")
				fmt.Scan(&daftarPemilih[i].waktu)
				//jika durasi oemilihan tidak mencukupi return daftar calon
				if durasi-daftarPemilih[i].waktu < 0 {
					fmt.Println("Waktu tidak cukup untuk melakukan pemilihan.")
					tampilkanSemuaCalon(true)
					return
				}
				//tampilkan calon
				fmt.Println("\nDaftar Calon:")
				for i := 0; i < jumlahCalon; i++ {
					fmt.Printf("%d. %s (%s)\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai)
				}
				//pilih nomor calon
				fmt.Print("Masukkan nomor calon yang dipilih: ")
				var pilihan int
				fmt.Scan(&pilihan)
				//jika nomor calon salah
				if pilihan < 1 || pilihan > jumlahCalon {
					fmt.Println("Nomor calon tidak valid")
					return
				}
				//update infromasi pemilih dan calon
				daftarCalon[pilihan-1].Suara++
				daftarPemilih[i].SudahMemilih = true
				daftarPemilih[i].Pilihan = daftarCalon[pilihan-1].namacalon
				durasi -= daftarPemilih[i].waktu
				fmt.Println("Pemilihan berhasil disimpan.")
				return
			}
		}
		fmt.Println("Pemilih tidak ditemukan.")
	}else{
		fmt.Println("\nDiluar waktu pemilihan! Set waktu pemilihan!")
		tampilkanSemuaCalon(true)
	}	
}

// Pengoperasian Data Calon
func tambahDataCalon() {
	if jumlahCalon == NMAX { // Cek apakah jumlah calon sudah mencapai batas maksimum
		fmt.Println("Jumlah calon sudah maksimal.") // Jika sudah mencapai batas, tampilkan pesan
		return
	}
	//input data calon baru
	fmt.Print("Masukkan nama calon: ")
	fmt.Scan(&daftarCalon[jumlahCalon].namacalon)
	fmt.Print("Masukkan nama partai: ")
	fmt.Scan(&daftarCalon[jumlahCalon].namapartai)
	daftarCalon[jumlahCalon].Suara = 0 //set suara calon baru ke 0
	jumlahCalon++
	fmt.Println("Calon berhasil ditambahkan.")
}

func editDataCalon() {
	//variabel untuk calon yang diedit
	var nama string
	fmt.Print("Masukkan nama calon yang ingin diubah: ")
	fmt.Scan(&nama)
	//
	for i := 0; i < jumlahCalon; i++ {
		if strings.EqualFold(daftarCalon[i].namacalon, nama) {//perbandingan string, jika nama calon terdapat pada daftarcalon
			//jika calon ditemukan, minta input untuk nama baru dan partai baru
			fmt.Print("Masukkan nama calon baru: ")
			fmt.Scan(&daftarCalon[i].namacalon)
			fmt.Print("Masukkan nama partai baru: ")
			fmt.Scan(&daftarCalon[i].namapartai)
			fmt.Println("Calon berhasil diubah.")
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

func hapusDataCalon() {
	var nama string
	fmt.Print("Masukkan nama calon yang ingin dihapus: ")
	fmt.Scan(&nama)
	for i := 0; i < jumlahCalon; i++ {
		if strings.EqualFold(daftarCalon[i].namacalon, nama) {//perbandingan string, jika nama calon terdapat pada daftarcalon
			// Menampilkan konfirmasi sebelum menghapus calon
			fmt.Print("Calon dengan nama " + nama + " dan partai " + daftarCalon[i].namapartai + " akan dihapus. Yakin? (y/n)")
			var konfirmasi string
			fmt.Scan(&konfirmasi)
			if konfirmasi != "y" {
				fmt.Println("Calon tidak jadi dihapus.")
				return
			}
			for j := i; j < jumlahCalon-1; j++ {
				daftarCalon[j] = daftarCalon[j+1] //menggeser calon setelah calon yang dihapus

			}
			jumlahCalon--
			fmt.Println("Calon berhasil dihapus.")
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

//sett durasi pemilihan
func setWaktu() {
	fmt.Print("Masukkan durasi pemilihan (Menit): ")
	fmt.Scan(&durasi) //input durasi pemilu
	fmt.Printf("Waktu berhasil diset ke Menit %d\n", durasi)
}

func pengoperasianDataCalon() { //Fungsi untuk mengoperasikan data calon (menambah, mengedit, menghapus)
	fmt.Println("\n================== PEMILU APP - PENGOPERASIAN DATA CALON ==================:")
	fmt.Println("1. Tambah data calon")
	fmt.Println("2. Edit data calon")
	fmt.Println("3. Hapus data calon")

	var pilihan int
	fmt.Print("Pilih operasi: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahDataCalon()
	case 2:
		editDataCalon()
	case 3:
		hapusDataCalon()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tambahDataPemilih() {
	if jumlahPemilih == NMAX {
		fmt.Println("Jumlah pemilih sudah maksimal.")
		return
	}
	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&daftarPemilih[jumlahPemilih].namapemilih)
	daftarPemilih[jumlahPemilih].SudahMemilih = false
	jumlahPemilih++
	fmt.Println("Pemilih berhasil ditambahkan.")
}

func editDataPemilih() {
	var nama string
	fmt.Print("Masukkan nama pemilih yang ingin diubah: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlahPemilih; i++ {
		if strings.EqualFold(daftarPemilih[i].namapemilih, nama) {
			fmt.Print("Masukkan nama pemilih baru: ")
			fmt.Scan(&daftarPemilih[i].namapemilih)
			fmt.Println("Pemilih berhasil diubah.")
			return
		}
	}
	fmt.Println("Pemilih tidak ditemukan.")
}

func hapusDataPemilih() {
	var nama string
	fmt.Print("Masukkan nama pemilih yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlahPemilih; i++ {
		if strings.EqualFold(daftarPemilih[i].namapemilih, nama) {
			fmt.Print("Pemilih dengan nama " + nama + " akan dihapus. Yakin? (y/n)")
			var konfirmasi string
			fmt.Scan(&konfirmasi)
			if konfirmasi != "y" {
				fmt.Println("Pemilih tidak jadi dihapus.")
				return
			}
			for j := i; j < jumlahPemilih-1; j++ {
				daftarPemilih[j] = daftarPemilih[j+1]
			}
			jumlahPemilih--
			fmt.Println("Pemilih berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pemilih tidak ditemukan.")
}

// Pengoperasian Data Pemilih
func pengoperasianDataPemilih() {
	fmt.Println("\n================== PEMILU APP - PENGOPERASIAN DATA PEMILIH ==================:")
	fmt.Println("1. Tambah data pemilih")
	fmt.Println("2. Edit data pemilih")
	fmt.Println("3. Hapus data pemilih")

	var pilihan int
	fmt.Print("Pilih operasi: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahDataPemilih()
	case 2:
		editDataPemilih()
	case 3:
		hapusDataPemilih()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Masukkan Threshold
func masukkanThreshold() {
	fmt.Print("Masukkan threshold: ")
	fmt.Scan(&threshold)
	fmt.Printf("Set Threshold calon %d%%\n", threshold)
}

// menampilkan calon berhak terpilih
func tampilkanCalonLolosThreshold() {
	totalSuara := 0
	for i := 0; i < jumlahCalon; i++ {
		totalSuara += daftarCalon[i].Suara
	}
	nilaiTHRSHLD := totalSuara * threshold / 100
	if nilaiTHRSHLD != 0{
		fmt.Printf("\nTotal suara sah sementara: %d\n", totalSuara)
		fmt.Printf("Ambang batas suara calon berdasarkan threshold %d%%: %d\n", threshold, nilaiTHRSHLD)
		fmt.Println("Calon berhak terpilih( ✔ ) dan tidak terpilih( ✖ ):")
		for i := 0; i < jumlahCalon; i++ {
			if daftarCalon[i].Suara > nilaiTHRSHLD {
				fmt.Printf("✔ %s dari partai %s dengan suara %d\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			}
			if daftarCalon[i].Suara < nilaiTHRSHLD {
				fmt.Printf("✖ %s dari partai %s dengan suara %d\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			}
	
		}

	}else{
		fmt.Println("Threshold calon belum ditentukan!")
	}
}

// data awal calon
func dummycalon() {
	daftarCalon[0] = Calon{namacalon: "Prasetyo", namapartai: "pks", Suara: 70}
	daftarCalon[1] = Calon{namacalon: "Aidit", namapartai: "pki", Suara: 90}
	daftarCalon[2] = Calon{namacalon: "Yunus", namapartai: "dpi", Suara: 100}
	daftarCalon[3] = Calon{namacalon: "Kesang", namapartai: "psi", Suara: 150}
	daftarCalon[4] = Calon{namacalon: "Hasan", namapartai: "psj", Suara: 15}
	jumlahCalon = 5
}

// data awal pemilih
func dummypemilih() {
	daftarPemilih[0] = Pemilih{namapemilih: "rini", SudahMemilih: false, waktu: 0}
	daftarPemilih[1] = Pemilih{namapemilih: "bobi", SudahMemilih: false, waktu: 0}
	daftarPemilih[2] = Pemilih{namapemilih: "dian", SudahMemilih: false, waktu: 0}
	daftarPemilih[3] = Pemilih{namapemilih: "adam", SudahMemilih: false, waktu: 0}
	daftarPemilih[4] = Pemilih{namapemilih: "vivi", SudahMemilih: false, waktu: 0}
	jumlahPemilih = 5
}

// Main Program
func main() {
	dummycalon()
	dummypemilih()
	for {
		fmt.Println("\n================== PEMILU APP - LOGIN PAGE ==================:")
		fmt.Println("Pilih user:")
		fmt.Println("1. Pemilih")
		fmt.Println("2. Petugas KPU")

		var pilihan int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuPemilih()
		} else if pilihan == 2 {
			menuPetugas()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
