# user-crud-service
The project provides a simple and efficient REST API for storing and retrieving data using a MongoDB database.
This was made as a test task for the **Data Impact by NielsenIQ**

<br>

### The API supports the following 4 CRUD operations and login:

### Create
- **Request:** POST `/add/users` 
- **Description:** Upload a JSON file via a POST request to `/add/users`. The data from the file is deserialized and concurrently saved to the MongoDB database for each user.
It is possible to upload JSON directly as POST request body.
The password is encrypted using **bcrypt**, and only the hash is stored in the database.
Additionally, a file is generated per user with the user's ID as the filename, containing only the "data" field.

### Read
- **Request:** GET `/users/list`
- **Request:** GET `/user/:id` 
- **Description:** Retrieves a user by their ID or a list of all users.
- To access GET `/user/:id` endpoint it is necessary to login first.

### Update
- **Request:** UPDATE `/user/:id` 
- **Description:** Modify a user's information by their ID. If the "data" field is changed, the associated file is updated accordingly.

### Delete
- **Request:** DELETE `/delete/user/:id`
- **Description:** Removes a user and their associated file by their unique ID

### Login
- **Request:** POST `/login`
- **Description:** The user is able to connect to their profile in order to be able to access all the data assigned to them. 
The login credentials are verified against the id and the stored password hash in the database.


## Launch:

```bash
$ git clone https://github.com/aytero/user-crud-service.git
$ cd user-crud-service
$ docker compose up --build
```

This will create two containers:
* MongoDB on 27017 port
* App on 8080 port

By default, the API will run on http://localhost:8080.


## Configuration
Make sure to configure the MongoDB connection settings before running the API. You can do this by modifying the .config.env file in the config package.
Provide info for MongoDB connection URL and any other relevant configurations.


## Documentation:
The project's documentation and description of JSON schema is located at
```bash
http://localhost:8080/swagger/index.html
```

