---
type: pull_request
number: 1537
title: "HMS-9952: Add SendTemplateUpdateEvents job"
state: merged
author: swadeley
created: 2026-06-11T14:51:03Z
updated: 2026-06-18T09:16:16Z
closed: 2026-06-18T09:16:15Z
merged: 2026-06-18T09:16:15Z
base_branch: main
head_branch: swadeley/HMS-9952
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1537
---

# Pull Request #1537: HMS-9952: Add SendTemplateUpdateEvents job

**Author**: @swadeley
**Created**: June 11, 2026 at 02:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/HMS-9952`

## Description

## Summary

[HMS-9952 add job to send update events for all templates with added advisories](https://redhat.atlassian.net/browse/HMS-9952])

## Testing steps

with backend running

snapshot a repo or two
make a template with the repos you create

In terminal 1:
```
~/content-sources-backend$ make kafka-topic-consume KAFKA_TOPIC=platform.content-sources.template
info:Using DATABASE_* defaults
  % Total    % Received % Xferd  Average Speed  Time    Time    Time   Current
                                 Dload  Upload  Total   Spent   Left   Speed
100  16101 100  16101   0      0  45918      0                              0
CONTENT_DATABASE_USER=content CONTENT_DATABASE_PASSWORD=content CONTENT_DATABASE_DATABASE_NAME=content CONTENT_DATABASE_PORT=5433 KAFKA_CONFIG_DIR=/home/<user>/content-sources-backend/compose_files/kafka/config KAFKA_DATA_DIR=/home/<user>/content-sources-backend/compose_files/kafka/data ZOOKEEPER_CLIENT_PORT=2181 KAFKA_TOPICS=platform.content-sources.introspect  podman-compose --project-name=cs -f /home/<user>/content-sources-backend/deployments/docker-compose.yml exec kafka \
  /opt/kafka/bin/kafka-console-consumer.sh \
  --property print.key=true --property print.partition=true --property print.headers=true \
  --topic platform.content-sources.template \
  --group content-sources \
  --bootstrap-server localhost:9092
Option --property is deprecated and will be removed in a future version. Use --formatter-property instead.
```
In terminal 2

```
~/content-sources-backend$ go run cmd/jobs/main.go update-template-content
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[3.243ms] [rows:1] SELECT "uuid","org_id" FROM "templates" WHERE deleted_at IS NULL AND "templates"."deleted_at" IS NULL ORDER BY uuid severity=info
7:46PM INF Found templates to process severity=info template_count=1
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[0.873ms] [rows:1] SELECT * FROM "tasks" WHERE "tasks"."id" = '020f187f-7b93-4bd1-9c6a-bb3080436ce7' severity=info
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[1.176ms] [rows:2] SELECT * FROM "repository_configurations" WHERE "repository_configurations"."uuid" IN ('8a90d784-aa2b-45c9-a241-4076ef3ae57b','d071cb05-901c-4529-b126-6d466faec790') AND "repository_configurations"."deleted_at" IS NULL severity=info
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[1.914ms] [rows:2] SELECT * FROM "snapshots" WHERE "snapshots"."uuid" IN ('3ae5f98a-ead6-49b8-9f17-338001d7cdb9','e03479d0-fbbe-4898-868a-b4b7117ba6b7') AND "snapshots"."deleted_at" IS NULL severity=info
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[2.480ms] [rows:2] SELECT * FROM "templates_repository_configurations" WHERE "templates_repository_configurations"."template_uuid" = '3defcbc9-1e6a-4579-aded-21511a86469a' AND "templates_repository_configurations"."deleted_at" IS NULL severity=info
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[4.335ms] [rows:1] SELECT * FROM "templates" WHERE (uuid = '3defcbc9-1e6a-4579-aded-21511a86469a' AND org_id = '9999') AND "templates"."deleted_at" IS NULL ORDER BY "templates"."uuid" LIMIT 1 severity=info
7:46PM WRN Creating correlation ID for pulp request 316f68ae-c30a-4db1-840e-2e31fe65cec1 severity=warn
7:46PM INF /home/<user>/content-sources-backend/pkg/db/db_logger.go:173
[0.705ms] [rows:2] SELECT * FROM "repository_configurations" WHERE uuid IN ('8a90d784-aa2b-45c9-a241-4076ef3ae57b','d071cb05-901c-4529-b126-6d466faec790')  AND "repository_configurations"."deleted_at" IS NULL severity=info
7:46PM INF Sent template-updated event org_id=9999 processed=1 severity=info template_uuid=3defcbc9-1e6a-4579-aded-21511a86469a total=1
7:46PM INF Finished sending template-updated events events_sent=1 severity=info total_templates=1
```

In terminal 1:
```

Partition:0	content-type:application/cloudevents+json	null	{"specversion":"1.0","id":"271f798f-b65b-4a7c-838e-cb58131a429a","source":"urn:redhat:source:console:app:repositories","type":"com.redhat.console.repositories.template-updated","subject":"urn:redhat:subject:console:rhel:template-updated","datacontenttype":"application/json","time":"2026-06-16T18:46:43.355319998Z","data":[{"uuid":"3defcbc9-1e6a-4579-aded-21511a86469a","name":"test-template","org_id":"9999","description":"created via curl","arch":"x86_64","version":"9","date":"0001-01-01T00:00:00Z","repository_uuids":["8a90d784-aa2b-45c9-a241-4076ef3ae57b","d071cb05-901c-4529-b126-6d466faec790"],"rhsm_environment_id":"3defcbc91e6a4579aded21511a86469a"}],"redhatorgid":"9999"}



```

---

## Discussion

### Comment by @xbhouse on June 11, 2026 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-9952

### Comment by @swadeley on June 12, 2026 at 01:34 PM UTC

> Bryttanie's PR (https://github.com/content-services/content-sources-backend/pull/1536/changes) adds calculating advisory changes. I wonder if this would just need to enqueue update-template-content jobs to trigger events for updates. What do you think?

Hi, yes. 
I could refactor now to add InternalOnlyListSnapshotErrataIDs or rebase to use it after it is merged.

EDIT: I will pull in InternalOnlyListSnapshotErrataIDs now rather than wait for that PR to be merged. 



### Comment by @rverdile on June 12, 2026 at 02:03 PM UTC

@swadeley I'm a little confused about what the ticket is asking for. Is this understanding correct? 

This job meant to run before the changes in Bryttanie's PR start sending template update events, meaning it populates the initial errata for the template, and then Bryttanie's PR starts calculating the diff.

### Comment by @swadeley on June 15, 2026 at 01:58 PM UTC

> @swadeley I'm a little confused about what the ticket is asking for. Is this understanding correct?
> 
> This job meant to run before the changes in Bryttanie's PR start sending template update events, meaning it populates the initial errata for the template, and then Bryttanie's PR starts calculating the diff.

Hi @rverdile , yes, Justin confirmed, my job is a standalone backfill job that will be run when Patch is ready to accept and make use of the updates Bryttanie's PR code is sending.

### Comment by @xbhouse on June 15, 2026 at 03:15 PM UTC

> @swadeley I'm a little confused about what the ticket is asking for. Is this understanding correct?
> 
> This job meant to run before the changes in Bryttanie's PR start sending template update events, meaning it populates the initial errata for the template, and then Bryttanie's PR starts calculating the diff.

sorry for the confusion @rverdile! as @swadeley said, Justin confirmed this job is meant to be run after my PR is merged and after patch can store these relationships, so the ticket to run this job should be done last. i've updated the dependencies for each ticket so the order is clearer

### Comment by @rverdile on June 15, 2026 at 04:53 PM UTC

Thanks guys, I'm still a little confused where this sits in context with everything else. What is the reason for the backfill here? Are the updates added in @xbhouse's PR already being sent before this job runs? I noticed that PR has a feature flag - does that stay off until this job runs?

### Comment by @swadeley on June 15, 2026 at 05:50 PM UTC

> Thanks guys, I'm still a little confused where this sits in context with everything else. What is the reason for the backfill here? Are the updates added in @xbhouse's PR already being sent before this job runs? I noticed that PR has a feature flag - does that stay off until this job runs?

Hi @rverdile , code in [PR# 1536](https://github.com/content-services/content-sources-backend/pull/1536/changes) can be merged and send updates and but nothing will be done with them until feature is turned on. 

Feature will not be turned on until:
1.  Patch is updated to accept and hold the data coming from content-sources
2.  My backfill job is run



### Comment by @xbhouse on June 15, 2026 at 05:55 PM UTC

> Thanks guys, I'm still a little confused where this sits in context with everything else. What is the reason for the backfill here? Are the updates added in @xbhouse's PR already being sent before this job runs? I noticed that PR has a feature flag - does that stay off until this job runs?

the reason for this job is to let patch know about the advisories for all existing templates. the diff updates from my PR will already be sending whenever a template is updated, but that wouldn't include the full set.

regarding the flag: patch should just drop message fields if it doesn't understand them, Michael suggested adding a feature flag though just in case. so we'd turn the flag on first, then run this job.


EDIT: with the recent change in approach, this job should just enqueue update-template-content tasks for all existing templates so patch gets an update event for each one. no advisories are sent from content and we don't need a feature flag since the message structure doesn't change

### Comment by @rverdile on June 17, 2026 at 05:21 PM UTC

looks good to me!

---

## Reviews

### Review by @rverdile - Commented on June 11, 2026 at 03:58 PM UTC

Bryttanie's PR (https://github.com/content-services/content-sources-backend/pull/1536/changes) adds calculating advisory changes. I wonder if this would just need to enqueue update-template-content jobs to trigger events for updates. What do you think?

### Review by @xbhouse - Commented on June 15, 2026 at 03:58 PM UTC

### Review by @swadeley - Commented on June 15, 2026 at 05:40 PM UTC

### Review by @swadeley - Commented on June 15, 2026 at 06:51 PM UTC

### Review by @swadeley - Commented on June 15, 2026 at 06:56 PM UTC

### Review by @xbhouse - Commented on June 16, 2026 at 03:57 PM UTC

### Review by @rverdile - Commented on June 16, 2026 at 05:29 PM UTC

### Review by @swadeley - Commented on June 16, 2026 at 05:37 PM UTC

### Review by @xbhouse - Commented on June 16, 2026 at 06:00 PM UTC

### Review by @xbhouse - Commented on June 16, 2026 at 06:29 PM UTC

### Review by @swadeley - Commented on June 16, 2026 at 06:54 PM UTC

### Review by @swadeley - Commented on June 16, 2026 at 07:03 PM UTC

### Review by @rverdile - Commented on June 16, 2026 at 07:46 PM UTC

### Review by @xbhouse - Commented on June 17, 2026 at 12:04 AM UTC

### Review by @swadeley - Commented on June 17, 2026 at 11:10 AM UTC

### Review by @swadeley - Commented on June 17, 2026 at 11:12 AM UTC

### Review by @xbhouse - Approved on June 17, 2026 at 01:09 PM UTC

ack from me :) not sure if @rverdile wants to take another look

### Review by @xbhouse - Commented on June 17, 2026 at 01:22 PM UTC

### Review by @swadeley - Commented on June 18, 2026 at 08:33 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1537*
