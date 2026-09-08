package minitask

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(list []int) (string, error)
}

type Bank struct {
	Name string
}

func (b Bank) Pay(product []int) (string, error) {
	total := 0

	for _, v := range product {
		total += v
	}

	return fmt.Sprintf("Pembayaran bank berhasil! Total = Rp.%d", total), nil
}

type Online struct {
	Name string
}

func (o Online) Pay(product []int) (string, error) {
	total := 0

	for _, v := range product {
		total += v
	}

	return fmt.Sprintf("Pembayaran online berhasil! Total = Rp.%d", total), nil
}

type Fiktif struct {
	listPrice []int
}

func (f *Fiktif) Pay(product []int) (string, error) {
	total := 0

	for _, v := range product {

		if v == 0 {
			return "", fmt.Errorf("Pembayaran harus lebih dari 0")
		}
		f.listPrice = append(f.listPrice, v)

		total += v
	}
	return "", nil
}

func (f Fiktif) GetList() string {
	total := 0

	for _, v := range f.listPrice {
		total += v
	}

	return fmt.Sprintf("Pembayaran fiktif berhasil! Total = Rp.%d", total)
}
