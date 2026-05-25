package main

import "fmt"

func main() {

	var pilih int

	for pilih != 8 {

		fmt.Println("\n1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("3. Edit Jadwal")
		fmt.Println("4. Hapus Jadwal")
		fmt.Println("5. Cari Jadwal")
		fmt.Println("6. Sorting Jadwal")
		fmt.Println("7. Statistik")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)
	}
	switch pilih {
	case 1:
		tambahJadwal()
	case 2:
		tampilJadwal()
	case 3:
		editJadwal()
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	}
	fmt.Println("Program selesai")

}

type Jadwal struct {
	kodeMK     string
	namaMK     string
	dosen      string
	hari       string
	jamMulai   int
	jamSelesai int
	ruangan    string
}

var dataJadwal [100]Jadwal
var jumlahData int

func tambahJadwal() {
	if jumlahData >= 100 {
		fmt.Println("Data penuh")
		return
	}

	var j Jadwal

	fmt.Println("\n=== Tambah Jadwal ===")

	fmt.Print("Kode MK        : ")
	fmt.Scan(&j.kodeMK)

	fmt.Print("Nama MK        : ")
	fmt.Scan(&j.namaMK)

	fmt.Print("Dosen          : ")
	fmt.Scan(&j.dosen)

	fmt.Print("Hari           : ")
	fmt.Scan(&j.hari)

	fmt.Print("Jam Mulai      : ")
	fmt.Scan(&j.jamMulai)

	fmt.Print("Jam Selesai    : ")
	fmt.Scan(&j.jamSelesai)

	fmt.Print("Ruangan        : ")
	fmt.Scan(&j.ruangan)

	if j.jamSelesai <= j.jamMulai {
		fmt.Println("Jam selesai harus lebih besar")
		return
	}

	if cekBentrok(j) {
		fmt.Println("Jadwal bentrok!")
		return
	}

	dataJadwal[jumlahData] = j
	jumlahData++

	fmt.Println("Jadwal berhasil ditambahkan")
}

func cekBentrok(j Jadwal) bool {
	var i int

	for i = 0; i < jumlahData; i++ {

		if dataJadwal[i].hari == j.hari {

			if dataJadwal[i].ruangan == j.ruangan {

				if j.jamMulai < dataJadwal[i].jamSelesai &&
					j.jamSelesai > dataJadwal[i].jamMulai {

					return true
				}
			}
		}
	}

	return false
}

func tampilJadwal() {
	var i int

	for i = 0; i < jumlahData; i++ {

		fmt.Println("----------------------")
		fmt.Println("Kode MK     :", dataJadwal[i].kodeMK)
		fmt.Println("Nama MK     :", dataJadwal[i].namaMK)
		fmt.Println("Dosen       :", dataJadwal[i].dosen)
		fmt.Println("Hari        :", dataJadwal[i].hari)
		fmt.Println("Jam Mulai   :", dataJadwal[i].jamMulai)
		fmt.Println("Jam Selesai :", dataJadwal[i].jamSelesai)
		fmt.Println("Ruangan     :", dataJadwal[i].ruangan)
	}
}

func editJadwal() {
	var kode string
	var idx int

	fmt.Print("Masukkan kode MK: ")
	fmt.Scan(&kode)

	idx = sequentialSearchKode(kode)

	if idx != -1 {

		fmt.Print("Nama MK Baru : ")
		fmt.Scan(&dataJadwal[idx].namaMK)

		fmt.Print("Dosen Baru : ")
		fmt.Scan(&dataJadwal[idx].dosen)

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
