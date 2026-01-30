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

## The Architecture

The system consists of three main components:

1.  **The Server (Publisher)**: Acts as the orchestrator. It accepts HTTP requests from users, manages the database, and publishes scraping tasks to the P2P network.
2.  **The Client (Subscriber/Worker)**: Distributed nodes that listen for tasks, perform the scraping using a headless browser, upload results to IPFS, and submit proofs to the blockchain.
3.  **The Trust Layer (Ethereum)**: Ensures workers are paid and tasks are verified.

### The Stack & Tools

To build this, I relied on a powerful set of tools:

*   **Golang**: The backbone of the application.
*   **Libp2p**: For peer-to-peer networking, discovery, and PubSub messaging (GossipSub).
*   **Chromedp**: For driving the headless Chrome browser.
*   **Go-Rod/Stealth**: I used the stealth plugins from Go-Rod (injected into Chromedp) to evade bot detection.
*   **Ethereum & Anvil**: I used **Anvil** (via Docker) for a fast local Ethereum development node to test smart contracts without spending real ETH.
*   **Supabase / GoTrue**: For handling user authentication on the platform.
*   **Bob**: A modern Go SQL ORM and query builder for interacting with PostgreSQL.
*   **IPFS**: For decentralized storage of the large scraped HTML content.

## Part 1: The Application Entry Point

The core of the system is orchestrated in the `App` struct. This handles the database connections, the P2P publisher, and the Ethereum manager.

```go
// app/app.go

type App struct {
	ListenAddr  string
	Publisher   *server.Server // Handles P2P publishing
	Manager     *task.EtherManager // Handles Ethereum interactions
	// ...
}

func NewApp() *App {
	db := getDBConnection(env.Get(env.AppDatabaseN))
	bdb := bob.NewDB(db)

    // Initialize the Ethereum Manager
    m := task.NewEtherManager(
		env.Get(env.EtherAddrN),
		env.Get(env.EtherContractAddrN),
	)
	m.SetDB(bdb)

	return &App{
		db:          db,
		bob:         bdb,
        // ...
		Publisher:   server.New(*m),
		Manager:     m,
	}
}
```

## Part 2: The Networking Layer (Libp2p)

We need a way for the Server to say "Hey, scrape this URL" and for any available Client to hear it instantly. **GossipSub** is perfect for this.

I created a common `Node` struct to handle the Libp2p boilerplate.

```go
// pkg/common/common.go

type Node struct {
	ID        peer.ID
	Topic     *pubsub.Topic
    // ...
}

func NewNode(...) *Node {
    // 1. Create the Libp2p Host
	host, err := libp2p.New(libp2p.ListenAddrStrings(address))

    // 2. Create the GossipSub service
	gossip, err := pubsub.NewGossipSub(ctx, host)

    // 3. Join the topic
	topic, err := gossip.Join(topicName)

	return &Node{...}
}
```

The **Server** publishes tasks:

```go
// server/server.go
func (s Server) Publish(ctx context.Context, tsk *task.Task) {
	s.Node.Topic.Publish(ctx, tsk.JSONMarshal())
}
```

The **Client** subscribes:

```go
// client/client.go
func (c Client) Subscribe(ctx context.Context, subscriber *pubsub.Subscription) {
	for {
		msg, _ := subscriber.Next(ctx)
		// Process task...
	}
}
```

## Part 3: The Scraper (Stealth Mode)

Standard headless Chrome is easy to detect. To fix this, I used **Chromedp** for the driving but injected **Go-Rod's stealth scripts** to mask the browser fingerprint.

```go
// scraper/scraper.go

func Scrape(ctx context.Context, url string) (string, error) {
    // Chrome options to look human
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

    // ... context setup ...

	err := chromedp.Run(ctx,
        // Inject stealth JS early
		chromedp.ActionFunc(stealthJS),
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent),
	)
    // ...
}
```

After scraping, the result is uploaded to **IPFS**, and only the content hash (CID) is returned.

## Part 4: Trust & Settlement (Ethereum)

We use an Ethereum Smart Contract to hold funds in escrow.

1.  **Create**: Requester locks ETH in the contract.
2.  **Work**: Scraper performs the task.
3.  **Complete**: Smart contract verifies the submission and releases funds.

We use `go-ethereum` to interact with the contract:

```go
// task/ether.go

func (e EtherManager) Create(ctx context.Context, pk *ecdsa.PrivateKey, tsk Task) (*types.Transaction, error) {
	auth := e.getAuth(ctx, pk)
	auth.Value = etherutil.BigIntFromString(tsk.Reward) // Attach ETH

	instance, _ := contract.NewContract(e.contractAddr, e.client)
	return instance.Create(auth, s2b32(tsk.ID), tsk.Input)
}
```

## Future Improvements

This PoC demonstrates the viability of the idea, but there is room for optimization:

1.  **Arbitrage & Bidding**: Currently, tasks are broadcasted to everyone. We could implement a bidding system where workers bid on tasks based on their bandwidth or region, creating an efficient market.
2.  **Priority Queues**: Requesters could pay a premium gas fee or higher task reward to get their scrapes processed faster.
3.  **Verification Games**: Instead of a central verifier, we could use a mechanism where multiple workers scrape the same site, and the consensus result is accepted (slashing the stake of any malicious worker).

## Conclusion

Building ScrapeIt allowed me to dive deep into the challenges of decentralized coordination. By combining the concurrency of Go, the resilience of Libp2p, and the trustlessness of Ethereum, we can build infrastructure that is harder to block and fairer for participants.

Check out the code and maybe run a node yourself!
