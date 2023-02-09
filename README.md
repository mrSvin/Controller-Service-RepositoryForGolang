<h1 align="center">Controller-server-repository pattern for CRUD using Golang</h1>

use PostgreSQL and pool connections.

## Create table in scheme auth
```
CREATE TABLE auth.accounts (user_id serial PRIMARY KEY,
                            username VARCHAR ( 50 ) UNIQUE NOT NULL,
                            password VARCHAR ( 50 ) NOT NULL,
                            email VARCHAR ( 255 ) UNIQUE NOT NULL,
                            created_on TIMESTAMP NOT NULL,
                            last_login TIMESTAMP
                            )
```
## Example request for test crete user
```
POST: localhost:8080/userCreate
{
    "name": "alex",
    "password": "qwert",
    "email": "sds@sd.sd"
}
```
## Example request for test update user by id
```
PUT: localhost:8080/userUpdate?id=25
{
    "name": "alex",
    "password": "qwert",
    "email": "sds@sd.sd"
}
```
## Example request for test read user by id
```
GET: localhost:8080/userRead?id=25
```
## Example request for test delete user by id
```
DELETE: localhost:8080/userDelete?id=25
```