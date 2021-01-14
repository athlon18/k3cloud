package voucher

import (
	"github.com/athlon18/k3cloud/service/struct/voucher"
	"github.com/athlon18/k3cloud/service/util"
)

type fDetailID struct {
	fDetailID voucher.FDetailID
}

func FDetailIDInit() *fDetailID {
	fDetailID := &fDetailID{}
	return fDetailID
}

func (this *fDetailID) FDETAILIDFFLEX9(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX9 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX10(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX10 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX11(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX11 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF100002(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF100002 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF101001(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF101001 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF101501(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF101501 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX4(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX4 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX12(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX12 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX5(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX5 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX13(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX13 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX6(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX6 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF101002(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF101002 = util.FNumber(str)
	return this
}
func (this *fDetailID) FDETAILIDFFLEX7(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX7 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF102501(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF102501 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFFLEX8(str string) *fDetailID {
	this.fDetailID.FDETAILIDFFLEX8 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF103001(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF103001 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF102001(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF102001 = util.FNumber(str)
	return this
}

func (this *fDetailID) FDETAILIDFF100501(str string) *fDetailID {
	this.fDetailID.FDETAILIDFF100501 = util.FNumber(str)
	return this
}

func (this *fDetailID) Get() voucher.FDetailID {
	return this.fDetailID
}
