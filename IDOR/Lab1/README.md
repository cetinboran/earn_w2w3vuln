# **IDOR Vulnerability in Web3 (Go + Fiber)**

## **Description**

This lab demonstrates an **Insecure Direct Object Reference (IDOR)** vulnerability in a simulated Web3 environment. The system mimics an educational platform where users earn NFTs for completing courses.

A vulnerable API endpoint allows users to fetch NFT data by specifying the NFT ID and their user ID—without validating ownership. This creates a scenario where attackers can access NFTs they do not own.

---

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

### 1. Clone the Repository

```bash
git clone https://github.com/cetinboran/earn_w2w3vuln.git
cd earn_w2w3vuln/Lab1
```

### 2. Install Go Modules

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

Server starts at: `http://localhost:3000`

---

## **Using the API**

### 🔴 Vulnerable Endpoint Exploit Example

```bash
curl -X POST http://localhost:3000/nfts/vuln/get \
  -H "Content-Type: application/json" \
  -d '{"nftid":2,"userID":1}'
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
curl -X POST http://localhost:3000/nfts/vuln/get -H "Content-Type: application/json" -d "{\"nftid\":1,\"userID\":1}"
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

- 🌀 [Fiber](https://gofiber.io) — Web framework for Go
- 🎓 NFT-like objects to simulate course rewards
- 🔐 Simulated identity through a hardcoded `currentUser`

---

## **Conclusion**

IDOR vulnerabilities can have **severe consequences in Web3** environments, including unauthorized NFT access and loss of digital assets. Proper access control is essential to protect decentralized applications and their users.

---

## **References**

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Solana Devnet Docs](https://docs.solana.com/clusters#devnet)
- [NFT Minting Example on GitHub](https://github.com/metaplex-foundation/js-examples)

```

```

```

```
