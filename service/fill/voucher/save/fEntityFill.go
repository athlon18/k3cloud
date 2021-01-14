package voucher

import (
	"github.com/athlon18/k3cloud/service/struct/voucher"
	"github.com/athlon18/k3cloud/service/util"
)

type fEntity struct {
	fEntity voucher.FEntity
}

func FEntityInit() *fEntity {
	fEntity := &fEntity{}
	return fEntity
}

func (this *fEntity) FEntryID(str string) *fEntity {
	this.fEntity.FEntryID = str
	return this
}

func (this *fEntity) FEXPLANATION(str string) *fEntity {
	this.fEntity.FEXPLANATION = str
	return this
}

func (this *fEntity) FACCOUNTID(str string) *fEntity {
	this.fEntity.FACCOUNTID = util.FNumber(str)
	return this
}

func (this *fEntity) FCURRENCYID(str string) *fEntity {
	this.fEntity.FCURRENCYID = util.FNumber(str)
	return this
}

func (this *fEntity) FEXCHANGERATETYPE(str string) *fEntity {
	this.fEntity.FEXCHANGERATETYPE = util.FNumber(str)
	return this
}

func (this *fEntity) FEXCHANGERATE(str string) *fEntity {
	this.fEntity.FEXCHANGERATE = str
	return this
}

func (this *fEntity) FUnitID(str string) *fEntity {
	this.fEntity.FUnitID = util.FNumber(str)
	return this
}

func (this *fEntity) FPrice(str string) *fEntity {
	this.fEntity.FPrice = str
	return this
}
func (this *fEntity) FQty(str string) *fEntity {
	this.fEntity.FQty = str
	return this
}

func (this *fEntity) FAMOUNTFOR(str string) *fEntity {
	this.fEntity.FAMOUNTFOR = str
	return this
}
func (this *fEntity) FDEBIT(str string) *fEntity {
	this.fEntity.FDEBIT = str
	return this
}
func (this *fEntity) FCREDIT(str string) *fEntity {
	this.fEntity.FCREDIT = str
	return this
}

func (this *fEntity) FSettleTypeID(str string) *fEntity {
	this.fEntity.FSettleTypeID = util.FNumber(str)
	return this
}
func (this *fEntity) FSETTLENO(str string) *fEntity {
	this.fEntity.FSETTLENO = str
	return this
}
func (this *fEntity) FEXPORTENTRYID(str string) *fEntity {
	this.fEntity.FEXPORTENTRYID = str
	return this
}

/**
辅助核算
*/
func (this *fEntity) FDetailID(str voucher.FDetailID) *fEntity {
	this.fEntity.FDetailID = str
	return this
}

func (this *fEntity) Get() voucher.FEntity {
	return this.fEntity
}
