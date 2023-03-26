package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	bio1 := Biodata{"Agung Wahyono", "DKI Jakarta", "Belum / Tidak Bekerja", "Ingin mempelajari Back-end Golang"}
	bio2 := Biodata{"Bara Ilham Bakti", "Riau", "Pegawai Negeri Sipil", "Bahasanya sedang populer"}
	bio3 := Biodata{"Cindy Loreta", "Jambi", "Pelajar / Mahasiswa", "Meningkatkan skill"}
	bio4 := Biodata{"Deny Mulyono", "Jawa Tengah", "Pelajar / Mahasiswa", "Punya Komunitas yang baik"}
	bio5 := Biodata{"Eko Supriyadi", "Jawa Tengah", "Pelajar / Mahasiswa", "Mirip dengan bahasa C"}
	bio6 := Biodata{"Fella Naziq Gadis", "Bali", "Mengurus Rumah Tangga", "Untuk mempelajari hal baru"}
	bio7 := Biodata{"Gunawan Bahri", "Jawa Barat", "Karyawan Swasta", "Untuk mempelajari hal baru"}
	bio8 := Biodata{"Hamdani Koestoro", "Jawa Timur", "Pensiunan", "Sering muncul di proyek Open Source"}
	bio9 := Biodata{"Ikhsan Bima Basra", "NTB", "Karyawan Swasta", "Meningkatkan skill"}
	bio10 := Biodata{"Jopi Ajitajan", "Maluku", "Karyawan Swasta", "Ingin meningkatkan skill"}

	argsRaw := os.Args[1:]

	for _, arg := range argsRaw {
		switch arg {
		case "1":
			display(bio1)
		case "2":
			display(bio2)
		case "3":
			display(bio3)
		case "4":
			display(bio4)
		case "5":
			display(bio5)
		case "6":
			display(bio6)
		case "7":
			display(bio7)
		case "8":
			display(bio8)
		case "9":
			display(bio9)
		case "10":
			display(bio10)
		default:
			fmt.Println("Data Tidak Ditemukan!")
		}
	}
}

func display(bio Biodata) {
	fmt.Println("Nama     :", bio.Nama)
	fmt.Println("Alamat   :", bio.Alamat)
	fmt.Println("Pekerjaan:", bio.Pekerjaan)
	fmt.Println("Alasan   :", bio.Alasan)
}
