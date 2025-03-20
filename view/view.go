package view

import (
	"bufio"
	"fmt"
	"golang_belajar/model"
	"golang_belajar/node"
	"os"
)

func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukan ID Pegawai:")
	fmt.Scan(&id)

	fmt.Println("Masukan Nama Pegawai:")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Println("Masukan Jalan:")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Println("Masukan Nomor Rumah:")
	fmt.Scan(&nomer)

	fmt.Println("Masukan Kota:")
	fmt.Scan(&kota)

	fmt.Println("Masukan Nomor Telp:")
	fmt.Scan(&notelp)

	fmt.Println("Masukan Email:")
	fmt.Scan(&email)

	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{jalan, kota, nomer},
		Notelp: notelp,
		Email:  email,
	}
	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("==Pegawai Berhasil Ditambahkan==")
	} else {
		fmt.Println("Pegawai Gagal Ditambahkan")
	}
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke -", i+1, "--- ")
		fmt.Println("ID Pegawai:", emp.ID)
		fmt.Println("Nama Pegawai:", emp.Nama)
		fmt.Println("Alamat Pegawai:", emp.Alamat.Jalan, emp.Alamat.Kota, emp.Alamat.Nomer)
		fmt.Println("Nomor Telp Pegawai:", emp.Notelp)
		fmt.Println("Email Pegawai:", emp.Email)
	}
}

func Update() {
	var id, nomer int
	var nama, kota, notelp, email string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukan ID Pegawai Yang Akan Di Update:")
	fmt.Scan(&id)

	fmt.Println("Masukan Nama Pegawai:")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Println("Masukan Nama Kota:")
	fmt.Scan(&kota)

	fmt.Println("Masukan Nama Jalan:")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Println("Masukan Nomor Rumah:")
	fmt.Scan(&nomer)

	fmt.Println("Masukan Nomor Telp:")
	fmt.Scan(&notelp)

	fmt.Println("Masukan Email:")
	fmt.Scan(&email)

	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{jalan, kota, nomer},
		Notelp: notelp,
		Email:  email,
	}

	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("==Pegawai Berhasil Diupdate==")
	} else {
		fmt.Println("Pegawai Gagal Diupdate")
	}
}
func Delete() {
	var id int
	fmt.Println("Masukan ID Yang Ingin Anda Hapus:")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("==Pegawai Berhasil Dihapus==")
	} else {
		fmt.Println("Pegawai Gagal Dihapus")
	}
}
