/*
TUGAS BESAR ALGORITMA PEMROGRAMAN IF-46-02 KELOMPOK 3
		APLIKASI INVENTORI BARANG

1301223235	Nazwa Betha Kirana
1301223410	Umar Khairur Rahman
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

type sembakoM struct {
	nama, jenis   string
	jmlh, stk     int
	tgl, bln, thn int
}

type sembakoK struct {
	nama, jenis         string
	jmlh, tgl, bln, thn int
}

type arrM [1000]sembakoM

var dataM = arrM{
	{"Chocopie", "Snack", 34, 34, 06, 9, 2022},
	{"Susu", "Minuman", 40, 40, 27, 9, 2022},
	{"Chitato", "Snack", 21, 21, 8, 10, 2022},
	{"Piatos", "Snack", 24, 24, 10, 10, 2022},
	{"Teh", "Minuman", 31, 26, 12, 10, 2022},
	{"Kopi", "Minuman", 51, 51, 12, 11, 2022},
	{"Beras", "Pangan", 20, 20, 12, 12, 2022},
	{"Mi", "Pangan", 59, 27, 15, 11, 2022},
	{"Garam", "Bumbu", 10, 10, 16, 11, 2022},
	{"Gula", "Bumbu", 14, 14, 16, 11, 2022},
	{"Lada", "Bumbu", 15, 15, 28, 12, 2022},
	{"Boncabe", "Bumbu", 7, 7, 01, 01, 2023},
	{"Galon", "Minuman", 20, 12, 01, 01, 2023},
	{"Sabun", "Kebersihan", 16, 16, 11, 02, 2023},
	{"Shampo", "Kebersihan", 23, 21, 14, 02, 2023},
}
var nM int = 15

type arrK [1000]sembakoK

var dataK = arrK{
	{"Teh", "Minuman", 5, 20, 10, 2022},
	{"Lada", "Bumbu", 8, 02, 12, 2022},
	{"Mi", "Pangan", 32, 25, 11, 2022},
	{"Galon", "Minuman", 8, 12, 02, 2023},
	{"Shampo", "Kebersihan", 2, 11, 03, 2023},
}
var nK int = 5

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func header() {
	fmt.Println("+---------------------+")
	fmt.Println("|   **                |")
	fmt.Println("|     ***********     |")
	fmt.Println("|      *********      |")
	fmt.Println("|       ******        |")
	fmt.Println("|       **  **        |")
	fmt.Println("|                     |")
	fmt.Println("|  INVENTORI SEMBAKO  |")
	fmt.Println("+---------------------+")
}

func intro() {
	header()
	fmt.Println("     Selamat Datang    ")
	fmt.Scanln()
	clear()
}

func menu() {
	var pick int
	for {
		header()
		fmt.Println("|        MENU         |")
		fmt.Println("+---------------------+")
		fmt.Println("| 1. Daftar Sembako   |")
		fmt.Println("| 2. Admin            |")
		fmt.Println("| 3. Keluar           |")
		fmt.Println("+---------------------+")
		fmt.Print("Pilihan Anda (1/2/3): ")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			tampilDaftar()
		case 2:
			clear()
			admin()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func pilih() {
	fmt.Println("+---------------------+")
	fmt.Println("| 1. Sembako Masuk    |")
	fmt.Println("| 2. Sembako Keluar   |")
	fmt.Println("| 3. Kembali          |")
	fmt.Println("+---------------------+")
	fmt.Print("Pilihan Anda: ")
}

func tampilDaftar() {
	var pick int
	for {
		header()
		fmt.Println("|   DAFTAR SEMBAKO    |")
		pilih()
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			printM()
		case 2:
			clear()
			printK()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func printM() {
	var pick int
	var s string
	var tampilM arrM = dataM
	for {
		header()
		fmt.Println("|    SEMBAKO MASUK    |")
		fmt.Println("+---------------------+---------------------------------------------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %-15s %0s\n", "No", "Nama", "Jenis", "Jumlah", "Stok", "Waktu")
		for i := 0; i < nM; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %-15d %d-%d-%d\n", i+1, tampilM[i].nama, tampilM[i].jenis, tampilM[i].jmlh, tampilM[i].stk, tampilM[i].tgl, tampilM[i].bln, tampilM[i].thn)
		}
		fmt.Println("+---------------------+")
		fmt.Println("| 1. Asc by Nama      |")
		fmt.Println("| 2. Dsc by Nama      |")
		fmt.Println("| 3. Asc by Jenis     |")
		fmt.Println("| 4. Dsc by Jenis     |")
		fmt.Println("| 5. Asc by Stok      |")
		fmt.Println("| 6. Dsc by Stok      |")
		fmt.Println("| 7. Kembali          |")
		fmt.Println("+---------------------+")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			s = "NAMA"
			selMascSTR(&tampilM, s)
		case 2:
			clear()
			s = "NAMA"
			selMdescSTR(&tampilM, s)
		case 3:
			clear()
			s = "JENIS"
			selMascSTR(&tampilM, s)
		case 4:
			clear()
			s = "JENIS"
			selMdescSTR(&tampilM, s)
		case 5:
			clear()
			selMascSTK(&tampilM)
		case 6:
			clear()
			selMdescSTK(&tampilM)
		}
		clear()
		if pick == 7 {
			break
		}
	}

}

func getSTR_M(sm sembakoM, flag string) string {
	if flag == "NAMA" {
		return sm.nama
	} else if flag == "JENIS" {
		return sm.jenis
	} else {
		return ""
	}
}

func selMascSTR(x *arrM, key string) {
	var pass, i, idx int
	var temp sembakoM
	i = 1
	for i <= nM-1 {
		idx = i - 1
		pass = i
		for pass < nM {
			if getSTR_M(x[idx], key) > getSTR_M(x[pass], key) {
				idx = pass
			}
			pass = pass + 1
		}
		temp = x[idx]
		x[idx] = x[i-1]
		x[i-1] = temp
		i = i + 1
	}
}

func selMdescSTR(x *arrM, key string) {
	var pass, i, idx int
	var temp sembakoM
	i = 1
	for i <= nM-1 {
		idx = i - 1
		pass = i
		for pass < nM {
			if getSTR_M(x[idx], key) < getSTR_M(x[pass], key) {
				idx = pass
			}
			pass = pass + 1
		}
		temp = x[idx]
		x[idx] = x[i-1]
		x[i-1] = temp
		i = i + 1
	}
}

func selMascSTK(x *arrM) {
	var pass, i, idx int
	var temp sembakoM
	i = 1
	for i <= nM-1 {
		idx = i - 1
		pass = i
		for pass < nM {
			if x[idx].stk > x[pass].stk {
				idx = pass
			}
			pass = pass + 1
		}
		temp = x[idx]
		x[idx] = x[i-1]
		x[i-1] = temp
		i = i + 1
	}
}

func selMdescSTK(x *arrM) {
	var pass, i, idx int
	var temp sembakoM
	i = 1
	for i <= nM-1 {
		idx = i - 1
		pass = i
		for pass < nM {
			if x[idx].stk < x[pass].stk {
				idx = pass
			}
			pass = pass + 1
		}
		temp = x[idx]
		x[idx] = x[i-1]
		x[i-1] = temp
		i = i + 1
	}
}

func printK() {
	var pick int
	var tampilK arrK = dataK
	for {
		header()
		fmt.Println("|    SEMBAKO KELUAR   |")
		fmt.Println("+---------------------+-----------------------------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %s\n", "No", "Nama", "Jenis", "Jumlah", "Waktu")
		for i := 0; i < nK; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %d-%d-%d\n", i+1, tampilK[i].nama, tampilK[i].jenis, tampilK[i].jmlh, tampilK[i].tgl, tampilK[i].bln, tampilK[i].thn)
		}
		fmt.Println("+---------------------+")
		fmt.Println("| 1. Asc by Nama      |")
		fmt.Println("| 2. Dsc by Nama      |")
		fmt.Println("| 3. Asc by Jumlah    |")
		fmt.Println("| 4. Dsc by Jumlah    |")
		fmt.Println("| 5. Asc by Waktu     |")
		fmt.Println("| 6. Dsc by Waktu     |")
		fmt.Println("| 7. Kembali          |")
		fmt.Println("+---------------------+")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			insKascNM(&tampilK)
		case 2:
			clear()
			insKdescNM(&tampilK)
		case 3:
			clear()
			insKascJMLH(&tampilK)
		case 4:
			clear()
			insKdescJMLH(&tampilK)
		case 5:
			clear()
			insKascWKT(&tampilK)
		case 6:
			clear()
			insKdescWKT(&tampilK)
		}
		clear()
		if pick == 7 {
			break
		}
	}

}

func insKascNM(x *arrK) {
	var pass, i int
	var temp sembakoK
	i = 1
	for i <= nK-1 {
		pass = i
		temp = x[pass]
		for pass > 0 && temp.nama < x[pass-1].nama {
			x[pass] = x[pass-1]
			pass = pass - 1
		}
		x[pass] = temp
		i = i + 1
	}
}

func insKdescNM(x *arrK) {
	var pass, i int
	var temp sembakoK
	i = 1
	for i <= nK-1 {
		pass = i
		temp = x[pass]
		for pass > 0 && temp.nama > x[pass-1].nama {
			x[pass] = x[pass-1]
			pass = pass - 1
		}
		x[pass] = temp
		i = i + 1
	}
}

func insKascJMLH(x *arrK) {
	var pass, i int
	var temp sembakoK
	i = 1
	for i <= nK-1 {
		pass = i
		temp = x[pass]
		for pass > 0 && temp.jmlh < x[pass-1].jmlh {
			x[pass] = x[pass-1]
			pass = pass - 1
		}
		x[pass] = temp
		i = i + 1
	}
}

func insKdescJMLH(x *arrK) {
	var pass, i int
	var temp sembakoK
	i = 1
	for i <= nK-1 {
		pass = i
		temp = x[pass]
		for pass > 0 && temp.jmlh > x[pass-1].jmlh {
			x[pass] = x[pass-1]
			pass = pass - 1
		}
		x[pass] = temp
		i = i + 1
	}
}

func bandingWKT(a, b sembakoK) int {
	if a.thn != b.thn {
		return a.thn - b.thn
	}
	if a.bln != b.bln {
		return a.bln - b.bln
	}
	if a.tgl != b.tgl {
		return a.tgl - b.tgl
	}
	return 0
}

func insKascWKT(x *arrK) {
	var pass int
	var temp sembakoK
	for i := 1; i < nK; i++ {
		temp = x[i]
		pass = i - 1
		for pass >= 0 && bandingWKT(x[pass], temp) > 0 {
			x[pass+1] = x[pass]
			pass = pass - 1
		}
		x[pass+1] = temp
	}
}

func insKdescWKT(x *arrK) {
	var pass int
	var temp sembakoK
	for i := 1; i < nK; i++ {
		temp = x[i]
		pass = i - 1
		for pass >= 0 && bandingWKT(x[pass], temp) < 0 {
			x[pass+1] = x[pass]
			pass = pass - 1
		}
		x[pass+1] = temp
	}
}

func admin() {
	var pick int
	for {
		header()
		fmt.Println("|        ADMIN        |")
		fmt.Println("+---------------------+")
		fmt.Println("| 1. Cari Data        |")
		fmt.Println("| 2. Tambah Data      |")
		fmt.Println("| 3. Ubah Data        |")
		fmt.Println("| 4. Hapus Data       |")
		fmt.Println("| 5. Kembali          |")
		fmt.Println("+---------------------+")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			tampilCari()
		case 2:
			clear()
			tampilTambah()
		case 3:
			clear()
			tampilUbah()
		case 4:
			clear()
			tampilHapus()
		}
		clear()
		if pick == 5 {
			break
		}
	}
}

func tampilCari() {
	var pick int
	for {
		header()
		fmt.Println("|      CARI DATA      |")
		pilih()
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			findM()
		case 2:
			clear()
			findK()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func findM() {
	var pick int
	var key string
	for {
		header()
		fmt.Println("| CARI SEMBAKO MASUK  |")
		fmt.Println("+---------------------+")
		fmt.Println("| 1. By Nama          |")
		fmt.Println("| 2. By Jenis         |")
		fmt.Println("| 3. By Tahun         |")
		fmt.Println("| 4. Kembali          |")
		fmt.Println("+---------------------+")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			key = "NAMA"
			cariSTR_M(key)
		case 2:
			clear()
			key = "JENIS"
			cariSTR_M(key)
		case 3:
			clear()
			cariTHN_M()
		}
		clear()
		if pick == 4 {
			break
		}

	}
}

func cariSTR_M(s string) {
	var jwb string
	var stop bool = false
	var R arrM
	var nR int
	for !stop {
		nR = 0
		header()
		fmt.Printf("|    CARI BY %s     |\n", s)
		fmt.Println("+---------------------------+")
		fmt.Print("Kata Kunci Pencarian: ")
		fmt.Scanln(&jwb)
		fmt.Println("+---------------------------+")
		for i := 0; i < nM; i++ {
			if jwb == getSTR_M(dataM[i], s) {
				R[nR] = dataM[i]
				nR++
			}
		}
		fmt.Printf("  %-3s %-15s %-15s %-15s %-15s %0s\n", "No", "Nama", "Jenis", "Jumlah", "Stok", "Waktu")
		for i := 0; i < nR; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %-15d %d-%d-%d\n", i+1, R[i].nama, R[i].jenis, R[i].jmlh, R[i].stk, R[i].tgl, R[i].bln, R[i].thn)
		}
		if nR == 0 {
			fmt.Println("Data Tidak Ada")
		}
		fmt.Println("+---------------------+")
		fmt.Println("Cari Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func cariTHN_M() {
	var jwb string
	var stop bool = false
	var R arrM
	var nR, thn int
	for !stop {
		nR = 0
		header()
		fmt.Println("|    CARI BY TAHUN    |")
		fmt.Println("+---------------------------+")
		fmt.Print("Kata Kunci Pencarian: ")
		fmt.Scanln(&thn)
		fmt.Println("+---------------------------+")
		for i := 0; i < nM; i++ {
			if thn == dataM[i].thn {
				R[nR] = dataM[i]
				nR++
			}
		}
		fmt.Printf("  %-3s %-15s %-15s %-15s %-15s %0s\n", "No", "Nama", "Jenis", "Jumlah", "Stok", "Waktu")
		for i := 0; i < nR; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %-15d %d-%d-%d\n", i+1, R[i].nama, R[i].jenis, R[i].jmlh, R[i].stk, R[i].tgl, R[i].bln, R[i].thn)
		}
		if nR == 0 {
			fmt.Println("Data Tidak Ada")
		}
		fmt.Println("+---------------------+")
		fmt.Println("Cari Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func findK() {
	var pick int
	var key string
	for {
		header()
		fmt.Println("| CARI SEMBAKO KELUAR |")
		fmt.Println("+---------------------+")
		fmt.Println("| 1. By Nama          |")
		fmt.Println("| 2. By Jenis         |")
		fmt.Println("| 3. By Tahun         |")
		fmt.Println("| 4. Kembali          |")
		fmt.Println("+---------------------+")
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			key = "NAMA"
			cariSTR_K(key)
		case 2:
			clear()
			key = "JENIS"
			cariSTR_K(key)
		case 3:
			clear()
			cariTHN_K()
		}
		clear()
		if pick == 4 {
			break
		}

	}
}

func getSTR_K(sk sembakoK, flag string) string {
	if flag == "NAMA" {
		return sk.nama
	} else if flag == "JENIS" {
		return sk.jenis
	} else {
		return ""
	}
}

func cariSTR_K(s string) {
	var jwb string
	var stop bool = false
	var R arrK
	var nR int
	for !stop {
		nR = 0
		header()
		fmt.Printf("|    CARI BY %s     |\n", s)
		fmt.Println("+---------------------------+")
		fmt.Print("Kata Kunci Pencarian: ")
		fmt.Scanln(&jwb)
		fmt.Println("+---------------------------+")
		for i := 0; i < nK; i++ {
			if jwb == getSTR_K(dataK[i], s) {
				R[nR] = dataK[i]
				nR++
			}
		}
		fmt.Printf("  %-3s %-15s %-15s %-15s %s\n", "No", "Nama", "Jenis", "Jumlah", "Waktu")
		for i := 0; i < nR; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %d-%d-%d\n", i+1, R[i].nama, R[i].jenis, R[i].jmlh, R[i].tgl, R[i].bln, R[i].thn)
		}
		if nR == 0 {
			fmt.Println("Data Tidak Ada")
		}
		fmt.Println("+---------------------+")
		fmt.Println("Cari Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func cariTHN_K() {
	var jwb string
	var stop bool = false
	var R arrK
	var nR, thn int
	for !stop {
		nR = 0
		header()
		fmt.Println("|    CARI BY TAHUN    |")
		fmt.Println("+---------------------------+")
		fmt.Print("Kata Kunci Pencarian: ")
		fmt.Scanln(&thn)
		fmt.Println("+---------------------------+")
		for i := 0; i < nK; i++ {
			if thn == dataK[i].thn {
				R[nR] = dataK[i]
				nR++
			}
		}
		fmt.Printf("  %-3s %-15s %-15s %-15s %s\n", "No", "Nama", "Jenis", "Jumlah", "Waktu")
		for i := 0; i < nR; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %d-%d-%d\n", i+1, R[i].nama, R[i].jenis, R[i].jmlh, R[i].tgl, R[i].bln, R[i].thn)
		}
		if nR == 0 {
			fmt.Println("Data Tidak Ada")
		}
		fmt.Println("+---------------------+")
		fmt.Println("Cari Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func tampilTambah() {
	var pick int
	for {
		header()
		fmt.Println("|     TAMBAH DATA     |")
		pilih()
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			inputM()
		case 2:
			clear()
			inputK()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func inputM() {
	var stop bool = false
	var jwb, nama, jenis string
	var jmlh, stk, tgl, bln, thn int
	for !stop {
		header()
		fmt.Println("|TAMBAH SEMBAKO MASUK |")
		fmt.Println("+---------------------+")
		fmt.Print("Nama Sembako: ")
		fmt.Scanln(&nama)
		fmt.Print("Jenis: ")
		fmt.Scanln(&jenis)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&jmlh)
		fmt.Print("Waktu Masuk (DD MM YYYY): ")
		fmt.Scanln(&tgl, &bln, &thn)
		fmt.Println("+---------------------+")
		fmt.Println("Simpan (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			masukStok(nama, jmlh, &stk)
			dataM[nM].nama = nama
			dataM[nM].jenis = jenis
			dataM[nM].jmlh = jmlh
			dataM[nM].stk = stk
			dataM[nM].tgl = tgl
			dataM[nM].bln = bln
			dataM[nM].thn = thn
			nM++
			fmt.Println("Data Tersimpan")
		}
		fmt.Println("Masukkan Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func inputK() {
	var stop bool = false
	var jwb, nama, jenis string
	var jmlh, tgl, bln, thn int
	for !stop {
		header()
		fmt.Println("|TAMBAH SEMBAKO KELUAR|")
		fmt.Println("+---------------------+")
		fmt.Print("Nama Sembako: ")
		fmt.Scanln(&nama)
		fmt.Print("Jenis: ")
		fmt.Scanln(&jenis)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&jmlh)
		fmt.Print("Waktu Keluar (DD MM YYYY): ")
		fmt.Scanln(&tgl, &bln, &thn)
		fmt.Println("+---------------------+")
		fmt.Println("Simpan (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			keluarStok(nama, tgl, bln, thn, jmlh)
			dataK[nK].nama = nama
			dataK[nK].jenis = jenis
			dataK[nK].jmlh = jmlh
			dataK[nK].tgl = tgl
			dataK[nK].bln = bln
			dataK[nK].thn = thn
			nK++
			fmt.Println("Data Tersimpan")
		}
		fmt.Println("Masukkan Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func masukStok(s string, x int, stk *int) {
	var in bool = false
	var ns int
	for i := 0; i < nM; i++ {
		if dataM[i].nama == s {
			if !in {
				dataM[i].stk += x
				ns = dataM[i].stk
				in = true
			} else {
				dataM[i].stk = ns
			}

		}
	}
	if !in {
		*stk = x
	} else {
		*stk = ns
	}
}

func keluarStok(s string, d, m, y, x int) {
	var in bool = false
	var ns int
	for i := 0; i < nM; i++ {
		if dataM[i].nama == s && duluanWKT(d, m, y, dataM[i].tgl, dataM[i].bln, dataM[i].thn) > 0 {
			if !in {
				dataM[i].stk -= x
				ns = dataM[i].stk
				in = true
			} else {
				dataM[i].stk = ns
			}
		}
	}
}

func duluanWKT(tgl1, bln1, thn1, tgl2, bln2, thn2 int) int {
	if thn1 != thn2 {
		return thn1 - thn2
	}
	if bln1 != bln2 {
		return bln1 - bln2
	}
	if tgl1 != tgl2 {
		return tgl1 - tgl2
	}
	return 0
}

func tampilUbah() {
	var pick int
	for {
		header()
		fmt.Println("|       UBAH DATA      |")
		pilih()
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			ubahM()
		case 2:
			clear()
			ubahK()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func ubahM() {
	var stop bool = false
	var jwb, nama, jenis string
	var jmlh, stk, tgl, bln, thn, idx int
	for !stop {
		header()
		fmt.Println("| UBAH SEMBAKO MASUK  |")
		fmt.Println("+---------------------+---------------------------------------------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %-15s %0s\n", "No", "Nama", "Jenis", "Jumlah", "Stok", "Waktu")
		for i := 0; i < nM; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %-15d %d-%d-%d\n", i+1, dataM[i].nama, dataM[i].jenis, dataM[i].jmlh, dataM[i].stk, dataM[i].tgl, dataM[i].bln, dataM[i].thn)
		}
		fmt.Println("+---------------------------+")
		fmt.Print("Nomor Data yang Ingin Diubah: ")
		fmt.Scanln(&idx)
		fmt.Println("+---------------------+")
		fmt.Println("|   INPUT DATA BARU   |")
		fmt.Println("+---------------------+")
		fmt.Print("Nama Sembako: ")
		fmt.Scanln(&nama)
		fmt.Print("Jenis: ")
		fmt.Scanln(&jenis)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&jmlh)
		fmt.Print("Waktu Masuk (DD MM YYYY): ")
		fmt.Scanln(&tgl, &bln, &thn)
		fmt.Println("+---------------------+")
		fmt.Println("Simpan (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			ubahStokM(nama, jmlh, dataM[idx-1].jmlh, &stk)
			dataM[idx-1].nama = nama
			dataM[idx-1].jenis = jenis
			dataM[idx-1].jmlh = jmlh
			dataM[idx-1].stk = stk
			dataM[idx-1].tgl = tgl
			dataM[idx-1].bln = bln
			dataM[idx-1].thn = thn
			fmt.Println("Data Terubah")
		}
		fmt.Println("Ubah Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func ubahK() {
	var stop bool = false
	var jwb, nama, jenis string
	var jmlh, tgl, bln, thn, idx int
	for !stop {
		header()
		fmt.Println("| UBAH SEMBAKO KELUAR |")
		fmt.Println("+---------------------+-----------------------------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %s\n", "No", "Nama", "Jenis", "Jumlah", "Waktu")
		for i := 0; i < nK; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %d-%d-%d\n", i+1, dataK[i].nama, dataK[i].jenis, dataK[i].jmlh, dataK[i].tgl, dataK[i].bln, dataK[i].thn)
		}
		fmt.Println("+--------------------------+")
		fmt.Print("Nomor Data yang Ingin Diubah: ")
		fmt.Scanln(&idx)
		fmt.Println("+---------------------+")
		fmt.Println("|   INPUT DATA BARU   |")
		fmt.Println("+---------------------+")
		fmt.Print("Nama Sembako: ")
		fmt.Scanln(&nama)
		fmt.Print("Jenis: ")
		fmt.Scanln(&jenis)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&jmlh)
		fmt.Print("Waktu Keluar (DD MM YYYY): ")
		fmt.Scanln(&tgl, &bln, &thn)
		fmt.Println("+---------------------+")
		fmt.Println("Simpan (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			ubahStokK(nama, tgl, bln, thn, dataK[idx-1].tgl, dataK[idx-1].bln, dataK[idx-1].thn, dataK[idx-1].jmlh, jmlh)
			dataK[idx-1].nama = nama
			dataK[idx-1].jenis = jenis
			dataK[idx-1].jmlh = jmlh
			dataK[idx-1].tgl = tgl
			dataK[idx-1].bln = bln
			dataK[idx-1].thn = thn
			fmt.Println("Data Terubah")
		}
		fmt.Println("Ubah Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func ubahStokM(s string, x, y int, stk *int) {
	var in bool = false
	var ns int
	for i := 0; i < nM; i++ {
		if dataM[i].nama == s {
			if !in {
				x -= y
				dataM[i].stk += x
				ns = dataM[i].stk
				in = true
			} else {
				dataM[i].stk = ns
			}

		}
	}
	if !in {
		*stk = x
	} else {
		*stk = ns
	}
}

func ubahStokK(s string, d1, m1, y1, d2, m2, y2, z, x int) {
	var in bool = false
	var ns int
	for i := 0; i < nM; i++ {
		if dataM[i].nama == s && duluanWKT(d1, m1, y1, dataM[i].tgl, dataM[i].bln, dataM[i].thn) > 0 {
			if !in {
				if duluanWKT(d2, m2, y2, dataM[i].tgl, dataM[i].bln, dataM[i].thn) > 0 {
					x -= z
				}
				dataM[i].stk -= x
				ns = dataM[i].stk
				in = true
			} else {
				dataM[i].stk = ns
			}
		}
	}
}

func tampilHapus() {
	var pick int
	for {
		header()
		fmt.Println("|      HAPUS DATA     |")
		pilih()
		fmt.Scanln(&pick)
		switch pick {
		case 1:
			clear()
			hapusM()
		case 2:
			clear()
			hapusK()
		}
		clear()
		if pick == 3 {
			break
		}
	}
}

func hapusM() {
	var stop bool = false
	var jwb string
	var idx int
	for !stop {
		header()
		fmt.Println("| HAPUS SEMBAKO MASUK |")
		fmt.Println("+---------------------+--------------------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %-15s %0s\n", "No", "Nama", "Jenis", "Jumlah", "Stok", "Waktu")
		for i := 0; i < nM; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %-15d %d-%d-%d\n", i+1, dataM[i].nama, dataM[i].jenis, dataM[i].jmlh, dataM[i].stk, dataM[i].tgl, dataM[i].bln, dataM[i].thn)
		}
		fmt.Println("+---------------------------+")
		fmt.Print("Nomor Data yang Ingin Dihapus: ")
		fmt.Scanln(&idx)
		fmt.Println("+---------------------+")
		fmt.Println("|      HAPUS DATA     |")
		fmt.Println("+---------------------+")
		fmt.Println("Nama Sembako:", dataM[idx-1].nama)
		fmt.Println("Jenis:", dataM[idx-1].jenis)
		fmt.Println("Jumlah:", dataM[idx-1].jmlh)
		fmt.Println("Stok:", dataM[idx-1].stk)
		fmt.Printf("Waktu Masuk : %d-%d-%d\n", dataM[idx-1].tgl, dataM[idx-1].bln, dataM[idx-1].thn)
		fmt.Println("+---------------------+")
		fmt.Println("Hapus (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			hapusStok(dataM[idx-1].nama, dataM[idx-1].jmlh)
			for i := idx - 1; i+1 < nM+1; i++ {
				dataM[i] = dataM[i+1]
			}
			nM--
			fmt.Println("Data Terhapus")
		}
		fmt.Println("Hapus Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func hapusK() {
	var stop bool = false
	var jwb string
	var idx int
	for !stop {
		header()
		fmt.Println("| HAPUS SEMBAKO KELUAR|")
		fmt.Println("+---------------------+----------------------+")
		fmt.Printf("  %-3s %-15s %-15s %-15s %s\n", "No", "Nama", "Jenis", "Jumlah", "Waktu")
		for i := 0; i < nK; i++ {
			fmt.Printf("  %-3d %-15s %-15s %-15d %d-%d-%d\n", i+1, dataK[i].nama, dataK[i].jenis, dataK[i].jmlh, dataK[i].tgl, dataK[i].bln, dataK[i].thn)
		}
		fmt.Println("+---------------------------+")
		fmt.Print("Nomor Data yang Ingin Dihapus: ")
		fmt.Scanln(&idx)
		fmt.Println("+---------------------+")
		fmt.Println("|      HAPUS DATA     |")
		fmt.Println("+---------------------+")
		fmt.Println("Nama Sembako:", dataK[idx-1].nama)
		fmt.Println("Jenis:", dataK[idx-1].jenis)
		fmt.Println("Jumlah:", dataK[idx-1].jmlh)
		fmt.Printf("Waktu Keluar : %d-%d-%d\n", dataK[idx-1].tgl, dataK[idx-1].bln, dataK[idx-1].thn)
		fmt.Println("+---------------------+")
		fmt.Println("Hapus (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		if jwb == "y" {
			hapusStok(dataK[idx-1].nama, dataK[idx-1].jmlh)
			for i := idx - 1; i+1 < nK+1; i++ {
				dataK[i] = dataK[i+1]
			}
			nK--
			fmt.Println("Data Terhapus")
		}
		fmt.Println("Hapus Data Lagi (y/t)?")
		fmt.Scanln(&jwb)
		for jwb != "y" && jwb != "t" {
			fmt.Scanln(&jwb)
		}
		clear()
		if jwb == "t" {
			stop = true
		}
	}
}

func hapusStok(s string, x int) {
	var in bool = false
	var ns int
	for i := 0; i < nM; i++ {
		if dataM[i].nama == s {
			if !in {
				dataM[i].stk -= x
				ns = dataM[i].stk
				in = true
			} else {
				dataM[i].stk = ns
			}

		}
	}
}

func outro() {
	fmt.Println("          **           ")
	fmt.Println("         ***           ")
	fmt.Println("        ***            ")
	fmt.Println("    ||**********       ")
	fmt.Println("    ||**********       ")
	fmt.Println("    ||*********        ")
	fmt.Println("      ********         ")
	fmt.Println("  Sampai Jumpa Lagi :] ")
	fmt.Scanln()
	clear()
}

func main() {
	intro()
	menu()
	outro()
}
