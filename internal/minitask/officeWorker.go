package minitask

import (
	"fmt"
	"sync"
	"time"
)

func Activity() {
	var wg = sync.WaitGroup{}

	defer fmt.Println("Berangkat kerja")

	wg.Go(mandi)
	wg.Go(buatKopi)
	wg.Go(sarapan)
	wg.Go(merapikanKamar)

	wg.Wait()
}

func mandi() {
	fmt.Println("Mulai mandi")
	time.Sleep(1 * time.Second)
	fmt.Println("Mandi selesai")
}

func buatKopi() {
	fmt.Println("Mulai buat kopi")
	time.Sleep(2 * time.Second)
	fmt.Println("Buat kopi selesai")
}

func sarapan() {
	fmt.Println("Mulai menyiapkan sarapan")
	time.Sleep(3 * time.Second)
	fmt.Println("Menyiapkan sarapan selesai")
}

func merapikanKamar() {
	fmt.Println("Mulai merapikan kamar tidur")
	time.Sleep(4 * time.Second)
	fmt.Println("Merapikan kamar tidur selesai")
}
