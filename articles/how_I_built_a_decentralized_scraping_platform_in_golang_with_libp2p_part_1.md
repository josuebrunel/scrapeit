---
title: "How I Built a Decentralized Scraping Platform in Golang with Libp2p - Part 1: The Architecture"
date: 2023-10-27T10:00:00Z
draft: false
tags: ["golang", "libp2p", "web-scraping", "architecture"]
---

Web scraping is a constant cat-and-mouse game. You write a script, it works for a week, and then you get IP banned. You rotate proxies, you change user agents, but eventually, centralized scraping infrastructure becomes fragile and expensive to maintain.

What if we could distribute the load across a network of decentralized nodes? What if anyone could run a "worker" node, contribute their bandwidth and IP reputation, and get paid for scraping tasks?

This is the premise behind **ScrapeIt**, a decentralized scraping platform I built using **Golang**, **Libp2p**, and **Ethereum**.

In this four-part series, I'll walk you through how I built it. We'll cover the architecture, the P2P networking, the stealth scraping engine, and the blockchain-based trust layer.

## The Problem with Centralization

When you run a scraper from a single server (or even a cluster of servers with a shared proxy pool), you have a single point of failure. Target websites can easily fingerprint your traffic patterns.

## The Solution: Decentralization

By leveraging a P2P network, we can achieve:
1.  **Resilience**: No single point of failure.
2.  **IP Diversity**: Traffic comes from residential IPs all over the world (the workers).
3.  **Trustless Coordination**: Using a blockchain to manage task assignments and payments ensures that workers get paid for valid work and requesters get valid data.

## High-Level Architecture

The system consists of three main components:

1.  **The Server (Publisher)**: This is the entry point for users who want to scrape data. It exposes an API to accept tasks and publishes them to the P2P network.
2.  **The Client (Subscriber/Worker)**: These are the nodes running on various machines. They listen for tasks, execute the scraping job, store the result on IPFS, and submit the proof to the blockchain.
3.  **The Blockchain (Ethereum)**: Handles the financial settlement and state of tasks. It ensures that a task was actually created (paid for) and completed (verified).

### The Stack

*   **Language**: Golang (for its concurrency and strong ecosystem).
*   **Networking**: `libp2p` (for GossipSub and peer discovery).
*   **Scraping**: `chromedp` (Chrome DevTools Protocol) + `go-rod/stealth` (for evasion).
*   **Storage**: IPFS (for storing large HTML results).
*   **Database**: PostgreSQL (for the server's local state) & Bob (ORM).

## Diving into the Code: The Application Entry Point

The core of the system is orchestrated in the `App` struct. This handles the database connections, the P2P publisher, and the Ethereum manager.

Here is how the application is initialized in `app/app.go`. It sets up the database, the Ethereum manager, and the Publisher (which we will cover in Part 2).

```go
// app/app.go

type App struct {
	ListenAddr  string
	Publisher   *server.Server // Handles P2P publishing
	Manager     *task.EtherManager // Handles Ethereum interactions
	privateKey  *ecdsa.PrivateKey
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
		AuthBackend: env.Get(env.AppAuthBackendN),
		DatabaseURL: env.Get(env.AppDatabaseN),
        // Initialize the P2P Server
		Publisher:   server.New(*m),
		privateKey:  etherutil.MustPrivateKeyFromHex(env.Get(env.EtherPrivateKeyN)),
		Manager:     m,
		ListenAddr:  env.Get(env.AppAddrN),
	}
}
```

The `Start` method in `app/app.go` ties everything together. It sets up the HTTP server (Echo) for the API and starts the background listener for contract events.

```go
// app/app.go

func (a *App) Start() error {
	e := echo.New()
    // ... Middleware setup ...

	// Task API Group
	taskV1 := e.Group("/api/v1/task", ApiKeyAuthMiddleware(a.bob))
	taskV1.POST("", a.Scrape()) // Endpoint to submit a new scrape task

    // ... Mount other routes ...

	xlog.Info("starting server", "addr", a.ListenAddr)

    // Start listening for smart contract events in the background
	go a.Publisher.ListenContractEvents(context.Background())

	log.Fatal(e.Start(a.ListenAddr))
	return nil
}
```

## What's Next?

We have our server running and ready to accept HTTP requests. But how does a request translate into a job that a worker node on the other side of the world can pick up?

In **Part 2**, we will dive into the **Networking Layer**. We will explore how to use `libp2p` and GossipSub to broadcast these tasks to a decentralized network of scrapers.
