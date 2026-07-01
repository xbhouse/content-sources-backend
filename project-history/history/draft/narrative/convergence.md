# Convergence (July 2025 – Present)

**2,600 commits · 9 repositories · ongoing**

[← Project History](../PROJECT_HISTORY.md) · [← Templates Era](templates-era.md)

---

## Opening

Convergence is the period when the workspace stopped building separate products and started **operating as one program**. All nine repositories are active. Commit activity hit all-time highs — 588 commits in Q3 2025, 702 in Q4 2025 — driven not by greenfield features but by integration work, testing modernization, and CI migration. Template update events flow from content-sources-backend to patchman-engine via Kafka. Playwright E2E tests gate PRs in three repos. Konflux pipelines replace legacy GitHub Actions builds. Database migrations moved out of rolling-deploy init containers into dedicated Jobs.

The name reflects the work: cross-product wiring, shared tooling, and platform modernization happening simultaneously. @swadeley, @xbhouse, @TenSt, @dominikvagner, and @Dugowitch appear across repo boundaries in review threads that would have been impossible when patchman and content-sources were separate teams in separate mental models.

This period remains open. The narrative below covers July 2025 through June 2026 — the most recent archived data — and describes trends still in motion.

---

## Major Themes

### Playwright E2E: testing paradigm shift

Backend PR #956 (@Andrewgdewar, February 2025 — activity accelerating through convergence) introduced `_playwright-tests` and GitHub Actions CI with **no production code changes**. @jlsherrill questioned directory naming ("why use _playwright-tests instead of playwright-tests or ./tests/playwright/ ?"); @Andrewgdewar chose consistency with frontend convention.

patchman-ui PR #1445 (@dominikvagner, December 2025) added Playwright configuration, CI on every PR, fixtures, and two example specs — the third repo in the adoption chain. content-sources-frontend already had Playwright CI; convergence made E2E a **program standard**, not a single-repo experiment.

Impact by the numbers: 147 Playwright-related commits in 2025 alone; theme analysis tags 944+ playwright-keyword commits in the maturity/convergence window. HMS-5430 introspect tests, template filter tests, RBAC migration tests (@swadeley, HMS-8765), and bulk import repository tests followed.

The tradeoff: E2E tests are slower, flakier, and more expensive to maintain than unit tests. The team accepted that cost because unit tests (even RTL at 96% patch coverage) missed integration failures across services — template assignment, snapshot lifecycle, RBAC roles — that only full-stack tests catch.

### Konflux CI: build infrastructure modernization

frontend-development-proxy PR #1 (@dominikvagner, merged July 28, 2025) established shared Konflux (Red Hat's Tekton-based CI) pipeline configs for pull-request and push pipelines. [Deep dive](../deep-dives/frontend-development-proxy-pr-1-analysis.md)

Konflux bot PRs (`@red-hat-konflux[bot]`) became a dominant contributor in 2025–2026 stats — dependency digest updates, reference bumps, security compliance pipelines. patchman-ui, patchman-engine, and content-sources-frontend all show konflux reference update PRs in the convergence timeline.

Developer workflow changed: legacy Clowder/GitHub Actions builds gave way to Tekton pipelines with `/ok-to-test` gating and bot-driven dependency management. The human review burden shifted from "does this build?" to "does this change make sense?" — bots handle mechanical updates.

Honest assessment: bot PR volume inflates commit statistics and creates review noise. Teams batch-merge konflux updates or auto-merge with policy. The modernization is real; the signal-to-noise ratio in PR lists is not.

### Template events: cross-product Kafka contracts

PR #1537 (@swadeley, June 2026) added `SendTemplateUpdateEvents` — a background job publishing to `platform.content-sources.template` when template content changes (new advisories in repos). Coordinated with @xbhouse's advisory ID endpoints (#1536) and message structure changes (#1533).

@rverdile's review confusion was productive: "I'm a little confused about what the ticket is asking for. Is this understanding correct?" — the team clarified backfill vs. ongoing event semantics. Job vs. inline publish: batch job for initial/backfill population; ongoing updates via separate paths.

PR #1538 (@jlsherrill, June 2026) added template Floorist analytics — SQL export mirroring patchman's Floorist pattern ([patchman-engine #898](../deep-dives/patchman-engine-pr-898-analysis.md), 2022) for platform reporting on template adoption. [Deep dive](../deep-dives/content-sources-backend-pr-1538-analysis.md)

patchman-engine PR #2194 (@katarinazaprazna) added `patchman.advisory.update` Kafka topic — advisory lifecycle events for downstream consumers. Typed `AdvisoryUpdateEvent` contract; topics provisioned externally by Platform team rather than in ClowdApp config. Connects to template-advisory relationship work (PR #2249) and content-sources template events.

This is convergence in architecture: **events, not REST polling**, connect products that were built independently. Template changes in content-sources propagate to patchman; advisory changes in patchman propagate to content-sources.

### Cyndi and inventory: system_platform split

patchman-engine PR #1988 (@Dugowitch, February 2026) adapted core inventory sync to Cyndi pipeline changes — splitting `system_platform` to match upstream inventory event structure. [Deep dive](../deep-dives/patchman-engine-pr-1988-analysis.md)

Inventory sync is patchman's lifeline — without current system data, compliance evaluation is wrong. Cyndi changes upstream force schema changes downstream. The convergence period's patchman-engine work is disproportionately **adaptation** (Cyndi, org_id tail work, template references) rather than new features.

@MichaelMraka's RHINENG-18242 work (July 2025) removed candlepin calls from the listener — decoupling patchman from synchronous subscription lookups as template data moves to content-sources events.

### Deployment reliability: DB migration as Job

PR #2236 (@TenSt, June 2026) addressed a production incident class: the manager pod ran database migrations in an init container, and rolling deploys with multiple replicas caused concurrent migration attempts competing for `pg_advisory_lock(123)`. ClowdApp cannot set maxSurge on the manager (public web service), making concurrent migrators inevitable.

Solution: dedicated ClowdApp **Job** (single runner), manager init replaced with lightweight `check-for-db`, retry logic in `entrypoint.sh`. New pods fail init if migration fails; old pods keep serving. @dominikvagner: "looks good to me! ✅"

Paired with runbook documentation (PR #2244). This fix applies to all future schema migrations — a deployment architecture change, not a one-off patch.

The mistake being corrected: running migrations in init containers on horizontally scaled web services. Many teams make this choice early for simplicity; patchman hit the wall at scale during large migrations in 2025–2026.

### frontend-development-proxy: local dev and shared CI

Beyond Konflux PR #1, the proxy repo provides local development proxy for Hybrid Cloud Console frontend work (@dominikvagner founder). Convergence treats developer experience as shared infrastructure — not each repo maintaining its own proxy configuration.

### RBAC, integration tests, and QE migration

Convergence invested heavily in quality engineering across repos. @swadeley's HMS-8765 added Content Sources RBAC RHEL roles testing. HMS-5999 migrated IQE tests to Playwright (@swadeley). HMS-8635 added bulk import repository tests. @dominikvagner fixed nightly CI env vars and documented integration test execution (@swadeley, "How to run integration tests").

patchman-ui moved to Node 22 (@Michael Johnson, RHINENG-16645, July 2025) and absorbed Konflux reference updates with occasional reverts when bot PRs broke compliance pipelines — the friction of CI migration visible in the timeline.

content-sources-frontend became the highest commit-volume repo in convergence (21% of commits) — not because backend features slowed, but because Playwright CI, Konflux digest updates, and integration test expansion concentrated in frontend workflows.

### Floorist template analytics and advisory relationships

PR #1538 (HMS-10345) added template Floorist SQL queries with extended release information — completing the analytics story Floorist started in patchman-engine PR #898 (2022). Platform reporting on template adoption required the same DataHub export pattern applied to content-sources deployment config.

PR #2249 (open at archive time) stores template-advisory relationships in patchman-engine — the data model connecting content-sources template content to patchman compliance evaluation. PR #2194's advisory update events feed this relationship reactively rather than through polling.

Together: template update events flow out of content-sources (#1537); advisory update events flow out of patchman (#2194); patchman stores relationships (#2249); Floorist exports adoption metrics (#1538). The event graph is the convergence architecture.

### Activity and sustainability

Quarterly commits: Q3 2025 (588), Q4 2025 (702), Q1 2026 (674) — sustained peak, not a spike. Bot contributors (dependabot, red-hat-konflux) account for a significant PR fraction, but human review culture scaled: @rverdile (1,430 review comments), @jlsherrill (1,376), @swadeley (1,174), @xbhouse (838 reviews, 346 approvals).

The program's bus factor improved versus genesis's four-person core — but coordination cost across nine repos replaced it. Convergence's challenge is organizational as much as technical.

---

## Key Milestones

### PR #956 — Playwright E2E testing (HMS-5287)

@Andrewgdewar's February 2025 merge was pure test infrastructure — `_playwright-tests` directory, GitHub Actions workflow, zero production code. [Deep dive](../deep-dives/content-sources-backend-pr-956-analysis.md)

Directory naming debate (@jlsherrill vs. @Andrewgdewar) was the most commented design decision in a PR with no production impact. Consistency with frontend won.

Enables: HMS-5430 introspect test, template filter tests, RBAC role tests, bulk import tests — the convergence period's QA strategy.

### PR #1445 — Playwright test initialization (HMS-9438)

@dominikvagner brought Playwright to patchman-ui — config, CI on every PR, fixtures, two example specs. [Deep dive](../deep-dives/patchman-ui-pr-1445-analysis.md)

Fixtures-first design: reusable helpers for auth, navigation, common flows. Example specs only — team expands coverage incrementally rather than blocking merge on full suite.

Aligns patchman-ui with content-sources testing strategy. Unit tests (RTL) catch component bugs; Playwright catches routing, API integration, and feature-flag behavior.

### PR #1537 — SendTemplateUpdateEvents job (HMS-9952)

@swadeley's template update events job is the cross-product capstone. [Deep dive](../deep-dives/content-sources-backend-pr-1537-analysis.md)

Coordination complexity: advisory ID endpoints (#1536), message structure (#1533), backfill vs. ongoing semantics clarified in review with @rverdile. Job for batch backfill; separate paths for ongoing updates.

Without this: patchman would poll content-sources for template changes or miss updates entirely. Event-driven integration completes the template domain story started in PR #510 (2024).

### PR #2194 — patchman.advisory.update Kafka topic (RHINENG-26118)

@katarinazaprazna opened advisory lifecycle to event-driven integration — the reverse direction from #1537. [Deep dive](../deep-dives/patchman-engine-pr-2194-analysis.md)

Connects to template-advisory relationships (PR #2249) — patchman tracking which advisories belong to which templates based on content-sources data, reacting to advisory changes via events.

### PR #2236 — DB migration as dedicated ClowdApp Job (RHINENG-27056)

@TenSt's migration Job fix is deployment architecture, not feature work — but it unblocks every future schema migration across the program. [Deep dive](../deep-dives/patchman-engine-pr-2236-analysis.md)

Sourcery bot suggested making `MIGRATION_MAX_RETRIES` configurable via ClowdApp parameter — small polish on a fix addressing incidents during rolling deploys.

Every developer maintaining a ClowdApp-deployed service with PostgreSQL migrations should read this PR's problem statement: init-container migrations + horizontal scaling + advisory locks = race conditions that are hard to debug in production.

### frontend-development-proxy PR #1 — Konflux pipeline adoption

@dominikvagner's same-day merge established shared Konflux pipeline configs — PR and push pipelines reused across frontend repos. Minimal discussion (`/ok-to-test` only); maximum downstream impact as konflux bot PRs proliferate across the workspace.

Marks the shift from legacy GitHub Actions/Clowder builds to Tekton-based CI — developer workflow change as significant as any feature PR in this period.

---

## Impact

Convergence is the busiest era on record because the workspace is **integrating under load** — not greenfield, not maintenance, but wiring nine repos together while customers use all of them in production. Playwright, Konflux, Kafka template/advisory events, Cyndi adaptation, and migration Job architecture are the enabling infrastructure for a program that outgrew the "two patchman repos" origin story.

What worked: event-driven cross-product integration (cheaper than synchronous REST coupling), shared test stacks (Playwright across repos), extracting deployment fixes (migration Job) that benefit all services, feature-flagged migrations (from templates era, still paying off).

What remains hard: bot PR noise, E2E test flakiness at scale, coordinating Kafka schema changes across teams, Cyndi upstream changes forcing downstream migrations on patchman's schedule.

The period is open. Template-advisory relationships (PR #2249), Floorist template analytics (#1538), and Konflux rollout continue. The workspace's arc — from patchman prototype (2019) to nine-repo integrated program (2026) — enters its next chapter with infrastructure mature enough that feature work can focus on customer value rather than foundational plumbing.

For developers: convergence proves that multi-repo programs converge on **shared events, shared tests, and shared CI** — or drown in integration overhead. The investments in Playwright, Konflux, and Kafka contracts are the program's immune system against the complexity of its own growth.

---

[← Project History](../PROJECT_HISTORY.md) · [← Templates Era](templates-era.md)

*Sources: [convergence timeline](../timeline/convergence.md) · [deep dives](../deep-dives/) (#956, #1445, #1537, #1538, #1988, #2194, #2236, proxy #1) · [architectural milestones](../curated/architectural_milestones.md) · [validation report](../validation_report.md)*
