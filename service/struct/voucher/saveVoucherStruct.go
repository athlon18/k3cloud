package voucher

type Model struct {
	FVOUCHERID       string      `json:"FVOUCHERID"`
	FAccountBookID   interface{} `json:"FAccountBookID"`
	FDate            string      `json:"FDate"`
	FVOUCHERGROUPID  interface{} `json:"FVOUCHERGROUPID"`
	FVOUCHERGROUPNO  string      `json:"FVOUCHERGROUPNO"`
	FATTACHMENTS     string      `json:"FATTACHMENTS"`
	FISADJUSTVOUCHER bool        `json:"FISADJUSTVOUCHER"`
	FDocumentStatus  string      `json:"FDocumentStatus"`
	FYEAR            string      `json:"FYEAR"`
	FSourceBillKey   interface{} `json:"FSourceBillKey"`
	FCreatorId       interface{} `json:"FCreatorId"`
	FPERIOD          string      `json:"FPERIOD"`
	FIMPORTVERSION   string      `json:"FIMPORTVERSION"`
	FEntity          []FEntity   `json:"FEntity"`
}

type FEntity struct {
	FEntryID          string      `json:"FEntryID,omitempty"`
	FEXPLANATION      string      `json:"FEXPLANATION,omitempty"`
	FACCOUNTID        interface{} `json:"FACCOUNTID,omitempty"`
	FDetailID         interface{} `json:"FDetailID,omitempty"`
	FCURRENCYID       interface{} `json:"FCURRENCYID,omitempty"`
	FEXCHANGERATETYPE interface{} `json:"FEXCHANGERATETYPE,omitempty"`
	FEXCHANGERATE     string      `json:"FEXCHANGERATE,omitempty"`
	FUnitID           interface{} `json:"FUnitId,omitempty"`
	FPrice            string      `json:"FPrice,omitempty"`
	FQty              string      `json:"FQty,omitempty"`
	FAMOUNTFOR        string      `json:"FAMOUNTFOR,omitempty"`
	FDEBIT            string      `json:"FDEBIT,omitempty"`
	FCREDIT           string      `json:"FCREDIT,omitempty"`
	FSettleTypeID     interface{} `json:"FSettleTypeID,omitempty"`
	FSETTLENO         string      `json:"FSETTLENO,omitempty"`
	FEXPORTENTRYID    string      `json:"FEXPORTENTRYID,omitempty"`
}

type FDetailID struct {
	FDETAILIDFFLEX9   interface{} `json:"FDETAILID__FFLEX9,omitempty"`
	FDETAILIDFFLEX10  interface{} `json:"FDETAILID__FFLEX10,omitempty"`
	FDETAILIDFFLEX11  interface{} `json:"FDETAILID__FFLEX11,omitempty"`
	FDETAILIDFF100002 interface{} `json:"FDETAILID__FF100002,omitempty"`
	FDETAILIDFF101001 interface{} `json:"FDETAILID__FF101001,omitempty"`
	FDETAILIDFF101501 interface{} `json:"FDETAILID__FF101501,omitempty"`
	FDETAILIDFFLEX4   interface{} `json:"FDETAILID__FFLEX4,omitempty"`
	FDETAILIDFFLEX12  interface{} `json:"FDETAILID__FFLEX12,omitempty"`
	FDETAILIDFFLEX5   interface{} `json:"FDETAILID__FFLEX5,omitempty"`
	FDETAILIDFFLEX13  interface{} `json:"FDETAILID__FFLEX13,omitempty"`
	FDETAILIDFFLEX6   interface{} `json:"FDETAILID__FFLEX6,omitempty"`
	FDETAILIDFF101002 interface{} `json:"FDETAILID__FF101002,omitempty"`
	FDETAILIDFFLEX7   interface{} `json:"FDETAILID__FFLEX7,omitempty"`
	FDETAILIDFF102501 interface{} `json:"FDETAILID__FF102501,omitempty"`
	FDETAILIDFFLEX8   interface{} `json:"FDETAILID__FFLEX8,omitempty"`
	FDETAILIDFF103001 interface{} `json:"FDETAILID__FF103001,omitempty"`
	FDETAILIDFF102001 interface{} `json:"FDETAILID__FF102001,omitempty"`
	FDETAILIDFF100501 interface{} `json:"FDETAILID__FF100501,omitempty"`
	FDETAILIDFF103501 interface{} `json:"FDETAILID__FF103501,omitempty"`
}
