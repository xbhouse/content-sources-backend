---
type: pull_request
number: 1656
title: "Update grafana/grafana Docker tag to v12"
state: merged
author: red-hat-konflux
created: 2025-05-06T00:37:43Z
updated: 2025-10-08T16:16:34Z
closed: 2025-05-06T09:07:22Z
merged: 2025-05-06T09:07:21Z
base_branch: master
head_branch: konflux/mintmaker/master/major-grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1656
---

# Pull Request #1656: Update grafana/grafana Docker tag to v12

**Author**: @red-hat-konflux
**Created**: May 06, 2025 at 12:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/major-grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | major | `11.6.1` -> `12.0.0` |

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.0.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1200-2025-05-05)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v11.6.1...v12.0.0)

##### Features and enhancements

-   **Alerting:** API to convert submitted Prometheus rules to GMA [#&#8203;102231](https://redirect.github.com/grafana/grafana/pull/102231), [@&#8203;fayzal-g](https://redirect.github.com/fayzal-g)
-   **Alerting:** Add HMAC signature config to the webhook integration [#&#8203;100960](https://redirect.github.com/grafana/grafana/pull/100960), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
-   **Alerting:** Add MissingSeriesEvalsToResolve to the APIs [#&#8203;102150](https://redirect.github.com/grafana/grafana/pull/102150), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
-   **Alerting:** Add UI migration feature toggle [#&#8203;102217](https://redirect.github.com/grafana/grafana/pull/102217), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Add backend support for keep_firing_for [#&#8203;100750](https://redirect.github.com/grafana/grafana/pull/100750), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
-   **Alerting:** Add details and edit pages for groups [#&#8203;100884](https://redirect.github.com/grafana/grafana/pull/100884), [@&#8203;konrad147](https://redirect.github.com/konrad147)
-   **Alerting:** Add keep_firing_for and Recovering state [#&#8203;103248](https://redirect.github.com/grafana/grafana/pull/103248), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Add migration to clean up rule versions table [#&#8203;102484](https://redirect.github.com/grafana/grafana/pull/102484), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Add missing_series_evals_to_resolve option to alert rule form [#&#8203;102808](https://redirect.github.com/grafana/grafana/pull/102808), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Delete permanently deleted alert rules. [#&#8203;102960](https://redirect.github.com/grafana/grafana/pull/102960), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Detect target folder rules and show warning [#&#8203;103673](https://redirect.github.com/grafana/grafana/pull/103673), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Migration UI [#&#8203;102010](https://redirect.github.com/grafana/grafana/pull/102010), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Recover deleted alert rules [#&#8203;101869](https://redirect.github.com/grafana/grafana/pull/101869), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Remove constraints for uniqueness of rule title [#&#8203;102067](https://redirect.github.com/grafana/grafana/pull/102067), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Remove feature flag `alertingNoDataErrorExecution` [#&#8203;102156](https://redirect.github.com/grafana/grafana/pull/102156), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Sequential evaluation of rules in group [#&#8203;98829](https://redirect.github.com/grafana/grafana/pull/98829), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Skip rules that are managed by plugins when importing datasource-managed rules [#&#8203;103573](https://redirect.github.com/grafana/grafana/pull/103573), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Stop allowing manual editing/restore of internal AM config via settings [#&#8203;103884](https://redirect.github.com/grafana/grafana/pull/103884), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Template preview enhancements [#&#8203;103817](https://redirect.github.com/grafana/grafana/pull/103817), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **Alerting:** Update alerting module to [`58ba6c6`](https://redirect.github.com/grafana/grafana/commit/58ba6c617ff05eb1d6f65c59d369a6a16923dff6) [#&#8203;102812](https://redirect.github.com/grafana/grafana/pull/102812), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Use 'Grafana IRM' wording in alerting contact point [#&#8203;102014](https://redirect.github.com/grafana/grafana/pull/102014), [@&#8203;brojd](https://redirect.github.com/brojd)
-   **Alerting:** Webhook Improvements - Templateable Payloads [#&#8203;103818](https://redirect.github.com/grafana/grafana/pull/103818), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **AppChrome:** Move kiosk button into profile menu [#&#8203;103600](https://redirect.github.com/grafana/grafana/pull/103600), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **AppPlatform:** Introduce experimental Github integration for dashboard configuration management [#&#8203;96329](https://redirect.github.com/grafana/grafana/pull/96329), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
-   **Authorization:** Add group to role DisplayName to make filtered list more clear [#&#8203;102950](https://redirect.github.com/grafana/grafana/pull/102950), [@&#8203;forsethc](https://redirect.github.com/forsethc)
-   **Azure Monitor:** Add logs query builder [#&#8203;99055](https://redirect.github.com/grafana/grafana/pull/99055), [@&#8203;alyssabull](https://redirect.github.com/alyssabull)
-   **Azure:** Mark Azure Prometheus exemplars as GA and enable by default [#&#8203;100595](https://redirect.github.com/grafana/grafana/pull/100595), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **AzureMonitor:** Improve selection of Basic Logs tables in the query builder [#&#8203;103820](https://redirect.github.com/grafana/grafana/pull/103820), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **BrowseDashboards:** Switch to list view if sort is set [#&#8203;102196](https://redirect.github.com/grafana/grafana/pull/102196), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
-   **Checkbox:** Add z-index to description [#&#8203;103847](https://redirect.github.com/grafana/grafana/pull/103847), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
-   **Chore:** Promoting feature toggle pluginsSriChecks GA [#&#8203;102212](https://redirect.github.com/grafana/grafana/pull/102212), [@&#8203;tolzhabayev](https://redirect.github.com/tolzhabayev)
-   **CloudMigrations:** Add sorting and error filtering to Snapshot Results backend [#&#8203;102753](https://redirect.github.com/grafana/grafana/pull/102753), [@&#8203;mmandrus](https://redirect.github.com/mmandrus)
-   **CloudMigrations:** Change onPremToCloudMigrations feature toggle to GA [#&#8203;103212](https://redirect.github.com/grafana/grafana/pull/103212), [@&#8203;dana-axinte](https://redirect.github.com/dana-axinte)
-   **CloudMigrations:** Enable high-level resource type selection [#&#8203;103011](https://redirect.github.com/grafana/grafana/pull/103011), [@&#8203;macabu](https://redirect.github.com/macabu)
-   **CloudMigrations:** Implement table sorting in the UI [#&#8203;103061](https://redirect.github.com/grafana/grafana/pull/103061), [@&#8203;mmandrus](https://redirect.github.com/mmandrus)
-   **CloudWatch:** Migrate to aws-sdk-go-v2 [#&#8203;103106](https://redirect.github.com/grafana/grafana/pull/103106), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
-   **Cloudwatch:** Do not parse log query grouping field to float [#&#8203;102244](https://redirect.github.com/grafana/grafana/pull/102244), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
-   **Cloudwatch:** Migrate to aws-sdk-go-v2 [#&#8203;99643](https://redirect.github.com/grafana/grafana/pull/99643), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
-   **Cloudwatch:** Revert aws sdk go v2 [#&#8203;103644](https://redirect.github.com/grafana/grafana/pull/103644), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
-   **Config:** Removes setting `viewers_can_edit` [#&#8203;102275](https://redirect.github.com/grafana/grafana/pull/102275), [@&#8203;eleijonmarck](https://redirect.github.com/eleijonmarck)
-   **Dashboard Restore:** Remove experimental functionality under feature flag `dashboardRestore` for now - this will be reworked [#&#8203;103204](https://redirect.github.com/grafana/grafana/pull/103204), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **Dashboards:** Add Dashboard Schema validation (1) [#&#8203;103662](https://redirect.github.com/grafana/grafana/pull/103662), [@&#8203;marcoabreu](https://redirect.github.com/marcoabreu)
-   **Dashboards:** Add a config setting that limits the number of series that will be displayed in a panel. Users can opt in to render all series. [#&#8203;103405](https://redirect.github.com/grafana/grafana/pull/103405), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
-   **Dashboards:** Prevent saving to a non-existent folder [#&#8203;103503](https://redirect.github.com/grafana/grafana/pull/103503), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **Dashboards:** Prevent version restore to same data [#&#8203;102665](https://redirect.github.com/grafana/grafana/pull/102665), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **Dependencies:** Bump github.com/redis/go-redis/v9 from 9.7.0 to 9.7.3 [#&#8203;102555](https://redirect.github.com/grafana/grafana/pull/102555), [@&#8203;dependabot\[bot\]](https://redirect.github.com/dependabot\[bot])
-   **Docs:** Standard Datetime units limited to millisecond precision [#&#8203;103610](https://redirect.github.com/grafana/grafana/pull/103610), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
-   **ElasticSearch:** Improve index pattern error messaging and docs [#&#8203;103899](https://redirect.github.com/grafana/grafana/pull/103899), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **ElasticSearch:** Make script field input a text area [#&#8203;103708](https://redirect.github.com/grafana/grafana/pull/103708), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Extensions:** Expose new observable APIs for accessing components and links [#&#8203;103063](https://redirect.github.com/grafana/grafana/pull/103063), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
-   **Feat:** Make expressions work with plugins that set `alerting:false` but `backend:true` in their `plugin.json` files [#&#8203;102232](https://redirect.github.com/grafana/grafana/pull/102232), [@&#8203;tolzhabayev](https://redirect.github.com/tolzhabayev)
-   **FlameGraphPanel:** Add units to standard options ([#&#8203;89815](https://redirect.github.com/grafana/grafana/issues/89815)) [#&#8203;102720](https://redirect.github.com/grafana/grafana/pull/102720), [@&#8203;snyderdan](https://redirect.github.com/snyderdan)
-   **Frontend:** Remove Angular [#&#8203;99760](https://redirect.github.com/grafana/grafana/pull/99760), [@&#8203;jackw](https://redirect.github.com/jackw)
-   **Go:** Bump to 1.24.2 [#&#8203;103521](https://redirect.github.com/grafana/grafana/pull/103521), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
-   **Go:** Bump to 1.24.2 (Enterprise)
-   **I18n:** Add 13 new languages for translations [#&#8203;102971](https://redirect.github.com/grafana/grafana/pull/102971), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
-   **Influx:** Support PDC for Influx SQL [#&#8203;103032](https://redirect.github.com/grafana/grafana/pull/103032), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **JWT:** Add org role mapping support to the JWT provider [#&#8203;101584](https://redirect.github.com/grafana/grafana/pull/101584), [@&#8203;QuentinBisson](https://redirect.github.com/QuentinBisson)
-   **K8s:** Dashboards: Add fine grained access control checks to /apis [#&#8203;104418](https://redirect.github.com/grafana/grafana/pull/104418), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **K8s:** Enable kubernetesClientDashboardsFolders by default [#&#8203;103843](https://redirect.github.com/grafana/grafana/pull/103843), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **LBAC for data sources:** PublicPreview and self serve enablement [#&#8203;102276](https://redirect.github.com/grafana/grafana/pull/102276), [@&#8203;eleijonmarck](https://redirect.github.com/eleijonmarck)
-   **Live:** Remove queryOverLive and live-service-web-worker experimental feature flags [#&#8203;103518](https://redirect.github.com/grafana/grafana/pull/103518), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
-   **Logs Panel:** Add ISO8601 date to log download files [#&#8203;102932](https://redirect.github.com/grafana/grafana/pull/102932), [@&#8203;gtk-grafana](https://redirect.github.com/gtk-grafana)
-   **Logs Table:** Add new Controls component to Explore [#&#8203;103467](https://redirect.github.com/grafana/grafana/pull/103467), [@&#8203;matyax](https://redirect.github.com/matyax)
-   **Logs:** Add new Controls component to Explore [#&#8203;103401](https://redirect.github.com/grafana/grafana/pull/103401), [@&#8203;matyax](https://redirect.github.com/matyax)
-   **Logs:** Always keep displayed fields with changed queries [#&#8203;102493](https://redirect.github.com/grafana/grafana/pull/102493), [@&#8203;svennergr](https://redirect.github.com/svennergr)
-   **Logs:** Clean up Explore meta information [#&#8203;103801](https://redirect.github.com/grafana/grafana/pull/103801), [@&#8203;matyax](https://redirect.github.com/matyax)
-   **Logs:** Prevent automatic scrolling on refresh after changing scroll position [#&#8203;102463](https://redirect.github.com/grafana/grafana/pull/102463), [@&#8203;matyax](https://redirect.github.com/matyax)
-   **MetricsDrilldown:** Advance `exploreMetricsUseExternalAppPlugin` feature toggle stage [#&#8203;102137](https://redirect.github.com/grafana/grafana/pull/102137), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **MetricsDrilldown:** Advance `exploreMetricsUseExternalAppPlugin` to GA [#&#8203;103653](https://redirect.github.com/grafana/grafana/pull/103653), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **MetricsDrilldown:** Mark `exploreMetricsUseExternalAppPlugin` as not frontend-only [#&#8203;102942](https://redirect.github.com/grafana/grafana/pull/102942), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **MetricsDrilldown:** Remove legacy Metrics Drilldown code paths [#&#8203;103845](https://redirect.github.com/grafana/grafana/pull/103845), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **MetricsDrilldown:** Restore link to Metrics Drilldown from Explore [#&#8203;104075](https://redirect.github.com/grafana/grafana/pull/104075), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **NodeGraph:** Add node graph algorithm layout option [#&#8203;102760](https://redirect.github.com/grafana/grafana/pull/102760), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
-   **Plugins:** Remove plugin dependency version (Enterprise)
-   **Plugins:** Remove sort by options from plugins catalog [#&#8203;102862](https://redirect.github.com/grafana/grafana/pull/102862), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Plugins:** Remove support for secrets manager plugins [#&#8203;101467](https://redirect.github.com/grafana/grafana/pull/101467), [@&#8203;wbrowne](https://redirect.github.com/wbrowne)
-   **Plugins:** Remove support for secrets manager plugins (Enterprise)
-   **Plugins:** Remove userStorageAPI feature toggle [#&#8203;102915](https://redirect.github.com/grafana/grafana/pull/102915), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Prometheus:** Add back [@&#8203;lezer/highlight](https://redirect.github.com/lezer/highlight) to dev dependency [#&#8203;102632](https://redirect.github.com/grafana/grafana/pull/102632), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Prometheus:** Add support for cloud partners Prometheus data sources [#&#8203;103482](https://redirect.github.com/grafana/grafana/pull/103482), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
-   **Prometheus:** Enable Combobox metric select by default [#&#8203;101045](https://redirect.github.com/grafana/grafana/pull/101045), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
-   **Prometheus:** Enable prometheusRunQueriesInParallel feature toggle by default [#&#8203;102127](https://redirect.github.com/grafana/grafana/pull/102127), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
-   **RecordedQueries:** Deprecate recorded queries UI messaging (Enterprise)
-   **Security:** Update JWT library (CVE-2025-30204) [#&#8203;102715](https://redirect.github.com/grafana/grafana/pull/102715), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
-   **Tempo:** Add support for ad-hoc filters [#&#8203;102448](https://redirect.github.com/grafana/grafana/pull/102448), [@&#8203;ifrost](https://redirect.github.com/ifrost)
-   **Tempo:** Remove aggregate by [#&#8203;98474](https://redirect.github.com/grafana/grafana/pull/98474), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
-   **TraceView:** Add scope attributes to span details [#&#8203;103173](https://redirect.github.com/grafana/grafana/pull/103173), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
-   **TraceView:** Render all links in span details [#&#8203;101881](https://redirect.github.com/grafana/grafana/pull/101881), [@&#8203;ifrost](https://redirect.github.com/ifrost)
-   **Traces:** Preinstall Traces Drilldown app with Grafana [#&#8203;102986](https://redirect.github.com/grafana/grafana/pull/102986), [@&#8203;ifrost](https://redirect.github.com/ifrost)

##### Bug fixes

-   **Alerting:** Fix Simple condition threshold inputs with negative values. [#&#8203;102976](https://redirect.github.com/grafana/grafana/pull/102976), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Fix display of `Normal (Updated)` in alert history [#&#8203;102476](https://redirect.github.com/grafana/grafana/pull/102476), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Fix rule instances table [#&#8203;102290](https://redirect.github.com/grafana/grafana/pull/102290), [@&#8203;konrad147](https://redirect.github.com/konrad147)
-   **Alerting:** Make nested folders work in Alert List Panel [#&#8203;103550](https://redirect.github.com/grafana/grafana/pull/103550), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Remove rule type switch for modified export mode [#&#8203;102287](https://redirect.github.com/grafana/grafana/pull/102287), [@&#8203;konrad147](https://redirect.github.com/konrad147)
-   **Alerting:** Simplified alert rule toggle bug fixes [#&#8203;102119](https://redirect.github.com/grafana/grafana/pull/102119), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alertmanager:** Add Role-Based Access Control via reqAction Field [#&#8203;101543](https://redirect.github.com/grafana/grafana/pull/101543), [@&#8203;olegpixel](https://redirect.github.com/olegpixel)
-   **App Platform:** Pin bleve to fix CVE-2022-31022 [#&#8203;102513](https://redirect.github.com/grafana/grafana/pull/102513), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
-   **AppChrome/MegaMenu:** Fixes issue with default state being initialised to undocked [#&#8203;103507](https://redirect.github.com/grafana/grafana/pull/103507), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **AppTitle:** Fix overflowing text [#&#8203;103583](https://redirect.github.com/grafana/grafana/pull/103583), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
-   **Azure:** Ensure basic logs queries are limited to a single resource [#&#8203;103588](https://redirect.github.com/grafana/grafana/pull/103588), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **CloudWatch:** Import new grafana-aws-sdk with PDC fix [#&#8203;103249](https://redirect.github.com/grafana/grafana/pull/103249), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
-   **ColorPicker:** Fixed height when switching tabs [#&#8203;103304](https://redirect.github.com/grafana/grafana/pull/103304), [@&#8203;DanMPA](https://redirect.github.com/DanMPA)
-   **Dashboard:** Fix Core Panel Migrations - table panel [#&#8203;102146](https://redirect.github.com/grafana/grafana/pull/102146), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
-   **DashboardScenePage:** Correct slug in self referencing data links [#&#8203;100048](https://redirect.github.com/grafana/grafana/pull/100048), [@&#8203;Sergej-Vlasov](https://redirect.github.com/Sergej-Vlasov)
-   **Dashboards:** Fix duplicate provisioning when errors occur on title-only based provisioning [#&#8203;102249](https://redirect.github.com/grafana/grafana/pull/102249), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **Dashboards:** Fix panel link to Grafana Metrics Drilldown [#&#8203;103759](https://redirect.github.com/grafana/grafana/pull/103759), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **Fix:** Change secure_json_data column data type to medium text only MYSQL [#&#8203;102557](https://redirect.github.com/grafana/grafana/pull/102557), [@&#8203;s4kh](https://redirect.github.com/s4kh)
-   **GrafanaUI:** Prevent ToolbarButton from submitting form [#&#8203;102228](https://redirect.github.com/grafana/grafana/pull/102228), [@&#8203;kozhuhds](https://redirect.github.com/kozhuhds)
-   **GrafanaUI:** Remove blurred background from overlay backdrops to improve performance [#&#8203;103563](https://redirect.github.com/grafana/grafana/pull/103563), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
-   **LDAP test:** Fix page crash [#&#8203;102587](https://redirect.github.com/grafana/grafana/pull/102587), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
-   **Navigation:** Fix bookmarks when Grafana is running under subpath [#&#8203;102679](https://redirect.github.com/grafana/grafana/pull/102679), [@&#8203;matejkubinec](https://redirect.github.com/matejkubinec)
-   **PanelEdit:** Fixes suggestions not applying options or field config [#&#8203;102675](https://redirect.github.com/grafana/grafana/pull/102675), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **PluginProxy:** Fix nil pointer in OAuth forwarding [#&#8203;103626](https://redirect.github.com/grafana/grafana/pull/103626), [@&#8203;moustafab](https://redirect.github.com/moustafab)
-   **Plugins:** Fix better UX for disabled Angular plugins [#&#8203;101333](https://redirect.github.com/grafana/grafana/pull/101333), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
-   **Plugins:** Fix support for adhoc filters with raw queries in InfluxDB [#&#8203;101966](https://redirect.github.com/grafana/grafana/pull/101966), [@&#8203;beejeebus](https://redirect.github.com/beejeebus)
-   **Renderer:** Fix regression on callback URL in plugin mode [#&#8203;103787](https://redirect.github.com/grafana/grafana/pull/103787), [@&#8203;AgnesToulet](https://redirect.github.com/AgnesToulet)
-   **SQL:** Fix builder crashes when any in selected [#&#8203;102871](https://redirect.github.com/grafana/grafana/pull/102871), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
-   **SSE:** Fix goroutine leak in math operation expression parsing [#&#8203;102380](https://redirect.github.com/grafana/grafana/pull/102380), [@&#8203;kylebrandt](https://redirect.github.com/kylebrandt)
-   **Tempo:** Add fixes for broken exemplars [#&#8203;103298](https://redirect.github.com/grafana/grafana/pull/103298), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)

##### Breaking changes

-   **Alerting:** Make $value return the query value in case when a single datasource is used [#&#8203;102301](https://redirect.github.com/grafana/grafana/pull/102301), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
-   **Alerting:** Relax permissions for access a rule [#&#8203;103664](https://redirect.github.com/grafana/grafana/pull/103664), [@&#8203;moustafab](https://redirect.github.com/moustafab)
-   **Alerting:** Remove feature toggles relating to Loki Alert State History [#&#8203;103540](https://redirect.github.com/grafana/grafana/pull/103540), [@&#8203;rwwiv](https://redirect.github.com/rwwiv)
-   **Alerting:** Remove the POST endpoint for the internal Grafana Alertmanager config [#&#8203;103819](https://redirect.github.com/grafana/grafana/pull/103819), [@&#8203;rwwiv](https://redirect.github.com/rwwiv)
-   **Anonymous:** Enforce org role Viewer setting [#&#8203;102070](https://redirect.github.com/grafana/grafana/pull/102070), [@&#8203;eleijonmarck](https://redirect.github.com/eleijonmarck)
-   **Chore:** Enable Grafana version check when installing plugins [#&#8203;103176](https://redirect.github.com/grafana/grafana/pull/103176), [@&#8203;andresmgot](https://redirect.github.com/andresmgot)
-   **Chore:** Enabling failWrongDSUID by default in Grafana 12 [#&#8203;102192](https://redirect.github.com/grafana/grafana/pull/102192), [@&#8203;tolzhabayev](https://redirect.github.com/tolzhabayev)
-   **Config:** Removes setting `viewers_can_edit` [#&#8203;101767](https://redirect.github.com/grafana/grafana/pull/101767), [@&#8203;eleijonmarck](https://redirect.github.com/eleijonmarck)
-   **Frontend:** Remove Angular (Enterprise)
-   **Plugin Extensions:** Clean up the deprecated APIs [#&#8203;102102](https://redirect.github.com/grafana/grafana/pull/102102), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
-   **Plugins:** Remove plugin dependency version [#&#8203;103728](https://redirect.github.com/grafana/grafana/pull/103728), [@&#8203;wbrowne](https://redirect.github.com/wbrowne)
-   **Tempo:** Remove traceQLStreaming feature toggle [#&#8203;103619](https://redirect.github.com/grafana/grafana/pull/103619), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)

##### Plugin development fixes & changes

-   **Combobox:** add grouping functionality [#&#8203;100603](https://redirect.github.com/grafana/grafana/pull/100603), [@&#8203;eledobleefe](https://redirect.github.com/eledobleefe)
-   **Grafana UI:** Add `columnGap` + `rowGap` to `Stack`/`Grid` [#&#8203;102883](https://redirect.github.com/grafana/grafana/pull/102883), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
-   **Grafana UI:** Clearly separate multiple warnings by using HTML tags [#&#8203;97979](https://redirect.github.com/grafana/grafana/pull/97979), [@&#8203;zenador](https://redirect.github.com/zenador)

<!-- 12.0.0 END -->

<!-- 11.6.1 START -->

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "before 5am" in timezone Europe/Prague, Automerge - "* 3-10 * * 1" in timezone Europe/Prague.

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->

## Summary by Sourcery

Update Grafana Docker image from version 11.6.1 to 12.0.0

New Features:
- Upgrade to Grafana 12.0.0 with significant improvements to alerting, dashboards, and plugin ecosystem

Chores:
- Update base Docker image tag

---

## Discussion

### Comment by @jira-linking on May 06, 2025 at 12:37 AM UTC

Commits missing Jira IDs:
3df006895b268e384989942fc8cdf8a1d45f7f34


### Comment by @sourcery-ai on May 06, 2025 at 12:37 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This pull request updates the base Grafana Docker image version specified in the Dockerfile.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Updated Grafana Docker image tag. | <ul><li>Changed the base image tag from `11.6.1` to `12.0.0` in the `FROM` instruction.</li></ul> | `dev/grafana/Dockerfile` |

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

### Comment by @codecov-commenter on May 06, 2025 at 12:42 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1656?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.21%. Comparing base [(`5410d9e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5410d9e2692e3142fe15bfedeb576031d144251f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`3df0068`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3df006895b268e384989942fc8cdf8a1d45f7f34?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1656   +/-   ##
=======================================
  Coverage   58.21%   58.21%           
=======================================
  Files         146      146           
  Lines       11397    11397           
=======================================
  Hits         6635     6635           
  Misses       4175     4175           
  Partials      587      587           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1656/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1656/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.21% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1656?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1656*
