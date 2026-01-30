---
title: "How I Built a Decentralized Scraping Platform in Golang with Libp2p - Part 2: The Networking Layer"
date: 2023-10-28T10:00:00Z
draft: false
tags: ["golang", "libp2p", "networking", "gossipsub"]
---

In Part 1, we looked at the high-level architecture of ScrapeIt. Now, let's get into the engine that makes decentralization possible: **Libp2p**.

We need a way for the Server (Publisher) to say "Hey, I need this URL scraped" and for any available Client (Worker) to hear that request instantly, without maintaining a direct connection to every single worker. This is a perfect use case for **Pub/Sub** (Publish/Subscribe).

## The Node Setup

Every participant in our network (both Server and Client) is a Libp2p "Node". I created a common `Node` struct in `pkg/common/common.go` to handle the boilerplate of setting up a host, discovery, and GossipSub.

```go
// pkg/common/common.go

type Node struct {
	ID        peer.ID
	Address   string
	Host      host.Host
	Topic     *pubsub.Topic
	TopicName string
}

func NewNode(ctx context.Context, topicName, address string, bootStrappingPeers []multiaddr.Multiaddr, serverMode bool) *Node {
    // 1. Create the Libp2p Host
	host, err := libp2p.New(libp2p.ListenAddrStrings(address))
	if err != nil {
		panic(err)
	}

    // 2. Create the GossipSub service
	gossip, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}

    // 3. Join the specific Topic (e.g., "scrapeit-tasks")
	topic, err := gossip.Join(topicName)
	if err != nil {
		panic(err)
	}

    // ... DHT and Discovery setup ...

	return &Node{
		ID:        host.ID(),
		Address:   address,
		Host:      host,
		Topic:     topic,
		TopicName: topicName,
	}
}
```

We use **GossipSub** because it's robust and efficient for mesh networks. It ensures messages propagate quickly without flooding the network.

## Discovery: Finding Peers

A decentralized network is useless if nodes can't find each other. We use two mechanisms:
1.  **MDNS**: For finding peers on the same local network (great for testing).
2.  **DHT (Distributed Hash Table)**: For finding peers across the internet using bootstrap nodes.

```go
// pkg/common/common.go

func SetupDiscover(h host.Host) error {
	m := mdns.NewMdnsService(h, DiscoveryTag, NewDiscoveryNotifee(h))
	return m.Start()
}
```

## The Publisher (Server)

On the server side (`server/server.go`), we initialize the node and provide a method to `Publish` tasks.

When a user submits a scrape request via the API, the server marshals the task into JSON and broadcasts it to the topic.

```go
// server/server.go

func (s Server) Publish(ctx context.Context, tsk *task.Task) {
	xlog.Info("publishing task", "task", *tsk, "topic", s.Node.TopicName)
	// Publish the JSON-encoded task to the GossipSub topic
	s.Node.Topic.Publish(ctx, tsk.JSONMarshal())
	xlog.Info("task published")
}
```

## The Subscriber (Client)

The client (`client/client.go`) joins the network and subscribes to the topic. It enters a loop where it waits for messages.

```go
// client/client.go

func (c Client) Subscribe(ctx context.Context, subscriber *pubsub.Subscription) {
	xlog.Info("subscribing to topic", "topic", c.TopicName)
	for {
        // Block until a new message is received
		msg, err := subscriber.Next(ctx)
		if err != nil {
			panic(err)
		}

        // Ignore messages from ourselves
		if msg.ReceivedFrom == c.Node.ID {
			continue
		}

		var tsk task.Task
		tsk.JSONUnmarshal(msg.Data)

        // Process the task...
		xlog.Info("go message", "from", msg.ReceivedFrom, "topic", msg.Topic, "task", tsk)

        // ... Logic to claim and execute the task ...
	}
}
```

Once a task is received, the client doesn't just start scraping immediately. It first attempts to **Claim** the task to prevent redundant work (we'll touch on this coordination in Part 4).

## Conclusion

With Libp2p, we've decoupled the task creator from the task executor. The server doesn't know who will perform the scrape, and the worker doesn't need to know the server's IP address beforehand—they just need to find the topic.

In **Part 3**, we will look at what happens *after* the client receives the task: **The Scraper**. We'll see how we use `chromedp` and stealth techniques to browse the web like a human.
