# user-crud-service

a REST API in Golang, allowing data to be saved in a MongoDB type database. You can take the help of the Gin Gonic router.
This API supports the following 4 CRUD functions:

1. Create
- Request: POST /add/users 
- Description: The POST method data is a file, the file format is provided in the dataset in JSON format. The data is de-serialized and then saved to a MongoDB database concurrently for each user. Entries already inserted are not edited again. The password is encrypted with bcrypt and only the hash is inserted in the base.
  In addition to the insertion in the database, it generates a file per user with the user's id as file name, this file contains only the “data” field

2. Login
- Request: Post /login 
- Description: The user is able to connect to their profile in order to be able to access all the data assigned to them. The user is considered to have an id and a password.

3. Delete
- Request: DELETE /delete/user/:id 
- Description: Deletes the user and the generated file by their ID.

4. Read
- Request: GET /users/list
- Request: GET /user/:id 
- Description: Retrieves a user by their id or a list of all users.

5. Update
- Request: UPDATE /user/:id 
- Description: Modifies a user by their id. If the data field changes the file is modified.


curl -X POST -d @tests/DataSetOne localhost:8080/add/users