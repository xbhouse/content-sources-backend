---
repo: tang
pr_number: 1
period: expansion
source: history/repos/tang/pull-requests/merged/
---

# Deep Dive: tang PR #1

## Problem / Motivation

Snapshot queries against Pulp were embedded in the monolithic backend — slow to test, hard to scale. A dedicated service was needed for RPM search across repository versions.

## Solution Approach

@rverdile created tang with exported `Tangy` interface (`RpmRepositoryVersionPackageSearch()`), internal test packages, podman-compose Pulp dev environment, CI integration test for RPM name search, lint CI.

## Key Design Decisions

- **Interface-first**: `Tangy` interface enables mocking in backend ( @jlsherrill: "would it make sense to include a Mock of the interface here? Instead of in our application" — mock added to tang).
- **Separate repo**: Microservice pattern matching yummy/zest.
- **Pulp compose for dev**: Self-contained dev/test environment.

## Impact

Offloaded snapshot errata queries from backend. @xbhouse later contributed errata sorting (#9) and CVE listing (#10). Backend integrates via HTTP client.

## Review Notes

@jlsherrill reviewed extensively. Merged Jan 2024 after Dec 2023 init.

> **Notable quote** (@jlsherrill): "would it make sense to include a Mock of the interface here? Instead of in our application."

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
