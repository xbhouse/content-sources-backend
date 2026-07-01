# Evolution Patterns & Proposed Period Structure

Analysis of `history/draft/timeline/*.md`, commit activity across 9 repositories (9,474 commits, Oct 2019 – Jun 2026), `theme_analysis.json`, and curated milestones/contributors.

---

## Executive Summary

This workspace is not a single product — it is a **multi-repo program** spanning two major product lines that converged over time:

1. **Patchman** (2019+) — patch compliance for Red Hat Insights
2. **Content Sources** (2022+) — custom repository introspection, snapshotting, and content templates

Activity grew from 119 commits in 2019 to **2,075 in 2025**. The current auto-discovered 3-period config (origins / growth / maturity) captures Patchman's early arc but collapses too much of the content-services story into a single "maturity" bucket. **Five periods** better reflect the inflection points below.

---

## 1. Phases of Development

### When did the project start?

**October 2019** — `patchman-engine` initial commit (Oct 31) by @jiridostal, followed immediately by Docker scaffolding and Python/Go prototypes from @semtexzv and @josef-hak. **November 2019** — `patchman-ui` initial commit and React build setup by @jiridostal.

No other repos existed until **April 2022**, when @jlsherrill created `content-sources-backend`.

### When did it move from prototype to production?

**May 2020 – early 2021** marks the transition:

| Signal | Date | Evidence |
|--------|------|----------|
| Postgres 10 → 12 upgrade | May 2020 | Production database hardening |
| Export endpoints, search, metrics | May–Sep 2020 | Customer-facing API completeness |
| OpenAPI spec checker in CI | Aug 2020 | API contract enforcement |
| system_* table normalization | Jan 2021 | Schema ready for scale |
| semantic-release automation | 2020–2021 | Release engineering maturity |
| Floorist/DataHub analytics | Feb 2022 | Platform production reporting (PR #898) |

The **Go prototype won** over Python/Rust by Nov 2019; production hardening consumed most of 2020–2021. Patchman was a live Insights service before Content Sources existed.

### When did major architectural shifts happen?

| Shift | When | Impact |
|-------|------|--------|
| **Kafka-driven ingestion** | Nov 2019 – Sep 2020 | Core event architecture for patchman-engine |
| **System table decomposition** | Jan 2021 | system_platform, system_advisories, system_repo splits |
| **org_id multi-tenancy** | May–Jun 2022 | Platform identity model (PR #960) |
| **Content Sources founding** | Apr–Jun 2022 | New product line: backend + frontend + yummy |
| **Kafka introspection** | Oct 2022 | Event-driven repo introspection (PR #113) |
| **Pulp + zest bindings** | Mar 2023 | Generated clients replace hand-written Pulp HTTP code |
| **system_package2 migration** | Dec 2023 | Largest patchman-engine schema change in years |
| **Content templates domain** | Jan 2024 | New API surface replacing patch sets (PR #510) |
| **Template ↔ Patch integration** | Nov 2024 – present | Feature-flagged UI bridge (PR #1224), cross-repo work |
| **Playwright E2E adoption** | Jan 2025 – present | Testing paradigm shift across 3 repos |
| **Konflux CI migration** | Jul 2025 – present | Build pipeline modernization |
| **DB migration as ClowdApp Job** | Jun 2026 | Deployment architecture fix (PR #2236) |

### When did community grow significantly?

| Inflection | Commits/month | Cause |
|------------|---------------|-------|
| **Jun–Jul 2022** | 171 → 191 (2× prior avg) | Content Sources team joins; org_id migration; patch templates UI |
| **Mar 2023** | 232 (2× prior avg) | zest creation; Pulp integration ramp; zest binding automation |
| **Q4 2025 – Q1 2026** | 702 → 674/quarter | Playwright rollout, Konflux, template Kafka events, highest sustained activity |

**Contributor growth**: 2 repos / ~5 people (2019) → 4 repos (2022) → 9 repos (2025). Bot contributors (dependabot, konflux) added ~1,800 PRs in 2024–2026 but human review culture scaled with @rverdile, @swadeley, @xbhouse, @jlsherrill as high-volume reviewers.

---

## 2. Recurring Themes

### What types of changes happen most frequently?

Commit message heuristics by year (approximate):

| Year | Fix | Feat | Migration | Test/Playwright | Chore/Deps | Other |
|------|-----|------|-----------|-----------------|------------|-------|
| 2020 | 18% | 4% | 2% | 0% | 16% | 60% |
| 2021 | 12% | 5% | 1% | 0% | 27% | 55% |
| 2022 | 13% | 2% | 2% | 0% | 12% | 71% |
| 2023 | 20% | 1% | 2% | 1% | 5% | 71% |
| 2024 | 23% | 0% | 1% | 0% | 10% | 65% |
| 2025 | 6% | 3% | 1% | **7%** | 14% | 69% |
| 2026 | 7% | 3% | 2% | 4% | 27% | 57% |

**Patterns:**
- **2020–2021**: Heavy feature + infrastructure work on patchman; rising dependency/chore commits as ecosystem matured
- **2022–2024**: "Other" dominates — merge commits, release tags, semantic-release, multi-repo coordination; fix ratio increases as products stabilize
- **2025+**: **Playwright/test commits surge** (147 in 2025 alone); Konflux/dependabot automation drives chore/deps spike

Theme analysis (`theme_analysis.json`) by auto-discovered period:

| Theme | Origins | Growth | Maturity |
|-------|---------|--------|----------|
| playwright | 92 | 221 | **944** |
| snapshots | 3 | 10 | **646** |
| pulp | — | — | **458** |
| content templates | 4 | 2 | **226** |
| introspection | — | — | **92** |

Maturity-period themes are overwhelmingly content-services concerns; origins/growth are patchman-dominated (playwright keyword matches patchman-ui snapshot tests, not Playwright E2E).

### Seasonal patterns

Commits pooled across all years by calendar month:

| High activity | Low activity |
|---------------|--------------|
| **March** (1,084) | **December** (539) |
| January (855) | November (678) |
| February (906) | September (698) |

**Q1 is consistently strong** — likely aligned with Red Hat fiscal/planning cycles and post-holiday sprint capacity. **Q4 dips** (especially December) suggest release freeze or holiday slowdown. No strong "fixes in Q4" pattern; fix rates are relatively flat year-round in 2023–2024.

Quarterly trend (recent): activity **accelerated sharply in Q3–Q4 2025** (588 → 702 commits/quarter) and remained high through Q1–Q2 2026 — the integration/modernization phase is the busiest era on record.

### How has focus shifted over time?

```
2019–2020   Patchman prototype → production
            [patchman-engine 70%] [patchman-ui 30%]

2020–2022   Patchman scale & platform alignment
            [engine 58%] [ui 39%] [CS repos barely exist]

2022–2023   Dual product lines emerge
            [engine 56%] [CS-backend 17%] [ui 16%] [CS-frontend 9%]

2023–2024   Pulp ecosystem & snapshot infrastructure
            [engine 45%] [CS-backend 23%] [zest 12%] [ui 11%]

2024–2025   Content templates + UI depth
            [engine 38%] [CS-backend 26%] [CS-frontend 16%]

2025–2026   Cross-product integration & tooling
            [engine 40%] [CS-frontend 21%] [CS-backend 16%] [proxy 4%]
```

**Narrative arc**: Single product → dual products → shared template domain → unified testing/CI infrastructure.

---

## 3. Inflection Points

### Sudden activity increases

| Date | Δ | Trigger |
|------|---|---------|
| **2022-06** | 171 commits (+133% vs 6mo avg) | content-sources-backend/frontend go live; org_id migration; patch template UI work begins |
| **2022-07** | 191 commits | CS feature sprint (introspection, deployment, pagination); patchman Go migration tooling |
| **2023-03** | 232 commits | zest repo created; automated Pulp binding releases begin |
| **2025-Q3** | 588 commits/quarter | Playwright CI across repos; Konflux adoption begins |
| **2025-Q4 – 2026-Q1** | 702 → 674/quarter | Sustained peak; template Kafka events; PF6/TypeScript migrations |

### Major refactorings

| Refactoring | Period | Repos |
|-------------|--------|-------|
| Python/Rust → Go (engine) | Nov 2019 | patchman-engine |
| system_* table splits | Jan 2021 | patchman-engine |
| org_id identity migration | May–Jun 2022 | patchman-engine |
| yummy extraction from backend | Nov–Dec 2022 | yummy, content-sources-backend |
| system_package → system_package2 | Dec 2023 | patchman-engine |
| Patch sets → content templates | 2023–2024 | content-sources-backend, patchman-ui |
| Enzyme → RTL | 2023–2024 | patchman-ui |
| PatternFly 5 → 6 | 2024–2025 | patchman-ui |
| Cyndi system_platform split | 2025–2026 | patchman-engine |

### New subsystem introductions

| Subsystem | Repo | First activity |
|-----------|------|----------------|
| Kafka ingestion | patchman-engine | Nov 2019 |
| Floorist analytics | patchman-engine | Feb 2022 |
| Content Sources API | content-sources-backend | Apr 2022 |
| Repository introspection (yummy) | yummy → backend | May 2022 |
| Kafka introspection | content-sources-backend | Oct 2022 |
| Pulp bindings generator | zest | Mar 2023 |
| Snapshot API | content-sources-backend | Nov 2023 |
| Snapshot errata service | tang | Nov 2023 |
| Content templates API | content-sources-backend | Jan 2024 |
| Candlepin bindings | caliri | Feb 2024 |
| Playwright E2E | content-sources-backend | Jan 2025 |
| Dev proxy | frontend-development-proxy | May 2025 |
| Template Kafka events | content-sources-backend | Jun 2026 |

---

## 4. Proposed Period Structure

### Recommendation: 5 periods

Refines the auto-discovered 3-period config in `.history-analysis-config.json`. Each period has a clear dominant narrative, 691–2,600 commits, and natural boundary dates tied to inflection points.

---

### Period 1: **Genesis** (Oct 2019 – Apr 2020)
**691 commits · 2 repos**

| | |
|---|---|
| **Theme** | Prototype, language selection, initial architecture |
| **Dominant repos** | patchman-engine (70%), patchman-ui (30%) |
| **Key actors** | @jiridostal, @semtexzv, @josef-hak |
| **Defining work** | Go/Python/Rust prototypes; Kafka + Postgres foundation; React/Redux UI bootstrap |

**Why this boundary**: Activity is exclusively patchman; commit rate is low (~100/quarter) and exploratory. Ends as production hardening begins May 2020.

---

### Period 2: **Production Patchman** (May 2020 – May 2022)
**1,837 commits · 2 repos (+ early CS seeds)**

| | |
|---|---|
| **Theme** | Scale, platform integration, production hardening |
| **Dominant repos** | patchman-engine (58%), patchman-ui (39%) |
| **Key actors** | @josef-hak, @MichaelMraka, @mhornick, @jiridostal, @mkholjuraev |
| **Defining work** | Postgres upgrades; system table migrations; Floorist analytics; export APIs; patch template UI begins; org_id migration at period end |

**Why this boundary**: Patchman reaches production maturity. Ends Jun 2022 when content-sources-backend (Apr 2022) and frontend (Jun 2022) accelerate and org_id migration lands — the workspace becomes multi-product.

---

### Period 3: **Expansion** (Jun 2022 – Dec 2023)
**2,284 commits · 6 repos**

| | |
|---|---|
| **Theme** | Content Sources birth, introspection, Pulp ecosystem |
| **Dominant repos** | patchman-engine (46%), content-sources-backend (22%), patchman-ui (14%), zest (9%) |
| **Key actors** | @jlsherrill, @Andrewgdewar, @rverdile, @psegedy, @mkholjuraev, @xbhouse (from Oct 2023) |
| **Defining work** | CS backend/frontend/yummy founding; Kafka introspection; zest Pulp bindings; snapshot endpoint; tang service; system_package2 migration; patch templates UI |

**Why this boundary**: Dec 2023 is a natural break — snapshot infrastructure complete (PR #458 merged), tang/yummy mature, system_package migration finishes, @xbhouse joins and lands groups/envs API. Content **templates CRUD** begins next month (PR #510, Jan 2024).

---

### Period 4: **Templates Era** (Jan 2024 – Jun 2025)
**2,062 commits · 8 repos**

| | |
|---|---|
| **Theme** | Content templates domain, snapshot maturity, UI modernization |
| **Dominant repos** | patchman-engine (38%), content-sources-backend (26%), content-sources-frontend (16%) |
| **Key actors** | @rverdile, @xbhouse, @swadeley, @Andrewgdewar, @leSamo, @dominikvagner |
| **Defining work** | Content templates CRUD (PR #510); template-for-date API; caliri Candlepin bindings; patch templates → content templates UI bridge (PR #1224); PF6/RTL/TypeScript on patchman-ui; snapshot cleanup automation |

**Why this boundary**: Jul 2025 marks Konflux pipeline adoption (frontend-development-proxy PR #1), Playwright backend POC (Feb 2025) scaling up, and the start of sustained Q3 2025 activity surge — shift from **building the template domain** to **integrating and modernizing** the whole program.

---

### Period 5: **Convergence** (Jul 2025 – present)
**2,600 commits · 9 repos** *(ongoing)*

| | |
|---|---|
| **Theme** | Cross-product integration, Playwright E2E, Konflux CI, template events |
| **Dominant repos** | patchman-engine (40%), content-sources-frontend (21%), content-sources-backend (16%), patchman-ui (12%) |
| **Key actors** | @swadeley, @xbhouse, @TenSt, @dominikvagner, @Dugowitch |
| **Defining work** | Playwright across backend/frontend/patchman-ui; Konflux pipelines; template Kafka events (PR #1537); template-advisory patchman integration; Cyndi table split; DB migration Job (PR #2236); dev proxy tooling |

**Why this boundary**: All 9 repos active; activity at all-time highs; work crosses product boundaries routinely. Period remains open.

---

## 5. Comparison to Auto-Discovered Config

| Auto-discovered (3 periods) | Proposed (5 periods) | Rationale for split |
|----------------------------|---------------------|---------------------|
| origins (2019–2020) | **genesis** (same) | Good fit |
| growth (2020–2022) | **production** (same) | Rename for clarity |
| *(part of maturity)* | **expansion** (2022–2023) | CS founding lost in "maturity" |
| *(part of maturity)* | **templates era** (2024–2025) | Template domain is distinct arc |
| maturity (2022–present) | **convergence** (2025–present) | Integration/modernization is current phase |

### Suggested updates to `.history-analysis-config.json`

If narrative generation proceeds with these periods, update the config:

```json
{
  "periods": {
    "genesis": {
      "start": "2019-10-01",
      "end": "2020-04-30",
      "description": "Patchman prototype and language selection"
    },
    "production": {
      "start": "2020-05-01",
      "end": "2022-05-31",
      "description": "Patchman production hardening and platform integration"
    },
    "expansion": {
      "start": "2022-06-01",
      "end": "2023-12-31",
      "description": "Content Sources founding and Pulp/snapshot infrastructure"
    },
    "templates-era": {
      "start": "2024-01-01",
      "end": "2025-06-30",
      "description": "Content templates domain and UI modernization"
    },
    "convergence": {
      "start": "2025-07-01",
      "end": "2030-12-31",
      "description": "Cross-product integration, Playwright, and Konflux modernization"
    }
  }
}
```

Re-run `./tools/generate-history-draft --verbose` after updating to regenerate timeline files per period.

---

## 6. Implications for Narrative Writing (Phase 4)

| Period | Narrative focus | Key files to deep-dive |
|--------|----------------|------------------------|
| genesis | Why Go won; Kafka-from-day-one; paired UI/engine | origins timeline, patchman-engine PRs #4–#27 |
| production | Scaling pain; identity; analytics | growth timeline, PR #898, #960, system migrations |
| expansion | Second product line; introspection architecture | maturity timeline 2022–2023, PRs #4, #55, #113, #458 |
| templates-era | Template domain design; patch UI transition | PRs #510, #486, #1224, patchman-ui template PRs |
| convergence | Integration complexity; testing/CI modernization | PRs #956, #1445, #1537, #2236, key_contributors.md |

---

*Generated: Phase 2, Prompt 3 · Sources: `history/draft/timeline/*.md`, git logs, `theme_analysis.json`, `architectural_milestones.md`, `key_contributors.md`*
