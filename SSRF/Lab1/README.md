### Introduction

This lab demonstrates a security vulnerability within the Web2-to-Web3 transition process, specifically highlighting how a Web2-based HTTP server can be exploited to interact with a Web3 gRPC server, leading to unintended access or data leakage. The lab aims to provide hands-on experience with identifying and exploiting such vulnerabilities within an actual system.

### Lab Overview

In this lab, you will be interacting with two components:

1. A vulnerable **Web2 HTTP server**, which receives HTTP requests and forwards them to a gRPC server.
2. A **gRPC server** that is part of a Web3 system, which provides sensitive information.

The goal of the lab is to exploit the vulnerability in the Web2 HTTP server to retrieve secret data from the Web3 gRPC server, showcasing how improper handling of requests and vulnerable server-side components can lead to unintended access to sensitive Web3 information.

For a detailed demonstration of this lab, check out the [Lab Demonstration Video Playlist](https://www.youtube.com/watch?v=37Lmy72kRZE&list=PLVW6r8PSLZe7MU2yoHFxRZVtPHn2xxmfX&index=2).

### Lab Objectives

- Understand the Web2-to-Web3 communication model and the security risks involved.
- Identify and exploit security vulnerabilities in the Web2 HTTP server to interact with the Web3 gRPC server.
- Learn about common security misconfigurations and how they can be exploited in a Web3 environment.

### Prerequisites

Before starting the lab, ensure you have the following:

- Basic understanding of Web2 and Web3 systems.
- Familiarity with **gRPC** and **HTTP** protocols.
- Knowledge of **Go** programming language for interacting with the server components.
- Tools:
  - Go programming environment set up on your machine.
  - A web browser or HTTP client (such as `curl` or Postman).
  - Ability to run `docker` containers for local service execution.

### Setup Instructions

1. **Clone the Repository:**

   - Clone the repository to your local machine.
     ```bash
     git clone https://github.com/cetinboran/earn_w2w3vuln.git
     cd Lab2
     ```

2. **Install Dependencies:**

   - Navigate to the Go backend directory and install required dependencies:
     ```bash
     cd backend
     go mod tidy
     ```

3. **Build and Run the Servers:**

   - To run both servers concurrently using a single command, you can follow these steps:

   ```bash
       go run main.go
   ```

4. **Verify the Services:**
   - Make sure both services are running. The Web2 HTTP server should be accessible on `http://localhost:3001`, and the gRPC server should be accessible on `localhost:50051`.

### Exploiting the Vulnerability

1. **Initiate an HTTP Request:**

   - Using a tool like `curl` or Postman, send a POST request to the vulnerable HTTP server with a JSON payload containing the URL of the gRPC server:

     ```bash
     curl -X POST http://localhost:3001/fetch -H "Content-Type: application/json" -d "{\"url\":\"http://localhost:50051\"}"
     ```

2. **Observe the Exploitation:**
   - The Web2 HTTP server will forward the request to the gRPC server, resulting in the secret data being retrieved from the Web3 system. This is an unintended behavior, as the Web2 server was not designed to properly handle or secure this interaction.

### Key Concepts

- **Web2 HTTP Server Vulnerability:**
  - The vulnerability arises from improper handling of incoming HTTP requests, allowing an attacker to interact with a Web3 gRPC service that should not be exposed to the Web2 environment.
- **gRPC Server:**
  - The gRPC server is designed to handle Web3-specific requests. It exposes sensitive data that is intended to be protected from unauthorized access. The vulnerability in the Web2 server allows this data to be fetched.
- **Exploiting Server Misconfigurations:**
  - The lab demonstrates how simple misconfigurations and lack of proper authentication/authorization checks in the Web2 HTTP server can lead to exposing critical Web3 data.

### How to Fix the Vulnerability

To mitigate this issue, several steps can be taken:

1. **Proper Input Validation:**

   - Ensure that the Web2 HTTP server performs proper validation of incoming requests to prevent unauthorized gRPC server calls.

2. **Implement Authentication and Authorization:**

   - Use authentication mechanisms (e.g., OAuth, API keys) to restrict access to sensitive Web3 resources.

3. **Use HTTPS for Communication:**

   - All communication between Web2 and Web3 servers should be encrypted to prevent man-in-the-middle attacks and other vulnerabilities.

4. **Secure gRPC Endpoints:**
   - Ensure that the gRPC endpoints are only accessible to authorized clients. Use authentication tokens or certificates to validate requests.

### Steps to Reproduce the Issue

1. **Start the servers** as outlined in the setup instructions.
2. **Send a request** to the Web2 server using `curl` or Postman with the URL of the gRPC server (`http://localhost:50051`).
3. Observe the response containing secret data from the gRPC server, demonstrating the vulnerability.

### References

- [gRPC Documentation](https://grpc.io/docs/)
- [Go Documentation](https://golang.org/doc/)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)

### Conclusion

This lab provides a hands-on demonstration of a common vulnerability encountered when transitioning from Web2 to Web3 systems. By exploiting a simple misconfiguration, we can access sensitive Web3 data through a vulnerable Web2 server. The lab also highlights the importance of proper validation, authentication, and secure communication in Web3 applications.
