# BackendX Template

This is a template for creating Go backend projects. It includes various features and integrations commonly used in backend development.

## Features

Here's a breakdown of the features included in this template:

### Routing with go-chi

[go-chi](https://github.com/go-chi/chi) is used for efficient routing and HTTP middleware management. It provides a lightweight, idiomatic, and composable router for building Go HTTP services.

### Monitoring with Prometheus

[Prometheus](https://prometheus.io/) is integrated into the project for gathering metrics and monitoring the system. It scrapes metrics HTTP endpoints exposed by your application, allowing you to track its performance and health.

### Visualization with Grafana

[Grafana](https://grafana.com/) is included for visualization and analysis of Prometheus metrics. Grafana dashboards allow you to visualize metrics in real-time, helping you gain insights into your application's behavior and performance.

### Logging with zerolog

[zerolog](https://github.com/rs/zerolog) is used for structured logging in the project. It provides fast and efficient logging while allowing you to add context to your log messages, making them easier to analyze.

### Log Aggregation with Loki

[Loki](https://grafana.com/oss/loki/) is integrated for log aggregation and analysis. It stores log data in a horizontally scalable manner, allowing you to easily query and visualize logs alongside your metrics in Grafana.

### Environment Validation

The template includes environment variable validation to ensure correct configuration. This helps prevent runtime errors due to missing or invalid environment variables.

## Makefile Commands

The Makefile included in this project provides convenient commands for building, testing, running, and load testing your Go backend application. Here are the available commands:

```makefile
## build: builds the API into an executable file in the build directory
build:
    @go build -o build/api cmd/main.go

## test: runs the defined test suite
test:
    @go test -v ./...

## run: builds the application then runs it
run: build
    @./build/api

## loadtest: uses the vegeta CLI utility to load test the API
loadtest:
    @echo "GET http://localhost:3000/health" | vegeta attack -duration=60s -rate=10000/s | vegeta report

## dos: uses the vegeta CLI utility to perform DoS on the API to test its limits
dos:
    @echo "GET http://localhost:3000/health" | vegeta attack -duration=60s -rate=50000/s | vegeta report

## help: displays help
help:
    @echo " Choose a command:"
    @sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
```

## Loadtesting this API

Load testing is an essential part of ensuring your API can handle the expected amount of traffic. You can use the make loadtest command to simulate a load on your API and measure its performance. The load test sends a constant rate of requests to a specified endpoint over a specified duration.

To perform a load test, run the following command:

```bash
make loadtest
```

This command will simulate a load on the API endpoint http://localhost:3000/health with a request rate of 10,000 requests per second for a duration of 60 seconds. Adjust the endpoint, rate, and duration as needed for your testing scenario.

## Prerequisites

Before running this project with Docker Compose, ensure you have the following prerequisites installed:

- **Docker**:
  is a platform for developing, shipping, and running applications in containers. It provides a consistent environment for your application to run across different systems. Install Docker from the official website.

## Getting Started

Follow these steps to get started with your Go backend project based on this template:

1. **Fork this repository**:

- Fork this repo so that it can be used as an template.

2. **Clone this forked repo**

   ```bash
   git clone https://github.com/your-username/backendx.git
   cd backendx
   ```

````

3. **Change the username**:
   Change the username the `docker-compose.yml` file to yout username
   or if you are on windows change volume path to a suitable path where
   you would like to mount the volumes

4. **Set up environment variables**:
   by creating a `.env.local` file based on `.env.example` and filling in the necessary values.

5. **Prepate Configuration File**:
   Ensure you have the necessary configuration files ready. This includes prometheus-config.yml, loki-config.yml, and promtail-config.yml. These files are mounted into their respective containers to configure Prometheus, Loki, and Promtail

6. **Accessing Services**:
   Accessing Services: Your Go backend service will be accessible at http://localhost:3000. Prometheus, Grafana, Loki, and Promtail will be available at http://localhost:9090, http://localhost:3100, http://localhost:3200, and http://localhost:xxxx respectively.

7. **Customization**:
   Customization: If needed, you can customize the Dockerfile (Dockerfile) and Docker Compose configuration (docker-compose.yml) according to your project requirements.

8. **Build and run the Docker containers**:

```bash
docker-compose up --build
```

## Usage

- Access the API endpoints through the defined routes.
- Monitor metrics in Prometheus and visualize them in Grafana.
- Analyze logs with Loki for better observability.

## Configuration

Ensure to configure the following environment variables:

- `PORT`: The port on which the server will run.
- `ENVIRONMENT`: This corresponds to the environment in which the application is running in (e.g., `TEST`, `DEV`, `PROD`).
- Other configuration variables specific to your application.

## Contributing

Contributions to this template are welcome! Feel free to open issues or pull requests to suggest improvements or add new features.

## License

This project template is licensed under the [MIT License](LICENSE).
````
