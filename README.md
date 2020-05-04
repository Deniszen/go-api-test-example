# First steps in Go API test example


Requirements
------------
Must be installed in your system: Docker, Go

Quick Start
------------
1. Pull and run docker image
    ```bash 
    docker pull zentreid/zoo:v1.01
    docker run -d -p 8080:8080 --name crud zentreid/zoo:v1.01
    ```
2. Clone project
    ```bash
    git clone git@github.com:Deniszen/go-api-test-example.git
    ```
3. Run test
    ```bash
    go test
    ```