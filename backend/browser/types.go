package browser

import "github.com/playwright-community/playwright-go"


type BrowserInstance struct {
	SessionID string
	PW        *playwright.Playwright
	Browser   playwright.Browser
	Context   playwright.BrowserContext
	Page      playwright.Page
}