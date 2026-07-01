---
type: pull_request
number: 1631
title: "Update module github.com/gin-contrib/gzip to v1.2.3"
state: merged
author: red-hat-konflux
created: 2025-04-06T09:25:07Z
updated: 2026-04-03T09:56:08Z
closed: 2025-04-23T08:19:50Z
merged: 2025-04-23T08:19:50Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gin-contrib-gzip-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1631
---

# Pull Request #1631: Update module github.com/gin-contrib/gzip to v1.2.3

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 09:25 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gin-contrib-gzip-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/gin-contrib/gzip](https://redirect.github.com/gin-contrib/gzip) | indirect | minor | `v1.0.1` -> `v1.2.3` |

---

### Release Notes

<details>
<summary>gin-contrib/gzip (github.com/gin-contrib/gzip)</summary>

### [`v1.2.3`](https://redirect.github.com/gin-contrib/gzip/releases/tag/v1.2.3)

[Compare Source](https://redirect.github.com/gin-contrib/gzip/compare/v1.2.2...v1.2.3)

#### Changelog

##### Enhancements

-   [`3b246bb`](https://redirect.github.com/gin-contrib/gzip/commit/3b246bb1ab0a98b1e3685d98711669008e26d84a): chore: update dependencies to latest versions focused on security patches ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

-   [`ea4c811`](https://redirect.github.com/gin-contrib/gzip/commit/ea4c811f5cc72f7f263ea105a51d354e33a1d24d): ci: update GitHub Actions and improve test configurations ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`f75942a`](https://redirect.github.com/gin-contrib/gzip/commit/f75942add4614f96c7b0a585a3dc3af6e6bd792c): build: drop support for older Go versions and upgrade dependencies ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`aa60ac1`](https://redirect.github.com/gin-contrib/gzip/commit/aa60ac102420547f81148d57c5e08fbaa48f0162): ci: reorganize and update golangci-lint configuration ([@&#8203;appleboy](https://redirect.github.com/appleboy))

### [`v1.2.2`](https://redirect.github.com/gin-contrib/gzip/releases/tag/v1.2.2)

[Compare Source](https://redirect.github.com/gin-contrib/gzip/compare/v1.2.1...v1.2.2)

#### Changelog

##### Features

-   [`c881664`](https://redirect.github.com/gin-contrib/gzip/commit/c881664f9a6e88a9aa562f78328278ec8cd46dbd): feat: add logic to handle weak ETag headers ([#&#8203;100](https://redirect.github.com/gin-contrib/gzip/issues/100)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Bug fixes

-   [`438926d`](https://redirect.github.com/gin-contrib/gzip/commit/438926dfeb8be56f0052fc211e90d353b99738c5): fix: default auto decompressor on uncompressed or double compressed bodies ([#&#8203;91](https://redirect.github.com/gin-contrib/gzip/issues/91)) ([@&#8203;alexdyukov](https://redirect.github.com/alexdyukov))

##### Refactor

-   [`3529136`](https://redirect.github.com/gin-contrib/gzip/commit/3529136c16a1f39b15e558b1e7a6dff6a28d8ae7): refactor: improve error handling and resource management in gzip utils ([#&#8203;101](https://redirect.github.com/gin-contrib/gzip/issues/101)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`f3f028a`](https://redirect.github.com/gin-contrib/gzip/commit/f3f028a0433535f941fef3e483d45defa35a7caf): refactor: refactor middleware and improve API handling ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`4fc267c`](https://redirect.github.com/gin-contrib/gzip/commit/4fc267cc077ea76a93b5fee5d5337c074e769368): refactor: improve error handling and documentation for content encoding ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`aa5d8ee`](https://redirect.github.com/gin-contrib/gzip/commit/aa5d8ee6bb29ab6c0d4055138b4ee24a2e1b07b3): refactor: optimize gzip writer pool management and resource handling ([#&#8203;102](https://redirect.github.com/gin-contrib/gzip/issues/102)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

### [`v1.2.1`](https://redirect.github.com/gin-contrib/gzip/releases/tag/v1.2.1)

[Compare Source](https://redirect.github.com/gin-contrib/gzip/compare/v1.2.0...v1.2.1)

#### Changelog

##### Features

-   [`1b7e801`](https://redirect.github.com/gin-contrib/gzip/commit/1b7e80168e465b05dc6808954b0993a41d8e73e5): feat: improve compression level validation and default handling ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`bc3ccd1`](https://redirect.github.com/gin-contrib/gzip/commit/bc3ccd11f2bce606d8ff2059380a4a53262b97ab): feat: add new endpoint and improve middleware functionality ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Bug fixes

-   [`3bd96db`](https://redirect.github.com/gin-contrib/gzip/commit/3bd96db1406edf854eb4070ebb84e16973009fb2): fix: use `Header().Add` instead of `Header()` for setting `Vary` header ([#&#8203;99](https://redirect.github.com/gin-contrib/gzip/issues/99)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Enhancements

-   [`fa0d885`](https://redirect.github.com/gin-contrib/gzip/commit/fa0d885d737306430f120ec88aae8b811b360209): chore: update dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Refactor

-   [`262eca6`](https://redirect.github.com/gin-contrib/gzip/commit/262eca6d329a97000761f4757c382412c558feff): refactor: improve gzip handling and optimize resource management ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`79e888f`](https://redirect.github.com/gin-contrib/gzip/commit/79e888fe1e39cfaee270886eebcb3762dbf7b1c9): refactor: refactor codebase to improve maintainability and performance ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`e389346`](https://redirect.github.com/gin-contrib/gzip/commit/e389346319d2ceb6cb1bf71929c5bed35af321be): refactor: refactor codebase and update API integration ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Others

-   [`3830b8f`](https://redirect.github.com/gin-contrib/gzip/commit/3830b8f0d1ea89fd896c149b39a76184d50a5103): fix http: invalid Content-Length of "-1" ([#&#8203;97](https://redirect.github.com/gin-contrib/gzip/issues/97)) ([@&#8203;evildao](https://redirect.github.com/evildao))

### [`v1.2.0`](https://redirect.github.com/gin-contrib/gzip/releases/tag/v1.2.0)

[Compare Source](https://redirect.github.com/gin-contrib/gzip/compare/v1.1.0...v1.2.0)

#### Changelog

##### Features

-   [`cb30140`](https://redirect.github.com/gin-contrib/gzip/commit/cb301401d3ac24cd228db8a389932fefd406e077): feat: add support for server-sent events ([#&#8203;72](https://redirect.github.com/gin-contrib/gzip/issues/72)) ([@&#8203;htrendev](https://redirect.github.com/htrendev))
-   [`969f96e`](https://redirect.github.com/gin-contrib/gzip/commit/969f96e95cadd39faae9b3b5f7c960ed787e37df): feat: gzip decompress-only supports ([#&#8203;90](https://redirect.github.com/gin-contrib/gzip/issues/90)) ([@&#8203;yzqzss](https://redirect.github.com/yzqzss))
-   [`1b4a866`](https://redirect.github.com/gin-contrib/gzip/commit/1b4a86607c155042323e02dc103ca01c556db788): feat: improve compression and testing mechanisms across the project ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`53b6bc1`](https://redirect.github.com/gin-contrib/gzip/commit/53b6bc1334fba7fa9b6c1633d1d3b3dd57ab48ae): feat: introduce custom compression decision function in handler ([#&#8203;96](https://redirect.github.com/gin-contrib/gzip/issues/96)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`f5e4784`](https://redirect.github.com/gin-contrib/gzip/commit/f5e478469ba563659b73ce1302fa74ef32e5b2ab): feat: implement HTTP connection hijacking for gzipWriter ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Enhancements

-   [`9ab9665`](https://redirect.github.com/gin-contrib/gzip/commit/9ab966554096a70be3629d1f0e511e69739ca631): chore: update various Go module dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Refactor

-   [`aa0f078`](https://redirect.github.com/gin-contrib/gzip/commit/aa0f078fc2ed36b54fc6136811323d0ccfa2ff4e): refactor: refactor and optimize exclusion handling logic ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`48d0307`](https://redirect.github.com/gin-contrib/gzip/commit/48d0307c68233eb4acd3f7454d253b54086f7969): refactor: refactor `WithDecompressOnly` function and update related tests ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`9bb40fc`](https://redirect.github.com/gin-contrib/gzip/commit/9bb40fc33059840150166668829493777210c30a): refactor: refactor gzipHandler to use config struct instead of Options struct ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`bc23737`](https://redirect.github.com/gin-contrib/gzip/commit/bc23737a65ae92b532d55a2c97f4456a6a386c37): refactor: refactor header handling with constants in handler module ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

-   [`071dbd1`](https://redirect.github.com/gin-contrib/gzip/commit/071dbd18a1124899462b6d1c18bef184e012e6f2): ci: add GitHub Actions workflow for Bearer PR check ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`3b2d6e7`](https://redirect.github.com/gin-contrib/gzip/commit/3b2d6e7fa16da64c92381356153f7783a20b5ddd): build: refactor project structure and update dependencies ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`47135c9`](https://redirect.github.com/gin-contrib/gzip/commit/47135c972f825edaff7cafe62120dfa3f92aa568): build: refactor codebase and optimize API interactions ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Documentation updates

-   [`fa26419`](https://redirect.github.com/gin-contrib/gzip/commit/fa26419408bdea4b871c2b8f2c59f89d3193948e): docs: improve documentation and examples for server push with Gin framework ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`eae4239`](https://redirect.github.com/gin-contrib/gzip/commit/eae4239dc53b1aa6ddc6cf3c3347b53a774f9df0): docs: add documentation for exclusion and decompression functions ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Others

-   [`807284b`](https://redirect.github.com/gin-contrib/gzip/commit/807284b2d1812976d0cdede6f4215b9995ca1134): style: refactor codebase for consistency and improved test coverage ([@&#8203;appleboy](https://redirect.github.com/appleboy))

### [`v1.1.0`](https://redirect.github.com/gin-contrib/gzip/releases/tag/v1.1.0)

[Compare Source](https://redirect.github.com/gin-contrib/gzip/compare/v1.0.1...v1.1.0)

#### Changelog

##### Enhancements

-   [`8a54a91`](https://redirect.github.com/gin-contrib/gzip/commit/8a54a9173a46b6812aff8ba1fb65db4686352d67): chore: update dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`6c2fc87`](https://redirect.github.com/gin-contrib/gzip/commit/6c2fc876b8ad3d71aa42a1583b0dc4abddca56b7): chore: upgrade Go and dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`096f030`](https://redirect.github.com/gin-contrib/gzip/commit/096f030d801155139cfedfb0024241cb74272f02): chore: restrict Go versions and simplify CI workflow ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`506e1f1`](https://redirect.github.com/gin-contrib/gzip/commit/506e1f1b7174d47d0ddf01b3ddc033246c1869a8): chore: bump Go version to 1.21.0 for compatibility updates ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`03fade5`](https://redirect.github.com/gin-contrib/gzip/commit/03fade5fcbb0f25825f7bef4ae1fd5aecdc230ae): chore: update dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`f45010e`](https://redirect.github.com/gin-contrib/gzip/commit/f45010e4beb3622cd2768b2b8b9cbc0e3fd5de19): chore(deps): update dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`a1786dc`](https://redirect.github.com/gin-contrib/gzip/commit/a1786dc85fd644875b8c07f21a8a9d3dfb45c06c): chore: avoid un-necessary allocation ([#&#8203;92](https://redirect.github.com/gin-contrib/gzip/issues/92)) ([@&#8203;manisharma](https://redirect.github.com/manisharma))
-   [`fb15f54`](https://redirect.github.com/gin-contrib/gzip/commit/fb15f54b910ff49c95b7c0688e7147b0634761af): chore: discard gzip footer when empty body ([#&#8203;89](https://redirect.github.com/gin-contrib/gzip/issues/89)) ([@&#8203;koenno](https://redirect.github.com/koenno))
-   [`4713b05`](https://redirect.github.com/gin-contrib/gzip/commit/4713b051dda6b10841110cb89bfaa211fe970292): chore: update dependencies and GitHub actions for improved CI/CD ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Refactor

-   [`a6b2037`](https://redirect.github.com/gin-contrib/gzip/commit/a6b20374a6768725d3c57e8e692231c9f2c3ada0): refactor: refactor data reading to use io package consistently ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`4766b70`](https://redirect.github.com/gin-contrib/gzip/commit/4766b70caac8c4d5e6bcbee1b34e83be91413f84): refactor: migrate to io package for improved compatibility ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

-   [`28e7bee`](https://redirect.github.com/gin-contrib/gzip/commit/28e7bee1bf3299f13a4dd86d81a34d0328bb3863): ci: refactor codebase for improved performance and maintainability ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`bd91a35`](https://redirect.github.com/gin-contrib/gzip/commit/bd91a351047713822fc59011c76676e9cc6d9c23): ci: update golangci-lint action to version 6 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
-   [`0e140cc`](https://redirect.github.com/gin-contrib/gzip/commit/0e140ccbb3b170212d3a9062668f3614367b6f70): ci: update CI workflow and test configurations ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Documentation updates

-   [`33a5d2e`](https://redirect.github.com/gin-contrib/gzip/commit/33a5d2eb2fe2cbca4bdbb77314f2e77850ceb393): docs: remove Gitter badge from README file ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Others

-   [`33b88bc`](https://redirect.github.com/gin-contrib/gzip/commit/33b88bc9f28a2a545b9af5a7d1a855d6122e98dc): test: add gzip compression and decompression tests ([@&#8203;appleboy](https://redirect.github.com/appleboy))

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @red-hat-konflux on April 06, 2025 at 09:25 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 10 additional dependencies were updated
- The `go` directive was updated for compatibility reasons


Details:


| **Package**                              | **Change**               |
| :--------------------------------------- | :----------------------- |
| `go`                                     | `1.22.9` -> `1.23.8`     |
| `github.com/bytedance/sonic`             | `v1.12.7` -> `v1.13.2`   |
| `golang.org/x/net`                       | `v0.35.0` -> `v0.38.0`   |
| `github.com/go-playground/validator/v10` | `v10.23.0` -> `v10.26.0` |
| `github.com/klauspost/cpuid/v2`          | `v2.2.9` -> `v2.2.10`    |
| `golang.org/x/arch`                      | `v0.13.0` -> `v0.15.0`   |
| `golang.org/x/crypto`                    | `v0.33.0` -> `v0.36.0`   |
| `golang.org/x/sync`                      | `v0.11.0` -> `v0.12.0`   |
| `golang.org/x/sys`                       | `v0.30.0` -> `v0.31.0`   |
| `golang.org/x/text`                      | `v0.22.0` -> `v0.23.0`   |
| `google.golang.org/protobuf`             | `v1.36.2` -> `v1.36.6`   |

### Comment by @jira-linking on April 06, 2025 at 09:25 AM UTC

Commits missing Jira IDs:
269ce3a0508cfdd174316c3d6cde0671213a8cfb


### Comment by @codecov-commenter on April 13, 2025 at 07:18 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1631?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.16%. Comparing base ([`dde5824`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dde582420b242db97254d841074ded8e0ba5d6eb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`269ce3a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/269ce3a0508cfdd174316c3d6cde0671213a8cfb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 874 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1631   +/-   ##
=======================================
  Coverage   58.16%   58.16%           
=======================================
  Files         146      146           
  Lines       11333    11333           
=======================================
  Hits         6592     6592           
  Misses       4159     4159           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1631/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1631/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1631?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @red-hat-konflux on April 14, 2025 at 09:30 AM UTC

### Edited/Blocked Notification

Renovate will not automatically rebase this PR, because it does not recognize the last commit author and assumes somebody else may have edited the PR.

You can manually request rebase by checking the rebase/retry box above.

 ⚠️ **Warning**: custom changes will be lost.

### Comment by @MichaelMraka on April 22, 2025 at 02:04 PM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1631*
