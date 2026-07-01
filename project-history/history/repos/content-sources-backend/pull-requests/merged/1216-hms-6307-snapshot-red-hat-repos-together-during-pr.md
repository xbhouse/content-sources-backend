---
type: pull_request
number: 1216
title: "HMS-6307: Snapshot Red Hat repos together during process-repos"
state: merged
author: rverdile
created: 2025-09-26T13:31:50Z
updated: 2025-10-02T13:53:23Z
closed: 2025-10-02T13:53:19Z
merged: 2025-10-02T13:53:19Z
base_branch: main
head_branch: rh-snaps-all
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1216
---

# Pull Request #1216: HMS-6307: Snapshot Red Hat repos together during process-repos

**Author**: @rverdile
**Created**: September 26, 2025 at 01:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `rh-snaps-all`

## Description

## Summary
Changes the functionality of the process-repos command:
 * snapshot all Red Hat repos every hour
 * snapshot failed Community Repos regardless of failed_snapshot_count

Also refactors the InternalOnly_ListReposToSnapshot method to modularize more of it and improve the readability

## Testing steps
1. Use this CLI cursor made to add tasks to repositories: https://gist.github.com/rverdile/5c4d43d1fba474bfdd24fea030d0dd6a
2. Create some combinations of custom, community, and red hat repositories with new, old, or failed tasks
3. run `make process-repos` 
4. Snapshots should be enqueued as expected based on the tasks you've created
5. For Red Hat repos specifically, if just one of the snapshots is older than 45 minutes, all Red Hat repos should be enqueued.


---

## Discussion

### Comment by @jlsherrill on September 26, 2025 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-6307

### Comment by @rverdile on September 26, 2025 at 02:28 PM UTC

@coderabbitai review

### Comment by @coderabbitai on September 26, 2025 at 02:29 PM UTC

<!-- This is an auto-generated reply by CodeRabbit -->
<details>
<summary>✅ Actions performed</summary>

Review triggered.

> Note: CodeRabbit is an incremental review system and does not re-review already reviewed commits. This command is applicable only when automatic reviews are paused.

</details>

### Comment by @coderabbitai on September 26, 2025 at 02:37 PM UTC

<!-- This is an auto-generated comment: summarize by coderabbit.ai -->
<!-- walkthrough_start -->

## Walkthrough
- Adds config.SnapshotInterval(redHat bool) string returning "45 minutes" for Red Hat and "23 hours" otherwise.
- Refactors repository snapshot selection logic: introduces parameterized origin filtering, standalone helpers (e.g., minimumSnapshotCount, extraReposToSnapshot), and new orchestrators (listDueRepos, listDueRedHatRepos, failedReposToSnapshot).
- Updates InternalOnly_ListReposToSnapshot and ListReposWithOutdatedSnapshots to use the new flows and signatures with improved error propagation.
- Expands tests: reworks TestListReposToSnapshot with multiple Red Hat and community repos, case-driven assertions, and stale/extra snapshot scenarios.
- Adds a DaoSuite helper createTestRedHatRepository for constructing Red Hat repositories in tests.

## Estimated code review effort
🎯 4 (Complex) | ⏱️ ~75 minutes

## Possibly related PRs
- content-services/content-sources-backend#1125 — Also refactors snapshot/origin handling in pkg/dao/repository_configs.go, aligning Red Hat vs community filtering.

## Poem
> Snapshots march by ticking clocks ⏱️,  
> Red Hat’s forty-five, the rest near a day.  
> Origins now pass as tidy blocks,  
> DAO whispers, “filter this way.”  
> Tests plant seeds in crimson and gray—  
> Repos line up; stale ghosts give way.  
> Ship it with a smile today! 😊

<!-- walkthrough_end -->


<!-- pre_merge_checks_walkthrough_start -->

## Pre-merge checks and finishing touches
<details>
<summary>❌ Failed checks (1 warning)</summary>

|     Check name     | Status     | Explanation                                                                           | Resolution                                                                     |
| :----------------: | :--------- | :------------------------------------------------------------------------------------ | :----------------------------------------------------------------------------- |
| Docstring Coverage | ⚠️ Warning | Docstring coverage is 40.00% which is insufficient. The required threshold is 80.00%. | You can run `@coderabbitai generate docstrings` to improve docstring coverage. |

</details>
<details>
<summary>✅ Passed checks (2 passed)</summary>

|     Check name    | Status   | Explanation                                                                                                                                                                                                                                                                      |
| :---------------: | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|    Title Check    | ✅ Passed | The title succinctly conveys the main change by referencing the ticket and clearly stating that Red Hat repositories will be snapshot together during the process-repos command, matching the pull request’s primary objective.                                                  |
| Description Check | ✅ Passed | The pull request description follows the repository template by providing a ## Summary section that outlines the functional changes and a ## Testing steps section that details how to validate them, with clear instructions and examples for both Red Hat and Community repos. |

</details>

<!-- pre_merge_checks_walkthrough_end -->

<!-- finishing_touch_checkbox_start -->

<details>
<summary>✨ Finishing touches</summary>

- [ ] <!-- {"checkboxId": "7962f53c-55bc-4827-bfbf-6a18da830691"} --> 📝 Generate Docstrings
<details>
<summary>🧪 Generate unit tests</summary>

- [ ] <!-- {"checkboxId": "f47ac10b-58cc-4372-a567-0e02b2c3d479", "radioGroupId": "utg-output-choice-group-3338972361"} -->   Create PR with unit tests
- [ ] <!-- {"checkboxId": "07f1e7d6-8a8e-4e23-9900-8731c2c87f58", "radioGroupId": "utg-output-choice-group-3338972361"} -->   Post copyable unit tests in a comment

</details>

</details>

<!-- finishing_touch_checkbox_end -->

<!-- announcements_start -->

> [!TIP]
> <details>
> <summary>👮 Agentic pre-merge checks are now available in preview!</summary>
> 
> Pro plan users can now enable pre-merge checks in their settings to enforce checklists before merging PRs.
> 
> 	- Built-in checks – Quickly apply ready-made checks to enforce title conventions, require pull request descriptions that follow templates, validate linked issues for compliance, and more.
> 	- Custom agentic checks – Define your own rules using CodeRabbit’s advanced agentic capabilities to enforce organization-specific policies and workflows. For example, you can instruct CodeRabbit’s agent to verify that API documentation is updated whenever API schema files are modified in a PR. Note: Upto 5 custom checks are currently allowed during the preview period. Pricing for this feature will be announced in a few weeks.
> 
> Please see the [documentation](https://docs.coderabbit.ai/pr-reviews/pre-merge-checks) for more information.
> 
> Example:
> 
> ```yaml
> reviews:
>   pre_merge_checks:
>     custom_checks:
>       - name: "Undocumented Breaking Changes"
>         mode: "warning"
>         instructions: |
>           Pass/fail criteria: All breaking changes to public APIs, CLI flags, environment variables, configuration keys, database schemas, or HTTP/GraphQL endpoints must be documented in the "Breaking Change" section of the PR description and in CHANGELOG.md. Exclude purely internal or private changes (e.g., code not exported from package entry points or explicitly marked as internal).
> ```
> 
> Please share your feedback with us on this [Discord post](https://discord.com/channels/1134356397673414807/1414771631775158383).
> 
> </details>

<!-- announcements_end -->

<!-- tips_start -->

---

Thanks for using CodeRabbit! It's free for OSS, and your support helps us grow. If you like it, consider giving us a shout-out.

<details>
<summary>❤️ Share</summary>

- [X](https://twitter.com/intent/tweet?text=I%20just%20used%20%40coderabbitai%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20the%20proprietary%20code.%20Check%20it%20out%3A&url=https%3A//coderabbit.ai)
- [Mastodon](https://mastodon.social/share?text=I%20just%20used%20%40coderabbitai%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20the%20proprietary%20code.%20Check%20it%20out%3A%20https%3A%2F%2Fcoderabbit.ai)
- [Reddit](https://www.reddit.com/submit?title=Great%20tool%20for%20code%20review%20-%20CodeRabbit&text=I%20just%20used%20CodeRabbit%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20proprietary%20code.%20Check%20it%20out%3A%20https%3A//coderabbit.ai)
- [LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https%3A%2F%2Fcoderabbit.ai&mini=true&title=Great%20tool%20for%20code%20review%20-%20CodeRabbit&summary=I%20just%20used%20CodeRabbit%20for%20my%20code%20review%2C%20and%20it%27s%20fantastic%21%20It%27s%20free%20for%20OSS%20and%20offers%20a%20free%20trial%20for%20proprietary%20code)

</details>

<sub>Comment `@coderabbitai help` to get the list of available commands and usage tips.</sub>

<!-- tips_end -->

<!-- internal state start -->


<!-- DwQgtGAEAqAWCWBnSTIEMB26CuAXA9mAOYCmGJATmriQCaQDG+Ats2bgFyQAOFk+AIwBWJBrngA3EsgEBPRvlqU0AgfFwA6NPEgQAfACgjoCEYDEZyAAUASpETZWaCrI4Go0dQBsSXABIAsgDKYABsAMwADADsXEEYaNyIsPi4kDZ0kH7UkBQk3PjIBKS4sJSQtNgU8BhEPBT4DNKIYHkFiO6QAHKOApRcAIwATAOhnQCqNgAyXLC4uEkcAPRLROqw2AIaTMxLTBg0B2CIlBLwTYh7+Afsx/hVF2ACaAwA1mS0S9zYXl5Lw6MJicKFwKFIKLR4D5OkF7hQmpABFQMAxYKDYMcEkkwGhfp1oM4SojkaiuMxtBgYbhqNhEFx8NwyJ0MmcSAB3Sh0yAACgw1xIkC8SBotAAlJ0pioSF4uQBHEhgGiIcS1MDkOh0ToAYTy1Ey1C4Q0iQwArGBIgBOMBDULQAbhDjhAYcE2RABaRgAItIGNVuOJrjyAGYvXDIXj4M5KMVuKBa2CYUhFMqIkgJs74Pj4IOQUoCiMXFptQoKJwYWhuSC6SDxRLJVLoX7pTLZNLF5AkcHyFJVDQGKtQWtJFJpENQzJaljMbAYdTydu5EhEZy0HyIZDZyBjny0AD6iCx9dwu6YM80zJIIbEmeQAEkbhQEl4APIYLyyXdTYUZdrQfBDo9c3wSBmEUH5nHgAAvAVMHoeBmAjKRFzQWgVChOc+wMaBpBVOplXycMGijOg3AGDRIHGE5cxTLUplvRgqkQTMtwaZhIDmBY6RWNZlQ0NZSk2bYWCWMFKEhHwlhNBgABZaGk8JaAGINnmk6JpIEINaFoIZpKDEg0EiKJaEiLTQjQID0C03M0EQV4imA9t1EzeBpD7IZyJ1fSaFLNQEgDDANxzBhaQIZgABpS2nWdcFkCLYObehW0XdonOqaRIDZdZIHINkIvwLxaDyvht0yalbMQPtwnImwZ0gAADcl3nqRpmlafJCjqvtpPIgA1Sh4CDeQDzrEdkGcAUyFlbASGm+gXiYCEajqAhqIFX0vNKmy7L7E1yIAMWYjJEpyBdEEZBgBvOXF3wigb0AweR+XsQ8RxQDcCvKUpMEgaSTRAmo8GkOKmyOrITva5B6x+eg+kgSbptmzCLEgSdWHYSHHHJFxY0gIQZTKChql+Lgvwwd56AAKXgKg3ocAUOMWFYkDpiq8loBNNB2JYkXwNkTiWQIQgiGI+ygUTFp8LgMimnDMjyVk2RYlhIFpcomCUKhVHUbQeQAIgAAXV5QtepHR5ZctlddFUWFA1lQ1FNgBtARUgAXS4ABBPBCFIcgqBFZL3xQcsrtwlHFBIGx7fUdAxHga46X7asWQt3NqiIUg2b7AdulSAUvrSSclCjk23vu4P1rYA5cUXBX7FkfC2Pi2h8HSvk2wVc32UbXVaHnTsLcyHZmHURAAG5Vsi+LEm4IV0uuIO2TKLA0G98lxAYWuLbGvIeFXk5aEw/RjHAKAPn4HNV4IYgyGUAPh/YLheH4YRRHEKQZHkI3NYdrQdGPkwHgEDIFQN9K+Ptb7+yHlOR+uQ0CKwcE4FwiIv4Rx/uoP+ugwCGBPqYAw3BXhECuBgIM8AiESFxNNE8CdcBUBqGGPi+A3C6xYQYZGHtbw3z9nqegiCsaPSCgmWo0gjBQA9lpMa2Vu7fAEEKTeQYZxx0DABEc94aBglxNyNmSUXb5VFPYWhS1qLg1wFUAKkBda/X+hgQGiBdZbmYto6gABeWh017r0F1kMcI7E4R2P4HmCgmUTjkTgKgBRKJ/Jl1VvQFaOxvjeXMsNYcDYGAoTIAieKqAahKEZOWdgF84YAA9hRGK8PgNYm8SBFKaP6BxfA8xSMVkGTM68w7KmqLUCqnQ1ENEqBcdATSUAIUzGkTcQZmBpHivsJiQpUJKmGQUCgYZg4rXMrIxorw4paSMRMqZ5SunwCULmeCApuR7OiQfFBNYXqpDUacXE1tc6MCEUmCy1TSm1EFBU84qZZDXFicAyykJ/JI0sOYSwHsvDqOoPHcxK1GlKAYF4ZwsKE6FOqUsgOzEZFyLhgcdQLkOidD2ooqJKFoxcDqiou5D4KFeC0XQHR+A9EGM6UQOqwc95vDQKQeq+xSEcuDOOJ+hDiGCqWPSqhMzaEUgYUQfA4oDAAFFlTwR4bbAUXdFaXhacsrgAQ6DwEcAYFhutRF4IIUQ1C+ARIQ1Sh+AVZCKoKuYaw9hnDfZ30yHw5wAiXmJhEdqa44IVknHBDXRyBAXBgCdXUPIKKA5sFKIoZAQZWKLiaJISgTx7jlh5HwKNmZZCThIWQz0aB8C3gQl4fRK1uAvFeLyhUPgpBeH4HwZUsFcRPQiUo8x3ISB8Q0BFEes5pzMBpbgScZ4PHFNlT+Qof4p2imBuooxKRFZ5iongdC4h0rjWDhIfA5M52NSMXmNiOT8gfAKZuRpnoPbPmDl2lEQ6ekHD6dgAZ5kgw+BKbIgUySjzUkA4u0tgq01krhcYtIfJFbzXyCs76zk1jmMQHI/OwFSHQvKEvc4sAO1kJqDvNaCcjmUDoBFNoKKLpfPMrweOxV4BFLoGAVDNRIAUOqJgc8UAfy0fngVdiK5Y0R3oBxrAOH11fMyqUQZDaqDJv6tBOa3AIwvEI3qrjEF7iQ1ubgNqibMgyzSsgQdw6IpCmVJ6aai7EBWeFLZyOTLqD2brcBBw6nRmIlSIR0GSV4p8gwGAALORJPIDOOZfYNAimGdoNUKQUmoQyaIDbDIV5o3JgFH+3m9TBSfPjRDCo7jtPAdGrdFEXhsCQi+UI1cRjNxdp8M9EaqQHPzqoMgL1/sYMrUQLCxAg1rHwUcK1lJaRTwHA61km4RBeuBk3GO0bV66U13KZUm24jaDIByluaD6L+vYG88s0ss9qkdtRDhBbWAmsGeQrV5oPJrO4Gc/ZxzNm7OudwO9rc2gdz2eXQZ/R0z2BUCFNBNN/2qj01gkKL5ZX7t5EMa5C8bJMx2UgPcx8uJXzvk/N+CGQO2tTPzQmjV9DKBPi3OUvmFkdi+QFMF0LLYcjxTC2kF7RiG2lA6/hlrvBpCnF2f9n1936vw7qLNmg83YUI8zE0EStVfSBjKq8b5m3OjiKECFZAX5lT2YAOrrGfHgeZdAp32RVlRRpx3ze8IM6Bnw4HriQfsGQvyMOtsSNyAR8olAGh8Al0Y+KyOzFpuYuG5Q7aeewHM0OtLEUA84oaA25cUTOMvbe8V0P0oNVlC8IyCgiBrZGFVeIde0DjnarhkGPVnAshkNgKa1h7hLViptXalK0bHWu+dbuJUmhXUt/NR6rh3reGYz9YU1EgbiX8fZBjzI2FlT65+0T/8Bmn7mz00HdaGqklLRa8WfLjTB8095hPeDjBdQLO+tU9R1PiwRVwOjhKYMO7tALRiYsYABhzooB/3amtFXTnWiynBnDnGSmAjk0I0xTkRjmfHTk4ylVcix0/TAh/SGUHy1BsgVAS2zRXnUwaE0wyiynMhwLwMhkwwsgIKQhsmBH8g6xoxeAvRTEY2Yj/RyBODGSCioJtl6UwIPR2X8hrnP243gBUDXDnQrw3WlCL3M3r1kD2mzQCABiVA9gVQilkFlghDQFkBB3zV/VGR0KbzSDHUBlWQeXbTKyu0qElw13OHS3yEE2QEhDyDEB4BzTSR3W4HN0CmgIUDLSICqDRSwFgMGXKQZH4HBFzBwlwJOA6xqGRRqyMTtzlzqBKj3HK1SGoTPGo3amYE+EQA9xpF3lAjeBm3zSYmTVOWQHSLDkzCIF3COQAPal3GwGwFaO0zV26TER932BBThXWx+U3n6xIH1CwAfxpjV3y0AMXSGAymXnfySmLH/1QBmQozZhKywwMVxCAwMyqNiWAgaIFFHkgElGVCnQJFsljgWlqyIHfBtmVTiw+GQBiwaBsNpy5TPxwgsnBAGnkALiWLIASlbEXXWI7CKXOgDkqF2KSWpBaxyNwAiiOnBPakWNARlGAmSOq2jB+JTDyAcGhT6MgC9iICrhWQBNITSQzywEaVIWpy7T5TWXXEoDSEFxODfUKRdnkzP0JBIADhP1ggiiXlBJnjkSkJIFRO+wWIJIOJJwK14k6AAGkJikgAkCY4jlRkAUUKBSAg56FQxEQ8BLJEgVlMU34wird5opxFA9RJ5pxoV4A2oCh7BBTjsZCUw9tB9Y08CKhEtQTcs2QbYrBNg8UPYrB6Jm5qAESKBv1TEiTFxyRONFFXk6AJ5Z9hFSMgjGTSpgJz8eDjsIofCBQeswi4ojC2Tlk4VukjA2FIVcMbsrdEVRA9TrSMVoTRlMgcVwzfl2BCUg0oAuh+Qy81VK96AjYt5u5dVRkDUjUTUzULUwAjArUlhO8HB1ASAB8cJGE3VR9IVPVIENVfVkFNwsykxREySfdzIfS/iC8i99tIkYN98aAV919aB0Tu9i0LIK1/wujvJgSZV4yxAxp80FDhRJEOdoCSie851zi41Qiol01lZvpIzOEhQmoagEkIpaQjFQIlAvBFRCB0KgjwQSiE5Kz6Aw9HxstFxiSw5SKi1kEiSCgAoh0YAUw+0okTg6BkBmL5BxhxhbxPQOseDkAkDiN6TgI0TqA4p1wPcJL9SRKqL3S+LBlkQW42IUU+gvByJbwOTKA9VmAig39PQAAhG/DaczezB1OdWynvCDMhJCuFUAlpX4XmTIOQBTBNfAFCSAUzeQaeLcGoGufYcivrYCUi0haUHbVMPVLVD0x8NgxKs6BODiwQ/pdKVWQpBg9k/LZPPgK7SoudTFWCIoP43dIUfdCPPgaCgSqysIzCR2QfQE12CFMkpsjshFFMJFdspgzsrFHsvgXFfsglGqowEc8gEfC1HBQBfFCTS+b2cfKBScmBA4UEeBewKfZBby7+aOTQHWABU+SKdQFonbXcbVOgfcakU7Y+AweagYBgAADgUiGGiFCB0hNFkmkkvA+vCCDHtGiHCAtAGBID6CDAYGNFCAYFoHCCGCNCwTmpOuHjOqOUQEuoHnZGuvPmOogHqG3LYH1O3OKrshuucDSHuoAG8k5dYkBbBzLyk3g6BUZKSrBCgRRdYuAQwZRpTaaoYCpGbNlbAua/teawpaakBnxwRqgtIyBRaeaTgJaqxdZIRaAaoMBPRGgghDEul4xRBXhRa3E+aVa1aNbPBcAfB9a3gjb4yTaLEzaZxvREBfR4B/Q4VrbDauBjblaLF4dyZbx1xppEAdaKBRaWFfbdYUVlRPaMgGK7EuBHYk4qwaaqw06LFSaug0A2Bw6LaWtPbdZfb07dYu1TEE605poi607dZ4DMAwjc6UxxBLagNv06MxA98Q0SAG5HSKQA1hFrk8g9I8gUQUqTlmbSdJyfBnAg5S62Cch6r7Vo0iUyCmxYZkSgISgtTKh2VJ4CxWoFxh4RSQJqBURR7vgmw8gZZlRABMAkInVTPNfjjikA0ELuTurqJPyl3WuHDpHL7reUPUvq6O2J6Mbu8BghATSA2OuAuiVvsHOkugYFUvmiQulH7j/TfjooWhywOyfCgNOJfqrpVoIpIHDrZGcFnFqFfvTpVo8L7xCLyAVtxCVrfpVsk1xE9qzpzq4F1ibp8HNXToAF8q7U7qHdZM7s6SHuHnbXb3bAwC7CGLFS7aRbbK6WGLFa6/I4UG78wfh21AG/ilAXa/RkL8pac6KGqaAa0HTvLEIjkQ9IBkYggdqho34+sOZ+A8B4d0oGScGwr0ywK5oHHLB3yjF8INSTh+1YMKhBT/tkBN1/jcQjkHTL1RSspkV9I+ASM3F+0Anils7ztarfN5NoL4pWbICYoYKCG1HdYP7qt/If7sMqhAlgVCV0VD1yBCw/V5LGw9GSAppqYfVXG2nd5OSCl4oHA69zgXIDgg4lBTYdwqnqGLFiHSHyGloqHRHaHgiYdGHxbqm2GvAOGJHw7DGZH6m36hG36RHi7xGuGLEtaGAOkjFJwI1SANni6lHy7Fb7bi6NH66pHGgnmvkmBXmzjkBpJIgNBIhIgABSJYgjMuEjbASZi6dgUJQkvpoB0qWAIklIYTVAZ6yF6FmFxZzZ6QT++p7hgATXuEYG+njKwDqkNjQQOp1nLO8hbked1qIEQE5RWngkQgFE5aBbqBBeUFIFJeLpWe4bIeSsoYUd1gOaObudVsBe5bsQuaTldkjujtwFsGkeMa0e4ZNGiCGAEHCBNHCGeuiCDBMmiBNBNCDCGDoEiGeueqNAEFCGBuiFoAtAiBIBNBIEiCaAmOeotAdeGGeoYCiHDdUBNGdZMhNAEGepNA2ajpsj1ZsDzskYsQiBerNbQGiAEAtF+tdAtGha0jQFCBIDBvCA+ues0n0hNAGCkmknLciCeoEF9ekiUjDerdoGLboHCFEBeAEAYHNQEdwSgEF13CJtIBPDKEqN3FxuwSAA== -->

<!-- internal state end -->

---

## Reviews

### Review by @coderabbitai - Commented on September 26, 2025 at 02:37 PM UTC

**Actionable comments posted: 8**

> [!CAUTION]
> Some comments are outside the diff and can’t be posted inline due to platform limitations.
> 
> 
> 
> <details>
> <summary>⚠️ Outside diff range comments (1)</summary><blockquote>
> 
> <details>
> <summary>pkg/dao/repository_configs.go (1)</summary><blockquote>
> 
> `266-277`: **Count NULL last_snapshot_task_uuid as non-failed.**
> 
> LEFT JOIN + `tasks.status != 'failed'` excludes NULLs. Include repos with no prior task in the count.
> 
> 
> 
> ```diff
>  func minimumSnapshotCount(pdb *gorm.DB, runsPerDay int) int {
>  	var totalCount int64
>  	query := pdb.Model(&models.RepositoryConfiguration{}).
>  		Joins("LEFT JOIN tasks on last_snapshot_task_uuid = tasks.id").
> -		Where("tasks.status != ?", config.TaskStatusFailed).
> +		Where("(tasks.status != ? OR repository_configurations.last_snapshot_task_uuid IS NULL)", config.TaskStatusFailed).
>  		Where("snapshot IS TRUE").Count(&totalCount)
>  	if query.Error != nil {
>  		log.Logger.Error().Err(query.Error).Msg("Could not calculate total repo count")
>  		return 0
>  	}
>  	return (int(totalCount) / runsPerDay) + 1 // remainder will be less than runsPerDay, so just add 1 each time
>  }
> ```
> 
> </blockquote></details>
> 
> </blockquote></details>

<details>
<summary>📜 Review details</summary>

**Configuration used**: CodeRabbit UI

**Review profile**: ASSERTIVE

**Plan**: Pro

<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between 20bb28b3ba9fa37782e31d1119a32e71263b886a and 1c83d2762454d4ef763f137391ebefc026cd3220.

</details>

<details>
<summary>📒 Files selected for processing (4)</summary>

* `pkg/config/value_constraints.go` (2 hunks)
* `pkg/dao/repository_configs.go` (3 hunks)
* `pkg/dao/repository_configs_test.go` (1 hunks)
* `pkg/dao/suite_test.go` (2 hunks)

</details>

<details>
<summary>🔇 Additional comments (6)</summary><blockquote>

<details>
<summary>pkg/dao/suite_test.go (1)</summary><blockquote>

`8-16`: **LGTM on imports.**

assert is used below; import looks right.

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs_test.go (3)</summary><blockquote>

`2499-2511`: **Nice multi-origin setup.**

Creating external, RH (x2), and community repos upfront makes the cases clear.

---

`2660-2692`: **Test logic looks solid; covers RH fan‑out scenario.**

Great call ensuring RH2 appears when any RH repo is stale.

---

`2690-2719`: **Nice verification for RH group inclusion.**

This assertion captures the “all RH repos together” requirement well.

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs.go (2)</summary><blockquote>

`307-340`: **Non‑RH staleness query looks correct.**

Good grouping: (queued_at <= now - interval AND completed) OR no task. Filter propagation and “extras” handling are tidy.

---

`461-466`: **LGTM on outdated snapshots query.**

Using the shared snapshottable filter keeps origin handling consistent.

</blockquote></details>

</blockquote></details>

</details>

<!-- This is an auto-generated comment by CodeRabbit for review status -->

### Review by @xbhouse - Approved on October 01, 2025 at 02:21 PM UTC

this looks good to me, nice job!! :tada: thanks for providing that CLI, it was very useful :) 

(also coderabbit's poems crack me up)

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1216*
