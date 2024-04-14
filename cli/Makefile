.PHONY: create_open_api_schema
create_open_api_schema:
	# https://github.com/parvez3019/go-swagger3
	docker run -t --rm -v ${PWD}:/app -w /app parvez3019/go-swagger3:latest --module-path . --output oas.json --schema-without-pkg --generate-yaml true