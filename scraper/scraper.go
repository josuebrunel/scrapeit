package scraper

import (
	"context"
	"scrapeit/pkg/env"
	"scrapeit/pkg/task"

	"github.com/chromedp/chromedp"
	"github.com/josuebrunel/gopkg/xlog"
)

func Scrape(ctx context.Context, tsk *task.Task) error {
	// Configure allocator options.
	// chromedp.DefaultExecAllocatorOptions are used as a base, which will
	// attempt to find an installed Chrome/Chromium browser.
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Headless,   // Run in headless mode (no GUI).
		chromedp.NoSandbox,  // Often required in containerized/CI environments.
		chromedp.DisableGPU, // Can improve stability and reduce resource usage.
		// You could add other flags here if needed, e.g.:
		// chromedp.Flag("user-agent", "MyScraperAgent/1.0"),
	)

	var (
		allocCtx    = ctx
		allocCancel context.CancelFunc
	)

	if env.Get(env.CDPRemoteN) != "" {
		// allocOpts = append(allocOpts, chromedp.Flag("remote-debugging-address", env.Get(env.CDPRemoteN)))
		wsUrl := env.Get(env.CDPRemoteN)
		xlog.Info("Using remote debugging", "address", wsUrl)
		allocCtx, allocCancel = chromedp.NewRemoteAllocator(allocCtx, wsUrl)
	} else {
		// Create an allocator context. This manages the browser process.
		allocCtx, allocCancel = chromedp.NewExecAllocator(allocCtx, allocOpts...)
	}
	defer allocCancel()

	// Create a new chromedp context from the allocator context.
	// This context is used for individual browser tab operations.
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(tsk.Input),
		chromedp.OuterHTML("html", &htmlContent),
	)

	if err != nil {
		xlog.Error("Failed to scrape URL", "url", tsk.Input, "error", err)
		tsk.Output = err.Error()
		return err
	}

	tsk.Output = htmlContent
	return nil
}
