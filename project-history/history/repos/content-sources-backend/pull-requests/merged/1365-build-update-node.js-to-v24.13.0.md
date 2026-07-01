---
type: pull_request
number: 1365
title: "Build: Update Node.js to v24.13.0"
state: merged
author: red-hat-konflux
created: 2026-01-19T07:46:53Z
updated: 2026-01-21T02:00:29Z
closed: 2026-01-21T01:59:19Z
merged: 2026-01-21T01:59:19Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1365
---

# Pull Request #1365: Build: Update Node.js to v24.13.0

**Author**: @red-hat-konflux
**Created**: January 19, 2026 at 07:46 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.12.0` -> `24.13.0` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.13.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.13.0): 2026-01-13, Version 24.13.0 &#x27;Krypton&#x27; (LTS), @&#8203;marco-ippolito

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.12.0...v24.13.0)

This is a security release.

##### Notable Changes

lib:

- (CVE-2025-59465) add TLSSocket default error handler (RafaelGSS) [nodejs-private/node-private#797](https://redirect.github.com/nodejs-private/node-private/pull/797)
- (CVE-2025-55132) disable futimes when permission model is enabled (RafaelGSS) [nodejs-private/node-private#748](https://redirect.github.com/nodejs-private/node-private/pull/748)
  lib,permission:
- (CVE-2025-55130) require full read and write to symlink APIs (RafaelGSS) [nodejs-private/node-private#760](https://redirect.github.com/nodejs-private/node-private/pull/760)
  src:
- (CVE-2025-59466) rethrow stack overflow exceptions in async\_hooks (Matteo Collina) [nodejs-private/node-private#773](https://redirect.github.com/nodejs-private/node-private/pull/773)
  src,lib:
- (CVE-2025-55131) refactor unsafe buffer creation to remove zero-fill toggle (Сковорода Никита Андреевич) [nodejs-private/node-private#759](https://redirect.github.com/nodejs-private/node-private/pull/759)
  tls:
- (CVE-2026-21637) route callback exceptions through error handlers (Matteo Collina) [nodejs-private/node-private#796](https://redirect.github.com/nodejs-private/node-private/pull/796)

##### Commits

- \[[`2092785d01`](https://redirect.github.com/nodejs/node/commit/2092785d01)] - **deps**: update c-ares to v1.34.6 (Node.js GitHub Bot) [#&#8203;60997](https://redirect.github.com/nodejs/node/pull/60997)
- \[[`3e58b7f2af`](https://redirect.github.com/nodejs/node/commit/3e58b7f2af)] - **deps**: update undici to 7.18.2 (Node.js GitHub Bot) [#&#8203;61283](https://redirect.github.com/nodejs/node/pull/61283)
- \[[`4ba536a5a6`](https://redirect.github.com/nodejs/node/commit/4ba536a5a6)] - **(CVE-2025-59465)** **lib**: add TLSSocket default error handler (RafaelGSS) [nodejs-private/node-private#797](https://redirect.github.com/nodejs-private/node-private/pull/797)
- \[[`89adaa21fd`](https://redirect.github.com/nodejs/node/commit/89adaa21fd)] - **(CVE-2025-55132)** **lib**: disable futimes when permission model is enabled (RafaelGSS) [nodejs-private/node-private#748](https://redirect.github.com/nodejs-private/node-private/pull/748)
- \[[`7302b4dae1`](https://redirect.github.com/nodejs/node/commit/7302b4dae1)] - **(CVE-2025-55130)** **lib,permission**: require full read and write to symlink APIs (RafaelGSS) [nodejs-private/node-private#760](https://redirect.github.com/nodejs-private/node-private/pull/760)
- \[[`ac030753c4`](https://redirect.github.com/nodejs/node/commit/ac030753c4)] - **(CVE-2025-59466)** **src**: rethrow stack overflow exceptions in async\_hooks (Matteo Collina) [nodejs-private/node-private#773](https://redirect.github.com/nodejs-private/node-private/pull/773)
- \[[`20075692fe`](https://redirect.github.com/nodejs/node/commit/20075692fe)] - **(CVE-2025-55131)** **src,lib**: refactor unsafe buffer creation to remove zero-fill toggle (Сковорода Никита Андреевич) [nodejs-private/node-private#759](https://redirect.github.com/nodejs-private/node-private/pull/759)
- \[[`20591b0618`](https://redirect.github.com/nodejs/node/commit/20591b0618)] - **(CVE-2026-21637)** **tls**: route callback exceptions through error handlers (Matteo Collina) [nodejs-private/node-private#796](https://redirect.github.com/nodejs-private/node-private/pull/796)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "on Monday after 3am and before 10am" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @swadeley - Approved on January 21, 2026 at 01:59 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1365*
