# OpenTelemetry & Jaeger Tracing

> Work in progress..

## Installation

### Setup
* Spin up Jaeger and Mysql containers -> `make m1-up` or `make intel-up`
* Run API server -> `make run`

### Test
* API -> http://localhost:8080/api/v1/users
* Jaeger UI -> http://localhost:16686
* MySQL connection
![Alt text](docs/mysql-connection.png?raw=true "MySQL Connection")
