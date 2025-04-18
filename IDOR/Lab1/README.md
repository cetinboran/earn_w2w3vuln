# **IDOR Vulnerability in Web3 (Go + Fiber)**

## **Description**

This lab demonstrates an **Insecure Direct Object Reference (IDOR)** vulnerability within the context of Web3 applications. The vulnerability allows attackers to manipulate object IDs, such as NFT IDs, to gain unauthorized access to sensitive resources. In decentralized platforms like Web3, this vulnerability can lead to severe consequences, including theft or unauthorized modification of blockchain-based assets such as NFTs.

IDOR typically occurs when a server lacks proper access control mechanisms, allowing users to access or modify resources they should not be able to. In Web3 applications, where ownership and control over digital assets (like NFTs) are fundamental, IDOR vulnerabilities are especially critical as they can directly affect the security and integrity of the blockchain environment.

### **Example Scenario: NFT Reward System**

Consider an educational platform where users earn NFTs as rewards for completing courses. Each course is linked to a unique NFT, and users can access their NFTs via a dedicated endpoint.

However, if access controls are not properly enforced, an attacker could manipulate the object ID (in this case, the NFT ID) in the request URL to access NFTs belonging to other users. This would allow the attacker to claim rewards that they have not earned.

### **IDOR in Web3**

While IDOR vulnerabilities are often seen in traditional Web2 applications, they can be even more dangerous in Web3 systems. This is because Web3 applications often involve blockchain-based assets, which are directly tied to a user’s identity, typically represented by a wallet address. Exploiting an IDOR vulnerability in this context could lead to unauthorized access to NFTs or other valuable assets.

By exploiting such a vulnerability, an attacker can potentially:

- View or retrieve NFTs that belong to other users
- Modify the ownership of NFTs
- Claim NFTs that they haven’t earned (e.g., claiming rewards for courses they haven’t completed)

This presents a severe security risk that can undermine the trust and integrity of a Web3 platform.

## **Setup Instructions**

Follow these steps to run the vulnerable server and demonstrate the IDOR vulnerability in a Web3 environment.

### **1. Clone the Repository**

Start by cloning the repository to your local machine:

```bash
git clone https://github.com/cetinboran/earn_w2w3vuln.git
cd earn_w2w3vuln/Lab1
```

### **2. Install Dependencies**

Make sure to install all the necessary dependencies:

```bash
go mod tidy
```

### **3. Run the Vulnerable Server**

Now, start the server that simulates the IDOR vulnerability:

```bash
go run main.go
```

This will start the server locally on `localhost:3000`.

### **4. Exploit the IDOR Vulnerability**

The vulnerability can be exploited by manipulating the `id` parameter in the URL when making a request to the `/nfts/:id` endpoint.

#### **Example Exploit:**

If a legitimate user wants to view their NFT, the request would look like this:

```bash
curl -X GET http://localhost:3000/nfts/1
```

This fetches the NFT associated with ID `1`.

However, if an attacker modifies the `id` parameter to view someone else’s NFT, they can send a request like:

```bash
curl -X GET http://localhost:3000/nfts/2
```

If the server fails to check whether the user is authorized to access this NFT, it will return the NFT data for `id 2`, which belongs to another user.

This is a classic **IDOR** vulnerability, where an attacker can gain access to resources that they shouldn't have access to simply by manipulating the object ID.

## **Exploiting IDOR in Web3**

In Web3, digital assets like NFTs are linked to user wallets, which provide a decentralized and tamper-proof means of ownership. Exploiting an IDOR vulnerability in a Web3 system can have serious consequences, such as:

- Unauthorized access to NFTs or other digital assets
- Modifying the ownership of NFTs
- Claiming rewards (e.g., NFTs for courses not completed)

Since Web3 applications rely on the integrity of blockchain assets, the exploitation of an IDOR vulnerability could result in significant financial or reputational loss for users and platforms.

## **Mitigation Strategies**

To prevent such vulnerabilities, robust access control mechanisms should be put in place. Here are some strategies to mitigate the risk of IDOR in Web3 applications:

### **1. Authorization Checks**

Always ensure that the user making the request is authorized to access the requested object. For example, before allowing access to an NFT, verify that the user’s wallet address matches the one associated with the NFT ID.

### **2. User Identity Verification**

In Web3 applications, the user’s identity is typically tied to their wallet address. Ensure proper verification of the user’s identity before processing requests. This can be done through wallet signatures, JWT tokens, or other secure methods.

### **3. Implement Access Control Lists (ACLs)**

Use ACLs to enforce access control rules based on user roles and ensure that only authorized users can view or modify their NFTs.

By implementing these security measures, you can greatly reduce the risk of exploitation of IDOR vulnerabilities in Web3 applications.

## **Steps to Reproduce the Vulnerability**

1. Start the vulnerable server by running:

   ```bash
   go run main.go
   ```

2. Attempt to access another user’s NFT by sending a request like:

   ```bash
   curl -X GET http://localhost:3000/nfts/2
   ```

   If the server does not perform proper access checks, you will see the details of the NFT that belongs to `user2`, even though you are not authorized to access it.

## **Conclusion**

This lab highlights the importance of proper access control in Web3 applications. IDOR vulnerabilities in such platforms can lead to unauthorized access to NFTs and other digital assets, undermining the integrity of the system. Developers must be vigilant in implementing security measures to ensure that Web3 platforms are secure and trustworthy for users.

## **References**

- [Solana Devnet Docs](https://docs.solana.com/clusters#devnet)
- [Solana General Docs](https://solana.com/tr/docs)
- [Solana CLI Tool Docs](https://solana.com/tr/docs/intro/installation)
- [NFT Minting Example on GitHub](https://github.com/metaplex-foundation/js-examples)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)

