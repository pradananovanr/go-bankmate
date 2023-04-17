
# BankMate

A simple API between Customer and Merchant to make payments and deposits.

## Requirements

 - [PostgreSQL](https://www.postgresql.org/download/)
 - [Go](https://go.dev/doc/install)

### Library Used for Go
 - Gin (https://github.com/gin-gonic/gin)
 - JWT-Go (https://github.com/dgrijalva/jwt-go)
 - Godotenv (https://github.com/joho/godotenv)
 - PQ (https://github.com/lib/pq)
## Installation
    1. Install PostgreSQL on your machine.
    2. Create a database by running the SQL script in  the db/db.sql file.
    3. Install the necessary Go packages by running go get
    4. Clone this repository and navigate to the bankmate directory.
## Pre-Usage

Before using the API, make sure to set the environment variables in the `.env` file. The following variables are required:

 - **DB_HOST** (PostgreSQL host
 - **DB_PORT** (PostgreSQL port
 - **DB_USER** (PostgreSQL username
 - **DB_PASSWORD** (PostgreSQL password)
 - **DB_NAME** (PostgreSQL database name)
 - **SSL_MODE** (PostgreSQL SSL mode)
 - **SERVER_PORT** (Port on which the server will run)
 - **API_SECRET** (Secret key for JWT authentication)
 - **TOKEN_HOUR_LIFESPAN** (Lifetime of JWT tokens in hours)
## API Reference

### Customer Endpoint

#### Customer registration

```http
  POST /ver1/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Username |
| `password` | `string` | **Required**. Password |
| `email` | `string` | **Required**. Email |
| `phone` | `string` | **Required**. Phone |

#### Customer login

```http
  POST /ver1/login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Username |
| `password` | `string` | **Required**. Password |

Return value :

`message` `string` that contains Bearer Token for Authorization

#### Customer logout

```http
  POST /ver1/logout
```

### Deposit Endpoint

#### New deposit

```http
  POST /ver1/deposit
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `deposit_amount` | `float` | **Required**. Deposit Amount |
| `deposit_description` | `string` | **Required**. Deposit Description |

#### Find deposit by ID

```http
  GET /ver1/deposit/{id}
```

#### Find all deposit history

```http
  GET /ver1/deposit
```

### Payment Endpoint

#### New payment

```http
  POST /ver1/payment
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `payment_code` | `string` | **Required**. Payment Code |
| `payment_merchant` | `string` | **Required**. Payment |
| `payment_amount` | `float` | **Required**. Payment Amount |
| `payment_description` | `string` | **Required**. Payment Description |

#### Find payment history by ID

```http
  GET /ver1/payment/{id}
```

#### Find all payment history

```http
  GET /ver1/payment
```

### Log Endpoint

#### Find all log history

```http
  GET /ver1/log
```
