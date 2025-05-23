package main

import "fmt"

const maxData = 1000
type Pengeluaran struct{
	transportasi, akomodasi, makanan, hiburan float64
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
	for {
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
				fmt.Println(" ")
				return
			case 1:
				fmt.Print("Masukan budget Transportasi:")
				fmt.Scan(&tambah)
				A[0].transportasi += tambah
				
			case 2:
				fmt.Print("Masukan budget Akomodasi:")
				fmt.Scan(&tambah)
				A[0].akomodasi += tambah
			case 3:
				fmt.Scan(&tambah)
				A[0].makanan += tambah
			case 4:
				fmt.Scan(&tambah)
				A[0].hiburan += tambah
		}
	}
}

func ubahPengeluaran() {

}

func hapusPengeluaran() {

}

func hitungTotalPengeluaran() {

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
	  for {
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
				return
			case 1:tambahPengeluaran(data)
			case 2:ubahPengeluaran
			case 3:hapusPengeluaran
			case 4:hitungTotalPengeluaran
			case 5:saranPenghematan
			case 6:sequentialSearch
			case 7:selectionSortJumlah
			case 8:insertionSortKategori
			case 9:tampilkanLaporan
		}
	}
}
  