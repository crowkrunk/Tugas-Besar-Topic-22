package main

import "fmt"

func statistik() {
	var i int
	var totalJam int

	for i = 0; i < jumlahData; i++ {

		totalJam += dataJadwal[i].jamSelesai - dataJadwal[i].jamMulai
	}

	fmt.Println("Total Jam Kuliah :", totalJam)
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

func main() {
	var pilih int

	for pilih != 8 {

		fmt.Println("\n===== JADWALKU =====")
		fmt.Println("1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("3. Edit Jadwal")
		fmt.Println("4. Hapus Jadwal")
		fmt.Println("5. Cari Jadwal")
		fmt.Println("6. Sorting Jadwal")
		fmt.Println("7. Statistik")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahJadwal()
		} else if pilih == 2 {
			tampilJadwal()
		} else if pilih == 3 {
			editJadwal()
		} else if pilih == 4 {
			hapusJadwal()
		} else if pilih == 5 {
			menuCari()
		} else if pilih == 6 {
			menuSorting()
		} else if pilih == 7 {
			statistik()
		} else if pilih == 8 {
			fmt.Println("Program selesai")
		}
	}
}
