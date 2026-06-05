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

func sequentialSearchNama(nama string) int {
	var i int

	for i = 0; i < jumlahData; i++ {

		if dataJadwal[i].namaMK == nama {
			return i
		}
	}

	return -1
}

func binarySearchKode(kode string) int {
	var left, right, mid int

	left = 0
	right = jumlahData - 1

	for left <= right {

		mid = (left + right) / 2

		if dataJadwal[mid].kodeMK == kode {
			return mid
		}

		if dataJadwal[mid].kodeMK < kode {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}


func menuCari() {
	var pilih int
	var key string
	var idx int

	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		fmt.Print("Nama MK : ")
		fmt.Scan(&key)

		idx = sequentialSearchNama(key)

		if idx != -1 {
			fmt.Println("Data ditemukan")
			fmt.Println(dataJadwal[idx])
		} else {
			fmt.Println("Data tidak ditemukan")
		}

	} else if pilih == 2 {

		fmt.Print("Kode MK : ")
		fmt.Scan(&key)

		idx = binarySearchKode(key)

		if idx != -1 {
			fmt.Println("Data ditemukan")
			fmt.Println(dataJadwal[idx])
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

