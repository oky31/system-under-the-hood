# Deep Dive: how the go garbage collector works


Dimana nilai itu di simpan ? go developer tidak perlu tau di mana sebuah nilai itu di simpan di dalam 
sebuah memory computer, karean go memiliki garbage collector.

Kalau bahasa kerenya "automatically recycling memory" itu adalah tugas garbage collector. Kenapa ini penting 
kerena sesunguhnya memory itu terbatas jadi setelah di gunakan memory harus di lepas atau di hapus 


Where Go value live ? 
non pointer value di simpan di dalam local variable yang artinya tidak di manage oleh go garbage collector , go lebih memilih memory di alokasikan berdasarkan scope dari lexical 

jadi varibale local yang lingkupnya lexical scope, tidak di butuhkan di luar fungsinya maka dari itu akan di alokasikan di "stack" . karena di alokasikan di stack memory bisa langsung di bersihkan begitu fungsinya selesai di gunakan

Contoh di alokasikan di stack :

```go
package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func CreateNewUser() User {
	return User{
		Id:    1,
		Name:  "Susilo",
		Email: "susilo@mail.com",
	}
}

func main() {
	user := CreateNewUser()
	fmt.Println(user)
}
```

Cara kamu melakukan testing 
```sh
go run -gcflags="-m" main.go
```


Hasil Testing

```sh
# command-line-arguments
./main.go:11:6: can inline CreateNewUser
./main.go:20:23: inlining call to CreateNewUser
./main.go:21:13: inlining call to fmt.Println
./main.go:21:13: ... argument does not escape
./main.go:21:14: user escapes to heap
{1 Susilo susilo@mail.com}

```


Dari sini Kita bisa pecah langkah-langkah nya sebagai berikut : 

Step by step:
1. can inline CreateNewUser (line 11)
- Fungsi CreateNewUser di-inline — sama seperti sebelumnya.
2. inlining call to CreateNewUser (line 20)
- Pemanggilan di main() di-inline.
3. inlining call to fmt.Println (line 21)
- fmt.Println(user) di-inline.
4. ... argument does not escape (line 21, argumen)
- Variadic argumen ...interface{} itu sendiri gak escape.
5. user escapes to heap (line 21, variabel user)
- Nah ini yang beda. Meskipun CreateNewUser() return value (User, bukan *User), tetap ada escape ke heap.
- Kenapa? Karena fmt.Println(user) menerima user sebagai interface{}. Begitu suatu concrete value dikonversi ke interface{}, compiler gak bisa jamin ukuran konkretnya, jadi value-nya dibungkus (boxing) dan pindah ke heap.
- Kalau di code gak ada fmt.Println, si user akan tetap di stack dan gak kena GC.
6. {1 Susilo susilo@mail.com}
- Output program.



---

Contoh di alokasikan di heap: 

```go
package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func CreateNewUser() *User {
	return &User{
		Id:    1,
		Name:  "Susilo",
		Email: "susilo@mail.com",
	}
}

func main() {
	user := CreateNewUser()
	fmt.Println(user)
}
```


Cara kamu melakukan testing 
```sh
go run -gcflags="-m" main.go
```


Hasil Testing 

```sh
# command-line-arguments
./main.go:11:6: can inline CreateNewUser
./main.go:20:23: inlining call to CreateNewUser
./main.go:21:13: inlining call to fmt.Println
./main.go:12:9: &User{...} escapes to heap
./main.go:20:23: &User{...} escapes to heap
./main.go:21:13: ... argument does not escape
&{1 Susilo susilo@mail.com}
```

Dari sini Kita bisa pecah langkah-langkah nya sebagai berikut : 

1. can inline CreateNewUser (line 11)
- Compiler menilai fungsi CreateNewUser cukup sederhana untuk di-inline — artinya tubuh fungsi bisa langsung ditempel di tempat pemanggilan, tanpa perlu call ke fungsi terpisah.
2. inlining call to CreateNewUser (line 20)
- Pemanggilan CreateNewUser() di main() benar-benar di-inline.
3. inlining call to fmt.Println (line 21)
- Pemanggilan fmt.Println(user) juga di-inline.
4. &User{...} escapes to heap (line 12 & 20)
- Ini yang penting: struct User tidak dialokasikan di stack, tapi escape ke heap.
- Kenapa? Karena fungsi CreateNewUser mengembalikan pointer (*User). Pointer-nya dipakai di luar fungsi (di main), jadi compiler gak bisa jamin nilai itu aman di stack — karena stack frame CreateNewUser udah cleanup setelah fungsi selesai.
- Jadinya, &User{} harus pindah ke heap supaya bisa diakses dari main.
5. ... argument does not escape (line 21)
- Argumen user yang dikirim ke fmt.Println tidak escape. Maksudnya, fmt.Println cuma baca nilainya, gak nyimpen pointer ke heap lebih lanjut.
6. &{1 Susilo susilo@mail.com}
- Output program normal.

Paradoks menarik: Meski CreateNewUser() cuma bikin struct sederhana, return type *User memaksa si struct pindah ke heap — dan itu artinya GC harus turun tangan buat bersihin nanti. Inilah trade-off pake pointer di Go.

# Resources : 

- https://go.dev/doc/gc-guide
- https://internals-for-interns.com/series/understanding-the-go-compiler/
- https://internals-for-interns.com/series/understanding-the-go-runtime/
