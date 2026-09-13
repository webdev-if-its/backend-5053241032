package main

import (
	"fmt"
)

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {

	return float64(qty) * hargaSatuan
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	if len(qty) != len(hargaSatuan) {
		return 0
	}
	var temp float64
	for i := range qty {
		temp += (float64(qty[i]) * hargaSatuan[i])
	}
	return temp
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {

	return total + (total * tarifPajak)
}

func HitungDiskon(total float64) float64 {
	var diskon float64

	if total < 500000 {
		diskon = 0
	} else if total >= 500000 && total < 1000000 {
		diskon = 0.05
	} else {
		diskon = 0.1
	}
	return total * diskon
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	var total = HitungTotalPesanan(qty, hargaSatuan)
	total -= HitungDiskon(total)
	var habisPajak = TerapkanPajak(total, tarifPajak)

	return habisPajak
}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {

	if len(qty) != len(hargaSatuan) {
		return false, "Jumlah barang dan harga satuan tidak sama"
	}

	for i := range qty {
		if qty[i] < 0 {
			return false, "Jumlah barang tidak boleh negatif"
		}
		if hargaSatuan[i] < 0 {
			return false, "Harga satuan tidak boleh negatif"
		}
	}

	return true, ""
}

func TentukanStatus(total float64) string {
	if total > 1000000 {
		return "Prioritas"
	} else if total > 100000 {
		return "Reguler"
	}
	return "Hemat"
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {

	var subTotalPesanan = HitungTotalPesanan(qty, hargaSatuan)
	var nominalDiskon = HitungDiskon(subTotalPesanan)
	var total = TotalSetelahDiskon(qty, hargaSatuan, tarifPajak)
	var status = TentukanStatus(total)

	return fmt.Sprintf("Subtotal= %.2f\nDiskon=%.2f\nTotal=%.2f\nStatus:%s\n", subTotalPesanan, nominalDiskon, total, status)
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	var jumlah float64
	for _, h := range harga {
		jumlah += h
	}
	return jumlah
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	panic("belum diimplementasikan")
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
