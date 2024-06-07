# To Do List
Project to learn more about go

## How to run this project
- clone the repo
````
git clone https://github.com/cairomedeiros/todo-list.git
````
### Running without docker

- download  <a href="https://www.enterprisedb.com/downloads/postgres-postgresql-downloads">PostgreSQL</a> and install
- create a database called 'tododb'

### Running with docker

- run ````docker-compose up -d````
- i put volume folder to be created at project root but you can change that

### Now a few more steps
  
- run 'go mod tidy' in the root
- run 'go run main.go'
- or you can use the makefile commands
- but first you have to download <a href="https://gnuwin32.sourceforge.net/packages/make.htm">make for windows</a>
- install the package manager chocolatey
````
choco install make
````
- now run
````
make run
````
