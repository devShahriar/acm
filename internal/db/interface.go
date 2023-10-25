package db

import "magic.pathao.com/carta/carta-acm/internal/contract"

type DB interface {
	GenerateApiKey(req contract.GenerateApiKeyReq) error
}
