
# Backend API Test 📝🚀

Untuk projek kali ini, projek ini dikerjakan menggunakan Bahasa Pemrograman Golang dengan framework Gin. Version Golang yang digunakan adalah versio go1.22. Pada service kali ini, tidak ada terhubung dengan Database, pure Golang dan Gin.


## Features API 🔑

- Get User Profile (Accounts)
- Get User By UUID (Accounts)
- Login (Accounts)

Untuk Link Postman dapat dilihat atau terlampir sebagai berikut:
https://documenter.getpostman.com/view/25430042/2sAY52bekX

Didalam postman dokumentasi tersebut, sudah terlampir jelas list api apa saja yang terlampir. Guna untuk dilakukan debuggin atau running development.
## Run Locally Without Makefile 🚀

Clone the project

```bash
  git clone https://gitlab.com/devops-test9315545/backend-api-test.git
```

Go to the project directory

```bash
  cd [my-project]
```

Install dependencies

```bash
  go mod tidy
```

Copy file .env.example -> .env

```bash
  cp -rf .env.example .env
```

Start the server

```bash
  go run ./cmd/app/main.go --signal SIGTERM
```


## Support

Apabila mengalami kendala atau ada pertanyaan silahkan email rakajanitraa@pharos.co.id.

