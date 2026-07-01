---
type: pull_request
number: 416
title: "Fixes 2717: add config to turn off notifs"
state: merged
author: jlsherrill
created: 2023-10-05T12:44:23Z
updated: 2023-10-11T14:30:23Z
closed: 2023-10-11T14:02:06Z
merged: 2023-10-11T14:02:06Z
base_branch: main
head_branch: 2717
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/416
---

# Pull Request #416: Fixes 2717: add config to turn off notifs

**Author**: @jlsherrill
**Created**: October 05, 2023 at 12:44 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2717`

## Description

## Summary
so we can disable it in prod

## Testing steps

in one terminal: 
```
make kafka-topic-consume KAFKA_TOPIC=platform.notifications.ingress
```
in another:
``` go run cmd/content-sources/main.go api consumer ```

Then create a repo for introspection:
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/repositories/" \
    -H "Content-Type: application/json" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiMSIsImludGVybmFsIjp7Im9yZ19pZCI6IjEifX19Cg==" \
    -H "x-Rh-Insights-Request-Id: 9876" \
    -d "{
          \"name\": \"pulp3.16-8\",
          \"url\": \"http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/\",
          \"snapshot\": false
        }"
```

You should see no messages comes through.   Finally, turn it on:

``` OPTIONS_ENABLE_NOTIFICATIONS=true  go run cmd/content-sources/main.go api consumer ```

create and introspect a new repo, now see messages

---

## Discussion

### Comment by @jlsherrill on October 05, 2023 at 12:47 PM UTC

we want to merge this https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/83144 before this PR

### Comment by @jlsherrill on October 05, 2023 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-2717

### Comment by @jlsherrill on October 05, 2023 at 05:35 PM UTC

/retest

### Comment by @jlsherrill on October 11, 2023 at 02:02 PM UTC

MR to app-interface is merged, merging this

### Comment by @jlsherrill on October 11, 2023 at 02:02 PM UTC

(as no qe is needed as per discussion)

---

## Reviews

### Review by @Andrewgdewar - Approved on October 10, 2023 at 09:29 PM UTC

Config confirmed as working both on and off.
Make sure to turn on notifications for stage!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/416*
