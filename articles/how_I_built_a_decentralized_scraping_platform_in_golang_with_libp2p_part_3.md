---
title: "How I Built a Decentralized Scraping Platform in Golang with Libp2p - Part 3: The Scraper"
date: 2023-10-29T10:00:00Z
draft: false
tags: ["golang", "web-scraping", "chromedp", "stealth"]
---

So far, we have a server publishing tasks and clients listening for them. Now comes the actual work: visiting a URL and extracting data.

For this project, I chose **Chromedp**, a Golang library that drives a headless Chrome browser. While simple HTTP requests (like `curl` or `net/http`) are faster, they fail on modern websites that rely heavily on JavaScript or have bot protection. A real browser is necessary.

## The Scraper Package

The core logic lives in `scraper/scraper.go`.

### Avoiding Detection with Stealth

Standard headless Chrome is easy to detect. It leaves "fingerprints" like `navigator.webdriver = true` that scream "I am a bot".

To counter this, I integrated the `go-rod/stealth` library. Even though I'm using `chromedp` for the driving, I inject the `stealth` JavaScript payload into the page before it loads.

```go
// scraper/scraper.go

import (
    "github.com/chromedp/chromedp"
    "github.com/go-rod/stealth"
)

func Scrape(ctx context.Context, url string) (string, error) {
    // ... Allocator setup ...

    var htmlContent string
    err := chromedp.Run(ctx,
        // Apply stealth script to all new documents early
        chromedp.ActionFunc(stealthJS),

        chromedp.EmulateViewport(1920, 1080),

        // Navigate to the target URL
        chromedp.Navigate(url),

        // Extract the full HTML
        chromedp.OuterHTML("html", &htmlContent),
    )

    if err != nil {
        return "", err
    }
    return htmlContent, nil
}

// Inject the stealth JS
func stealthJS(ctx context.Context) error {
	_, err := page.AddScriptToEvaluateOnNewDocument(stealth.JS).Do(ctx)
	return err
}
```

### Configuring Chrome

In `scraper/scraper.go`, we also configure Chrome to look as "human" as possible. We disable automation flags and set a common User-Agent.

```go
	allocOpts := chromedp.DefaultExecAllocatorOptions[:]
	allocOpts = append(allocOpts,
		chromedp.NoSandbox,
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // Key for evasion
		chromedp.Flag("disable-infobars", "true"),
		chromedp.UserAgent("Mozilla/5.0 ... Chrome/114.0.0.0 ..."),
	)
```

## Storing the Result

Once the scraping is complete, we have a large string of HTML. We don't want to send this entire payload back through the blockchain (too expensive) or the PubSub channel (too noisy).

Instead, the client uploads the result to **IPFS** (InterPlanetary File System).

```go
// client/client.go

// Scrape the content
content, err := scraper.Scrape(ctx, tsk.Input)

// Upload to IPFS
tsk.Output, err = ipfsCli.UploadFile(ctx, content)
```

The `tsk.Output` now contains the **CID** (Content Identifier), a small string representing the file on IPFS. This CID is what we will eventually record on the blockchain.

## Conclusion

We now have a stealthy scraper that can execute JavaScript and hide its bot nature. We have a way to store the large results off-chain using IPFS.

But how do we ensure the worker actually did the work? How do they get paid? That's the job of the **Ethereum Smart Contract**, which we will cover in the final part.
