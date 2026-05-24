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
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	}
	fmt.Println("Program selesai")

}
