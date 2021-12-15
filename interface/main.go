package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luasPersegiPanjang := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)

	persegi := Persegi{Sisi: 5}
	luasPersegi := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi:", luasPersegi)

	asal := Asal{Angka: 5}
	luasAsal := SeberapaLuas(asal)
	fmt.Println("Luas Asal:", luasAsal)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}