package _struct

type Root struct {
	Creator               string      `json:"Creator"`
	NeedUpDateFields      interface{} `json:"NeedUpDateFields"`
	NeedReturnFields      interface{} `json:"NeedReturnFields"`
	IsDeleteEntry         bool        `json:"IsDeleteEntry"`
	SubSystemID           string      `json:"SubSystemId"`
	IsVerifyBaseDataField bool        `json:"IsVerifyBaseDataField"`
	IsEntryBatchFill      bool        `json:"IsEntryBatchFill"`
	ValidateFlag          bool        `json:"ValidateFlag"`
	NumberSearch          bool        `json:"NumberSearch"`
	InterationFlags       string      `json:"InterationFlags"`
	IsAutoSubmitAndAudit  bool        `json:"IsAutoSubmitAndAudit"`
	Model                 interface{} `json:"Model"`
}
