package main

import "fmt"

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
