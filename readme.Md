# Description
A simple API Product with CRUD and JWT auth

# Features
This API developed with Go, Mux, Gorm, JWT, Postgresql, Docker
 
# Tech Used
 ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) !![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
      
# Getting Start:
Before you running the program, make sure you've run this command:
- `go get -u all`
- `database.go with config db`
- `docker compose up -d`

### Run the program
`go run .`

The program will run on http://localhost:8080


### API Route List
| Method | URL | Description |
| ----------- | ----------- | ----------- | 
| POST | localhost:8080/api/auth/register  | Register 
| POST | localhost:8080/api/auth/login  | Register 
| GET | localhost:8080/api/products/  | Get All Product | Add Authorization {{token}}
| POST | localhost:8080/api/products/  | Create Product | Add Authorization {{token}}
| POST | localhost:8080/api/products/{id}  | Update Product | Add Authorization {{token}}
| GET | localhost:8080/api/products/{id}  | Get Detail a Product | Add Authorization {{token}}
| DELETE | localhost:8080/api/products/{id}  | Delete a Product | Add Authorization {{token}}
 
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->
