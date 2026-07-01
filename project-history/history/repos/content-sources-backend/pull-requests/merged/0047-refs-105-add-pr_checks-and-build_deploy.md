---
type: pull_request
number: 47
title: "Refs 105: Add pr_checks and build_deploy"
state: merged
author: mshriver
created: 2022-06-29T23:36:48Z
updated: 2022-07-25T23:07:41Z
closed: 2022-07-25T23:07:41Z
merged: 2022-07-25T23:07:41Z
base_branch: main
head_branch: add-appsre-scripts
labels: ["enhancement", "no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/47
---

# Pull Request #47: Refs 105: Add pr_checks and build_deploy

**Author**: @mshriver
**Created**: June 29, 2022 at 11:36 PM UTC
**Status**: Merged
**Labels**: `enhancement`, `no-qe-needed`
**Base**: `main` ← **Head**: `add-appsre-scripts`

## Description

Add a docker-login make target, for use in build_deploy.sh

build_deploy.sh will be used by appsre pipelines to push commits to
quay/redhat.io. The GH action building to quay can be removed (once
appinterface updates are accepted and repo created).

pr_checks leverages bonfire for PR image building, deployment, and test execution. appinterface pipeline definitions will execute this script against PRs.

---

## Discussion

### Comment by @jlsherrill on July 05, 2022 at 03:23 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-105

### Comment by @jlsherrill on July 05, 2022 at 03:26 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on July 18, 2022 at 06:42 PM UTC

Built and pushed image for bb691bd53008b196649a3e6a70112894a2124d67

### Comment by @jlsherrill on July 22, 2022 at 08:22 PM UTC

Built and pushed image for 20cf2ba085260631cac10e29d5296a82c1c50d7f

---

## Reviews

### Review by @mshriver - Commented on June 29, 2022 at 11:37 PM UTC

### Review by @mshriver - Commented on June 29, 2022 at 11:38 PM UTC

### Review by @mshriver - Commented on June 29, 2022 at 11:39 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 03:26 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 03:27 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 03:28 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 03:29 PM UTC

### Review by @avisiedo - Commented on July 01, 2022 at 03:29 PM UTC

### Review by @avisiedo - Commented on July 13, 2022 at 10:07 AM UTC

### Review by @avisiedo - Commented on July 13, 2022 at 12:57 PM UTC

### Review by @avisiedo - Commented on July 13, 2022 at 12:58 PM UTC

### Review by @avisiedo - Commented on July 13, 2022 at 01:10 PM UTC

### Review by @avisiedo - Commented on July 13, 2022 at 01:19 PM UTC

A couple of comments! and there are parts that I would need to learn more honestly!

### Review by @mshriver - Commented on July 13, 2022 at 01:59 PM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:08 AM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:09 AM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:11 AM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:15 AM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:18 AM UTC

### Review by @avisiedo - Commented on July 18, 2022 at 12:23 AM UTC

### Review by @mshriver - Commented on July 22, 2022 at 10:51 AM UTC

### Review by @avisiedo - Commented on July 22, 2022 at 11:07 AM UTC

### Review by @avisiedo - Commented on July 22, 2022 at 11:07 AM UTC

### Review by @mshriver - Commented on July 22, 2022 at 11:38 AM UTC

### Review by @avisiedo - Approved on July 22, 2022 at 11:50 AM UTC

LGTM

### Review by @jlsherrill - Approved on July 25, 2022 at 11:51 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/47*
