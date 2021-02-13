# Hello World
Contoh source code

```go
package main // <nama-package> 

// digunakan untuk import library lain
import "fmt"
/*
  Contoh penggunaan lain
  import (
    "fmt"
    "strings
  )
*/

func main() {
  fmt.Println("Hello world")

  // variable
  var hello string // atau "hello := new(string)" yang berupa pointer, gunakan *hello untuk menampilkan nilai
  hello = "hello"
  world := "world"
  const whitespace = " "
  fmt.Printf("%s %s, said Mr. Robot.", hello, world)
  fmt.Println(hello + whitespace + world)
}
```

## Reference
1. [Hello World](https://dasarpemrogramangolang.novalagung.com/A-hello-world.html)
2. [Comment](https://dasarpemrogramangolang.novalagung.com/A-komentar.html)