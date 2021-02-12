# Instalasi Linux
Proses instalasi pada Linux
1. Download .tar.gz
2. Extract file .tar.gz menggunakan `tar -C /usr/internalFolder -xzf go..tar.gz`
3. Simpan ke .bashrc menggunakan `export PATH=$PATH:/usr/internalFolder/go/bin >> ~/.bashrc`

## Menjalankan Go CLI
Jika `go version` menunjukkan bahwa `go` tidak dapat dikenali, jalankan perintah berikut:
```bash
$ source ~/.bashrc
```

## Update Go
1. Hapus seluruh file Go pada `/usr/internalFolder/go`
2. Lakukan proses instalasi kembali

## Referensi
1. [Instalasi Golang](https://dasarpemrogramangolang.novalagung.com/2-instalasi-golang.html)
