# Web2 to Web3 Race Condition Demo

## Project Overview

This project demonstrates how **race condition vulnerabilities** in Web2 applications can impact Web3 systems. It highlights how these vulnerabilities, commonly seen in centralized systems, can also exist in decentralized environments. The demonstration uses **Solana (SOL)** and **Bitcoin (BTC)** tokens to showcase how Web2 issues, such as race conditions, can lead to unintended or excessive actions in Web3 systems.

- **SOL (Solana)**: The SOL token claim process is secured with a **mutex** to prevent multiple requests from being processed simultaneously.
- **BTC (Bitcoin)**: The BTC token claim process lacks mutex protection, making it vulnerable to race conditions. This vulnerability allows an attacker to send multiple simultaneous requests and claim the token multiple times.

The focus of this demo is not on the tokens themselves but on how Web2 vulnerabilities like race conditions can be exploited in Web3 contexts.

For a detailed demonstration of this lab, check out the [Lab Demonstration Video Playlist](https://www.youtube.com/watch?v=hAcSu7oFjr8&list=PLVW6r8PSLZe7MU2yoHFxRZVtPHn2xxmfX&index=1).

## Technology Stack

- **Go (Golang)**: The backend is built with Go, a powerful language known for its concurrency management, making it suitable for handling race conditions.
- **Fiber**: A lightweight web framework for Go, chosen for its speed and simplicity.
- **Mutex**: Used in the SOL token endpoint to prevent race conditions and ensure safe token claiming.
- **JavaScript (Frontend)**: The frontend is built with simple HTML and JavaScript to allow users to interact with the backend and claim tokens.

## Understanding Race Conditions

A **race condition** occurs when two or more processes attempt to modify shared data simultaneously without proper synchronization, leading to unpredictable outcomes. In web applications, race conditions can happen when multiple requests are made to the same endpoint at the same time, and the server fails to manage concurrent access properly.

In Web3 environments, where trust and security are critical, race conditions can result in serious vulnerabilities, such as unauthorized transactions or multiple claims of the same token.

## Web2 to Web3 Scenario: The Attack

This demo showcases a Web2 race condition scenario affecting a Web3 environment:

1. **SOL Token**: The SOL token claim endpoint is secured using a **mutex**, ensuring that only one request is processed at a time. This protects against race conditions and ensures the claim process is handled correctly.
2. **BTC Token**: The BTC token claim endpoint **lacks mutex protection**, creating a race condition vulnerability. Attackers can exploit this by sending simultaneous requests, potentially claiming multiple tokens before the system can prevent them.

The BTC claim vulnerability illustrates how a Web2 race condition can lead to unexpected and malicious behavior in Web3 systems.

## Exploit: `exploit.go`

The **`exploit.go`** file contains a deliberate **race condition exploit** targeting the unprotected BTC claim endpoint. By sending multiple requests simultaneously, the exploit demonstrates how attackers can claim more BTC tokens than allowed by the system.

### How to Use the Exploit

1. Run the **`exploit.go`** file to send multiple requests to the BTC claim endpoint.
2. Due to the race condition, the exploit will allow you to claim multiple BTC tokens by bypassing the expected behavior of the system.

This demonstrates the risk of race conditions if they are not properly mitigated.

## Running the Project

### Prerequisites

- **Go 1.18+** should be installed on your machine.

### Steps to Run the Backend

1. Clone the repository:

   ```bash
   git clone https://github.com/cetinboran/earn_w2w3vuln.git
   cd RaceCondition/Lab1
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

- **`POST /claim/SOL`**: Claims the SOL token. This endpoint is protected by a mutex, preventing race conditions.
- **`POST /claim/BTC`**: Claims the BTC token. This endpoint is **vulnerable to race conditions** due to the lack of mutex protection.
- **`GET /balance`**: Displays the current balances of both tokens (SOL and BTC).
- **`POST /balance/restart`**: Resets the token balances back to their initial values.

## Frontend Interface

The frontend can be hosted with any static file server (e.g., using `Live Server` extension in VSCode or any HTTP server). The frontend allows users to interact with the backend and view their balances.

### Frontend Features

- **Claim SOL and BTC tokens**: Users can claim both tokens using the respective buttons. The SOL token claim is secure (protected by a mutex), while the BTC token claim is vulnerable to race conditions.
- **View current balances**: Displays the current balance of both SOL and BTC tokens.
- **Restart token balances**: Allows users to reset both token balances to their initial values.

## Conclusion

This project highlights how race conditions, a classic Web2 vulnerability, can still affect Web3 systems. By exploiting the race condition in the BTC token endpoint, attackers can bypass the system's controls and claim more tokens than allowed. This serves as a reminder that Web3 systems must take care to secure all critical operations, ensuring that decentralized environments are protected from traditional Web2 vulnerabilities like race conditions.

## Mitigation Recommendations

To prevent such vulnerabilities in Web3 applications:

1. **Use Mutex for Critical Operations**: Ensure that all token claim processes are synchronized to prevent race conditions.
2. **Validate Token Claims**: Implement logic to verify that a token claim request is valid and has not been exploited.
3. **Monitor for Suspicious Activity**: Track requests and flag any patterns that could indicate an attack, such as rapid multiple claims.
4. **Enforce Access Control**: Ensure that only authorized users can make critical changes like claiming tokens.

By securing decentralized applications against race conditions, developers can protect users and maintain the integrity of Web3 systems.
