# How to Create a Golang Web API with Fiber, PostgreSQL, and Gorm

**Author**: Cori1109


## Introduction

This guide explains how to build an example web API in Go to create, update, delete book records from a database. At the end of this article, you should be able to create endpoints and understand how routing works with the Fiber framework. You'll also use Gorm for object-relational mapping with PostgreSQL.


## Prerequisites

  1. Have [Golang version 1.1x installed on your machine](https://www.vultr.com/docs/how-to-set-up-a-fiber-server-with-golang-on-ubuntu-21-04)
  
  2. Basic knowledge of Golang.

  3. Basic knowledge of SQL.

  4. Have [PostgreSQL installed](https://www.vultr.com/docs/how-to-install-configure-backup-and-restore-postgresql-on-ubuntu-20-04-lts) on your PC.


## What is Fiber?

Fiber is an Express-inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's "designed to ease things up for fast development with zero memory allocation and performance in mind" according to the Fiber documentation.


## Set up the Project

1. In your terminal, create a Gofiber directory and change to it.
```
$ mkdir Gofiber
$ cd Gofiber
```
2. Initialize the project.
```
$ go mod init github.com/<your GitHub username>/Gofiber
```
3. Generate the required folders.
```
$ mkdir service storage models
```


## Install the Required Libraries


## Create the Database


## Create the Model


## Service


## Set up the Routes


## Set up the Main Function

```
func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal(err)
  }
  config := &storage.Config{
    Host:     os.Getenv("DB_HOST"),
    Port:     os.Getenv("DB_PORT"),
    Password: os.Getenv("DB_PASS"),
    User:     os.Getenv("DB_USER"),
    SSLMode:  os.Getenv("DB_SSLMODE"),
    DBName:   os.Getenv("DB_NAME"),
  }
  db, err := storage.NewConnection(config)
  if err != nil {
    log.Fatal("could not load database")
  }
  err = models.MigrateBooks(db)
  if err != nil {
    log.Fatal("could not migrate db")
  }
  r := &service.Repository{
    DB: db,
  }
  app := fiber.New()
  r.SetupRoutes(app)
  app.Listen(":8080")
}
```

What comes first here is loading the `.env` file and filling the storage config struct with the values from the env file, that you will create the connection and migrate the database. After that comes the filling of the repository with the connected DB and then setting up the routes by passing the initialized Fiber function as the argument value in the `SetupRoute` function. Finally, you will start the app with the `Listen` function, which accepts an argument of the port number desired to run the app.


## Conclusion

Now you should be able to freely use the Fiber framework to get, post, update, and delete requests and use the Gorm library to connect PostgreSQL to the code.