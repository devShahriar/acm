package test

import (
	"fmt"
	"testing"

	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func TestParseJWTToken(t *testing.T) {

	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyMy0xMS0wMVQwMjozNzoyNC4zNzIzNzkrMDY6MDAiLCJlbWFpbCI6InNpYW1AZ21haWwuY29tIiwiaWQiOiI4OTg5ODgiLCJtZW1iZXJfaWQiOiIzN2NhNTNjOS0wMDc3LTQ4YzQtYTc5OC01ZDg1MmFmYWI4ZmUiLCJvcmdfaWQiOiI4YmNiMjdlOC0wYWFhLTQ5MTEtOTZjMC03ZTUyZTlhMWYwYzQiLCJwZXJtaXNzaW9ucyI6W3siaWQiOjEsIm5hbWUiOiJnZXRBcGlLZXlzIn0seyJpZCI6MiwibmFtZSI6ImdlbmVyYXRlQXBpS2V5cyJ9LHsiaWQiOjMsIm5hbWUiOiJhZGRNZW1iZXJzIn1dLCJyb2xlX2lkIjoxfQ.u9cBfWOGufCwdMwrnIMuvOUssHm3aH5y4evOIgtiU5c"

	signature := config.GetAppConfig().JWTSecretKey

	res, err := helper.ParseJWTToken(testToken, []byte(signature))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(res)
}
