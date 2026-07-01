---
type: issue
number: 1501
title: "add patch simulation and impact preview"
state: closed
author: dmzoneill
created: 2026-02-12T21:14:40Z
updated: 2026-02-13T16:57:18Z
closed: 2026-02-13T16:57:18Z
labels: []
url: https://github.com/RedHatInsights/patchman-ui/issues/1501
---

# Issue #1501: add patch simulation and impact preview

**Author**: @dmzoneill
**Created**: February 12, 2026 at 09:14 PM UTC
**Status**: Closed
**Labels**: None

## Description

before applying patches would be helpful to see what would change - which services need restart, which packages have dependency chains, estimated downtime impact.

use case: ops teams want to understand blast radius before scheduling maintenance windows. right now you find out after applying patches if something critical needs a reboot or if a library update cascades to application services.

could show a diff-style preview of current vs proposed package versions, highlight services that would be affected (systemd units that depend on updated packages), and estimate whether reboot is required based on kernel/init updates.

backend probably needs to query package metadata and systemd dependencies from the system's package manager. might be too complex if this info isn't already available from insights data.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/issues/1501*
