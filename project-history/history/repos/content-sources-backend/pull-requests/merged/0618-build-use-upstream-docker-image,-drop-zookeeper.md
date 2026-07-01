---
type: pull_request
number: 618
title: "Build: Use upstream docker image, drop zookeeper"
state: merged
author: jlsherrill
created: 2024-03-28T17:19:17Z
updated: 2024-04-22T15:14:22Z
closed: 2024-04-22T15:14:19Z
merged: 2024-04-22T15:14:19Z
base_branch: main
head_branch: custom_kafka
labels: []
url: https://github.com/content-services/content-sources-backend/pull/618
---

# Pull Request #618: Build: Use upstream docker image, drop zookeeper

**Author**: @jlsherrill
**Created**: March 28, 2024 at 05:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `custom_kafka`

## Description

## Summary

This moves us to the official kafka container, as well as drops zookeeper.  Zookeeper is no longer needed, and in fact will removed from the kafka project entirely at 4.0 

## Testing steps

```
make compose-clean compose-up
make run
```

in another terminal
```
KAFKA_TOPICS=platform.content-sources.template make kafka-topic-consume
```

and then create a content template:

```
####
POST http://localhost:8000/api/content-sources/v1.0/templates/
x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==
Content-Type: application/json

{
 "name":"FOO",
  "repository_uuids": []
}
```

You should see the message make it to the kafka consume command:

```
Partition:0	content-type:application/cloudevents+json	null	{"specversion":"1.0","id":"ddcf8416-c153-4386-b1c8-9bf338eca641","source":"urn:redhat:source:console:app:repositories","type":"com.redhat.console.repositories.template-created","subject":"urn:redhat:subject:console:rhel:template-created","datacontenttype":"application/json","time":"2024-04-18T11:33:51.38720455Z","data":[{"uuid":"7bdbf4d9-04de-4405-ba4b-1d6505c61477","name":"FOO","org_id":"9","description":"","arch":"","version":"","date":"0001-01-01T00:00:00Z","repository_uuids":[]}],"redhatorgid":"9"}

```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 01, 2024 at 01:24 PM UTC

if we do want to merge this, i'd want to build an automated way to rebuild the container  'make kafka' or something

### Comment by @xbhouse on April 18, 2024 at 06:50 PM UTC

looks good on my end!

```
Partition:0	content-type:application/cloudevents+json	null	{"specversion":"1.0","id":"f6ec129c-256b-4ab5-b1df-a7d3dda9fe24","source":"urn:redhat:source:console:app:repositories","type":"com.redhat.console.repositories.template-created","subject":"urn:redhat:subject:console:rhel:template-created","datacontenttype":"application/json","time":"2024-04-18T18:47:42.587779Z","data":[{"uuid":"d0044364-8544-49cf-89bf-7a16c741d257","name":"test","org_id":"17684632","description":"description","arch":"aarch64","version":"8","date":"0001-01-01T00:00:00Z","repository_uuids":["34c37f8c-2d95-47bb-8c8e-e95f3ba14f3e"]}],"redhatorgid":"17684632"}
```

### Comment by @xbhouse on April 18, 2024 at 06:54 PM UTC

i did see this warning when creating the topic, but platform.content-sources.introspect doesn't have an underscore so not sure why it came up: 

```
WARNING: Due to limitations in metric names, topics with a period ('.') or underscore ('_') could collide. To avoid issues it is best to use either, but not both.
```

### Comment by @xbhouse on April 18, 2024 at 07:04 PM UTC

we should probably update the README too since we wouldn't need to build the container anymore

### Comment by @jlsherrill on April 22, 2024 at 11:52 AM UTC

fixed the readme.  That warning is fine, we don't necessarily control the topic name format, and just went with the console dot standard.

### Comment by @jlsherrill on April 22, 2024 at 03:14 PM UTC

@Andrewgdewar says he is fine with @xbhouse's testing and to go ahead and merge

---

## Reviews

### Review by @Andrewgdewar - Dismissed on April 15, 2024 at 04:30 PM UTC

### Review by @xbhouse - Approved on April 22, 2024 at 01:17 PM UTC

ack from me! @Andrewgdewar when you get a chance can you check if the kafka changes work for you?

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/618*
