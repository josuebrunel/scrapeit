package scraper

import (
	"context"
	"scrapeit/pkg/env"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/go-rod/stealth"
	"github.com/josuebrunel/gopkg/xlog"
)

func Scrape(ctx context.Context, url string) (string, error) {
	// A common user agent string. You might want to rotate this or pick from a list of current, common UAs.
	const commonUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"

	// Start with default options (includes Headless, DisableGPU, NoFirstRun, NoDefaultBrowserCheck)
	allocOpts := chromedp.DefaultExecAllocatorOptions[:]
	// Add specific options for better fingerprinting and stability
	allocOpts = append(allocOpts,
		chromedp.NoSandbox, // Often required in containerized/CI environments.
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // Removes navigator.webdriver flag.
		chromedp.Flag("disable-infobars", "true"),                       // Removes "Chrome is being controlled..." infobar.
		chromedp.Flag("disable-extensions", true),                       // Disable extensions.
		chromedp.Flag("disable-component-extensions-with-background-pages", true),
		chromedp.UserAgent(commonUserAgent),
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

	// Create a new context with the specified user agent.
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		// Apply stealth script to all new documents. This should be early.
		chromedp.ActionFunc(stealthJS),

		chromedp.EmulateViewport(1920, 1080),              // A common desktop resolution.
		emulation.SetTimezoneOverride("America/New_York"), // Common timezone. Consider making this configurable.

		// Navigate to the target URL.
		chromedp.Navigate(url),
		// Extract the outer HTML of the page.
		chromedp.OuterHTML("html", &htmlContent),
	)

	if err != nil {
		xlog.Error("Failed to scrape URL", "url", url, "error", err)
		return "", err
	}

	return htmlContent, nil
}

func stealthJS(ctx context.Context) error {
	_, err := page.AddScriptToEvaluateOnNewDocument(stealth.JS).Do(ctx)
	if err != nil {
		xlog.Error("Failed to add stealth script", "error", err)
		return err
	}

	return nil
}
