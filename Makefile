GENERATOR=/opt/atricore/tools/openapi-generator-cli/openapi-generator-cli.sh
#GENERATOR=/data/atricore/tools/openapi-generator-cli/openapi-generator-cli.sh
#SWAGGER_FILE=~/.m2/repository/com/atricore/idbus/console/console-api/1.4.3-SNAPSHOT/console-api-1.4.3-SNAPSHOT-swagger.yaml
#SWAGGER_FILE=./console-api-1.4.3-SNAPSHOT-swagger.yaml

SWAGGER_FILE=./console-api-1.5.0-SNAPSHOT-swagger.json

PGK_NAME=jossoappi

default: build

build:
	go install

dep: # Download required dependencies
	go mod tidy
	go mod vendor

test:
	go test

generate:
	$(GENERATOR) generate -i $(SWAGGER_FILE) -g go -o . --additional-properties=packageName=$(PGK_NAME) --additional-properties=disallowAdditionalPropertiesIfNotPresent=false --git-repo-id=josso-api-go --git-user-id=atricore
