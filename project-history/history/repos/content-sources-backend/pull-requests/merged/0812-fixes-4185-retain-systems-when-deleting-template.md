---
type: pull_request
number: 812
title: "Fixes 4185: retain systems when deleting template"
state: merged
author: rverdile
created: 2024-09-11T19:51:25Z
updated: 2024-09-18T13:29:57Z
closed: 2024-09-18T13:29:53Z
merged: 2024-09-18T13:29:53Z
base_branch: main
head_branch: candlepin-systems-check
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/812
---

# Pull Request #812: Fixes 4185: retain systems when deleting template

**Author**: @rverdile
**Created**: September 11, 2024 at 07:51 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `candlepin-systems-check`

## Description

## Summary
When deleting a template, consumers associated to that template will not be deleted.

## Testing steps
1. Create a template
2. Create a consumer
```
curl -X POST --location "http://localhost:8181/candlepin/consumers/?owner=content-sources-test" \
    -H "Content-Type: application/json" \
    -H "Authorization: Basic YWRtaW46YWRtaW4=" \
    -d '{
          "name": "example-consumer",
          "type": {
            "label": "system"
          }
        }'
```
3. Associate the consumer to the template
```
go run cmd/candlepin/main.go add-system <consumer uuid> <template name>
```
4. Delete the template
5. Check that the consumer still exists
```
curl -X GET --location "http://localhost:8181/candlepin/consumers?owner=content-sources-test" \
    -H "Content-Type: application/json" \
    -H "Authorization: Basic YWRtaW46YWRtaW4="
```

You could also test this by
1. Create a template
2. Register a RHEL9 VM to the template as described in [register_client.md](https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md)
3. Delete the template
4. The system will still have the content until you refresh, afterwards the content would be gone.
5. If you associate a new template to the system and refresh, it will have the new template's content


---

## Discussion

### Comment by @jlsherrill on September 11, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4185

---

## Reviews

### Review by @jlsherrill - Approved on September 17, 2024 at 06:37 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/812*
