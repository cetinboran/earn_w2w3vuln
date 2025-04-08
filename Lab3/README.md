# Web2 to Web3 Race Condition Demo

## Project Description

This project is designed to demonstrate how **race condition vulnerabilities** in Web2 applications can have an impact on Web3 systems. The main objective is to showcase how a **race condition** can be exploited in a decentralized environment, where multiple requests can lead to unintended or excessive actions.

In this project, **two tokens** (SOL and BTC) are used purely as an example to highlight the **race condition vulnerability**. These tokens serve as a practical demonstration of how Web2 vulnerabilities can be leveraged in Web3 scenarios.

- **SOL (Solana)**: This token is secured using a mutex, ensuring that only one request can claim the token at a time.
- **BTC (Bitcoin)**: The BTC token endpoint lacks mutex protection, allowing race conditions to occur and enabling attackers to claim tokens multiple times simultaneously.

While tokens are involved, the core focus of this project is to showcase the **race condition vulnerability** itself, rather than the tokens or blockchain-specific aspects.

## Technology Stack

- **Go (Golang)**: The backend is developed using Go, which is a fast and efficient language for building secure and concurrent applications.
- **Fiber**: A web framework for Go, chosen for its lightweight nature and performance.
- **Mutex**: Used in the SOL token claim process to ensure safe and synchronized access.
- **JavaScript (Frontend)**: The frontend is a simple HTML and JavaScript interface that allows users to interact with the backend to claim tokens and view balances.

## What is a Race Condition?

A **race condition** is a situation in which the outcome of a process is determined by the non-deterministic ordering of events. In the context of computing, it refers to a situation where multiple processes or threads attempt to modify shared resources simultaneously without proper synchronization. This can lead to unexpected behavior, data corruption, or security vulnerabilities.

In web applications, race conditions can occur when multiple requests are made to the same endpoint at the same time, and the server fails to manage access to shared resources properly. If not mitigated, this can allow an attacker to perform actions that they should not be able to perform, such as claiming a token multiple times or manipulating the state of the system.

In Web3 applications, where security and integrity are paramount, race conditions can lead to severe consequences, such as unauthorized transactions or token claims.

## Web2 to Web3 Scenario: The Attack

In this project, we simulate a Web2 race condition scenario that can directly affect a Web3 system. Here's the process:

1. **SOL Token**: This token is protected by a **mutex**, ensuring that only one request can be processed at a time. This makes the claim process for SOL token secure, as it prevents multiple concurrent requests from succeeding.
2. **BTC Token**: In contrast, the BTC token claim endpoint is **not protected by a mutex**. This creates a race condition vulnerability. An attacker can send multiple simultaneous requests to claim the BTC token, and due to the lack of synchronization, the system might allow multiple claims, even though the token should only be claimed once.

The BTC claim vulnerability demonstrates how Web2 race conditions can be exploited in Web3 environments, where decentralization and trustless operations are critical. By manipulating the order in which requests are processed, an attacker can exploit the race condition to claim extra tokens or cause unintended behavior in the system.

## Exploit: `exploit.go`

The **`exploit.go`** file contains a **deliberate exploit** that targets the **BTC token**'s lack of mutex protection. By sending multiple simultaneous requests to the claim endpoint, the exploit demonstrates how an attacker can claim **more BTC tokens** than allowed by the system.

### How to Use the Exploit

1. Run the **`exploit.go`** file, which sends multiple requests to the **BTC claim endpoint** simultaneously.
2. Since the BTC endpoint is vulnerable to race conditions, the exploit will allow you to claim extra tokens by bypassing the expected behavior of the system.

This serves as a demonstration of how race conditions, if left unaddressed, can lead to security vulnerabilities in Web3 environments.

## Running the Project

### Prerequisites

- **Go 1.18+** should be installed on your machine.

### Steps to Run the Backend

1. Clone the repository:

   ```bash
   git clone https://github.com/cetinboran/earn_w2w3vuln.git
   cd Lab3
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

### API Endpoints

- **`POST /claim/SOL`**: Claims the SOL token. This endpoint is secured using a mutex, preventing race conditions.
- **`POST /claim/BTC`**: Claims the BTC token. This endpoint is **vulnerable to race conditions** due to the lack of mutex protection.
- **`GET /balance`**: Displays the current balances of both tokens (SOL and BTC).
- **`POST /balance/restart`**: Resets the token balances back to their initial values.

## Frontend Interface

You can easily run the frontend by hosting it with a live server of your choice. The frontend allows users to interact with the backend to claim tokens and view their current balances.

### Frontend Features

- **Claim SOL and BTC tokens**: The user can trigger a claim for both tokens using the respective buttons. The SOL token is protected by a mutex, while the BTC token is not, making it vulnerable to race conditions.
- **View current balances**: The current token balances are displayed in real-time.
- **Restart the tokens**: The user can reset the token balances by clicking the "Restart Token" button, which will restore the initial balances of both tokens.

## Conclusion

This project serves as a demonstration of how race conditions in Web2 applications can affect decentralized Web3 environments. By exploiting the race condition vulnerability in the BTC token endpoint, attackers can manipulate the system to gain extra tokens. This project highlights the importance of securing Web3 applications and ensuring that critical resources are properly synchronized to prevent such attacks.
