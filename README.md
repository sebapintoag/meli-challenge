# meli-challenge

## Requirements
- Docker
- Docker compose
- Modify your hosts file (/etc/hosts) and add ```127.0.0.1       me.li```

## Setup
1. Clone this repository
2. Open a terminal and go to root directory
3. Create a **.env** file based on **.env.example** and replace the variable values
4. Type ```make dc-build``` to build the containers
5. Type ```make dc-up``` to run containers

## Instructions
- In your browser, go to **http://me.li:40** to access the admin frontend. You can create, find and delete short URLs
- To access a short url, use **http://me.li/<short_url_key>**
- To use the service as an API Rest, see the Postman file located in **api/MeLi Challenge.postman_collection.json**. This service will work with the complete short URL (**http://me.li/<short_url_key>**) or just the short URL key (**<short_url_key>**).