package controllers

//go:generate go run github.com/matryer/moq -out moq_test_gorm.go -skip-ensure -pkg controllers ../gateways/repositories DBer

//go:generate go run github.com/matryer/moq -out moq_test_repo.go -skip-ensure -pkg controllers ../../usecases UserGetter UserAdder TaskLister TaskAdder TokenGenerator UserNameGetter
