package main

import "fmt"

func main() {
	var mahasiswa = map[string]int{"aldo": 182, "yosep": 178}

	tampil_mahasiswa(mahasiswa)
}

func tampil_mahasiswa(mahasiswa map[string]int) {
	fmt.Println("Aldo\t:", mahasiswa["aldo"], "cm")
	fmt.Println("Yosep\t:", mahasiswa["yosep"], "cm")
}
