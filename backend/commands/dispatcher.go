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

	// 3. Route ONLY (NO EXECUTION HERE)
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

func (d *Dispatcher) handleNavigate(instance *browser.BrowserInstance, cmd Command) error {
	exec := NewExecutor()

	data, ok := cmd.Data.(map[string]interface{})
	if !ok {
		return errors.New("invalid navigate payload")
	}

	url, _ := data["url"].(string)
	return exec.Navigate(instance, url)
}

func (d *Dispatcher) handleReload(instance *browser.BrowserInstance) error {
	exec := NewExecutor()
	return exec.Reload(instance)
}

func (d *Dispatcher) handleClick(instance *browser.BrowserInstance, cmd Command) error {
	exec := NewExecutor()

	data, ok := cmd.Data.(map[string]interface{})
	if !ok {
		return errors.New("invalid click payload")
	}

	x, _ := data["x"].(float64)
	y, _ := data["y"].(float64)

	return exec.Click(instance, x, y)
}

func (d *Dispatcher) handleType(instance *browser.BrowserInstance, cmd Command) error {
	exec := NewExecutor()

	data, ok := cmd.Data.(map[string]interface{})
	if !ok {
		return errors.New("invalid type payload")
	}

	text, _ := data["text"].(string)

	return exec.Type(instance, text)
}

func (d *Dispatcher) handleScroll(instance *browser.BrowserInstance, cmd Command) error {
	exec := NewExecutor()

	data, ok := cmd.Data.(map[string]interface{})
	if !ok {
		return errors.New("invalid scroll payload")
	}

	y, _ := data["y"].(float64)

	return exec.Scroll(instance, y)
}

func (d *Dispatcher) handleEval(instance *browser.BrowserInstance, cmd Command) error {
	exec := NewExecutor()

	data, ok := cmd.Data.(map[string]interface{})
	if !ok {
		return errors.New("invalid eval payload")
	}

	script, _ := data["script"].(string)

	return exec.Eval(instance, script)
}