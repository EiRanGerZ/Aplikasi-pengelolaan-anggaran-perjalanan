package main
import "fmt"

const NMAX = 1000

type Pengeluaran struct {
	kategori  [NMAX]string
	jumlah    [NMAX]float64
	nKategori int
}

func main() {
	var data Pengeluaran
	var anggaran float64
	var pilihan, idx int
	var nama, metode string

	fmt.Print("Masukkan anggaran perjalanan Anda (Rp): ")
	fmt.Scan(&anggaran)

	for {
		menuAwal()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahUpdate(&data)
		case 2:
			fmt.Print("Nama kategori yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusKategori(&data, nama)
		case 3:
			tampilkanPengeluaran(data)
		case 4:
			fmt.Print("Nama kategori yang dicari: ")
			fmt.Scan(&nama)
			fmt.Print("Metode (sequential/binary): ")
			fmt.Scan(&metode)
			if metode == "sequential" {
				idx = sequentialSearch(data, nama)
			} else if metode == "binary" {
				// urutkan array asli SEBELUM cari
				selectionSortKategori(&data)
				idx = binarySearch(data, nama)
			} else {
				fmt.Println("Metode tidak dikenal, default ke sequential")
				idx = sequentialSearch(data, nama)
			} 
			if idx != -1 {
				fmt.Println("\n+--------------------------------------------------+")
				fmt.Printf("Kategori ditemukan: %s Rp%.2f", data.kategori[idx], data.jumlah[idx])
				fmt.Println("\n+--------------------------------------------------+")
			} else {
				fmt.Println("\n+--------------------------------------------------+")
				fmt.Println("|             Kategori tidak ditemukan             |")
				fmt.Println("+--------------------------------------------------+")
			}
		case 5:
			selectionSortKategori(&data)
			fmt.Println("\n+--------------------------------------------------+")
			fmt.Println("|             Data berhasil diurutkan              |")
			fmt.Println("+--------------------------------------------------+")
		case 6:
			selectionSortJumlah(&data)
			fmt.Println("\n+--------------------------------------------------+")
			fmt.Println("|             Data berhasil diurutkan              |")
			fmt.Println("+--------------------------------------------------+")
		case 7:
			insertionSort(&data)
			fmt.Println("\n+--------------------------------------------------+")
			fmt.Println("|              Data berhasil diurutkan             |")
			fmt.Println("+--------------------------------------------------+")
		case 8:
			tampilkanLaporan(data, anggaran)

		case 9:
			fmt.Println("\n+--------------------------------------------------+")
			fmt.Println("|     Terima kasih telah menggunakan aplikasi      |")
			fmt.Println("+--------------------------------------------------+")
			return

		default:
			fmt.Println("\n+--------------------------------------------------+")
			fmt.Println("|                Pilihan tidak valid               |")
			fmt.Println("+--------------------------------------------------+")
		}
	}
}

func menuAwal() {
	fmt.Println("\n+--------------------------------------------------+")
	fmt.Println("|      --- Menu Aplikasi Budget Traveling ---      |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("| 1. Tambah / Update Pengeluaran                   |")
	fmt.Println("| 2. Hapus Pengeluaran                             |")
	fmt.Println("| 3. Tampilkan Semua Pengeluaran                   |")
	fmt.Println("| 4. Cari Kategori (Sequential/Binary Search)      |")
	fmt.Println("| 5. Urutkan Data (Nama Kategori)                  |")
	fmt.Println("| 6. Urutkan Data (Pengeluaran Terendah)           |")
	fmt.Println("| 7. Urutkan Data (Pengeluaran Tertinggi)          |")
	fmt.Println("| 8. Laporan Anggaran & Saran                      |")
	fmt.Println("| 9. Keluar                                        |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Print("Pilih menu: ")
}

func tambahUpdate(p *Pengeluaran) {
	var kategoriDefault [4]string
	var tambah, namaBaru string
	var nilaiBaru, jumlah float64
	var i int
	
	kategoriDefault[0] = "Transportasi"
	kategoriDefault[1] = "Akomodasi"
	kategoriDefault[2] = "Makanan"
	kategoriDefault[3] = "Hiburan"

	for i = 0; i < 4; i++ {
		fmt.Printf("Masukkan jumlah pengeluaran untuk kategori %s: ", kategoriDefault[i])
		fmt.Scan(&jumlah)
		tambahAtauUpdateKategori(p, kategoriDefault[i], jumlah)
	}

	// Input tambahan kategori lain jika user ingin
	for {
		fmt.Print("Apakah ingin menambahkan kategori lain? (ya/tidak): ")
		fmt.Scan(&tambah)
		if tambah != "ya" {
			break
		}
		fmt.Print("Masukkan nama kategori: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Masukkan jumlah pengeluaran: ")
		fmt.Scan(&nilaiBaru)
		tambahAtauUpdateKategori(p, namaBaru, nilaiBaru)
	}
}

// Fungsi dasar
func tambahAtauUpdateKategori(p *Pengeluaran, nama string, nilai float64) {
	var i int
	
	for i = 0; i < p.nKategori; i++ {
		if p.kategori[i] == nama {
			p.jumlah[i] += nilai
			return
		}
	}
	p.kategori[p.nKategori] = nama
	p.jumlah[p.nKategori] = nilai
	p.nKategori++
}

func hapusKategori(p *Pengeluaran, nama string) {
	var i, j int
	
	for i = 0; i < p.nKategori; i++ {
		if p.kategori[i] == nama {
			for j = i; j < p.nKategori-1; j++ {
				p.kategori[j] = p.kategori[j+1]
				p.jumlah[j] = p.jumlah[j+1]
			}
			p.nKategori--
			fmt.Println("+--------------------------------------------------+")
			fmt.Println("|            Kategori berhasil dihapus             |")
			fmt.Println("+--------------------------------------------------+")
			return
		}
	}
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("|            Kategori tidak ditemukan              |")
	fmt.Println("+--------------------------------------------------+")
}

func tampilkanPengeluaran(p Pengeluaran) float64 {
	var total float64
	var i int
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("|                 Daftar Pengeluaran               |")
	fmt.Println("+--------------------------------------------------+")
	for i = 0; i < p.nKategori; i++ {
		fmt.Printf("%d. %s: Rp%.2f\n", i+1, p.kategori[i], p.jumlah[i])
		total += p.jumlah[i]
	}
	fmt.Println("+--------------------------------------------------+")
	fmt.Printf("Total Pengeluaran: Rp%.2f\n", total)
	fmt.Println("+--------------------------------------------------+")
	fmt.Println()
	return total
}

// Search
func sequentialSearch(p Pengeluaran, nama string) int {
	var i int
	for i = 0; i < p.nKategori; i++ {
		if p.kategori[i] == nama {
			return i
		}
	}
	return -1
}

func binarySearch(p Pengeluaran, nama string) int {
	selectionSortKategori(&p)
	var left, right, mid int
	
	left = 0
	right = p.nKategori - 1

	for left <= right {
		mid = (left + right) / 2
		if p.kategori[mid] == nama {
			return mid
		} else if p.kategori[mid] < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Sorting
//ASCENDING
func selectionSortKategori(p *Pengeluaran) {
	var idx, i, pass int
	var tempKategori string
	var tempJumlah float64
	pass = 1

	for pass <= p.nKategori {
		idx = pass - 1
		i = pass
		for i < p.nKategori {
			if p.kategori[i] < p.kategori[idx] {
				idx = i
			}
			i = i + 1
		}

		tempKategori = p.kategori[pass-1]
		p.kategori[pass-1] = p.kategori[idx]
		p.kategori[idx] = tempKategori

		tempJumlah = p.jumlah[pass-1]
		p.jumlah[pass-1] = p.jumlah[idx]
		p.jumlah[idx] = tempJumlah

		pass = pass + 1
	}
}
func selectionSortJumlah(p *Pengeluaran) {
	var idx, i, pass int
	var tempKategori string
	var tempJumlah float64
	
	pass = 1

	for pass <= p.nKategori {
		idx = pass - 1
		i = pass
		for i < p.nKategori {
			if p.jumlah[i] < p.jumlah[idx] {
				idx = i
			}
			i = i + 1
		}

		tempKategori = p.kategori[pass-1]
		p.kategori[pass-1] = p.kategori[idx]
		p.kategori[idx] = tempKategori

		tempJumlah = p.jumlah[pass-1]
		p.jumlah[pass-1] = p.jumlah[idx]
		p.jumlah[idx] = tempJumlah

		pass = pass + 1
	}
}

//DESCENDING
func insertionSort(p *Pengeluaran) {
	var i, pass int
	var tempKategori string
	var tempJumlah float64

	pass = 1
	for pass < p.nKategori {
		i = pass
		tempJumlah = p.jumlah[pass]
		tempKategori = p.kategori[pass]
		for i > 0 && tempJumlah > p.jumlah[i-1] {
			p.jumlah[i] = p.jumlah[i-1]
			p.kategori[i] = p.kategori[i-1]
			i = i - 1
		}
		p.jumlah[i] = tempJumlah
		p.kategori[i] = tempKategori
		pass = pass + 1
	}
}

// Laporan dan saran
func tampilkanLaporan(p Pengeluaran, anggaran float64) {
	var total, selisih float64
	var i int
	
	total = tampilkanPengeluaran(p)
	selisih = anggaran - total
	fmt.Println("+==================================================+")
	fmt.Println("|                 LAPORAN ANGGARAN                 |")
	fmt.Println("+==================================================+")
	fmt.Println("|                Daftar Pengeluaran                |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Printf("Anggaran: Rp%.2f\n", anggaran)
	fmt.Printf("Selisih: Rp%.2f\n", selisih)
	if selisih < 0 {
		fmt.Println("+--------------------------------------------------------------------------------+")
		fmt.Println("⚠️ Pengeluaran melebihi anggaran! Pertimbangkan mengurangi kategori besar seperti:")
		selectionSortJumlah(&p)
		fmt.Println("+--------------------------------------------------------------------------------+")
		for i = p.nKategori - 1; i >= 0; i-- {
			fmt.Println("+--------------------------------------------------+")
			fmt.Printf("- %s: Rp%.2f\n", p.kategori[i], p.jumlah[i])
			fmt.Println("+--------------------------------------------------+")
		}
	} else {
		fmt.Println("+--------------------------------------------------+")
		fmt.Println("     ✅ Pengeluaran masih dalam batas anggaran")
		fmt.Println("+--------------------------------------------------+")
	}
}
