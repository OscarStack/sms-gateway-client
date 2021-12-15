package models

import (
	"github.com/OscarStack/sms-gateway-client/util"
)

type ReceivedMessage struct {
	Index  string
	Date   string
	Sender string
	Text   string
	Status string
}

func (s *ReceivedMessage) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := util.SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type TotalMessage struct {
	Used  string
	Total string
}

func (s *TotalMessage) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := util.SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
