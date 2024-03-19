package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Perpustakaan struct {
	Kode  string
	Judul string
}

var ListBuku []Perpustakaan
var kodeBuku int

func PerbaruiPerpstakaan() {

	inputanUser := bufio.NewReader(os.Stdin)
	judulBuku := ""
	fmt.Println("====================")
	fmt.Println("	Perpustakaan	")
	fmt.Println("=====================")

	isiPerpustakaan := []Perpustakaan{}

	for {
		fmt.Print("Tekan Enter -> ")

		menuPengguna, err := inputanUser.ReadString('\r')
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		menuPengguna = strings.Replace(
			menuPengguna,
			"\r",
			"",
			1,
		)

		fmt.Print("Silahkan Masukan Kode Buku : ")
		_, err = fmt.Scanln(&kodeBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Print("Masukan Masukan Judul Buku : ")
		_, err = fmt.Scanln(&judulBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		// NewLine
		// _, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		// Simpan Kode dan Judul ( SIMPAN BUKU )
		isiPerpustakaan = append(isiPerpustakaan, Perpustakaan{
			Kode:  strconv.Itoa(kodeBuku),
			Judul: judulBuku,
		})

		var pilihanMenuPengguna = 0
		fmt.Println("Ketik 1 untuk tambah buku, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilihanMenuPengguna)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilihanMenuPengguna == 0 {
			break
		}
	}

	fmt.Println("Data Sedang di Proses... ")

	err := os.MkdirAll("perpustakaan", 0777)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	ch := make(chan Perpustakaan)

	wg := sync.WaitGroup{}

	jumlahStaff := 5

	// Mendaftarkan petugas/staff data
	for i := 0; i < jumlahStaff; i++ {
		wg.Add(1)
		go simpanBuku(ch, &wg, i)
	}

	// Mengrim Data ke Channel
	for _, perpustakaan := range isiPerpustakaan {
		ch <- perpustakaan
	}

	close(ch)

	wg.Wait()

	fmt.Println("Perpustakaan Telah di Perbarui")
}

func simpanBuku(ch <-chan Perpustakaan, wg *sync.WaitGroup, noStaff int) {

	for perpustakaan := range ch {
		dataJson, err := json.Marshal(perpustakaan)
		if err != nil {
			fmt.Println("Terjadi error: ", err)
		}

		err = os.WriteFile(fmt.Sprintf("perpustakaan/%s.json", perpustakaan.Kode), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error: ", err)
		}

		fmt.Printf("Staff No %d Memproses Database Kode : %s!\r", noStaff, perpustakaan.Kode)
	}
	wg.Done()
}

func lihatPerpustakaan(ch <-chan string, chPerpustakaan chan Perpustakaan, wg *sync.WaitGroup) {
	var perpustakaan Perpustakaan
	for kodePerpustakaan := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("perpustakaan/%s", kodePerpustakaan))
		if err != nil {
			fmt.Println("Terjadi error: ", err)
		}

		err = json.Unmarshal(dataJSON, &perpustakaan)
		if err != nil {
			fmt.Println("Terjadi Error: ", err)
		}

		chPerpustakaan <- perpustakaan
	}
	wg.Done()
}

func SeePerpustakaan() {
	fmt.Println("====================")
	fmt.Println("Lihat Data Buku Perpustakaan")
	fmt.Println("====================")
	fmt.Println(" Memuat Perpustakaan... ")
	ListBuku = []Perpustakaan{}

	listJsonPerpustakaan, err := os.ReadDir("perpustakaan")
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chPerpustakaan := make(chan Perpustakaan, len(listJsonPerpustakaan))

	jumlahStaff := 5

	for i := 0; i < jumlahStaff; i++ {
		wg.Add(1)
		go lihatPerpustakaan(ch, chPerpustakaan, &wg)
	}

	for _, databasePerpustakaan := range listJsonPerpustakaan {
		ch <- databasePerpustakaan.Name()
	}

	close(ch)

	wg.Wait()

	close(chPerpustakaan)

	for dataPerpustakaan := range chPerpustakaan {
		ListBuku = append(ListBuku, dataPerpustakaan)
	}

	for urutan, perpustakaan := range ListBuku {
		fmt.Printf("%d. Kode Buku : %s, Judul Buku : %s\n",
			urutan+1,
			perpustakaan.Kode,
			perpustakaan.Judul,
		)
	}
}

func HapusDataPerpustakaan() {
	fmt.Println("====================")
	fmt.Println("Hapus Data Perpustakaan")
	fmt.Println("====================")
	SeePerpustakaan()
	fmt.Println("====================")
	var urutanPerpustakaan int
	fmt.Print("Masukan Urutan Pesanan : ")
	_, err := fmt.Scanln(&urutanPerpustakaan)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
	}

	if (urutanPerpustakaan-1) < 0 ||
		(urutanPerpustakaan-1) > len(ListBuku) {
		fmt.Println("Urutan Perpustakaan Tidak Sesuai")
		HapusDataPerpustakaan()
		return
	}

	err = os.Remove(fmt.Sprintf("pesanan/%s.jsonn", ListBuku[urutanPerpustakaan-1].Kode))
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
	}

	fmt.Println("Data Perpustakaan Berhasil di Hapus")

}

func main() {
	var pilihanMenu int
	fmt.Println("==============================")
	fmt.Println(" Software Perpustakaan Bamboo ")
	fmt.Println("==============================")
	fmt.Println("Silahkan Memilih : ")
	fmt.Println("1. Tambah Data Buku")
	fmt.Println("2. Lihat Data Buku")
	fmt.Println("3. Hapus Data Buku")
	fmt.Println("4. Keluar")
	fmt.Println("==============================")
	fmt.Print("Masukan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error : ", err)
	}

	switch pilihanMenu {
	case 1:
		PerbaruiPerpstakaan()
	case 2:
		SeePerpustakaan()
	case 3:
		HapusDataPerpustakaan()
	case 4:
		os.Exit(0)
	}

	main()
}
