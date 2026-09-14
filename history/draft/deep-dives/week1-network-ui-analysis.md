# frontend#1059–#1094 — Lightwell Network UI, week one (2026-07-01 – 2026-07-08)

Period: lightwell  
Scope: Lightwell Network console this week, including the insights-chrome `/lightwell` shell. Notifications, Lens, Beacon, JFrog, and the Network API are in lightwell-later-surfaces-analysis.md. Insights partner-repos are a different product.

## What shipped

| Date | PR | Author | What |
|---|---|---|---|
| 07-01 | frontend#1059 | @xbhouse | Empty table, app entry, module |
| 07-02 | frontend#1061 | @xbhouse | Repo list (HMS-10939) |
| 07-03 | frontend#1063 | @marusak | Hide chrome that leaves Lightwell |
| 07-03 | frontend#1065 | @rverdile | Python on the repo list |
| 07-03 | frontend#1066 | @xbhouse | Package list (HMS-10957) |
| 07-04 | frontend#1067 | @xbhouse | Filtering |
| 07-05 | frontend#1068 | @rverdile | Java package details |
| 07-06 | frontend#1069 | @xbhouse | Python package details |
| 07-06 | frontend#1075 | @rverdile | Version switching; group in route |
| 07-06 | frontend#1074 | @ochosi | Chrome stylesheet overrides |
| 07-07 | frontend#1079 | @ochosi | Render routes during RBAC fetch |
| 07-07 | frontend#1088 | @arburka | Connect popover: auth + Maven/Gradle/pip/Artifactory/Nexus snippets |
| 07-08 | frontend#1094 | @arburka | Full Maven coordinate as package name |

Same week in insights-chrome: #3582 `/lightwell` route (@florkbr), #3585 felt theme, #3589 hidden Insights nav (@marusak), #3592 flicker fix (@ochosi). See [insights-chrome-lightwell-analysis.md](insights-chrome-lightwell-analysis.md).

Same week on the backend: #1557 feature filter, #1558 count cache, #1559 import, #1560/#1561 package + fetch APIs, #1562 Python repo type, tang#25/#26. All of those read Pulp.

## Connect snippets are not a store

#1088 (and later #1116) document how to point Maven, Gradle, pip, Artifactory, and Nexus *clients* at Lightwell Pulp distribution URLs. Artifactory in the UI is a how-to tab. Package versions still come from tang → Pulp today; the planned path is Pulp catalog APIs (pulp_maven#451, pulp_python#1359).

## After 7/8

Polish (copy buttons, `/lightwell` paths, empty states) continued in mid-July. New product surfaces started later: notifications (late July), Lens/Beacon (mid-August), JFrog bridge and Network API phase 1 (late August).
