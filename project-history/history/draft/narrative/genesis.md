# Genesis (October 2019 – April 2020)

**688 commits · 2 repositories · patchman-engine, patchman-ui**

[← Project History](../PROJECT_HISTORY.md) · Next: [Production →](production.md)

---

## Opening

The genesis period is the story of a greenfield service built under time pressure: Red Hat Insights needed patch compliance — which systems are missing which errata — and the team had roughly six months to go from nothing to a credible prototype. What emerged was not a single codebase but a **paired architecture**: a Go backend consuming Kafka inventory events and evaluating packages against VMaaS advisories, and a React/Redux frontend embedded in the Insights platform shell.

This period was defined by **exploration over polish**. Three language prototypes (Python, Go, Rust) competed in the same repository. Kafka and PostgreSQL were chosen from day one, not added later. The UI was built in parallel with the engine, not as an afterthought. By April 2020 the language question was settled, the core data model existed, and the team was ready to harden for production — but almost nothing in this period would survive unchanged. That is the point: genesis established **direction**, not **final form**.

---

## Major Themes

### Language selection: Go wins through measurement, not mandate

@josef-hak added the Go prototype on November 1, 2019, the same week @semtexzv (Michal Hornicky) scaffolded Python from vuln-engine and a basic Rust experiment. The team did not pick Go on faith. @josef-hak built benchmark infrastructure — Apache Bench scripts, parallel NEVRA parsing tests, optional batch DB writes — and documented results in the README. @semtexzv's November benchmarks compared goroutine-per-message processing against Python's single-threaded path.

Go won because it handled **Kafka throughput + PostgreSQL batch writes** without the operational complexity Rust would have added to a small team, and without Python's GIL limitations at message volume. The Rust prototype never progressed beyond basics; Python served as a reference implementation borrowed from vuln-engine. By December, @semtexzv had created the base application structure from the Go prototype, and Python/Rust directories faded from active development.

The tradeoff was real: the team spent November running three stacks instead of one. In hindsight that month of parallel prototypes prevented a costly wrong-language migration later — patchman-engine would eventually reach 900+ commits from @MichaelMraka alone, all in Go.

### Kafka-from-day-one: event-driven before it was fashionable

On November 4, @semtexzv got Kafka and the database running together. @josef-hak added Kafka message sending after platform container startup. By December 2, @semtexzv implemented proper Kafka message handling — not a polling cron, not a REST webhook, but a **listener** processing inventory upload events as they arrived.

This decision locked in patchman-engine's architecture for the next seven years. Every subsequent scale challenge — org_id migration, Cyndi inventory sync, advisory update events — extended the Kafka layer rather than replacing it. The genesis team accepted operational complexity (Kafka consumer groups, message ordering, poison messages) in exchange for decoupling from inventory's release cadence.

Local development got the same attention as production: docker-compose with SELinux workarounds, Python scripts to send test messages over a single producer, inventory mocks. @josef-hak documented selinux instructions in the README — a small detail that signals how much friction early OpenShift development created.

### PostgreSQL schema borrowed from experience

@semtexzv added the database schema in November; @josef-hak copied schema from the vulnerability engine on November 28. The team was not inventing patch data modeling from scratch — they adapted patterns from Insights' existing vulnerability service. Tables were renamed from errata to advisory in December. @MichaelMraka joined in December with schema updates and removal of business_risk concepts that did not fit patch compliance.

@semtexzv added batch insert with `ON CONFLICT DO UPDATE` in December — upsert semantics for idempotent Kafka replays, essential once duplicate messages became a production concern. @josef-hak switched tests from SQLite to Postgres in December, rejecting the shortcut of in-memory testing that would miss PostgreSQL-specific behavior.

### Paired UI: PatternFly + Redux from the first week

@jiridostal created patchman-ui on November 4, four days after patchman-engine's initial commit. The UI was never a separate project — it shared the same sprint cadence, the same demo deadlines, the same Insights platform conventions.

November brought webpack build setup, landing page, Header component, encrypted deployment keys, and the first Redux store. December accelerated: expandable advisory tables, checkboxes, pagination with limit/offset, sorting, loading states, systems page, system detail routes, inventory integration. @jiridostal connected the UI to real API data by November 29 and switched from a temporary Vulnerability API to the SPM API by December 12.

The PatternFly + `@redhat-cloud-services/frontend-components` stack chosen here persisted through 2026. Dependabot was already active in genesis — @dependabot-preview bumped PatternFly from 3.16 to 3.124 in November, establishing the dependency-churn rhythm that would dominate later periods.

### OpenAPI-first API design

@josef-hak added API mockups and Swagger in November. @semtexzv implemented API client generation in December. @josef-hak converted Swagger 2 to OpenAPI 3.0.0 on December 2. The engine's external contract was defined before most handlers were complete — handlers like `systems`, `advisories`, `advisory_detail`, and `system_advisories` were implemented by @josef-hak in mid-December against an already-published spec.

This contract-first approach enabled the UI and engine teams to work in parallel: @jiridostal built tables and pagination against mock responses while @josef-hak wired handlers to PostgreSQL. The pattern would reappear in content-sources-backend three years later with the same api/dao/handler layering.

### Platform integration scaffolding

Genesis work was not localhost-only. @semtexzv added OpenShift deployment templates and ocdeployer integration in November. @semtexzv implemented 3scale identity parsing in December — multi-tenancy via identity headers was a platform requirement from the start, not a later retrofit. @josef-hak added VMaaS client integration and Prometheus middleware configuration in December.

The vmaas_sync component (@semtexzv, December 3) established the advisory sync pipeline: patchman does not own CVE/advisory data, it consumes VMaaS. That boundary — inventory events in, VMaaS advisories as reference data, compliance state in PostgreSQL — is still the core data flow.

---

## Key Milestones

### Initial commit and Docker scaffolding (October 31, 2019)

@jiridostal's initial commit on October 31 was the anchor. Within hours @semtexzv created base Dockerfiles and a Python scaffold based on vuln-engine. The repo structure — `deploy/`, `dev/`, application packages, database migrations — was established in the first week.

This matters because patchman-engine's directory layout from October 2019 is recognizable in the 2026 codebase. The team did not reorganize; they extended. Docker-compose, RHEL Dockerfiles, and OpenShift templates created the deployment path that ClowdApp configs would inherit years later.

### Go prototype and benchmark structure (November 1–15, 2019)

@josef-hak's Go prototype and benchmark structure (November 6) was the inflection point. NEVRA parsing with regex, parallel parsing, batch DB writes, Apache Bench integration — these were not features, they were **evidence** for the language decision.

@semtexzv's November 15 benchmark work ("make go save items immediately") and goroutine-per-message spawning (November 18) validated concurrency assumptions. The README's go prototype guide (@josef-hak, November 6) documented how new contributors could reproduce results — important for a team that would grow across time zones.

### Kafka listener and VMaaS sync (December 2019)

December converted the prototype into a system. @semtexzv's "Implement proper handling of Kafka messages" (December 2) and "Sync advisories from VMaaS" (@semtexzv, January 3 — just after period boundary but rooted in December work) connected the two external dependencies.

@josef-hak separated evaluation from upload processing (January 3) — a design split that would evolve into distinct listener and evaluator processes in production. The genesis listener was monolithic; the **conceptual** separation of ingest vs. evaluate was present early.

### First end-to-end UI demo (December 2019)

@jiridostal's December sprint — systems page, advisory tables with checkboxes, pagination, sorting, system detail routes — produced the first demo-ready UI. "Upload code for first API demo" (December 2) and "Use SPM API" (December 12) mark the transition from mock data to real backend integration.

The UI team's test investment in genesis (hooks tests, helpers tests, component tests) set expectations that patchman-ui would maintain high test coverage — though the testing technology would change twice (Enzyme → RTL → Playwright) over the following six years.

### Evaluator split and remediation integration (February–April 2020)

February 2020 was the busiest genesis month (180 commits) and marked the transition from "prototype" to "pre-production feature set." @josef-hak split the evaluator into `evaluator_upload` and `evaluator_recalc` (February 3) — the first concrete process separation beyond the conceptual ingest/evaluate distinction. The `EVAL_LABEL` env var allowed running labeled evaluator instances in OpenShift, foreshadowing the multi-process ClowdApp deployment model.

@josef-hak added re-evaluation message sending, vmaas_sync duration metrics, and unified `entrypoint.sh` scripts for OpenShift. @Michal Hornicky simplified system_advisories counting, refresh functions, and system update triggers while adding periodic culling of removed inventory systems — operational concerns that only matter when real customers accumulate stale data.

On the UI side, @jiridostal integrated Remediations — bulk remediation on affected systems, kebab dropdown actions, custom remediation connectors across system advisory and systems pages. Patchman was not just displaying compliance state; it was wiring into Insights' remediation workflow. Search query params appeared in OpenAPI docs (@josef-hak, February 3). Table footers, last-seen dates replacing status columns, and selection on affected systems polished the UX toward production expectations.

March and April continued hardening: more handler tests, docker-compose improvements for debugging (volume mounts), Dockerfile validation in CI, and advisory sync reliability. By April 30 the commit rate was steady — not the exploratory spikes of November, but the sustained output of a team shipping toward a production date.

### Who built genesis

The contributor map for this period is small and identifiable:

| Contributor | Role in genesis |
|-------------|-----------------|
| @jiridostal | patchman-ui founder; UI architecture, Redux, API integration |
| @josef-hak | Go prototype, benchmarks, handlers, VMaaS/Kafka, evaluator split |
| @semtexzv / @mhornick | Schema, Kafka listener, DB batch ops, OpenShift deploy, benchmarks |
| @MichaelMraka | Schema refinement, handler implementation (joined December 2019) |

No content-sources contributors existed. No bot-driven dependency PRs at the volume of 2024–2026. Every commit was human-authored feature or infrastructure work — a contrast that makes genesis timelines readable in a way convergence timelines are not.

---

## Impact

Genesis shaped the workspace in ways that no later refactor reversed. **Go, Kafka, PostgreSQL, OpenAPI, and PatternFly** were chosen under uncertainty and confirmed under load. The paired engine/UI repo split established a development pattern that content-sources would replicate in 2022 with backend + frontend repos.

The period's honest tradeoffs — three language prototypes, SQLite rejected for tests, monolithic listener with conceptual evaluation split, schema borrowed from vuln-engine — were pragmatic choices for a deadline-driven greenfield. What genesis did not solve (production metrics, export APIs, schema normalization, org-scoped identity, analytics export) became the entire agenda of the **Production** period starting May 2020.

For developers reading this history: genesis is the reminder that architectural decisions made in the first 60 days, especially **data flow boundaries** (Kafka in, VMaaS reference, PostgreSQL state) and **contract-first APIs**, constrain everything that follows. The content-services program would later inherit these same instincts — event-driven ingestion, org-scoped APIs, shared Go libraries — because the people and patterns originated here.

---

[← Project History](../PROJECT_HISTORY.md) · Next: [Production →](production.md)

*Sources: [genesis timeline](../timeline/genesis.md) · [architectural milestones](../curated/architectural_milestones.md) (milestones 1–2) · [key contributors](../curated/key_contributors.md)*
