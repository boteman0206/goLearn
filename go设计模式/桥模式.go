package main

import "fmt"

/**
桥接模式
桥接模式分离抽象部分和实现部分。使得两部分独立扩展。
桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。
策略模式使抽象部分和实现部分分离，可以独立变化。
*/

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct {
}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

func main() {

	sms := ViaSMS()
	message := NewCommonMessage(sms)
	message.SendMessage("sms", "jack")

	email := ViaEmail()
	urgencyMessage := NewUrgencyMessage(email)
	urgencyMessage.SendMessage("email", "bob")

}
