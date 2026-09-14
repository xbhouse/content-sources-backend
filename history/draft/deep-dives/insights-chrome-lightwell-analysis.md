# insights-chrome — Lightwell shell (2026-07-01 onward)

Repo: [RedHatInsights/insights-chrome](https://github.com/RedHatInsights/insights-chrome)  
Period: lightwell only. Pre-July 2026 chrome history is out of scope.  
Source: local git (PR files are stubs; GitHub fetch was not run).

## Why this repo is here

Lightwell is not only a content-sources frontend module. It needed a console layout that is not Insights: no Beta switcher, no Help/Notifications masthead, felt/glass theme, Lightwell logo, later a horizontal subnav and a console nav entry. That work lives in insights-chrome.

## Week one (7/1–7/8)

| Date | PR | Author | What |
|---|---|---|---|
| 07-01 | #3582 | @florkbr | Scalprum route / layout for `/lightwell` |
| 07-02 | #3585 | chrome platform | Felt theme on `/lightwell` |
| 07-03 | #3589 | @marusak | Template with hidden Insights navigation |
| 07-06 | #3590 | @marusak | Do not pre-render footer when unused |
| 07-07 | #3591 | chrome platform | Simplify header |
| 07-07 | #3592 | @ochosi | Kill theme/banner flicker |
| 07-07 | #3596 | @ochosi | Red Hat Display header font |

## After week one

| Date | PR | Author | What |
|---|---|---|---|
| 07-24 | #3617 | @ochosi | Branding |
| 08-05 | #3638 | @karelhala | Main logo on Lightwell pages |
| 08-07 | #3616 | chrome platform | Add Lightwell to console navigation |
| 08-10 | #3645 | @karelhala | Clickable wordmark/logo |
| 08-14 | #3648 | @karelhala | Logout on Lightwell pages |
| 08-21 | #3656 | chrome platform | Horizontal subnav |
| 08-25 | #3670 | @marusak | Theme-aware icons |
| 08-31 | #3663 | @marusak | Horizontal nav adjustments |
| 09-01 | #3675 | @marusak | Same visual at all widths |
| 09-08 | #3676 | @dominikvagner | Drop last breadcrumb segment (Lens/Beacon routes) |
| 09-11 | #3679 | @katarinazaprazna | Minimal footer outside the page card |
