package browser

import "sync"

var (
	mu        sync.Mutex
	instances = make(map[string]*BrowserInstance)
)

// START browser for session
func Start(sessionID string) (*BrowserInstance, error) {
	instance, err := Launch(sessionID)
	if err != nil {
		return nil, err
	}

	mu.Lock()
	instances[sessionID] = instance
	mu.Unlock()

	return instance, nil
}

// STOP browser safely
func Stop(sessionID string) {
	mu.Lock()
	inst, ok := instances[sessionID]
	if ok {
		delete(instances, sessionID)
	}
	mu.Unlock()

	if !ok {
		return
	}

	if inst.Page != nil {
		_ = inst.Page.Close()
	}
	if inst.Context != nil {
		_ = inst.Context.Close()
	}
	if inst.Browser != nil {
		_ = inst.Browser.Close()
	}
	if inst.PW != nil {
		inst.PW.Stop()
	}
}

// GET browser instance
func Get(sessionID string) (*BrowserInstance, bool) {
	mu.Lock()
	defer mu.Unlock()

	inst, ok := instances[sessionID]
	return inst, ok
}