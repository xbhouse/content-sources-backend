---
type: pull_request
number: 383
title: "Fixes 2319: add repo filtering by origin & type"
state: merged
author: jlsherrill
created: 2023-09-06T16:59:44Z
updated: 2024-02-26T14:48:08Z
closed: 2023-09-20T15:56:50Z
merged: 2023-09-20T15:56:50Z
base_branch: main
head_branch: 2319
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/383
---

# Pull Request #383: Fixes 2319: add repo filtering by origin & type

**Author**: @jlsherrill
**Created**: September 06, 2023 at 04:59 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2319`

## Description

## Summary

This adds two new attributes to repositories:  content_type (rpm), and origin (external, red_hat).

I adjusted the values of 'origin' from the original task as the only other 'origin' i can think of would be for 'uploaded' content.  In the future when we support uploading content, we want image builder to still be able to get only 'custom repos with external urls', so they could fetch repos with origin=external to get that.

## Testing steps

For dev:

After migrating, run this main.go file:
```
package main

import (
	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/content-services/content-sources-backend/pkg/seeds"
	"github.com/openlyinc/pointy"
)

func main() {
	db.Connect()
	d := db.DB
	seeds.SeedRepositoryConfigurations(d, 1, seeds.SeedOptions{OrgID: "1", ContentType: pointy.Pointer("rpm"), Origin: pointy.Pointer("external")})
	seeds.SeedRepositoryConfigurations(d, 1, seeds.SeedOptions{OrgID: "1", ContentType: pointy.Pointer("not-rpm"), Origin: pointy.Pointer("red_hat")})
}
```

Launch the server, with the feature disabled:
```
FEATURES_NEW_REPO_FILTERING_ENABLED=false go run cmd/content-sources/main.go  api
```

Fetch the default repo list:

```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiMSIsImludGVybmFsIjp7Im9yZ19pZCI6IjEifX19Cg=="
```

This should show only 1 repo, as its filtering by to rpm & external by default.

Launch the server with the feature enabled:
```
FEATURES_NEW_REPO_FILTERING_ENABLED=true go run cmd/content-sources/main.go  api
```

Fetch the default repo list:

```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiMSIsImludGVybmFsIjp7Im9yZ19pZCI6IjEifX19Cg=="
```

Now you should get both repos that were created as part of the main.go run


Try adding  origin=external  & content_type=rpm to the url:

```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/repositories/?origin=external" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiMSIsImludGVybmFsIjp7Im9yZ19pZCI6IjEifX19Cg=="
```

```
curl -X GET --location "http://localhost:8000/api/content-sources/v1.0/repositories/?content_type=rpm" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiMSIsImludGVybmFsIjp7Im9yZ19pZCI6IjEifX19Cg=="
```

You should get expected results.
Other things to test:

1.  trying to update the content_type or origin via PUT or PATCH should be ignored
2.  specifying content_type or origin when creating a repo should be ignored.

---

## Discussion

### Comment by @jlsherrill on September 06, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2319

### Comment by @jlsherrill on September 14, 2023 at 05:18 PM UTC

No, these values should not be able to be updated at all.   Updates should be ignored

### Comment by @jlsherrill on September 15, 2023 at 08:06 PM UTC

@swadeley this actually needs a bindings update.   The added parameters are kinda useless right now and not really worthy of testing though.

### Comment by @swadeley on September 18, 2023 at 03:23 PM UTC

Hi

I have just deployed, not built API docs yet, ran API tests and noticed this:

    # compare if repo introspection has happened in past 8 or 24 hours or not
    >           assert repo_introspection_time >= max_last_introspection_time
    E           assert datetime.datetime(2023, 9, 17, 16, 0, 7) >= datetime.datetime(2023, 9, 17, 16, 21, 36, 380229)


EDIT: After building APIdocs, error persists.
EDIT2: my mistake, disregard (wrong test target env)

### Comment by @swadeley on September 18, 2023 at 05:57 PM UTC

Hi, after building apidocs and installing the updated plugin, i get:
`TypeError: __init__() got an unexpected keyword argument 'origin'`


### Comment by @swadeley on September 20, 2023 at 02:00 PM UTC

/retest

### Comment by @swadeley on September 20, 2023 at 03:56 PM UTC

Hi, my mistake. I had to add support QE code side to accept the two new strings: origin and content_type

### Comment by @swadeley on February 22, 2024 at 05:10 PM UTC

For QE testing:

In [2]: app.content_sources.rest_client.repositories_api.list_repositories(content_type='rpm')

In [4]: app.content_sources.rest_client.repositories_api.list_repositories(origin='red_hat')

In [5]: app.content_sources.rest_client.repositories_api.list_repositories(origin='external')


---

## Reviews

### Review by @rverdile - Commented on September 12, 2023 at 08:36 PM UTC

### Review by @jlsherrill - Commented on September 14, 2023 at 02:58 PM UTC

### Review by @rverdile - Commented on September 14, 2023 at 03:52 PM UTC

When making a PUT request, without specifying origin or content_type, would expect the values to be replaced with the default values? MapForUpdate in the model would need to be updated for that to work I think.

### Review by @rverdile - Approved on September 14, 2023 at 05:37 PM UTC

lgtm!

### Review by @swadeley - Commented on September 18, 2023 at 03:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/383*
