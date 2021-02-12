# Program Baru
Hal yang perlu dilakukan dan diperhatikan untuk membuat program baru.
1. go mod
Setelah membuat folder baru, jalankan perintah `go mod init <nama-project>`

2. .editorconfig
Setelah membuat folder baru, buat file `.editorconfig` pada root dir dari project dengan nilai:
```editorconfig
root = true

[*]
insert_final_newline = true
charset = utf-8
trim_trailing_whitespace = true
indent_style = space
indent_size = 2

[{Makefile,go.mod,go.sum,*.go}]
indent_style = tab
indent_size = 8
``` 
## Perintah Go
1. go run
Menjalankan program go untuk development dengan `go run main.go`

2. go test
Menjalankan unit testing pada file dengan suffix `_test.go` dengan menjalankan `go test <nama file *opt-arg>`

3. go build
Menghasilkan executable file atau binary file yang bisa dijalankan menggunakan `./<nama-file>.exe`. Contoh `go build -o <nama-executable-tanpa-exe>`

4. go get
Untuk download package / 3rd library dengan menggunakan `go get github.com/user/repo`. Untuk download package versi terbaru dapat menambahkan flag `-u`.

5. go mod tidy
Validasi depedensi, mendownload depedensi yang belum ter-download.

6. go mod vendor
soon

## Referensi
1. [go mod](https://dasarpemrogramangolang.novalagung.com/A-setup-go-project-dengan-go-modules.html)
2. [editoroconfig](https://dasarpemrogramangolang.novalagung.com/A-instalasi-editor.html)
3. [perintah go](https://dasarpemrogramangolang.novalagung.com/A-go-command.html)
