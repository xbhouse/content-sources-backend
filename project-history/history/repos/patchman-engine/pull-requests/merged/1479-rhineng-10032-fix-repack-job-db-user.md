---
type: pull_request
number: 1479
title: "RHINENG-10032: fix repack job db user"
state: merged
author: Dugowitch
created: 2024-09-18T12:08:51Z
updated: 2026-04-02T11:25:17Z
closed: 2024-09-19T10:21:38Z
merged: 2024-09-19T10:21:38Z
base_branch: master
head_branch: RHINENG-10032
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1479
---

# Pull Request #1479: RHINENG-10032: fix repack job db user

**Author**: @Dugowitch
**Created**: September 18, 2024 at 12:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-10032`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [ ] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [x] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Discussion

### Comment by @jira-linking on September 18, 2024 at 12:08 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-10032


### Comment by @codecov-commenter on September 18, 2024 at 12:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1479?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 65.06%. Comparing base ([`45c6c36`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/45c6c3638763cb40c0ae2b4417012405d2fd2b27?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e3b81ea`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e3b81eac3d26d384b8b24504233e4fb6cb101c47?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1203 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1479      +/-   ##
==========================================
+ Coverage   65.03%   65.06%   +0.02%     
==========================================
  Files         114      114              
  Lines        7877     7877              
==========================================
+ Hits         5123     5125       +2     
+ Misses       2217     2215       -2     
  Partials      537      537              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1479/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1479/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.06% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1479?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @MichaelMraka on September 18, 2024 at 03:14 PM UTC

~~According to reorg/pg_repack#201 you may need to use  `pg_repack -k `.~~
Edit: I see there is `--no-superuser-check`.

---

## Reviews

### Review by @psegedy - Commented on September 18, 2024 at 01:41 PM UTC

it might be possible that this won't help because `admin` user is probably not a superuser and the option to use regular user seems broken in pg_repack https://github.com/reorg/pg_repack/issues/223
Can we use default `postgres` user? Try to look at how db migrations are performed, I think we're using the superuser there

### Review by @MichaelMraka - Approved on September 19, 2024 at 10:02 AM UTC

### Review by @psegedy - Approved on September 19, 2024 at 10:03 AM UTC

### Review by @psegedy - Approved on September 19, 2024 at 10:21 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1479*
