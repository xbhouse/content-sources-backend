---
type: pull_request
number: 1064
title: "Update module github.com/klauspost/compress to v1.18.0"
state: closed
author: red-hat-konflux
created: 2025-04-06T16:09:32Z
updated: 2025-04-08T17:56:04Z
closed: 2025-04-08T17:11:30Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1064
---

# Pull Request #1064: Update module github.com/klauspost/compress to v1.18.0

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 04:09 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | indirect | minor | `v1.17.11` -> `v1.18.0` |

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.0`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.0)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.17.11...v1.18.0)

#### What's Changed

-   Deprecate Go 1.21 and add 1.24 by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1055](https://redirect.github.com/klauspost/compress/pull/1055)
-   Add unsafe little endian loaders by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1036](https://redirect.github.com/klauspost/compress/pull/1036)
-   fix: check `r.err != nil` but return a nil value error `err` by [@&#8203;alingse](https://redirect.github.com/alingse) in [https://github.com/klauspost/compress/pull/1028](https://redirect.github.com/klauspost/compress/pull/1028)
-   refactor: use built-in `min` function by [@&#8203;Juneezee](https://redirect.github.com/Juneezee) in [https://github.com/klauspost/compress/pull/1038](https://redirect.github.com/klauspost/compress/pull/1038)
-   zstd: use `slices.Max` for max value in slice by [@&#8203;Juneezee](https://redirect.github.com/Juneezee) in [https://github.com/klauspost/compress/pull/1041](https://redirect.github.com/klauspost/compress/pull/1041)
-   flate: Simplify L4-6 loading by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1043](https://redirect.github.com/klauspost/compress/pull/1043)
-   flate: Simplify matchlen (remove asm) by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1045](https://redirect.github.com/klauspost/compress/pull/1045)
-   s2: Add block decode fuzzer by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1044](https://redirect.github.com/klauspost/compress/pull/1044)
-   s2: Improve small block compression speed w/o asm by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1048](https://redirect.github.com/klauspost/compress/pull/1048)
-   flate: Fix matchlen L5+L6 by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1049](https://redirect.github.com/klauspost/compress/pull/1049)
-   flate: Cleanup & reduce casts by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1050](https://redirect.github.com/klauspost/compress/pull/1050)

#### New Contributors

-   [@&#8203;tcpdumppy](https://redirect.github.com/tcpdumppy) made their first contribution in [https://github.com/klauspost/compress/pull/1021](https://redirect.github.com/klauspost/compress/pull/1021)
-   [@&#8203;sam9291](https://redirect.github.com/sam9291) made their first contribution in [https://github.com/klauspost/compress/pull/1022](https://redirect.github.com/klauspost/compress/pull/1022)
-   [@&#8203;dezza](https://redirect.github.com/dezza) made their first contribution in [https://github.com/klauspost/compress/pull/1023](https://redirect.github.com/klauspost/compress/pull/1023)
-   [@&#8203;alingse](https://redirect.github.com/alingse) made their first contribution in [https://github.com/klauspost/compress/pull/1028](https://redirect.github.com/klauspost/compress/pull/1028)
-   [@&#8203;hyunsooda](https://redirect.github.com/hyunsooda) made their first contribution in [https://github.com/klauspost/compress/pull/1031](https://redirect.github.com/klauspost/compress/pull/1031)
-   [@&#8203;Juneezee](https://redirect.github.com/Juneezee) made their first contribution in [https://github.com/klauspost/compress/pull/1038](https://redirect.github.com/klauspost/compress/pull/1038)
-   [@&#8203;Bbulatov](https://redirect.github.com/Bbulatov) made their first contribution in [https://github.com/klauspost/compress/pull/1052](https://redirect.github.com/klauspost/compress/pull/1052)

**Full Changelog**: https://github.com/klauspost/compress/compare/v1.17.11...v1.18.0

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYWluIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on April 08, 2025 at 05:56 PM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v1.18.0`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1064*
