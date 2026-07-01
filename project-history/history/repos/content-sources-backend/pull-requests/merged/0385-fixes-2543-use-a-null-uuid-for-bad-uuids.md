---
type: pull_request
number: 385
title: "Fixes 2543: use a null uuid for bad uuids"
state: merged
author: jlsherrill
created: 2023-09-08T19:29:24Z
updated: 2023-09-18T21:10:34Z
closed: 2023-09-18T21:10:31Z
merged: 2023-09-18T21:10:31Z
base_branch: main
head_branch: 2543
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/385
---

# Pull Request #385: Fixes 2543: use a null uuid for bad uuids

**Author**: @jlsherrill
**Created**: September 08, 2023 at 07:29 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2543`

## Description

## Summary

instead of converting to text in queries, as
that has a performance penalty

## Testing steps

Tests pass

Here's the explain plan comparing two queries, the one with text(uuid) has a MUCH higher 'cost':
```
content=# explain SELECT * FROM repository_configurations WHERE (text(UUID) = '59a54071-06fb-45b6-b721-09bd0fb767cf' AND ORG_ID = '16791776') AND repository_configurations.deleted_at IS NULL ORDER BY repository_configurations.uuid LIMIT 1
;
                                                        QUERY PLAN                                                         
---------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=3529.23..3529.23 rows=1 width=172)
   ->  Sort  (cost=3529.23..3529.23 rows=1 width=172)
         Sort Key: uuid
         ->  Bitmap Heap Scan on repository_configurations  (cost=3197.68..3529.22 rows=1 width=172)
               Recheck Cond: (((org_id)::text = '16791776'::text) AND (deleted_at IS NULL))
               Filter: ((uuid)::text = '59a54071-06fb-45b6-b721-09bd0fb767cf'::text)
               ->  Bitmap Index Scan on repo_config_repo_org_id_deleted_null_unique  (cost=0.00..3197.68 rows=100 width=0)
                     Index Cond: ((org_id)::text = '16791776'::text)
(8 rows)
content=# explain SELECT * FROM repository_configurations WHERE (UUID = '59a54071-06fb-45b6-b721-09bd0fb767cf' AND ORG_ID = '16791776') AND repository_configurations.deleted_at IS NULL ORDER BY repository_configurations.uuid LIMIT 1
;
                                                       QUERY PLAN                                                       
------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=0.42..8.44 rows=1 width=172)
   ->  Index Scan using repository_configurations_pkey on repository_configurations  (cost=0.42..8.44 rows=1 width=172)
         Index Cond: (uuid = '59a54071-06fb-45b6-b721-09bd0fb767cf'::uuid)
         Filter: ((deleted_at IS NULL) AND ((org_id)::text = '16791776'::text))
(4 rows)
```

---

## Discussion

### Comment by @jlsherrill on September 08, 2023 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-2543

### Comment by @swadeley on September 11, 2023 at 01:01 PM UTC

/retest

### Comment by @jlsherrill on September 18, 2023 at 04:30 PM UTC

/retest

### Comment by @jlsherrill on September 18, 2023 at 09:10 PM UTC

![https://media.giphy.com/media/f1YGk2o7rVqMtTrqAY/giphy.gif](https://media.giphy.com/media/f1YGk2o7rVqMtTrqAY/giphy.gif)

---

## Reviews

### Review by @Andrewgdewar - Commented on September 15, 2023 at 08:33 PM UTC

### Review by @jlsherrill - Commented on September 18, 2023 at 02:28 PM UTC

### Review by @jlsherrill - Commented on September 18, 2023 at 02:43 PM UTC

### Review by @Andrewgdewar - Approved on September 18, 2023 at 08:38 PM UTC

![ackd](https://github.com/content-services/content-sources-backend/assets/38083295/7a948a80-2ce6-4342-92d9-46941538cf6d)


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/385*
