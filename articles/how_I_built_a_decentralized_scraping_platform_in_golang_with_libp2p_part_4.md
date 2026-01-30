---
title: "How I Built a Decentralized Scraping Platform in Golang with Libp2p - Part 4: Ethereum & Trust"
date: 2023-10-30T10:00:00Z
draft: false
tags: ["golang", "ethereum", "smart-contracts", "blockchain"]
---

In a decentralized system, you can't trust the nodes. A worker might say "I scraped the site" but return garbage data. A requester might say "I'll pay you" but never send the money.

To solve this, ScrapeIt uses an **Ethereum Smart Contract** to act as an escrow and truth layer.

## The Task Lifecycle

1.  **Creation**: The Server (Requester) sends a transaction to the contract with the task details and the reward (ETH). The contract holds the funds.
2.  **Processing**: The worker scrapes the site and uploads the result to IPFS.
3.  **Completion**: The Server verifies the result and signals the contract to release the funds to the worker.

*Note: In a fully trustless version, the verification would happen on-chain or via a challenge mechanism. For this MVP, the Server acts as the Oracle verifying the work.*

## Interacting with Ethereum

I used `go-ethereum` (geth) to generate Go bindings for the smart contract. The interaction logic is encapsulated in `task/ether.go`.

### Creating a Task

When a user submits a task via the API, the server creates it on-chain.

```go
// task/ether.go

func (e EtherManager) Create(ctx context.Context, pk *ecdsa.PrivateKey, tsk Task) (*types.Transaction, error) {
	auth := e.getAuth(ctx, pk)
	auth.Value = etherutil.BigIntFromString(tsk.Reward) // Attach ETH reward

	instance, err := contract.NewContract(e.contractAddr, e.client)

    // Call the smart contract's Create function
	tx, err := instance.Create(auth, s2b32(tsk.ID), tsk.Input)

	xlog.Info("Task created", "tx", tx.Hash().Hex())
	return tx, nil
}
```

### Listening for Events

The server doesn't just fire and forget. It watches the blockchain for events to update its local database state. This is handled in `server/event.go`.

```go
// server/event.go

func (s Server) ListenContractEvents(ctx context.Context) {
    // ... Dial WebSocket ...
	filterer, _ := contract.NewContractFilterer(ctAddr, evCli)

	taskCreatedEvents := make(chan *contract.ContractTaskCreated)
	subCreated, _ := filterer.WatchTaskCreated(nil, taskCreatedEvents, nil, nil)

	for {
		select {
		case event := <-taskCreatedEvents:
			xlog.Info("Task created event detected", "id", event.Id)
            // Update local DB to mark task as 'OnChain'
		}
	}
}
```

### Completing the Task

Once the worker submits the result (via the API, which the server validates), the server triggers the payout.

```go
// task/ether.go

func (e EtherManager) Complete(ctx context.Context, id string, assignee string, output string, success bool, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
    // ... Auth setup ...

	instance, _ := contract.NewContract(e.contractAddr, e.client)

    // Call the smart contract's Complete function
    // This releases the funds to the 'assignee' address
	tx, err := instance.Complete(auth, s2b32(id), common.HexToAddress(assignee), output, success)

	return tx, nil
}
```

## Wrapping Up

By combining **Golang** for the robust backend, **Libp2p** for decentralized coordination, **Chromedp** for scraping, and **Ethereum** for trustless settlement, we've built a functional decentralized scraping marketplace.

The code demonstrates how modern P2P and blockchain technologies can be integrated into standard Go applications to build resilient distributed systems.

You can check out the full source code in the repository. Happy scraping!
