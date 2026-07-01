---
type: pull_request
number: 220
title: "Refs 1203: add additional metrics"
state: merged
author: jlsherrill
created: 2023-02-23T14:52:05Z
updated: 2023-03-02T21:37:41Z
closed: 2023-03-02T21:37:37Z
merged: 2023-03-02T21:37:37Z
base_branch: main
head_branch: 1203
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/220
---

# Pull Request #220: Refs 1203: add additional metrics

**Author**: @jlsherrill
**Created**: February 23, 2023 at 02:52 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1203`

## Description

Adds the following metrics:

* kafka message latency
* kafka message success (with error or without error)
* 36 hour intropsection status (intropsected or not)
  * this reworks a lot of the existing 24 hours metric gathering


---

## Discussion

### Comment by @jlsherrill on February 23, 2023 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-1203

### Comment by @jlsherrill on February 23, 2023 at 07:00 PM UTC

To test:

1) 
Send some badly formatted kafka messages:
```
$ podman run   -i --rm   --net=host   docker.io/edenhill/kcat:1.7.1   -k "c67cd587-3741-493d-9302-f655fcd3bd68"   -H X-Rh-Identity="eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIxMTExMTEiLCJpbnRlcm5hbCI6eyJvcmdfaWQiOiIyMjIyMjIifX19Cg=="   -H X-Rh-Insight-Request-Id="demo"   -H Type="Introspect"   -b localhost:9092   -t platform.content-sources.introspect   -P <<< "bad-json"
```
send some good ones:
```
 make kafka-produce-msg 
```

then check the metrics:
```
curl localhost:9000/metrics | grep kafka_message_status
```

2) After doing that, check the latency:

```
curl localhost:9000/metrics | grep kafka_message_latency
```

3) insert a couple of repositories:

```
make db-cli-connect
```

NOTE: You may need to change 2023-02-23 to the current day/previous day for it to be less than 36 hours ago.

```
insert into repositories (uuid, url, last_introspection_time) values ('905f7459-1cd2-46bd-ba9c-e54ca93ad6a2', 'introspected/', '2023-02-23');
insert into repositories (uuid, url, last_introspection_time) values ('905f7459-1cd2-46bd-ba9c-e54ca93ad6a1', 'missed/', '2023-02-10');
insert into repositories (uuid, url, last_introspection_time, public) values ('905f7459-1cd2-46bd-ba9c-e54ca93ad6a3', 'public_introspected/', '2023-02-23', true);
INSERT 0 1
insert into repositories (uuid, url, last_introspection_time, public) values ('905f7459-1cd2-46bd-ba9c-e54ca93ad6a4', 'public_missed/', '2023-02-10', true);
```

then after a second or so, check the metrics:
```
curl localhost:9000/metrics | grep introsp
```

You may also want to introspect real repos via the ui creation and examine the output

### Comment by @jlsherrill on March 02, 2023 at 07:54 PM UTC

updated!

---

## Reviews

### Review by @rverdile - Commented on February 28, 2023 at 04:49 PM UTC

a couple of help texts need updating, but otherwise I think the code looks good. Still testing.

### Review by @rverdile - Commented on February 28, 2023 at 07:09 PM UTC

~~The `content_sources_custom_repositories_36_hour_introspection_total` doesn't seem to be working correctly for me. Seems like it's not decrementing when it should. Trying to figure out why. I haven't looked as closely at public_repositories, but it may also be an issue there.~~

Nevermind, just confused myself.

### Review by @rverdile - Commented on February 28, 2023 at 08:08 PM UTC

### Review by @rverdile - Commented on February 28, 2023 at 08:13 PM UTC

### Review by @rverdile - Commented on February 28, 2023 at 08:14 PM UTC

Looks like everything is working as intended! I just had some questions mostly around the clarity of the names of the metrics

### Review by @jlsherrill - Commented on March 02, 2023 at 01:40 PM UTC

### Review by @rverdile - Commented on March 02, 2023 at 02:57 PM UTC

### Review by @jlsherrill - Commented on March 02, 2023 at 02:58 PM UTC

### Review by @rverdile - Commented on March 02, 2023 at 03:32 PM UTC

### Review by @rverdile - Approved on March 02, 2023 at 08:04 PM UTC

looks good!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/220*
