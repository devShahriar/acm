mod-tidy:
	@echo "Running go mod tidy"
	go mod tidy -v


mod-vendor:
	@echo "Running go mod vendor"
	go mod vendor -v

# Resolve dependencies
dep: mod-tidy mod-vendor