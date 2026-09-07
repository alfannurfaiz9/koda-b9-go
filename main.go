package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alfannurfaiz9/koda-b9-go.git/internal/minitask"
	"github.com/alfannurfaiz9/koda-b9-go.git/internal/model"
)

func main() {
	choises()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "0":
			return
		case "1":
			fmt.Println("--------------------------------")
			luas, keliling := minitask.Gabungan(1, 3)

			fmt.Println("Func Luas = ", minitask.HitungLuas(1, 3))
			fmt.Println("Func Keliling = ", minitask.HitungKeliling(1, 3))

			fmt.Printf("Func Gabungan, Luas = %d \n", luas)
			fmt.Printf("Func Gabungan, Keliling = %d \n", keliling)
			fmt.Println("--------------------------------")
			choises()
		case "2":
			fmt.Println("--------------------------------")
			minitask.PrintJendela(4, 8)
			fmt.Println("--------------------------------")
			choises()
		case "3":
			fmt.Println("--------------------------------")
			minitask.PrintNum()
			fmt.Println("--------------------------------")
			choises()
		case "4":
			fmt.Println("--------------------------------")
			alfan := model.Biodata{
				Nama:             "alfan nurfaiz",
				Foto:             "https://media.licdn.com/dms/image/v2/D5603AQG1_s2U-uEcEw/profile-displayphoto-crop_800_800/B56aBGgkhNGUAI-/0/1787889339987?e=1790208000&v=beta&t=-qhe1Kp1bLhTatqFLJqznW62EZU606PSmQegDUIiyU8",
				Email:            "alfannurfaizwork@gmail.com",
				Umur:             26,
				NomorTelepon:     "62895350096363",
				StatusPernikahan: false,
				RiwayatPendidikan: model.RiwayatPendidikan{
					Nama:    "SMK Muhammadiyah 1 Purwokerto",
					Jurusan: "Multimedia",
				},
			}
			fmt.Println(alfan)
			fmt.Println("--------------------------------")
			choises()
		default:
			fmt.Println("--------------------------------")
			fmt.Println("Pilihan tidak tersedia!")
			fmt.Println("--------------------------------")
			choises()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func choises() {
	fmt.Println("--------------------------------")
	fmt.Println("Masukkan pilihan:")
	fmt.Println("1. Luas dan keliling persegi panjang")
	fmt.Println("2. Print jendela")
	fmt.Println("3. Sisipkan angka ke dalam slice")
	fmt.Println("4. Biodata")
	fmt.Println("0. Exit")
	fmt.Println("--------------------------------")
}
