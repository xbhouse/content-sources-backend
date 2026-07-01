---
type: pull_request
number: 1667
title: "chore(deps): update node.js to v24.17.0"
state: merged
author: red-hat-konflux
created: 2026-06-22T01:12:24Z
updated: 2026-06-22T09:34:21Z
closed: 2026-06-22T07:23:46Z
merged: 2026-06-22T07:23:45Z
base_branch: master
head_branch: konflux/mintmaker/master/node-24.x
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1667
---

# Pull Request #1667: chore(deps): update node.js to v24.17.0

**Author**: @red-hat-konflux
**Created**: June 22, 2026 at 01:12 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/node-24.x`

## Description

This PR contains the following updates:

| Package | Update | Change |
|---|---|---|
| [node](https://nodejs.org) ([source](https://redirect.github.com/nodejs/node)) | minor | `24.16.0` → `24.17.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>nodejs/node (node)</summary>

### [`v24.17.0`](https://redirect.github.com/nodejs/node/releases/tag/v24.17.0): 2026-06-18, Version 24.17.0 &#x27;Krypton&#x27; (LTS), @&#8203;aduh95

[Compare Source](https://redirect.github.com/nodejs/node/compare/v24.16.0...v24.17.0)

This is a security release.

##### Notable Changes

- (CVE-2026-48618) tls: normalize hostname for server identity checks (Matteo Collina) – High
- (CVE-2026-48933) crypto: guard WebCrypto cipher output length (Filip Skokan) – High
- (CVE-2026-48615) lib,test: redact proxy credentials in tunnel errors (Matteo Collina) – Medium
- (CVE-2026-48619) http2: cap originSet size to prevent unbounded memory growth (Matteo Collina) – Medium
- (CVE-2026-48928) tls: fix case-sensitive SNI context matching (Matteo Collina) – Medium
- (CVE-2026-48930) dns,net: reject hostnames with embedded NUL bytes (Matteo Collina) – Medium
- (CVE-2026-48934) tls: bind reusable sessions to authenticated host (Matteo Collina) – Medium
- (CVE-2026-48937) deps: fix integration issues with the latest nghttp2 – Medium
- (CVE-2026-48617) permission: handle process.chdir on writereport (RafaelGSS) – Low
- (CVE-2026-48931) http: fix response queue poisoning in http.Agent (Matteo Collina) – Low
- (CVE-2026-48935) permission: disable FileHandle utimes with permission model (RafaelGSS) – Low

##### Commits

- \[[`9e4dfc7bba`](https://redirect.github.com/nodejs/node/commit/9e4dfc7bba)] - **(CVE-2026-48933)** **crypto**: guard WebCrypto cipher output length (Filip Skokan) [nodejs-private/node-private#878](https://redirect.github.com/nodejs-private/node-private/pull/878)
- \[[`cb2aed980c`](https://redirect.github.com/nodejs/node/commit/cb2aed980c)] - **deps**: update llhttp to 9.4.2 (Antoine du Hamel) [nodejs-private/node-private#890](https://redirect.github.com/nodejs-private/node-private/pull/890)
- \[[`a8a0d12875`](https://redirect.github.com/nodejs/node/commit/a8a0d12875)] - **(CVE-2026-48937)** **deps**: fix integration issues with the latest nghttp2 (Tim Perry) [#&#8203;62891](https://redirect.github.com/nodejs/node/pull/62891)
- \[[`66e6203c1c`](https://redirect.github.com/nodejs/node/commit/66e6203c1c)] - **(SEMVER-MAJOR)** **deps**: update nghttp2 to 1.69.0 (Node.js GitHub Bot) [#&#8203;62891](https://redirect.github.com/nodejs/node/pull/62891)
- \[[`dd627ced27`](https://redirect.github.com/nodejs/node/commit/dd627ced27)] - **deps**: update archs files for openssl-3.5.7 (Node.js GitHub Bot) [#&#8203;63820](https://redirect.github.com/nodejs/node/pull/63820)
- \[[`684bae568f`](https://redirect.github.com/nodejs/node/commit/684bae568f)] - **deps**: upgrade openssl sources to openssl-3.5.7 (Node.js GitHub Bot) [#&#8203;63820](https://redirect.github.com/nodejs/node/pull/63820)
- \[[`3a631e7f83`](https://redirect.github.com/nodejs/node/commit/3a631e7f83)] - **deps**: fix aix implicit declaration in OpenSSL (Abdirahim Musse) [#&#8203;62656](https://redirect.github.com/nodejs/node/pull/62656)
- \[[`cf44df3996`](https://redirect.github.com/nodejs/node/commit/cf44df3996)] - **deps**: update undici to 7.28.0 (Node.js GitHub Bot) [#&#8203;63703](https://redirect.github.com/nodejs/node/pull/63703)
- \[[`138c70294b`](https://redirect.github.com/nodejs/node/commit/138c70294b)] - **(CVE-2026-48930)** **dns,net**: reject hostnames with embedded NUL bytes (Matteo Collina) [nodejs-private/node-private#868](https://redirect.github.com/nodejs-private/node-private/pull/868)
- \[[`be7e719c3f`](https://redirect.github.com/nodejs/node/commit/be7e719c3f)] - **(CVE-2026-48931)** **http**: fix response queue poisoning in http.Agent (Matteo Collina) [nodejs-private/node-private#846](https://redirect.github.com/nodejs-private/node-private/pull/846)
- \[[`cc7c11b4d1`](https://redirect.github.com/nodejs/node/commit/cc7c11b4d1)] - **(CVE-2026-48619)** **http2**: cap originSet size to prevent unbounded memory growth (Matteo Collina) [nodejs-private/node-private#855](https://redirect.github.com/nodejs-private/node-private/pull/855)
- \[[`9224427b92`](https://redirect.github.com/nodejs/node/commit/9224427b92)] - **(CVE-2026-48615)** **lib,test**: redact proxy credentials in tunnel errors (Matteo Collina) [nodejs-private/node-private#867](https://redirect.github.com/nodejs-private/node-private/pull/867)
- \[[`cf85d54839`](https://redirect.github.com/nodejs/node/commit/cf85d54839)] - **(CVE-2026-48935)** **permission**: disable FileHandle utimes with permission model (RafaelGSS) [nodejs-private/node-private#873](https://redirect.github.com/nodejs-private/node-private/pull/873)
- \[[`a1bbc24f96`](https://redirect.github.com/nodejs/node/commit/a1bbc24f96)] - **(CVE-2026-48617)** **permission**: handle process.chdir on writereport (RafaelGSS) [nodejs-private/node-private#870](https://redirect.github.com/nodejs-private/node-private/pull/870)
- \[[`e3723ff2d6`](https://redirect.github.com/nodejs/node/commit/e3723ff2d6)] - **test**: add session reuse host verification regressions (Matteo Collina) [nodejs-private/node-private#854](https://redirect.github.com/nodejs-private/node-private/pull/854)
- \[[`a77af4867b`](https://redirect.github.com/nodejs/node/commit/a77af4867b)] - **(CVE-2026-48934)** **tls**: bind reusable sessions to authenticated host (Matteo Collina) [nodejs-private/node-private#854](https://redirect.github.com/nodejs-private/node-private/pull/854)
- \[[`31beb4f707`](https://redirect.github.com/nodejs/node/commit/31beb4f707)] - **(CVE-2026-48928)** **tls**: fix case-sensitive SNI context matching (Matteo Collina) [nodejs-private/node-private#857](https://redirect.github.com/nodejs-private/node-private/pull/857)
- \[[`8e75c73f91`](https://redirect.github.com/nodejs/node/commit/8e75c73f91)] - **(CVE-2026-48618)** **tls**: normalize hostname for server identity checks (Matteo Collina) [nodejs-private/node-private#869](https://redirect.github.com/nodejs-private/node-private/pull/869)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on June 22, 2026 at 01:15 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1667?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`49d0fb1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/49d0fb1e3781c62e9029b9593a6cc546021cf53b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b8cb399`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b8cb399db60b32dcad9f090d41349a45b348a097?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 9 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1667   +/-   ##
=======================================
  Coverage   77.27%   77.27%           
=======================================
  Files         103      103           
  Lines        3287     3287           
  Branches      740      740           
=======================================
  Hits         2540     2540           
  Misses        668      668           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1667?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @github-actions - Approved on June 22, 2026 at 01:12 AM UTC

This PR from Konflux has been automatically approved.

### Review by @swadeley - Approved on June 22, 2026 at 07:23 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1667*
