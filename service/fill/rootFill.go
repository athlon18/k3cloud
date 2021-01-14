package _fill

import _struct "github.com/athlon18/k3cloud/service/struct"

type root struct {
	root _struct.Root
}

func RootInit() *root {
	root := &root{}
	root.root.NeedUpDateFields = []string{}
	root.root.NeedReturnFields = []string{}
	root.root.IsDeleteEntry = true
	root.root.IsVerifyBaseDataField = false
	root.root.IsEntryBatchFill = true
	root.root.ValidateFlag = true
	root.root.NumberSearch = true
	root.root.IsAutoSubmitAndAudit = false
	return root
}

func (this *root) NeedUpDateFields(needUpDateFields ...string) *root {
	this.root.NeedUpDateFields = needUpDateFields
	return this
}
func (this *root) NeedReturnFields(needReturnFields ...string) *root {
	this.root.NeedReturnFields = needReturnFields
	return this
}

func (this *root) Get() _struct.Root {
	return this.root
}
