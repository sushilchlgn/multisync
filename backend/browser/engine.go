package browser

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func Launch(sessionID string) (*BrowserInstance, error) {

	// Install browsers (safe to call repeatedly in dev)
	if err := playwright.Install(); err != nil {
		log.Println("Playwright install warning:", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Channel:  playwright.String("chrome"),
	})

	if err != nil {
		return nil, err
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, err
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}

	return &BrowserInstance{
		SessionID: sessionID,
		PW:        pw,
		Browser:   browser,
		Context:   context,
		Page:      page,
	}, nil
}