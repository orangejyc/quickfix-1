package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ApplicationMessageRequest struct {
	quickfixgo.Message
}

func (m *ApplicationMessageRequest) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) ApplReqType() (*field.ApplReqType, error) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
