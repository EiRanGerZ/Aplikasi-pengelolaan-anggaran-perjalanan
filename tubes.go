package main

import "fmt"

const maxData = 1000
type Pengeluaran struct{
	transportasi, akomodasi, makanan, hiburan float64
	idxTransportasi, idxAkomodasi, idxMakanan, idxHiburan int
}
type arrPengeluaran [maxData]Pengeluaran


func main() {
	var data arrPengeluaran
	var totalBudget float64
	
    fmt.Print("Masukkan total budget perjalanan: ")
    fmt.Scan(&totalBudget)
	menuAwal(&data)
	
	
}

func tambahPengeluaran(A *arrPengeluaran) {
	var pilihkategori int
	var tambah float64
	for pilihkategori != 0{
        fmt.Println("\nPilih Kategori:")
        fmt.Println("1. Transportasi")
        fmt.Println("2. Akomodasi")
        fmt.Println("3. Makanan")
        fmt.Println("4. Hiburan")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih kategori: ")
        fmt.Scan(&pilihkategori)
		
		switch pilihkategori {
			case 0:
				fmt.Println("Kembali ke menu utama")
			case 1:
				fmt.Print("Masukan budget Transportasi: ")
				fmt.Scan(&tambah)
				A.transportasi[A.idxTransportasi] = tambah
				A.idxTransportasi++
				
			case 2:
				fmt.Print("Masukan budget Akomodasi: ")
				fmt.Scan(&tambah)
				A.akomodasi[A.idxAkomodasi] = tambah
				A.idxAkomodasi++
				
			case 3:
				fmt.Scan(&tambah)
				A.makanan[A.idxMakanan] = tambah
				A.idxMakanan++
				
			case 4:
				fmt.Scan(&tambah)
				A.hiburan[A.idxHiburan] = tambah
				A.idxHiburan++
			default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}

func ubahPengeluaran(A *arrPengeluaran) {
	var pilihkategori, idx int
	var nilaiBaru float64
        fmt.Println("\nUbah Budget Pengeluaran:")
        fmt.Println("1. Transportasi")
        fmt.Println("2. Akomodasi")
        fmt.Println("3. Makanan")
        fmt.Println("4. Hiburan")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih kategori: ")
        fmt.Scan(&pilihkategori)
	switch pilihkategori {
		case 1:
			fmt.Println("Pengeluaran Budget Transportasi:")
			for i := 0; i < A.idxTransportasi; i++ {
				fmt.Print(i+1, ". ", A.transportasi[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxTransportasi {
				fmt.Print("Masukkan budget baru: ")
				fmt.Scan(&nilaiBaru)
				A.transportasi[idx] = nilaiBaru
				fmt.Println("Pengeluaran berhasil diubah")
			} else {
				fmt.Println("Data kosong")
			}
		case 2:
			fmt.Println("Pengeluaran Budget Akomodasi:")
			for i := 0; i < A.idxAkomodasi; i++ {
				fmt.Print(i+1, ". ", A.akomodasi[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin diubah:: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxAkomodasi {
				fmt.Print("Masukkan budget baru: ")
				fmt.Scan(&nilaiBaru)
				A.akomodasi[idx] = nilaiBaru
				fmt.Println("Pengeluaran berhasil diubah")
			} else {
				fmt.Println("Data kosong")
			}
		case 3:
			fmt.Println("Pengeluaran Budget Makanan:")
			for i := 0; i < A.idxMakanan; i++ {
				fmt.Print(i+1, ". ", A.makanan[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxMakanan {
				fmt.Print("Masukkan budget baru: ")
				fmt.Scan(&nilaiBaru)
				A.makanan[idx] = nilaiBaru
				fmt.Println("Pengeluaran berhasil diubah")
			} else {
				fmt.Println("Data kosong")
			}
		case 4:
			fmt.Println("Pengeluaran Budget Hiburan:")
			for i := 0; i < A.idxHiburan; i++ {
				fmt.Print(i+1, ". ", A.hiburan[i])
			}
			fmt.Print("Masukkan nomor yang ingin diubah: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxHiburan {
				fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
				fmt.Scan(&nilaiBaru)
				A.hiburan[idx] = nilaiBaru
				fmt.Println("Pengeluaran berhasil diubah")
			} else {
				fmt.Println("Data kosong")
			}
	}

func hapusPengeluaran(A *arrPengeluaran) {
	var pilihkategori, idx int
	var nilaiBaru float64
        fmt.Println("\Hapus Budget Pengeluaran:")
        fmt.Println("1. Transportasi")
        fmt.Println("2. Akomodasi")
        fmt.Println("3. Makanan")
        fmt.Println("4. Hiburan")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih kategori: ")
        fmt.Scan(&pilihkategori)
	switch pilihkategori {
		case 1:
			fmt.Println("Pengeluaran Budget Transportasi:")
			for i := 0; i < A.idxTransportasi; i++ {
				fmt.Print(i+1, ". ", A.transportasi[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxTransportasi {
				for i = idx; i < A.idxTransportasi-1; i++{
				A.transportasi[i] = A.transportasi[i+1]
				}
				A.idxTransportasi--
				fmt.Println("Pengeluaran berhasil dihapus")
			} else {
				fmt.Println("Data kosong")
			}
		case 2:
			fmt.Println("Pengeluaran Budget Akomodasi:")
			for i := 0; i < A.idxAkomodasi; i++ {
				fmt.Print(i+1, ". ", A.akomodasi[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxAkomodasi {
				for i = idx; i < A.idxAkomodasi-1; i++{
				A.akomodasi[i] = A.akomodasi[i+1]
				}
				A.idxAkomodasi--
				fmt.Println("Pengeluaran berhasil dihapus")
			} else {
				fmt.Println("Data kosong")
			}
		case 3:
			fmt.Println("Pengeluaran Budget Makanan:")
			for i := 0; i < A.idxMakanan; i++ {
				fmt.Print(i+1, ". ", A.makanan[i])
			}
			fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxMakanan {
				for i = idx; i < A.idxMakanan-1; i++{
				A.makanan[i] = A.makanan[i+1]
				}
				A.idxMakanan--
				fmt.Println("Pengeluaran berhasil dihapus")
			} else {
				fmt.Println("Data kosong")
			}
		case 4:
			fmt.Println("Pengeluaran Budget Hiburan:")
			for i := 0; i < A.idxHiburan; i++ {
				fmt.Print(i+1, ". ", A.hiburan[i])
			}
			fmt.Print("Masukkan nomor yang ingin dihapus: ")
			fmt.Scan(&idx)
			idx = idx - 1
			if idx >= 0 && idx < A.idxHiburan {
				for i = idx; i < A.idxHiburan-1; i++{
				A.hiburan[i] = A.hiburan[i+1]
				}
				A.idxHiburan--
				fmt.Println("Pengeluaran berhasil dihapus")
			} else {
				fmt.Println("Data kosong")
			}
	}
}

func hitungTotalPengeluaran() {
	var totalTransportasi, totalAkomodasi, totalMakanan, totalHiburan float64
	var totalKeseluruhan float64
	var i int
	
	for i = 0; i < A.idxTransportasi; i++{
		totalTransportasi += A.transportasi[i]
	}
	for i = 0; i < A.idxAkomodasi; i++{
		totalAkomodasi += A.akomodasi[i]
	}
	for i = 0; i < A.idxMakanan; i++{
		totalMakanan += A.makanan[i]
	}
	for i = 0; i < A.idxHiburan; i++{
		totalTransportasi += A.hiburan[i]
	}
	
	fmt.Println("Total Pengeluaran Transportasi :", totalTransportasi)
	fmt.Println("Total Pengeluaran Akomodasi    :", totalAkomodasi)
	fmt.Println("Total Pengeluaran Makanan      :", totalMakanan)
	fmt.Println("Total Pengeluaran Hiburan      :", totalHiburan)

	totalKeseluruhan = totalTransportasi + totalAkomodasi + totalMakanan + totalHiburan
	fmt.Println("--------------------------------------")
	fmt.Println("Total Seluruh Pengeluaran     :", totalKeseluruhan)
}

func saranPenghematan() {

}

func sequentialSearch() {
	
}

func binarySearch() {

}

func selectionSortJumlah() {

}

func insertionSortKategori() {

}

func tampilkanLaporan() {

}

func menuAwal(data *arrPengeluaran) {
	var pilihMenu int
	  for pilihMenu != 0{
        fmt.Println("\nMenu:")
        fmt.Println("1. Tambah Pengeluaran")
        fmt.Println("2. Ubah Pengeluaran")
        fmt.Println("3. Hapus Pengeluaran")
        fmt.Println("4. Tampilkan Laporan")
        fmt.Println("5. Saran Penghematan")
        fmt.Println("6. Cari Pengeluaran")             /*sequential*/
        fmt.Println("7. Cari Pengeluaran")             /*binary*/
        fmt.Println("8. Urutkan berdasarkan Jumlah")   /*Selection Sort*/
        fmt.Println("9. Urutkan berdasarkan Kategori") /*Insertion Sort*/
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihanMenu)
		
		switch pilihMenu {
			case 0:
				fmt.Print("Program Selesai")
			case 1:tambahPengeluaran(&data)
			case 2:ubahPengeluaran(&data)
			case 3:hapusPengeluaran(&data)
			case 4:hitungTotalPengeluaran
			case 5:saranPenghematan
			case 6:sequentialSearch
			case 7:selectionSortJumlah
			case 8:insertionSortKategori
			case 9:tampilkanLaporan
		}
	}
}
  
