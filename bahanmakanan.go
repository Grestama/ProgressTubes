package main

import "fmt"

type bahanMakanan struct{
	nama string
	stok int
	tanggalKadaluwarsa string
}

const NMAX int = 1000
type inventori[NMAX]bahanMakanan
var banyakData int

func main(){
	var data inventori
	var pilih int
	for {
		menu()
		fmt.Scan(&pilih)
		fmt.Println()
		switch pilih{
			case 1:
				tambahStok(&data)
			case 2:
				tampilkanStok(data)
			case 3:
				ubahData(&data)
			case 0: 
				fmt.Println("👋 Terima kasih telah menggunakan aplikasi. Sampai jumpa!")
			default:
				fmt.Println("Pilihan Tidak Valid")
		}
	}
}


func menu(){
	fmt.Println("\n===== Aplikasi Manajemen Stok Bahan Makanan =====")
	fmt.Println("1. Tambah Stok Bahan Makanan")
	fmt.Println("2. Tampilkan Stok")
	fmt.Println("3. Ubah Stok Bahan Makanan")
	fmt.Println("4. Hapus Stok Bahan Makanan")
	fmt.Println("5. Urutkan Stok")
	fmt.Println("6. Peringatan Kadaluwarsa")
	fmt.Println("7. Laporan Total")
	fmt.Println("8. Cari Stok")
	fmt.Println("0. Keluar")
	fmt.Println("\n===== Silahkan Pilih Menu Diatas =====")
	fmt.Println("Pilih (Masukkan Angka): ")
}

func tambahStok(b *inventori){
	var i, banyak int
	fmt.Print("\nBerapa banyak data yang ingin ditambahkan? ")
	fmt.Scan(&banyak)
	for i = 0; i < banyak; i++{
		if banyakData > NMAX{
		fmt.Println("GUDANG OVERLOAD !!!")
		break
		}
		
		fmt.Print("Nama: ")
		fmt.Scan(&b[i].nama)

		fmt.Print("Jumlah Stok: ")
		fmt.Scan(&b[i].stok)

		fmt.Print("Tanggal Kedaluwarsa (dd-mm-yyyy): ")
		fmt.Scan(&b[i].tanggalKadaluwarsa)
		
		banyakData++
		fmt.Println("✅ Data berhasil ditambahkan!")
	}
	fmt.Println()
}

func tampilkanStok(b inventori){
	var i int
	var nomor int
	
	nomor = 1
	fmt.Println("Data Tercatat: ")
	for i = 0; i < banyakData; i++{
		if b[i].nama != ""{
		fmt.Printf("%d. %s, %d, %s",nomor, b[i].nama, b[i].stok, b[i].tanggalKadaluwarsa)
		fmt.Println()
		nomor++
		}
	}
	fmt.Println()
}

func ubahData(b *inventori){
	var i int
	var target string
	var ditemukan bool = false
	
	fmt.Println("Masukkan Nama Bahan yang ingin diganti! ")
	fmt.Scan(&target)
	
	for i = 0; i < banyakData; i++{
		if b[i].nama == target{
			ditemukan = true
			fmt.Println("Data Ditemukan!!!")
			fmt.Println("Nama: ", b[i].nama)
			fmt.Println("Nama: ", b[i].stok)
			fmt.Println("Nama: ", b[i].tanggalKadaluwarsa)
			
			fmt.Printf("\nMasukkan Nama Bahan Baru: ")
			fmt.Printf("\nNama: ")
			fmt.Scan(&b[i].nama)

			fmt.Printf("\nMasukkan Stok Baru: ")
			fmt.Scan(&b[i].stok)

			fmt.Printf("\nMasukkan Tanggal Kadaluwarsa Baru:  ")
			fmt.Scan(&b[i].tanggalKadaluwarsa)
		
			fmt.Println("✅ Data berhasil diubah!")
			break
		}
	}
	if ditemukan == false{
			fmt.Println("Data Tidak Ditemukan ❌")
		}
}