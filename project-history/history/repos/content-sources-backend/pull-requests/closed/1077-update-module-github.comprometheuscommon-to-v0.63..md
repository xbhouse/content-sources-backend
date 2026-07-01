---
type: pull_request
number: 1077
title: "Update module github.com/prometheus/common to v0.63.0"
state: closed
author: red-hat-konflux
created: 2025-04-13T18:06:01Z
updated: 2025-10-07T12:16:28Z
closed: 2025-04-14T15:06:21Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-common-0.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1077
---

# Pull Request #1077: Update module github.com/prometheus/common to v0.63.0

**Author**: @red-hat-konflux
**Created**: April 13, 2025 at 06:06 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | indirect | minor | `v0.62.0` -> `v0.63.0` |

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.63.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.63.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.62.0...v0.63.0)

#### What's Changed

-   Making the map a public variable for promtheus-operator by [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) in [https://github.com/prometheus/common/pull/741](https://redirect.github.com/prometheus/common/pull/741)
-   setup ossf scorecard and codeql workflows by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [https://github.com/prometheus/common/pull/564](https://redirect.github.com/prometheus/common/pull/564)
-   feat(promslog): implement reserved keys, rename duplicates by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [https://github.com/prometheus/common/pull/746](https://redirect.github.com/prometheus/common/pull/746)
-   Bump golang.org/x/oauth2 from 0.24.0 to 0.25.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/750](https://redirect.github.com/prometheus/common/pull/750)
-   Bump golang.org/x/net from 0.33.0 to 0.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/749](https://redirect.github.com/prometheus/common/pull/749)
-   Bump google.golang.org/protobuf from 1.36.1 to 1.36.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/751](https://redirect.github.com/prometheus/common/pull/751)
-   promslog: Make AllowedLevel concurrency safe. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [https://github.com/prometheus/common/pull/754](https://redirect.github.com/prometheus/common/pull/754)
-   Fix typo 'the an' by [@&#8203;petern48](https://redirect.github.com/petern48) in [https://github.com/prometheus/common/pull/752](https://redirect.github.com/prometheus/common/pull/752)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/common/pull/757](https://redirect.github.com/prometheus/common/pull/757)
-   build(deps): bump google.golang.org/protobuf from 1.36.3 to 1.36.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/756](https://redirect.github.com/prometheus/common/pull/756)
-   build(deps): bump google.golang.org/protobuf from 1.36.4 to 1.36.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/761](https://redirect.github.com/prometheus/common/pull/761)
-   build(deps): bump github.com/google/go-cmp from 0.6.0 to 0.7.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/763](https://redirect.github.com/prometheus/common/pull/763)
-   build(deps): bump golang.org/x/net from 0.34.0 to 0.35.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/762](https://redirect.github.com/prometheus/common/pull/762)
-   model: Clarify the purpose of model.NameValidationScheme by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [https://github.com/prometheus/common/pull/765](https://redirect.github.com/prometheus/common/pull/765)
-   Fix spelling mistake in godoc by [@&#8203;grobinson-grafana](https://redirect.github.com/grobinson-grafana) in [https://github.com/prometheus/common/pull/766](https://redirect.github.com/prometheus/common/pull/766)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/common/pull/767](https://redirect.github.com/prometheus/common/pull/767)
-   otlptranslator: Add dependency free package that translates OTLP data into Prometheus metric/label names by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [https://github.com/prometheus/common/pull/768](https://redirect.github.com/prometheus/common/pull/768)

#### New Contributors

-   [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) made their first contribution in [https://github.com/prometheus/common/pull/741](https://redirect.github.com/prometheus/common/pull/741)
-   [@&#8203;petern48](https://redirect.github.com/petern48) made their first contribution in [https://github.com/prometheus/common/pull/752](https://redirect.github.com/prometheus/common/pull/752)

**Full Changelog**: https://github.com/prometheus/common/compare/v0.62.0...v0.63.0

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

*Archived from: https://github.com/content-services/content-sources-backend/pull/1077*
