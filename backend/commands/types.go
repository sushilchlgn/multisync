package commands

// ===============================
// 1. COMMAND TYPES (ENUM STYLE)
// ===============================

type CommandType string

const (
	Navigate CommandType = "NAVIGATE"
	Reload   CommandType = "RELOAD"

	Click  CommandType = "CLICK"
	Type   CommandType = "TYPE"
	Scroll CommandType = "SCROLL"

	Eval CommandType = "EVAL"
)


// ===============================
// 2. BASE COMMAND STRUCT
// ===============================

type Command struct {
	SessionID string      `json:"sessionId"`
	Type      CommandType `json:"type"`
	Data      interface{} `json:"data"`
}


// ===============================
// 3. STRONGLY TYPED PAYLOADS
// ===============================

// NAVIGATE
type NavigateData struct {
	URL string `json:"url"`
}

// CLICK
type ClickData struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// TYPE (keyboard input)
type TypeData struct {
	Selector string `json:"selector"`
	Text     string `json:"text"`
}

// SCROLL
type ScrollData struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// EVAL (run JS)
type EvalData struct {
	Script string `json:"script"`
}