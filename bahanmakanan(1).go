package main

import (
	"fmt"
	"os"
	"os/exec"
)

type bahanMakanan struct{
	nama string
	stok int
	tanggalKadaluwarsa string
}

const NMAX int = 1000
type inventori[NMAX]bahanMakanan
var banyakData int

func main(){
	menu()
	pilihan()
	penutup()
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
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

func pilihan(){
	var data inventori
	var pilih int
	for {
		fmt.Scan(&pilih)
		fmt.Println()
		switch pilih{
			case 1:
				tambahStok(&data)
			case 2:
				tampilkanStok(data)
			case 3:
				ubahData(&data)
			case 4:
				hapusStock(&data)
			case 5:
				urutkanStok(&data)
			case 6:
				peringatanKadaluwarsa(&data)
			case 7:
				laporanTotal(&data)
			case 8:
				cariStok(data)
			}
			if pilih == 0{
			break
		}
	}
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

func hapusStock(b *inventori){
	var i int
	var target string
	var ditemukan bool = false
	
	fmt.Println("Masukkan Nama Bahan yang ingin dihapus! ")
	fmt.Scan(&target)
	
	for i = 0; i < banyakData; i++{
		if b[i].nama == target{
			ditemukan = true
			b[i].nama = ""
			b[i].stok = 0
			b[i].tanggalKadaluwarsa = ""
			banyakData--
			fmt.Println("✅ Data berhasil dihapus!")
			break
		}
	}
	if ditemukan == false{
			fmt.Println("Data Tidak Ditemukan ❌")
		}
}

func urutkanStok(b *inventori){
	for i := 0; i < banyakData-1; i++ {
		for j := 0; j < banyakData-i-1; j++ {
			if b[j].stok > b[j+1].stok {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
}

func peringatanKadaluwarsa(b *inventori){
	var tanggalSekarang string
	fmt.Print("Masukkan tanggal sekarang (dd-mm-year) : ")
	fmt.Scan(&tanggalSekarang)

	fmt.Println("Bahan makanan yang mendekati kadaluarsa : ")
	for i := 0 ; i < banyakData ; i++ {
		if b[i].tanggalKadaluwarsa < tanggalSekarang{
			fmt.Printf ("%s, %d, %s", b[i].nama, b[i].stok, b[i].tanggalKadaluwarsa)
		}
	}
}

func laporanTotal(b *inventori){
	var totalStok int
	for i := 0 ; i < banyakData ; i++{
		totalStok += b[i].stok
	}
	fmt.Printf("Total stok bahan makanan : %i\n", totalStok)
}

func cariStok(b inventori){
	var target string 
	var ditemukan bool = false

	fmt.Println("Masukkan nama bahan yang ingin di cari : ")
	fmt.Scan(&target)

	for i := 0 ; i < banyakData ; i++ {
		if b[i].nama == target{
			ditemukan = true
			fmt.Printf("Bahan ditemukan : %s, %d, %s", b[i].nama, b[i].stok, b[i].tanggalKadaluwarsa)
			break
		}
	}
	if ditemukan == false {
		fmt.Println("Bahan tidak di temukan ❌ ")
	}
}


func penutup(){
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("👋 Terima kasih telah menggunakan aplikasi. Sampai jumpa!")
	fmt.Println()
	fmt.Println("=========================================================")
}