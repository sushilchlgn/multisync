package commands

import (
	// "time"

	"multisync-backend/browser"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Navigate(instance *browser.BrowserInstance, url string) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	if url == "" {
		return nil
	}

	_, err := instance.Page.Goto(url)
	return err
}

func (e *Executor) Reload(instance *browser.BrowserInstance) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	_, err := instance.Page.Reload()
	return err
}

func (e *Executor) Click(instance *browser.BrowserInstance, x, y float64) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	return instance.Page.Mouse().Click(x, y)
}

func (e *Executor) Type(instance *browser.BrowserInstance, text string) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	return instance.Page.Keyboard().Type(text)
}

func (e *Executor) Scroll(instance *browser.BrowserInstance, y float64) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	_, err := instance.Page.Evaluate(
		`(y) => window.scrollTo(0, y)`,
		y,
	)

	// time.Sleep(50 * time.Millisecond)

	return err
}

func (e *Executor) Eval(instance *browser.BrowserInstance, script string) error {
	if instance == nil || instance.Page == nil {
		return nil
	}

	if script == "" {
		return nil
	}

	_, err := instance.Page.Evaluate(script)
	return err
}