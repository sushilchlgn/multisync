package commands

import (
	"errors"

	"multisync-backend/browser"
)

// ===============================
// Dispatcher (CORE BRAIN)
// ===============================

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(cmd Command) error {

	// 1. Validate session ID
	if cmd.SessionID == "" {
		return errors.New("session id is required")
	}

	// 2. Get browser instance
	instance, exists := browser.Get(cmd.SessionID)
	if !exists {
		return errors.New("browser instance not found for session")
	}

	// 3. Route ONLY (NO EXECUTION YET)
	switch cmd.Type {

	case Navigate:
		return d.handleNavigate(instance, cmd)

	case Reload:
		return d.handleReload(instance)

	case Click:
		return d.handleClick(instance, cmd)

	case Type:
		return d.handleType(instance, cmd)

	case Scroll:
		return d.handleScroll(instance, cmd)

	case Eval:
		return d.handleEval(instance, cmd)

	default:
		return errors.New("unknown command type")
	}
}

// ===============================
// Handlers (STUB ONLY)
// ===============================

func (d *Dispatcher) handleNavigate(instance *browser.BrowserInstance, cmd Command) error {
	_ = instance
	_ = cmd
	return nil
}

func (d *Dispatcher) handleReload(instance *browser.BrowserInstance) error {
	_ = instance
	return nil
}

func (d *Dispatcher) handleClick(instance *browser.BrowserInstance, cmd Command) error {
	_ = instance
	_ = cmd
	return nil
}

func (d *Dispatcher) handleType(instance *browser.BrowserInstance, cmd Command) error {
	_ = instance
	_ = cmd
	return nil
}

func (d *Dispatcher) handleScroll(instance *browser.BrowserInstance, cmd Command) error {
	_ = instance
	_ = cmd
	return nil
}

func (d *Dispatcher) handleEval(instance *browser.BrowserInstance, cmd Command) error {
	_ = instance
	_ = cmd
	return nil
}