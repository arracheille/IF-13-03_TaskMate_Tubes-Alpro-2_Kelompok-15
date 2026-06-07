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

type arrTugas []Tugas

func tambahTugas(T *arrTugas) {
	var t Tugas

	fmt.Print("Masukkan ID Tugas   : ")
	fmt.Scan(&t.ID)

	fmt.Print("Masukkan Nama Tugas : ")
	fmt.Scan(&t.NamaTugas)

	fmt.Print("Masukkan Kategori Ruangan : ")
	fmt.Scan(&t.Kategori)

	fmt.Print("Masukkan Deskripsi (gunakan underscore untuk spasi) : ")
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

	*T = append(*T, t)

	fmt.Println("Data tugas berhasil ditambahkan!")
}

func tampilTugas(T arrTugas) {
	var i int

	fmt.Println("\n===== DAFTAR TUGAS RUMAH TANGGA =====")

	if len(T) == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i = 0; i < len(T); i += 1 {
		labelKesulitan := "Sulit"
		if T[i].Kesulitan == 1 {
			labelKesulitan = "Mudah"
		} else if T[i].Kesulitan == 2 {
			labelKesulitan = "Sedang"
		}

		statusSelesai := "Belum Selesai"
		if T[i].Selesai {
			statusSelesai = "Selesai"
		}

		fmt.Println("ID Tugas   :", T[i].ID)
		fmt.Println("Nama Tugas :", T[i].NamaTugas)
		fmt.Println("Kategori   :", T[i].Kategori)
		fmt.Println("Deskripsi  :", T[i].Deskripsi)
		fmt.Println("Kesulitan  :", labelKesulitan)
		fmt.Println("Durasi     :", T[i].Durasi, "menit")
		fmt.Println("Status     :", statusSelesai)
		fmt.Println("---------------------------")
	}
}

func ubahTugas(T *arrTugas) {
	var id string

	fmt.Print("Masukkan ID Tugas yang ingin diubah: ")
	fmt.Scan(&id)

	idx := seqSearchID(*T, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Print("Nama Tugas baru : ")
	fmt.Scan(&(*T)[idx].NamaTugas)

	fmt.Print("Kategori baru   : ")
	fmt.Scan(&(*T)[idx].Kategori)

	fmt.Print("Deskripsi baru  : ")
	fmt.Scan(&(*T)[idx].Deskripsi)

	fmt.Println("Skala Kesulitan:")
	fmt.Println("1. Mudah")
	fmt.Println("2. Sedang")
	fmt.Println("3. Sulit")
	fmt.Print("Kesulitan baru (1-3): ")
	fmt.Scan(&(*T)[idx].Kesulitan)

	fmt.Print("Durasi baru (menit) : ")
	fmt.Scan(&(*T)[idx].Durasi)

	fmt.Println("Data berhasil diubah!")
}

func hapusTugas(T *arrTugas) {
	var j int
	var id string

	fmt.Print("Masukkan ID Tugas yang ingin dihapus: ")
	fmt.Scan(&id)

	idx := seqSearchID(*T, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	j = idx
	for j < len(*T)-1 {
		(*T)[j] = (*T)[j+1]
		j = j + 1
	}
	*T = (*T)[:len(*T)-1]

	fmt.Println("Data berhasil dihapus!")
}

func tandaiSelesai(T *arrTugas) {
	var id string

	fmt.Print("Masukkan ID Tugas yang sudah selesai: ")
	fmt.Scan(&id)

	idx := seqSearchID(*T, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	(*T)[idx].Selesai = true
	fmt.Println("Tugas berhasil ditandai selesai!")
}

func seqSearchID(T arrTugas, X string) int {
	var found int
	var j int

	j = 0
	found = -1

	for j < len(T) && found == -1 {
		if T[j].ID == X {
			found = j
		}
		j = j + 1
	}
	return found
}

func sequentialSearch(T arrTugas) {
	var pilih int
	var keyword string
	var i int

	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama Pekerjaan")
	fmt.Println("2. Kategori Ruangan")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan nama pekerjaan yang dicari: ")
		fmt.Scan(&keyword)

		ditemukan := false
		for i = 0; i < len(T); i += 1 {
			if T[i].NamaTugas == keyword {
				fmt.Println("Data ditemukan!")
				fmt.Println("ID       :", T[i].ID)
				fmt.Println("Kategori :", T[i].Kategori)
				fmt.Println("Durasi   :", T[i].Durasi, "menit")
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
		var j int
		for j = 0; j < len(T); j += 1 {
			if T[j].Kategori == keyword {
				fmt.Println("Data ditemukan!")
				fmt.Println("ID         :", T[j].ID)
				fmt.Println("Nama Tugas :", T[j].NamaTugas)
				fmt.Println("Durasi     :", T[j].Durasi, "menit")
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

func binarySearchNama(T arrTugas, X string) int {
	var found int
	var kr int
	var kn int

	found = -1
	kr = 0
	kn = len(T) - 1

	for kr <= kn && found == -1 {
		med := (kr + kn) / 2
		if X < T[med].NamaTugas {
			kn = med - 1
		} else if X > T[med].NamaTugas {
			kr = med + 1
		} else {
			found = med
		}
	}

	return found
}

func binarySearch(T *arrTugas) {
	var cari string
	insertionSortNama(T)

	fmt.Print("Masukkan nama pekerjaan yang dicari: ")
	fmt.Scan(&cari)

	idx := binarySearchNama(*T, cari)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data ditemukan!")
		fmt.Println("ID         :", (*T)[idx].ID)
		fmt.Println("Nama Tugas :", (*T)[idx].NamaTugas)
		fmt.Println("Kategori   :", (*T)[idx].Kategori)
		fmt.Println("Durasi     :", (*T)[idx].Durasi, "menit")
	}
}

func insertionSortDurasi(T *arrTugas) {
	var i int
	var j int
	for i = 1; i <= len(*T)-1; i += 1 {
		j = i
		temp := (*T)[j]
		for j > 0 && temp.Durasi < (*T)[j-1].Durasi {
			(*T)[j] = (*T)[j-1]
			j = j - 1
		}
		(*T)[j] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan estimasi durasi (ascending)")
}

func insertionSortNama(T *arrTugas) {
	var i int
	for i = 1; i <= len(*T)-1; i += 1 {
		var j int = i
		temp := (*T)[j]
		for j > 0 && temp.NamaTugas < (*T)[j-1].NamaTugas {
			(*T)[j] = (*T)[j-1]
			j = j - 1
		}
		(*T)[j] = temp
	}
}

func selectionSortKesulitan(T *arrTugas) {
	var i int
	for i = 1; i <= len(*T)-1; i += 1 {
		idx_max := i - 1
		var j int = i
		for j < len(*T) {
			if (*T)[j].Kesulitan > (*T)[idx_max].Kesulitan {
				idx_max = j
			}
			j = j + 1
		}
		t := (*T)[idx_max]
		(*T)[idx_max] = (*T)[i-1]
		(*T)[i-1] = t
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tingkat kesulitan (descending)")
}

func statistik(T arrTugas) {
	var i int
	var totalSelesai, totalDurasi int

	if len(T) == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	totalSelesai = 0
	totalDurasi = 0

	for i = 0; i < len(T); i += 1 {
		if T[i].Selesai {
			totalSelesai++
			totalDurasi += T[i].Durasi
		}
	}

	fmt.Println("\n===== STATISTIK TASKMATE =====")
	fmt.Println("Total Tugas           :", len(T))
	fmt.Println("Tugas Selesai         :", totalSelesai)
	fmt.Println("Tugas Belum Selesai   :", len(T)-totalSelesai)

	if totalSelesai > 0 {
		rataRata := float64(totalDurasi) / float64(totalSelesai)
		fmt.Printf("Rata-rata Waktu Kerja : %.2f menit\n", rataRata)
	} else {
		fmt.Println("Belum ada tugas yang selesai")
	}
}

func main() {
	var dataTugas arrTugas
	var pilih int

	for {
		fmt.Println("\n===== MENU TASKMATE =====")
		fmt.Println("1.  Tambah Tugas")
		fmt.Println("2.  Tampilkan Tugas")
		fmt.Println("3.  Ubah Tugas")
		fmt.Println("4.  Hapus Tugas")
		fmt.Println("5.  Tandai Tugas Selesai")
		fmt.Println("6.  Sequential Search")
		fmt.Println("7.  Binary Search")
		fmt.Println("8.  Insertion Sort (Durasi)")
		fmt.Println("9.  Selection Sort (Kesulitan)")
		fmt.Println("10. Statistik")
		fmt.Println("0.  Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahTugas(&dataTugas)
		} else if pilih == 2 {
			tampilTugas(dataTugas)
		} else if pilih == 3 {
			ubahTugas(&dataTugas)
		} else if pilih == 4 {
			hapusTugas(&dataTugas)
		} else if pilih == 5 {
			tandaiSelesai(&dataTugas)
		} else if pilih == 6 {
			sequentialSearch(dataTugas)
		} else if pilih == 7 {
			binarySearch(&dataTugas)
		} else if pilih == 8 {
			insertionSortDurasi(&dataTugas)
		} else if pilih == 9 {
			selectionSortKesulitan(&dataTugas)
		} else if pilih == 10 {
			statistik(dataTugas)
		} else if pilih == 0 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}