
#!/bin/sh
COVERAGE_FILE=.profile.cov
go test -v -coverpkg=./... -coverprofile=${COVERAGE_FILE} ./... && go tool cover -func ${COVERAGE_FILE}