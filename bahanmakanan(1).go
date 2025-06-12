package main
import ("fmt")

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
	pilihanMenu()
	penutup()
}

func menu(){
	fmt.Println("\n===== Aplikasi Manajemen Stok Bahan Makanan =====")
	fmt.Println("1. Tambah Stok Bahan Makanan")
	fmt.Println("2. Tampilkan Stok")
	fmt.Println("3. Ubah Stok Bahan Makanan")
	fmt.Println("4. Hapus Stok Bahan Makanan")
	fmt.Println("5. Urutkan Stok")
	fmt.Println("6. Peringatan Kadaluwarsa")
	fmt.Println("7. Laporan Total Jumlah stok")
	fmt.Println("8. Cari Stok")
	fmt.Println("9. Cari Stok Terbanyak/ Terendah")
	fmt.Println("0. Keluar")
	fmt.Println("\n===== Silahkan Pilih Menu Diatas =====")
	fmt.Println("Pilih (Masukkan Angka): ")
}

func pilihanMenu(){
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
			case 9:
				cariStokEkstrim(&data)
			case 0:
				return
			default:
				fmt.Println("Pilihan tidak valid!")
			}
			menu()
		}
}

func urutkanStok(b *inventori){
	pilihanSort(b)
}

func pilihanSort(b *inventori){
    var pilih int
    for {
		fmt.Println("==========================")
        fmt.Println("1. Urutkan Stok Descending")
        fmt.Println("2. Urutkan Stok Ascending")
        fmt.Println("0. Kembali ke Menu Utama")
		fmt.Println("==========================")
        fmt.Print("Pilih (Masukkan Angka): ")
        fmt.Scan(&pilih)
        fmt.Println()
		
        switch pilih {
        case 1:
            urutkanStokDescending(b)
            tampilkanStok(*b)
        case 2:
            urutkanStokAscending(b)
            tampilkanStok(*b)
        case 0:
            return
        default:
            fmt.Println("Pilihan tidak valid!")
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
		fmt.Scan(&b[banyakData].nama)

		fmt.Print("Jumlah Stok: ")
		fmt.Scan(&b[banyakData].stok)

		fmt.Print("Tanggal Kedaluwarsa (dd-mm-yyyy): ")
		fmt.Scan(&b[banyakData].tanggalKadaluwarsa)
		
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
			fmt.Println("Jumlah Stok: ", b[i].stok)
			fmt.Println("Tanggal Kadaluwarsa: ", b[i].tanggalKadaluwarsa)
			
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
	var i,j int
	var target string
	var ditemukan bool = false
	
	fmt.Println("Masukkan Nama Bahan yang ingin dihapus! ")
	fmt.Scan(&target)
	
	for i = 0; i < banyakData; i++{
		if b[i].nama == target{
			ditemukan = true
			for j = i; j < banyakData-1; j++{
				b[j] = b[j+1]
			}
			b[banyakData-1] = bahanMakanan{}
			banyakData--
			fmt.Println("✅ Data berhasil dihapus!")
			break
		}
	}
	if !ditemukan{
			fmt.Println("Data Tidak Ditemukan ❌")
	}
}

func urutkanStokAscending(b *inventori){
	var i, j int
	var temp bahanMakanan
	
	for i = 1; i < banyakData; i++ {
		temp = b[i]
		j = i - 1
		for j >= 0 && b[j].stok > temp.stok{
			b[j+1] = b[j]
			j--
		}
		b[j+1] = temp
	}
}

func urutkanStokDescending(b *inventori){
	var i, j int
	var temp bahanMakanan
	
	for i = 1; i < banyakData; i++ {
		temp = b[i]
		j = i - 1
		for j >= 0 && b[j].stok < temp.stok{
			b[j+1] = b[j]
			j--
		}
		b[j+1] = temp
	}
}

func peringatanKadaluwarsa(b *inventori){
	var tanggalSekarang string
	fmt.Print("Masukkan tanggal sekarang (dd-mm-year) : ")
	fmt.Scan(&tanggalSekarang)

	fmt.Println("Bahan makanan yang mendekati kadaluarsa : ")
	for i := 0 ; i < banyakData ; i++ {
		if b[i].tanggalKadaluwarsa < tanggalSekarang{
			fmt.Printf ("%s, %d, %s\n", b[i].nama, b[i].stok, b[i].tanggalKadaluwarsa)
		}
	}
}

func laporanTotal(b *inventori){
	var totalStok int
	for i := 0 ; i < banyakData ; i++{
		totalStok += b[i].stok
	}
	fmt.Printf("Total stok bahan makanan : %d\n", totalStok)
}

func cariStok(b inventori){
	var target string 
	var ditemukan bool = false
	var i, j int
	var temp bahanMakanan
	var kiri, kanan, tengah	int

	for i = 1; i < banyakData-1; i++ {
		temp = b[i]
		j = i - 1
		for j >= 0 && b[j].nama > temp.nama{
			b[j+1] = b[j]
			j--
		}
		b[j+1] = temp
	}
	
	fmt.Println("Masukkan nama bahan yang ingin di cari : ")
	fmt.Scan(&target)
	
	kiri = 0
	kanan = banyakData - 1
	
	for kiri <= kanan{
		tengah = (kiri + kanan) / 2
		if b[tengah].nama == target{
			ditemukan = true
			fmt.Printf("Barang Ditemukan : %s, %d, %s\n", b[tengah].nama, b[tengah].stok, b[tengah].tanggalKadaluwarsa)
			break
		}else if b[tengah].nama < target{
			kiri = tengah + 1
		}else{
			kanan = tengah - 1
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

func cariStokEkstrim(b *inventori){
	var minIdx, maxIdx, i int
	
	if banyakData == 0{
		fmt.Println("❌ Data kosong, tidak dapat mencari nilai ekstrem.")
		return
	}
	
	minIdx = 0
	maxIdx = 0
	
	for i = 0; i < banyakData; i++{
		if b[i].stok > b[maxIdx].stok{
			maxIdx = i
		}
		
		if b[i].stok < b[minIdx].stok{
			minIdx = i
		}
	}
	
	fmt.Println("📈 Stok Terbanyak:")
	fmt.Printf("- %s, %d, %s\n", b[maxIdx].nama, b[maxIdx].stok, b[maxIdx].tanggalKadaluwarsa)
	
	fmt.Println("📉 Stok Tersedikit:")
	fmt.Printf("- %s, %d, %s\n", b[minIdx].nama, b[minIdx].stok, b[minIdx].tanggalKadaluwarsa)
}