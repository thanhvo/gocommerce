# GoCommerce
This project creates a simple ecommerce system to demonstrate the capbilities of Golang to build micro services.  

## Techonologies
The tech stack includes:
* Web Framework: Gin  
* Database: GORM  
   
## Deployment
First, we setup the data servers using docker-compose.   
$\qquad$`cd docker`  
$\qquad$`docker-compose up -d`  
Then, we can build and run the application  
$\qquad$`make`

## API Verification
Please use tools like Postman to verify the application with following APIs:  
* Register user  
`POST http://localhost:8090/api/user/register`  
* Sign in  
`POST http://localhost:8090/api/user/signin`
* List all users  
`GET http://localhost:8090/api/users`
