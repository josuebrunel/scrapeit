package scraper

import (
	"context"
	"scrapeit/pkg/env"

	"github.com/chromedp/chromedp"
	"github.com/josuebrunel/gopkg/xlog"
)

func Scrape(ctx context.Context, url string) (string, error) {

	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Headless,   // Run in headless mode (no GUI).
		chromedp.NoSandbox,  // Often required in containerized/CI environments.
		chromedp.DisableGPU, // Can improve stability and reduce resource usage.
	)

	var (
		allocCtx    = ctx
		allocCancel context.CancelFunc
	)

	if env.Get(env.CDPRemoteN) != "" {
		wsUrl := env.Get(env.CDPRemoteN)
		xlog.Info("Using remote debugging", "address", wsUrl)
		allocCtx, allocCancel = chromedp.NewRemoteAllocator(allocCtx, wsUrl)
	} else {
		allocCtx, allocCancel = chromedp.NewExecAllocator(allocCtx, allocOpts...)
	}
	defer allocCancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent),
	)

	if err != nil {
		xlog.Error("Failed to scrape URL", "url", url, "error", err)
		return "", err
	}

	return htmlContent, nil
}
