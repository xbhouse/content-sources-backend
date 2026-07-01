---
type: pull_request
number: 1946
title: "Update grafana/grafana Docker tag to v12.3.0"
state: merged
author: red-hat-konflux
created: 2025-11-24T04:44:36Z
updated: 2025-12-01T08:41:52Z
closed: 2025-12-01T04:46:50Z
merged: 2025-12-01T04:46:50Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1946
---

# Pull Request #1946: Update grafana/grafana Docker tag to v12.3.0

**Author**: @red-hat-konflux
**Created**: November 24, 2025 at 04:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | minor | `12.2.1` -> `12.3.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.3.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1230-2025-11-19)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.2.2...v12.3.0)

##### Features and enhancements

- **API Clients:** Add lazy hooks to clients [#&#8203;113226](https://redirect.github.com/grafana/grafana/pull/113226), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **API clients:** Automatically set PATCH headers [#&#8203;111879](https://redirect.github.com/grafana/grafana/pull/111879), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **API clients:** Extract into a package [#&#8203;111810](https://redirect.github.com/grafana/grafana/pull/111810), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **API clients:** Extract into a package (Enterprise)
- **API clients:** Update API clients to include all endpoints & add hooks [#&#8203;113061](https://redirect.github.com/grafana/grafana/pull/113061), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **AccessControl:** Include hidden roles in service account role display [#&#8203;112924](https://redirect.github.com/grafana/grafana/pull/112924), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **AccessControl:** Increase limit of LBAC for Datasources rules [#&#8203;111560](https://redirect.github.com/grafana/grafana/pull/111560), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **Accessibility:** Wrap data source info onto 2 lines at small viewports [#&#8203;113033](https://redirect.github.com/grafana/grafana/pull/113033), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Alert Enrichment:** Add mutator to insert rule UID labels to allow for efficient use of labelSelector (Enterprise)
- **Alerting:** Add enrichment components to rule view page (Enterprise)
- **Alerting:** Add enrichment section to rule view page (Enterprise)
- **Alerting:** Add jitter support for periodic alert state storage to reduce database load spikes [#&#8203;111357](https://redirect.github.com/grafana/grafana/pull/111357), [@&#8203;softho0n](https://redirect.github.com/softho0n)
- **Alerting:** Add position-based matching for identical alert rules [#&#8203;112407](https://redirect.github.com/grafana/grafana/pull/112407), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Create alertingAlertRuleFormSchema in restrictedGrafanaApis [#&#8203;112794](https://redirect.github.com/grafana/grafana/pull/112794), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Display error message in central state history view [#&#8203;111445](https://redirect.github.com/grafana/grafana/pull/111445), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Enrichment per rule wip-2 (Enterprise)
- **Alerting:** Hide metadata if grouping by folder [#&#8203;113216](https://redirect.github.com/grafana/grafana/pull/113216), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Improve template ai helper prompt and add some examples (Enterprise)
- **Alerting:** Move enrichment tab between details and versions [#&#8203;110886](https://redirect.github.com/grafana/grafana/pull/110886), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Remove ai feedback button from alert form [#&#8203;112713](https://redirect.github.com/grafana/grafana/pull/112713), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Remove unused components [#&#8203;111320](https://redirect.github.com/grafana/grafana/pull/111320), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Remove useRulesSourcesWithRuler for SmartAlertTypeDetector [#&#8203;111623](https://redirect.github.com/grafana/grafana/pull/111623), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Surface remote AM silence creation errors properly [#&#8203;112757](https://redirect.github.com/grafana/grafana/pull/112757), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Triage [#&#8203;110339](https://redirect.github.com/grafana/grafana/pull/110339), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Alerting:** Triage rule details drawer [#&#8203;112055](https://redirect.github.com/grafana/grafana/pull/112055), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Update prompt examples for template AI Helper (Enterprise)
- **Alerting:** Update width to instance details drawer in Triage page [#&#8203;113209](https://redirect.github.com/grafana/grafana/pull/113209), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Use new enrichment endpoints in FE (Enterprise)
- **Alerting:** Use ruleUid as a prop instead of extracting it from the rule context (Enterprise)
- **Analytics:** Aggregate daily summary in datasources analytics (Enterprise)
- **Analytics:** Apply proper batching to Loki exports and add configurable settings (Enterprise)
- **Annotations:** Exclude internal dashboard id when saved via UID [#&#8203;111535](https://redirect.github.com/grafana/grafana/pull/111535), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Azure:** Use SSO settings in plugin context [#&#8203;112058](https://redirect.github.com/grafana/grafana/pull/112058), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Buttons:** Active style for buttons [#&#8203;111235](https://redirect.github.com/grafana/grafana/pull/111235), [@&#8203;gtk-grafana](https://redirect.github.com/gtk-grafana)
- **Caching:** Disable cache if datasource has oauthPassThru=true (Enterprise)
- **Canvas:** Allow non-icon bg image fields [#&#8203;112308](https://redirect.github.com/grafana/grafana/pull/112308), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Chore:** Add logsdrilldown replace to apps/iam/go.mod [#&#8203;112581](https://redirect.github.com/grafana/grafana/pull/112581), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **CloudWatch Logs:** Don't add console link to every field in the logs response [#&#8203;112230](https://redirect.github.com/grafana/grafana/pull/112230), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **CloudWatch Logs:** Support Log Anomalies query type [#&#8203;113067](https://redirect.github.com/grafana/grafana/pull/113067), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **CloudWatch:** Add syntax highlighting and autocomplete for logs diff command [#&#8203;111207](https://redirect.github.com/grafana/grafana/pull/111207), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch:** Add tracking for logs anomalies [#&#8203;113181](https://redirect.github.com/grafana/grafana/pull/113181), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Dashboard Controls:** Add annotations to the dashboard controls menu [#&#8203;112816](https://redirect.github.com/grafana/grafana/pull/112816), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **Dashboard Picker:** Update to use correct search + dashboards APIs [#&#8203;112341](https://redirect.github.com/grafana/grafana/pull/112341), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Dashboard:** Backend always set `metricEditorMode: 0` regardless `metricQueryType` and `expression` [#&#8203;111613](https://redirect.github.com/grafana/grafana/pull/111613), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Dashboards:** Add a new variable type called "Switch" [#&#8203;111366](https://redirect.github.com/grafana/grafana/pull/111366), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **Dashboards:** Hide error notifications in kiosk mode on dashboards [#&#8203;112390](https://redirect.github.com/grafana/grafana/pull/112390), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Dynamic Dashboards:** Expand dashboards\_init\_dashboard\_completed tracking info [#&#8203;111102](https://redirect.github.com/grafana/grafana/pull/111102), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **ErrorBoundary:** Report specific boundary type to Faro [#&#8203;112071](https://redirect.github.com/grafana/grafana/pull/112071), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
- **Explore:** Use compact mode only when targeting Tempo [#&#8203;113037](https://redirect.github.com/grafana/grafana/pull/113037), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **FeatureToggles:** Remove deprecated experimental apiserver [#&#8203;111617](https://redirect.github.com/grafana/grafana/pull/111617), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Fields Selector:** Add component and integrate with Logs and Logs table visualization [#&#8203;112534](https://redirect.github.com/grafana/grafana/pull/112534), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Flame Graph:** Anchor exact match when clicking a table symbol in search [#&#8203;111101](https://redirect.github.com/grafana/grafana/pull/111101), [@&#8203;samarthbagga-meesho](https://redirect.github.com/samarthbagga-meesho)
- **FlameGraph:** Improve prompt for open assistant to analyze flamegraph [#&#8203;113071](https://redirect.github.com/grafana/grafana/pull/113071), [@&#8203;simonswine](https://redirect.github.com/simonswine)
- **FolderPicker:** Don't show expand button for empty folders and move search icon [#&#8203;111872](https://redirect.github.com/grafana/grafana/pull/111872), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **FolderPicker:** Show parent folder when searching [#&#8203;111026](https://redirect.github.com/grafana/grafana/pull/111026), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Geomap:** Add a MapLibre style base layer [#&#8203;109841](https://redirect.github.com/grafana/grafana/pull/109841), [@&#8203;remogeissbuehler](https://redirect.github.com/remogeissbuehler)
- **Geomap:** Move beta layers to GA [#&#8203;113186](https://redirect.github.com/grafana/grafana/pull/113186), [@&#8203;drew08t](https://redirect.github.com/drew08t)
- **Go:** Update to 1.25.2 + golangci-lint v2.5.0 + golang.org/x/net v0.45.0 [#&#8203;112149](https://redirect.github.com/grafana/grafana/pull/112149), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Go:** Update to 1.25.3 [#&#8203;112359](https://redirect.github.com/grafana/grafana/pull/112359), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Grafana Advisor:** Prometheus Type Migration check [#&#8203;110853](https://redirect.github.com/grafana/grafana/pull/110853), [@&#8203;bossinc](https://redirect.github.com/bossinc)
- **Grafana Data Source:** Add random walk configuration options [#&#8203;113009](https://redirect.github.com/grafana/grafana/pull/113009), [@&#8203;nmarrs](https://redirect.github.com/nmarrs)
- **IAM:** Add uid column in team\_member DB table [#&#8203;112439](https://redirect.github.com/grafana/grafana/pull/112439), [@&#8203;dmihai](https://redirect.github.com/dmihai)
- **Jaeger:** Migrate API calls to gRPC endpoint [#&#8203;113297](https://redirect.github.com/grafana/grafana/pull/113297), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **LBAC for data sources:** Provide user feedback of potential performance loss from LBAC rules (Enterprise)
- **Library Panels:** Remove direct use of legacy search [#&#8203;112231](https://redirect.github.com/grafana/grafana/pull/112231), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Logs panel:** Respect selected fields for downloading logs [#&#8203;111753](https://redirect.github.com/grafana/grafana/pull/111753), [@&#8203;matyax](https://redirect.github.com/matyax)
- **Nav:** Render menu items as `p` tags so truncation logic can work [#&#8203;113248](https://redirect.github.com/grafana/grafana/pull/113248), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Navigation:** Move Cost management and billing plugin to root [#&#8203;111739](https://redirect.github.com/grafana/grafana/pull/111739), [@&#8203;gubjanos](https://redirect.github.com/gubjanos)
- **PanelTimeCompare:** Support saving time compare window [#&#8203;113150](https://redirect.github.com/grafana/grafana/pull/113150), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **PanelTimeSettings:** Support panel time range settings changes from dashboard in view mode [#&#8203;113027](https://redirect.github.com/grafana/grafana/pull/113027), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Plugins:** Install Grafana Pathfinder behind a feature flag [#&#8203;109909](https://redirect.github.com/grafana/grafana/pull/109909), [@&#8203;Jayclifford345](https://redirect.github.com/Jayclifford345)
- **PostgreSQL:** Support PGPASSFILE by making password optional [#&#8203;108856](https://redirect.github.com/grafana/grafana/pull/108856), [@&#8203;taraspos](https://redirect.github.com/taraspos)
- **Provisioning:** Watch file system for changes [#&#8203;112184](https://redirect.github.com/grafana/grafana/pull/112184), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Reporting:** Add support for schema v2 dashboards (Enterprise)
- **Reporting:** Wait for streaming to end before exporting CSVs (Enterprise)
- **SQL Expressions:** Add Functions to Allow list [#&#8203;113291](https://redirect.github.com/grafana/grafana/pull/113291), [@&#8203;kylebrandt](https://redirect.github.com/kylebrandt)
- **Snapshots:** Use appSubUrl for View all snapshots [#&#8203;111652](https://redirect.github.com/grafana/grafana/pull/111652), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **Span Details:** Bring back span id to span details [#&#8203;112411](https://redirect.github.com/grafana/grafana/pull/112411), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Span Details:** Wrap label values [#&#8203;112413](https://redirect.github.com/grafana/grafana/pull/112413), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Stars:** Refactor StarsToolbarButton and unify nav update logic [#&#8203;112582](https://redirect.github.com/grafana/grafana/pull/112582), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Stat/BarGauge:** Border radius tweak [#&#8203;112562](https://redirect.github.com/grafana/grafana/pull/112562), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Table:** Add some error-case handling to ImageCell [#&#8203;110461](https://redirect.github.com/grafana/grafana/pull/110461), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Allow FieldType.other containing arrays to use Pills [#&#8203;111205](https://redirect.github.com/grafana/grafana/pull/111205), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Disable virtualization, hover overflow, and scrollbar width resizing on Safari 26 [#&#8203;111834](https://redirect.github.com/grafana/grafana/pull/111834), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Pill and JSON Cells should allow formatting [#&#8203;111951](https://redirect.github.com/grafana/grafana/pull/111951), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Support DataLinks and Actions in SparklineCell [#&#8203;112244](https://redirect.github.com/grafana/grafana/pull/112244), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Update ad-hoc filter to use name instead of displayName [#&#8203;112815](https://redirect.github.com/grafana/grafana/pull/112815), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Tempo:** Migrates tags and tag values to datasource backend CallResource requests (Enterprise)
- **Theme:** Changes light theme canvas color a more white shade [#&#8203;111318](https://redirect.github.com/grafana/grafana/pull/111318), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Themes:** Update themes border radius [#&#8203;111478](https://redirect.github.com/grafana/grafana/pull/111478), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **TimeComparison:** Automatically show/hide menu on hover [#&#8203;112750](https://redirect.github.com/grafana/grafana/pull/112750), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **TimeSeries:** Allow custom time units on x-axis [#&#8203;112913](https://redirect.github.com/grafana/grafana/pull/112913), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
- **Timeseries:** Numeric duration values could render as NaN ([#&#8203;73795](https://redirect.github.com/grafana/grafana/issues/73795)) [#&#8203;112076](https://redirect.github.com/grafana/grafana/pull/112076), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Transformations:** Hide "Match all/any" conditions for less than two filters [#&#8203;109754](https://redirect.github.com/grafana/grafana/pull/109754), [@&#8203;sudoice](https://redirect.github.com/sudoice)
- **UI Extensions:** Remove path validation from link extensions [#&#8203;112259](https://redirect.github.com/grafana/grafana/pull/112259), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)

##### Bug fixes

- **Access Control:** Fix the permission checks for saving/updating/deleting annotations [#&#8203;112953](https://redirect.github.com/grafana/grafana/pull/112953), [@&#8203;IevaVasiljeva](https://redirect.github.com/IevaVasiljeva)
- **Accessibility:** Improve no-unreduced-motion rule and fix violations [#&#8203;110304](https://redirect.github.com/grafana/grafana/pull/110304), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Alerting Provisioning:** Don't error on recording rules without conditions [#&#8203;109410](https://redirect.github.com/grafana/grafana/pull/109410), [@&#8203;djpnicholls](https://redirect.github.com/djpnicholls)
- **Alerting:** Clear outdated settings when switching contact point type [#&#8203;111869](https://redirect.github.com/grafana/grafana/pull/111869), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix enrichment tab to be rendered only for grafana alerting rules [#&#8203;113030](https://redirect.github.com/grafana/grafana/pull/113030), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix instances matching in notification policies [#&#8203;112326](https://redirect.github.com/grafana/grafana/pull/112326), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix threshold params [#&#8203;111645](https://redirect.github.com/grafana/grafana/pull/111645), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix unmarshalling of GettableStatus to include time intervals [#&#8203;112602](https://redirect.github.com/grafana/grafana/pull/112602), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Migrate `spec.title` and `spec.name` fieldSelectors [#&#8203;111993](https://redirect.github.com/grafana/grafana/pull/111993), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Alerting:** Normalize health when filtering rules [#&#8203;113087](https://redirect.github.com/grafana/grafana/pull/113087), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Alerting:** Prohibit receivers with empty name [#&#8203;113064](https://redirect.github.com/grafana/grafana/pull/113064), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Provisioning to fix contact point type on save [#&#8203;112246](https://redirect.github.com/grafana/grafana/pull/112246), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Remove \_\_grafana\_origin when duplicating rule [#&#8203;112396](https://redirect.github.com/grafana/grafana/pull/112396), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **AnnoList:** Fix annotations not loading when in a repeated row [#&#8203;111540](https://redirect.github.com/grafana/grafana/pull/111540), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **Annotations:** Fix issue with transformation logic in scenes [#&#8203;112288](https://redirect.github.com/grafana/grafana/pull/112288), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Auth:** Fix render user OAuth passthrough [#&#8203;111636](https://redirect.github.com/grafana/grafana/pull/111636), [@&#8203;charandas](https://redirect.github.com/charandas)
- **ComboBox:** Add loading state to dropdown and prefixIcon [#&#8203;112967](https://redirect.github.com/grafana/grafana/pull/112967), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Connections:** Fix connections home page on enterprise [#&#8203;111751](https://redirect.github.com/grafana/grafana/pull/111751), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
- **Dashboard:** Fix editor specific permissions in /api [#&#8203;113292](https://redirect.github.com/grafana/grafana/pull/113292), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboards:** Fix bug with anon users with editor permissions creating dashboards [#&#8203;113260](https://redirect.github.com/grafana/grafana/pull/113260), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboards:** Fix missing Ctrl+O keyboard shortcut for crosshair toggle [#&#8203;111310](https://redirect.github.com/grafana/grafana/pull/111310), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Dashboards:** Fix moving to root folder [#&#8203;111515](https://redirect.github.com/grafana/grafana/pull/111515), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboards:** Fix preload field not being persisted via /v1beta1 [#&#8203;112475](https://redirect.github.com/grafana/grafana/pull/112475), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Flame Graph:** Use suffix for values formatted with a short formatter [#&#8203;110999](https://redirect.github.com/grafana/grafana/pull/110999), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **FlameGraph:** Ensure total is only counted once for recursive function calls [#&#8203;111548](https://redirect.github.com/grafana/grafana/pull/111548), [@&#8203;simonswine](https://redirect.github.com/simonswine)
- **FolderPermissions:** Return 404 error when folder does not exist instead of 500 [#&#8203;112919](https://redirect.github.com/grafana/grafana/pull/112919), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **FolderPicker:** Fix expand toggle also selecting folder [#&#8203;111755](https://redirect.github.com/grafana/grafana/pull/111755), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Graphite:** Fix legacy response unmarshalling [#&#8203;112968](https://redirect.github.com/grafana/grafana/pull/112968), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Histogram:** Properly handle sparse heatmap-cells frames [#&#8203;112907](https://redirect.github.com/grafana/grafana/pull/112907), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
- **LDAP Authentication:** Fix URL to propagate username context as parameter [#&#8203;111723](https://redirect.github.com/grafana/grafana/pull/111723), [@&#8203;bradleypettit](https://redirect.github.com/bradleypettit)
- **Node graph:** Fix context menu position after scrolling [#&#8203;112374](https://redirect.github.com/grafana/grafana/pull/112374), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)
- **Playlist:** Fix navigation issues with emoji-titled dashboards during dual-write migration [#&#8203;111659](https://redirect.github.com/grafana/grafana/pull/111659), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Plugin Details Page:** Fix tabs not loading on hard refresh [#&#8203;112915](https://redirect.github.com/grafana/grafana/pull/112915), [@&#8203;sunker](https://redirect.github.com/sunker)
- **Plugin navigation:** Fix active nav item selection when there are more than 10 items in a group [#&#8203;112886](https://redirect.github.com/grafana/grafana/pull/112886), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Plugins:** Dependencies do not inherit parent URL for preinstall [#&#8203;111762](https://redirect.github.com/grafana/grafana/pull/111762), [@&#8203;wbrowne](https://redirect.github.com/wbrowne)
- **Plugins:** Set isProvisioned for local plugins without remote counterpart [#&#8203;111268](https://redirect.github.com/grafana/grafana/pull/111268), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
- **Prometheus:** Fix incremental querying logic for public dashboards [#&#8203;111642](https://redirect.github.com/grafana/grafana/pull/111642), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **Prometheus:** Fix parsing logic of prometheus expressions to honor the order of binary operations [#&#8203;112220](https://redirect.github.com/grafana/grafana/pull/112220), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **Security:** fix for CVE-2025-41115 in SCIM (System for Cross-domain Identity Management) (Enterprise)
- **SoloPanel:** Fixes issue with solo route and scopes variable [#&#8203;112769](https://redirect.github.com/grafana/grafana/pull/112769), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Stars:** Fix starred state not being updated [#&#8203;111936](https://redirect.github.com/grafana/grafana/pull/111936), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **Stat:** Fix math for percent change value heights when sparkline is not rendered [#&#8203;112599](https://redirect.github.com/grafana/grafana/pull/112599), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **StateTimeline:** Fix color display in tooltip [#&#8203;112878](https://redirect.github.com/grafana/grafana/pull/112878), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Fix cell inspect for Sparkline and inferred JSON cells [#&#8203;113059](https://redirect.github.com/grafana/grafana/pull/113059), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **TextPanel:** Fix `CodeEditor` not appearing properly [#&#8203;111937](https://redirect.github.com/grafana/grafana/pull/111937), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **UnitPicker/Cascader:** Fixes type to search for unit feature [#&#8203;112614](https://redirect.github.com/grafana/grafana/pull/112614), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **VizTooltip:** Better overflow handling on long series names [#&#8203;112240](https://redirect.github.com/grafana/grafana/pull/112240), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)

##### Breaking changes

- **Faro:** Update configuration with best practices [#&#8203;112108](https://redirect.github.com/grafana/grafana/pull/112108), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **LibraryPanels:** Remove unique name constraints [#&#8203;113077](https://redirect.github.com/grafana/grafana/pull/113077), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **RBAC:** Only write action sets [#&#8203;112429](https://redirect.github.com/grafana/grafana/pull/112429), [@&#8203;IevaVasiljeva](https://redirect.github.com/IevaVasiljeva)

##### Plugin development fixes & changes

- **Checkbox:** Improve accessibility of the `indeterminate` state [#&#8203;112388](https://redirect.github.com/grafana/grafana/pull/112388), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Collapse:** Improve layout and deprecate `collapsible` prop [#&#8203;113164](https://redirect.github.com/grafana/grafana/pull/113164), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Docs:** Add storybook links to components [#&#8203;113102](https://redirect.github.com/grafana/grafana/pull/113102), [@&#8203;samsch](https://redirect.github.com/samsch)
- **Modal:** Fix button focus being clipped [#&#8203;112867](https://redirect.github.com/grafana/grafana/pull/112867), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Slider:** Expose prop to control visibility of input [#&#8203;113084](https://redirect.github.com/grafana/grafana/pull/113084), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Slider:** Make `inputId` a required param and fix minor a11y violations [#&#8203;112006](https://redirect.github.com/grafana/grafana/pull/112006), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)

<!-- 12.3.0 END -->

<!-- 12.1.4 START -->

### [`v12.2.2`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1222-2025-11-19)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.2.1...v12.2.2)

##### Features and enhancements

- **Access control:** Reduce memory usage when fetching user's permissions [#&#8203;113414](https://redirect.github.com/grafana/grafana/pull/113414), [@&#8203;hairyhenderson](https://redirect.github.com/hairyhenderson)
- **Table:** Pill and JSON Cells should allow formatting [#&#8203;113130](https://redirect.github.com/grafana/grafana/pull/113130), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)

##### Bug fixes

- **AnalyticsSummaries:** Fix dashboard rollup not resetting "last X days" metrics to zero (Enterprise)
- **AnalyticsSummaries:** Fix dashboard rollup totals resetting incorrectly (Enterprise)
- **Security:** fix for CVE-2025-41115 in SCIM (System for Cross-domain Identity Management) (Enterprise)

<!-- 12.2.2 END -->

<!-- 11.6.8 START -->

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on December 01, 2025 at 04:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1946?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`29b134c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/29b134cd85fb4993e8bc7ab6e6d006aedb8ff205?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`7efd06b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7efd06b8e9cf79351709acfd503f0d0357d1cc6a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1946      +/-   ##
==========================================
+ Coverage   58.79%   58.83%   +0.03%     
==========================================
  Files         131      131              
  Lines        8407     8407              
==========================================
+ Hits         4943     4946       +3     
+ Misses       2930     2927       -3     
  Partials      534      534              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1946/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1946/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1946?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1946*
