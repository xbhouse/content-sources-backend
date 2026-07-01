---
type: pull_request
number: 945
title: "HMS-5348: ensure order on module stream search"
state: merged
author: jlsherrill
created: 2025-01-17T19:14:00Z
updated: 2025-01-17T21:44:41Z
closed: 2025-01-17T21:12:36Z
merged: 2025-01-17T21:12:36Z
base_branch: main
head_branch: HMS-5348
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/945
---

# Pull Request #945: HMS-5348: ensure order on module stream search

**Author**: @jlsherrill
**Created**: January 17, 2025 at 07:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-5348`

## Description

## Summary

we have to aggregate the the streams into modules, and are using a map. This doesn't preserve module ordering though, so this change ensures that the name order is consistent. 


## Testing steps

run the  dao/TestSearchRepositoryModuleStreams test 10 times and see no failures



---

## Discussion

### Comment by @jlsherrill on January 17, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5348

### Comment by @mayurilahane on January 17, 2025 at 09:12 PM UTC

testing it in stage 

### Comment by @jlsherrill on January 17, 2025 at 09:21 PM UTC

snipped

### Comment by @mayurilahane on January 17, 2025 at 09:41 PM UTC

Verified on stage by repeatedly executing ascending and descending commands, yielding accurate results.


### Comment by @mayurilahane on January 17, 2025 at 09:44 PM UTC

cmd's used - 

list modules in descending order:

` curl  --proxy http://squid.corp.redhat.com:3128/  https://console.stage.redhat.com/api/content-sources/v1.0/module_streams/search    -u content7:changeme     -d '{"urls":["https://cdn.redhat.com/content/dist/rhel8/8.5/x86_64/appstream/os"], "sort_by":"name:desc"}'  -H "Content-Type: application/json" | jq '.[].module_name'
`

list in ascending:

`
curl  --proxy http://squid.corp.redhat.com:3128/  https://console.stage.redhat.com/api/content-sources/v1.0/module_streams/search    -u content7:changeme     -d '{"urls":["https://cdn.redhat.com/content/dist/rhel8/8.5/x86_64/appstream/os"], "sort_by":"name:asc"}'  -H "Content-Type: application/json" | jq '.[].module_name'
`

---

## Reviews

### Review by @xbhouse - Approved on January 17, 2025 at 07:38 PM UTC

test passes consistently 👍 lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/945*
