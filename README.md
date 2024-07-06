
# Golang REST API Project: Notes

This is a simple REST API project built with Go, using the Gorilla Mux router and PostgreSQL as the database.

## Getting Started

To get started with this project, first clone the repository to your local machine:

```
git clone https://github.com/timur-danilchenko/go-nodes.git
```

### Prerequisites

Before you can run this project, you will need to have Go and PostgreSQL installed on your machine.

### Installing

To install the project dependencies, run the following command:

```
make setup
```

This will install all required dependencies and create the necessary database and table.

### Running

To start the server, run the following command:

```
make start
```

This will start the server on port 8080.


### Usage
    
1.  Start the server by running `make start`. This will start the server at `http://localhost:8080`.
    
2.  Explore the REST API using your preferred tool, such as `curl`, `httpie`, or a REST client like `Postman`. The available endpoints are: 

Method|Endpoint|Description 
---|---|---
GET|/notes|Retrieve all notes
GET|/notes/{id}|Retrieve a specific note by ID
POST|/notes|Create a new note
PUT|/notes/{id}|Update an existing note by ID
DELETE|/notes/{id}|Delete an existing note by ID
3.  The request and response payloads for each endpoint are described in the `openapi.yaml` file, which is an OpenAPI specification file. You can use tools like Swagger UI or ReDoc to explore the API documentation visually.

4.  The server can be configured by setting the following environment variables:
   
Variable|Default value|Description
---|---|---
DB_USER|postgres|PostgreSQL user
DB_PASS|postgres|PostgreSQL password
DB_HOST|localhost|PostgreSQL host
DB_PORT|5432|PostgreSQL port
DB_NAME|notes|PostgreSQL database name

5.  Stop the server by pressing `Ctrl+C` or sending a `SIGINT` signal.
