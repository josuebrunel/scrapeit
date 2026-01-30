---
title: "How I Built a Decentralized Scraping Platform in Golang with Libp2p"
date: 2023-10-27T10:00:00Z
draft: false
tags: ["golang", "libp2p", "ethereum", "web-scraping", "architecture"]
---

Web scraping is a constant cat-and-mouse game. You write a script, it works for a week, and then you get IP banned. You rotate proxies, you change user agents, but eventually, centralized scraping infrastructure becomes fragile and expensive to maintain.

What if we could distribute the load across a network of decentralized nodes? What if anyone could run a "worker" node, contribute their bandwidth and IP reputation, and get paid for scraping tasks?

This is the premise behind **ScrapeIt** (or Scrapions), a decentralized scraping platform I built using **Golang**, **Libp2p**, and **Ethereum**.

> **Note**: This project is a **Proof of Concept (PoC)** to explore the intersection of P2P networking and web scraping. You can see it live at [https://scrapions.xyz](https://scrapions.xyz).

## The Architecture & "The Why"

When designing this system, I wanted to avoid the pitfalls of traditional scraping architectures:
1.  **Centralization**: A single point of failure (IP ban) can take down the whole operation.
2.  **Trust**: How do you trust a stranger to run code for you? How do they trust you to pay them?
3.  **Scalability**: Coordinating thousands of scrapers requires a robust messaging layer.

To solve this, the system is split into three main components:

1.  **The Server (Publisher)**: Acts as the orchestrator. It accepts HTTP requests from users, manages the database, and publishes scraping tasks to the P2P network.
2.  **The Client (Subscriber/Worker)**: Distributed nodes that listen for tasks, perform the scraping using a headless browser, upload results to IPFS, and submit proofs to the blockchain.
3.  **The Trust Layer (Ethereum)**: Ensures workers are paid and tasks are verified.

### The Stack: Tools of the Trade

To build this, I relied on a powerful set of tools. Here is why I chose them:

*   **Golang**: The backbone of the application. Go's concurrency model (goroutines and channels) is perfect for handling hundreds of peer connections and scraping tasks simultaneously without the overhead of threads.
*   **Libp2p**: Building a custom P2P protocol is hard. Libp2p handles the heavy lifting: NAT traversal, peer discovery, and pub/sub messaging.
*   **Chromedp & Go-Rod/Stealth**: Simple HTTP requests (`curl`) fail on modern sites that use React/Vue or anti-bot checks. We need a real browser. I used Chromedp to drive Chrome and injected Go-Rod's stealth scripts to evade detection.
*   **Ethereum & Anvil**: For the trust layer. I used **Anvil** (via Docker) for a fast local Ethereum development node. It allowed me to iterate on smart contracts instantly without waiting for testnet block times.
*   **Supabase / GoTrue**: Managing user auth is tedious. Supabase handles the email/password logic so I could focus on the P2P code.
*   **Bob**: A modern Go SQL ORM. It generates type-safe Go code from the database schema, preventing SQL injection and runtime errors.
*   **IPFS**: Storing large HTML files on-chain is prohibitively expensive. IPFS allows us to store the data cheaply and decentralized, only putting the small content hash (CID) on the blockchain.

## Part 1: The Networking Layer (Libp2p & Autodiscovery)

The biggest challenge in P2P is: **How do nodes find each other?**

In a centralized system, the client knows the server's IP. In a decentralized network, nodes come and go. We need **Autodiscovery**.

### Discovery Mechanisms

I implemented two layers of discovery using Libp2p:

1.  **MDNS (Multicast DNS)**: This is for **local network discovery**. When you spin up a node, it broadcasts a "I'm here!" message to the local subnet. This is crucial for development and testing (e.g., running multiple docker containers on your laptop).
2.  **DHT (Distributed Hash Table) / Bootstrap Nodes**: For the **global internet**. New nodes connect to a known "bootstrap" node (the Publisher). From there, they ask "Who else is on the network?" and the DHT helps them find other peers.

Here is the setup code in `pkg/common/common.go`:

```go
// pkg/common/common.go

func SetupDiscover(h host.Host) error {
    // Start MDNS service with a specific tag "scrapeit"
	m := mdns.NewMdnsService(h, DiscoveryTag, NewDiscoveryNotifee(h))
	return m.Start()
}

// Handler for when a peer is found via MDNS
func (d *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	xlog.Info("discovered new peer", "peer", pi.ID)
    // Connect to the peer immediately
	if err := d.h.Connect(context.Background(), pi); err != nil {
		xlog.Error("error while connecting to peer", "err", err)
	}
}
```

### PubSub with GossipSub

Once nodes are connected, they need to talk. We use **GossipSub**, a scalable PubSub protocol.

Why GossipSub? In a naive "flood" network, every node sends every message to every other node (O(N^2)). This crashes the network at scale. GossipSub creates a mesh: you only talk to a few neighbors, and they talk to theirs. It's efficient and robust.

**The Publisher (Server)** broadcasts a task:

```go
// server/server.go
func (s Server) Publish(ctx context.Context, tsk *task.Task) {
    // Marshal task to JSON and publish to the topic
	s.Node.Topic.Publish(ctx, tsk.JSONMarshal())
}
```

**The Subscriber (Client)** listens:

```go
// client/client.go
func (c Client) Subscribe(ctx context.Context, subscriber *pubsub.Subscription) {
	for {
		msg, _ := subscriber.Next(ctx)
        // Check if message is from self, then process
		if msg.ReceivedFrom == c.Node.ID { continue }
		// ...
	}
}
```

## Part 2: The Scraper (Stealth Mode)

Why did I use `chromedp` *and* `go-rod`?

`Chromedp` is the standard for driving Chrome in Go. However, out of the box, it leaks information that reveals it's a bot (e.g., `navigator.webdriver = true`).

`Go-Rod` has an excellent **Stealth** library that injects JavaScript to patch these leaks. I combined the two: using Chromedp for the driving (because of its robust API) and Go-Rod's scripts for the evasion.

```go
// scraper/scraper.go

func Scrape(ctx context.Context, url string) (string, error) {
    // 1. Configure Chrome to disable automation flags
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

    // ... context setup ...

	err := chromedp.Run(ctx,
        // 2. Inject the Stealth JS before the page loads
		chromedp.ActionFunc(stealthJS),
        // 3. Navigate and extract HTML
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent),
	)
    // ...
}
```

## Part 3: Trust & Settlement (Ethereum)

The final piece is trust. We use an Ethereum Smart Contract to hold funds in escrow. This ensures:
1.  **Payment Guarantee**: The worker knows the money is there before they start working.
2.  **Verification**: The server verifies the result before the money is released.

We use `go-ethereum` to generate Go bindings for the Solidity contract, making interaction type-safe.

```go
// task/ether.go

func (e EtherManager) Create(ctx context.Context, pk *ecdsa.PrivateKey, tsk Task) (*types.Transaction, error) {
	auth := e.getAuth(ctx, pk)
	auth.Value = etherutil.BigIntFromString(tsk.Reward) // Attach ETH value

	instance, _ := contract.NewContract(e.contractAddr, e.client)
    // Call the 'create' function on the smart contract
	return instance.Create(auth, s2b32(tsk.ID), tsk.Input)
}
```

## Future Improvements

This PoC demonstrates the viability of the idea, but there is room for optimization:

1.  **Arbitrage & Bidding**: Currently, tasks are broadcasted to everyone. We could implement a bidding system where workers bid on tasks based on their bandwidth or region (e.g., "I can scrape US-east sites faster"). This creates an efficient market for scraping.
2.  **Priority Queues**: Requesters could pay a premium (higher ETH reward) to get their scrapes processed by high-reputation nodes first.
3.  **Verification Games**: Instead of the Server being the sole verifier (which is centralized), we could have multiple workers scrape the same site. If their results match, they split the reward. If one differs, they are slashed.

## Conclusion

Building ScrapeIt allowed me to dive deep into the challenges of decentralized coordination. By combining the concurrency of Go, the resilience of Libp2p, and the trustlessness of Ethereum, we can build infrastructure that is harder to block and fairer for participants.

Check out the code and maybe run a node yourself!
