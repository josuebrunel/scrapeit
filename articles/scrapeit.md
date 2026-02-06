---
title: "Deep Dive into Scrapeit: A Decentralized Scraping Platform"
date: 2024-05-23
draft: false
tags: ["go", "libp2p", "ethereum", "scraping", "architecture"]
---

# Scrapeit: Decentralized Web Scraping

**Scrapeit** is a Proof of Concept (PoC) platform designed to decentralize web scraping tasks. By distributing the workload across a peer-to-peer (P2P) network, it aims to overcome common challenges faced by centralized scraping infrastructures, such as IP blocking and censorship.

## What does this project solve?

Traditional web scraping relies on centralized servers or proxy pools. Target websites can easily identify and block these known data center IP ranges.

**Scrapeit solves this by:**
1.  **Decentralization**: Distributing requests across a network of client nodes (potentially residential IPs), making it much harder for anti-bot systems to fingerprint and block the scraping traffic.
2.  **Incentivization**: Using Ethereum smart contracts to create a marketplace where anyone can contribute computing power and bandwidth in exchange for rewards.
3.  **Resilience**: Removing the single point of failure. If one node is blocked, others can pick up the task.

## The Tech Stack: Why these choices?

The project leverages a specific set of technologies to achieve its goals:

### 1. Go (Golang)
Go is the language of choice for high-performance networking and distributed systems (used by Ethereum, IPFS, Kubernetes).
*   **Concurrency**: Go's goroutines allow the Server and Client nodes to handle thousands of concurrent connections and tasks with minimal overhead.
*   **Ecosystem**: Excellent libraries for P2P (`libp2p`) and Blockchain (`go-ethereum`).

### 2. Libp2p
Libp2p is the networking stack that powers the P2P layer.
*   **PubSub**: Scrapeit uses the Publish-Subscribe pattern to broadcast tasks. The Server publishes a task to a topic, and all listening Clients receive it instantly. This decouples the producer from the consumers.
*   **Peer Discovery**: Enables nodes to find each other on the network without a central registry, using mechanisms like mDNS (local) and DHT (global).

### 3. Ethereum & Smart Contracts
Ethereum provides the trustless "financial" and "state" layer.
*   **Trust**: The `Scraper` smart contract acts as an escrow. When a task is created, funds are locked. They are released to the worker only when the work is completed.
*   **Transparency**: The history of tasks and payments is immutable and verifiable on-chain.

### 4. Chromedp & Go-Rod
These are Go libraries for controlling headless Chrome/Chromium browsers.
*   **Modern Web Support**: Unlike simple HTTP clients (curl), these tools render JavaScript, allowing Scrapeit to handle complex Single Page Applications (SPAs) and interact with dynamic content.
*   **Stealth**: The project integrates `go-rod/stealth` to evade bot detection scripts.

### 5. IPFS (InterPlanetary File System)
Storing large HTML blobs or screenshots on the Ethereum blockchain is prohibitively expensive.
*   **Off-chain Storage**: Scrapeit stores the scraping results on IPFS and only saves the unique content hash (CID) on the smart contract. This keeps gas costs low while ensuring data availability.

## Architecture Overview

The system consists of three main components:

1.  **The Server (Publisher/Oracle)**:
    *   Acts as the gateway for users to submit tasks via a web UI/API.
    *   Deploys the task to the Smart Contract.
    *   Publishes the task payload to the Libp2p PubSub topic.
    *   Currently acts as an "Oracle" to verify completion and trigger the payout.

2.  **The Client (Subscriber/Worker)**:
    *   Connects to the P2P network and subscribes to the task topic.
    *   Upon receiving a task, it claims it and launches a headless browser.
    *   Scrapes the target URL, uploads the result to IPFS.
    *   Reports the result back to the Server.

3.  **The Smart Contract**:
    *   Written in Solidity.
    *   Manages task state (`Pending`, `Completed`, `Failed`).
    *   Holds the bounty and transfers it to the assigned worker upon completion.

## Setting Up Locally

To run Scrapeit locally, you need Go, Docker, and Solidity tools installed.

**1. Install Dependencies**
Use the Makefile to install necessary Go tools and dependencies:
```bash
make deps
```

**2. Start Infrastructure**
Start the database (PostgreSQL) and any other services defined in Docker:
```bash
docker-compose up -d
```

**3. Build the Binaries**
Compile the Server, Client, and Contract tools:
```bash
make build
```

**4. Run the System**
*   **Deploy Contract**: Use the `bin/scrapeit-contract` tool to deploy the smart contract to your local chain or testnet.
*   **Start Server**:
    ```bash
    ./bin/scrapeit-server
    ```
*   **Start Client**:
    You need to provide an Ethereum private key for the worker wallet.
    ```bash
    export ETH_PRIVATE_KEY="your_private_key_here"
    ./bin/scrapeit-client
    ```

## Improvements & Future Direction

While the current PoC works, there are several areas for improvement:

*   **Task Distribution Strategies**:
    *   *Current*: PubSub broadcasts to everyone. This is inefficient at scale.
    *   *Improvement*: Use a Distributed Hash Table (DHT) or Kademlia routing to send tasks to specific nodes based on their reputation, geography, or capabilities.

*   **Decentralized Verification**:
    *   *Current*: The system relies on the Server to verify the scraping result.
    *   *Improvement*: Implement a consensus mechanism where multiple nodes scrape the same URL and compare results (e.g., checking the hash of the DOM). This removes the trust assumption on the Server.

*   **Security**:
    *   *Contract*: The `complete` function in the smart contract should be secured to prevent front-running. Only the authorized assignee or the original submitter should be able to trigger the completion logic.

*   **Privacy**:
    *   Encrypting task parameters (URL) so that only the assigned worker can read them, protecting the privacy of the requester.
