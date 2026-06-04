package main

import "fmt"
type Tugas struct {
	ID string
	NamaTugas string
	Kategori string
	Deskripsi string
	Kesulitan int
	Durasi int
	Selesai bool
}

var dataTugas []Tugas

func tambahTugas() {
	var t Tugas

	fmt.Print("Masukkan ID Tugas   : ")
	fmt.Scan(&t.ID)

	fmt.Print("Masukkan Nama Tugas : ")
	fmt.Scan(&t.NamaTugas)

	fmt.Print("Masukkan Kategori Ruangan : ")
	fmt.Scan(&t.Kategori)

	fmt.Print("Masukkan Deskripsi Pekerjaan : ")
	fmt.Scan(&t.Deskripsi)

	fmt.Println("Skala Kesulitan:")
	fmt.Println("1. Mudah")
	fmt.Println("2. Sedang")
	fmt.Println("3. Sulit")
	fmt.Print("Pilih skala kesulitan (1-3): ")
	fmt.Scan(&t.Kesulitan)

	fmt.Print("Estimasi Durasi Pengerjaan (menit): ")
	fmt.Scan(&t.Durasi)

	t.Selesai = false

	dataTugas = append(dataTugas, t)

	fmt.Println("Data tugas berhasil ditambahkan!")
}

func tampilTugas() {
	fmt.Println("\n===== DAFTAR TUGAS RUMAH TANGGA =====")

	if len(dataTugas) == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < len(dataTugas); i++ {
		labelKesulitan := ""
		if dataTugas[i].Kesulitan == 1 {
			labelKesulitan = "Mudah"
		} else if dataTugas[i].Kesulitan == 2 {
			labelKesulitan = "Sedang"
		} else {
			labelKesulitan = "Sulit"
		}

		statusSelesai := "Belum Selesai"
		if dataTugas[i].Selesai {
			statusSelesai = "Selesai"
		}

		fmt.Println("ID Tugas   :", dataTugas[i].ID)
		fmt.Println("Nama Tugas :", dataTugas[i].NamaTugas)
		fmt.Println("Kategori   :", dataTugas[i].Kategori)
		fmt.Println("Deskripsi  :", dataTugas[i].Deskripsi)
		fmt.Println("Kesulitan  :", labelKesulitan)
		fmt.Println("Durasi     :", dataTugas[i].Durasi, "menit")
		fmt.Println("Status     :", statusSelesai)
		fmt.Println("---------------------------")
	}
}

func ubahTugas() {
	var id string

	fmt.Print("Masukkan ID Tugas yang ingin diubah: ")
	fmt.Scan(&id)

	for i := 0; i < len(dataTugas); i++ {
		if dataTugas[i].ID == id {

			fmt.Print("Nama Tugas baru : ")
			fmt.Scan(&dataTugas[i].NamaTugas)

			fmt.Print("Kategori baru   : ")
			fmt.Scan(&dataTugas[i].Kategori)

			fmt.Print("Deskripsi baru  : ")
			fmt.Scan(&dataTugas[i].Deskripsi)

			fmt.Println("Skala Kesulitan:")
			fmt.Println("1. Mudah")
			fmt.Println("2. Sedang")
			fmt.Println("3. Sulit")
			fmt.Print("Kesulitan baru (1-3): ")
			fmt.Scan(&dataTugas[i].Kesulitan)

			fmt.Print("Durasi baru (menit) : ")
			fmt.Scan(&dataTugas[i].Durasi)

			fmt.Println("Data berhasil diubah!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapusTugas() {
	var id string

	fmt.Print("Masukkan ID Tugas yang ingin dihapus: ")
	fmt.Scan(&id)

	for i := 0; i < len(dataTugas); i++ {
		if dataTugas[i].ID == id {

			dataTugas = append(dataTugas[:i], dataTugas[i+1:]...)

			fmt.Println("Data berhasil dihapus!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func tandaiSelesai() {
	var id string

	fmt.Print("Masukkan ID Tugas yang sudah selesai: ")
	fmt.Scan(&id)

	for i := 0; i < len(dataTugas); i++ {
		if dataTugas[i].ID == id {
			dataTugas[i].Selesai = true
			fmt.Println("Tugas berhasil ditandai selesai!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func sequentialSearch() {
	var pilih int
	var keyword string

	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama Pekerjaan")
	fmt.Println("2. Kategori Ruangan")
	fmt.Print("Pilih: ")

	fmt.Scan(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan nama pekerjaan yang dicari: ")
		fmt.Scan(&keyword)

		ditemukan := false
		for i := 0; i < len(dataTugas); i++ {
			if dataTugas[i].NamaTugas == keyword {
				fmt.Println("Data ditemukan!")
				fmt.Println("ID       :", dataTugas[i].ID)
				fmt.Println("Kategori :", dataTugas[i].Kategori)
				fmt.Println("Durasi   :", dataTugas[i].Durasi, "menit")
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Data tidak ditemukan")
		}

	} else if pilih == 2 {
		fmt.Print("Masukkan kategori ruangan yang dicari: ")
		fmt.Scan(&keyword)

		ditemukan := false
		for i := 0; i < len(dataTugas); i++ {
			if dataTugas[i].Kategori == keyword {
				fmt.Println("Data ditemukan!")
				fmt.Println("ID         :", dataTugas[i].ID)
				fmt.Println("Nama Tugas :", dataTugas[i].NamaTugas)
				fmt.Println("Durasi     :", dataTugas[i].Durasi, "menit")
				ditemukan = true
			}
		}

		if !ditemukan {
			fmt.Println("Data tidak ditemukan")
		}

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch() {
	var cari string

	for i := 0; i < len(dataTugas)-1; i++ {
		min := i
		for j := i + 1; j < len(dataTugas); j++ {
			if dataTugas[j].NamaTugas < dataTugas[min].NamaTugas {
				min = j
			}
		}
		dataTugas[i], dataTugas[min] = dataTugas[min], dataTugas[i]
	}

	fmt.Print("Masukkan nama pekerjaan yang dicari: ")
	fmt.Scan(&cari)

	kiri := 0
	kanan := len(dataTugas) - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if dataTugas[tengah].NamaTugas == cari {
			fmt.Println("Data ditemukan!")
			fmt.Println("ID         :", dataTugas[tengah].ID)
			fmt.Println("Nama Tugas :", dataTugas[tengah].NamaTugas)
			fmt.Println("Kategori   :", dataTugas[tengah].Kategori)
			fmt.Println("Durasi     :", dataTugas[tengah].Durasi, "menit")
			return

		} else if cari < dataTugas[tengah].NamaTugas {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func insertionSortDurasi() {
	for i := 1; i < len(dataTugas); i++ {
		temp := dataTugas[i]
		j := i - 1

		for j >= 0 && dataTugas[j].Durasi > temp.Durasi {
			dataTugas[j+1] = dataTugas[j]
			j--
		}

		dataTugas[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan estimasi durasi (ascending)")
}

func selectionSortKesulitan() {
	for i := 0; i < len(dataTugas)-1; i++ {
		max := i

		for j := i + 1; j < len(dataTugas); j++ {
			if dataTugas[j].Kesulitan > dataTugas[max].Kesulitan {
				max = j
			}
		}

		dataTugas[i], dataTugas[max] = dataTugas[max], dataTugas[i]
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tingkat kesulitan (descending)")
}

func statistik() {
	if len(dataTugas) == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	totalSelesai := 0
	totalDurasi := 0

	for i := 0; i < len(dataTugas); i++ {
		if dataTugas[i].Selesai {
			totalSelesai++
			totalDurasi += dataTugas[i].Durasi
		}
	}

	fmt.Println("\n===== STATISTIK TASKMATE =====")
	fmt.Println("Total Tugas           :", len(dataTugas))
	fmt.Println("Tugas Selesai         :", totalSelesai)
	fmt.Println("Tugas Belum Selesai   :", len(dataTugas)-totalSelesai)

	if totalSelesai > 0 {
		rataRata := float64(totalDurasi) / float64(totalSelesai)
		fmt.Printf("Rata-rata Waktu Kerja : %.2f menit\n", rataRata)
	} else {
		fmt.Println("Belum ada tugas yang selesai")
	}
}

func main() {
	var pilih int

	for {
		fmt.Println("\n===== MENU TASKMATE =====")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Tampilkan Tugas")
		fmt.Println("3. Ubah Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Tandai Tugas Selesai")
		fmt.Println("6. Sequential Search")
		fmt.Println("7. Binary Search")
		fmt.Println("8. Insertion Sort (Durasi)")
		fmt.Println("9. Selection Sort (Kesulitan)")
		fmt.Println("10. Statistik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahTugas()
		} else if pilih == 2 {
			tampilTugas()
		} else if pilih == 3 {
			ubahTugas()
		} else if pilih == 4 {
			hapusTugas()
		} else if pilih == 5 {
			tandaiSelesai()
		} else if pilih == 6 {
			sequentialSearch()
		} else if pilih == 7 {
			binarySearch()
		} else if pilih == 8 {
			insertionSortDurasi()
		} else if pilih == 9 {
			selectionSortKesulitan()
		} else if pilih == 10 {
			statistik()
		} else if pilih == 0 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}