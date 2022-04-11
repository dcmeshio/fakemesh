package fakemesh

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {

	x, err := CreateToken(100003, 100004, int64(0))
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}

	// xYZA0L/bYC6CyDOZZF18vj026tFtTXXiQYHOd8EqfWs=

	println(fmt.Sprintf("Encrypt: %s", x))
	uc, qc, err := CheckToken("ewRzOM/dZ9D0MQKIG8tls7U1U4DRBIHKdfUOHYXanDmts96vK5mFurgr4hzZHquc")
	if err != nil {
		println(fmt.Sprintf("%s", err))
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
