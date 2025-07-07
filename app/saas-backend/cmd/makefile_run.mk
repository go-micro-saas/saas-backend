.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/saas-backend/cmd/store-configuration/... -conf=./app/saas-backend/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/saas-backend/cmd/database-migration/... -conf=./app/saas-backend/configs

.PHONY: run-saas-backend
# run service :-->: run saas-backend
run-saas-backend:
	go run ./app/saas-backend/cmd/saas-backend/... -conf=./app/saas-backend/configs

.PHONY: testing-saas-backend
# testing service :-->: testing saas-backend
testing-saas-backend:
	@echo "==> testing-saas-backend"


.PHONY: run-service
# run service :-->: run saas-backend
run-service:
	#@$(MAKE) run-saas-backend
	go run ./app/saas-backend/cmd/saas-backend/... -conf=./app/saas-backend/configs

.PHONY: testing-service
# testing service :-->: testing saas-backend
testing-service:
	$(MAKE) testing-saas-backend
