---
type: pull_request
number: 24
title: "Update module github.com/Microsoft/go-winio to v0.6.2"
state: merged
author: red-hat-konflux
created: 2025-10-16T12:19:48Z
updated: 2025-10-24T08:47:31Z
closed: 2025-10-24T08:47:28Z
merged: 2025-10-24T08:47:28Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-microsoft-go-winio-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/24
---

# Pull Request #24: Update module github.com/Microsoft/go-winio to v0.6.2

**Author**: @red-hat-konflux
**Created**: October 16, 2025 at 12:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-microsoft-go-winio-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/Microsoft/go-winio](https://redirect.github.com/Microsoft/go-winio) | `v0.6.0` -> `v0.6.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fMicrosoft%2fgo-winio/v0.6.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fMicrosoft%2fgo-winio/v0.6.0/v0.6.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>Microsoft/go-winio (github.com/Microsoft/go-winio)</summary>

### [`v0.6.2`](https://redirect.github.com/microsoft/go-winio/releases/tag/v0.6.2)

[Compare Source](https://redirect.github.com/Microsoft/go-winio/compare/v0.6.1...v0.6.2)

#### What's Changed

- \[etw] Add String() functions, JSON field option by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#285](https://redirect.github.com/microsoft/go-winio/pull/285)
- enable dependency updates by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#287](https://redirect.github.com/microsoft/go-winio/pull/287)
- Isolate tools dependencies in tools package by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#293](https://redirect.github.com/microsoft/go-winio/pull/293)
- Update tests; run fuzzing by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#294](https://redirect.github.com/microsoft/go-winio/pull/294)
- Add support for flushing and disconnecting named pipes by [@&#8203;dgolub](https://redirect.github.com/dgolub) in [microsoft#292](https://redirect.github.com/microsoft/go-winio/pull/292)
- Add ResolvePath tests by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#276](https://redirect.github.com/microsoft/go-winio/pull/276)
- \[lint] Fix errors from [#&#8203;276](https://redirect.github.com/Microsoft/go-winio/issues/276) by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#296](https://redirect.github.com/microsoft/go-winio/pull/296)
- Switch from sycall to windows by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#295](https://redirect.github.com/microsoft/go-winio/pull/295)
- \[lint] Remove deprecated tar.TypeRegA by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#300](https://redirect.github.com/microsoft/go-winio/pull/300)
- sd.go: fix calculation of security descriptor length in SddlToSecurit… by [@&#8203;dblohm7](https://redirect.github.com/dblohm7) in [microsoft#299](https://redirect.github.com/microsoft/go-winio/pull/299)
- fix: already typo by [@&#8203;testwill](https://redirect.github.com/testwill) in [microsoft#303](https://redirect.github.com/microsoft/go-winio/pull/303)
- pipe.go: add DialPipeAccessImpLevel by [@&#8203;dblohm7](https://redirect.github.com/dblohm7) in [microsoft#302](https://redirect.github.com/microsoft/go-winio/pull/302)
- Bug: Close hvsock handle on listen error; fix tests by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#310](https://redirect.github.com/microsoft/go-winio/pull/310)
- fileinfo: internally fix FileBasicInfo memory alignment by [@&#8203;dagood](https://redirect.github.com/dagood) in [microsoft#312](https://redirect.github.com/microsoft/go-winio/pull/312)
- Update go1.21 and CI by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#315](https://redirect.github.com/microsoft/go-winio/pull/315)

#### New Contributors

- [@&#8203;dgolub](https://redirect.github.com/dgolub) made their first contribution in [microsoft#292](https://redirect.github.com/microsoft/go-winio/pull/292)
- [@&#8203;dblohm7](https://redirect.github.com/dblohm7) made their first contribution in [microsoft#299](https://redirect.github.com/microsoft/go-winio/pull/299)
- [@&#8203;testwill](https://redirect.github.com/testwill) made their first contribution in [microsoft#303](https://redirect.github.com/microsoft/go-winio/pull/303)
- [@&#8203;dagood](https://redirect.github.com/dagood) made their first contribution in [microsoft#312](https://redirect.github.com/microsoft/go-winio/pull/312)

**Full Changelog**: <https://github.com/microsoft/go-winio/compare/v0.6.1...v0.6.2>

### [`v0.6.1`](https://redirect.github.com/microsoft/go-winio/releases/tag/v0.6.1)

[Compare Source](https://redirect.github.com/Microsoft/go-winio/compare/v0.6.0...v0.6.1)

#### What's Changed

- Soften linter by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#264](https://redirect.github.com/microsoft/go-winio/pull/264)
- Bump linter, remove structcheck, ignore unhandled\_errors by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#265](https://redirect.github.com/microsoft/go-winio/pull/265)
- pkg/etw/sample: remove dependency on github.com/sirupsen/logrus by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [microsoft#267](https://redirect.github.com/microsoft/go-winio/pull/267)
- Update go.mod dependencies by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#277](https://redirect.github.com/microsoft/go-winio/pull/277)
- Add some basic bind filter functions by [@&#8203;gabriel-samfira](https://redirect.github.com/gabriel-samfira) in [microsoft#274](https://redirect.github.com/microsoft/go-winio/pull/274)
- Apply two upstream CL to mkwinsyscall by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#278](https://redirect.github.com/microsoft/go-winio/pull/278)
- Unable to create named pipe on certain WS2012, Win 10 Pro machines by [@&#8203;rcarman-r7](https://redirect.github.com/rcarman-r7) in [microsoft#280](https://redirect.github.com/microsoft/go-winio/pull/280)
- Add `fs.ResolvePath` to resolve symbolic links by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#275](https://redirect.github.com/microsoft/go-winio/pull/275)
- update linter and fix lint errors by [@&#8203;helsaawy](https://redirect.github.com/helsaawy) in [microsoft#284](https://redirect.github.com/microsoft/go-winio/pull/284)

#### New Contributors

- [@&#8203;gabriel-samfira](https://redirect.github.com/gabriel-samfira) made their first contribution in [microsoft#274](https://redirect.github.com/microsoft/go-winio/pull/274)
- [@&#8203;rcarman-r7](https://redirect.github.com/rcarman-r7) made their first contribution in [microsoft#280](https://redirect.github.com/microsoft/go-winio/pull/280)

**Full Changelog**: <https://github.com/microsoft/go-winio/compare/v0.6.0...v0.6.1>

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

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @Hyperkid123 - Approved on October 24, 2025 at 08:47 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/24*
