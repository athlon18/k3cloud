package service

import (
	"fmt"
	_fill "github.com/athlon18/k3cloud/service/fill"
	voucher2 "github.com/athlon18/k3cloud/service/fill/voucher/save"
	"testing"
)
const url = "xxx"
const accID = "xxx"

const username = "xxx"
const password = "xxx"

func TestApi(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常:", err)
		}
	}()
	root := _fill.
		RootInit().
		NeedReturnFields("FVOUCHERGROUPNO").
		Get()
	model := voucher2.
		ModelInit().
		FDate("2020-12-12 00:00:00").
		FYEAR("2020").
		FPERIOD("12").
		FVOUCHERGROUPID("PRE001").
		FSourceBillKey("78050206-2fa6-40e3-b7c8-bd608146fa38").
		FAccountBookID("102").
		Get()
	//mapList := make(map[string]string, 0)
	//for _, _ = range mapList {
	fEntity1 := voucher2.
		FEntityInit().
		FEXPLANATION("测试").
		FACCOUNTID("1221.01.01").
		FCURRENCYID("PRE001").
		FEXCHANGERATETYPE("HLTX01_SYS").
		FEXCHANGERATE("1").
		FAMOUNTFOR("500").
		FDEBIT("500").
		Get()
	model.FEntity = append(model.FEntity, fEntity1)

	fEntity2 := voucher2.
		FEntityInit().
		FEXPLANATION("测试2").
		FACCOUNTID("1002").
		FCURRENCYID("PRE001").
		FEXCHANGERATETYPE("HLTX01_SYS").
		FEXCHANGERATE("1").
		FAMOUNTFOR("500").
		FCREDIT("500").
		Get()
	fEntity2.FDetailID = voucher2.
		FDetailIDInit().
		FDETAILIDFF100002("03383000040013701").
		FDETAILIDFF101002("YH006").
		Get()
	model.FEntity = append(model.FEntity, fEntity2)
	//}

	root.Model = model
	//bb, _ := json.Marshal(root)
	//fmt.Println(string(bb))
	K3configInit(url, accID, username, password).
		Login().
		SaveVoucher(root).
		Print()
}
