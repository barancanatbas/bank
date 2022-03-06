package api

import (
	"mongoexample/config"
	"mongoexample/handler"
	"mongoexample/repository"
)

func BankInit() handler.HandlerBank {
	ctx := config.Init()

	repo := repository.Bank(config.Db, ctx)
	hadler := handler.Bank(repo)

	return hadler
}
