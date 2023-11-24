# Product Management App

This is a comprehensive Product Management App implemented in Go, designed to handle various aspects of product creation, modification, and analysis.

## Features

- **Product CRUD Operations:** Create, Read, Update, and Delete product information.
- **Image Analysis:** Asynchronously perform image analysis tasks, including downloading and compressing product images. Results are stored in the database with local file paths.

## Key Components

- **RESTful API:** Handles product operations with a structured and scalable API.
- **Database Integration:** Utilizes a database for persistent storage of product details.
- **Image Analysis Component:** Processes images in the background, ensuring non-blocking user operations.
- **Unit Tests:** Includes a suite of unit tests for robust code verification.

## Getting Started

### Prerequisites

- Go installed on your machine
- Database (Specify the database and provide setup instructions if needed)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/SakthivelMadhu/product-management-app.git
```
2. Navigate to the project directory:

```bash
cd Product_Management_App
```

3. Install dependencies:

```bash
go mod download
```

4. Run the application:

```bash
go run main.go
```

## Dependencies

This project relies on the following libraries, frameworks, and tools:

- **Gorilla Mux:** A powerful URL router and dispatcher for Go.  
  https://github.com/gorilla/mux

- **GORM:** A fantastic ORM library for Golang, simplifying database operations.  
   https://github.com/go-gorm/gorm

- **MySQL Driver for GORM:** The MySQL database driver used in conjunction with GORM.  
    https://github.com/go-sql-driver/mysql






