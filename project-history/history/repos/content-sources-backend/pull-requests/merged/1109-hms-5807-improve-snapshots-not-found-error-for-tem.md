---
type: pull_request
number: 1109
title: "HMS-5807: improve snapshots not found error for template create"
state: merged
author: rverdile
created: 2025-05-14T18:53:20Z
updated: 2025-06-17T13:18:16Z
closed: 2025-06-17T09:25:20Z
merged: 2025-06-17T09:25:20Z
base_branch: main
head_branch: snapshot-404
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1109
---

# Pull Request #1109: HMS-5807: improve snapshots not found error for template create

**Author**: @rverdile
**Created**: May 14, 2025 at 06:53 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `snapshot-404`

## Description

## Summary
When creating a template, if the repository used does not have snapshots, the repo will be listed in the error response

## Testing steps
1. Create a repository for introspection only
2. Try to create a template using that repository. You will get a 404 with the error message "No snapshots found for the following repositories: ... "



---

## Discussion

### Comment by @jlsherrill on May 14, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-5807

### Comment by @rverdile on May 29, 2025 at 05:40 PM UTC

@coderabbitai review

### Comment by @coderabbitai on May 29, 2025 at 05:43 PM UTC

<!-- This is an auto-generated comment: summarize by coderabbit.ai -->
<!-- walkthrough_start -->

## Walkthrough

The `validateRepositoryUUIDs` method in the `templateDaoImpl` implementation was updated to improve how it checks for repositories missing snapshots. Instead of simply counting repositories with a non-null `last_snapshot_uuid`, the method now retrieves all relevant repositories and explicitly identifies those with missing snapshots. If any repositories lack snapshots, the method returns an error message listing their names. The initial check for the existence of repository UUIDs remains unchanged, and no public interfaces were modified.

## Poem

> Repositories checked, one by one,  
> Now missing snapshots can't hide or run!  
> Names are listed, clear as day,  
> No more guessing which went astray.  
> With careful eyes and gentle care,  
> The code now tells you who's not there!  
> 🗂️✨

<!-- walkthrough_end -->


---

<details>
<summary>📜 Recent review details</summary>

**Configuration used: CodeRabbit UI**
**Review profile: ASSERTIVE**
**Plan: Pro**


<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between 11b0c56c88579300b3e18a7a6782c868a6a06f6f and ba5ee3a791521c4b68fb1c8bf5b9dfc38a0a1ec4.

</details>

<details>
<summary>📒 Files selected for processing (1)</summary>

* `pkg/dao/templates.go` (1 hunks)

</details>

<details>
<summary>🧰 Additional context used</summary>

<details>
<summary>🧬 Code Graph Analysis (1)</summary>

<details>
<summary>pkg/dao/templates.go (4)</summary><blockquote>

<details>
<summary>pkg/models/repository_configuration.go (1)</summary>

* `RepositoryConfiguration` (17-38)

</details>
<details>
<summary>pkg/config/value_constraints.go (1)</summary>

* `RedHatOrg` (22-22)

</details>
<details>
<summary>pkg/dao/helpers.go (1)</summary>

* `UuidifyStrings` (24-30)

</details>
<details>
<summary>pkg/errors/dao_error.go (1)</summary>

* `DaoError` (8-14)

</details>

</blockquote></details>

</details>

</details>

<details>
<summary>🔇 Additional comments (4)</summary><blockquote>

<details>
<summary>pkg/dao/templates.go (4)</summary>

`142-143`: **Excellent variable naming and preparation for the enhanced logic! 🌟**

The introduction of `repos` and `missingSnapshotRepos` variables clearly sets up the improved error reporting functionality. This makes the code much more readable and sets the stage beautifully for the enhanced user experience.

---

`144-147`: **Smart approach to fetch all repositories for detailed analysis! 👍**

Changing from a count-based check to fetching all matching repositories is a great decision. This allows the method to provide specific details about which repositories are problematic, rather than just knowing that *something* is wrong. The query structure remains clean and efficient.

---

`151-155`: **Perfect implementation of the repository snapshot validation! ✨**

This logic elegantly identifies repositories missing snapshots by checking for empty `LastSnapshotUUID` values. The approach of collecting repository names (rather than UUIDs) will make error messages much more user-friendly, especially in UI contexts as mentioned in the previous review discussion.

---

`156-159`: **Wonderful improvement to the error message clarity! 🎉**

The new detailed error message format is fantastic! Instead of a generic "snapshots missing" message, users now get specific information about which repositories are causing the issue. The use of `strings.Join` with comma separation makes the output clean and readable. This will be especially helpful for developers and users trying to understand what needs to be fixed.

</details>

</blockquote></details>

</details>
<!-- internal state start -->


<!-- DwQgtGAEAqAWCWBnSTIEMB26CuAXA9mAOYCmGJATmriQCaQDG+Ats2bgFyQAOFk+AIwBWJBrngA3EsgEBPRvlqU0AgfFwA6NPEgQAfACgjoCEYDEZyAAUASpETZWaCrKNwSPbABsvkCiQBHbGlcSHFcLzpIACIACQBZAGUwAFYADgAGAHYueGZefCl7DDRuRFh8XGQMSsgAM3xsDHpKCnw+Br4afK9qDwZ/PoAaaPtsAWZ1Gno5SGxESj8pClp4SKH0Wlp/RAXkTBRd4MgyJiaaf3pabAp4DCIwkh6+xkHxfAwNSETuUXg6+AMNA+WQbADusDI6D8JG4+EQ6na8nmUU6rxI1DuDzQj2eNEgtHw0kgNVCTAwuG0WEw8kQJTKFSqG1wkJOFDafB2cIwCxJ+DBJwAHtwvID1F55KLEFUwqz/HCEQQXLLqIw0Cj6CyPAUBJFmF84KgyLBMAwnux0HlkAQeG0JPAlIxIs5FnUSHQBGgGABrML4OYLCiIDRuVk0aVY+w0Mr8PCi8j0O4SfBeIoDDHie7Q+XwxHKu40ZpRD4S+rtFAUtqIX5ieAfdDNWVQ6jdbiZh429MvHGt3r4+aRlmqnOKpEG1kkYWiaax3BMNiHaEAFgyS7ZHMgYPUsGhbF2aFIUcxWeiADl/XTSuVKsgGk16GitWWfPzIyO8/BpFwND/Rg0X2CUSzE+75KvIJRsAAFIgACUIYGO4ng+DCQQhIuaASNovS6h4qqwLguBlBwAD0xFENu4waPOxHkoWuBgIG9pmogNEfHRDGNBQzFgJ6PpkLQxHcN4XjEQAjKJGQAJzwRYkAAMIsGwFLIA4TguKG/SKewyD1k+QnIf4qHSug/iQJMGB5MCXwAPLkAorAWrMKJ8EIXjlK0ay+AU9pKPskDxr6NogSQfZREgDgeHcTaQDYUSxKq4XHLgVB8RQXD4YRiAkcRiXSBolwmpo1ECG0YILMRCTJOk2TjqgAUKBg5BiNarL6b4hnBMZQWsiV+BoEofC0ZOoT4HU0W5ZAAgkJG/XbNICy0PBiHqiy5ajdFbUoZ1uAbBQyyrOsYQHi1HjOQoA0qGolI6DaHVoTi/j2iQArrU+DAmvceWQOe9Q3FqfCrIgDDzAiHwbG6Hpet6GzlkZ7w8iZ2o7BaUVvVpynwfoxjgFA/H8GNK2EKQ5BUDO85KZwtr8MI06SMSsxMBdqjqFoOhYyYUCGsgqAHITxBkMoZPo5TVACqpzDOPIDOKMozOaNouhgIY2OmAY3DekQxG0Gg+DEb2fTBkQ+AcAY0RmwYskAIIAJL8yTfT0OLkv44wH2kIgGlmSQq30AABphorazQsUKnmsgAKrh9bAAiiC+5uaDIMwij/J+mr+u9mCHhUArqJAyyp8SQ6hKB7SfsgJpFJeDI3l8ViPXW8wShsednBSUSl7cxJbiy0I1BgYAYMJkC+700oAPrV9euDj9g2AOvHmD0PO3AupqhUKOcfrRUPzBTXw62RzHyDeQ6dC7d7NwWVmBytOW/xNvIqz0KSZnUO944eOQuc9OaFKYvWVCXd9jIQlnOBAWZO7l03NuRooQAReAuJGeso9E4zynoyWe89aC+2ZJCLA6hBbEkKIsYKDhEHWgzimSIYgd5oD3C7KBxJJi7EjBgm8kBIIQkoB4X2AAZNBiR6TTyPtHeORp8i4FkHBSA1sCYYFpNgd6MJQ5KmgZOJAO0UAl0vhQBGOIlDXUiC0dk5Y9yIAPB4KU7YmzwD4BBL6XMVG9GYhtBujRkDE0oICdcZj5qWPQAIOBZlwpsOEYyYMMBWR3HUPAYE+dgQOgAVgdand5CiKFJo04Hh/ASzuMgJomdPqLXMJYS2iDBZ1gRt1DwSgGC9FJlUnSY1JxwgoDOcsQldQ+PYLE6QRgoA/SKe7beT4EREBKLgG4xCWnCnaB0vgXTRQMBOBSPpyBAKmQlkoT+rss6RW5u3PR8SakjwDkk4OsJcxgVEXHL2PsXa+31jQaOOtrY9F9iGM20QBlKxVrjRs60+ZeNJlEcm7AuCizGGpKW8hGayyuqzRWhgOb2UmDPB0iBx4N2enQSelJ2nIv+ZNNAKR3QAGY0BZEkqJFIAAmUSDAlwCAAGxpDqAIRlaQBB1BSAISStA6gMHJWkNAGQ0CiVEGuLGBhUXk3UOPTF2KSBPUArQceeN9BAA== -->

<!-- internal state end -->
<!-- finishing_touch_checkbox_start -->

<details open="true">
<summary>✨ Finishing Touches</summary>

- [ ] <!-- {"checkboxId": "7962f53c-55bc-4827-bfbf-6a18da830691"} --> 📝 Generate Docstrings

</details>

<!-- finishing_touch_checkbox_end -->
<!-- tips_start -->

---



<details>
<summary>🪧 Tips</summary>

### Chat

There are 3 ways to chat with [CodeRabbit](https://coderabbit.ai?utm_source=oss&utm_medium=github&utm_campaign=content-services/content-sources-backend&utm_content=1109):

- Review comments: Directly reply to a review comment made by CodeRabbit. Example:
  - `I pushed a fix in commit <commit_id>, please review it.`
  - `Explain this complex logic.`
  - `Open a follow-up GitHub issue for this discussion.`
- Files and specific lines of code (under the "Files changed" tab): Tag `@coderabbitai` in a new review comment at the desired location with your query. Examples:
  - `@coderabbitai explain this code block.`
  -	`@coderabbitai modularize this function.`
- PR comments: Tag `@coderabbitai` in a new PR comment to ask questions about the PR branch. For the best results, please provide a very specific query, as very limited context is provided in this mode. Examples:
  - `@coderabbitai gather interesting stats about this repository and render them as a table. Additionally, render a pie chart showing the language distribution in the codebase.`
  - `@coderabbitai read src/utils.ts and explain its main purpose.`
  - `@coderabbitai read the files in the src/scheduler package and generate a class diagram using mermaid and a README in the markdown format.`
  - `@coderabbitai help me debug CodeRabbit configuration file.`

### Support

Need help? Create a ticket on our [support page](https://www.coderabbit.ai/contact-us/support) for assistance with any issues or questions.

Note: Be mindful of the bot's finite context window. It's strongly recommended to break down tasks such as reading entire modules into smaller chunks. For a focused discussion, use review comments to chat about specific files and their changes, instead of using the PR comments.

### CodeRabbit Commands (Invoked using PR comments)

- `@coderabbitai pause` to pause the reviews on a PR.
- `@coderabbitai resume` to resume the paused reviews.
- `@coderabbitai review` to trigger an incremental review. This is useful when automatic reviews are disabled for the repository.
- `@coderabbitai full review` to do a full review from scratch and review all the files again.
- `@coderabbitai summary` to regenerate the summary of the PR.
- `@coderabbitai generate docstrings` to [generate docstrings](https://docs.coderabbit.ai/finishing-touches/docstrings) for this PR.
- `@coderabbitai generate sequence diagram` to generate a sequence diagram of the changes in this PR.
- `@coderabbitai resolve` resolve all the CodeRabbit review comments.
- `@coderabbitai configuration` to show the current CodeRabbit configuration for the repository.
- `@coderabbitai help` to get help.

### Other keywords and placeholders

- Add `@coderabbitai ignore` anywhere in the PR description to prevent this PR from being reviewed.
- Add `@coderabbitai summary` to generate the high-level summary at a specific location in the PR description.
- Add `@coderabbitai` anywhere in the PR title to generate the title automatically.

### CodeRabbit Configuration File (`.coderabbit.yaml`)

- You can programmatically configure CodeRabbit by adding a `.coderabbit.yaml` file to the root of your repository.
- Please see the [configuration documentation](https://docs.coderabbit.ai/guides/configure-coderabbit) for more information.
- If your editor has YAML language server enabled, you can add the path at the top of this file to enable auto-completion and validation: `# yaml-language-server: $schema=https://coderabbit.ai/integrations/schema.v2.json`

### Documentation and Community

- Visit our [Documentation](https://docs.coderabbit.ai) for detailed information on how to use CodeRabbit.
- Join our [Discord Community](http://discord.gg/coderabbit) to get help, request features, and share feedback.
- Follow us on [X/Twitter](https://twitter.com/coderabbitai) for updates and announcements.

</details>

<!-- tips_end -->

### Comment by @swadeley on June 16, 2025 at 12:52 PM UTC

/retest

### Comment by @swadeley on June 16, 2025 at 03:05 PM UTC

/retest

### Comment by @swadeley on June 17, 2025 at 08:45 AM UTC

/retest

### Comment by @swadeley on June 17, 2025 at 01:18 PM UTC

Hi

before:
```
HTTP response body: {"errors":[{"status":404,"title":"Error creating template","detail":"One or more repositories does not have a snapshot."}]}
```

after:
```
HTTP response body: {"errors":[{"status":404,"title":"Error creating template","detail":"No snapshots found for the following repositories: [test-no-snap]"}]}
```

---

## Reviews

### Review by @xbhouse - Commented on May 14, 2025 at 09:12 PM UTC

### Review by @rverdile - Commented on May 15, 2025 at 01:21 PM UTC

### Review by @xbhouse - Commented on May 15, 2025 at 01:38 PM UTC

### Review by @xbhouse - Approved on May 15, 2025 at 01:39 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1109*
