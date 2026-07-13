# Deep Dive: how the go garbage collector works


Dimana nilai itu di simpan ? go developer tidak perlu tau di mana sebuah nilai itu di simpan di dalam 
sebuah memory computer, karean go memiliki garbage collector.

Kalau bahasa kerenya "automatically recycling memory" itu adalah tugas garbage collector. Kenapa ini penting 
kerena sesunguhnya memory itu terbatas jadi setelah di gunakan memory harus di lepas atau di hapus 


Where Go value live ? 
non pointer value di simpan di dalam local variable yang artinya tidak di manage oleh go garbage collector , go lebih memilih memory di alokasikan berdasarkan scope dari lexical 

jadi varibale local yang lingkupnya lexical scope , tidak di butuhkan di luar fungsinya maka dari itu akan di alokasikan di "stack" . karena di alokasikan di stack memory bisa langsung di bersihkan begitu fungsinya selesai di gunakan

Contoh di alokasikan di stack : 
```go
package main

import "fmt"

func buatAngka() int {
    x := 42 // Non-pointer, hanya dipakai di dalam fungsi ini
    return x
}

func main() {
    hasil := buatAngka()
    fmt.Println(hasil)
}

```

Contoh di alokasikan di heap 
```go
package main

import "fmt"

func buatPointerAngka() *int {
    x := 42 // Kita mengembalikan POINTER (&x) ke luar fungsi
    return &x
}

func main() {
    ptr := buatPointerAngka()
    fmt.Println(*ptr)
}
```

Cara kamu melakukan testing 
```sh
go run -gcflags="-m" main.go
```


# Resources : 

- https://go.dev/doc/gc-guide
- https://internals-for-interns.com/series/understanding-the-go-compiler/
- https://internals-for-interns.com/series/understanding-the-go-runtime/
