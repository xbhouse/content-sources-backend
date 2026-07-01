---
type: pull_request
number: 675
title: "Fixes 3600: delete distributions and env on template delete"
state: merged
author: rverdile
created: 2024-05-22T13:35:32Z
updated: 2024-05-30T13:47:40Z
closed: 2024-05-30T13:47:37Z
merged: 2024-05-30T13:47:37Z
base_branch: main
head_branch: delete-template-candlepin
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/675
---

# Pull Request #675: Fixes 3600: delete distributions and env on template delete

**Author**: @rverdile
**Created**: May 22, 2024 at 01:35 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `delete-template-candlepin`

## Description

## Summary
Adds additional actions to the DeleteTemplate task so that the associated pulp distributions and candlepin environment are also deleted.

## Testing steps
1. Create a template with a mix of red hat and custom repositories
2. You can verify the existence of the distributions by browsing to `http://pulp.content:8080/pulp/content/<domain>/templates/<template_uuid>/`
3. You can verify the existence of the environment by GET requesting `http://localhost:8181/candlepin/environments/<env id>`. The env_id is the template_uuid without the hyphens.
4. Delete the template
5. Recheck those endpoints and you should get a 404.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 22, 2024 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-3600

### Comment by @jlsherrill on May 28, 2024 at 08:23 PM UTC

I'm not seeing the distributions  getting deleted.  In fact i never saw a DELETE call to pulp, which makes me think dao.template.GetDistributionHref() is returning  ""  ?  (just a guess)  I did check the payload and it seemed to be there:

```
{"PoolID": "8ad981858fc03bd6018fc06c821e1e03", "TemplateUUID": "09e69d33-e236-4219-bdfe-b1ea2a5bc0d5", "RepoConfigUUIDs": ["73c32f6c-e1ba-4d51-b3de-8cf66f848854", "4d858709-65b2-495f-91fa-28aa7a731654", "9cb8141b-9c73-4152-9fa9-50d3ccbc3966"]}
```

### Comment by @rverdile on May 29, 2024 at 04:56 PM UTC

this should work now

### Comment by @jlsherrill on May 29, 2024 at 07:56 PM UTC

yep, seems to work great.  Would you mind adding an integration test?   Even if its just with a custom repo (no red hat repos), that would have caught this.

### Comment by @rverdile on May 29, 2024 at 08:04 PM UTC

I added to the test here: https://github.com/content-services/content-sources-backend/blob/025902696cfe7856a5acda0c005870f708459b1d/test/integration/update_template_content_test.go#L249-L269

It doesn't catch it because our integration tests don't test the API handler. It was missing the "RepositoryUUIDs" part of the payload.

### Comment by @jlsherrill on May 29, 2024 at 08:09 PM UTC

ahhh, i see!

---

## Reviews

### Review by @jlsherrill - Approved on May 29, 2024 at 08:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/675*
