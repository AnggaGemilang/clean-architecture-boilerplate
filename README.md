# BE JSC

BE JSC is a service for ...

## Version

0.1.0

## Maintained by

JSC BackEnd Development Team
- Email: [leniot@len.co.id](mailto:leniot@len.co.id)

## Table of Contents

- [BE JSC](#be-jsc)
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
1. GCC 11.4.0
2. CMAKE v3.31.0
3. VS Code with CMAKE integrated
4. VCPKG v2024.06.15
5. OpenDDS v3.29.1

### Installing

Follow these steps to get a development environment running:

1. Clone the repo:

    ```bash
    git clone https://gitea.de.id/c2-nextgen/be-jsc.git
    ```

2. Navigate to the project directory:

    ```bash
    cd be-jsc
    ```

3. Install all dependencies using VCPKG:

    ```bash
    ./vcpkg install [dependency]
    ```

4. Create a new file "CMakeLists.txt":

    ```text
    cmake_minimum_required(VERSION 3.31.0)
    project(be-jsc VERSION 0.2.0 LANGUAGES C CXX)

    set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
    set(CMAKE_CXX_STANDARD 17)

    include(CTest)
    enable_testing()

    include_directories(
        "[vcpkg root directory]/installed/x64-linux/include"
        "[opendds root directory]/OpenDDS-3.29.1"
        "[opendds root directory]/OpenDDS-3.29.1/ACE_wrappers"
        "[opendds root directory]/OpenDDS-3.29.1/ACE_wrappers/TAO"
    )
    link_directories(
        "[vcpkg root directory]/installed/x64-linux/lib" 
        "[opendds root directory]/OpenDDS-3.29.1/lib"
        "[opendds root directory]/OpenDDS-3.29.1/ACE_wrappers/lib"
    )

    file(GLOB_RECURSE SOURCES src/*.cpp src/*.cc src/*.h src/*.hpp)
    list(REMOVE_ITEM SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/src/app.cpp)
    add_library(shared-objects OBJECT ${SOURCES})

    add_executable(${PROJECT_NAME} ${CMAKE_CURRENT_SOURCE_DIR}/src/app.cpp $<TARGET_OBJECTS:shared-objects>)

    find_package(boost_asio REQUIRED CONFIG)
    target_link_libraries(${PROJECT_NAME} PRIVATE Boost::asio)
    target_link_libraries(${PROJECT_NAME} PRIVATE uSockets uv z)
    find_package(RapidJSON CONFIG REQUIRED)
    target_link_libraries(${PROJECT_NAME} PRIVATE rapidjson)
    find_package(Boost REQUIRED COMPONENTS thread)
    target_link_libraries(${PROJECT_NAME} PRIVATE Boost::boost Boost::thread)
    target_link_libraries(${PROJECT_NAME} PRIVATE OpenDDS_Dcps TAO_Valuetype TAO ACE)
    find_package(spdlog CONFIG REQUIRED)
    target_link_libraries(${PROJECT_NAME} PRIVATE spdlog::spdlog)
    find_package(gRPC CONFIG REQUIRED)
    target_link_libraries(${PROJECT_NAME} PRIVATE gRPC::gpr gRPC::grpc gRPC::grpc++ gRPC::grpc++_alts)
    find_package(protobuf CONFIG REQUIRED)
    target_link_libraries(${PROJECT_NAME} PRIVATE protobuf::libprotoc protobuf::libprotobuf protobuf::libprotobuf-lite)

    set(CPACK_PROJECT_NAME ${PROJECT_NAME})
    set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
    include(CPack)
    ```

5. Build the project using that "CMakeLists.txt", the executable file now exists in directory "build".

6. Export environment variables from .env file:

    ```bash
    export $(grep -v '^#' .env | xargs)
    ```

7. Run the application:

    ```bash
    cd build
    ./be-jsc
    ```

The application should now be running and using several ports for:
1. TCP Server
2. UDP Server
3. OpenDDS
4. Websocket Server

## Scripts


## Dependencies

Main libraries and frameworks used in this project:

- Boost (v1.83.0) - Set of libraries (asio, algorithm are used in this project)
- Spdlog (v1.14.1) - Fast C++ logging library
- gRPC (v1.51.1) - Framework for Remote Procedure Call (RPC)
- Protobuf (v3.21.12) - Define and generate message structures in proto files
- uWebSockets (v20.62.0) - WebSocket library
- RapidJSON (v1.1.0) - JSON parser and serializer library
- OpenDDS (v3.29.1) - Data Distribution Service (DDS)

## Dev Dependencies

Development tools and libraries used:

- CMAKE - Builder for C++
- SonarLint - Code linting tool

## Feedback and Contributions

Your feedback is highly appreciated! If you find any bugs or have suggestions for improvements, feel free to open an issue or create a pull request.

If you wish to contribute, please follow the existing coding style and ensure all tests pass before submitting your pull request.
