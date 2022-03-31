# Wallet-API


#### Setting Up Project

Go version 1.17
    
    git clone https://github.com/Toflex/Wallet-API.git
    
    cd ./Wallet-API
    
    go get .
    
    go run main.go

#### Setup environment variables

`RedisDB - Redis DB Name`

`DBName - SQL DB Name`

`DBHost - SQL DB Host`

`DBPort - SQL DB Port`

`DBUser - SQL DB Username`

`DBPass - SQL DB Password`

`RedisHost - Redis Host`

`RedisPort - Redis Port`

`RedisPass - Redis Password`

`ServerPort - Server Port`

`ServerHost - Server Host`

`Secret - JWT secret key`


#### API Documentation
`<BASE-URL>/swagger/index.html#/`


#### Usage
1. Endpoints are protected, in other to access endpoint goto, authentication endpoint to generate bearer token. email address and password can be any random string.

`URL: <BASE-URL>/auth

Request:
{
   "email_address": "string",
   "password": "string"
}

Response: 
{
   "token": "string"
}`


2. Pass bearer token into the header of the other endpoints to get a valid response.
Below is a sample curl request to get a wallet balance.

`curl -X GET \
   http://localhost:8000/api/v1/wallets/11/balance \
   -H 'authorization: Bearer <Token>' \
   -H 'content-type: application/json' \
   -d '{
   "amount": 1290.89
   }'`


