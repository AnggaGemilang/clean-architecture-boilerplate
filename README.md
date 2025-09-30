# Instruction (SHOULD BE REMOVED SOON...)
- Adjust all the service name, version, description, and other stuff
- You can check https://gitea.len-iot.id/kementan/be-weather-insight as the reference of go project

# BE Weather Insight

BE Weather Insight is developed to do some potention prediction, and handle request for visualizing weather trend prediction.

## Version

0.1.0

## Maintained by

Kementan BackEnd Development Team
- Email: [leniot@len.co.id](mailto:leniot@len.co.id)

## Table of Contents

- [BE Weather Insight](#be-weather-insight)
    - [Version](#version)
    - [Maintained by](#maintained-by)
    - [Table of Contents](#table-of-contents)
    - [Getting Started](#getting-started)
        - [Prerequisites](#prerequisites)
        - [Installing](#installing)
    - [Scripts](#scripts)
    - [Dependencies](#dependencies)
    - [Dev Dependencies](#dev-dependencies)
    - [Feedback and Contributions](#feedback-and-contributions)

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

Ensure that you have Linux environment and below components installed to build and run the service:
1. Go v.1.24.1
2. VS Code with Go integrated

### Installing

Follow these steps to get a development environment running:

1. Clone the repo:

    ```bash
    git clone https://gitea.len-iot.id/kementan/be-weather-insight.git
    ```

2. Navigate to the project directory:

    ```bash
    cd be-weather-insight
    ```

3. Initiate go modules:

    ```bash
    go mod init [SERVICE_NAME] e.g. be-weather-insight
    ```

4. Install all dependencies:

    ```bash
    go mod tidy
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

The application should now be running and using several ports for:
1. Websocket Server

## Scripts


## Dependencies

Main libraries and frameworks used in this project:

- github.com/gin-gonic/gin v1.10.1 - Framework for HTTP
- github.com/gorilla/websocket v1.5.3 - Library for Websocket
- github.com/jackc/pgx/v5 v5.7.5 - Library for postgresql connection

## Dev Dependencies

Development tools and libraries used:

- SonarLint - Code linting tool

## Feedback and Contributions

Your feedback is highly appreciated! If you find any bugs or have suggestions for improvements, feel free to open an issue or create a pull request.

If you wish to contribute, please follow the existing coding style and ensure all tests pass before submitting your pull request.
