# Basic Hack
variabel, tipe data, dan sebagainya

## fmt.Fprint format
- **%d** => format data numerik non desimal
- **%f** => format data numerik desimal (%.3f => 3 digit desimal)
- **%t** => format data bool

## Tipe Data
Tipe data pada Go (cheatsheet)
- uint => bilangan positif non desimal
- int => bilangan bulat (positif dan negatif, non desimal)

masing-masing uint dan int ada variasi 8 bit (uint8) 0~255, 16 bit 0~65535, 32 bit 0~4294967295, dan 64 bit. sedangkan untuk int mulai dari -128~127, -32768~32767, dst.
- uint => 32/64 bit
- int => 32/64 bit
- byte => uint8
- rune => int32

desimal
- float32
- float64

lainnya
- bool
- string (diapit oleh tanda petik dua => "string!" atau backtick sehingga karakter tidak di-escape \`Saya \n tidak di escape\`)

## Zero Value
null? it's nil!
- string => ""
- bool => false
- int? uint? => 0
- float => 0.0

selain itu, beberapa menggunakan nil sebagai zero value.

## Operator
Operator untuk mengolah dua variable yang berbeda
- tambah, kurang, kali, bagi, modulus
- sama dengan, tidak sama dengan, kurang dari, lebih dari (sama seperti bahasa lain, python)
- &&, ||, ! 

## Switch
Switch pada Go
```go
switch input {
  case 1:
    fmt.Println("Bukan Prima")
  case 2, 3, 5, 7, 11:
    fmt.Println("Prima")
  default:
    fmt.Println("Belum dihitung")
}
```

atau 
```go
switch {
  case input == 1 {
    ...
}
```

gunakan `fallthrough` untuk mengaktifkan 1 case selanjutnya, kondisi apapun.

## Reference
1. [Variable](https://dasarpemrogramangolang.novalagung.com/A-variabel.html)
2. [Tipe Data](https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html)
3. [Operator](https://dasarpemrogramangolang.novalagung.com/A-operator.html)
4. [Kondisi](https://dasarpemrogramangolang.novalagung.com/A-seleksi-kondisi.html)