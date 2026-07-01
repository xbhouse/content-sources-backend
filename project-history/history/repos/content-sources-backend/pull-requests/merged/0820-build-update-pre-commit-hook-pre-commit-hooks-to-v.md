---
type: pull_request
number: 820
title: "Build: Update pre-commit hook pre-commit-hooks to v4.6.0"
state: merged
author: red-hat-konflux
created: 2024-09-17T12:23:52Z
updated: 2024-09-18T13:19:00Z
closed: 2024-09-18T13:18:56Z
merged: 2024-09-18T13:18:56Z
base_branch: main
head_branch: konflux/mintmaker/main/pre-commit-pre-commit-hooks-4.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/820
---

# Pull Request #820: Build: Update pre-commit hook pre-commit-hooks to v4.6.0

**Author**: @red-hat-konflux
**Created**: September 17, 2024 at 12:23 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/pre-commit-pre-commit-hooks-4.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [pre-commit/pre-commit-hooks](https://togithub.com/pre-commit/pre-commit-hooks) | repository | minor | `v4.3.0` -> `v4.6.0` |

Note: The `pre-commit` manager in Renovate is not supported by the `pre-commit` maintainers or community. Please do not report any problems there, instead [create a Discussion in the Renovate repository](https://togithub.com/renovatebot/renovate/discussions/new) if you have any questions.

---

### Release Notes

<details>
<summary>pre-commit/pre-commit-hooks (pre-commit/pre-commit-hooks)</summary>

### [`v4.6.0`](https://togithub.com/pre-commit/pre-commit-hooks/releases/tag/v4.6.0): pre-commit-hooks v4.6.0

[Compare Source](https://togithub.com/pre-commit/pre-commit-hooks/compare/v4.5.0...v4.6.0)

##### Features

-   `requirements-txt-fixer`: remove duplicate packages.
    -   [#&#8203;1014](https://togithub.com/pre-commit/pre-commit-hooks/issues/1014) PR by [@&#8203;vhoulbreque-withings](https://togithub.com/vhoulbreque-withings).
    -   [#&#8203;960](https://togithub.com/pre-commit/pre-commit-hooks/issues/960) issue [@&#8203;csibe17](https://togithub.com/csibe17).

##### Migrating

-   `fix-encoding-pragma`: deprecated -- will be removed in 5.0.0.  use
    [pyupgrade](https://togithub.com/asottile/pyupgrade) or some other tool.
    -   [#&#8203;1033](https://togithub.com/pre-commit/pre-commit-hooks/issues/1033) PR by [@&#8203;mxr](https://togithub.com/mxr).
    -   [#&#8203;1032](https://togithub.com/pre-commit/pre-commit-hooks/issues/1032) issue by [@&#8203;mxr](https://togithub.com/mxr).

### [`v4.5.0`](https://togithub.com/pre-commit/pre-commit-hooks/releases/tag/v4.5.0): pre-commit-hooks v4.5.0

[Compare Source](https://togithub.com/pre-commit/pre-commit-hooks/compare/v4.4.0...v4.5.0)

##### Features

-   `requirements-txt-fixer`: also sort `constraints.txt` by default.
    -   [#&#8203;857](https://togithub.com/pre-commit/pre-commit-hooks/issues/857) PR by [@&#8203;lev-blit](https://togithub.com/lev-blit).
    -   [#&#8203;830](https://togithub.com/pre-commit/pre-commit-hooks/issues/830) issue by [@&#8203;PLPeeters](https://togithub.com/PLPeeters).
-   `debug-statements`: add `bpdb` debugger.
    -   [#&#8203;942](https://togithub.com/pre-commit/pre-commit-hooks/issues/942) PR by [@&#8203;mwip](https://togithub.com/mwip).
    -   [#&#8203;941](https://togithub.com/pre-commit/pre-commit-hooks/issues/941) issue by [@&#8203;mwip](https://togithub.com/mwip).

##### Fixes

-   `file-contents-sorter`: fix sorting an empty file.
    -   [#&#8203;944](https://togithub.com/pre-commit/pre-commit-hooks/issues/944) PR by [@&#8203;RoelAdriaans](https://togithub.com/RoelAdriaans).
    -   [#&#8203;935](https://togithub.com/pre-commit/pre-commit-hooks/issues/935) issue by [@&#8203;paduszyk](https://togithub.com/paduszyk).
-   `double-quote-string-fixer`: don't rewrite inside f-strings in 3.12+.
    -   [#&#8203;973](https://togithub.com/pre-commit/pre-commit-hooks/issues/973) PR by [@&#8203;asottile](https://togithub.com/asottile).
    -   [#&#8203;971](https://togithub.com/pre-commit/pre-commit-hooks/issues/971) issue by [@&#8203;XuehaiPan](https://togithub.com/XuehaiPan).

#### Migrating

-   now requires python >= 3.8.
    -   [#&#8203;926](https://togithub.com/pre-commit/pre-commit-hooks/issues/926) PR by [@&#8203;asottile](https://togithub.com/asottile).
    -   [#&#8203;927](https://togithub.com/pre-commit/pre-commit-hooks/issues/927) PR by [@&#8203;asottile](https://togithub.com/asottile).

### [`v4.4.0`](https://togithub.com/pre-commit/pre-commit-hooks/releases/tag/v4.4.0): pre-commit-hooks v4.4.0

[Compare Source](https://togithub.com/pre-commit/pre-commit-hooks/compare/v4.3.0...v4.4.0)

#### Features

-   forbid-submodules: new hook which outright bans submodules.
    -   [#&#8203;815](https://togithub.com/pre-commit/pre-commit-hooks/issues/815) PR by [@&#8203;asottile](https://togithub.com/asottile).
    -   [#&#8203;707](https://togithub.com/pre-commit/pre-commit-hooks/issues/707) issue by [@&#8203;ChiefGokhlayeh](https://togithub.com/ChiefGokhlayeh).

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

This PR has been generated by [Renovate Bot](https://togithub.com/renovatebot/renovate).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @app-sre-bot on September 17, 2024 at 12:24 PM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @jlsherrill - Approved on September 18, 2024 at 01:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/820*
