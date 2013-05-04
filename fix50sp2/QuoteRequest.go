package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteRequest struct {
	quickfixgo.Message
}

func (m *QuoteRequest) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) PrivateQuote() (*field.PrivateQuote, error) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) RespondentType() (*field.RespondentType, error) {
	f := new(field.RespondentType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) PreTradeAnonymity() (*field.PreTradeAnonymity, error) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteRequest) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
