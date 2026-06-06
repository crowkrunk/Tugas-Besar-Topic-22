package main
import "fmt"

func statistik() {
	var i int
	var totalJam int

	var senin, selasa, rabu, kamis, jumat int

	for i = 0; i < jumlahData; i++ {

		totalJam += dataJadwal[i].jamSelesai - dataJadwal[i].jamMulai

		if dataJadwal[i].hari == "Senin" {
			senin++
		} else if dataJadwal[i].hari == "Selasa" {
			selasa++
		} else if dataJadwal[i].hari == "Rabu" {
			rabu++
		} else if dataJadwal[i].hari == "Kamis" {
			kamis++
		} else if dataJadwal[i].hari == "Jumat" {
			jumat++
		}
	}

	fmt.Println("\n=== Statistik Jadwal ===")
	fmt.Println("Total Jam Kuliah :", totalJam)
	fmt.Println("Jumlah Jadwal Hari Senin  :", senin)
	fmt.Println("Jumlah Jadwal Hari Selasa :", selasa)
	fmt.Println("Jumlah Jadwal Hari Rabu   :", rabu)
	fmt.Println("Jumlah Jadwal Hari Kamis  :", kamis)
	fmt.Println("Jumlah Jadwal Hari Jumat  :", jumat)
}

func hapusJadwal() {
	var kode string
	var idx, i int

	fmt.Print("Masukkan kode MK: ")
	fmt.Scan(&kode)

	idx = sequentialSearchKode(kode)

	if idx != -1 {

		for i = idx; i < jumlahData-1; i++ {
			dataJadwal[i] = dataJadwal[i+1]
		}

		jumlahData--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
