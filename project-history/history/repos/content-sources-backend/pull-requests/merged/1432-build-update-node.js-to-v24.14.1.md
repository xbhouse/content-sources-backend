---
type: pull_request
number: 1432
title: "Build: Update Node.js to v24.14.1"
state: merged
author: red-hat-konflux
created: 2026-03-30T05:13:15Z
updated: 2026-03-30T07:17:44Z
closed: 2026-03-30T07:17:42Z
merged: 2026-03-30T07:17:42Z
base_branch: main
head_branch: konflux/mintmaker/main/node-24.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1432
---

# Pull Request #1432: Build: Update Node.js to v24.14.1

**Author**: @red-hat-konflux
**Created**: March 30, 2026 at 05:13 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | patch | `24.14.0` -> `24.14.1` |

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.14.1`](https://redirect.github.com/nodejs/node/releases/tag/v24.14.1): 2026-03-24, Version 24.14.1 &#x27;Krypton&#x27; (LTS), @&#8203;RafaelGSS prepared by @&#8203;juanarbol

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.14.0...v24.14.1)

This is a security release.

##### Notable Changes

- (CVE-2026-21710) use null prototype for headersDistinct/trailersDistinct (Matteo Collina) - High
- (CVE-2026-21637) wrap SNICallback invocation in try/catch (Matteo Collina) - High
- (CVE-2026-21717) test array index hash collision (Joyee Cheung) - Medium
- (CVE-2026-21713) use timing-safe comparison in Web Cryptography HMAC and KMAC (Filip Skokan) - Medium
- (CVE-2026-21714) handle NGHTTP2\_ERR\_FLOW\_CONTROL error code (RafaelGSS) - Medium
- (CVE-2026-21712) handle url crash on different url formats (RafaelGSS) - Medium
- (CVE-2026-21716) include permission check on lib/fs/promises (RafaelGSS) - Low
- (CVE-2026-21715) add permission check to realpath.native (RafaelGSS) - Low

##### Commits

- \[[`6fae244080`](https://redirect.github.com/nodejs/node/commit/6fae244080)] - **(CVE-2026-21717)** **build,test**: test array index hash collision (Joyee Cheung) [nodejs-private/node-private#828](https://redirect.github.com/nodejs-private/node-private/pull/828)
- \[[`cc0910c62e`](https://redirect.github.com/nodejs/node/commit/cc0910c62e)] - **(CVE-2026-21713)** **crypto**: use timing-safe comparison in Web Cryptography HMAC and KMAC (Filip Skokan) [nodejs-private/node-private#822](https://redirect.github.com/nodejs-private/node-private/pull/822)
- \[[`80cb042cf3`](https://redirect.github.com/nodejs/node/commit/80cb042cf3)] - **deps**: update undici to 7.24.4 (Node.js GitHub Bot) [#&#8203;62271](https://redirect.github.com/nodejs/node/pull/62271)
- \[[`f5b8667dc2`](https://redirect.github.com/nodejs/node/commit/f5b8667dc2)] - **deps**: update undici to 7.24.3 (Node.js GitHub Bot) [#&#8203;62233](https://redirect.github.com/nodejs/node/pull/62233)
- \[[`08852637d9`](https://redirect.github.com/nodejs/node/commit/08852637d9)] - **deps**: update undici to 7.22.0 (Node.js GitHub Bot) [#&#8203;62035](https://redirect.github.com/nodejs/node/pull/62035)
- \[[`61097db9fb`](https://redirect.github.com/nodejs/node/commit/61097db9fb)] - **deps**: upgrade npm to 11.11.0 (npm team) [#&#8203;61994](https://redirect.github.com/nodejs/node/pull/61994)
- \[[`9ac0f9f81e`](https://redirect.github.com/nodejs/node/commit/9ac0f9f81e)] - **deps**: upgrade npm to 11.10.1 (npm team) [#&#8203;61892](https://redirect.github.com/nodejs/node/pull/61892)
- \[[`3dab3c4698`](https://redirect.github.com/nodejs/node/commit/3dab3c4698)] - **deps**: V8: override `depot_tools` version (Richard Lau) [#&#8203;62344](https://redirect.github.com/nodejs/node/pull/62344)
- \[[`87521e99d1`](https://redirect.github.com/nodejs/node/commit/87521e99d1)] - **deps**: V8: backport [`1361b2a`](https://redirect.github.com/nodejs/node/commit/1361b2a49d02) (Joyee Cheung) [nodejs-private/node-private#828](https://redirect.github.com/nodejs-private/node-private/pull/828)
- \[[`045013366f`](https://redirect.github.com/nodejs/node/commit/045013366f)] - **deps**: V8: backport [`185f0fe`](https://redirect.github.com/nodejs/node/commit/185f0fe09b72) (Joyee Cheung) [nodejs-private/node-private#828](https://redirect.github.com/nodejs-private/node-private/pull/828)
- \[[`af22629ea8`](https://redirect.github.com/nodejs/node/commit/af22629ea8)] - **deps**: V8: backport [`0a8b1cd`](https://redirect.github.com/nodejs/node/commit/0a8b1cdcc8b2) (snek) [nodejs-private/node-private#828](https://redirect.github.com/nodejs-private/node-private/pull/828)
- \[[`380ea72eef`](https://redirect.github.com/nodejs/node/commit/380ea72eef)] - **(CVE-2026-21710)** **http**: use null prototype for headersDistinct/trailersDistinct (Matteo Collina) [nodejs-private/node-private#821](https://redirect.github.com/nodejs-private/node-private/pull/821)
- \[[`d6b6051e08`](https://redirect.github.com/nodejs/node/commit/d6b6051e08)] - **(CVE-2026-21716)** **permission**: include permission check on lib/fs/promises (RafaelGSS) [nodejs-private/node-private#795](https://redirect.github.com/nodejs-private/node-private/pull/795)
- \[[`bfdecef9da`](https://redirect.github.com/nodejs/node/commit/bfdecef9da)] - **(CVE-2026-21715)** **permission**: add permission check to realpath.native (RafaelGSS) [nodejs-private/node-private#794](https://redirect.github.com/nodejs-private/node-private/pull/794)
- \[[`c015edf313`](https://redirect.github.com/nodejs/node/commit/c015edf313)] - **(CVE-2026-21714)** **src**: handle NGHTTP2\_ERR\_FLOW\_CONTROL error code (RafaelGSS) [nodejs-private/node-private#832](https://redirect.github.com/nodejs-private/node-private/pull/832)
- \[[`cba66c48a5`](https://redirect.github.com/nodejs/node/commit/cba66c48a5)] - **(CVE-2026-21712)** **src**: handle url crash on different url formats (RafaelGSS) [nodejs-private/node-private#816](https://redirect.github.com/nodejs-private/node-private/pull/816)
- \[[`df8fbfb93d`](https://redirect.github.com/nodejs/node/commit/df8fbfb93d)] - **(CVE-2026-21637)** **tls**: wrap SNICallback invocation in try/catch (Matteo Collina) [nodejs-private/node-private#819](https://redirect.github.com/nodejs-private/node-private/pull/819)

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

### Review by @swadeley - Approved on March 30, 2026 at 07:17 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1432*
