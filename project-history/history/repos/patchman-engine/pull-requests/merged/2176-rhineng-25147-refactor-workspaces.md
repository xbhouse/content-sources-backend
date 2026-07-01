---
type: pull_request
number: 2176
title: "RHINENG-25147: refactor workspaces"
state: merged
author: Dugowitch
created: 2026-04-28T14:31:17Z
updated: 2026-05-12T08:50:01Z
closed: 2026-05-12T08:49:55Z
merged: 2026-05-12T08:49:54Z
base_branch: master
head_branch: workspaces
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2176
---

# Pull Request #2176: RHINENG-25147: refactor workspaces

**Author**: @Dugowitch
**Created**: April 28, 2026 at 02:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `workspaces`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices

## Summary by Sourcery

Refactor workspace handling across the service to replace JSONB-based group/workspace storage with explicit workspace_id and workspace_name columns, and propagate the new workspace-based filtering through queries, middleware, models, and exports.

Enhancements:
- Simplify workspace representation in system_inventory by introducing dedicated workspace_id and workspace_name columns and indexing them.
- Update database queries, controllers, and middleware to use workspace ID lists instead of inventory group JSON for scoping systems, advisories, packages, and templates.
- Adjust system upload and model mapping to persist a single validated workspace per host and log anomalies like multiple workspaces.
- Extend API response and export payloads to include explicit workspace identifiers and names alongside existing group data.
- Revise tests and test fixtures to align with the new workspace schema and selection semantics, including updated test data and helper context setup.

Build:
- Bump schema migration version and add a migration that backfills workspace_id and workspace_name from existing JSONB workspaces before dropping the old column.

---

## Discussion

### Comment by @codecov-commenter on April 28, 2026 at 02:37 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `82.30769%` with `23 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 59.06%. Comparing base ([`6d0becb`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6d0becbac729dddf5e8b88c952b87638530a9717?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`19291c2`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/19291c23e4d3f5c564afaa01aa5f47b839919462?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 30 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&filepath=base%2Fdatabase%2Futils.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 18.18% | [8 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 60.00% | [6 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_detail.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&filepath=manager%2Fcontrollers%2Fsystem_detail.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | 50.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zLmdv) | 71.42% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/kessel.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Fkessel.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9rZXNzZWwuZ28=) | 66.66% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2176      +/-   ##
==========================================
- Coverage   59.13%   59.06%   -0.07%     
==========================================
  Files         134      136       +2     
  Lines        8738     8761      +23     
==========================================
+ Hits         5167     5175       +8     
- Misses       3028     3040      +12     
- Partials      543      546       +3     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.06% <82.30%> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2176?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @sourcery-ai on May 07, 2026 at 12:31 PM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

Refactors workspace handling across the service by replacing JSONB-based `workspaces` groups with explicit `workspace_id`/`workspace_name` columns, updating all query helpers, controllers, middleware, schema, migrations, test data, and tests to use workspace IDs (and names) instead of inventory-group JSON, and wiring workspace information consistently from uploads through RBAC/Kessel to API responses and exports.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Replace JSONB `workspaces` field with scalar `workspace_id`/`workspace_name` in schema, models, data access layer, and test data. | <ul><li>Update `system_inventory` schema to drop `workspaces` JSONB and add `workspace_id` UUID and `workspace_name` TEXT with indexes.</li><li>Add migration 153 to backfill `workspace_id`/`workspace_name` from existing `workspaces` JSONB and then drop the old column.</li><li>Adjust `SystemInventory` model struct to expose `WorkspaceID` and `WorkspaceName` instead of `Workspaces`, and update related JSON/query tags for workspace-based selection and ordering.</li><li>Rewrite dev `test_data.sql` seed rows to populate `workspace_id` and `workspace_name` columns instead of `workspaces` JSON arrays.</li></ul> | `database_admin/schema/create_schema.sql`<br/>`database_admin/migrations/153_simplify_workspaces.up.sql`<br/>`base/models/models.go`<br/>`dev/test_data.sql` |
| Refactor workspace filtering and system/advisory/package queries to operate on workspace IDs instead of group JSON. | <ul><li>Change `Systems`, `SystemAdvisories`, `SystemPackages`, and `SystemAdvisoriesByInventoryID` helpers to take `[]string workspaceIDs` and delegate to a simplified `ApplyInventoryWorkspaceFilter` that filters by `si.workspace_id IN (?)`.</li><li>Simplify `ApplyInventoryWorkspaceFilter` to drop grouped/ungrouped semantics and log a warning when called with an empty workspace slice.</li><li>Update all controller query builders (systems, advisories, packages, package versions, system advisories, system packages, templates, template systems, systems-advisories view, template systems update/delete/subscribed-systems, system tags, exports) to accept/read workspace ID slices from Gin context instead of inventory group maps.</li><li>Adjust caching decisions (`shouldUseCache`, advisory cached counts) to use workspace-based inputs, with TODOs noting conditions that are now always true.</li><li>Update group-name filter construction to filter on `workspace_name` rather than JSONB `workspaces` content.</li></ul> | `base/database/utils.go`<br/>`manager/controllers/systems.go`<br/>`manager/controllers/advisories.go`<br/>`manager/controllers/advisories_export.go`<br/>`manager/controllers/advisory_systems.go`<br/>`manager/controllers/advisory_systems_export.go`<br/>`manager/controllers/system_advisories.go`<br/>`manager/controllers/system_advisories_export.go`<br/>`manager/controllers/system_detail.go`<br/>`manager/controllers/system_packages.go`<br/>`manager/controllers/system_packages_export.go`<br/>`manager/controllers/packages.go`<br/>`manager/controllers/packages_export.go`<br/>`manager/controllers/package_versions.go`<br/>`manager/controllers/package_systems.go`<br/>`manager/controllers/package_systems_export.go`<br/>`manager/controllers/systems_advisories_view.go`<br/>`manager/controllers/template_systems.go`<br/>`manager/controllers/template_systems_update.go`<br/>`manager/controllers/template_systems_delete.go`<br/>`manager/controllers/template_subscribed_systems_update.go`<br/>`manager/controllers/template_systems_export.go`<br/>`manager/controllers/templates.go`<br/>`manager/controllers/systemtags.go`<br/>`manager/controllers/utils.go` |
| Change how workspaces are propagated from RBAC/Kessel and uploads into requests and persisted systems, ensuring single-workspace semantics per host. | <ul><li>Replace `KeyInventoryGroups`/group map usage with `KeyInventoryWorkspaces` and a `[]string` of workspace IDs in Gin context and testing helpers.</li><li>Remove Kessel `processWorkspaces` JSON-group builder; instead, collect workspace IDs directly from Kessel responses, validate non-empty, and place them under `KeyInventoryWorkspaces`.</li><li>Keep RBAC inventory-group discovery but store it under local middleware constants (`KeyInventoryGroups`, `KeyGrouped`, `KeyUngrouped`) to decouple from the new workspace key used elsewhere.</li><li>In upload pipeline, stop storing `Groups` as JSONB `workspaces`; instead, validate and parse the first `host.Groups` entry into `workspace_id` (UUID) and `workspace_name`, log and error on invalid UUIDs, and log a warning if multiple workspaces are received.</li><li>Ensure `SystemPlatformV2` upsert paths write `WorkspaceID` and `WorkspaceName` and update their column update lists accordingly.</li></ul> | `base/utils/gin.go`<br/>`base/core/gintesting.go`<br/>`manager/middlewares/kessel.go`<br/>`manager/middlewares/rbac.go`<br/>`listener/upload.go`<br/>`listener/upload_test.go`<br/>`listener/common_test.go`<br/>`manager/middlewares/kessel_test.go` |
| Expose workspace information in API DTOs and CSV exports while preserving existing groups JSON for backward compatibility. | <ul><li>Introduce `SystemWorkspace` struct with `workspace_id` and `workspace_name` fields and embed it where system attributes are exposed (systems, template systems, advisory systems, package systems).</li><li>Change `SystemGroups` query spec to synthesize groups JSON from `workspace_id`/`workspace_name` via `jsonb_build_array(jsonb_build_object(...))`, keeping the previous groups representation as a derived value.</li><li>Extend CSV headers and expectations in systems, template systems, advisory systems, and package systems export tests to include workspace ID and name columns alongside the groups field, and adjust line expectations to use UUID-based IDs instead of legacy inventory-group names.</li><li>Update tests that use group-name filters to use workspace IDs and new expectations for counts and order.</li></ul> | `manager/controllers/common_attributes.go`<br/>`manager/controllers/systems_export.go`<br/>`manager/controllers/systems_export_test.go`<br/>`manager/controllers/template_systems.go`<br/>`manager/controllers/template_systems_export.go`<br/>`manager/controllers/advisory_systems.go`<br/>`manager/controllers/advisory_systems_export.go`<br/>`manager/controllers/package_systems.go`<br/>`manager/controllers/package_systems_export.go`<br/>`manager/controllers/utils_test.go` |
| Remove obsolete JSONB workspace helper types and update tests to the new workspace model and data. | <ul><li>Delete the `inventory.Groups` JSONB valuer/scanner type now that `system_inventory` no longer stores `workspaces` as JSONB.</li><li>Adjust listener and controller tests to assert on `WorkspaceID`/`WorkspaceName` instead of `Workspaces`, including creation helpers (e.g., `createTestUploadEvent`) and store/update assertions.</li><li>Update workspace filter tests in `utils_test.go` and `systems_advisories_view_test.go` to use workspace IDs, new expectations for grouped/ungrouped counts, and the new test router context default for workspaces.</li><li>Drop now-unused Kessel `processWorkspaces` tests and update Kessel middleware permission tests to read `KeyInventoryWorkspaces` and assert workspace IDs instead of grouped JSON.</li></ul> | `base/inventory/inventory.go`<br/>`listener/common_test.go`<br/>`listener/upload_test.go`<br/>`base/database/utils_test.go`<br/>`manager/middlewares/kessel_test.go`<br/>`manager/controllers/template_systems_export_test.go`<br/>`manager/controllers/advisory_systems_export_test.go`<br/>`manager/controllers/package_systems_export_test.go`<br/>`manager/controllers/systems_advisories_view_test.go` |

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

---

## Reviews

### Review by @sourcery-ai - Commented on April 28, 2026 at 02:31 PM UTC

Sorry @Dugowitch, your pull request is larger than the review limit of 150000 diff characters

### Review by @MichaelMraka - Changes Requested on April 30, 2026 at 10:00 AM UTC

### Review by @MichaelMraka - Commented on April 30, 2026 at 12:54 PM UTC

### Review by @sourcery-ai - Commented on May 07, 2026 at 12:31 PM UTC

Hey - I've found 3 issues, and left some high level feedback:

- ApplyInventoryWorkspaceFilter now always applies `WHERE si.workspace_id IN (?)` even when `workspaceIDs` is empty (and only logs a warning); consider short‑circuiting and returning the unmodified tx (or an always‑false predicate) to avoid unexpected SQL/behavior on empty slices.
- RBAC still computes and stores inventory groups under its own KeyInventoryGroups, but the rest of the code now reads workspace IDs from utils.KeyInventoryWorkspaces and ApplyInventoryWorkspaceFilter no longer accepts groups; you should either adapt RBAC to emit workspace IDs or add a translation layer, otherwise RBAC group restrictions are effectively ignored in queries.
- The TODOs around caching conditions (e.g., advisoriesCommon and AdvisoriesExportHandler) indicate the cache path is now always disabled because there is always at least the root workspace; it would be good to either remove the dead cache branch or update the conditions to make caching usable again under the new workspace model.

<details>
<summary>Prompt for AI Agents</summary>

~~~markdown
Please address the comments from this code review:

## Overall Comments
- ApplyInventoryWorkspaceFilter now always applies `WHERE si.workspace_id IN (?)` even when `workspaceIDs` is empty (and only logs a warning); consider short‑circuiting and returning the unmodified tx (or an always‑false predicate) to avoid unexpected SQL/behavior on empty slices.
- RBAC still computes and stores inventory groups under its own KeyInventoryGroups, but the rest of the code now reads workspace IDs from utils.KeyInventoryWorkspaces and ApplyInventoryWorkspaceFilter no longer accepts groups; you should either adapt RBAC to emit workspace IDs or add a translation layer, otherwise RBAC group restrictions are effectively ignored in queries.
- The TODOs around caching conditions (e.g., advisoriesCommon and AdvisoriesExportHandler) indicate the cache path is now always disabled because there is always at least the root workspace; it would be good to either remove the dead cache branch or update the conditions to make caching usable again under the new workspace model.

## Individual Comments

### Comment 1
<location path="listener/upload.go" line_range="363-372" />
<code_context>

 	updatesReqJSONString := string(updatesReqJSON)
-	hostWorkspaces := inventory.Groups(host.Groups)
+	var workspaceID uuid.UUID
+	var workspaceName *string
+	if l := len(host.Groups); l >= 1 {
+		if idString := host.Groups[0].ID; idString != "" {
+			workspaceID, err = uuid.Parse(idString)
+			if err != nil {
+				utils.LogError("workspaceID", idString, "invalid workspace UUID")
+				return nil, errors.New("received invalid workspace UUID")
+			}
+		}
+		if host.Groups[0].Name != "" {
+			workspaceName = &host.Groups[0].Name
+		}
+		if l != 1 {
+			utils.LogWarn(
+				"host_id", host.ID, "org_id", host.OrgID, "workspaces", host.Groups,
</code_context>
<issue_to_address>
**issue (bug_risk):** Avoid persisting a zero UUID when no or invalid workspace ID is present

Because `workspaceID` is always initialized and passed by address, `SystemInventory` will always see a non-nil pointer:

- With no groups, you persist `00000000-0000-0000-0000-000000000000` instead of `NULL`.
- With an empty `host.Groups[0].ID`, you silently use the zero UUID, which is indistinguishable in the DB from a real value and doesn’t clearly mean “no workspace”.

Instead, either:
- Declare `var workspaceID *uuid.UUID` and only set it on successful parse, or
- Keep `workspaceID uuid.UUID` but only set `WorkspaceID` to a non-nil pointer when a valid ID was parsed.

This preserves the intended `NULL` semantics for missing/unknown workspaces and avoids treating the zero UUID as a valid value.
</issue_to_address>

### Comment 2
<location path="manager/middlewares/kessel_test.go" line_range="112" />
<code_context>
-	inventoryGroups, found := c.Get(utils.KeyInventoryGroups)
</code_context>
<issue_to_address>
**suggestion (testing):** Consider adding a negative test for missing or empty workspaces in hasPermissionKessel

The updated `hasPermissionKessel` now aborts with `http.StatusUnauthorized` when the collected workspace IDs slice is empty, but `TestHasPermissionKessel` only covers the success path. Please add a test where the `workspaces` stream is empty (or decoded to an empty slice) and assert that the handler returns 401 and does not set `utils.KeyInventoryWorkspaces` to ensure this failure path is covered and guarded against regressions.

Suggested implementation:

```golang
func TestBuildPermission(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{Method: http.MethodGet, Header: http.Header{}}

	permission := buildPermission(c)
	require.NotNil(t, permission)

	c.Request.Header.Set("x-rh-identity", "ewogICAgImVudGl0bGVtZW50cyI6IHsKICAgICAgICAiaW5zaWdodHMiOiB7CiAgICAgICAgICAgICJpc19lbnRpdGxlZCI6IHRydWUKICAgICAgICB9LAogICAgICAgICJjb3N0X21hbmFnZW1lbnQiOiB7CiAgICAgICAgICAgICJpc19lbnRpdGxlZCI6IHRydWUKICAgICAgICB9LAogICAgICAgICJhbnNpYmxlIjogewogICAgICAgICAgICAiaXNfZW50aXRsZWQiOiB0cnVlCiAgICAgICAgfSwKICAgICAgICAib3BlbnNoaWZ0IjogewogICAgICAgICAgICAiaXNfZW50aXRsZWQiOiB0cnVlCiAgICAgICAgfSwKICAgICAgICAic21hcnRfbWFuYWdlbWVudCI6IHsKICAgICAgICAgICAgImlzX2VudGl0bGVkIjogdHJ1ZQogICAgICAgIH0sCiAgICAgICAgIm1pZ3JhdGlvbnMiOiB7CiAgICAgICAgICAgICJpc19lbnRpdGxlZCI6IHRydWUKICAgICAgICB9CiAgICB9LAogICAgImlkZW50aXR5IjogewogICAgICAgICJpbnRlcm5hbCI6IHsKICAgICAgICAgICAgImF1dGhfdGltZSI6IDI5OSwKICAgICAgICAgICAgImF1dGhfdHlwZSI6ICJiYXNpYy1hdXRoIiwKICAgICAgICAgICAgIm9yZ19pZCI6ICIxMTc4OTc3MiIKICAgICAgICB9LAogICAgICAgICJhY2NvdW50X251bWJlciI6ICI2MDg5NzE5IiwKICAgICAgICAidXNlciI6IHsKICAgICAgICAgICAgImZpcnN0X25hbWUiOiAiSW5zaWdodHMiLAogICAgICAgICAgICAiaXNfYWN0aXZlIjogdHJ1ZSwKICAgICAgICAgICAgImlzX2ludGVybmFsIjogZmFsc2UsCiAgICAgICAgICAgICJsYXN0X25hbWUiOiAiUUEiLAogICAgICAgICAgICAibG9jYWxlIjogImVuX1VTIiwKICAgICAgICAgICAgImlzX29yZ19hZG1pbiI6IHRydWUsCiAgICAgICAgICAgICJ1c2VybmFtZSI6ICJpbnNpZ2h0cy1xYSIsCiAgICAgICAgICAgICJlbWFpbCI6ICJqbmVlZGxlK3FhQHJlZGhhdC5jb20iLAogICAgICAgICAgICAidXNlcl9pZCI6ICI2MDg5NzE5IgogICAgICAgIH0sCiAgICAgICAgInR5cGUiOiAiVXNlciIKICAgIH0KfQ==") //nolint:lll

	hasPermissionKessel(c)

	assert.Equal(t, http.StatusOK, w.Code)

	workspaces, found := c.Get(utils.KeyInventoryWorkspaces)
	require.True(t, found)
	workspaceIDs, ok := workspaces.([]string)
	require.True(t, ok)
	require.Greater(t, len(workspaceIDs), 0)
	assert.Equal(t, "inventory-group-1", workspaceIDs[0])
}

func TestHasPermissionKessel_NoWorkspaces(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{Method: http.MethodGet, Header: http.Header{}}

	// Do not set any identity / workspaces so that hasPermissionKessel
	// collects an empty workspace IDs slice and should abort as unauthorized.
	hasPermissionKessel(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	_, found := c.Get(utils.KeyInventoryWorkspaces)
	require.False(t, found, "inventory workspaces should not be set when authorization fails")
}

```

To compile these tests, ensure `manager/middlewares/kessel_test.go` imports `net/http/httptest` (and `net/http` if not already imported):

1. Add `httptest` to the import list:

<<<<<<< SEARCH
import (
	"net/http"

	"github.com/gin-gonic/gin"
=======
import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
>>>>>>> REPLACE

If the actual import block differs, adjust the import edits accordingly to include `net/http/httptest`.
</issue_to_address>

### Comment 3
<location path="listener/common_test.go" line_range="89-92" />
<code_context>
 // assertSystemInventoryProfileMatchesHost checks host-derived system_inventory columns written by
 // storeOrUpdateSysPlatform (must stay in sync on ON CONFLICT DO UPDATE, not only on first insert).
-// nolint: unparam
+// nolint: unparam,funlen
 func assertSystemInventoryProfileMatchesHost(t *testing.T, inventoryID string, host *Host) {
 	t.Helper()
</code_context>
<issue_to_address>
**suggestion:** Test (and make helper robust for) hosts without any workspace groups

The helper currently assumes `host.Groups[0]` exists, but `updateSystemPlatform` now allows hosts with zero groups (nil `WorkspaceID`/`WorkspaceName`). Please guard against `len(host.Groups) == 0` (asserting `inv.WorkspaceID`/`WorkspaceName` are nil in that case), and add a test with `Groups: nil`/`[]` to confirm the values stay nil and no panic occurs.

Suggested implementation:

```golang
 // assertSystemInventoryProfileMatchesHost checks host-derived system_inventory columns written by
 // storeOrUpdateSysPlatform (must stay in sync on ON CONFLICT DO UPDATE, not only on first insert).
 // nolint: unparam,funlen
 func assertSystemInventoryProfileMatchesHost(t *testing.T, inventoryID string, host *Host) {
 	t.Helper()
 	var inv models.SystemInventory

 	assert.JSONEq(t, string(utils.MarshalNilToJSONB(host.Tags)), string(inv.Tags))

+	// A host may legitimately have no workspace groups; in that case the inventory workspace
+	// fields should remain nil and we must not index into host.Groups[0].
+	if len(host.Groups) == 0 {
+		assert.Nil(t, inv.WorkspaceID)
+		// NOTE: if inv.WorkspaceName is a *string / sql.NullString or similar, it should also be nil/invalid here.
+		// See additional_changes below for aligning this with the actual type.
+	} else if hostWorkspaceID := host.Groups[0].ID; hostWorkspaceID != "" {
+		require.NotNil(t, inv.WorkspaceID)
+		assert.Equal(t, hostWorkspaceID, inv.WorkspaceID.String())
+	} else {

```

1. In `assertSystemInventoryProfileMatchesHost`, you should also assert the correct behavior for `inv.WorkspaceName` in the `len(host.Groups) == 0` branch, matching its actual type:
   - If it's a pointer: `assert.Nil(t, inv.WorkspaceName)`
   - If it's an `sql.NullString`: `assert.False(t, inv.WorkspaceName.Valid)`
   - Or equivalent for whatever type is used.
2. Anywhere else in this helper where `host.Groups[0]` is accessed (e.g., to derive workspace name), wrap that logic in a `len(host.Groups) > 0` guard in the same pattern.
3. Add tests in `listener/common_test.go` that exercise the helper with no groups, e.g.:
   - One test with `host := &Host{Groups: nil, /* other required fields */}` that eventually calls `assertSystemInventoryProfileMatchesHost` and verifies no panic and that `inv.WorkspaceID`/`WorkspaceName` remain nil.
   - Another test with `host := &Host{Groups: []*HostGroup{}, /* other required fields */}` doing the same.
   These tests should mirror the existing tests that cover the non-empty-groups case, reusing whatever setup helpers you already have for inserting a `Host`, calling `storeOrUpdateSysPlatform`, and retrieving the resulting `models.SystemInventory`.
</issue_to_address>
~~~

</details>

***

<details>
<summary>Sourcery is free for open source - if you like our reviews please consider sharing them ✨</summary>

- [X](https://twitter.com/intent/tweet?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [Mastodon](https://mastodon.social/share?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https://sourcery.ai)
- [Facebook](https://www.facebook.com/sharer/sharer.php?u=https://sourcery.ai)

</details>

<sub>
Help me be more useful! Please click 👍 or 👎 on each comment and I'll use the feedback to improve your reviews.
</sub>

### Review by @MichaelMraka - Approved on May 11, 2026 at 12:33 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2176*
