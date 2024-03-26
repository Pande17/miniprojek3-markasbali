package main

import (
	"fmt"
	"miniprojek3/config"
	"miniprojek3/usecase"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}
func main() {

	Init()
	chooseOrder := 0
	fmt.Println("Daftar buku manajemen")
	fmt.Println("================================================")
	fmt.Println("Silahkan Pilih :")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Edit Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Daftar Buku")
	fmt.Println("5. Import File Buku")
	fmt.Println("6. Keluar")
	fmt.Println("Tekan pilihanmu")
	_, err := fmt.Scanln(&chooseOrder)

	if err != nil {
		fmt.Println("error: ", err)
	}

	if chooseOrder == 1 {
		usecase.TambahBuku(&gorm.DB{})
	} else if chooseOrder == 2 {
		usecase.EditBuku(&gorm.DB{})
	} else if chooseOrder == 3 {
		usecase.DeleteBuku(&gorm.DB{})
	} else if chooseOrder == 4 {
		usecase.ListBuku(&gorm.DB{})
	} else if chooseOrder == 5 {
		usecase.ImportFile()
	} else if chooseOrder == 7 {
		os.Exit(0)
	}
	main()
}