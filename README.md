# Audit Server

## Getting Started

These instructions will get you a copy of the project up and running on 
your local machine for development and testing purposes. See deployment 
for notes on how to deploy the project on a live system.

#### Prerequisites

```
Install Docker, MongoDB
```

## Deployment

### Run 

- In Docker environment:

  1.  Open /app/config/app.toml

  2.  Set active = 2

  3.  Start Docker

  4. ```
     docker-compose up
     ```

- In non-Docker environment:

  1.  Open /app/config/app.toml

  2.  Set active = 0

  3.  Start MongoDB by running 

     ```
     mongodb
     ```

  2.  ​

     ```
     go run server.go
     ```

### Built With 

- [Gin](https://github.com/gin-gonic/gin): a web framework written in Go (Golang) 


## API

In your local machine, all APIs starts with [http://localhost:8082](http://localhost:8082)

- Error Event: /errorevent
- Quote Server: /quoteserver
- Account Transaction: /accounttransaction
- System Event:  /systemevent
- User Command: /usercommand
- Log all events: /log
- Log events by user id: /log/:userId

#### To get log file from docker container

1. Get the id of the docker container by  running 

   ```
   docker ps
   ```

2. ```
   docker exec -it (the id of the container from step 1) /bin/bash
   ```

3. Get the file path

4. ```
   docker cp (the id of the container):(file path from step 3) [destination file path]
   ```

   ​