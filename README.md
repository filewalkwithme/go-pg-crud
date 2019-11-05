# go-pg-crud

This is an example of CRUD web application writen in pure Golang. It makes use of Postgres as database.

![01.png](https://github.com/maiconio/go-pg-crud/blob/master/screenshots/01.png)

![02.png](https://github.com/maiconio/go-pg-crud/blob/master/screenshots/02.png)

# How to Use

* Install PostgreSQL on your system. You can skip this step if already installed.
* Clone this repository.
* Execute the `info.sql` file into your PostgreSQL client. This will import sample database and tables that will be used for this example.
* Modify `main.go` file, line 12. You must configure the PostgreSQL database connection.
* Run `go build` to build the executable file.
* Run/execute the generated program via terminal/command line.
* Open your web browser, and navigate to `http://localhost:8080`.
