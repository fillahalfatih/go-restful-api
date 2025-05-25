# 📚 Go RESTful API - Book Management

A simple RESTful API for managing book data using **Golang**, **GORM**, **MySQL**, and **Gorilla Mux**.

## 🚀 Features

- ✅ Get all books
- ✅ Get a book by ID
- ✅ Create a new book
- ✅ Update a book by ID
- ✅ Delete a book by ID

## 🏗️ Tech Stack

- [Gorilla Mux](https://github.com/gorilla/mux) v1.8.1 – HTTP router and URL matcher  
- [Google UUID](https://github.com/google/uuid) v1.6.0 – Generate and parse UUIDs  
- [GORM](https://gorm.io/) v1.30.0 – ORM library for Golang  
- [GORM MySQL Driver](https://gorm.io/docs/connecting_to_the_database.html#MySQL) v1.5.7 – GORM dialect for MySQL  
- [Go SQL Driver for MySQL](https://github.com/go-sql-driver/mysql) v1.9.2 – Native MySQL driver for Go  
- [Joho Godotenv](https://github.com/joho/godotenv) v1.5.1 – Load environment variables from `.env`  

## ⚙️ Project Structure

```

go-restful-api/
├── config/
│   └── db.go
├── handlers/
│   └── book_handler.go
├── models/
│   └── book.go
├── routes/
│   └── routes.go
├── .env
├── go.mod
├── LICENSE
├── main.go
└── README.md

````

## 🛠️ Setup and Run

### 1. Clone the Repository

```bash
git clone https://github.com/fillahalfatih/go-restful-api.git
cd go-restful-api
````

### 2. Setup `.env`

Create a `.env` file and fill in your database credentials:

```env
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=db_name
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Server

```bash
go run main.go
```

The server will run at:

```
http://localhost:8080
```

## 📡 API Endpoints

| Method | Endpoint        | Description         |
| ------ | --------------- | ------------------- |
| GET    | /api/books      | Get all books       |
| GET    | /api/books/{id} | Get a book by ID    |
| POST   | /api/books      | Add a new book      |
| PUT    | /api/books/{id} | Update a book by ID |
| DELETE | /api/books/{id} | Delete a book by ID |

## 📝 Sample Request JSON

```json
{
    "id": "a07e28fc-4b4f-4bf8-8fd7-72f9bd1235fc",
    "isbn": "448-1234567890",
    "title": "Book One",
    "year": 2025,
    "author_id": 2,
    "author": {
        "id": 2,
        "first_name": "Grant",
        "last_name": "Gustin"
    }
}
```

## 🧾 License

Free to use for learning and personal exploration.
