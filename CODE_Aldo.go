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

func selectionSortAsc() {
	var i, j, idxMin int
	var temp Jadwal

	for i = 0; i < jumlahData-1; i++ {

		idxMin = i

		for j = i + 1; j < jumlahData; j++ {

			if dataJadwal[j].jamMulai < dataJadwal[idxMin].jamMulai {
				idxMin = j
			}
		}

		temp = dataJadwal[i]
		dataJadwal[i] = dataJadwal[idxMin]
		dataJadwal[idxMin] = temp
	}
}

func selectionSortDesc() {
	var i, j, idxMax int
	var temp Jadwal

	for i = 0; i < jumlahData-1; i++ {

		idxMax = i

		for j = i + 1; j < jumlahData; j++ {

			if dataJadwal[j].jamMulai > dataJadwal[idxMax].jamMulai {
				idxMax = j
			}
		}

		temp = dataJadwal[i]
		dataJadwal[i] = dataJadwal[idxMax]
		dataJadwal[idxMax] = temp
	}
}

func insertionSortNamaAsc() {
	var i, j int
	var temp Jadwal

	for i = 1; i < jumlahData; i++ {

		temp = dataJadwal[i]
		j = i - 1

		for j >= 0 && dataJadwal[j].namaMK > temp.namaMK {

			dataJadwal[j+1] = dataJadwal[j]
			j--
		}

		dataJadwal[j+1] = temp
	}
}

func insertionSortNamaDesc() {
	var i, j int
	var temp Jadwal

	for i = 1; i < jumlahData; i++ {

		temp = dataJadwal[i]
		j = i - 1

		for j >= 0 && dataJadwal[j].namaMK < temp.namaMK {

			dataJadwal[j+1] = dataJadwal[j]
			j--
		}

		dataJadwal[j+1] = temp
	}
}

func menuSorting() {
	var pilih int

	fmt.Println("1. Selection Sort Asc")
	fmt.Println("2. Selection Sort Desc")
	fmt.Println("3. Insertion Sort Asc")
	fmt.Println("4. Insertion Sort Desc")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortAsc()
	} else if pilih == 2 {
		selectionSortDesc()
	} else if pilih == 3 {
		insertionSortNamaAsc()
	} else if pilih == 4 {
		insertionSortNamaDesc()
	}
}
