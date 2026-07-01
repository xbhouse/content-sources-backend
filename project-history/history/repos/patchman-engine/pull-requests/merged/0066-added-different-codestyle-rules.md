---
type: pull_request
number: 66
title: "Added different codestyle rules"
state: merged
author: josef-hak
created: 2020-01-16T16:27:18Z
updated: 2020-01-17T12:33:25Z
closed: 2020-01-17T10:36:00Z
merged: 2020-01-17T10:36:00Z
base_branch: master
head_branch: lint-config
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/66
---

# Pull Request #66: Added different codestyle rules

**Author**: @josef-hak
**Created**: January 16, 2020 at 04:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `lint-config`

## Description

    - bodyclose # checks whether HTTP response body is closed successfully [fast: true, auto-fix: false]
    - depguard # Go linter that checks if package imports are in a list of acceptable packages [fast: true, auto-fix: false]
    - dupl # Tool for code clone detection [fast: true, auto-fix: false]
    - funlen # Tool for detection of long functions [fast: true, auto-fix: false]
    - gocognit # Computes and checks the cognitive complexity of functions [fast: true, auto-fix: false]
    - goconst # Finds repeated strings that could be replaced by a constant [fast: true, auto-fix: false]
    - gocritic # The most opinionated Go source code linter [fast: true, auto-fix: false]
    - gocyclo # Computes and checks the cyclomatic complexity of functions [fast: true, auto-fix: false]
    - gofmt # Gofmt checks whether code was gofmt-ed. By default this tool runs with -s option to check for code simplification [fast: true, auto-fix: true]
    - golint # Golint differs from gofmt. Gofmt reformats Go source code, whereas golint prints out style mistakes [fast: true, auto-fix: false]
    - gosec # (gas): Inspects source code for security problems [fast: true, auto-fix: false]
    - interfacer # Linter that suggests narrower interface types [fast: true, auto-fix: false]
    - lll # Reports long lines [fast: true, auto-fix: false]
    - maligned # Tool to detect Go structs that would take less memory if their fields were sorted [fast: true, auto-fix: false]
    - misspell # Finds commonly misspelled English words in comments [fast: true, auto-fix: true]
    - nakedret # Finds naked returns in functions greater than a specified function length [fast: true, auto-fix: false]
    - prealloc # Finds slice declarations that could potentially be preallocated [fast: true, auto-fix: false]
    - scopelint # Scopelint checks for unpinned variables in go programs [fast: true, auto-fix: false]
    - stylecheck # Stylecheck is a replacement for golint [fast: true, auto-fix: false]
    - unconvert # Remove unnecessary type conversions [fast: true, auto-fix: false]
    - unparam # Reports unused function parameters [fast: true, auto-fix: false]
    - whitespace # Tool for detection of leading and trailing whitespace [fast: true, auto-fix: true]

---

## Reviews

### Review by @semtexzv - Commented on January 17, 2020 at 08:51 AM UTC

### Review by @josef-hak - Commented on January 17, 2020 at 08:59 AM UTC

### Review by @semtexzv - Approved on January 17, 2020 at 10:35 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/66*
