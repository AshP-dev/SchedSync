# SchedSync

SchedSync is a powerful organizational tool designed to help you streamline your life and schedule. By combining flashcard-like priority-based spaced repetition with seamless calendar integration, SchedSync ensures you never miss a chore, task, or appointment.

## Features

- Recipe management
- Recurring appointments
- Task tracking
- Calendar integration
- Spaced repetition system
- Go backend with REST API
- React frontend

## Setup

### Prerequisites

- Go 1.16+
- Node.js 14+
- PostgreSQL

### Local setup

#### Backend Setup

```sh
git clone https://github.com/yourusername/schedsync.git
cd schedsync/backend
go mod download
```

#### Run Server

```sh
go run main.go
```

#### Frontend Setup

```sh
cd ../frontend
npm install
npm start
```

### Dockerised setup

#### Build and run the containers

```sh
docker compose up --build
```

This command will:

Build the backend Dockerfile and frontend Dockerfile
Start the Go backend server on port 3010
Start the React frontend on port 3000
Set up the proper networking between containers

#### Accessing the Application

Frontend: http://localhost:3000
Backend API: http://localhost:3010

#### Development

To rebuild the containers after making changes:

```sh
docker compose down
docker compose up --build
```

To view container logs:

```sh
docker compose logs -f
```

To stop all container:

```sh
docker compose down
```
