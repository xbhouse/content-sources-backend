---
repo: patchman-ui
pr_number: 1224
period: templates-era
source: history/repos/patchman-ui/pull-requests/merged/
---

# Deep Dive: patchman-ui PR #1224

## Problem / Motivation

RHINENG-7807 removed legacy patch template UI and switched systems list references to content-sources template data. This broke stable/production environments expecting the old UI. HMS-5033 required a controlled transition: preview uses new content templates; stable retains old UI until feature-flag flip.

## Solution Approach

@Andrewgdewar wrapped old and new template UIs behind Unleash feature flag `patchman-ui.template-update.enabled`. Preview environments see content template data; stable keeps legacy patch templates until deliberately switched.

## Key Design Decisions

- **Feature flag over big-bang cutover**: Production safety during multi-quarter migration.
- **No UI change in stable by default**: Acceptance criteria explicitly preserve stable behavior.
- **Easy future flip**: Flag enables switching stable to new templates when ready.

## Impact

Critical bridge PR unifying Patchman's template UX with content-sources backend. Enabled parallel development of content templates API (#510) without forcing immediate customer migration. PR #1516 later removed deprecated patch-set code.

## Review Notes

Testing via proxy against prod and Unleash flags on stage/prod. Associated with RHINENG-7807 removal work.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
