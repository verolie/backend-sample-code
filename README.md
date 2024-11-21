# Sample Code
Backend using Golang and PostgreSQL for the database

## Script for start up
To run the project you can run:

### ` go run main.go `

Runs the app in the development mode.\
you can connect to the end point using [http://localhost:8080](http://localhost:8080) in postman or other app

for the database you can set in utils/database.go you have to have postgreSQL database first

## API 
There are several APIs available: delete, get, create, and edit product, as well as login.
For the login endpoint, no Authorization is needed. For the other APIs, Authorization is required.

### API Route
**login**
```
POST http://localhost:8080/login
```
This endpoint is used for logging in a user to get a token and username. Below is an example of the request body:
```
{
  "username": "admin",
  "password": "admin123"
}
```
**Create Product**
```
POST http://localhost:8080/stock-product
```
This endpoint is used to create a product. Below is an example of the request body:
```
{
  "product_name": "test 4",
  "quantity": 10,
  "status": "draft"
}
```
when you create a product the status should be draft or active
**Get Product**
```
GET http://localhost:8080/stock-product?product_id=1&page=1&pageSize=2
```
This endpoint is used to get products. Below is an example of the request body
you can remove product_id

**Delete Product**
```
DELETE http://localhost:8080/stock-product/4
```
This endpoint is used for deleting products. 
when you delete a product when the product have draft it will hard delete and if the product status is active then the status turn into inactive

**Edit Product**
```
PUT http://localhost:8080/stock-product/3
```
This endpoint is used to edit a product. Below is an example of the request body:
```
{
  "product_name": "test 4",
  "quantity": 12,
  "status": "draft"
}
```
when you edit a product the status should be draft or active

