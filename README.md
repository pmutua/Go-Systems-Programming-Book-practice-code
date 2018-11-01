## Set up instructions

Install the following dependencies

```bash
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
```

#### Api structure

```md
├───app
│ ├───handler
│ |───models
| └───app.go
|───config
└───main.go

main.go - just pulls the config and bootstraps the api,
config - contains the configurations,
handlers - contains all the route handlers,
models - contains models to carry the data,
app.go - assembles the route with handlers
```
