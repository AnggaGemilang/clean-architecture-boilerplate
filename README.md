# Instruction (SHOULD BE REMOVED SOON...)
- Adjust all the service name, version, description, and other stuff
- You can check https://github.com/AnggaGemilang/nest-example as the reference of nest project

# BE Weather Insight

BE Weather Insight is developed to do some potention prediction, and handle request for visualize trend prediction.

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
1. NodeJS v20.18.1
2. NPM v10.8.2
3. NestJS Framework v11.1.6
2. VS Code with NodeJS integrated

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

3. Install all dependencies:

    ```bash
    npm install
    ```

4. Run the application:

    ```bash
    npm run start
    ```

The application should now be running and using several ports for:
1. Websocket Server

## Scripts


## Dependencies

Main libraries and frameworks used in this project:

- @nestjs/common v11.0.1 - Built-in package for decorator and basic class in NestJS

## Dev Dependencies

Development tools and libraries used:

- SonarLint - Code linting tool

## Feedback and Contributions

Your feedback is highly appreciated! If you find any bugs or have suggestions for improvements, feel free to open an issue or create a pull request.

If you wish to contribute, please follow the existing coding style and ensure all tests pass before submitting your pull request.