---
type: pull_request
number: 97
title: "Refs 151: Add makefile rules for kafka"
state: merged
author: avisiedo
created: 2022-09-06T12:10:14Z
updated: 2022-10-31T18:00:29Z
closed: 2022-09-20T06:50:17Z
merged: 2022-09-20T06:50:17Z
base_branch: main
head_branch: hmscontent-151-local-kafka
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/97
---

# Pull Request #97: Refs 151: Add makefile rules for kafka

**Author**: @avisiedo
**Created**: September 06, 2022 at 12:10 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-151-local-kafka`

## Description

- Add kafka-up to start local container to leverage kafka.
- Add kafka-down to stop local container.
- Add kafka-build to build the container.
- Add Dockerfile for the kafka container.
- Configurations are stored at kafka/config
- Some scripts are stored at kafka/scripts

> The configuration passed through environment variables would need some clean-up, but some could be impacting as I got error starting the containers.

> It would need refinement, but this commit start and stop the containers;another commits that I had broke this

---

## Discussion

### Comment by @jlsherrill on September 06, 2022 at 12:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-151

### Comment by @jlsherrill on September 06, 2022 at 12:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @avisiedo on September 14, 2022 at 12:14 PM UTC

I fixed one more stuff related with the directory owner; For any reason the directory that is bound into the container has the mapped user (185 (container) -> 100184 (host)) (I checked the container (podman image inspect ... ) and I didn't see a `USER 185` directive); Not sure where this is comming from and this was evoking situations when starting the container. I have fixed this in the kafka Dockerfile by using `USER 0`, so the root user inside the container (0 (container) > $UID (host)) is mapped to my user account out of the container, allowing to write and create directories into the bound data directory inside the container; not sure how it was working previously in the workstation; probably some residual states. now it is fixed.

I have added `kafka-clean` to stop and clean the kafka container (remove the data directory).
I have added `kafka-topics-create` so the command is launched inside the container if necessary (it uses KAFKA_TOPICS variable, we can override at variables.mk or directly from the terminal).
I have added `kafka-topics-list` so we get the list of existing topics.

### Comment by @jlsherrill on September 15, 2022 at 05:40 PM UTC

/retest
[test]

### Comment by @jlsherrill on September 15, 2022 at 05:41 PM UTC

while the main issue could be tested, this particular Pr should not require any QE

---

## Reviews

### Review by @jlsherrill - Commented on September 12, 2022 at 05:58 PM UTC

### Review by @avisiedo - Commented on September 14, 2022 at 09:55 AM UTC

### Review by @jlsherrill - Approved on September 15, 2022 at 05:40 PM UTC

ACK from me.  I haven't actually tested the running server, but the containers start up for me with no issue

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/97*
