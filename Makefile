.DEFAULT_GOAL := help

test: ## テスト実行
	go test -race -shuffle=on ./...

dry-migrate: ## マイグレーション用DLL表示
	mysqldef -u sample -p sample -h db -P 3306 sample --dry-run < ./_tools/mysql/schema.sql

migrate: ## マイグレーション実行
	mysqldef -u sample -p sample -h db -P 3306 sample < ./_tools/mysql/schema.sql

generate: ## コード生成
	go generate ./...

help: ## ヘルプ表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
