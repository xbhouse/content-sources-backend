# backend#935 — HMS-4863: remove YYYY-MM-DD in snapshots/for_date (context)

Author: @xbhouse  
Merged: 2025-01-13  
URL: https://github.com/content-services/content-sources-backend/pull/935  
Period: content-sources-expansion  
Weight: one-pager — the documented snapshot API break

## Problem

`POST /snapshots/for_date/` accepted RFC3339 and `YYYY-MM-DD` so Image Builder would keep working. Dual formats made date selection ambiguous (timezone, “which snapshot for that calendar day”). IB UI and backend had already switched.

## Approach

Reject `YYYY-MM-DD`; require RFC3339 (`2025-01-09T00:00:00Z`). @mayurilahane confirmed 400 vs 200. Summary: “To be safe, don't merge on a Friday.”

## Why it matters here

This is the one breaking change on the snapshot-by-date API that Image Builder and templates share. It is not Pulp architecture and not Lightwell. Cite it if the expansion chapter needs a single “we broke a contract on purpose” example.

## Quote

> “Before this PR, we supported both RFC3339 and YYYY-MM-DD formats to maintain backward compatibility with Image Builder.” — @xbhouse
