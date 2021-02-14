# Loop dan Array
Loop untuk mengakses data dari suatu array

## Loop
Ada beberapa cara,

1. Basic Loop

```go
for i:= 0; i < 10; i++ {
  fmt.Println(i)
}
```

2. Seperti while (dengan kondisi)

```go
i := 0
for i < 5 {
  fmt.Println(i)
  i++
}
``` 

3. Seperti while (tanpa kondisi)

```go
i := 0
for {
  fmt.Println(i)

  i++
  if i == 5 {
    break
  }
}
```

4. For - range

```go
data := []int{0, 1, 2, 3, 4}
for index, value := range data {
  fmt.Println(index, value)
}
```

5. Break & Continue
Break => menghentikan perulangan
Continue => maju ke perulangan berikutnya

```go
for i:= 1; i <= 10; i++ {
  if i % 2 == 0 {
    continue
  }

  if i > 8 {
    break
  }

  fmt.Println(i)
}
```

## Array
Deklarasi:
```go
var scores [5]int // memiliki 5 index yang dapat diakses dengan nilai awal default dari int
var nilai [5]int{10, 7, 8, 5, 5} // memberikan nilai awal yang berbeda dari default
var ujian [5]int{
  10,
  7,
  8,
  5,
  5,
}
var buah [...]string{"mangga", "melon", "jeruk"} // jumlah array otomatis terisi
var vector := [2][2]int{{0, 1}, {1, 0}} // multi dimensi
// atau
var vector2 := [2][2]int{[2]int{0, 1}, [2]int{1, 0}}
var sayuran = make([]string, 5)
```

Fungsi untuk array:
```go
len(scores) // mengembalikan panjang dari array
```


## Reference
1. [Loop](https://dasarpemrogramangolang.novalagung.com/A-perulangan.html)
2. [Array](https://dasarpemrogramangolang.novalagung.com/A-array.html)