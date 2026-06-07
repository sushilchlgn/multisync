package browser

var browsers = make(map[string]*BrowserInstance)

func Start(sessionID string) *BrowserInstance {

	instance := &BrowserInstance{
		SessionID: sessionID,
		Status:    "running",
	}

	browsers[sessionID] = instance

	return instance
}

func Stop(sessionID string) {

	delete(browsers, sessionID)
}

func Get(sessionID string) (*BrowserInstance, bool) {

	browser, exists := browsers[sessionID]

	return browser, exists
}