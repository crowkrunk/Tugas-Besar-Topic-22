package main

import "fmt"

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
