# **IDOR Vulnerability in Web3 (Go + Fiber)**

## **Description**

This lab demonstrates an **Insecure Direct Object Reference (IDOR)** vulnerability in a simulated Web3 environment. The system mimics an educational platform where users earn NFTs for completing courses.

A vulnerable API endpoint allows users to fetch NFT data by specifying the NFT ID and their user ID—without validating ownership. This creates a scenario where attackers can access NFTs they do not own.

For a detailed demonstration of this lab, check out the [Lab Demonstration Video Playlist](https://www.youtube.com/watch?v=hmD7bBqwzmc&list=PLVW6r8PSLZe7MU2yoHFxRZVtPHn2xxmfX&index=4).

## **Example Scenario: NFT Reward System**

In this example:

- Users are rewarded with NFTs for completing courses.
- NFTs are stored with information about the `OwnerID` and `OwnerName`.
- Users can retrieve their NFTs via API.

> If access controls are missing or improperly enforced, a malicious user can simply change the `nftid` in a request and retrieve NFTs belonging to other users—this is an IDOR vulnerability.

---

## **Endpoints Overview**

There are two endpoints provided:

### 🔴 `/nfts/vuln/get` (Vulnerable)

This endpoint does **not** check if the user requesting the NFT is its rightful owner. It simply returns the NFT object based on the given ID.

### 🟢 `/nfts/safe/get` (Secure)

This endpoint **verifies** that the user making the request matches the `OwnerID` of the NFT.

---

## **Setup Instructions**

### **Step 1: Clone the Repository**

Start by cloning the repository to your local machine:

```bash
git clone https://github.com/cetinboran/earn_w2w3vuln.git
cd earn_w2w3vuln/Lab1
```

This will download all necessary files for the lab.

---

### **Step 2: Install Go Modules**

Make sure all Go dependencies are installed by running:

```bash
go mod tidy
```

This will fetch and install the required Go packages listed in `go.mod`.

---

### **Step 3: Run the Application Locally**

You can run the lab locally by executing:

```bash
go run main.go
```

The server will start running on `http://localhost:3000`.

> **Note**: If you face any issues while running locally, ensure Go is installed correctly and the necessary dependencies are in place.

---

### **Step 4: Running the Application in Docker (Optional)**

If you prefer to run the application in a Docker container, follow these steps:

1. **Create a Dockerfile**

   If not already present, create a `Dockerfile` for the application:

   ```Dockerfile
   FROM golang:1.18-alpine AS builder
   WORKDIR /app
   COPY . .
   RUN go mod tidy
   RUN go build -o app .

   FROM alpine:latest
   WORKDIR /root/
   COPY --from=builder /app/app .
   EXPOSE 3000
   CMD ["./app"]
   ```

2. **Build the Docker Image**

   In the project directory, build the Docker image using:

   ```bash
   docker build -t web3-vuln-lab .
   ```

3. **Run the Docker Container**

   Once the image is built, run the container:

   ```bash
   docker run -p 3000:3000 web3-vuln-lab
   ```

   This will start the server inside the Docker container and bind it to `http://localhost:3000`.

---

### **Step 5: Access the Application**

Once the server is running, you can access the lab's endpoints via `http://localhost:3000`. Use the following tools to interact with the API:

- **curl**: Use `curl` commands to interact with the vulnerable and secure endpoints.
- **Postman**: Alternatively, you can use tools like Postman to test the API by sending requests to the endpoints.

---

### **Step 6: Verify Deployment**

To ensure the application is deployed correctly, test both vulnerable and secure endpoints:

1. **Vulnerable Endpoint**:

   ```bash
   curl -X POST http://localhost:3000/nfts/vuln/get \
     -H "Content-Type: application/json" \
     -d '{"nftid":2,"userID":1}'
   ```

   This should return NFT data for `nftid:2` even though it belongs to a different user.

2. **Secure Endpoint**:

   ```bash
   curl -X POST http://localhost:3000/nfts/safe/get \
     -H "Content-Type: application/json" \
     -d '{"nftid":1,"userID":1}'
   ```

   This request should only succeed if `userID: 1` owns NFT `1`.

---

## **Using the API**

### 🔴 Vulnerable Endpoint Exploit Example

```bash
curl -X POST http://localhost:3000/nfts/vuln/get -H "Content-Type: application/json" -d "{\"nftid\":3,\"userID\":2}"
```

Even though NFT with ID 2 belongs to `userID: 2`, this request from `userID: 1` will still return the data — demonstrating the IDOR vulnerability.

---

### 🟢 Secure Endpoint Example

```bash
curl -X POST http://localhost:3000/nfts/safe/get -H "Content-Type: application/json" -d "{\"nftid\":1,\"userID\":1}"
```

This request will succeed because `userID: 1` owns NFT 1.

Trying to access an NFT not owned by the current user:

```bash
curl -X POST http://localhost:3000/nfts/safe/get -H "Content-Type: application/json" -d "{\"nftid\":3,\"userID\":2}"
```

Will return:

```json
IDOR Detected
```

---

## **Mitigation Strategies**

### ✅ Authorization Checks

Always validate whether the authenticated user is allowed to access the resource.

### ✅ Identity Verification

Use wallet addresses, JWT tokens, or signatures to reliably identify users in Web3 apps.

### ✅ Access Control Lists (ACLs)

Map users to resources securely and enforce strict access policies.

---

## **Code Overview**

The application uses:

- 🌀 [Fiber](https://gofiber.io) — A fast, lightweight web framework for Go, ideal for building performant REST APIs.
- 🎓 NFT-like objects — Simulate course rewards as NFTs, each containing an `nftid`, `OwnerID`, `OwnerName`, and other metadata like `CourseName`.
- 🔐 Simulated identity — For simplicity, user authentication is simulated by a hardcoded user identity (`currentUser`) with the following details:

  ```go
  currentUser := User{
    UserID:   1,
    UserName: "cetinboran",
  }
  ```

  In this case:

  - **UserID** is set to `1`, representing the user.
  - **UserName** is set to `"cetinboran"`, a simulated username.

  The `currentUser` is used to simulate a logged-in user when performing requests to the secure endpoint (`/nfts/safe/get`). In a real-world Web3 application, the identity could be verified via wallet addresses, JWT tokens, or other mechanisms.

---

## **Conclusion**

IDOR vulnerabilities can have **severe consequences in Web3** environments, including unauthorized NFT access and loss of digital assets. Proper access control is essential to protect decentralized applications and their users.

---

## **References**

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Solana Devnet Docs](https://docs.solana.com/clusters#devnet)
- [NFT Minting Example on GitHub](https://github.com/metaplex-foundation/js-examples)
