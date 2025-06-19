#!/bin/sh

echo "Running Go linters before push..."

# Run golangci-lint
golangci-lint run --issues-exit-code=1

# Check the exit code of the linter
if [ $? -ne 0 ]; then
	echo "Linting failed. Please fix the issues before pushing."
	exit 1
fi

echo "Linting passed. Proceeding with push."
exit 0