package node

type Address struct {
	Jalan, Kota string
	Nomer       int
}
type Pegawai struct {
	ID     int
	Nama   string
	Alamat Address
	Notelp string
	Email  string
}
