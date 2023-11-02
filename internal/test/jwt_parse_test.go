package test

import (
	"fmt"
	"testing"

	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func TestParseJWTToken(t *testing.T) {

	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyMy0xMS0wMlQxNDowMjo1MC4zMzQxMzYrMDY6MDAiLCJlbWFpbCI6InNodWRpcEBnbWFpbC5jb20iLCJpZCI6IjEyMzQ1IiwibWVtYmVyX2lkIjoiMGJiYjQzYzMtNWU5ZC00ZjAwLWI3ZWUtYjdlNmM4NDE3MWY5Iiwib3JnX2lkIjoiNmU2ZWNlNTAtYjUxYS00OTRkLWFlYzQtODRjODdiMTlhNGFhIiwicGVybWlzc2lvbnMiOlt7ImlkIjoxLCJuYW1lIjoiZ2V0TWVtYmVycyJ9LHsiaWQiOjMsIm5hbWUiOiJ1cGRhdGVSb2xlIn1dLCJyb2xlX2lkIjoxfQ.adA2iwdH3UzHJ-nnBqL2S0RjKSuhcYNaJFA_nMrgZz0"

	signature := config.GetAppConfig().JWTSecretKey

	res, err := helper.ParseJWTToken(testToken, []byte(signature))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(res)
}
