package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type NetworkCounterpartySystemStatusRequest struct {
	quickfixgo.Message
}

func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestType() (*field.NetworkRequestType, error) {
	f := new(field.NetworkRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestID() (*field.NetworkRequestID, error) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}
