

# Demonstration of Web Server Using Gorilla/Mux in Golang

---

###Step 1: Install the Requirements for the Project
```shell
cat reqiurements.txt | xargs go get
```

###Step 2: Set Environmental Variables
```shell
------ OPTIONAL ------
export APP_HOST=0.0.0.0
export APP_PORT=5000
export DB_HOST=0.0.0.0
export DB_PORT=27017

------ MANDATORY ------
export DB_NAME="mylibrary"
export DB_COLLECTION="books"
```

###Step 3: Run the Mux Server
```shell
go run app.go
```