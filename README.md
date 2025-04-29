# **Web2 Vulnerabilities in Web3 - Labs Repository**

## **Introduction**

Welcome to the **Web2 Vulnerabilities in Web3** repository! This project explores how common security vulnerabilities found in traditional Web2 applications, such as **Insecure Direct Object Reference (IDOR)**, **Race Conditions**, and **Server-Side Request Forgery (SSRF)**, can still be exploited within modern Web3 ecosystems. The aim of this project is to raise awareness about these vulnerabilities and demonstrate how they persist in decentralized systems, impacting blockchain-based applications, smart contracts, and decentralized finance (DeFi) platforms.

As Web3 continues to develop and expand, understanding how Web2 security flaws carry over into decentralized applications (dApps) is vital. Despite the promises of greater security and decentralization, many Web2 vulnerabilities are still prevalent in Web3 environments, posing significant risks to user assets and platform integrity. By studying and exploiting these vulnerabilities, we can better understand the associated risks and develop strategies to mitigate them.

## **Purpose**

This project investigates how well-known Web2 security issues impact Web3 systems and explores the challenges faced by developers when securing decentralized applications. While Web3 technologies promise to disrupt traditional centralized systems with increased transparency, ownership, and security, the complexity of these systems introduces new risks, including the persistence of Web2 vulnerabilities.

Through this repository, we demonstrate practical examples of how Web2 security flaws—such as IDOR, Race Conditions, and SSRF—can be exploited in Web3 environments. The primary objective is to show how these vulnerabilities can compromise decentralized applications (dApps), smart contracts, and blockchain-based systems, and provide insights into mitigating these risks to enhance overall platform security.

By focusing on hands-on labs that replicate real-world exploitation scenarios, this repository offers an interactive learning experience for developers, security researchers, and blockchain enthusiasts interested in understanding Web3 vulnerabilities from a security standpoint.

## **Labs**

This repository contains detailed labs designed to help you understand how Web2 vulnerabilities manifest in Web3 environments. Each lab is organized to walk you through the exploitation process, from identifying the vulnerability to exploiting it, as well as offering recommendations for remediation.

The labs are structured as follows:

### **1. IDOR (Insecure Direct Object Reference)**

- **Lab 1:** **IDOR Vulnerability in NFT Minting**  
  This lab simulates a common Web2 vulnerability—IDOR—within a Web3 context. By manipulating an API request, attackers can steal NFTs that belong to other users due to the lack of proper ownership validation. This demonstrates how improper access control can lead to significant security breaches in decentralized applications.
- **Lab 2:** **Advanced IDOR Exploits in dApps**  
  This lab explores further IDOR attack vectors in decentralized applications (dApps) and demonstrates additional techniques to exploit improper access control mechanisms in blockchain-based systems.

### **2. Race Condition**

- **Lab 1:** **Race Condition in Web3**  
  Race conditions are notorious in Web2 applications, and they persist in Web3 environments as well. In this lab, we demonstrate how attackers can exploit timing-related vulnerabilities in Web3 systems, causing unintended behavior such as double-spending or unauthorized actions within a blockchain environment.

### **3. SSRF (Server-Side Request Forgery)**

- **Lab 1:** **Exploiting SSRF in Web3 Systems**  
  SSRF attacks allow attackers to manipulate a server into making unintended requests to internal services or external systems. In the context of Web3, this can be particularly dangerous, as it can lead to information disclosure, internal network compromise, or even the triggering of smart contract vulnerabilities.

## **Why Is This Important?**

As the Web3 space rapidly evolves, the integration of decentralized technologies with traditional systems means that Web2 vulnerabilities can still pose significant threats. Many Web3 projects rely on off-chain logic, centralized servers, and third-party services, all of which may be vulnerable to Web2-style attacks.

This repository aims to bridge the gap between Web2 and Web3 security by highlighting that, while Web3 offers advantages like decentralization and greater user control, it does not inherently eliminate the need for rigorous security practices. As blockchain technologies become more prevalent in finance, supply chains, and other critical industries, securing Web3 applications is paramount to safeguarding digital assets, sensitive information, and user trust.

Through this project, we offer a comprehensive understanding of how traditional Web2 vulnerabilities can be exploited in Web3 environments. By participating in these labs, you will learn how to recognize, mitigate, and protect against such vulnerabilities, ultimately contributing to more secure decentralized systems.

## **Repository Structure**

The repository is organized into the following directories, with each folder containing relevant labs and resources:

- **IDOR/**
  - `lab1/`: Demonstrates the IDOR vulnerability in Web3, specifically within an NFT minting dApp.
  - `lab2/`: Further explores IDOR use cases in decentralized applications, covering scenarios where improper access control is common.
- **RaceCondition/**

  - `lab1/`: Demonstrates a race condition vulnerability in Web3 systems and how it can be exploited to compromise the integrity of transactions.

- **SSRF/**

  - `lab1/`: Exploits SSRF vulnerabilities within Web3 environments, showing how attackers can manipulate servers to access internal systems.

- **README.md**: This file, providing general repository information and instructions for getting started.
- **LICENSE**: Details of the open-source license for this project, ensuring its free use and distribution under specific terms.
