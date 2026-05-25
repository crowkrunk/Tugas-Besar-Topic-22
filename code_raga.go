package main
import "fmt"

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

func cekBentrok(j Jadwal) bool {
	var i int

	for i = 0; i < jumlahData; i++ {

		if dataJadwal[i].hari == j.hari &&
			dataJadwal[i].ruangan == j.ruangan {

			if j.jamMulai < dataJadwal[i].jamSelesai &&
				j.jamSelesai > dataJadwal[i].jamMulai {

				return true
			}
		}
	}

	return false
}
