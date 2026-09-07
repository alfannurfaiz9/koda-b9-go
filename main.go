package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alfannurfaiz9/koda-b9-go.git/internal"
	"github.com/alfannurfaiz9/koda-b9-go.git/internal/minitask"
	"github.com/alfannurfaiz9/koda-b9-go.git/internal/model"
)

func main() {
	internal.Choises()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "0":
			fmt.Println("--------------------------------")
			fmt.Println("Anda telah keluar!")
			fmt.Println("--------------------------------")
			return
		case "1":
			fmt.Println("--------------------------------")
			luas, keliling := minitask.Gabungan(1, 3)

			fmt.Println("Func Luas = ", minitask.HitungLuas(1, 3))
			fmt.Println("Func Keliling = ", minitask.HitungKeliling(1, 3))

			fmt.Printf("Func Gabungan, Luas = %d \n", luas)
			fmt.Printf("Func Gabungan, Keliling = %d \n", keliling)
			fmt.Println("--------------------------------")
			internal.Choises()
		case "2":
			fmt.Println("--------------------------------")
			minitask.PrintJendela(4, 8)
			fmt.Println("--------------------------------")
			internal.Choises()
		case "3":
			fmt.Println("--------------------------------")
			minitask.PrintNum()
			fmt.Println("--------------------------------")
			internal.Choises()
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
			internal.Choises()
		default:
			fmt.Println("--------------------------------")
			fmt.Println("Pilihan tidak tersedia!")
			fmt.Println("--------------------------------")
			internal.Choises()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
