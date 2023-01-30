package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "HelloWorld(Alqo)",
			request: "Alqo",
		},
		{
			name: "HelloWorld(Mulky)",
			request: "Mulky",
		},
		{
			name: "HelloWorld(Aria)",
			request: "Aria",
		},
		{
			name: "RandaRamadhan",
			request: "Randa Ramadhan",
		},
		{
			name: "CobaCobaCobaCobaCoba",
			request: "Coba Coba Coba Coba Coba",
		},
	}
	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Alqo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Alqo")
		}
	})
	b.Run("Mulky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Mulky")
		}
	})
}

// benchmark function
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Alqo")
	}
}

func BenchmarkHelloWorldMulky(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mulky")
	}
}

// table test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Alqo",
			request:  "Alqo",
			expected: "Hello Alqo",
		},
		{
			name:     "Mulky",
			request:  "Mulky",
			expected: "Hello Mulky",
		},
		{
			name:     "Aria",
			request:  "Aria",
			expected: "Hello Aria",
		},
		{
			name:     "Rivan",
			request:  "Rivan",
			expected: "Hello Rivan",
		},
		{
			name:     "Randa",
			request:  "Randa",
			expected: "Hello Randa",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Main Test
func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	m.Run() // eksekusi semua unit test
	// after
	fmt.Println("Setelah Unit Test")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Unit test tidak bisa jalan di Linux")
	}

	result := HelloWorld("Alqo")
	require.Equal(t, "Hello Alqo", result, "Result must be 'Hello Alqo'")
}

// Require Test
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Alqo")
	require.Equal(t, "Hello Alqo", result, "Result must be 'Hello Alqo'")
	fmt.Println(("TestHelloWorld with Require Done"))
}

// Assertion Test
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Alqo")
	assert.Equal(t, "Hello Alqo", result, "Result must be 'Hello Alqo'")
	fmt.Println(("TestHelloWorld with Assert Done"))
}

// t.Fail() / t.Error()
func TestHelloWorldAlqo(t *testing.T) {
	result := HelloWorld("Alqo")
	if result != "Hello Alqo" {
		// unit test failed
		// panic("Result is not 'Hello Alqo'")
		// t.Fail()
		t.Error("Result must be 'Hello Alqo'")
	}

	fmt.Println("TestHelloWorldAlqo Done")
}

// t.FailNow() / t.Fatal()
func TestHelloWorldMulky(t *testing.T) {
	result := HelloWorld("Mulky")
	if result != "Hello Mulky" {
		// unit test failed
		// t.FailNow()
		t.Fatal("Result must be 'Hello Mulky'")
	}

	fmt.Println("TestHelloWorldAlqo Done")
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Alqo", func(t *testing.T) {
		result := HelloWorld("Alqo")
		require.Equal(t, "Hello Alqo", result, "Result must be 'Hello Alqo'")
	})
	t.Run("Mulky", func(t *testing.T) {
		result := HelloWorld("Mulky")
		require.Equal(t, "Hello Mulky", result, "Result must be 'Hello Mulky'")
	})
}

/**
- pengenalan software testing :
    1. software testing adalah salah satu disiplin ilmu dalam software engineering.
    2. tujuan utama dari software testing adalah memastikan kualitas kode dan aplikasi kita baik (minim bug dan minim error).
    3. ilmu untuk software testing sendiri sangatlah luat, pada materi ini kita hanya akan fokus ke unit testing.

- unit test :
    1. unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah method.
    2. unit test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kadang kode unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unit test
    3. unit test bisa digunakan sebagai cara untuk meningkatkan kualitas kode program kita

- pengenalan testing package :
    1. di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test
    2. berbeda dengan go-lang, di go-lang sudah ada untuk unit test yang disediakan sebuah package khusus bernama testing
    3. selain itu untuk menjalankan unit test, di go-lang juga sudah disediakan perintahnya.
    4. hal ini membuat implementasi unit testing di golang sangat mudah dibanding dengan bahasa pemrograman yang lain
    5. https://golang.org/pkg/testing/

- testing.T :
    1. go-lang menyediakan sebuah struct yang bernama testing.T
    2. struct ini digunakan untuk unit test di go-lang

- testing.M :
    1. testing.M adalah struct yang disediakan go-lang untuk mengatur life cycle testing
    2. materi ini nanti akan kita bahas di chapter Main

- testing.B :
    1. testing.B adalah struct yang disediakan go-lang untuk melakukan benchmarking
    2. di go-lang untuk melaukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan

- aturan file test :
    1. go-lang memiliki aturan cara membuat file khusus untuk unit test
    2. untuk membuat file unit test, kita harus menggunakan akhiran _test
    3. jadi kita misal ingin membuat file hello_world.go, artinya untuk membuat unit testnya, kita harus membuat file hello_world_test.go

- aturan function unit test :
    1. selain aturan nama file, di go-lang juga sudah diatur untuk nama function unit test
    2. nama function untuk unit test harus diawali dengan nama Test
    3. misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama TestHelloWorld
    4. Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawali dengan kata Test
    5. selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value

- menjalankan unit test
    1. untuk menjalankan unit test kita bisa menggunakan perintah :
        go test
    2. jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah : go test -v
    3. dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah : go test -v -run TestNamaFunction

- menjalankan semua unit test :
    1. jika kita ingin menjalankan semua unit test dari top folder modulenya, kita bisa gunakan perintah : go test ./...

- menggagalkan unit test
    1. menggagalkan unit test menggunakan panic bukanlah hal yang bagus
    2. go-lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
    3. terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

- t.Fail() dan t.FailNow()
    1. terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(), Lantas apa bedanya?
    2. Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test, Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
    3. FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

- t.Error(args...) dan t.Fatal(args...)
    1. Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
    2. Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal.
    3. Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
    4. Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

- Assertion
    1. Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
    2. apalagi jika result data yang harus di cek itu banyak
    3. oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan
    4. sayangnya, go-lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini.

- Testify
    1. Salah satu library assertion yang paling populer di Go-Lang adalah Testify
    2. Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
    3. https://github.com/stretchr/testify
    4.  kita bisa menambahkannya di Go module kita :
            go get github.com/stretchr/testify

- Assert vs Require
    1. Testify menyediakan dua package untuk assertion, yaitu assert dan require
    2. saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
    3. sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan

- Skip Test
    1. Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
    2. di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
    3. Untuk membatalkan unit test kita bisa menggunakan function Skip()

- Before dan After Test
    1. Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
    2. jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual  di unit test functionnya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
    3. untungnya di go-lang terdapat fitur yang bernama testing.M
    4. Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test

- testing.M
    1. Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
    2. jika terdapat function TestMain tersebut, maka secara otomatis go-lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
    3. dengan ini kita bisa mengatur Befpre dan After unit test sesuai dengan yang kita mau
    4. Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test

- Sub Test
    1. Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
    2. fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test bahasa pemrograman yang lainnya
    3. untuk membuat sub test, kita bisa menggunakan function Run()

- Menjalankan Hanya Sub Test
    1. Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah : go test -v -run=TestNamaFunction
    2. jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah : go test -v -run TestNamaFunction/NamaSubTest
    3. Atau untuk semua sub test di semua function, kita bisa gunakan perintah : go test -v -run /NamaSubTest

- Table Test
    1. Sebelumnya kita sudah belajar tentang sub test
    2. jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
    3. dan fitur sub test ini, biasa digunakan oleh programmer go-lang untuk membuat test dengan konsep table test
    4. table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan ekspektasi hasil dari unit test
    5. lalu slice tersebut kita iterasi menggunakan sub test
*/
