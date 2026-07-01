---
type: pull_request
number: 243
title: "Fixes 1568: lower log level on introspect errors"
state: merged
author: jlsherrill
created: 2023-04-05T15:48:41Z
updated: 2023-04-10T15:37:31Z
closed: 2023-04-10T15:37:28Z
merged: 2023-04-10T15:37:27Z
base_branch: main
head_branch: 1568
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/243
---

# Pull Request #243: Fixes 1568: lower log level on introspect errors

**Author**: @jlsherrill
**Created**: April 05, 2023 at 03:48 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1568`

## Description

## Summary
This separates introspection errors from internal errors to better report them separately

## Testing steps

run the server locally, and run the following request:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K" \
    -d "{
          \"name\": \"needed3\",
          \"url\": \"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata3/\",
          \"distribution_versions\": [
            \"8\"
          ],
          \"distribution_arch\": \"x86_64\"
        }"
```

Then monitor the server logs:

```
11:22AM INF Forcing introspection for 'https://jlsherrill.fedorapeople.org/fake-repos/needed-errata3/'
11:22AM WRN Error 0 introspecting repository https://jlsherrill.fedorapeople.org/fake-repos/needed-errata3/ error="Error introspecting https://jlsherrill.fedorapeople.org/fake-repos/needed-errata3/: Cannot fetch https://jlsherrill.fedorapeople.org/fake-repos/needed-errata3/repodata/repomd.xml: 404"
```

you should see a WARN message, and not a FAIL message.

Also try running: 

```
go run cmd/external-repos/main.go introspect-all --force
```

you should see:

```
{"level":"warn","time":"2023-04-05T11:47:02-04:00","message":"Introspection Error: Error introspecting https://jlsherrill.fedorapeople.org/fake-repos/needed-errata2/: Cannot fetch https://jlsherrill.fedorapeople.org/fake-repos/needed-errata2/repodata/repomd.xml: 404"}

```
which is now warn


---

## Discussion

### Comment by @jlsherrill on April 05, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1568

---

## Reviews

### Review by @rverdile - Commented on April 06, 2023 at 04:18 PM UTC

### Review by @rverdile - Commented on April 06, 2023 at 04:18 PM UTC

### Review by @jlsherrill - Commented on April 06, 2023 at 05:41 PM UTC

### Review by @rverdile - Approved on April 06, 2023 at 07:11 PM UTC

lgtm!

### Review by @rverdile - Commented on April 06, 2023 at 07:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/243*
