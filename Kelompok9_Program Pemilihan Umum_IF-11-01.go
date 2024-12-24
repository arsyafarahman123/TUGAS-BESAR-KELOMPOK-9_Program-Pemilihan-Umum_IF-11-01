package main // Package utama untuk menjalankan program

import (
	"fmt"
	"strings"
)

const NMAX = 1000 // Deklarasi konstanta maksimal data calon dan pemilih

type Calon struct { // Struktur data untuk menyimpan informasi calon
	namacalon  string
	namapartai string
	Suara      int
}

type Pemilih struct { // Struktur data untuk menyimpan informasi pemilih
	namapemilih  string
	SudahMemilih bool
	waktu        int
}

var ( // Deklarasi variabel global
	daftarCalon   [NMAX]Calon
	daftarPemilih [NMAX]Pemilih
	jumlahCalon   int
	jumlahPemilih int
	threshold     int
	waktuSekarang int
)

// Menu Pemilih
func menuPetugas() {
	// Perulangan untuk menampilkan menu sampai admin logout
	for { // Menampilkan menu utama admin
		fmt.Println("\n================== PEMILU APP - MENU PAGE (ADMIN) ==================:")
		fmt.Println("1. Tampilkan semua daftar calon (berurutan berdasarkan abjad)")
		fmt.Println("2. Cari data calon/pemilih")
		fmt.Println("3. Tampilkan perolehan suara calon (berurutan berdasarkan perolehan suara)")
		fmt.Println("4. Pengoperasian data calon")
		fmt.Println("5. Pengoperasian data pemilih")
		fmt.Println("6. Masukkan threshold calon")                // Memasukkan ambang batas suara calon
		fmt.Println("7. Set waktu")                               // Menyetel waktu pemilihan
		fmt.Println("8. Tampilkan calon yang memenuhi threshold") // Menu baru
		fmt.Println("9. Logout")

		var pilihan int
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&pilihan)
		// Switch case untuk menangani pilihan admin
		switch pilihan {
		case 1:
			tampilkanSemuaCalon(true) // Menampilkan daftar calon (urut abjad)
		case 2:
			cariDataCalonPemilih()
		case 3:
			tampilkanPerolehanSuara(true) // Menampilkan suara calon (urut suara)
		case 4:
			pengoperasianDataCalon() // Operasi data calon (tambah, edit, hapus)
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

// Update menu Pemilih
func menuPemilih() {
	// Perulangan untuk menampilkan menu sampai pemilih logout
	for {
		fmt.Println("\n================== PEMILU APP - MENU PAGE (PEMILIH) ==================:")
		fmt.Println("1. Tampilkan semua daftar calon (berurutan berdasarkan abjad)")
		fmt.Println("2. Cari data calon/pemilih")
		fmt.Println("3. Tampilkan perolehan suara calon (berurutan berdasarkan perolehan suara)")
		fmt.Println("4. Pemilihan calon")
		fmt.Println("5. Tampilkan calon yang memenuhi threshold") // Menu baru
		fmt.Println("9. Logout")

		var pilihan int
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanSemuaCalon(true)
		case 2:
			cariDataCalonPemilih()
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

// Fungsi untuk menampilkan daftar calon, terurut berdasarkan nama calon
func tampilkanSemuaCalon(ascending bool) {
	// Cek apakah jumlah calon adalah 0
	if jumlahCalon == 0 {
		// Jika tidak ada calon, tampilkan pesan dan keluar dari fungsi
		fmt.Println("Tidak ada calon.")
		return
	}
	for i := 0; i < jumlahCalon-1; i++ {
		for j := i + 1; j < jumlahCalon; j++ {
			// Jika ascending true, urutkan secara menaik
			// Jika ascending false, urutkan secara menurun
			if (ascending && strings.Compare(daftarCalon[i].namacalon, daftarCalon[j].namacalon) > 0) ||
				(!ascending && strings.Compare(daftarCalon[i].namacalon, daftarCalon[j].namacalon) < 0) {
				daftarCalon[i], daftarCalon[j] = daftarCalon[j], daftarCalon[i]
			}
		}
	}
	fmt.Println("Daftar Calon:")
	for i := 0; i < jumlahCalon; i++ {
		// Menampilkan nama calon dan partainya
		fmt.Printf("%d. %s (%s)\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai)
	}
}

// Cari Data Calon atau Pemilih
func cariDataCalonPemilih() {
	// Menampilkan pilihan pencarian kepada pengguna
	fmt.Println("\n1. Cari berdasarkan nama calon")
	fmt.Println("2. Cari berdasarkan nama partai")
	fmt.Println("3. Cari peserta berdasarkan nama calon")
	// Variabel untuk menyimpan pilihan pengguna
	var pilihan int
	// Minta input pilihan dari pengguna
	fmt.Print("Pilih pencarian: ")
	fmt.Scan(&pilihan)
	// Proses pencarian berdasarkan pilihan yang diberikan
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
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&nama)
		cariPeserta(nama)
	default: // Jika pilihan tidak valid
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
	found := false
	for i := 0; i < jumlahCalon; i++ {
		// Jika nama partai ditemukan, tampilkan informasi calon
		if strings.EqualFold(daftarCalon[i].namapartai, partai) {
			fmt.Printf("Nama: %s, Partai: %s, Suara: %d\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			// Tandai bahwa partai ditemukan
			found = true
		}
	} // Jika partai tidak ditemukan, tampilkan pesan
	if !found {
		fmt.Println("Partai tidak ditemukan.")
	}
}

func cariPeserta(nama string) {
	found := false
	// Loop untuk mencari pemilih yang sudah memilih
	for i := 0; i < jumlahPemilih; i++ {
		// Jika pemilih belum memilih, lanjutkan ke pemilih berikutnya
		if !daftarPemilih[i].SudahMemilih {
			continue
		}
		// Jika nama pemilih ditemukan, tampilkan nama pemilih
		if strings.EqualFold(daftarPemilih[i].namapemilih, nama) {
			fmt.Printf("Pemilih: %s\n", daftarPemilih[i].namapemilih)
			found = true
		}
	}
	if !found {
		fmt.Println("Pemilih tidak ditemukan.")
	}
}

// Fungsi untuk menampilkan perolehan suara calon berdasarkan urutan suara
func tampilkanPerolehanSuara(ascending bool) {
	if jumlahCalon == 0 {
		// Jika tidak ada calon, tampilkan pesan dan keluar dari fungsi
		fmt.Println("Tidak ada calon.")
		return
	}
	// Algoritma Bubble Sort untuk mengurutkan calon berdasarkan perolehan suara
	for i := 0; i < jumlahCalon-1; i++ {
		for j := i + 1; j < jumlahCalon; j++ {
			// Jika ascending true, urutkan secara menaik berdasarkan suara
			// Jika ascending false, urutkan secara menurun berdasarkan suara
			if (ascending && daftarCalon[i].Suara > daftarCalon[j].Suara) ||
				(!ascending && daftarCalon[i].Suara < daftarCalon[j].Suara) {
				daftarCalon[i], daftarCalon[j] = daftarCalon[j], daftarCalon[i]
			}
		}
	}
	// Tampilkan perolehan suara calon setelah diurutkan
	fmt.Println("Perolehan Suara:")
	for i := 0; i < jumlahCalon; i++ {
		// Menampilkan nama calon, nama partai, dan jumlah suara yang diterima
		fmt.Printf("%d. %s (%s): %d suara\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
	}
}

// Pemilihan Calon
func pemilihanCalon() {
	// Minta input nama pemilih
	var nama string
	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&nama)

	// Cari pemilih berdasarkan nama
	var pemilihIndex = -1
	for i := 0; i < jumlahPemilih; i++ {
		// Jika ditemukan pemilih berdasarkan nama, simpan index pemilih
		if strings.EqualFold(daftarPemilih[i].namapemilih, nama) {
			pemilihIndex = i
			break
		}
	}
	// Jika pemilih tidak ditemukan, tampilkan pesan dan keluar dari fungsi
	if pemilihIndex == -1 {
		fmt.Println("Pemilih tidak ditemukan.")
		return
	}

	// Cek apakah pemilih sudah memilih sebelumnya
	if daftarPemilih[pemilihIndex].SudahMemilih {
		// Jika pemilih sudah memilih, tampilkan pesan dan keluar dari fungsi
		fmt.Println("Anda sudah melakukan pemilihan sebelumnya.")
		return
	}

	// Cek apakah waktu pemilihan sesuai dengan jadwal
	if daftarPemilih[pemilihIndex].waktu != waktuSekarang {
		// Jika waktu pemilih tidak sesuai, tampilkan pesan
		fmt.Printf("Anda tidak dapat memilih pada jam ini. Jadwal pemilihan anda adalah jam %d\n",
			daftarPemilih[pemilihIndex].waktu)
		return
	}

	// Tampilkan daftar calon yang tersedia untuk dipilih
	fmt.Println("\nDaftar Calon:")
	for i := 0; i < jumlahCalon; i++ {
		// Tampilkan nomor calon, nama calon, dan nama partai
		fmt.Printf("%d. %s (%s)\n", i+1, daftarCalon[i].namacalon, daftarCalon[i].namapartai)
	}

	// Proses pemilihan oleh pemilih
	var pilihan int
	fmt.Print("Masukkan nomor calon yang dipilih: ") // Meminta input nomor calon yang dipilih
	fmt.Scan(&pilihan)                               // Menerima input nomor calon

	// Validasi apakah pilihan nomor calon berada dalam rentang yang valid
	if pilihan < 1 || pilihan > jumlahCalon {
		fmt.Println("Nomor calon tidak valid.") // Jika nomor calon tidak valid, tampilkan pesan
		return
	}

	// Update suara dan status pemilih
	daftarCalon[pilihan-1].Suara++                  // Menambah suara pada calon yang dipilih (indeks array mulai dari 0)
	daftarPemilih[pemilihIndex].SudahMemilih = true // Menandai pemilih sudah memilih
	fmt.Println("Pemilihan berhasil dilakukan.")    // Menampilkan pesan bahwa pemilihan berhasil
}

// Pengoperasian Data Calon
func tambahDataCalon() {
	if jumlahCalon == NMAX { // Cek apakah jumlah calon sudah mencapai batas maksimum
		fmt.Println("Jumlah calon sudah maksimal.") // Jika sudah mencapai batas, tampilkan pesan
		return
	}
	// Meminta input data calon baru
	fmt.Print("Masukkan nama calon: ")
	fmt.Scan(&daftarCalon[jumlahCalon].namacalon)
	fmt.Print("Masukkan nama partai: ")
	fmt.Scan(&daftarCalon[jumlahCalon].namapartai)
	daftarCalon[jumlahCalon].Suara = 0 // Inisialisasi suara calon dengan 0
	jumlahCalon++
	fmt.Println("Calon berhasil ditambahkan.")
}

func editDataCalon() {
	var nama string
	fmt.Print("Masukkan nama calon yang ingin diubah: ")
	fmt.Scan(&nama) // Meminta input nama calon yang akan diedit

	for i := 0; i < jumlahCalon; i++ {
		if strings.EqualFold(daftarCalon[i].namacalon, nama) { // Perbandingan nama calon tanpa memperhatikan kapitalisasi
			// Jika calon ditemukan, minta input untuk mengubah nama dan partai
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
		if strings.EqualFold(daftarCalon[i].namacalon, nama) { // Perbandingan nama calon tanpa memperhatikan kapitalisasi
			// Menampilkan konfirmasi sebelum menghapus calon
			fmt.Print("Calon dengan nama " + nama + " dan partai " + daftarCalon[i].namapartai + " akan dihapus. Yakin? (y/n)")
			var konfirmasi string
			fmt.Scan(&konfirmasi)  // Menerima input konfirmasi
			if konfirmasi != "y" { // Jika pengguna tidak mengonfirmasi dengan 'y'
				fmt.Println("Calon tidak jadi dihapus.")
				return
			}
			for j := i; j < jumlahCalon-1; j++ {
				daftarCalon[j] = daftarCalon[j+1] // Menggeser calon setelah calon yang dihapus

			}
			jumlahCalon--
			fmt.Println("Calon berhasil dihapus.")
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

// Fungsi untuk mengatur waktu pemilihan
func setWaktu() {
	fmt.Print("Masukkan waktu sekarang (jam): ")
	fmt.Scan(&waktuSekarang) // Meminta input jam saat ini untuk mengatur waktu pemilihan
	fmt.Printf("Waktu berhasil diset ke jam %d\n", waktuSekarang)
}

func pengoperasianDataCalon() { // Fungsi untuk mengoperasikan data calon (menambah, mengedit, menghapus)
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
	fmt.Print("Masukkan waktu pemilih (jam): ")
	fmt.Scan(&daftarPemilih[jumlahPemilih].waktu)
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
			fmt.Print("Masukkan waktu pemilih baru (jam): ")
			fmt.Scan(&daftarPemilih[i].waktu)
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
	fmt.Printf("Threshold diset ke %d\n", threshold)
}

// Fungsi untuk menampilkan calon yang memenuhi threshold
func tampilkanCalonLolosThreshold() {
	// jika threshold belum diatur tampilkan pesan
	if threshold <= 0 {
		fmt.Println("Threshold belum diatur. Silakan atur threshold terlebih dahulu.")
		return
	}

	fmt.Printf("\nCalon yang memenuhi threshold (%d suara):\n", threshold)
	adaCalon := false
	// menampilkan calon lolos Threshold
	for i := 0; i < jumlahCalon; i++ {
		if daftarCalon[i].Suara >= threshold {
			fmt.Printf("- %s (%s) dengan %d suara\n", daftarCalon[i].namacalon, daftarCalon[i].namapartai, daftarCalon[i].Suara)
			adaCalon = true
		}
	}
	// jika tidak ada calon yang lolos threshold
	if !adaCalon {
		fmt.Println("Tidak ada calon yang memenuhi threshold.")
	}
}
// data awal calon
func dummycalon() {
	daftarCalon[0] = Calon{namacalon: "john", namapartai: "merah", Suara: 120}
	daftarCalon[1] = Calon{namacalon: "smith", namapartai: "kuning", Suara: 90}
	daftarCalon[2] = Calon{namacalon: "robert", namapartai: "hijau", Suara: 75}
	daftarCalon[3] = Calon{namacalon: "emily", namapartai: "biru", Suara: 50}
	daftarCalon[4] = Calon{namacalon: "michael", namapartai: "hitam", Suara: 0}
	jumlahCalon = 5
}
// data awal pemilih
func dummypemilih() {
	daftarPemilih[0] = Pemilih{namapemilih: "alice", SudahMemilih: false, waktu: 1}
	daftarPemilih[1] = Pemilih{namapemilih: "bob", SudahMemilih: false, waktu: 2}
	daftarPemilih[2] = Pemilih{namapemilih: "kevin", SudahMemilih: false, waktu: 3}
	daftarPemilih[3] = Pemilih{namapemilih: "adam", SudahMemilih: false, waktu: 4}
	daftarPemilih[4] = Pemilih{namapemilih: "eve", SudahMemilih: false, waktu: 5}
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
