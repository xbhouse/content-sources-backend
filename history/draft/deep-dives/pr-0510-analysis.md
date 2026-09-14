# backend#510 — Fixes 2965: add create/fetch/list for templates (context)

Author: @rverdile  
Merged: 2024-01-15  
URL: https://github.com/content-services/content-sources-backend/pull/510  
Period: content-sources-expansion  
Weight: one-pager — why Pulp snapshots were used, not a Pulp origin story

## Problem

Snapshots (#248–#387) freeze repo state. Customers needed an object that says “these repos, this arch/version, this date” so a fleet can install from that pin. That object is a content template.

## Approach

Create, list, and fetch `/templates/` with filters on name, version, and arch. Delete (#535), update (#547), and the frontend table (#193) followed in weeks. First related backend work was #486 (Dec 2023): POST snapshots for `content_templates`. @swadeley hit OpenAPI `UUID` vs `uuid`; @xbhouse found fetch-on-bad-uuid returning “Could not find repository.”

## Why it matters here

Templates only work because Pulp distributions exist. Candlepin (#570/#639) later turns a template into a RHEL environment. For this history, #510 is the pay-off of 2023, not a new content engine.

## Quote

> “should probably say `Could not find template`” — @xbhouse, on a 404 that reused repository error text
