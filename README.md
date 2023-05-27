# User Authentication

Aplikasi autentikasi dengan Golang Fiber menggunakan JWT (JSON Web Tokens) adalah sebuah aplikasi yang memanfaatkan framework web Golang Fiber untuk menyediakan sistem autentikasi berbasis token. JWT digunakan sebagai mekanisme untuk mengenkripsi dan memvalidasi token autentikasi yang digunakan dalam proses otentikasi.

Secara umum, aplikasi tersebut akan memiliki beberapa fitur utama, antara lain:

1. Registrasi Pengguna: Pengguna dapat membuat akun baru dengan menyediakan informasi seperti nama, alamat email, phone number, dan kata sandi. Informasi ini kemudian akan disimpan dalam database untuk penggunaan selanjutnya.

2. Login Pengguna: Pengguna yang sudah memiliki akun dapat melakukan login dengan menggunakan alamat email dan kata sandi yang valid. Setelah berhasil login, pengguna akan menerima token JWT yang akan digunakan untuk mengotentikasi akses ke fitur-fitur tertentu di dalam aplikasi.

3. Otentikasi dan Otorisasi: Setiap kali pengguna mengakses fitur yang memerlukan otentikasi, aplikasi akan memvalidasi token JWT yang dikirimkan dalam header permintaan (misalnya, Authorization header). Jika token valid, pengguna dianggap terotentikasi dan diberikan akses ke fitur tersebut. Selain itu, aplikasi juga dapat melakukan otorisasi untuk memastikan pengguna hanya dapat mengakses sumber daya yang sesuai dengan peran atau hak akses yang dimilikinya.

4. Penanganan Permintaan dan Respon: Aplikasi menggunakan framework Golang Fiber untuk menangani permintaan dan merespon dengan data yang sesuai. Hal ini melibatkan pengaturan middleware untuk otentikasi JWT, pengiriman token JWT setelah login berhasil, dan menangani permintaan untuk fitur-fitur yang memerlukan otentikasi.

5. Penyimpanan Data Pengguna: Informasi pengguna seperti akun, kata sandi, dan detail profil biasanya akan disimpan dalam database MySql

6. Mendapatkan semua data users: Pengguna yang sudah login dapat mengakses semua data user.
7. Update user: Pengguna dapat mengubah nama, email, phone number, dan password ketika sudah login
8. Delete user: Pengguna dapat menghapus akunnya ketika sudah melakukan login
 


