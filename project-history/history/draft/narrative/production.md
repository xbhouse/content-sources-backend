# Production Patchman (May 2020 – May 2022)

**1,826 commits · 2 repositories (+ content-sources seeds at period end)**

[← Project History](../PROJECT_HISTORY.md) · [← Genesis](genesis.md) · [Expansion →](expansion.md)

---

## Opening

If genesis was about choosing directions, the production period was about **surviving contact with customers**. Patchman went from demo to deployed Insights service: PostgreSQL upgraded for production workloads, export endpoints shipped, search and metrics added, the schema decomposed for scale, and the UI gained bulk operations and release automation. By May 2022, patchman-engine and patchman-ui were mature platform citizens — and the workspace was about to become multi-product.

This period is patchman-dominated: roughly 58% of commits touched patchman-engine, 39% patchman-ui. Content-sources-backend appeared in April 2022 (@jlsherrill's initial commit) and content-sources-frontend in June — straddling the period boundary. The defining tension was **scale vs. platform alignment**: storing more systems, serving more queries, while migrating identity models and analytics expectations that the platform team controlled.

---

## Major Themes

### Production hardening: databases, exports, and observability

May 2020 opened with @josef-hak upgrading PostgreSQL from 10 to 12 and @Michal Hornicky adding export endpoints, search, and metrics. Export was not a nice-to-have — enterprise customers needed CSV exports of advisories and systems for compliance reporting. @jiridostal wired export into the UI the same week (May 28).

Observability matured in parallel: @Michal Hornicky added metrics for database processes, inter-service call reporting, and warnings for "no rows modified" conditions that indicated stale Kafka replays. @josef-hak hardened the listener against invalid upload messages and added host-OS listener support for developer workflows.

The `/patch` prefix added to the manifest (@josef-hak, May 27) and gzip compression options for VMaaS calls reflected platform integration details that only surface in production — routing conventions, bandwidth optimization on advisory sync.

### Schema normalization for scale

January 2021 brought the most significant pre-2023 schema change: @MichaelMraka split monolithic system storage into `system_platform`, `system_advisories`, and `system_repo` tables. This was not a rename — it was decomposition to enable finer-grained indexing and later Cyndi (inventory sync) integration.

The migration required coordinated deploys: data migration scripts, table switchover, bulk-delete tuning for large accounts. The API surface stayed stable, which was the point — customers saw no breaking change while the engine gained headroom for accounts with hundreds of thousands of systems.

This pattern — **normalize under load, keep APIs stable** — would repeat with system_package2 (2023) and system_platform split for Cyndi (2025–2026).

### Release engineering and UI feature velocity

patchman-ui adopted semantic-release early (@semantic-release-bot releases appear from May 2020). The UI shipped select-all on advisories and affected systems, indeterminate checkbox states, pagination fixes, and moment.js removal (replaced with native Date) — a steady cadence of customer-facing improvements.

@mkholjuraev joined patchman-ui work in this period, beginning the patch-set/template UI features that would evolve for years. @jiridostal remained the primary UI contributor but the team was expanding.

### Patch templates: the first template concept

In mid-2021, patchman-ui began patch-set (later patch template) work — named collections of advisories applied to system groups. This was patchman's native template model, built entirely within patchman-engine's API. The production period planted the seed that the **Templates Era** would eventually replace with content-sources content templates, but in 2020–2022 patch sets were a patchman-ui/patchman-engine concern only.

### Platform identity: org_id migration

The period's capstone was the migration from `account_number` to `org_id` as the canonical tenant identifier — Red Hat Insights' platform-wide identity overhaul. Patchman stored accounts keyed by account number; inventory events and identity headers were transitioning to org-scoped IDs.

This was not a simple column add. It required dual-write logic, backfill jobs, cache safety when Kafka messages still carried only account_number, and explicit avoidance of cross-tenant cache collisions. The migration landed at the period boundary (May–June 2022) and is detailed below in the PR #960 milestone section.

### Floorist analytics: joining the platform reporting ecosystem

February 2022's Floorist (DataHub) integration ([patchman-engine #898](../deep-dives/patchman-engine-pr-898-analysis.md)) was easy to underestimate because it changed no application code — only ClowdApp deployment YAML. But it established how patch compliance data would reach platform analytics: scheduled SQL queries exported to S3 for downstream reporting pipelines.

This deployment-only pattern would be reused for template metrics on both patchman-engine and content-sources-backend years later.

---

## Key Milestones

### System table decomposition (January 2021)

Before PR numbers cleanly capture the work, @MichaelMraka's January 2021 commits split monolithic system storage into normalized tables. The timeline shows concentrated activity: `system_platform`, `system_advisories`, `system_repo` with data migration and bulk-delete tuning.

Why it mattered: the original schema stored too much per-system data in wide rows. Query performance for "show me all systems missing advisory X" degraded as Insights onboarded larger customers. Normalization enabled indexes on the columns queries actually filtered — platform, advisory relationships, repository membership — without changing REST response shapes.

The tradeoff was migration risk. Schema changes on a live multi-tenant service require coordinated deploys, rollback plans, and often dual-write periods. The team chose migration complexity over query complexity — the right call given patchman's read-heavy workload.

@MichaelMraka would author 914 commits lifetime — more than any other human contributor — and many of the hardest schema changes trace to his work in this period and again in expansion (system_package2) and convergence (Cyndi splits). The January 2021 decomposition established the expectation that patchman-engine's senior maintainers own migration design, not just application features.

### PR #898 — Floorist (DataHub) analytics integration (February 2022)

@josef-hak added Floorist job definitions to `deploy/clowdapp.yaml` — scheduled SQL export queries publishing patch compliance aggregates to S3. No application code changed. [Deep dive](../deep-dives/patchman-engine-pr-898-analysis.md) · [PR archive](../repos/patchman-engine/pull-requests/merged/0898-spm-1268-added-floorist-(datahub)-config-to-deploy.md)

The design decision worth noting: **SQL-in-manifest**. Analytics queries live in deployment config, not in Go code. This keeps the evaluation engine focused on compliance logic while reporting queries can be updated independently by platform analytics teams. Floorist became the standard export path; later PRs (#1654 on patchman-engine, #1538 on content-sources-backend) optimized queries as template features grew.

Review was minimal — @MichaelMraka approved, merged within a day. Floorist was already the established platform pattern; patchman was catching up, not innovating.

The long-term impact exceeded the PR's diff size. When content-sources added template Floorist queries in PR #1538 (2026), the pattern was identical: SQL in clowdapp.yaml, scheduled export, S3 upload. Analytics concerns stay out of application code; deployment teams can tune queries without engine redeploys.

### PR #960 — org_id migration (SPM-1482, May 2022)

@psegedy added an `org_id` column to accounts with dual-write logic for new accounts, backfill for existing records, and special handling when Kafka messages still carried only `account_number`. A follow-up org_id populator CJI (SPM-1550) completed the backfill in production. [Deep dive](../deep-dives/patchman-engine-pr-960-analysis.md)

Three design decisions stand out:

1. **Dual identifier period**: Continue creating accounts from Kafka `account_number` while accepting org_id from identity headers — a pragmatic bridge, not a clean cutover.
2. **Cache safety**: Explicitly avoided `findAccount` in listeners to prevent cross-tenant cache collisions when account_number of one user could numerically match org_id of another.
3. **Non-destructive org_id**: Never overwrite org_id when multiple null-account records exist and an empty account_number header arrives.

@MichaelMraka's review was terse: "Please rebase it on top of current master." The technical substance was in @psegedy's PR description documenting cache leak concerns and inline todos.

Impact: this migration enabled all subsequent org-scoped APIs. Content-sources-backend's create endpoint (PR #19, June 2022) reused the same org_id patterns from platform-go-middlewares. Getting identity wrong would have been a data isolation catastrophe; the dual-write period was slower but safer.

### Go-native database tooling (June 2022, period boundary)

@psegedy's SPM-1537 work at the period end — migrate DB using Go code, wait for DB with golang, remove postgresql dependency from scripts — replaced shell-based migration tooling with Go CLIs. This foreshadowed the dedicated migration Job architecture of 2026 (PR #2236) but in 2022 the goal was simpler: remove fragile shell scripts from the deploy path.

### Content-sources seeds (April–June 2022)

@jlsherrill created content-sources-backend in April 2022. By June, @Andrewgdewar had bootstrapped content-sources-frontend, @Ryan Verdile and @Mike Shriver were landing early backend PRs, and @Justin Sherrill was adding repository CRUD. The production period ends as a second product line accelerates — the **Expansion** period picks up this story.

### Kafka library consolidation and patch template seeds

May 2022 also saw [patchman-engine #955](../repos/patchman-engine/pull-requests/merged/0955-spm1495-remove-confluent-kafka-go.md) (@psegedy) remove the confluent-kafka-go experiment (SPM-1495), consolidating on `segmentio/kafka-go`. Two Kafka client libraries had created abstraction overhead without benefit — a small but representative production-period decision to **reduce moving parts** as operational complexity grew.

patchman-ui patch-set work accelerated in 2021–2022 under @mkholjuraev — systems page cleanup, Redux store simplification (removing localStorage-persisted store), patch set assignment UX. The patch template concept that content-sources would later replace was actively used by customers during this period; production hardened both the engine and the template UI that expansion would eventually bridge away from.

### Contributor landscape

Production-period humans centered on patchman:

| Contributor | Production-period signal |
|-------------|-------------------------|
| @josef-hak | Kafka, exports, Postgres upgrades, Floorist, Kafka lib consolidation |
| @MichaelMraka | Schema migrations, system table splits, highest engine commit volume |
| @Michal Hornicky / @mhornick | Metrics, exports, search, listener reliability |
| @jiridostal | UI features, semantic-release cadence, export integration |
| @mkholjuraev | patchman-ui patch-set/template UX (joined during period) |
| @psegedy | org_id migration, Go-native DB tooling (peak activity at period end) |

@jlsherrill's April 2022 content-sources-backend initial commit appears at the boundary — the production narrative ends as the workspace's most prolific PR author enters the story.

---

## Impact

The production period transformed patchman from prototype to platform service. Export APIs, search, metrics, normalized schema, semantic-release, and Floorist analytics are the infrastructure of a **maintained product**, not a demo. The org_id migration aligned patchman with Insights' identity model just as content-sources was born — timing that let the new product skip account_number entirely.

Honest tradeoffs from this era: shell-script migrations (later replaced), monolithic listener process (later split), patch-set templates as a patchman-native concept (later superseded), and dependency churn from dependabot/semantic-release that would only accelerate. The production period's mistakes were mostly **deferrals** — schema normalization happened twice more, identity migration needed a populator CJI, analytics queries needed later optimization — not wrong directions.

For the workspace as a whole, production established patchman as the anchor product whose patterns (Kafka ingestion, PostgreSQL normalization, ClowdApp deployment, org-scoped APIs) content-services would adopt and extend. Everything after May 2022 is a multi-product story; everything before was patchman finding its production shape.

---

[← Project History](../PROJECT_HISTORY.md) · [← Genesis](genesis.md) · [Expansion →](expansion.md)

*Sources: [production timeline](../timeline/production.md) · [deep dives](../deep-dives/patchman-engine-pr-898-analysis.md) (#898, #960) · [architectural milestones](../curated/architectural_milestones.md) (milestones 3–6, 12)*
