---
type: issue
number: 36
title: "RFE: Option to turn on debug"
state: closed
author: catastrophe-brandon
created: 2025-10-24T17:15:46Z
updated: 2025-11-05T17:18:14Z
closed: 2025-11-05T17:18:14Z
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/issues/36
---

# Issue #36: RFE: Option to turn on debug

**Author**: @catastrophe-brandon
**Created**: October 24, 2025 at 05:15 PM UTC
**Status**: Closed
**Labels**: None

## Description

I'm trying to consume this container in a Tekton pipeline. In order to turn on debugging, I need to copy the Caddyfile from this project, drop it into a configmap in my project, add the debug flag, then set up my tekton pipeline to consume the configmap. This feels pretty heavy just to turn on debug. Would it be possible to add an env var to turn on debugging output?

`DEBUG=1` maybe?

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/issues/36*
