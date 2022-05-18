package fakemesh

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {

	//x, err := CreateToken(100003, 100004, int64(0))
	//if err != nil {
	//	println(fmt.Sprintf("%s", err))
	//	return
	//}
	//
	//// xYZA0L/bYC6CyDOZZF18vj026tFtTXXiQYHOd8EqfWs=
	//
	//println(fmt.Sprintf("Encrypt: %s", x))
	uc, qc, err := CheckToken("CQmzs8xf/hSSugETy7d5QUM2WEmI7k9TmUg7tW1EUFbVW6orJPiYPv7+XEIn/zU=")
	if err != nil {
		println(fmt.Sprintf("error: %s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d, qc: %d", uc, qc))

}

func TestCheckToken(t *testing.T) {
	uc, qc, err := CheckToken("zFBOTz7tZQw3BEBLAm+7TeJoqvaELIYhnXs+Ne6+WqEqOrYAUN7hEQkkB2CYQq46")
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d, qc: %d", uc, qc))
}
