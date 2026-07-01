---
type: pull_request
number: 877
title: "Fixes 4775: add weekly cron for pulp orphan cleanup"
state: merged
author: rverdile
created: 2024-11-06T21:38:59Z
updated: 2024-11-13T09:44:17Z
closed: 2024-11-13T09:44:16Z
merged: 2024-11-13T09:44:16Z
base_branch: main
head_branch: orphan-cleanup
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/877
---

# Pull Request #877: Fixes 4775: add weekly cron for pulp orphan cleanup

**Author**: @rverdile
**Created**: November 06, 2024 at 09:38 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `orphan-cleanup`

## Description

## Summary
Adds a new command and weekly cron job to run orphan cleanup for pulp. Runs in batches of 5 orgs.

## Testing steps
1. You can use this script to create a bunch of orgs

```
TOTAL_ORGS=33

for ((i=0; i<$TOTAL_ORGS; ++i))
do

ORG_ID=`tr -dc A-Za-z0-9 </dev/urandom | head -c 13`

HEADER=`./scripts/header.sh $ORG_ID`

UUID=$(curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "${HEADER}" \
    -H "Content-Type: application/json" \
    -d '{
          "name": "comps repo 2",
          "url": "https://rverdile.fedorapeople.org/dummy-repos/comps/repo2/",
          "snapshot": true
  }' | jq -r .uuid)

curl -X DELETE --location "http://localhost:8000/api/content-sources/v1.0/repositories/$UUID" \
    -H "${HEADER}" \
    -H "Content-Type: application/json"
done
```
2. Run `go cmd/external-repos/main.go pulp-orphan-cleanup`
3. In the logs you'll see the pulp tasks running in groups of 5


---

## Discussion

### Comment by @jlsherrill on November 06, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-4775

---

## Reviews

### Review by @dominikvagner - Commented on November 07, 2024 at 12:58 PM UTC

### Review by @rverdile - Commented on November 07, 2024 at 02:07 PM UTC

### Review by @dominikvagner - Commented on November 11, 2024 at 09:22 AM UTC

tested it out with adding a wait time to the testing script between the repo creation and it's deletion and then it works as expected 👍🏼 
overall it looks good! just remove that script and it's good to go, nice job! 😄😇 

### Review by @rverdile - Commented on November 11, 2024 at 03:33 PM UTC

### Review by @dominikvagner - Approved on November 12, 2024 at 08:49 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/877*
