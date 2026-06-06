package main
import "fmt"

const NMAX = 100

type Jadwal struct {
	kodeMK     string
	namaMK     string
	dosen      string
	hari       string
	jamMulai   int
	jamSelesai int
	ruangan    string
}

var dataJadwal [NMAX]Jadwal
var jumlahData int

func tambahJadwal() {
	if jumlahData >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	var j Jadwal

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
		fmt.Println("Jam tidak valid")
		return
	}

	if cekBentrok(j) {
		fmt.Println("Jadwal bentrok")
		return
	}

	dataJadwal[jumlahData] = j
	jumlahData++

	fmt.Println("Data berhasil ditambah")
}

func tampilJadwal() {
	var i int

	if jumlahData == 0 {
		fmt.Println("Belum ada data")
		return
	}

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

func sequentialSearchKode(kode string) int {
	var i int

	for i = 0; i < jumlahData; i++ {

		if dataJadwal[i].kodeMK == kode {
			return i
		}
	}

	return -1
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

	fmt.Println("Data berhasil diurutkan ascending")
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

	fmt.Println("Data berhasil diurutkan descending")
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

	fmt.Println("Data berhasil diurutkan nama ascending")
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

	fmt.Println("Data berhasil diurutkan nama descending")
}

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
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}