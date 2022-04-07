package fakec

import (
	"fmt"
	"github.com/dcmeshio/fakemesh"
	"github.com/dcmeshio/fakemesh/common"
	"io"
)

// 外部获取 Writer、Host、CreateTokenParam
func Send(writer io.Writer, host string, uc, qc int) error {
	option := fakemesh.GetOption()
	builder := CreateBuilder()
	// 首行
	index := common.Rint(len(option.Lines))
	Line := option.Lines[index]
	ParamValue := common.Rint8()
	FirstLine := fmt.Sprintf("GET %s?%s=%s HTTP/1.1", Line.PathName, Line.ParamName, ParamValue)
	builder.AddFirstLine(FirstLine)
	// 请求头 Host
	Host := fmt.Sprintf("Host: %s", host)
	builder.AddHeader(Host)
	// 请求头 Token
	tv, err := fakemesh.CreateToken(uc, qc, option.TimestampDifference)
	if err != nil {
		return err
	}
	th := ""
	if option.Type == 0 {
		th = "X-token"
	} else {
		th = "Ps"
	}
	Token := fmt.Sprintf("%s: %s", th, tv)
	builder.AddHeader(Token)
	// 请求头 Other
	for _, v := range option.RequestHeaders {
		Header := ""
		if v.Single {
			Header = fmt.Sprintf("%s: %s", v.Name, v.Value)
		} else {
			index = common.Rint(len(v.Values))
			Header = fmt.Sprintf("%s: %s", v.Name, v.Values[index])
		}
		builder.AddHeader(Header)
	}

	bytes := builder.Build()
	_, err = writer.Write(bytes)
	return err
}
