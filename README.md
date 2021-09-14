# Go Restfull API Laithan

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## Modul

- [Golang JWT](https://github.com/golang-jwt/jwt)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/).
- [Godotenv](https://github.com/joho/godotenv)
- [Driver Mysql](https://github.com/go-sql-driver/mysql)
- [Go UUID](github.com/nu7hatch/gouuid)

## Installation
Use Golang **version 1.17**

```sh
git clone https://github.com/kurnia123/go-restfull-api.git
cd go-restfull-api
go get .
```

Create File **.env** in root project

```sh
echo > .env
```

Add configuration in **.env** file
```
HOST=localhost
PORT=8080

DB_USER=[your username]
DB_HOST=localhost
DB_PASSWORD=[your password]
DB_DATABASE=belajar_restfull_golang
DB_PORT=3306

ACCESS_TOKEN_KEY=1caa40d6db685726f28387396335ff4ab5b66f52ae508478a98906dad1eb2996cbb82ee1adfa47368ab36bc5dc64b588380cc1880ff9adaa1bd30a8d470abcdc
REFRESH_TOKEN_KEY=31b8365e3ba61ea4350b4e2b39354c608c5394b85a0fae3e5210e9296d1f992da141b51fa414650bfe617325b5327d481b1b53fc4dc748c20ee58a28fb0576a9
ACCESS_TOKEN_AGE=1800
```

Run Server

```sh
go run main.go
```

## License

MIT

**Free Software, Hell Yeah!**
