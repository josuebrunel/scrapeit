package common

import "encoding/json"

const (
	ActionScrape     = "scrape"
	ActionScreenshot = "screenshot"
)

type Message struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Data   string `json:"data"`
	Type   string `json:"type"`
}

func (m Message) JSONMarshal() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

func (m *Message) JSONUnmarshal(b []byte) {
	if err := json.Unmarshal(b, m); err != nil {
		panic(err)
	}
}
