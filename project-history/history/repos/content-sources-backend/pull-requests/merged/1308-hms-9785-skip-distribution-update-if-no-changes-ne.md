---
type: pull_request
number: 1308
title: "HMS-9785: skip distribution update if no changes needed"
state: merged
author: jlsherrill
created: 2025-11-20T21:50:59Z
updated: 2025-11-20T22:54:29Z
closed: 2025-11-20T22:54:20Z
merged: 2025-11-20T22:54:20Z
base_branch: main
head_branch: 9785
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1308
---

# Pull Request #1308: HMS-9785: skip distribution update if no changes needed

**Author**: @jlsherrill
**Created**: November 20, 2025 at 09:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `9785`

## Description

## Summary

Skips triggering a distribution update task if neither the content guard href, nor the publication href have changed.  The name shouldn't ever change, and we identify them by their path, so it wouldn't make sense to be different than expected.

## Testing steps

With this pr, launch the server and run:
```
 podman logs -f cs_pulp_api_1 | grep distribution
```

now from the UI (or api), create a repository to snapshot.  You should see two distributions get created in the pulp logs

```
('pulp [69ff09c2a77e46ea895030841e6b0fc6]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:42:34 +0000] "GET /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/?base_path=458ed86c-5000-47c2-9dab-221f9eefc209%2F3837d6f2-ef5e-4a29-9c6c-47869550a46b HTTP/1.0" 200 52 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
('pulp [45bc34dc51674931877dd8f304a58c52]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:42:34 +0000] "POST /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/ HTTP/1.0" 202 85 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
('pulp [5155cd44876849b08a6a7bc5e7748586]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:42:35 +0000] "GET /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/?base_path=458ed86c-5000-47c2-9dab-221f9eefc209%2Flatest HTTP/1.0" 200 52 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
('pulp [5d8091328565497d8e5715c4d9ad4576]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:42:35 +0000] "POST /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/ HTTP/1.0" 202 85 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
```

now force a new snapshot to be done.  Since there no changes, you should see no distribution POST request.

Now change the url to a new url, which should force a snapshot.   Now you will see one distribution getting created and then one get updated:

```
('pulp [69ddf3abf475499a93405514eb72c829]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:44:41 +0000] "GET /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/?base_path=458ed86c-5000-47c2-9dab-221f9eefc209%2Fe7265d30-012c-4b54-af80-0d37dc169b2e HTTP/1.0" 200 52 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
('pulp [9ad92475a40b4ea38cb376d7dee6a0b9]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:44:42 +0000] "POST /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/ HTTP/1.0" 202 85 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
('pulp [8c71e81361154882a8fc6584940a11f6]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:44:43 +0000] "GET /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/?base_path=458ed86c-5000-47c2-9dab-221f9eefc209%2Flatest HTTP/1.0" 200 912 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
pulp [49c22bf0c6a4454c8f6d754476e6ff0b]: pulpcore.tasking.tasks:INFO: Starting task id: 019aa33a-121e-727a-a57a-288d0d238905 in domain: cs-d6b5e6b84e, task_type: pulpcore.app.tasks.base.ageneral_update, immediate: True, deferred: True
pulp [49c22bf0c6a4454c8f6d754476e6ff0b]: pulpcore.tasking.tasks:INFO: Task completed 019aa33a-121e-727a-a57a-288d0d238905 in domain: cs-d6b5e6b84e, task_type: pulpcore.app.tasks.base.ageneral_update, immediate: True, deferred: True, execution_time: 43129 μs
('pulp [49c22bf0c6a4454c8f6d754476e6ff0b]: ::ffff:10.89.0.63 - - [20/Nov/2025:21:44:43 +0000] "PATCH /api/pulp/cs-d6b5e6b84e/api/v3/distributions/rpm/rpm/019aa338-2116-7f04-b41b-87d7ef60c745/ HTTP/1.0" 202 85 "-" "OpenAPI-Generator/386d86254354e29d3b864523aed47835d638abdf67968b5e2da9ab863deab9a5/go"',)
```



---

## Discussion

### Comment by @xbhouse on November 20, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-9785

---

## Reviews

### Review by @rverdile - Approved on November 20, 2025 at 10:37 PM UTC

Looks good. Added a repo to a use-latest template and forced a snapshot, didn't see pulp update the distribution.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1308*
