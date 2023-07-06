package main

import "fmt"

type ErrCode int

func (e ErrCode) Error() string {
	return fmt.Sprintf("error code %d", e)
}

var ERROR_CODE_404 = ErrCode(404)
var ERROR_CODE_500 = ErrCode(500)

func GetErrMsg(e error) string {
	return e.Error()
}

func main() {
	fmt.Println(GetErrMsg(ERROR_CODE_404))
	fmt.Println(GetErrMsg(ERROR_CODE_500))
}
