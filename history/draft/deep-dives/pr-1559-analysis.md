# backend#1559 — HMS-10936/10938: import for Lightwell repositories

Author: @rverdile  
Merged: 2026-07-02  
URL: https://github.com/content-services/content-sources-backend/pull/1559  
Period: lightwell

## Problem

Lightwell content already lives in Pulp. Pulp is the source of truth for versions and builds. It is not a customer-typed URL that content-sources should introspect. Reusing `origin=red_hat` (`org_id=-1`, #407) would mix a second product’s catalog into Insights RH repos. The backend needed a way to *import* a fixed set of Lightwell repos into `repository_configurations` without exposing them on the existing list/fetch APIs yet.

## Approach

@rverdile added an import path (`external-repos import`) that creates two Lightwell repo rows. Re-running import is idempotent. List/fetch and feature-name on import were deliberately omitted: “focus on the importing and the database changes.” Local testing pointed compose Pulp at stage Pulp, migrated, and imported.

@jlsherrill immediately hit `enforce_consistent_org_id`: listing repos returned 500 “Organization ID mismatch” because imported Lightwell org IDs are not the caller’s Insights org. That middleware was written for custom-repo tenancy, not a shared Lightwell catalog.

## Design decisions

- Import, do not introspect. Same command family as RH `repos-import` (#407), different product rows.
- Hide from list/fetch in v1. Avoid shipping the wrong list contract. HMS-10941 / fetch query (#1561) followed within days.
- Org-id middleware is Insights-shaped. Lightwell orgs required an exception; this is the first sign the shared backend has two tenancy models.

## Impact

This is the Lightwell on-ramp. Package APIs (#1560), Python on the list (#1562), and the JFrog advisory bridge (#1664) all assume these rows exist. Frontend empty table / new app module and the insights-chrome `/lightwell` shell landed the same week (2026-07-01). Insights partner-repos (#1596 and related HMS-107xx) are a different feature; do not treat them as Lightwell catalog APIs.

## Quote

> “I intentionally left out the lightwell repos from list/fetch methods because I wanted to think about the right way to handle that.” — @rverdile

> “we need to adjust the response checker to handle lightwell orgs” — @jlsherrill
