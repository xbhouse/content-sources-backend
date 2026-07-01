# Project History: Content Services & Patchman

**Scope**: Nine GitHub repositories · ~9,500 commits · ~6,800 pull requests · October 2019 – June 2026  
**Organizations**: `RedHatInsights` (patchman-engine, patchman-ui, frontend-development-proxy) · `content-services` (content-sources-backend, content-sources-frontend, yummy, zest, tang, caliri)

---

## Table of Contents

- [Executive Summary](#executive-summary)
- [Timeline Overview](#timeline-overview)
- [Period Summaries](#period-summaries)
- [Architectural Journey](#architectural-journey)
- [Community & Contributors](#community--contributors)
- [Lessons Learned](#lessons-learned)
- [Reference Materials](#reference-materials)

**Full period narratives**: [genesis](draft/narrative/genesis.md) · [production](draft/narrative/production.md) · [expansion](draft/narrative/expansion.md) · [templates-era](draft/narrative/templates-era.md) · [convergence](draft/narrative/convergence.md)

---

## Executive Summary

This workspace is a **multi-repository program** within Red Hat Insights that grew from one product into two converging product lines sharing platform infrastructure, identity models, and a unified content template domain.

**Patchman** (2019+) answers: *Which of my systems are missing which security patches and errata?* **patchman-engine** consumes Kafka inventory events, evaluates installed packages against VMaaS advisory data, and stores compliance state in PostgreSQL. **patchman-ui** is the PatternFly/React frontend embedded in the Hybrid Cloud Console.

**Content Sources** (2022+) answers: *What packages exist in my custom YUM/DNF repositories, and how do I compose that content into reproducible templates?* **content-sources-backend** introspects external repositories, snapshots them via Pulp, and exposes a content templates API. Supporting libraries (**yummy** for repo parsing, **zest** for Pulp bindings, **tang** for snapshot queries, **caliri** for Candlepin bindings) and **content-sources-frontend** complete the stack.

The program exists because Insights customers need both patch compliance visibility and custom content management — and because Red Hat's platform architecture (Kafka event bus, org-scoped identity, ClowdApp deployment on OpenShift, Floorist analytics) rewards services that integrate rather than stand alone.

Evolution followed five distinct phases:

1. **Genesis** (2019–2020): Patchman prototype; Go selected over Python/Rust through benchmarks; Kafka and PostgreSQL chosen from day one.
2. **Production** (2020–2022): Patchman hardened for customers — exports, metrics, schema normalization, Floorist analytics, org_id identity migration.
3. **Expansion** (2022–2023): Content Sources founded; event-driven introspection; Pulp ecosystem (zest, snapshots, tang); patchman scale migrations.
4. **Templates Era** (2024–2025): Content templates API replaced patch sets; UI modernization (RTL, feature-flagged migration); snapshot maturity.
5. **Convergence** (2025–present): Cross-product Kafka events, Playwright E2E across three repos, Konflux CI, deployment architecture fixes.

Activity grew from ~116 commits in 2019 to ~2,050 in 2025. The workspace expanded from 2 repos and ~5 contributors to 9 repos with sustained quarterly peaks above 650 commits. The current phase is integration under load — wiring products together while customers use all of them in production.

Five cross-cutting themes span all periods: **Kafka event architecture** (inventory → introspection → template/advisory events), **PostgreSQL schema evolution** (repeated normalization under stable APIs), **Pulp content storage** (snapshots and generated clients), **template domain unification** (patch sets → content templates), and **Playwright/Konflux modernization** (testing and CI convergence in 2025–2026).

---

## Timeline Overview

PR numbers below are **repo-scoped** — the same number in different repos refers to different changes.


| Period            | Date Range          | Commits | Repos         | Key Theme                      | Major Milestones                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ----------------- | ------------------- | ------- | ------------- | ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Genesis**       | Oct 2019 – Apr 2020 | 688     | 2             | Prototype & language selection | Go prototype wins; Kafka listener; PostgreSQL schema; patchman-ui bootstrap                                                                                                                                                                                                                                                                                                                                                                                       |
| **Production**    | May 2020 – May 2022 | 1,826   | 2 (+CS seeds) | Patchman production hardening  | Postgres 12 upgrade; system table splits (2021); [patchman-engine #898](draft/deep-dives/patchman-engine-pr-898-analysis.md) Floorist; [patchman-engine #960](draft/deep-dives/patchman-engine-pr-960-analysis.md) org_id                                                                                                                                                                                                                                         |
| **Expansion**     | Jun 2022 – Dec 2023 | 2,284   | 6             | Content Sources founding       | [CS-backend #4](draft/deep-dives/content-sources-backend-pr-4-analysis.md), [#19](draft/deep-dives/content-sources-backend-pr-19-analysis.md); [#55](draft/deep-dives/content-sources-backend-pr-55-analysis.md), [#113](draft/deep-dives/content-sources-backend-pr-113-analysis.md) introspection; [#458](draft/deep-dives/content-sources-backend-pr-458-analysis.md) snapshots; [patchman-engine #1353](draft/deep-dives/patchman-engine-pr-1353-analysis.md) |
| **Templates Era** | Jan 2024 – Jun 2025 | 2,062   | 8             | Content templates domain       | [CS-backend #510](draft/deep-dives/content-sources-backend-pr-510-analysis.md) templates CRUD; [#486](draft/deep-dives/content-sources-backend-pr-486-analysis.md) snapshot-for-date; [patchman-ui #1164](draft/deep-dives/patchman-ui-pr-1164-analysis.md) RTL; [#1224](draft/deep-dives/patchman-ui-pr-1224-analysis.md) UI bridge                                                                                                                              |
| **Convergence**   | Jul 2025 – present  | 2,600+  | 9             | Cross-product integration      | [CS-backend #956](draft/deep-dives/content-sources-backend-pr-956-analysis.md), [patchman-ui #1445](draft/deep-dives/patchman-ui-pr-1445-analysis.md) Playwright; [#1537](draft/deep-dives/content-sources-backend-pr-1537-analysis.md) template events; [patchman-engine #2194](draft/deep-dives/patchman-engine-pr-2194-analysis.md), [#2236](draft/deep-dives/patchman-engine-pr-2236-analysis.md)                                                             |


---

## Period Summaries

### Genesis (October 2019 – April 2020)

@jiridostal created patchman-engine on October 31, 2019; patchman-ui followed on November 4. @josef-hak, @semtexzv (Michal Hornicky), and @MichaelMraka built Go, Python, and Rust prototypes in parallel — Go won through Apache Bench benchmarks and Kafka throughput tests. Kafka message handling, PostgreSQL schema (adapted from vuln-engine), OpenAPI contracts, and OpenShift deployment templates were established before the first customer existed.

The UI team (@jiridostal) shipped PatternFly tables, Redux state, pagination, remediation integration, and API connectivity in the same sprint cadence as the engine. February 2020 split the evaluator into upload and recalc paths — the first process separation beyond a monolithic listener.

**Full narrative**: [genesis.md](draft/narrative/genesis.md)

---

### Production Patchman (May 2020 – May 2022)

Patchman transitioned from prototype to production Insights service. @josef-hak upgraded PostgreSQL 10→12; @Michal Hornicky added export endpoints, search, and Prometheus metrics; @jiridostal wired exports and bulk selection into the UI. @MichaelMraka's January 2021 system table decomposition (`system_platform`, `system_advisories`, `system_repo`) was the largest schema change until 2023.

@josef-hak integrated Floorist/DataHub analytics ([patchman-engine #898](draft/deep-dives/patchman-engine-pr-898-analysis.md), February 2022). @psegedy's org_id migration ([#960](draft/deep-dives/patchman-engine-pr-960-analysis.md), May 2022) aligned multi-tenancy with platform identity. content-sources-backend appeared in April 2022 as the period closed.

**Full narrative**: [production.md](draft/narrative/production.md)

---

### Expansion (June 2022 – December 2023)

Commit activity jumped sharply as Content Sources went live. @jlsherrill founded the backend; @Andrewgdewar bootstrapped the frontend; @rverdile established api/dao/handler patterns ([content-sources-backend #4](draft/deep-dives/content-sources-backend-pr-4-analysis.md), [#19](draft/deep-dives/content-sources-backend-pr-19-analysis.md)). @jlsherrill's introspection CLI ([#55](draft/deep-dives/content-sources-backend-pr-55-analysis.md)) and @avisiedo's Kafka-driven introspection ([#113](draft/deep-dives/content-sources-backend-pr-113-analysis.md)) — 13,106 packages discovered in first ephemeral test — became the product core.

@rverdile extracted **yummy** ([#142](draft/deep-dives/content-sources-backend-pr-142-analysis.md)). **zest** automated Pulp OpenAPI bindings (March 2023). @Andrewgdewar's snapshot endpoint ([#458](draft/deep-dives/content-sources-backend-pr-458-analysis.md)) introduced Pulp-backed content storage. @psegedy's system_package2 migration ([patchman-engine #1353](draft/deep-dives/patchman-engine-pr-1353-analysis.md)) addressed patchman scale.

**Full narrative**: [expansion.md](draft/narrative/expansion.md)

---

### Templates Era (January 2024 – June 2025)

Content-services named its central domain: **content templates** — named collections of repository snapshots with arch/version constraints. @rverdile's templates CRUD API ([content-sources-backend #510](draft/deep-dives/content-sources-backend-pr-510-analysis.md), January 2024) replaced patchman's legacy patch sets. @Andrewgdewar's feature-flag bridge ([patchman-ui #1224](draft/deep-dives/patchman-ui-pr-1224-analysis.md)) preserved stable environments during multi-quarter UI migration.

@mkholjuraev migrated patchman-ui tests from Enzyme to React Testing Library ([#1164](draft/deep-dives/patchman-ui-pr-1164-analysis.md)), unblocking PatternFly 5. **caliri** automated Candlepin bindings. @xbhouse drove introspection API work (groups/environments from comps metadata) and cross-repo integration.

**Full narrative**: [templates-era.md](draft/narrative/templates-era.md)

---

### Convergence (July 2025 – Present)

All nine repositories active at sustained all-time commit highs. @Andrewgdewar introduced Playwright E2E to content-sources-backend ([#956](draft/deep-dives/content-sources-backend-pr-956-analysis.md)); @dominikvagner extended it to patchman-ui ([#1445](draft/deep-dives/patchman-ui-pr-1445-analysis.md)). @dominikvagner's [frontend-development-proxy #1](draft/deep-dives/frontend-development-proxy-pr-1-analysis.md) established shared Konflux CI pipelines (merged July 28, 2025).

@swadeley's template update events job ([#1537](draft/deep-dives/content-sources-backend-pr-1537-analysis.md)) and @katarinazaprazna's advisory update topic ([patchman-engine #2194](draft/deep-dives/patchman-engine-pr-2194-analysis.md)) connected products via Kafka. @jlsherrill added template Floorist analytics ([#1538](draft/deep-dives/content-sources-backend-pr-1538-analysis.md)). @TenSt moved database migrations to dedicated ClowdApp Jobs ([#2236](draft/deep-dives/patchman-engine-pr-2236-analysis.md)). @Dugowitch led the Cyndi system_platform split ([#1988](draft/deep-dives/patchman-engine-pr-1988-analysis.md)).

This period remains open; template-advisory relationships (patchman-engine #2249) continue in flight.

**Full narrative**: [convergence.md](draft/narrative/convergence.md)

---

## Architectural Journey

### Core data flows

The workspace's architecture is best understood as three overlapping pipelines that share platform infrastructure but serve different domains.

**Patch compliance pipeline** (patchman-engine, unchanged since genesis):

```
Inventory (Kafka) → Listener → PostgreSQL → Evaluator → VMaaS advisories
                                      ↓
                              REST API ← patchman-ui
                                      ↓
                              Floorist → S3 analytics
```

**Content introspection pipeline** (content-sources-backend, expansion onward):

```
Repository CRUD (REST) → Kafka introspect → yummy parsing → PostgreSQL
                                                    ↓
                                            Pulp snapshot (zest clients)
                                                    ↓
                                            tang (RPM/errata queries)
```

**Template domain** (templates era → convergence):

```
Content templates API ← repository snapshots + constraints
         ↓                          ↓
  patchman-ui (feature flags)   Candlepin (caliri bindings)
         ↓
  Kafka: platform.content-sources.template (CS-backend #1537)
         ↓
  patchman-engine: template refs + advisory events (#2194)
```

### Major technical decisions


| Decision                                        | When      | Rationale                                    | Tradeoff                                            |
| ----------------------------------------------- | --------- | -------------------------------------------- | --------------------------------------------------- |
| **Go over Python/Rust**                         | Nov 2019  | Benchmark-proven Kafka + Postgres throughput | One month running three prototypes                  |
| **Kafka ingestion**                             | Nov 2019  | Decouple from inventory release cadence      | Operational complexity (consumers, poison messages) |
| **OpenAPI-first APIs**                          | Dec 2019  | Parallel UI/engine development               | Spec maintenance burden                             |
| **api/dao/handler layering**                    | Jun 2022  | Consistent backend structure                 | Boilerplate per endpoint                            |
| **Micro-libraries (yummy, zest, tang, caliri)** | Dec 2022+ | Independent testing and versioning           | Service/module proliferation                        |
| **Pulp for snapshots**                          | Dec 2023  | Production content storage                   | Infrastructure dependency                           |
| **Feature-flagged UI migration**                | Nov 2024  | Avoid customer-breaking cutover              | Dual code paths until flag flip                     |
| **Playwright E2E**                              | Feb 2025  | Catch cross-service integration failures     | Slower, flakier CI                                  |
| **Migration as ClowdApp Job**                   | Jun 2026  | Prevent concurrent migrators on scale-out    | Deploy flow complexity                              |


### Chronological breaking changes


| Date         | Change                                                                   | Impact                                                       |
| ------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------ |
| **May 2020** | PostgreSQL 10 → 12                                                       | Database upgrade; coordinated deploy                         |
| **Jan 2021** | system_* table decomposition                                             | Schema migration; API unchanged                              |
| **May 2022** | org_id replaces account_number (patchman-engine #960)                    | Identity header semantics; dual-identifier transition window |
| **May 2022** | Kafka client consolidation (patchman-engine #955)                        | Internal only; single library                                |
| **Oct 2022** | Kafka introspection topic (CS-backend #113)                              | New Kafka contract; operators configure connectivity         |
| **Dec 2023** | system_package → system_package2 (patchman-engine #1353)                 | Long-running background migration; dual-write period         |
| **Jan 2024** | Content templates API (CS-backend #510)                                  | New API surface; patch sets deprecated over time             |
| **Nov 2024** | Patch template UI behind feature flag (patchman-ui #1224)                | Stable/prod behavior differs from preview                    |
| **Feb 2026** | Cyndi system_platform split (patchman-engine #1988)                      | Inventory sync behavior changes for edge cases               |
| **Jun 2026** | DB migration Job replaces init container (patchman-engine #2236)         | Deployment flow: pods block until Job completes              |
| **Jun 2026** | Template/advisory Kafka topics (CS-backend #1537, patchman-engine #2194) | New event contracts for downstream consumers                 |


Most schema migrations kept REST APIs stable. Breaking changes concentrated in identity (org_id), deployment flow (migration Job), and the patch-set → content-template domain transition (managed via feature flags rather than hard cutover).

### Repository roles (current)


| Repository                   | Role                                                                               |
| ---------------------------- | ---------------------------------------------------------------------------------- |
| `patchman-engine`            | Patch compliance evaluation; Kafka inventory ingestion; template reference storage |
| `patchman-ui`                | Hybrid Cloud Console UI for patch management and template assignment               |
| `content-sources-backend`    | Repository CRUD, introspection, snapshots, content templates API                   |
| `content-sources-frontend`   | Admin UI for repositories, snapshots, templates, admin tasks                       |
| `yummy`                      | YUM/DNF repository parsing library                                                 |
| `zest`                       | Pulp OpenAPI → Go bindings generator                                               |
| `tang`                       | Snapshot RPM/errata query microservice                                             |
| `caliri`                     | Candlepin OpenAPI → Go bindings generator                                          |
| `frontend-development-proxy` | Local HCC dev proxy; shared Konflux CI pipeline configs                            |


---

## Community & Contributors

### Growth arc


| Era       | Repos | Active humans | Character                                          |
| --------- | ----- | ------------- | -------------------------------------------------- |
| 2019      | 2     | ~5            | Founding team; every commit is feature work        |
| 2020–2021 | 2     | ~8            | Production hardening; semantic-release bots appear |
| 2022      | 4     | ~15           | Content Sources team joins; commit rate doubles    |
| 2023      | 6     | ~20           | Pulp ecosystem; zest/tang/yummy                    |
| 2024–2025 | 8     | ~25           | Templates domain; cross-repo reviewers             |
| 2025–2026 | 9     | ~25 + bots    | Konflux/dependabot bots; highest commit volume     |


Bot accounts (dependabot, red-hat-konflux, semantic-release) account for ~1,800 PRs in 2024–2026. Human review culture scaled with the program — review comments, not just authorship, measure engagement.

### Key contributors

**Patchman founders and maintainers**

- **@MichaelMraka** (Michael Mraka) — 914 commits, 336 PRs, 609 approvals. patchman-engine's highest-volume contributor; schema migrations, evaluation logic, Cyndi adaptation.
- **@josef-hak** (Josef Hak) — 763 commits, 487 PRs. Go prototype, Kafka, Floorist (#898), production hardening.
- **@psegedy** (Patrik Segedy) — 557 commits, 343 PRs. org_id (#960), system_package2 (#1353), Kafka client consolidation (#955).
- **@jiridostal** (Jiri Dostal) — 351 commits. patchman-ui founder; UI architecture through production.
- **@mhornick / @semtexzv** (Michal Hornicky) — ~349 commits. Genesis schema, Kafka listener, benchmarks.
- **@mkholjuraev** (Mukhammadkhol Kholiujurakhmad) — 305 commits, 363 PRs. patchman-ui templates, RTL migration (#1164).
- **@TenSt** (Stepan Makymchuk) — 133 commits, 411 review comments. Migration Job architecture (#2236).
- **@Dugowitch** (Jakub Dugovič) — 140 commits. Cyndi system_platform split (#1988).

**Content Sources leaders**

- **@jlsherrill** (Justin Sherrill) — ~632 commits, 619 PRs (most authored). content-sources-backend founder; introspection (#55), template Floorist (#1538).
- **@rverdile** (Ryan Verdile) — 259 commits, 260 PRs, 1,430 review comments (most reviews). api/dao/handler patterns, templates API (#510), yummy (#142).
- **@xbhouse** (Bryttanie House) — 604 commits, 205 PRs, 838 reviews. Full-stack content-sources; yummy maintainer, snapshots and introspection APIs.
- **@Andrewgdewar** (Andrew Dewar) — 602 commits, 218 PRs. Frontend bootstrap, snapshots (#458, #486), Playwright (#956), UI bridge (#1224).
- **@swadeley** (Stephen Wadeley) — 268 commits, 270 PRs, 542 approvals. Template events (#1537), Playwright CI.
- **@dominikvagner** (Dominik Vagner) — 158 commits, 411 reviews. frontend-development-proxy, Konflux (#1), patchman-ui Playwright (#1445).

Verified username mappings: [contributor_mapping.json](draft/curated/contributor_mapping.json) · Extended profiles: [key_contributors.md](draft/curated/key_contributors.md)

---

## Lessons Learned

### What worked well

**Benchmark-driven language selection.** Go, Python, and Rust prototypes with Apache Bench in November 2019 prevented a costly rewrite.

**Contract-first development.** OpenAPI specs before handlers (patchman, 2019) and Kafka schema codegen (content-sources, 2022) enabled parallel development across thousands of PRs.

**Normalize schema under stable APIs.** system table splits (2021), system_package2 (2023), and Cyndi adaptations (2026) kept REST contracts stable while restructuring PostgreSQL internals.

**Feature flags for domain replacement.** [patchman-ui #1224](draft/deep-dives/patchman-ui-pr-1224-analysis.md) avoided breaking stable environments during the patch-set → content-template migration.

**Extract libraries when reuse is real.** yummy, zest, tang, and caliri each justified independent versioning after the problem was understood, not before.

**Event-driven cross-product integration.** Template update events (#1537) and advisory update events (#2194) replaced fragile REST polling between patchman and content-sources.

### What would be done differently

**Init-container database migrations on scaled services.** The dedicated Job fix (#2236) arrived in 2026 — six years after genesis. Stateful migration jobs should be day-one architecture for horizontally scaled ClowdApp services.

**Single Kafka client library from the start.** The confluent-kafka-go experiment (removed in patchman-engine #955, 2022) added abstraction overhead without benefit.

**Test stack migrations earlier.** Enzyme blocked PatternFly 5 for years; RTL migration (#1164, 2024) should have started when Enzyme's deprecation was announced.

**Bot PR volume management.** Konflux and dependabot PRs dominate 2025–2026 statistics. Auto-merge policies and batch review rituals help preserve human signal.

### Insights for similar projects

Multi-repo programs converge on **shared events, shared tests, and shared CI** — or drown in integration overhead. The hardest transitions are **domain ownership shifts** — when patch sets became content templates, feature flags and parallel APIs mattered more than code quality alone.

Document migrations and dual-write periods in PR descriptions — org_id (#960) and system_package2 (#1353) succeeded because authors documented cache safety rules and per-account batching inline.

Start cross-repo testing before coverage feels complete. Playwright adoption (#956, #1445) began with two example specs; CI gates and fixtures came first, coverage expanded incrementally.

---

## Reference Materials


| Resource                              | Location                                                                               |
| ------------------------------------- | -------------------------------------------------------------------------------------- |
| Period narratives (full)              | [draft/narrative/](draft/narrative/)                                                   |
| PR deep dives (22 analyses)           | [draft/deep-dives/](draft/deep-dives/)                                                 |
| Architectural milestones (30 curated) | [draft/curated/architectural_milestones.md](draft/curated/architectural_milestones.md) |
| Key contributors                      | [draft/curated/key_contributors.md](draft/curated/key_contributors.md)                 |
| Contributor username mapping          | [draft/curated/contributor_mapping.json](draft/curated/contributor_mapping.json)       |
| Period structure analysis             | [draft/curated/period_structure.md](draft/curated/period_structure.md)                 |
| Validation report                     | [draft/validation_report.md](draft/validation_report.md)                               |
| Timeline data                         | [draft/timeline/](draft/timeline/)                                                     |
| PR archives                           | [repos/](repos/)                                                                       |
| Generation guide                      | [tools/HISTORY_GENERATION_GUIDE.md](../tools/HISTORY_GENERATION_GUIDE.md)              |


---

*Phase 4–5 · Prompts 7–9 · Validated against PR archives and git history (see [validation_report.md](draft/validation_report.md))*

---

#### Generated with https://github.com/ochosi/project-history