---
type: pull_request
number: 1125
title: "HMS-8668: split red_hat repos into multiple files"
state: merged
author: jlsherrill
created: 2025-06-12T15:03:34Z
updated: 2025-07-16T01:51:51Z
closed: 2025-07-16T01:51:48Z
merged: 2025-07-16T01:51:48Z
base_branch: main
head_branch: 8668
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1125
---

# Pull Request #1125: HMS-8668: split red_hat repos into multiple files

**Author**: @jlsherrill
**Created**: June 12, 2025 at 03:03 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `8668`

## Description

Also fixes HMS-8716
## Summary
splits the red hat repos json file into multiple
also adds 6 new repos, supplementary repos for 8,9,10 x86_64 and aarm64.
repositories are largely organized by feature, while splitting RHEL into mulitple versions as well.
unifies red hat repo and community repo importing

## Testing steps

make repos-import

make db-cli-connect

```
content=# select count(*) from repositories where origin = 'red_hat';
 count 
-------
    31
(1 row)
```

On main, this would be 25, 25 + 6 supplementary repos is 31.

---

## Discussion

### Comment by @jlsherrill on June 12, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-8668

### Comment by @jlsherrill on June 23, 2025 at 11:23 AM UTC

/retest

### Comment by @jlsherrill on July 07, 2025 at 07:28 PM UTC

/retest

### Comment by @jlsherrill on July 07, 2025 at 07:59 PM UTC

/retest

### Comment by @jlsherrill on July 07, 2025 at 08:00 PM UTC

@coderabbitai review

### Comment by @coderabbitai on July 07, 2025 at 08:00 PM UTC

<!-- This is an auto-generated reply by CodeRabbit -->
<details>
<summary>✅ Actions performed</summary>

Review triggered.

> Note: CodeRabbit is an incremental review system and does not re-review already reviewed commits. This command is applicable only when automatic reviews are paused.

</details>

### Comment by @jlsherrill on July 07, 2025 at 08:02 PM UTC

/retest

### Comment by @coderabbitai on July 07, 2025 at 08:03 PM UTC

<!-- This is an auto-generated comment: summarize by coderabbit.ai -->
<!-- walkthrough_start -->

## Walkthrough

This change refactors repository management by consolidating Red Hat and community repository handling into a unified structure and workflow. It renames the `Arch` field to `DistributionArch` and adds an `Origin` field in repository-related structs and JSON data, ensuring clear source identification. The DAO interface and implementation are updated to merge separate refresh methods into a single method that supports both Red Hat and community repositories. Several new JSON files are added to represent snapshotted repositories for various architectures and versions, while obsolete files and code for community repositories are removed. Test cases and import logic are updated to align with the new structure.

## Possibly related PRs

- **content-services/content-sources-backend#1117**: Refactors and renames repository origin and architecture fields, consolidates repository refresh methods, and introduces schema changes to support multiple origins. This PR shares a strong code-level connection in how repository origins and refresh logic are managed.

## Poem

> 🌟  
> Snapshots and origins, now side by side,  
> Distribution arches, no longer must hide.  
> JSONs are sorted, the old ways replaced,  
> One method to rule them, efficiency embraced.  
> With structure refined and tests all anew,  
> Here’s to the code—so shiny, so true!  
> 🚀

<!-- walkthrough_end -->
<!-- internal state start -->


<!-- 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 -->

<!-- internal state end -->
<!-- finishing_touch_checkbox_start -->

<details open="true">
<summary>✨ Finishing Touches</summary>

- [ ] <!-- {"checkboxId": "7962f53c-55bc-4827-bfbf-6a18da830691"} --> 📝 Generate Docstrings

</details>

<!-- finishing_touch_checkbox_end -->
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

<details>
<summary>🪧 Tips</summary>

### Chat

There are 3 ways to chat with [CodeRabbit](https://coderabbit.ai?utm_source=oss&utm_medium=github&utm_campaign=content-services/content-sources-backend&utm_content=1125):

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

### Comment by @rverdile on July 10, 2025 at 08:15 PM UTC

This looks good, there's just one small comment from coderabbit that I also had a question about

### Comment by @rverdile on July 10, 2025 at 08:26 PM UTC

Oh and I want to point out, I'm seeing
```
content=# select count(*) from repositories where origin = 'red_hat';
 count 
-------
    83
```
I think is okay. The community EPEL PR [changed the logic](https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/repository_configs.go#L834) in SavePublicRepos, because there were red hat repos being saved with origin `external`. You might double check you ran against main.

### Comment by @jlsherrill on July 15, 2025 at 01:02 PM UTC

yep, i see 77 on main, and 83 on this branch, so i think it checks out.

### Comment by @jlsherrill on July 15, 2025 at 07:23 PM UTC

/retest

### Comment by @rverdile on July 15, 2025 at 07:55 PM UTC

I see "Popular Repositories › Test shared EPEL repos exist and cannot be edited" UI test running even though it should be skipped? The environment variable would need to be set for that test to pass. 

### Comment by @jlsherrill on July 16, 2025 at 01:51 AM UTC

merging as the previous commit was discussed in slack and now tests are green

---

## Reviews

### Review by @rverdile - Commented on June 23, 2025 at 08:08 PM UTC

This looks good. Just to differentiate from community repos, can we rename the directory to something like `snapshotted_repos/redhat`?

### Review by @coderabbitai - Commented on July 07, 2025 at 08:03 PM UTC

**Actionable comments posted: 5**

<details>
<summary>♻️ Duplicate comments (1)</summary><blockquote>

<details>
<summary>pkg/external_repos/snapshotted_repos/rhel10-aarch64.json (1)</summary><blockquote>

`9-9`: **Same feature_name verification needed for RHEL 10 ARM**

Similar to the RHEL 9 ARM file, this has `feature_name` set to "RHEL-OS-x86_64" for ARM 64 repositories. Please verify this is the intended behavior across your repository management system.




Also applies to: 19-19, 29-29, 39-39, 48-48

</blockquote></details>

</blockquote></details>

<details>
<summary>📜 Review details</summary>

**Configuration used: CodeRabbit UI**
**Review profile: ASSERTIVE**
**Plan: Pro**


<details>
<summary>📥 Commits</summary>

Reviewing files that changed from the base of the PR and between 9578835fde12990811fb1f7a7323c4498712d0df and cb7a749f1e67958e7a22a295ad7cc33d95b2b46b.

</details>

<details>
<summary>📒 Files selected for processing (20)</summary>

* `pkg/api/admin_task.go` (1 hunks)
* `pkg/clients/feature_service_client/features.go` (1 hunks)
* `pkg/dao/interfaces.go` (1 hunks)
* `pkg/dao/repository_configs.go` (1 hunks)
* `pkg/dao/repository_configs_mock.go` (1 hunks)
* `pkg/dao/repository_configs_test.go` (4 hunks)
* `pkg/external_repos/commands/import.go` (1 hunks)
* `pkg/external_repos/community_repos.go` (0 hunks)
* `pkg/external_repos/redhat_repos.json` (0 hunks)
* `pkg/external_repos/snapshotted_repos/community_epel-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/high_availability-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/openshift_ocp-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel10-aarch64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel10-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel8-aarch64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel8-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel9-aarch64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos/rhel9-x86_64.json` (1 hunks)
* `pkg/external_repos/snapshotted_repos_importer.go` (3 hunks)
* `pkg/handler/admin_tasks_test.go` (1 hunks)

</details>

<details>
<summary>💤 Files with no reviewable changes (2)</summary>

* pkg/external_repos/community_repos.go
* pkg/external_repos/redhat_repos.json

</details>

<details>
<summary>🧰 Additional context used</summary>

<details>
<summary>🧬 Code Graph Analysis (6)</summary>

<details>
<summary>pkg/external_repos/commands/import.go (1)</summary><blockquote>

<details>
<summary>pkg/external_repos/snapshotted_repos_importer.go (1)</summary>

* `NewSnapshotRepoImporter` (57-61)

</details>

</blockquote></details>
<details>
<summary>pkg/api/admin_task.go (1)</summary><blockquote>

<details>
<summary>pkg/config/value_constraints.go (2)</summary>

* `DistributionArch` (47-50)
* `DistributionVersion` (43-46)

</details>

</blockquote></details>
<details>
<summary>pkg/clients/feature_service_client/features.go (2)</summary><blockquote>

<details>
<summary>pkg/api/admin_task.go (1)</summary>

* `RedHatRepoStructure` (93-101)

</details>
<details>
<summary>pkg/config/value_constraints.go (3)</summary>

* `DistributionArch` (47-50)
* `DistributionVersion` (43-46)
* `OriginRedHat` (18-18)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/interfaces.go (1)</summary><blockquote>

<details>
<summary>pkg/api/repositories.go (2)</summary>

* `RepositoryRequest` (43-57)
* `RepositoryResponse` (9-40)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs_mock.go (1)</summary><blockquote>

<details>
<summary>pkg/api/repositories.go (2)</summary>

* `RepositoryRequest` (43-57)
* `RepositoryResponse` (9-40)

</details>

</blockquote></details>
<details>
<summary>pkg/dao/repository_configs_test.go (3)</summary><blockquote>

<details>
<summary>pkg/utils/utils.go (1)</summary>

* `Ptr` (55-57)

</details>
<details>
<summary>pkg/config/value_constraints.go (3)</summary>

* `OriginRedHat` (18-18)
* `ContentTypeRpm` (13-13)
* `OriginCommunity` (20-20)

</details>
<details>
<summary>pkg/models/snapshot.go (1)</summary>

* `Snapshot` (13-26)

</details>

</blockquote></details>

</details>

</details>

<details>
<summary>⏰ Context from checks skipped due to timeout of 90000ms. You can increase the timeout in your CodeRabbit configuration to a maximum of 15 minutes (900000ms). (1)</summary>

* GitHub Check: Red Hat Konflux / content-sources-backend-on-pull-request

</details>

<details>
<summary>🔇 Additional comments (26)</summary><blockquote>

<details>
<summary>pkg/external_repos/snapshotted_repos/community_epel-x86_64.json (1)</summary>

`8-9`: **Excellent addition of origin tracking! 🎉**

The addition of `"origin": "community"` to each EPEL repository entry is perfectly aligned with the broader refactoring effort. This enables clear distinction between community and Red Hat repositories in the unified handling system.




Also applies to: 17-18, 26-27

</details>
<details>
<summary>pkg/external_repos/commands/import.go (1)</summary>

`50-50`: **SnapshotRepoImporter signature verified! 🎉**

- Located `func NewSnapshotRepoImporter(daoReg *dao.DaoRegistry) SnapshotRepoImporter` in  
  `pkg/external_repos/snapshotted_repos_importer.go:57-61`  
- Signature and return type match the import usage in `import.go`

Everything looks good—changes approved!

</details>
<details>
<summary>pkg/api/admin_task.go (1)</summary>

`97-97`: **Fantastic API improvements for better clarity and functionality! 🌟**

The renaming of `Arch` to `DistributionArch` provides much better semantic clarity, and the addition of the `Origin` field perfectly supports the new repository origin tracking system. These changes enhance the API's expressiveness while maintaining clean JSON serialization.




Also applies to: 100-100

</details>
<details>
<summary>pkg/handler/admin_tasks_test.go (1)</summary>

`442-442`: **Perfect test alignment with the API changes! 🎯**

The test updates correctly reflect the renamed `DistributionArch` field and the new `Origin` field with the appropriate "red_hat" value. This ensures test coverage remains intact while supporting the enhanced repository structure.




Also applies to: 445-445

</details>
<details>
<summary>pkg/clients/feature_service_client/features.go (1)</summary>

`213-213`: **Excellent client-side implementation of the API changes! 💫**

The update to use `DistributionArch` and the addition of the `Origin` field with `config.OriginRedHat` constant demonstrates great attention to consistency and best practices. Using the constant instead of a hardcoded string is particularly thoughtful!




Also applies to: 217-217

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/rhel8-aarch64.json (1)</summary>

`8-8`: **Double-check `feature_name`: it still says `x86_64` for aarch64 repos**  
Looks like all four entries keep `feature_name` as `RHEL-OS-x86_64`, even though the `distribution_arch` is `aarch64`.  If the feature service relies on this field to gate access or group content, the mismatch may surface as a “no matching feature” error when users on ARM attempt to refresh these repos.  Could you verify whether a dedicated `RHEL-OS-aarch64` (or similar) feature string is expected?  



Also applies to: 17-17, 26-26, 35-35

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/openshift_ocp-x86_64.json (1)</summary>

`1-29`: **Looks perfect – nothing to tweak here**  
Fields, naming, and architecture all line up. Nice addition!

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/rhel9-x86_64.json (1)</summary>

`32-40`: **Missing `selector` on the Supplementary repo**  
The first three entries include `"selector": "rhel9"`, but the Supplementary snapshot does not. If that field is required for importer filtering, the repo might be skipped.

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/rhel9-aarch64.json (1)</summary>

`9-9`: **Review feature_name across all ARM JSON files**

- All `*-aarch64.json` files (RHEL 8/9/10) consistently use  
  `"feature_name": "RHEL-OS-x86_64"`.  
- If your feature-management system is architecture-agnostic, no change is needed.  
- Otherwise, please update `feature_name` to the correct aarch64 variant in:  
  • pkg/external_repos/snapshotted_repos/rhel8-aarch64.json  
  • pkg/external_repos/snapshotted_repos/rhel9-aarch64.json  
  • pkg/external_repos/snapshotted_repos/rhel10-aarch64.json  

Could you confirm whether using the x86_64 feature name for ARM64 is intentional? Thank you!

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/rhel10-x86_64.json (1)</summary>

`1-51`: **Excellent repository organization for RHEL 10 x86_64!**

The repository structure looks comprehensive and well-organized. The inclusion of Extensions repositories alongside the standard BaseOS, AppStream, CodeReady Linux Builder, and Supplementary repositories provides good coverage for RHEL 10 x86_64.

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos/rhel10-aarch64.json (1)</summary>

`1-51`: **Great addition of RHEL 10 ARM support!**

Wonderful to see comprehensive ARM 64 support for RHEL 10 with all the same repository types as the x86_64 version. This ensures feature parity across architectures.

</details>
<details>
<summary>pkg/dao/repository_configs_test.go (2)</summary>

`2715-2715`: **Excellent explicit origin handling!**

Love how you're now explicitly setting the `Origin` field using the config constants. This makes the test intent much clearer and aligns perfectly with the new unified repository management approach.




Also applies to: 2747-2747

---

`2719-2719`: **Smooth transition to unified refresh method!**

The migration from separate `InternalOnly_RefreshRedHatRepo` and `InternalOnly_RefreshCommunityRepo` methods to the unified `InternalOnly_RefreshPredefinedSnapshotRepo` method is very clean. The additional parameters for label and feature name provide good flexibility while maintaining the same test coverage.




Also applies to: 2730-2730, 2751-2751, 2761-2761

</details>
<details>
<summary>pkg/dao/interfaces.go (1)</summary>

`93-93`: **Excellent consolidation of repository refresh methods!**

This unification of `InternalOnly_RefreshRedHatRepo` and `InternalOnly_RefreshCommunityRepo` into a single `InternalOnly_RefreshPredefinedSnapshotRepo` method is a wonderful architectural improvement. It reduces interface complexity while maintaining flexibility through the `Origin` field in the repository request.

</details>
<details>
<summary>pkg/dao/repository_configs_mock.go (1)</summary>

`333-337`: **Perfect mock method synchronization!**

The mock implementation has been updated beautifully to match the interface changes. The method rename and corresponding panic message update maintain excellent consistency between the interface and its mock.

</details>
<details>
<summary>pkg/dao/repository_configs.go (5)</summary>

`1350-1350`: **Fantastic method unification and naming!**

The rename to `InternalOnly_RefreshPredefinedSnapshotRepo` perfectly captures the method's expanded responsibility of handling multiple repository origins while maintaining clarity about its purpose.

---

`1357-1363`: **Excellent origin-aware organization assignment!**

This conditional logic beautifully handles the different organizational contexts for Red Hat and Community repositories. The validation at the end ensures data integrity by restricting the method to only handle the intended repository origins.

---

`1371-1371`: **Smart conflict resolution strategy!**

Only updating the `public` column on conflict is a thoughtful approach that preserves existing repository data while ensuring the public flag stays current.

---

`1378-1378`: **Improved query specificity!**

Adding the origin filter alongside the URL filter is a great enhancement that ensures we're loading the exact repository record we need, preventing potential conflicts between repositories with the same URL but different origins.

---

`1388-1388`: **Enhanced upsert logic with comprehensive field updates!**

The addition of `label` and `feature_name` to the conflict resolution update columns is wonderful! This ensures that predefined snapshot repositories stay synchronized with their latest metadata.

</details>
<details>
<summary>pkg/external_repos/snapshotted_repos_importer.go (6)</summary>

`27-37`: **Wonderful struct evolution and field alignment!**

The rename from `RedHatRepo` to `SnapshottedRepo` perfectly reflects the expanded scope, and the field updates (`Arch` → `DistributionArch`, addition of `Origin`) create excellent consistency with the API layer. This makes the code much more intuitive and maintainable!

---

`39-51`: **Beautiful dynamic origin handling!**

The updated `ToRepositoryRequest` method now properly uses the dynamic `Origin` field instead of hardcoding Red Hat origin. This flexibility is exactly what's needed for the unified repository approach!

---

`53-61`: **Perfect type and constructor naming consistency!**

The rename from `RedHatRepoImporter` to `SnapshotRepoImporter` with the corresponding constructor update maintains excellent naming consistency throughout the codebase.

---

`62-82`: **Excellent enhanced loading logic with proper error handling!**

The new `LoadAndSave` method beautifully handles the directory-based approach and includes smart GPG key assignment based on origin. The conditional GPG key logic for Red Hat repositories is particularly well-implemented!

---

`96-111`: **Fantastic directory-based loading approach!**

The new `loadFromFiles` method is a wonderful improvement that provides much better scalability and organization. The error handling with contextual messages using `fmt.Errorf` is excellent and will make debugging much easier.

---

`128-141`: **Smart origin-aware filtering logic!**

The enhanced filtering that considers both repository selectors and origin-specific feature matching is brilliantly implemented. This allows for much more granular control over which repositories get imported based on their origin and features.

</details>

</blockquote></details>

</details>

<!-- This is an auto-generated comment by CodeRabbit for review status -->

### Review by @rverdile - Commented on July 08, 2025 at 03:12 PM UTC

### Review by @coderabbitai - Commented on July 08, 2025 at 03:13 PM UTC

### Review by @jlsherrill - Commented on July 15, 2025 at 12:44 PM UTC

### Review by @coderabbitai - Commented on July 15, 2025 at 12:44 PM UTC

### Review by @rverdile - Commented on July 15, 2025 at 07:15 PM UTC

### Review by @rverdile - Approved on July 15, 2025 at 07:15 PM UTC

this looks good. test failures seem unrelated.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1125*
