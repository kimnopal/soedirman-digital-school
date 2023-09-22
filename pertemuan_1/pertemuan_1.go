package pertemuan_1

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	fmt.Println("=== Pertemuan 1 ===")
	fmt.Println(HitungLuasPersegi(2))
	fmt.Println(HitungLuasSegitiga(2, 8))
	fmt.Println(HitungLuasLingkaran(3))
	fmt.Println(HitungEnergiKinetik(20, 10, 5))
	fmt.Println(HitungEnergiPotensial(20, 10))
	fmt.Println(HitungFrekuensi(50))
	fmt.Println(KonversiSuhu(16, "C"))
	fmt.Println()
}

func HitungLuasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func HitungLuasSegitiga(alas float64, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}

func HitungLuasLingkaran(jari_jari float64) float64 {
	return math.Phi * jari_jari * jari_jari
}

func HitungEnergiKinetik(massa float64, gravitasi float64, tinggi float64) float64 {
	return massa * gravitasi * tinggi
}

func HitungEnergiPotensial(massa float64, kecepatan float64) float64 {
	return 0.5 * massa * kecepatan * kecepatan
}

func HitungFrekuensi(periode float64) float64 {
	return 1 / periode
}

func KonversiSuhu(suhu float64, satuan string) (map[string]float64, error) {
	switch satuan {
	case "C":
		return map[string]float64{
			"F": (9 / 5 * suhu) + 32,
			"R": (4 / 5 * suhu),
			"K": suhu + 273,
		}, nil
	case "F":
		return map[string]float64{
			"C": (5 / 9 * suhu) - 32,
			"R": (4 / 9 * suhu),
			"K": 5/9*(suhu-32) + 273,
		}, nil
	case "R":
		return map[string]float64{
			"C": (5 / 4 * suhu) * suhu,
			"F": (9 / 4 * suhu) + 32,
			"K": (5 / 4 * suhu) + 273,
		}, nil
	case "K":
		return map[string]float64{
			"C": suhu - 273,
			"F": 9/5*(suhu-273) + 32,
			"R": (4 / 5 * suhu) - 273,
		}, nil
	default:
		return map[string]float64{}, errors.New("satuan tidak valid (gunakan C/F/R/K)")
	}
}
