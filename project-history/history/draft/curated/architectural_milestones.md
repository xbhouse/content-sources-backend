# Architectural Milestones

Curated from `important_prs.json`, `theme_analysis.json`, and timeline data across nine repositories in the content-services / Red Hat Insights upstream workspace. Automated importance scoring over-weighted dependency bumps; this list was hand-curated for architectural significance.

**Scope**: ~9,500 commits, ~6,800 PRs, Oct 2019 – Jun 2026  
**Periods**: genesis · production · expansion · templates-era · convergence (see [period_structure.md](period_structure.md))

---

## Patchman Engine & UI — Foundations (2019–2022)

### 1. patchman-engine genesis — Initial commit and Go prototype
**Reference**: Commits Oct–Nov 2019 (`patchman-engine`)  
**Why significant**: Establishes the vulnerability/patch evaluation service that anchors half the workspace's history. The first commits created Docker scaffolding, a Python prototype based on vuln-engine, then a Go prototype with Kafka message handling, PostgreSQL schema, and OpenShift deployment templates.  
**What changed**: Core data model (NEVRA parsing, advisory metadata), Kafka-driven ingestion, benchmark infrastructure for language selection (Go won over Python/Rust prototypes).  
**Breaking changes**: N/A — greenfield.

### 2. patchman-ui genesis — React/Redux frontend bootstrap
**Reference**: Commits Nov 2019 (`patchman-ui`)  
**Why significant**: Paired UI for patch management on the Red Hat Insights platform; established the PatternFly + Redux architecture still in use today.  
**What changed**: Landing page, Header component, encrypted deployment keys, Redux store configuration, webpack build pipeline.  
**Breaking changes**: N/A — greenfield.

### 3. PR #898 — Floorist (DataHub) analytics integration
**Repo**: `RedHatInsights/patchman-engine` · Feb 2022  
**Why significant**: Introduced Floorist/DataHub SQL export jobs in `deploy/clowdapp.yaml`, enabling platform-wide analytics on patch compliance data. This pattern later expanded to template metrics and cross-service reporting.  
**What changed**: Scheduled SQL queries exported to S3 for downstream analytics pipelines.  
**Breaking changes**: None; additive deployment config.

### 4. PR #960 — org_id migration (SPM-1482)
**Repo**: `RedHatInsights/patchman-engine` · May 2022  
**Why significant**: Migrated multi-tenancy from `account_number` to `org_id`, aligning Patch with the platform-wide identity model. Critical for data isolation as Insights moved to organization-scoped access.  
**What changed**: Added `org_id` column, dual-write during transition, org_id populator CJI, cache handling when Kafka messages still carried only `account_number`.  
**Breaking changes**: Identity header semantics changed; downstream consumers had to handle both identifiers during migration window.

### 5. PR #955 — Remove confluent-kafka-go experiment
**Repo**: `RedHatInsights/patchman-engine` · May 2022 · @psegedy  
**Why significant**: Consolidated on `segmentio/kafka-go` after evaluating Confluent's client (SPM-1495). Simplified the messaging layer to a single Kafka library.  
**What changed**: Removed unused confluent-kafka-go dependency and related abstraction code.  
**Breaking changes**: None for operators; internal library consolidation.

### 6. System table migrations (Jan 2021)
**Reference**: Commits by @Michael Mraka — `system_platform`, `system_advisories`, `system_repo` table splits (`patchman-engine`)  
**Why significant**: Decomposed monolithic system storage into normalized tables, enabling finer-grained indexing and later Cyndi (inventory sync) integration.  
**What changed**: New tables with data migration scripts, table switchover, bulk-delete tuning for large accounts.  
**Breaking changes**: Schema migration required coordinated deploy; API surface unchanged.

---

## Patchman Engine — Scale & Data Model (2022–2026)

### 7. PR #1353 — system_package per-account migration (RHINENG-5394)
**Repo**: `RedHatInsights/patchman-engine` · Dec 2023  
**Why significant**: Part of a multi-PR effort to replace the JSON-heavy `system_package` table with normalized `system_package2`, improving query performance at scale for accounts with hundreds of thousands of systems.  
**What changed**: Per-account migration job, POC → production path through PRs #1316–#1356, eventual drop of legacy table.  
**Breaking changes**: Long-running background migration; dual-write period.

### 8. PR #1654 — Grouped query for host template Floorist
**Repo**: `RedHatInsights/patchman-engine` · May 2025  
**Why significant**: Optimized Floorist analytics for template assignment metrics, reducing join complexity in compliance reporting.  
**What changed**: Rewrote SQL aggregation to pre-group system counts before joining to templates.  
**Breaking changes**: None; analytics query shape change only.

### 9. PR #1988 — Split system_platform for Cyndi (RHINENG-21214)
**Repo**: `RedHatInsights/patchman-engine` · Feb 2026  
**Why significant**: Adapted the core inventory sync table to Cyndi pipeline changes, splitting `system_platform` to match upstream inventory event structure.  
**What changed**: Table split migration, trigger simplification (PR #1817), updated listener logic.  
**Breaking changes**: Database migration; inventory sync behavior changed for edge cases.

### 10. PR #2236 — DB migration as dedicated ClowdApp Job (RHINENG-27056)
**Repo**: `RedHatInsights/patchman-engine` · Jun 2026  
**Why significant**: Moved schema migrations out of the manager pod's init container into a single Job, preventing concurrent migrators during rolling deploys when multiple replicas competed for `pg_advisory_lock`.  
**What changed**: Manager init replaced with `check-for-db`; new `db-migration` Job; entrypoint retry logic; runbook documentation (PR #2244).  
**Breaking changes**: Deployment flow change — new pods block until migration Job completes.

### 11. PR #2194 — patchman.advisory.update Kafka topic (RHINENG-26118)
**Repo**: `RedHatInsights/patchman-engine` · 2025  
**Why significant**: New event stream for advisory lifecycle, enabling downstream services (including content-sources) to react to advisory changes.  
**What changed**: New Kafka topic and producer; template-advisory relationship storage (PR #2249, open).  
**Breaking changes**: None; additive event contract.

---

## Patchman UI — Templates, Testing & Platform UI (2020–2026)

### 12. Patch Templates → Content Templates transition (2022–2024)
**Reference**: PR #832 area (rename patch set → patch template, Jun 2022); PR #1224 (Nov 2024)  
**Why significant**: Renamed and rebuilt the template workflow from legacy "patch sets" to "patch templates," then bridged to the new content-sources template API. PR #1224 re-exposed old UI via feature flag while preview environments switched to content template data (RHINENG-7807).  
**What changed**: Wizard internationalization, feature flags, systems list template column, eventual removal of patch-set code (PR #1516).  
**Breaking changes**: API backend for patch sets removed; UI controlled by Unleash feature flags during transition.

### 13. PR #1092 / #1114 — Inventory groups column and filter (ESSNTL-3716)
**Repo**: `RedHatInsights/patchman-ui` · 2023  
**Why significant**: Integrated Host Inventory group data into the patch systems view, a major UX feature tying patch compliance to inventory organization structure.  
**What changed**: Custom filter and sortable column for inventory groups; paginated API requests (PR #1163).  
**Breaking changes**: None; additive UI.

### 14. PR #1164 — Enzyme → React Testing Library migration (RHINENG-7975)
**Repo**: `RedHatInsights/patchman-ui` · Feb 2024  
**Why significant**: Modernized the entire test stack from deprecated Enzyme to RTL, unblocking PatternFly 5 and React 18 upgrades. Part of a multi-PR effort (#1132–#1158).  
**What changed**: Rewrote component tests, removed unnecessary Redux mocks, commented out fed-modules inventory tests pending PF5 migration.  
**Breaking changes**: Test infrastructure only; no production API changes.

### 15. PR #1445 — Playwright test initialization (HMS-9438)
**Repo**: `RedHatInsights/patchman-ui` · Dec 2025  
**Why significant**: Introduced end-to-end Playwright testing with CI pipeline, fixtures, and example specs — aligning patchman-ui with the content-sources testing strategy.  
**What changed**: Playwright config, CI workflow, helper fixtures, 2 initial tests.  
**Breaking changes**: None; additive test infrastructure.

---

## Content Sources — Backend Architecture (2022–2026)

### 16. PR #4 — Repository configuration database and model (CONTENT-39)
**Repo**: `content-services/content-sources-backend` · May 2022  
**Why significant**: First persistent data layer — PostgreSQL migrations, GORM models, transaction-wrapped migration templates, and force-rollback on failure. Foundation for everything that followed.  
**What changed**: `Repository` model, migration tooling (`dbmigrate`), global DB connection pattern (later refined).  
**Breaking changes**: N/A — initial schema.

### 17. PR #19 — Create content source endpoint (CONTENT-40)
**Repo**: `content-services/content-sources-backend` · Jun 2022  
**Why significant**: Established the api/dao/handler layering pattern and platform identity integration (`x-rh-identity` header via platform-go-middlewares).  
**What changed**: Create endpoint with OpenAPI comments, org-scoped unique URL constraint, dao layer separation.  
**Breaking changes**: N/A — new API.

### 18. PR #55 — Repository introspection command (HMSCONTENT-49)
**Repo**: `content-services/content-sources-backend` · Jul 2022  
**Why significant**: Core capability — downloading and parsing external YUM/DNF repositories to discover packages. This is the heart of the content-sources value proposition.  
**What changed**: CLI introspection command, SSL cert configuration, public repo support, revision tracking.  
**Breaking changes**: None; new capability.

### 19. PR #113 — Kafka-driven introspection (HMSCONTENT-151)
**Repo**: `content-services/content-sources-backend` · Oct 2022  
**Why significant**: Event-driven architecture — repositories introspect on Kafka messages instead of polling, enabling platform integration and async processing at scale.  
**What changed**: `kafka-introspect` command, schema-generated Go types, kafkacat test tooling, ephemeral environment validation across kafka pool types.  
**Breaking changes**: New Kafka topic contract; operators must configure Kafka connectivity.

### 20. PR #142 — Yummy library extraction (HMSCONTENT-200)
**Repo**: `content-services/content-sources-backend` · Dec 2022  
**Why significant**: Extracted repository parsing from `dao/external_resources.go` into standalone `yummy` library, establishing the pattern of shared Go modules across content-services repos.  
**What changed**: Moved introspection/parsing logic to `github.com/content-services/yummy`; backend imports library via go module.  
**Breaking changes**: Internal refactor; external API unchanged.

### 21. PR #458 — Repository snapshot endpoint (HMS-1973)
**Repo**: `content-services/content-sources-backend` · Dec 2023  
**Why significant**: Introduced Pulp-backed snapshotting — capturing point-in-time repository state for template composition and content delivery.  
**What changed**: `POST /repositories/{uuid}/snapshot/` endpoint, Pulp integration path, frontend testing hooks.  
**Breaking changes**: None; additive API. Required Pulp infrastructure.

### 22. PR #510 — Content templates CRUD (HMS-2965)
**Repo**: `content-services/content-sources-backend` · Jan 2024  
**Why significant**: Defined the content template domain model — named collections of repository snapshots with arch/version constraints. This is the API that replaced patchman's legacy patch templates.  
**What changed**: Create/list/fetch endpoints, filtering by name/version/arch, search.  
**Breaking changes**: New API surface; patch UI migration depended on this.

### 23. PR #486 — Snapshot-for-date API for templates (HMS-2969)
**Repo**: `content-services/content-sources-backend` · Dec 2023  
**Why significant**: Enabled templates to resolve "which snapshot existed on date X" — critical for reproducible content views and template publishing workflows.  
**What changed**: `POST /repositories/snapshots/for_date/` with repository UUID list and date parameter.  
**Breaking changes**: None; additive API.

### 24. PR #956 — Playwright E2E testing (HMS-5287)
**Repo**: `content-services/content-sources-backend` · Feb 2025  
**Why significant**: Added Playwright integration tests and GitHub Actions CI pipeline to the backend — the start of cross-repo Playwright adoption (944 theme-keyword commits in maturity period).  
**What changed**: `_playwright-tests` directory, CI workflow, no production code changes.  
**Breaking changes**: None; CI-only.

### 25. PR #1537 — SendTemplateUpdateEvents job (HMS-9952)
**Repo**: `content-services/content-sources-backend` · Jun 2026  
**Why significant**: Kafka events when template content changes (new advisories), enabling patchman and other consumers to react to template updates asynchronously.  
**What changed**: Background job publishing to `platform.content-sources.template` topic; message structure updates (PR #1533).  
**Breaking changes**: New Kafka event contract for downstream consumers.

### 26. PR #1538 — Template Floorist query (HMS-10345)
**Repo**: `content-services/content-sources-backend` · Jun 2026  
**Why significant**: Analytics export for content templates, mirroring patchman's Floorist pattern and enabling platform reporting on template adoption.  
**What changed**: Floorist SQL query with extended release information in clowdapp deployment config.  
**Breaking changes**: None; additive analytics.

---

## Shared Libraries & Tooling (2023–2026)

### 27. zest — Pulp OpenAPI bindings generator
**Repo**: `content-services/zest` · Mar 2023 (init) – ongoing  
**Why significant**: Automated Go client generation from Pulp's OpenAPI spec, eliminating hand-maintained HTTP clients. All Pulp interactions in content-sources-backend flow through zest-generated bindings.  
**What changed**: OpenAPI generator templates, Konflux CI pipelines (PR #27), automated binding releases consumed by backend via versioned modules.  
**Breaking changes**: Pulp API upgrades require zest regeneration and coordinated backend bumps.

### 28. tang — Snapshot errata query service (PR #1)
**Repo**: `content-services/tang` · Jan 2024  
**Why significant**: Dedicated microservice for querying RPM/errata data from Pulp snapshot databases — offloading complex snapshot queries from the monolithic backend.  
**What changed**: `Tangy` interface with `RpmRepositoryVersionPackageSearch()`, Pulp podman-compose for dev, integration tests.  
**Breaking changes**: N/A — new service; backend integrates via HTTP client.

### 29. caliri — Candlepin bindings automation
**Repo**: `content-services/caliri` · 2024–2026  
**Why significant**: Mirror of zest for Candlepin — automated Go bindings generation for subscription/content template integration with Candlepin APIs.  
**What changed**: Cron-triggered binding builds (PR #3), Konflux workflows (PR #6), release-tracked binding updates (v4.7.5 → v5.0.0).  
**Breaking changes**: Candlepin API upgrades require caliri regeneration.

### 30. frontend-development-proxy PR #1 — Konflux pipeline adoption
**Repo**: `RedHatInsights/frontend-development-proxy` · Jul 2025  
**Why significant**: Established shared Konflux (Red Hat's Tekton-based CI) pipeline configs reused across frontend repos, representing the platform-wide shift from legacy GitHub Actions/Clowder builds.  
**What changed**: Pull-request and push pipeline definitions for Konflux builds.  
**Breaking changes**: Build infrastructure migration; developer workflow changes.

---

## Cross-Cutting Themes

| Theme | Primary repos | Key inflection |
|-------|--------------|----------------|
| **Kafka event architecture** | patchman-engine, content-sources-backend | Inventory events → org_id (2022); introspection events (2022); template update events (2026) |
| **PostgreSQL schema evolution** | patchman-engine, content-sources-backend | Normalization migrations (2021, 2023); dedicated migration jobs (2026) |
| **Pulp content storage** | content-sources-backend, zest, tang | Snapshot endpoint (2023); templates (2024); tang query service (2024) |
| **Template domain unification** | content-sources-*, patchman-ui, patchman-engine | Content templates API (2024) replacing patch sets; UI feature-flag bridge (2024) |
| **Playwright E2E testing** | content-sources-backend, content-sources-frontend, patchman-ui | Backend POC (2025) → patchman-ui init (2025) → 944+ commits tagged in maturity |
| **OpenAPI binding generators** | zest (Pulp), caliri (Candlepin) | Automated client generation replacing hand-written HTTP code |

---

## Items Needing Maintainer Review

- **Exact PR for system_package table partitioning** (Sep 2020): significant commit in timeline but may not have a single merged PR — verify against `patchman-engine` migration files.
- **Pulp integration "first landing"**: multiple PRs (#438, #436, etc.) in late 2023 — may want a single deep-dive on the Pulp adoption sequence.
- **content-sources-frontend architectural milestones**: lighter coverage here; frontend-specific milestones (ClowdApp deployment PR #13, module federation) deserve a dedicated pass in Prompt 3 deep-dives.
- **Correlation coverage**: only 20% of commits matched to PRs — some milestones above are timeline/commit-derived rather than PR-correlated.

---

*Generated: Phase 2, Prompt 1 · Sources: `history/draft/data/important_prs.json`, `theme_analysis.json`, `history/draft/timeline/*.md`, PR markdown archives in `history/repos/`*
