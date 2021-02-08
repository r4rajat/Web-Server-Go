

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

----
##API Documentation

###Get All Books [GET]
```api
Endpoint: /api/books/view
```

###Get Book with ID [GET]
```api
Endpoint: /api/book/view/{_id}
```

###Create a New Book [POST]
```api
Endpoint: /api/book/create
Body:-
{
    "name": "Book 1",
    "isbn": 67535487,
    "publisher": "Pelican",
    "author":{
        "first_name": "Rajat",
        "last_name": "Gupta"
    }
}
```

###Update Book Parameters [PATCH]
```api
Endpoint: /api/book/update/{_id}
Body: Use only required to update
{
    "name": "Book x",
    "isbn": 11111,
    "publisher": "aaaaa",
    "author":{
        "first_name": "xxxxx",
        "last_name": "yyyyy"
    }
}
```

###Delete Book with ID [DELETE]
```api
Endpoint: /api/book/delete/{_id}
```