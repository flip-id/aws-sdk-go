.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make mocking-aws-ses                             Create mocking for AWS SES.'
	@echo '    make coverage-services-ses                       Create coverage for unit test Service SES.'
	@echo

.PHONY: coverage-services-ses
coverage-services-ses:
	go test -race -coverprofile=coverage-reports.out ./services/ses/...
	go tool cover -html=coverage-reports.out