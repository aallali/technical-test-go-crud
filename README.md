# technical-test-go-crud

### Description
simple CRUD rest api to manage users table in mysql:
- create user (first name, last name, email and phone)
- update user by Id
- delete user by Id
- fetch a user by Id
- fetch all users in table
- export all users data to excel + upload it to AWS S3 Bucket

### TODO: 
##### (i forgot that it was required to export the data from users table to Excel file, then upload it to AWS S3 bucket 😅)
- [x] Export data from Database (users table) to excel 
    - [x] stream row by row to avoid load all data to RAM
- [x] Upload file created to AWS S3 bucket  
    - [x] create access keys to AWS for our local env
        - `DONT FORGET TO CONFIG THAT IN UR MACHINE BEFORE TESTING THE EXPORT`
    - [x] create the S3 bucket (let's name it `nuitee-test-file-upload`)
    - [x] connect to S3 from golang (thanks GPT)
    - [x] check if file exists in AWS bucket
### Rules:
##### 1-Schema
- first name : string, length 1-30, required
- last name : string, length 1-30, required
- email : string, valid email, required, unique
- phone: string, valid numbers, required
#### 2-Table
the database table is created using this schema (created one if not exists):
```sql
CREATE TABLE IF NOT EXISTS users (
    id INT(20) unsigned AUTO_INCREMENT,
    firstname VARCHAR(20),
    lastname VARCHAR(20),
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(40),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);
```

##### 3-Duplication:
- update/or create a user requires the email to be unique to this user

### Folder structure:
```shell
.
├── go.mod
├── go.sum
├── main.go                 # main declared here
├── src
│   ├── create.go           # create user route
│   ├── delete.go           # delete user by id route
│   ├── read.go             # fetch users routes (all or single by Id)
│   ├── update.go           # update a user row by id
│   ├── export.go           # export data from 'users' table to excel then upload it to AWS
│   └── helper  
│       ├── config.go       # User model declared here
│       ├── db.go           # all config of MySql database
│       ├── lib.go          # helper functions in routes
│       └── validator.go    # validators of input functions
├── users-backup-10-01-2023_19-18-46.xlsx # file exported containing all data of users table
└── test
    └── nuitee.postman_collection.json  # postman request tests (import and test)
```


### Secrets:
- create an `.env` file based on `.env.example` and update those variables
```
mysql_user=root
mysql_pass=root
mysql_host=localhost:3306

S3_bucket_region="your bucket region here"
S3_bucket_name="your bucket name here"
```


### Run:
- in root folder where docker-compose file exists , excute this, if you dont have mysql installed (it will install mysql+phpmyadmin using docker)
`docker-compose up -d`
- go to folder `crud` and excute this command to install necessary packages
`go mod tidy`
- build the app, this will generate an excutable called server ready to run
`go build -o server`
- run it :) 
`./server`
    ```shell
    Connected to the database
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] POST   /users                    --> nuite/crud/src.CreateUser (3 handlers)
    [GIN-debug] PATCH  /users                    --> nuite/crud/src.UpdateUser (3 handlers)
    [GIN-debug] GET    /users                    --> nuite/crud/src.GetUsers (3 handlers)
    [GIN-debug] DELETE /users/:id                --> nuite/crud/src.DeleteUser (3 handlers)
    [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
    Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
    [GIN-debug] Listening and serving HTTP on :1337
    [GIN] 2023/09/27 - 17:29:54 | 201 |   35.846584ms |             ::1 | POST     "/users"
    [GIN] 2023/09/27 - 17:29:58 | 208 |   15.658792ms |             ::1 | PATCH    "/users"
    [GIN] 2023/09/27 - 17:30:02 | 200 |   26.138917ms |             ::1 | GET      "/users"
    [GIN] 2023/09/27 - 17:30:04 | 200 |   23.183708ms |             ::1 | DELETE   "/users/4"
    ```