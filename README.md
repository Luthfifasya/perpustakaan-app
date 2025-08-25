# perpustakaan-app
-Menggunakan Gin FrameWork
-Menggunakan Gorm sebagai Relational Object
-Menggunakan database Postgres
-Menggunakan Middleware JWT

post localhost:8080/api/users/login
--> untuk login dan dapatkan token jwt
post localhost:8080/api/users/register	
--> untuk register akun	
get localhost:8080/api/users
--> untuk melihat akun yang sudah didaftarkan/yg ada di database	

contoh body - raw - json di postman
{
  "username": "admin",
  "password": "admin"
}

==>localhost:8080/api/categories
Method yang digunakan adalah get. 
Api ini digunakan untuk menampilkan seluruh kategori

==> localhost:8080/api/categories
Method yang digunakan adalah post. 
Api ini digunakan untuk menambahkan kategori.

==> localhost:8080/api/categories/:id
Method yang digunakan adalah get. 
Api ini digunakan untuk menampilkan detail kategori.

==> localhost:8080/api/categories/:id
Method yang digunakan adalah delete. 
Api ini digunakan untuk menghapus kategori.

==> localhost:8080/api/categories/:id/books
Method yang digunakan adalah get. 
Api ini digunakan untuk menampilkan buku yang tersedia berdasarkan ketegori tertentu.

contoh body - raw - json di postman
{
  "name": "Pemrograman",
  "created_by": "admin"
}

==> localhost:8080/api/books
Method yang digunakan adalah get. 
Api ini digunakan untuk menampilkan seluruh buku

==> localhost:8080/api/books
Method yang digunakan adalah post. 
Api ini digunakan untuk menambahkan buku.

==> localhost:8080/api/books/:id
Method yang digunakan adalah get. 
Api ini digunakan untuk menampilkan detail buku..

==> localhost:8080/api/books/:id
Method yang digunakan adalah delete. 
Api ini digunakan untuk menghapus buku.

contoh body - raw - json do postman
{
  "title": "Calistung Dasar",
  "description": "Mempelajari sambil bermain",
  "image_url": "https://example.com/calistung.jpg",
  "release_year": 2009,
  "price": 40000,
  "total_page": 40,
  "category_id": 4,      ---> ini adalah category buku anak yang sudah saya bikin, id nya adalah 4
  "created_by": "admin"
}
