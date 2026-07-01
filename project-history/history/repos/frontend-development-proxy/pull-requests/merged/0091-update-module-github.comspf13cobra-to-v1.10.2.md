---
type: pull_request
number: 91
title: "Update module github.com/spf13/cobra to v1.10.2"
state: merged
author: red-hat-konflux
created: 2025-12-04T08:55:10Z
updated: 2026-01-02T08:26:19Z
closed: 2026-01-02T08:26:18Z
merged: 2026-01-02T08:26:18Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-spf13-cobra-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/91
---

# Pull Request #91: Update module github.com/spf13/cobra to v1.10.2

**Author**: @red-hat-konflux
**Created**: December 04, 2025 at 08:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-spf13-cobra-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/spf13/cobra](https://redirect.github.com/spf13/cobra) | `v1.10.1` -> `v1.10.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fspf13%2fcobra/v1.10.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fspf13%2fcobra/v1.10.1/v1.10.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>spf13/cobra (github.com/spf13/cobra)</summary>

### [`v1.10.2`](https://redirect.github.com/spf13/cobra/releases/tag/v1.10.2)

[Compare Source](https://redirect.github.com/spf13/cobra/compare/v1.10.1...v1.10.2)

#### 🔧 Dependencies

- chore: Migrate from `gopkg.in/yaml.v3` to `go.yaml.in/yaml/v3` by [@&#8203;dims](https://redirect.github.com/dims) in [#&#8203;2336](https://redirect.github.com/spf13/cobra/pull/2336) - the `gopkg.in/yaml.v3` package has been deprecated for some time: this should significantly cleanup dependency/supply-chains for consumers of `spf13/cobra`

#### 📈 CI/CD

- Fix linter and allow CI to pass by [@&#8203;marckhouzam](https://redirect.github.com/marckhouzam) in [#&#8203;2327](https://redirect.github.com/spf13/cobra/pull/2327)
- fix: actions/setup-go v6 by [@&#8203;jpmcb](https://redirect.github.com/jpmcb) in [#&#8203;2337](https://redirect.github.com/spf13/cobra/pull/2337)

#### 🔥✍🏼 Docs

- Add documentation for repeated flags functionality by [@&#8203;rvergis](https://redirect.github.com/rvergis) in [#&#8203;2316](https://redirect.github.com/spf13/cobra/pull/2316)

#### 🍂 Refactors

- refactor: replace several vars with consts by [@&#8203;htoyoda18](https://redirect.github.com/htoyoda18) in [#&#8203;2328](https://redirect.github.com/spf13/cobra/pull/2328)
- refactor: change minUsagePadding from var to const by [@&#8203;ssam18](https://redirect.github.com/ssam18) in [#&#8203;2325](https://redirect.github.com/spf13/cobra/pull/2325)

#### 🤗 New Contributors

- [@&#8203;rvergis](https://redirect.github.com/rvergis) made their first contribution in [#&#8203;2316](https://redirect.github.com/spf13/cobra/pull/2316)
- [@&#8203;htoyoda18](https://redirect.github.com/htoyoda18) made their first contribution in [#&#8203;2328](https://redirect.github.com/spf13/cobra/pull/2328)
- [@&#8203;ssam18](https://redirect.github.com/ssam18) made their first contribution in [#&#8203;2325](https://redirect.github.com/spf13/cobra/pull/2325)
- [@&#8203;dims](https://redirect.github.com/dims) made their first contribution in [#&#8203;2336](https://redirect.github.com/spf13/cobra/pull/2336)

**Full Changelog**: <https://github.com/spf13/cobra/compare/v1.10.1...v1.10.2>

Thank you to our amazing contributors!!!!! 🐍 🚀

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @Hyperkid123 - Approved on January 02, 2026 at 08:26 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/91*
