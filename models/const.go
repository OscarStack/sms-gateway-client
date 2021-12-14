package models

type Path string

var (
	List   Path = "/sms_list"
	Read   Path = "/sms_read"
	Send   Path = "/sms_send"
	Total  Path = "/sms_total"
	Delete Path = "/sms_delete"
)
