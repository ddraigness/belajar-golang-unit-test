package main

func main() {
	
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

- Mock
    1. Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal.
    2. Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing.
    3. Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita mau.
    4. Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

- Testify Mock
    1. Untuk membuat mock object, tidak ada fitur bawaan go-lang, namun ktia bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion.
    2. Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
    3. Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik.
    4. Mari kita buat contoh kasus

contoh kasus
- Aplikasi Query ke Database
    1. Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database.
    2. dimana kita akan buat layer Service sebagai business login, dan layer Repository sebagai jembatan ke database.
    3. agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface
-------

- Benchmark 
    1. selain unit test, go-lang testing package juga mendukung melakukan benchmark
    2. Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
    3. Benchmark di go-lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
    4. ktia tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

- testing.B
    1. testing.B adalah struct yang digunakan untuk melakukan benchmark.
    2. testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dan lain-lain
    3. yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
    4. salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

- cara kerja benchmark
    1. cara kerja benchmark di go-lang sangat sederhana
    2. gimana kita hanya perlu membuat perulangan sejumlah N attribute
    3. nanti secara otomatis go-lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmarknya dalam waktu

- benchmark function
    1. mirip seperti unit test, untuk benchmark pun, di go-lang sudah ditentukan nama functionnya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
    2. Selain itu, harus memiliki parameter (b *testing.B)
    3. dan tidak boleh mengembalikan return value
    4. untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go

- menjalankan benchmark
    1. untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namuun ditambahkan parameter bench : 'go test -v -bench=.'
    2. jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah : 'go test -v -run=NotMathUnitTest -bench=.'
    3. kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah : 'go test -v -run=NotMathUnitTest - bench=BenchmarkTest'
    4. Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah : 'go test -v -bench=. ./...'

- sub benchmark
    1. sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()

- menjalankan hanya sub benchmark
    1. saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
    2. namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah : 'go test -v -bench=BenchmarkNama/NamaSub'

- table benchmark
    1. sama seperti di unit test, programmer go-lang terbiasa membuat table benchmark juga
    2. ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function
*/