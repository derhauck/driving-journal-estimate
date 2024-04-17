BASE_URL :=https://driving.kateops.com
CONTAINER_TOOL ?= docker
IMG ?= ghcr.io/derhauck/dje:latest

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: updatek3s
updatek3s: docker-build
	docker save --output img.tar ${IMG}
	sudo k3s ctr images import ${PWD}/img.tar

.PHONY: deploy
deploy: updatek3s
	kubectl delete -f ./../deployment/dje.kubernetes.yaml
	kubectl apply -f ./../deployment/dje.kubernetes.yaml

.PHONY: create_open_api_schema
create_open_api_schema:
	# https://github.com/parvez3019/go-swagger3
	docker run --user $$(id -u):$$(id -g)  -t --rm -v ${PWD}:/app -w /app parvez3019/go-swagger3:latest --module-path . --output oas.json --schema-without-pkg --generate-yaml true

.PHONY: schema_test
schema_test:
	docker run  --user $$(id -u):$$(id -g) -v ${PWD}:/app -w /app schemathesis/schemathesis:stable run --checks all oas.yml --workers 8 --base-url $(BASE_URL) --cassette-path cas.yaml

.PHONY: run_tests
run_tests: create_open_api_schema schema_test