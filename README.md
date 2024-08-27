# Basic CRUD API written in Golang using Gin, GORM, and PostgreSQL

Hi my name is Rashid Ali, I recently started learning about Golang and I have decided to document it so that it can be helpful for other people who are getting started as well.

## Features

- RESTful API with CRUD operations.
- Database migrations using GORM.
- Dockerized environments.
- Environment variable support for configuration.

## Prerequisites

Before you start, make sure you have the following installed:

- [Git](https://git-scm.com/)
- [Golang](https://go.dev/)
- [Docker](https://docs.docker.com/get-docker/) (Optional)
- [Docker Compose](https://docs.docker.com/compose/install/) (Optional)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/rashidkalwar/go-crud.git
cd go-crud
```

### 2. Setup environment variables

- Create a .env file in the root directory of the repository and define two variables.
- First is the port the server will listen to eg. `PORT=8080`
- Second is the Database URL eg. `DB_URL="postgres://<username>:<password>@<IpAddress>:<PORT>"`.

> **Note:** If you have PostgreSQL installed on your system you can use that **OR** use Docker to spin up a PostgreSQL database **OR** use a free online PostgreSQL database. I would suggest [Supabase](https://supabase.com/) it has a generous free tier and a ton of other features. Here is a [Youtube](https://www.youtube.com/watch?v=9T6tjTQ4Zo4) video showing you how to do that.

### 3. Download the dependencies

Ensure that you have Go installed. Download the necessary dependencies using the following command:

```bash
go mod tidy
```

### 4. Run the application

To start the application locally, run the following command,
Once the application is running, you can access the API at `http://localhost:8080`

```bash
go run cmd/main.go
```

#### Run using docker container:

To run the application using docker container, first create an image using following command:

```bash
docker build -t image-name .
```

After creating an image, spin up docker a container using this image, make sure to pass in the Environment variables:

```bash
docker run --name container-name -p 8080:8080 -e PORT=8080 -e DB_URL="postgres://<username>:<password>@<IpAddress>:<PORT>" image-name
```

#### Run using docker-compose:

If you have docker and docker-compose installed you don't need to do much inorder to run the application, just run the following command and it will start the application.

```bash
docker-compose up --build
```

### 5. Testing the API

You can use tools like [Postman](https://www.postman.com) or [cURL](https://curl.se/) to test the API endpoints. Below is an example using "cURL":

```bash
curl -X GET http://localhost:8080/todos
```

### Contact

If you have any questions or feedback, feel free to reach out:

- GitHub: [rashidkalwar](https://github.com/rashidkalwar)
