# **IDOR Vulnerability in Web3 (Go + Fiber)**

## **Description**

This lab demonstrates an **Insecure Direct Object Reference (IDOR)** vulnerability within a Web3 scenario. The vulnerability allows attackers to directly reference and manipulate object IDs (such as NFT IDs) to gain unauthorized access to sensitive data. In the context of Web3, this can have serious consequences, particularly when NFTs (Non-Fungible Tokens) or other digital assets are involved.

The vulnerability arises when there is inadequate access control on the server side, allowing users to access or modify resources that they shouldn’t be able to. This is especially critical in decentralized applications (dApps) and Web3 platforms, where the security of digital assets (such as NFTs, tokens, etc.) is paramount.

### **Example Scenario: NFT Reward System**

Imagine an educational platform where users complete courses and earn NFTs as rewards. Each course is associated with a unique NFT, and users can view their NFTs by accessing them through an endpoint that uses the NFT ID.

However, due to a lack of proper access control mechanisms, an attacker can manipulate the ID in the URL to access NFTs that they have not earned. In this case, the attacker could exploit the system by fetching NFTs that belong to other users, essentially claiming rewards that they did not complete.

### **IDOR in Web3**

While IDOR vulnerabilities are commonly found in traditional web applications, their consequences in Web3 can be far more impactful. For instance, in a Web3 system, NFTs or other blockchain-based assets are involved, which are typically tied to a user’s wallet or identity. Exploiting an IDOR vulnerability could lead to unauthorized access to these digital assets, resulting in loss of control, unauthorized ownership transfers, or theft.

By exploiting this vulnerability, an attacker could manipulate requests to gain access to someone else’s NFT, thereby gaining access to digital assets that should not belong to them.

## **Setup Instructions**

To run this vulnerable server and demonstrate the IDOR vulnerability in Web3, follow the instructions below.

### **1. Clone the Repository**

Clone the repository to your local machine:

```bash
git clone https://github.com/cetinboran/earn_w2w3vuln.git
cd earn_w2w3vuln/Lab1
```

### **2. Install Dependencies**

Make sure all dependencies are installed:

```bash
go mod tidy
```

### **3. Run the Vulnerable Server**

Start the server that simulates the IDOR vulnerability:

```bash
go run main.go
```

This will start the server on `localhost:3000`.

### **4. Exploit the IDOR Vulnerability**

In this lab, the vulnerability can be exploited by sending a request to the `/nfts/:id` endpoint and modifying the `id` parameter in the URL.

**Example Exploit:**

If a user wants to view their NFT, the request might look like this:

```bash
curl -X GET http://localhost:3000/nfts/1
```

This would return the details of the NFT associated with ID `1`.

However, if the attacker changes the `id` parameter to access another user’s NFT, they might send a request like this:

```bash
curl -X GET http://localhost:3000/nfts/2
```

If the server does not perform proper access control checks, the attacker will receive the NFT data of another user, despite not being authorized to view it. This is a classic example of an **Insecure Direct Object Reference** (IDOR) vulnerability.

## **Exploiting IDOR in Web3**

In a Web3 environment, NFTs and other digital assets are often used to represent ownership of scarce or valuable items. These items are stored on the blockchain and are tied to specific users. When an attacker exploits an IDOR vulnerability in such a system, they could potentially:

- View NFTs that belong to other users
- Modify the ownership or transfer of NFTs
- Claim NFTs that they haven't earned (e.g., claiming rewards for courses they didn’t complete)

The result is a serious security risk that compromises the integrity of the Web3 application, undermining trust in the system.

## **Mitigation Strategies**

To prevent such vulnerabilities, proper access control mechanisms must be implemented. Specifically, when handling sensitive resources such as NFTs, the server should verify that the requesting user is authorized to access the specific object ID.

Here are a few mitigation strategies:

1. **Authorization Checks:** Before allowing access to an object (like an NFT), verify that the user requesting the data is the owner of that object.
2. **User Identity Verification:** Ensure that the user’s identity is correctly verified, possibly through wallet signatures, JWT tokens, or similar methods.
3. **Access Control Lists (ACLs):** Implement ACLs to control access to resources based on user roles, ensuring that only authorized users can view or modify their data.

By adding these layers of security, you can prevent attackers from exploiting IDOR vulnerabilities and gaining unauthorized access to users’ digital assets.

## **Steps to Reproduce the Vulnerability**

1. Start the vulnerable server by running:

   ```bash
   go run main.go
   ```

2. Try sending a request to view an NFT that belongs to another user:

   ```bash
   curl -X GET http://localhost:3000/nfts/2
   ```

   This should return the details of the NFT that belongs to `user2`, even though the current user is not authorized to view it.

## **Conclusion**

IDOR vulnerabilities pose a significant risk in Web3 applications, where digital assets like NFTs are involved. This lab demonstrates how a simple lack of access control can lead to unauthorized access to valuable resources. By understanding and mitigating IDOR vulnerabilities, developers can ensure that their Web3 applications are secure and that users' digital assets are protected.

## **References**
