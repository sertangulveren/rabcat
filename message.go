package main

import "encoding/json"

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

func (m *Message) generatePayload() []byte {
	b, _ := json.Marshal(m)
	return b
}

func parseMessage(payload []byte) *Message {
	res := Message{}
	_ = json.Unmarshal(payload, &res)
	return &res
}
