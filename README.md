# k3cloud
go版 金蝶 K3cloud API SDK

注: 目前只有凭证写入
```go
import (
    "github.com/athlon18/k3cloud/service"
    _struct "github.com/athlon18/k3cloud/service/struct"
)
const url = "xxx"
const accID = "xxx"

const username = "xxx"
const password = "xxx"

func main() {
    k3cloud := service.K3configInit(url, accID, username, password)
    k3Api := k3cloud.Login()
    k3Api.Get() //登录信息
    var root _struct.Root
    k3Api.SaveVoucher(root).Get() //凭证写入
    // ...
}


```