[![Build Status](https://travis-ci.org/tkeech1/golambda_helper.svg?branch=master)](https://travis-ci.org/tkeech1/golambda_helper)
[![codecov](https://codecov.io/gh/tkeech1/golambda_helper/branch/master/graph/badge.svg)](https://codecov.io/gh/tkeech1/golambda_helper)
[![Go Report Card](https://goreportcard.com/badge/github.com/tkeech1/golambda_helper)](https://goreportcard.com/report/github.com/tkeech1/golambda_helper)
[![CircleCI](https://circleci.com/gh/tkeech1/golambda_helper.svg?style=svg)](https://circleci.com/gh/tkeech1/golambda_helper)

## golambda_helper
Helper functions for AWS Go Lambda

# Run tests in Docker
```
make test-local
```

# Use mockery to mock AWS services
```
~/go/bin/mockery -name=DynamoInterface
```
