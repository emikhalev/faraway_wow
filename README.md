# Word of Wisdom (test)

## Algorithm Selection

When selecting the algorithm, the main objective was considered, and the following criteria were defined:

- Widely adopted and easy to implement.
- Provides adequate security.
- Allows flexible difficulty adjustment.

Since there were no specific requirements for the algorithm, the PoW algorithm based on SHA-256 was chosen as sufficient for this server, meeting the criteria.

Other algorithms were not selected for the following reasons:

- Specialized for certain tasks or dependent on specific hardware (ASIC, GPU).
- Require high computational resources.
- More complex to implement and would take more time to develop.

## How to Use

There are multiple ways to run the project.

### Running Locally (Without Docker)
To run the server and client locally without Docker:

1. Open a terminal and start the server:
   ```sh
   make run
   ```
2. Open another terminal and run the client:
   ```sh
   make run-client
   ```
   The client will send a request to the server and then exit.

### Running with Docker
To run the server and client using Docker:
   ```sh
   make composer-up
   ```
   In this mode, the client will keep restarting due to `restart: always` in the `docker-compose.yml` configuration. Requests will continue until you manually stop the containers.

### Running Unit Tests
To run unit tests, use the following command:
   ```sh
   make test
   ```

## Project Structure

### Applications
- **`cmd/wow`** – Server application and its Dockerfile
- **`cmd/client`** – Example client and its Dockerfile

### Internal Packages
- **`internal/closer`** – Closer for graceful shutdown and termination signal handling
- **`internal/config`** – Server configuration management
- **`internal/interceptors`** – Server interceptors (Proof of Work is implemented here)
- **`internal/logger`** – Logging package
- **`internal/server`** – TCP server
- **`internal/service`** – WoW service (request handling from the client is implemented here)
- **`internal/tracer`** – Tracing

### Public Packages
- **`pkg/client`** – Go client package for interacting with the server
- **`pkg/utils`** – Utilities for both the server and `pkg/client` (Proof of Work `verify` and `solve` functions are here)


