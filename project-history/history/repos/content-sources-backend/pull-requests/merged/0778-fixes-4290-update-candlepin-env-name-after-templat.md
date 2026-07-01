---
type: pull_request
number: 778
title: "Fixes 4290: update candlepin env name after template rename"
state: merged
author: dominikvagner
created: 2024-08-16T08:45:27Z
updated: 2024-08-16T15:32:11Z
closed: 2024-08-16T15:32:11Z
merged: 2024-08-16T15:32:10Z
base_branch: main
head_branch: 4290
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/778
---

# Pull Request #778: Fixes 4290: update candlepin env name after template rename

**Author**: @dominikvagner
**Created**: August 16, 2024 at 08:45 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4290`

## Description

## Summary
When template name is updated, the associated environment name in Candlepin isn't being updated. This PR fixes that.
Added a new `RenameEnvironment()` method to our Candlepin client. This method is called from the `update_template_content` task, which is started after updating the template, when the names are different.

## Testing steps
_Change integration test passes._
_Manual test:_
#### 1. Create a new template:
```sh
http :8000/api/content-sources/v1.0/templates/ "$( ./scripts/header.sh acme 1111)" arch='x86_64' version='8' repository_uuids:='[]' description='testing rename' name='test' use_latest:=true
```

#### 2. Verify the template and the environment with the same name exist:
```sh
http :8000/api/content-sources/v1.0/templates/{_New template ID_} "$( ./scripts/header.sh acme 1111)"
http -a admin:admin :8181/candlepin/environments/{_New template ID without hyphens_}
```

#### 3. Update the name of the template:
```sh
http PATCH :8000/api/content-sources/v1.0/templates/{_New template ID_}  "$( ./scripts/header.sh acme 1111)" name='renamed'
```

#### 4. Verify the name has changed both in the template and the associated environment:
```sh
http :8000/api/content-sources/v1.0/templates/{_New template ID_} "$( ./scripts/header.sh acme 1111)"
http -a admin:admin :8181/candlepin/environments/{_New template ID without hyphens_}
```

---

## Discussion

### Comment by @jlsherrill on August 16, 2024 at 09:00 AM UTC

https://issues.redhat.com/browse/HMS-4290

---

## Reviews

### Review by @jlsherrill - Approved on August 16, 2024 at 03:32 PM UTC

worked great and code looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/778*
