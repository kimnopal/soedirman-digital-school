package pertemuan_2

import "fmt"

func init() {
	fmt.Println("=== Pertemuan 2 ===")
	data := []int{64, 34, 22, 11, 90}

	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	fmt.Println("Hasil pengurutan:", data)

	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Slice sebelum di append : ", slice)
	slice = append(slice, "f", "g", "h")
	fmt.Println("Slice sesudah di append : ", slice)
	fmt.Println()
}
