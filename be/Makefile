init:
	go install go.uber.org/mock/mockgen@v0.5.0
	go get github.com/onsi/ginkgo/v2
	go install github.com/onsi/ginkgo/v2/ginkgo

gen-mock:
	mockgen -destination internal/mock/repository/reserved_money.go -package=repository_mock -source=internal/app/repository/reserved_money.go ReservedMoneyRepository
	mockgen -destination internal/mock/repository/product.go -package=repository_mock -source=internal/app/repository/product.go ProductRepository
	mockgen -destination internal/mock/repository/user.go -package=repository_mock -source=internal/app/repository/user.go UserRepository