package voucher

import (
	"github.com/athlon18/k3cloud/service/struct/voucher"
	"github.com/athlon18/k3cloud/service/util"
)

type model struct {
	model voucher.Model
}

func ModelInit() *model {
	model := &model{}
	model.model.FVOUCHERID = "0"
	model.model.FAccountBookID = util.FNumber("")
	model.model.FVOUCHERGROUPID = util.FNumber("PRE001")
	model.model.FSourceBillKey = util.FNumber("78050206-2fa6-40e3-b7c8-bd608146fa38")
	model.model.FVOUCHERGROUPNO = "0"
	model.model.FDocumentStatus = "Z"
	model.model.FISADJUSTVOUCHER = false
	return model
}

func (this *model) FYEAR(str string) *model {
	this.model.FYEAR = str
	return this
}

func (this *model) FPERIOD(str string) *model {
	this.model.FPERIOD = str
	return this
}

func (this *model) FDate(str string) *model {
	this.model.FDate = str
	return this
}

func (this *model) FSourceBillKey(str string) *model {
	this.model.FSourceBillKey = util.FNumber(str)
	return this
}

func (this *model) FAccountBookID(str string) *model {
	this.model.FAccountBookID = util.FNumber(str)
	return this
}

func (this *model) FVOUCHERGROUPID(str string) *model {
	this.model.FVOUCHERGROUPID = util.FNumber(str)
	return this
}
func (this *model) FATTACHMENTS(str string) *model {
	this.model.FATTACHMENTS = str
	return this
}

func (this *model) Get() voucher.Model {
	return this.model
}
