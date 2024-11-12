# Bonus Service - DevOps

This repository contains two services, `Frontend` and `Backend`.

## First Things

- **Docker**: Make sure Docker is installed and running on your machine. [Download Docker](https://www.docker.com/products/docker-desktop/)
- **Docker Compose**: Make sure Docker compose is installed and running on your machine. [Download Docker Compose](https://docs.docker.com/compose/install/)


## Steps

1. Clone the repository

```bash
git clone https://github.com/Zaidannzzz/bonus-service.git
cd bonus-service
```

2. **Build and Run Containers**: Use `docker-compose` to build and run both services (frontend and backend) together.

```bash
docker-compose up --build
```

3. **Access the Services**:

- **Frontend**: Open your browser and go to [http://localhost:3000](http://localhost:3000)
- **Backend**: Access the backend API at [http://localhost:8080](http://localhost:8080)

4. **Stopping the Containers**: To stop and remove containers, run:

```bash
docker-compose down
```


## Contact
If you have any questions about this project, please send an email to: `zaidan.zulhakim@gmail.com`.