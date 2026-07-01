---
type: pull_request
number: 1807
title: "Update module github.com/gabriel-vasile/mimetype to v1.4.10"
state: merged
author: red-hat-konflux
created: 2025-08-31T08:33:14Z
updated: 2025-10-08T16:16:10Z
closed: 2025-08-31T12:40:47Z
merged: 2025-08-31T12:40:47Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1807
---

# Pull Request #1807: Update module github.com/gabriel-vasile/mimetype to v1.4.10

**Author**: @red-hat-konflux
**Created**: August 31, 2025 at 08:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/gabriel-vasile/mimetype](https://redirect.github.com/gabriel-vasile/mimetype) | `v1.4.9` -> `v1.4.10` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.10?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.9/v1.4.10?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gabriel-vasile/mimetype (github.com/gabriel-vasile/mimetype)</summary>

### [`v1.4.10`](https://redirect.github.com/gabriel-vasile/mimetype/releases/tag/v1.4.10): perfomance inprovements, tests and new formats

[Compare Source](https://redirect.github.com/gabriel-vasile/mimetype/compare/v1.4.9...v1.4.10)

This release adds support for XHTML, Lotus-1-2-3, KML, shell scripts, VSDX, OneNote, CHM and Netpbm file formats.
Changes were made to make mimetype behave more file linux `$ file --mime` utility.

https://github.com/gabriel-vasile/mimetype\_tests repo is now used for running comparisons between `mimetype` and `$ file --mime`. It contains 50 000 samples and `mimetype` identifies the same format as `$ file --mime` for ~97% of them. Results are in the [Actions tab](https://redirect.github.com/gabriel-vasile/mimetype_tests/actions).

#### What's Changed

- charset: remove dependency on x/net for parsing html in [https://github.com/gabriel-vasile/mimetype/pull/669](https://redirect.github.com/gabriel-vasile/mimetype/pull/669)
- CSV: replace stdlib reader with a parser that allocates less in [https://github.com/gabriel-vasile/mimetype/pull/672](https://redirect.github.com/gabriel-vasile/mimetype/pull/672)
- svg: make detection harder in [https://github.com/gabriel-vasile/mimetype/pull/674](https://redirect.github.com/gabriel-vasile/mimetype/pull/674)
- pdf: relax check to match file in [https://github.com/gabriel-vasile/mimetype/pull/677](https://redirect.github.com/gabriel-vasile/mimetype/pull/677)
- csv: stop mutating input byte slices; for [#&#8203;680](https://redirect.github.com/gabriel-vasile/mimetype/issues/680) in [https://github.com/gabriel-vasile/mimetype/pull/681](https://redirect.github.com/gabriel-vasile/mimetype/pull/681)
- charset: remove dependency on mime  in [https://github.com/gabriel-vasile/mimetype/pull/684](https://redirect.github.com/gabriel-vasile/mimetype/pull/684)
- mso\_office: increase limit of checked entries from 4 to 100 in[https://github.com/gabriel-vasile/mimetype/pull/685](https://redirect.github.com/gabriel-vasile/mimetype/pull/685)5
- jar: replace application/jar with application/java-archive in [https://github.com/gabriel-vasile/mimetype/pull/686](https://redirect.github.com/gabriel-vasile/mimetype/pull/686)
- Zip container improvements in [https://github.com/gabriel-vasile/mimetype/pull/687](https://redirect.github.com/gabriel-vasile/mimetype/pull/687)
- Jar first entry inside a zip in [https://github.com/gabriel-vasile/mimetype/pull/688](https://redirect.github.com/gabriel-vasile/mimetype/pull/688)
- svg+html: better handling for comments in [https://github.com/gabriel-vasile/mimetype/pull/689](https://redirect.github.com/gabriel-vasile/mimetype/pull/689)
- xhtml: add support in [https://github.com/gabriel-vasile/mimetype/pull/690](https://redirect.github.com/gabriel-vasile/mimetype/pull/690)
- misc: behave more like file in [https://github.com/gabriel-vasile/mimetype/pull/691](https://redirect.github.com/gabriel-vasile/mimetype/pull/691)
- lotus-1-2-3: add support in [https://github.com/gabriel-vasile/mimetype/pull/695](https://redirect.github.com/gabriel-vasile/mimetype/pull/695)
- Add support for zipped KML files by [@&#8203;dmlambea](https://redirect.github.com/dmlambea) in [https://github.com/gabriel-vasile/mimetype/pull/693](https://redirect.github.com/gabriel-vasile/mimetype/pull/693)
- shell: add support by [@&#8203;scop](https://redirect.github.com/scop) in [https://github.com/gabriel-vasile/mimetype/pull/694](https://redirect.github.com/gabriel-vasile/mimetype/pull/694)
- ruby: add support by [@&#8203;scop](https://redirect.github.com/scop) in [https://github.com/gabriel-vasile/mimetype/pull/700](https://redirect.github.com/gabriel-vasile/mimetype/pull/700)
- python: associate with python2 and python3 shebangs by [@&#8203;scop](https://redirect.github.com/scop) in [https://github.com/gabriel-vasile/mimetype/pull/699](https://redirect.github.com/gabriel-vasile/mimetype/pull/699)
- vsdx: add support in [https://github.com/gabriel-vasile/mimetype/pull/702](https://redirect.github.com/gabriel-vasile/mimetype/pull/702)
- oneNote: add support in [https://github.com/gabriel-vasile/mimetype/pull/703](https://redirect.github.com/gabriel-vasile/mimetype/pull/703)
- chm: add support for Microsoft Compiled HTML Help in [https://github.com/gabriel-vasile/mimetype/pull/704](https://redirect.github.com/gabriel-vasile/mimetype/pull/704)
- Netpbm: add support by [@&#8203;kenshaw](https://redirect.github.com/kenshaw) in [https://github.com/gabriel-vasile/mimetype/pull/705](https://redirect.github.com/gabriel-vasile/mimetype/pull/705)

#### New Contributors

- [@&#8203;dmlambea](https://redirect.github.com/dmlambea) made their first contribution in [https://github.com/gabriel-vasile/mimetype/pull/693](https://redirect.github.com/gabriel-vasile/mimetype/pull/693)
- [@&#8203;scop](https://redirect.github.com/scop) made their first contribution in [https://github.com/gabriel-vasile/mimetype/pull/694](https://redirect.github.com/gabriel-vasile/mimetype/pull/694)
- [@&#8203;kenshaw](https://redirect.github.com/kenshaw) made their first contribution in [https://github.com/gabriel-vasile/mimetype/pull/705](https://redirect.github.com/gabriel-vasile/mimetype/pull/705)

**Full Changelog**: https://github.com/gabriel-vasile/mimetype/compare/v1.4.9...v1.4.10

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1hc3RlciIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @jira-linking on August 31, 2025 at 08:33 AM UTC

Commits missing Jira IDs:
e031d6ed38af24fcad4d111ecef772142b813c18


### Comment by @codecov-commenter on August 31, 2025 at 08:38 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1807?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.85%. Comparing base ([`bd53bb0`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bd53bb00f9e5a8a8599e29ff47bcdba1db63d723?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e031d6e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e031d6ed38af24fcad4d111ecef772142b813c18?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1807   +/-   ##
=======================================
  Coverage   54.85%   54.85%           
=======================================
  Files         140      140           
  Lines       10875    10875           
=======================================
  Hits         5966     5966           
  Misses       4373     4373           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1807/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1807/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.85% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1807?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1807*
