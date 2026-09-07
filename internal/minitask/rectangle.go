package minitask

func HitungLuas(p, l uint16) uint16 {
	return p * l
}

func HitungKeliling(p, l uint16) uint16 {
	return 2 * (p + l)
}

func Gabungan(p, l uint16) (luas uint16, keliling uint16) {
	return HitungLuas(p, l), HitungKeliling(p, l)
}
