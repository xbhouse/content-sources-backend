# Templates Era (January 2024 – June 2025)

**2,062 commits · 8 repositories**

[← Project History](../PROJECT_HISTORY.md) · [← Expansion](expansion.md) · [Convergence →](convergence.md)

---

## Opening

The templates era is the period when content-services **named its central domain** and patchman began **giving up its own**. Content templates — named collections of repository snapshots with arch, version, and repository constraints — shipped as a first-class API ([content-sources-backend #510](../deep-dives/content-sources-backend-pr-510-analysis.md)). Patchman-ui bridged legacy patch templates to the new backend via feature flags ([#1224](../deep-dives/patchman-ui-pr-1224-analysis.md)). Snapshot infrastructure matured with for-date queries, Pulp content guards, and tang RPM search. Meanwhile patchman-ui modernized its test stack (Enzyme → React Testing Library) to unblock PatternFly 5 and React 18 upgrades that had been deferred for years.

Eight repositories were active; patchman-engine still contributed 38% of commits but content-sources-backend (26%) and content-sources-frontend (16%) together exceeded patchman's UI share. @rverdile, @xbhouse, @swadeley, and @Andrewgdewar drove template and snapshot work; @mkholjuraev and @leSamo carried patchman-ui modernization. The period ends July 2025 as Konflux CI adoption and Playwright E2E scaling mark the shift to **Convergence**.

---

## Major Themes

### Content templates domain: CRUD and everything after

PR #510 (@rverdile, January 2024) defined the content template domain — create, list, fetch endpoints with filtering by name, version, and architecture, plus search. Org-scoped via identity headers. Month-long review (December 2023 → January 2024); @swadeley requested rebase; @Andrewgdewar validated integration.

@Andrewgdewar's approval quote captured the milestone: "Tested, integrated, ACK!" This was not a backend-only API — the frontend was exercising it, patchman-ui would consume it, and Candlepin integration (via caliri bindings, February 2024) would connect templates to subscription management.

The CRUD foundation deliberately shipped before update/delete (#535) and advanced features — enterprise-ready list API from day one with filter, sort, and search. Templates belong to organizations; they reference repository snapshots; they carry arch/version constraints. That model replaced patchman's legacy patch sets entirely.

patchman-engine added its own content templates table (@MichaelMraka, RHINENG-7684, January 2024) — patchman storing template references locally while content-sources remained the source of truth for template content.

### Snapshot maturity: for-date, guards, cleanup

PR #486 (snapshot-for-date API, late 2023 / templates-era activity) enabled "which snapshot existed on date X" queries — critical for reproducible content views and template publishing workflows. PR #502 (Pulp content guards, January 2024) and central Pulp server migration (@jlsherrill, Refs 3372) hardened production snapshot storage.

Snapshot cleanup automation, orphan repo version detection (#520), arm64 repository support (#538), and manual snapshot UI (#186, @kwarnerredhat) filled out the operational surface. Snapshots were no longer a demo feature — they were production infrastructure with edge cases (orphans, guards, multi-arch) receiving dedicated fixes.

tang's RPM search API (PR #1 area, January 2024) complemented backend snapshot endpoints — complex errata queries offloaded to a dedicated service.

### Patch UI transition: feature flags over big-bang

The most politically complex work was patchman-ui's template migration. Patch sets had existed since 2021; customers used them. Content templates in content-sources-backend were the successor, but switching UI and API atomically was unacceptable risk.

PR #1224 (@Andrewgdewar, November 2024) wrapped old and new template UIs behind Unleash feature flag `patchman-ui.template-update.enabled`. Preview environments see content template data; stable keeps legacy patch templates until deliberate flip. RHINENG-7807 had removed legacy UI in a way that broke stable environments; this PR restored controlled transition.

Design decisions:
- **Feature flag over big-bang cutover** — production safety during multi-quarter migration
- **No UI change in stable by default** — acceptance criteria explicitly preserve stable behavior
- **Easy future flip** — flag enables switching stable when ready

PR #1516 (later, convergence period) removed deprecated patch-set code once the bridge proved stable.

### UI modernization: Enzyme → RTL → PF5

PR #1164 (@mkholjuraev, February 2024) migrated patchman-ui's entire test stack from deprecated Enzyme to React Testing Library — part of multi-PR effort #1132–#1158. RHINENG-7975 blocked PatternFly 5 and React 18 upgrades until tests migrated.

Pragmatic scope cut: inventory table tests were **commented out** temporarily rather than delaying the entire migration — fed-modules testing deferred to unblock PF5. 96% patch coverage on changed files per Codecov.

This was test infrastructure only — no production API changes — but it unblocked two years of frontend dependency upgrades. PatternFly 5 migration (#1372) and later Playwright E2E (#1445) depended on RTL being in place.

### caliri and Candlepin: subscription integration

caliri (2024–2026) mirrored zest for Candlepin — automated Go bindings for subscription/content template integration. Cron-triggered binding builds, Konflux workflows, release-tracked updates (v4.7.5 → v5.0.0). Candlepin API upgrades require caliri regeneration and coordinated backend bumps — the same operational contract as zest + Pulp.

### Patchman-engine: templates table and package evaluation

Patchman-engine work in this period was supporting infrastructure: content templates table (RHINENG-7684), package evaluation refactors (RHINENG-6806), installable/applicable package counts in API (RHINENG-6831).

### Pulp production migration and content guards

January 2024 opened with significant Pulp infrastructure work — not feature headlines, but production prerequisites. @jlsherrill's Refs 3372 series switched to a new central Pulp server, new S3 bucket names, and dropped stale stage snapshots. PR #502 added Pulp content guards. PR #520 detected orphan repo versions.

These changes reflect a maturing dependency: Pulp was no longer "the snapshot prototype backend" but production content storage with guard policies, multi-environment bucket strategy, and cleanup automation. Getting Pulp operations wrong meant customer content loss — the templates era invested heavily in operational safety before expanding template CRUD surface area.

### patchman-engine template references

While content-sources owned template **content**, patchman-engine needed template **references** for compliance evaluation — which systems are assigned to which templates, what advisories apply. @MichaelMraka's RHINENG-7684 (January 2024) added a content templates table in patchman-engine. RHINENG-6831 added installable/applicable package counts to the API — evaluation metrics customers used for reporting.

This split — content-sources is source of truth for template definition; patchman-engine evaluates compliance against assigned templates — is the architectural boundary the convergence period's Kafka events (#1537, #2194) would eventually keep synchronized.

### TypeScript and PatternFly 6 groundwork

patchman-ui's RTL migration (#1164) unblocked PatternFly 5 (#1372 area) and Node version upgrades. content-sources-frontend adopted TypeScript patterns and snapshot test utilities (@xbhouse, @dominikvagner). The templates era's frontend work was largely **dependency and test infrastructure** preparing for convergence's PF6 and Playwright scale-up — less visible than PR #510 but equally load-bearing.

---

### PR #510 — Content templates CRUD (HMS-2965)

@rverdile's January 2024 merge was the API event the entire period built toward. [Deep dive](../deep-dives/content-sources-backend-pr-510-analysis.md)

The month-long review reflected API design scrutiny — template models are hard to change once customers depend on them. Filter, sort, and search on list endpoints from v1 avoided "add query params later" breaking changes.

Downstream: patchman-ui #1224 bridged UI; caliri connected Candlepin; patchman-engine stored template references; convergence-period Kafka events (#1537) would notify consumers of template changes.

### PR #486 — Snapshot-for-date API (HMS-2969)

Templates need to answer: "what content existed in this repository on date X?" The for-date endpoint accepted repository UUID lists and a date parameter, returning the appropriate snapshot references. Reproducible content views — essential for compliance scenarios where "what did we publish on March 1?" is an audit question.

Additive API, no breaking changes. Paired with PR #458's snapshot creation to complete the snapshot lifecycle: create → query by date → reference in template.

### PR #1164 — Enzyme → React Testing Library (RHINENG-7975)

@mkholjuraev rewrote Enzyme tests to RTL across patchman-ui. [Deep dive](../deep-dives/patchman-ui-pr-1164-analysis.md)

The fed-modules inventory test deferral was controversial but correct — blocking PF5 migration for fed-modules test parity would have delayed the entire UI stack. Technical debt acknowledged explicitly; Playwright E2E would partially cover integration scenarios later.

Tradeoff: temporary test coverage gap in inventory integration vs. two-year dependency upgrade unblock. The team chose velocity with documented debt.

### PR #1224 — Patch templates → content templates UI bridge (RHINENG-7807 / HMS-5033)

@Andrewgdewar's feature flag bridge was the customer-facing capstone of the template domain migration. [Deep dive](../deep-dives/patchman-ui-pr-1224-analysis.md)

Testing via proxy against prod and Unleash flags on stage/prod. Preview environments validated content template data paths while stable customers saw no change — the multi-quarter migration strategy in one PR.

Impact: enabled parallel development of content templates API (#510) without forcing immediate customer migration. The template domain was **defined** in backend APIs and **transitioned** in patchman-ui through flags — not through a single breaking release.

PR #1224's testing strategy — proxy against prod, Unleash flags on stage/prod — reflected mature platform feature-flag infrastructure. The team did not fork code paths per environment; they parameterized behavior via centralized flag service. When stable flips to content templates, it is a flag change plus validation, not a deploy of new UI code.

---

## Impact

The templates era answered the question expansion left open: **what is content-services for?** Not just introspection and snapshots — but **templates** that compose snapshots into named, constrained, org-scoped content views that replace patchman's legacy patch sets and integrate with Candlepin subscriptions.

Patchman did not shrink — it adapted, storing template references and preparing for advisory event integration. But the center of gravity shifted: content-sources-backend owned template content; patchman-engine owned compliance evaluation against that content; patchman-ui bridged UX during transition.

Honest tradeoffs: feature-flagged UI migration (complexity for safety), commented-out inventory tests (debt for velocity), eight repos coordinating on template contracts (communication overhead), Pulp/Candlepin dependencies (operational surface area). None were mistakes — they were the cost of replacing a live feature (patch sets) with a better architecture (content templates) without customer downtime.

July 2025 opens **Convergence**: Playwright at scale, Konflux CI, template Kafka events, cross-product advisory integration. The template domain was built; the program's job became wiring it together.

---

[← Project History](../PROJECT_HISTORY.md) · [← Expansion](expansion.md) · [Convergence →](convergence.md)

*Sources: [templates-era timeline](../timeline/templates-era.md) · [deep dives](../deep-dives/) (#510, #1164, #1224) · [architectural milestones](../curated/architectural_milestones.md) (milestones 12, 14, 22–23, 29)*
