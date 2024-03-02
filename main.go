package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Buku struct {
	KodeBuku      string
	JudulBuku     string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var DaftarBuku []Buku

// FUNGSI TEST
// FUNGSI TEST
// FUNGSI TEST
// FUNGSI TEST
// BUKAN KODE PRIBADI
// BUKAN KODE PRIBADI
// BUKAN KODE PRIBADI
// BUKAN KODE PRIBADI

// Untuk menulis data buku ke file external
func SimpanDataBuku() {
	// Membuka file dengan mode write dan create
	file, err := os.OpenFile("data_buku.love", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	// Mengubah data buku menjadi format JSON
	dataJSON, err := json.Marshal(DaftarBuku)
	if err != nil {
		fmt.Println("Gagal mengubah data buku menjadi JSON:", err)
		return
	}

	// Menulis data JSON ke file
	_, err = file.Write(dataJSON)
	if err != nil {
		fmt.Println("Gagal menulis data JSON ke file:", err)
		return
	}
	fmt.Println("==========================================================================")
	fmt.Println("Data buku berhasil disimpan ke file data_buku.love")
	fmt.Println("==========================================================================")
}

// Untuk membaca data buku dari data_buku.love
func BacaDataBuku() {
	// Membuka file dengan mode read only
	file, err := os.OpenFile("data_buku.love", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	// Membaca data JSON dari file
	dataJSON := make([]byte, 1024)
	n, err := file.Read(dataJSON)
	if err != nil {
		fmt.Println("Gagal membaca data JSON dari file:", err)
		return
	}

	// Mengubah data JSON menjadi data buku
	err = json.Unmarshal(dataJSON[:n], &DaftarBuku)
	if err != nil {
		fmt.Println("Gagal mengubah data JSON menjadi data buku:", err)
		return
	}
	fmt.Println("==========================================================================")
	fmt.Println("Data buku berhasil dibaca dari file data_buku.love")
	fmt.Println("==========================================================================")
}

// FUNGSI TEST END
// FUNGSI TEST END
// FUNGSI TEST END
// FUNGSI TEST END

func AddBuku(kode, judul, pengarang, penerbit string, halaman, tahun int) {
	buku := Buku{
		KodeBuku:      kode,
		JudulBuku:     judul,
		Pengarang:     pengarang,
		Penerbit:      penerbit,
		JumlahHalaman: halaman,
		TahunTerbit:   tahun,
	}
	DaftarBuku = append(DaftarBuku, buku)
	fmt.Println("Ditambahkan buku baru dengan kode ", kode)
}

// Tamilan untuk Buku yang Terdaftar
// Tamilan untuk Buku yang Terdaftar
func TampilanDaftar() {
	fmt.Println("==========================================================================")
	fmt.Println("Daftar Buku Perpustakaan : ")
	fmt.Println("==========================================================================")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("==========================================================================")
	fmt.Println(" Kode  |   Judul Buku   |  Pengarang  |  Penerbit  | Halaman | Tahun Terbit")
	for _, buku := range DaftarBuku {
		fmt.Println(fmt.Sprintf(" %-6s| %-15s| %-12s| %-11s| %-8d| %-8d", buku.KodeBuku, strings.ReplaceAll(buku.JudulBuku, " ", "+"), buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit))
		fmt.Println("==========================================================================")
		fmt.Println("==========================================================================")
	}

	//delay
	fmt.Println(" Scroll ke Atas untuk lihat list kembali, Atau Di Menu Utama Tekan ' 2 '")
	fmt.Println("Tunggu 5 detik untuk melanjutkan .....................................")
	time.Sleep(5 * time.Second)

	// fmt.Println("	Apakah Anda Ingin Kembali Ke Menu? ")
	// fmt.Print("	[Y] Yes / [N] No	: ")
	// var balikmenu string
	// fmt.Scanln(&balikmenu)
	// if balikmenu == "Y" || balikmenu == "y" {
	// 	main()
	// } else if balikmenu == "N" || balikmenu == "n" {
	// 	fmt.Println("==========================================================================")
	// 	fmt.Println("==========================================================================")
	// 	fmt.Printf("Terimakasih Telah Menggunakan Aplikasi Go")
	// 	os.Exit(0)
	// }
}

// fungsi HapusBuku
func HapusBuku(kode string) {
	for i, buku := range DaftarBuku {
		if buku.KodeBuku == kode {
			DaftarBuku = append(DaftarBuku[:i], DaftarBuku[i+1:]...)
			fmt.Println("Buku Berhasil di Hapus")
			return
		}
	}
	fmt.Println("================================================")
	fmt.Println("Buku tidak ada / tidak di temukan")
	fmt.Println("================================================")
	time.Sleep(2 * time.Second)
}

// EditBuku  digunakan untuk mengedit data buku berdasarkaan kode buku
func EditBuku(kode, judul, pengarang, penerbit string, halaman, tahun int) {
	for i, buku := range DaftarBuku {
		if buku.KodeBuku == kode {
			DaftarBuku[i].JudulBuku = judul
			DaftarBuku[i].Pengarang = pengarang
			DaftarBuku[i].Penerbit = penerbit
			DaftarBuku[i].JumlahHalaman = halaman
			DaftarBuku[i].TahunTerbit = tahun
			fmt.Println("Data Buku berhasil dirubah.")
			return
		}
	}
	fmt.Println("================================================")
	fmt.Println("Buku tidak ada / tidak di temukan")
	fmt.Println("================================================")
	time.Sleep(2 * time.Second)
}

// Tampilan
// function
func TambahBuku() {
	var kode, judul, pengarang, penerbit string
	var halaman, tahun int
	fmt.Println("============================================")
	fmt.Println("Silahkan Masukan Informasi Buku : ")
	fmt.Print("     1. Kode Buku		: ")
	fmt.Scanln(&kode)
	fmt.Print("     2. Judul Buku		: ")
	fmt.Scanln(&judul)
	fmt.Print("     3. Pengarang Buku		: ")
	fmt.Scanln(&pengarang)
	fmt.Print("     4. Penerbit Buku		: ")
	fmt.Scanln(&penerbit)
	fmt.Print("     5. Jumlah Halaman Buku 	: ")
	fmt.Scanln(&halaman)
	fmt.Print("     6. Tahun Terbit Buku 	: ")
	fmt.Scanln(&tahun)
	AddBuku(kode, judul, pengarang, penerbit, halaman, tahun)
	SimpanDataBuku()
}

func RemoveBuku() {
	fmt.Println("==========================================================================")
	TampilanDaftar()
	fmt.Println("==========================================================================")
	var kode string
	fmt.Print("Masukan kode buku yang ingin anda  hapus : ")
	fmt.Scanln(&kode)
	HapusBuku(kode)
}

// Edit Buku
func EditsBookoo() {
	TampilanDaftar()
	fmt.Print("Masukkan kode buku yang ingin Anda edit: ")
	var kode string
	fmt.Scanln(&kode)
	for i, buku := range DaftarBuku {
		if buku.KodeBuku == kode {
			fmt.Println("=======================================================")
			fmt.Println("Buku dengan kode ", kode, " ditemukan.")
			fmt.Println("=======================================================")
			fmt.Println("Pilih bagian mana yang ingin Anda edit:")
			fmt.Println("1. Kode Buku")
			fmt.Println("2. Judul Buku")
			fmt.Println("3. Pengarang Buku")
			fmt.Println("4. Penerbit Buku")
			fmt.Println("5. Halaman Buku")
			fmt.Println("6. Tahun Terbit")
			fmt.Println("7. Kembali ke menu")
			fmt.Print("Masukkan pilihan Anda: ")
			var pilihan int
			fmt.Scanln(&pilihan)
			fmt.Println("=======================================================")

			switch pilihan {
			case 1:
				fmt.Print("Masukkan kode buku yang baru: ")
				var kodeBaru string
				fmt.Scanln(&kodeBaru)
				DaftarBuku[i].KodeBuku = kodeBaru
				fmt.Println("Kode buku berhasil diubah.")
			case 2:
				fmt.Print("Masukkan judul buku yang baru: ")
				var judulBaru string
				fmt.Scanln(&judulBaru)
				DaftarBuku[i].JudulBuku = judulBaru
				fmt.Println("Judul buku berhasil diubah.")
			case 3:
				fmt.Print("Masukkan pengarang buku yang baru: ")
				var pengarangBaru string
				fmt.Scanln(&pengarangBaru)
				DaftarBuku[i].Pengarang = pengarangBaru
				fmt.Println("Pengarang buku berhasil diubah.")
			case 4:
				fmt.Print("Masukkan penerbit buku yang baru: ")
				var penerbitBaru string
				fmt.Scanln(&penerbitBaru)
				DaftarBuku[i].Penerbit = penerbitBaru
				fmt.Println("Penerbit buku berhasil diubah.")
			case 5:
				fmt.Print("Masukkan halaman buku yang baru: ")
				var halamanBaru int
				fmt.Scanln(&halamanBaru)
				DaftarBuku[i].JumlahHalaman = halamanBaru
				fmt.Println("Halaman buku berhasil diubah.")
			case 6:
				fmt.Print("Masukkan tahun terbit buku yang baru: ")
				var tahunBaru int
				fmt.Scanln(&tahunBaru)
				DaftarBuku[i].TahunTerbit = tahunBaru
				fmt.Println("Tahun terbit buku berhasil diubah.")
			case 7:
				fmt.Println("Anda kembali ke menu.")
				main()
			default:
				fmt.Println("Pilihan Anda tidak valid. Silakan masukkan angka dari 1 sampai 7.")
			}
			return
		}
	}
	fmt.Println("================================================")
	fmt.Println("Buku tidak ada / tidak di temukan")
	fmt.Println("================================================")
	time.Sleep(2 * time.Second)
}

// Func Main
func main() {
	//Baca Data Buku Dulu
	fmt.Println("=======================================================")
	fmt.Println("Mencoba Membaca Data Buku")
	fmt.Println("-- --- -- --- -- --- -- --- -- --- -- --- -- --- -- ---")
	BacaDataBuku()

	fmt.Println("=======================================================")
	fmt.Println("    |  Aplikasi Go - Daftar Buku Perpustakaan  |   ")
	fmt.Println("=======================================================")
	fmt.Println("List Menu:")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Edit / Ubah Buku")
	fmt.Println("5. Simpan Data Buku")
	fmt.Println("6. Baca Data Buku")
	fmt.Println("7. Keluar / Exit")
	fmt.Println("=======================================================")
	fmt.Println("=======================================================")
	fmt.Print("Masukan Pilihan Anda Di Sini :  ")

	//Input
	var pilihan int
	fmt.Scanln(&pilihan)

	//Case
	switch pilihan {
	case 1:
		//Tambah Buku
		TambahBuku()
		main()

	case 2:
		// Memperlihatkan list buku tadi
		TampilanDaftar()
		main()

	case 3:
		//Hapus Buku dari kode
		RemoveBuku()
		main()

	case 4:
		//Edit Buku dari kode
		EditsBookoo()
		main()

	case 5:
		SimpanDataBuku()
		main()

	case 6:
		BacaDataBuku()
		main()

	case 7:
		// Keluar / Exit
		fmt.Println("======================================================================================================")
		fmt.Println("Terimakasih Telah menggunakan Aplikasi ini. Tolong donasinya untuk membantu Developer ASEKK wkwk")
		fmt.Println("======================================================================================================")
		os.Exit(0)

	default:
		// Nuansa Error Klo buku ga ada
		fmt.Println("ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ")
		fmt.Println("Windows anda terkena Virus")
		fmt.Println("ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ")
		fmt.Println("==================================================")
		fmt.Println("Bercanda.. Cie Panik.. -- Pilihan anda Tidak Valid --")
		fmt.Println("==================================================")
		main()
	}
}
