# Online-shop app
 This is an example shop app


## Create database command:
- docker run --name=human-machine-db -e POSTGRES_PASSWORD='*****' -p 5438:5432 -d --rm postgres
## Local build command:
- docker-compose --env-file .env.dev up --build
## Production build command:
- docker-compose up --build
