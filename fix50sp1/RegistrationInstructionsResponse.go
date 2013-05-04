package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type RegistrationInstructionsResponse struct {
	quickfixgo.Message
}

func (m *RegistrationInstructionsResponse) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) RegistTransType() (*field.RegistTransType, error) {
	f := new(field.RegistTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) RegistRefID() (*field.RegistRefID, error) {
	f := new(field.RegistRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) RegistStatus() (*field.RegistStatus, error) {
	f := new(field.RegistStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) RegistRejReasonCode() (*field.RegistRejReasonCode, error) {
	f := new(field.RegistRejReasonCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *RegistrationInstructionsResponse) RegistRejReasonText() (*field.RegistRejReasonText, error) {
	f := new(field.RegistRejReasonText)
	err := m.Body.Get(f)
	return f, err
}
