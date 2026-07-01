---
type: pull_request
number: 1174
title: "HMS-9053: log slow queries"
state: merged
author: jlsherrill
created: 2025-08-13T12:45:03Z
updated: 2025-08-14T20:44:35Z
closed: 2025-08-14T20:44:29Z
merged: 2025-08-14T20:44:29Z
base_branch: main
head_branch: 9053
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1174
---

# Pull Request #1174: HMS-9053: log slow queries

**Author**: @jlsherrill
**Created**: August 13, 2025 at 12:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `9053`

## Description

## Summary

This fixes the gorm slow query logging that was not working previously with the gorm-zerolog bridge library.  This takes gorms logger (https://github.com/go-gorm/gorm/blob/master/logger/logger.go) and copies it into our code base with changes to work with zerolog.  

This properly passes in the zerolog level in such a way that slow queries are logged at WARN level, while other logging still respects the configured log level (so all queries aren't logged unless you are at the DEBUG level)

This also adds a zerlog hook to properly add the request id to all logs

It also fixed a metric query that would trigger an error when there was no data in the db, this now handles it properly.

## Testing steps

add this config in ./config/config.yaml under 'database':
```
slow_query_duration: 5ms
```
and set log_level to 'warn':
```
logging:
  level: warn
```


run the app and run some request and you should see most queries logged as being slow:

```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJ0eXBlIjoiVXNlciIsInVzZXIiOnsidXNlcm5hbWUiOiIxIn0sImFjY291bnRfbnVtYmVyIjoiZm9vIiwiaW50ZXJuYWwiOnsib3JnX2lkIjoiMSJ9fX0K" \
    -H "x-Rh-Insights-Request-Id: 9d229d34-7219-405d-ba3d-8f99cf7c36ae" \
    -H "Content-Type: application/json" \
    -d '{"urls":["https://cdn.redhat.com/content/dist/rhel9/9/x86_64/baseos/os","https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os"],
          "include_package_sources": true,
          "search":"a"}'
```


```
9:14AM WRN /home/jlsherri/git/content-sources-backend/pkg/dao/rpms.go:213 SLOW SQL >= 20ms
[67.607ms] [rows:100] SELECT DISTINCT ON(rpms.name) rpms.name as package_name,rpms.summary,rpms.uuid FROM "rpms" inner join repositories_rpms on repositories_rpms.rpm_uuid = rpms.uuid WHERE repositories_rpms.repository_uuid in (SELECT repositories.uuid FROM "repositories" left join repository_configurations on repositories.uuid = repository_configurations.repository_uuid WHERE (repository_configurations.org_id IN ('1', '-1', '-2') OR repositories.public OR repositories.url in ('https://dl.fedoraproject.org/pub/epel/10/Everything/x86_64/','https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/','https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/')) AND (repositories.url in ('https://cdn.redhat.com/content/dist/rhel9/9/x86_64/baseos/os/','https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os/') OR repository_configurations.uuid in (NULL))) AND rpms.name ILIKE 'a%' ORDER BY rpms.name ASC LIMIT 100 request_id=9d229d34-7219-405d-ba3d-8f99cf7c36ae severity=warn
```


Now set the log level to debug, and bump up the slow query setting:

```
slow_query_duration: 200ms
```

```
logging:
  level: debug
```


Now you will see all requests logged.






---

## Discussion

### Comment by @jlsherrill on August 13, 2025 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-9053

### Comment by @jlsherrill on August 14, 2025 at 01:43 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 14, 2025 at 07:37 PM UTC

works nicely!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1174*
