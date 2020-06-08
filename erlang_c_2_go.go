package erlang_c_2_go

import (
	"fmt"
	"io/ioutil"
	"rogchap.com/v8go"
)

func GetNumberOfAgents(
	volume int,
	intervalLength int,
	aht int,
	targetServiceLevel float32,
	targetTime int,
	maxOccupancy float32,
	shrinkage float32,
) (string, error) {
	erlangc, _ := ioutil.ReadFile("erlang-c.js")
	ctx, _ := v8go.NewContext(nil)
	_, err := ctx.RunScript(string(erlangc), "erlang-c.js")

	_, err = ctx.RunScript(fmt.Sprintf("const result = ErlangC.getNumberOfAgents(%v,%v,%v,%v,%v,%v,%v)", volume, intervalLength, aht, targetServiceLevel, targetTime, maxOccupancy, shrinkage), "main.js")
	if err != nil {
		return "", err
	}

	val, err := ctx.RunScript("result", "value.js")
	if err != nil {
		return "", err
	}

	return val.String(), nil
}
