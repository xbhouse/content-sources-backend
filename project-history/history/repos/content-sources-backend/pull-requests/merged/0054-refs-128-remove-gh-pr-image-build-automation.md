---
type: pull_request
number: 54
title: "Refs 128: remove GH pr image build automation"
state: merged
author: jlsherrill
created: 2022-07-18T18:54:33Z
updated: 2022-07-19T20:20:31Z
closed: 2022-07-19T15:29:48Z
merged: 2022-07-19T15:29:48Z
base_branch: main
head_branch: 128
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/54
---

# Pull Request #54: Refs 128: remove GH pr image build automation

**Author**: @jlsherrill
**Created**: July 18, 2022 at 06:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `128`

## Description

and other changes to fix current actions
* pins linter for current breakage: https://github.com/golangci/golangci-lint-action/issues/519
* updates dependencies
* new swag caused openapi json check to fail, updates openapi.json to match

---

## Discussion

### Comment by @jlsherrill on July 18, 2022 at 07:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-128

### Comment by @jlsherrill on July 18, 2022 at 08:44 PM UTC

Built and pushed image for 1ec3041a0455888236e125224118913a86dabfa0

### Comment by @jlsherrill on July 18, 2022 at 08:46 PM UTC

This can't work with github actions, so adding my own automation behind the scenes

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14751782

---

## Reviews

### Review by @jlsherrill - Commented on July 18, 2022 at 07:17 PM UTC

### Review by @jlsherrill - Commented on July 18, 2022 at 08:25 PM UTC

### Review by @Andrewgdewar - Approved on July 19, 2022 at 03:26 PM UTC

Much fix! 
Confirmed that its working, LGTM.

### Review by @swadeley - Approved on July 19, 2022 at 03:29 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/54*
