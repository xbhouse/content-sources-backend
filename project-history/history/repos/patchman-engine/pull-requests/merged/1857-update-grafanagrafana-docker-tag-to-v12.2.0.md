---
type: pull_request
number: 1857
title: "Update grafana/grafana Docker tag to v12.2.0"
state: merged
author: red-hat-konflux
created: 2025-09-24T00:15:20Z
updated: 2025-10-08T16:16:34Z
closed: 2025-09-24T00:23:41Z
merged: 2025-09-24T00:23:41Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1857
---

# Pull Request #1857: Update grafana/grafana Docker tag to v12.2.0

**Author**: @red-hat-konflux
**Created**: September 24, 2025 at 12:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | minor | `12.1.1` -> `12.2.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.2.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1220-2025-09-23)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.1.2...v12.2.0)

##### Features and enhancements

- \*\* Alerting:\*\* Add feedback buttons for the new AI helpers (Enterprise)
- **Access:** Remove plugin app access in plugin basic role seeder (Enterprise)
- **Actions:** Infinity authentication [#&#8203;109493](https://redirect.github.com/grafana/grafana/pull/109493), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
- **Alerting:** Add GMA export to the new list page [#&#8203;109784](https://redirect.github.com/grafana/grafana/pull/109784), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Add alerting AI buttons for cloud (Enterprise)
- **Alerting:** Add contact point filter to Active Notifications page [#&#8203;109775](https://redirect.github.com/grafana/grafana/pull/109775), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Add enrichment per rule extension component (Enterprise)
- **Alerting:** Add extension point link from alert rule to grafana-metricsdrilldown-app [#&#8203;108566](https://redirect.github.com/grafana/grafana/pull/108566), [@&#8203;bohandley](https://redirect.github.com/bohandley)
- **Alerting:** Add feature toggle and extension point [#&#8203;110141](https://redirect.github.com/grafana/grafana/pull/110141), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add keepFiringFor and missing\_series\_evals\_to\_resolve to file provisioning [#&#8203;109699](https://redirect.github.com/grafana/grafana/pull/109699), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Add observability to enrichment UI (Enterprise)
- **Alerting:** Add tooltips in enrichment list for enrichment type (Enterprise)
- **Alerting:** Alert enrichment list page (Enterprise)
- **Alerting:** Allow filter by rule source in Filter V2 [#&#8203;110336](https://redirect.github.com/grafana/grafana/pull/110336), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Auto refresh contact points in the rule form [#&#8203;109539](https://redirect.github.com/grafana/grafana/pull/109539), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Check if TimeInterval is used in ActiveTimings when deleting [#&#8203;110691](https://redirect.github.com/grafana/grafana/pull/110691), [@&#8203;fayzal-g](https://redirect.github.com/fayzal-g)
- **Alerting:** Disable group consistency check for GMA rules [#&#8203;109599](https://redirect.github.com/grafana/grafana/pull/109599), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Display Error Message in Alert History View [#&#8203;110123](https://redirect.github.com/grafana/grafana/pull/110123), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Enrichment Config Form (Enterprise)
- **Alerting:** Filter out private labels before writing recording rules [#&#8203;109295](https://redirect.github.com/grafana/grafana/pull/109295), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** List V2 - Add a group link to the rule list item [#&#8203;108960](https://redirect.github.com/grafana/grafana/pull/108960), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** List V2 - datasource icons for rules [#&#8203;109033](https://redirect.github.com/grafana/grafana/pull/109033), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Load labels in drop-downs without blocking the interaction with the form inputs [#&#8203;110648](https://redirect.github.com/grafana/grafana/pull/110648), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Mark Prometheus to Grafana conversion API as stable [#&#8203;103499](https://redirect.github.com/grafana/grafana/pull/103499), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Move alerting file to an alerting folder [#&#8203;110257](https://redirect.github.com/grafana/grafana/pull/110257), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Support JSON responses in the Prometheus conversion API [#&#8203;109070](https://redirect.github.com/grafana/grafana/pull/109070), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Support extra labels in the Prometheus conversion API [#&#8203;109136](https://redirect.github.com/grafana/grafana/pull/109136), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Support retry with backoff in alert rule evaluation [#&#8203;99710](https://redirect.github.com/grafana/grafana/pull/99710), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Triage alert history with Assistant if available (Enterprise)
- **Auditing:** Add settings to control recording of datasource query request and response body (Enterprise)
- **Auth:** Add setting to disable username based brute force login protection [#&#8203;109152](https://redirect.github.com/grafana/grafana/pull/109152), [@&#8203;TheoBrigitte](https://redirect.github.com/TheoBrigitte)
- **Auth:** Support JWT configs `tls_client_ca` and `jwk_set_bearer_token_file` [#&#8203;109095](https://redirect.github.com/grafana/grafana/pull/109095), [@&#8203;Baarsgaard](https://redirect.github.com/Baarsgaard)
- **Azure:** Resource picker improvements ([#&#8203;109458](https://redirect.github.com/grafana/grafana/issues/109458)) [#&#8203;109520](https://redirect.github.com/grafana/grafana/pull/109520), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Show resource group in picker [#&#8203;110442](https://redirect.github.com/grafana/grafana/pull/110442), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Canvas:** Add option to disable tooltips for one-click elements [#&#8203;109937](https://redirect.github.com/grafana/grafana/pull/109937), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
- **Canvas:** Dynamic connection direction [#&#8203;108423](https://redirect.github.com/grafana/grafana/pull/108423), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
- **Chore:** Remove prometheusCodeModeMetricNamesSearch feature toggle [#&#8203;109024](https://redirect.github.com/grafana/grafana/pull/109024), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Chore:** Removes HideAngularDeprecation configuration [#&#8203;110665](https://redirect.github.com/grafana/grafana/pull/110665), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
- **CloudConfig:** Add config from defaults.ini to StackInfo (Enterprise)
- **CloudWatch:** Append query type to the request id [#&#8203;109068](https://redirect.github.com/grafana/grafana/pull/109068), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **CloudWatch:** Use default region when query region is unset [#&#8203;109089](https://redirect.github.com/grafana/grafana/pull/109089), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **CloudWatch:** Use the correct metric name for errors per function panel in the AWS Lambda sample dashboard [#&#8203;110718](https://redirect.github.com/grafana/grafana/pull/110718), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CommandPalette:** Use fuzzySearch util from grafana/data [#&#8203;108884](https://redirect.github.com/grafana/grafana/pull/108884), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **Dashboard:** Inspect drawer can no longer be opened with url or linked to [#&#8203;109617](https://redirect.github.com/grafana/grafana/pull/109617), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboards:** Add support for full screen panel view and embedded (solo panel) route to repeated panels and new layouts (via new SoloPanelContex) [#&#8203;107375](https://redirect.github.com/grafana/grafana/pull/107375), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboards:** Conserve timestamp on time range copy-paste across timezones [#&#8203;109769](https://redirect.github.com/grafana/grafana/pull/109769), [@&#8203;alik-r](https://redirect.github.com/alik-r)
- **Dashboards:** Enable kubernetesDashboards by default [#&#8203;107618](https://redirect.github.com/grafana/grafana/pull/107618), [@&#8203;dprokop](https://redirect.github.com/dprokop)
- **Dashboards:** Make it possible to render variables under a drop-down [#&#8203;109225](https://redirect.github.com/grafana/grafana/pull/109225), [@&#8203;leventebalogh](https://redirect.github.com/leventebalogh)
- **Database:** Add primary key to Settings table (Enterprise)
- **Database:** Add primary key to settings table (Enterprise)
- **Dependencies:** Bump Go to v1.24.5 (Enterprise)
- **Docs:** Deprecate `grafana/grafana-oss` docker repo in favor of `grafana/grafana` [#&#8203;110065](https://redirect.github.com/grafana/grafana/pull/110065), [@&#8203;kminehart](https://redirect.github.com/kminehart)
- **Flame Graph:** Analyze with Grafana Assistant [#&#8203;108684](https://redirect.github.com/grafana/grafana/pull/108684), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Folders:** Add team folders feature toggle [#&#8203;109389](https://redirect.github.com/grafana/grafana/pull/109389), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Update folder using app platform APIs [#&#8203;110449](https://redirect.github.com/grafana/grafana/pull/110449), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Folders:** Use app platform search endpoint and update tests [#&#8203;108814](https://redirect.github.com/grafana/grafana/pull/108814), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Go:** Update to 1.24.6 [#&#8203;109313](https://redirect.github.com/grafana/grafana/pull/109313), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
- **InfluxDB:** Ad hoc filters support for expressions [#&#8203;109344](https://redirect.github.com/grafana/grafana/pull/109344), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Metrics:** Add http\_response\_size\_bytes metric [#&#8203;110428](https://redirect.github.com/grafana/grafana/pull/110428), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **Nested folders:** Remove feature flag [#&#8203;109212](https://redirect.github.com/grafana/grafana/pull/109212), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **NestedFolderPicker:** Add rootFolderUID prop [#&#8203;109991](https://redirect.github.com/grafana/grafana/pull/109991), [@&#8203;ywzheng1](https://redirect.github.com/ywzheng1)
- **P2P Filter:** Add adhoc filter option toggle [#&#8203;110160](https://redirect.github.com/grafana/grafana/pull/110160), [@&#8203;Develer](https://redirect.github.com/Develer)
- **PieChart:** Add panel options for ascending/descending sort, and no sorting [#&#8203;109564](https://redirect.github.com/grafana/grafana/pull/109564), [@&#8203;cglukas](https://redirect.github.com/cglukas)
- **Plugin Extensions:** DataSource Configuration Components [#&#8203;108350](https://redirect.github.com/grafana/grafana/pull/108350), [@&#8203;shelldandy](https://redirect.github.com/shelldandy)
- **Plugins:** Add Connections homepage [#&#8203;108316](https://redirect.github.com/grafana/grafana/pull/108316), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
- **Plugins:** Record plugin version in request metrics [#&#8203;110210](https://redirect.github.com/grafana/grafana/pull/110210), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **Preferences:** Move codegen to apps [#&#8203;109178](https://redirect.github.com/grafana/grafana/pull/109178), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Prometheus data source:** Migration service [#&#8203;107364](https://redirect.github.com/grafana/grafana/pull/107364), [@&#8203;bossinc](https://redirect.github.com/bossinc)
- **Prometheus:** Refactor metrics modal to handle high cardinality metrics [#&#8203;108437](https://redirect.github.com/grafana/grafana/pull/108437), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Pyroscope:** Process and display sampling annotations [#&#8203;109707](https://redirect.github.com/grafana/grafana/pull/109707), [@&#8203;aleks-p](https://redirect.github.com/aleks-p)
- **Reporting:** Permit valid but weird emails (Enterprise)
- **Reporting:** Show correct recipient count (Enterprise)
- **Revert:** DataSource: Support config CRUD from apiservers ([#&#8203;106996](https://redirect.github.com/grafana/grafana/issues/106996)) [#&#8203;110342](https://redirect.github.com/grafana/grafana/pull/110342), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **Revert:** DataSource: Support config CRUD from apiservers ([#&#8203;8860](https://redirect.github.com/grafana/grafana/issues/8860)) (Enterprise)
- **SCIM:** Add flag for rejecting non provisioned users from logging in (Enterprise)
- **SCIM:** Allow empty externalId on update operation (Enterprise)
- **SCIM:** Delete user instead of disabling it on SCIM DELETE user request (Enterprise)
- **SQL Expressions:** Switch feature toggle to public preview [#&#8203;110473](https://redirect.github.com/grafana/grafana/pull/110473), [@&#8203;kylebrandt](https://redirect.github.com/kylebrandt)
- **Table:** Frozen columns [#&#8203;109276](https://redirect.github.com/grafana/grafana/pull/109276), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Max row height for variable height rows [#&#8203;109639](https://redirect.github.com/grafana/grafana/pull/109639), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Tooltip from Field [#&#8203;109428](https://redirect.github.com/grafana/grafana/pull/109428), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Table:** Update UX for uniform-reducer case in new footer and overflow [#&#8203;110493](https://redirect.github.com/grafana/grafana/pull/110493), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **TableNG:** Footer enhancements [#&#8203;102948](https://redirect.github.com/grafana/grafana/pull/102948), [@&#8203;alexjonspencer1](https://redirect.github.com/alexjonspencer1)
- **Text:** Add Inter italic font variants to Grafana UI [#&#8203;110313](https://redirect.github.com/grafana/grafana/pull/110313), [@&#8203;kapowaz](https://redirect.github.com/kapowaz)
- **TraceView:** Refine UI visual hierarchy inside details section [#&#8203;108929](https://redirect.github.com/grafana/grafana/pull/108929), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Transformations:** Add empty values options to Transpose [#&#8203;108421](https://redirect.github.com/grafana/grafana/pull/108421), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Trend/TimeSeries:** Add "Show values" option [#&#8203;108090](https://redirect.github.com/grafana/grafana/pull/108090), [@&#8203;HasithDeAlwis](https://redirect.github.com/HasithDeAlwis)
- **Trend:** Add support for a logarithmic x axis [#&#8203;101433](https://redirect.github.com/grafana/grafana/pull/101433), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Variables:** shows warning when user tries to save erroneous variables [#&#8203;110154](https://redirect.github.com/grafana/grafana/pull/110154), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
- **VizTooltip:** Replace `ExemplarHoverView` with `VizTooltip` components [#&#8203;109369](https://redirect.github.com/grafana/grafana/pull/109369), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)

##### Bug fixes

- **Alerting:** Fix bug where rules with identical mute/active intervals produced conflicting routes [#&#8203;110971](https://redirect.github.com/grafana/grafana/pull/110971), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix copying of recording rule fields [#&#8203;110311](https://redirect.github.com/grafana/grafana/pull/110311), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Fix field names on webhook HMAC/TLS config HCL export [#&#8203;110722](https://redirect.github.com/grafana/grafana/pull/110722), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Fix newly created alert rules not immediately showing up in folder view [#&#8203;109584](https://redirect.github.com/grafana/grafana/pull/109584), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **Alerting:** Fix permission checks for the Import to GMA [#&#8203;109950](https://redirect.github.com/grafana/grafana/pull/109950), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix permissions for enrichment routes (Enterprise)
- **Alerting:** Fix subpath handling in the alerting package [#&#8203;109448](https://redirect.github.com/grafana/grafana/pull/109448), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix wrong import (Enterprise)
- **Alerting:** Hide list view loader if we don't have anything yet [#&#8203;110464](https://redirect.github.com/grafana/grafana/pull/110464), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Alerting:** Set dataSourceName to GRAFANA\_RULES\_SOURCE\_NAME when switch… [#&#8203;109900](https://redirect.github.com/grafana/grafana/pull/109900), [@&#8203;laurenashleigh](https://redirect.github.com/laurenashleigh)
- **Alerting:** Update alerting module to [`1091588`](https://redirect.github.com/grafana/grafana/commit/10915888e4f099586ad37bea5f4a70f45101d2f5) [#&#8203;109989](https://redirect.github.com/grafana/grafana/pull/109989), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Azure:** Fix logs editor rendering [#&#8203;109491](https://redirect.github.com/grafana/grafana/pull/109491), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Canvas:** Fix element selection being cleared on panel resize [#&#8203;110010](https://redirect.github.com/grafana/grafana/pull/110010), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
- **CloudConfig:** Fix panic in defaults.ini merge (Enterprise)
- **CloudWatch:** Fix handling region for legacy alerts [#&#8203;109217](https://redirect.github.com/grafana/grafana/pull/109217), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **CloudWatch:** Fix logs query requestId to prevent setting undefined-logs as a requestId [#&#8203;109930](https://redirect.github.com/grafana/grafana/pull/109930), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **CloudWatch:** Update grafana/aws-sdk-go with STS endpoint bugfix [#&#8203;109120](https://redirect.github.com/grafana/grafana/pull/109120), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Config:** Fix date\_formats options being moved to a different section [#&#8203;109339](https://redirect.github.com/grafana/grafana/pull/109339), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **Dashboard List:** Fix how link query part is created when variables are included [#&#8203;109861](https://redirect.github.com/grafana/grafana/pull/109861), [@&#8203;aocenas](https://redirect.github.com/aocenas)
- **Dashboard versions:** Fix list for large dashboards [#&#8203;109433](https://redirect.github.com/grafana/grafana/pull/109433), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboard:** Fix AngularJS deprecation in grafana-overview dashboard [#&#8203;106462](https://redirect.github.com/grafana/grafana/pull/106462), [@&#8203;schoen2](https://redirect.github.com/schoen2)
- **Dashboard:** Fixes url links to embedded panels in scene based dashboards [#&#8203;109837](https://redirect.github.com/grafana/grafana/pull/109837), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboards:** Fix UTF-8 characters not working with excel downloads by replacing download for excel with excel compatibility mode. [#&#8203;110099](https://redirect.github.com/grafana/grafana/pull/110099), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
- **Dashboards:** Fix issue where the time range picker would seemingly be hidden behind the side menu if it was set to always open. [#&#8203;108607](https://redirect.github.com/grafana/grafana/pull/108607), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
- **Dashboards:** Fix kiosk mode not persisting through refresh [#&#8203;110284](https://redirect.github.com/grafana/grafana/pull/110284), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
- **Dashboards:** Fixing saving and viewing snapshots for repeated panels [#&#8203;109856](https://redirect.github.com/grafana/grafana/pull/109856), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Explore:** Fix units overflow for trace durations [#&#8203;108515](https://redirect.github.com/grafana/grafana/pull/108515), [@&#8203;martincostello](https://redirect.github.com/martincostello)
- **Fix:** Install plugins when they have no plugin archive info(catalog en… [#&#8203;109200](https://redirect.github.com/grafana/grafana/pull/109200), [@&#8203;s4kh](https://redirect.github.com/s4kh)
- **InfluxDB:** Fix Unable to use self-signed CA for adding influxdb data source [#&#8203;105586](https://redirect.github.com/grafana/grafana/pull/105586), [@&#8203;geekeryy](https://redirect.github.com/geekeryy)
- **Prometheus:** Don't use incremental querying if one of the queries has $\_\_range variable [#&#8203;108823](https://redirect.github.com/grafana/grafana/pull/108823), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Prometheus:** Fix eager auto completion [#&#8203;109128](https://redirect.github.com/grafana/grafana/pull/109128), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
- **Prometheus:** QueryEditor fix error when switching from code to builder for undefined aggregation operations [#&#8203;110179](https://redirect.github.com/grafana/grafana/pull/110179), [@&#8203;jcolladokuri](https://redirect.github.com/jcolladokuri)
- **Pyroscope:** Add start and end date to profiletypes call [#&#8203;110277](https://redirect.github.com/grafana/grafana/pull/110277), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
- **Pyroscope:** Fix incorrect rate calculation from flamegraph totals [#&#8203;110470](https://redirect.github.com/grafana/grafana/pull/110470), [@&#8203;marcsanmi](https://redirect.github.com/marcsanmi)
- **Service Accounts:** Fix typo on page indicating none are present [#&#8203;109560](https://redirect.github.com/grafana/grafana/pull/109560), [@&#8203;eamonryan](https://redirect.github.com/eamonryan)
- **Tempo:** Fix instant query streaming [#&#8203;108924](https://redirect.github.com/grafana/grafana/pull/108924), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)
- **TimeSeries:** Use exported time shift and fix time comparison tooltip [#&#8203;109947](https://redirect.github.com/grafana/grafana/pull/109947), [@&#8203;drew08t](https://redirect.github.com/drew08t)
- **Transformations:** Account for group by / count when assessing if calculation is needed [#&#8203;110546](https://redirect.github.com/grafana/grafana/pull/110546), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Transforms:** GroupToMatrix transform should retain keyRowField config [#&#8203;109066](https://redirect.github.com/grafana/grafana/pull/109066), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)

##### Breaking changes

- **Alerting:** Enable alertingSaveStateCompressed by default [#&#8203;109390](https://redirect.github.com/grafana/grafana/pull/109390), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Dashboards:** Repeating with no clone keys [#&#8203;109839](https://redirect.github.com/grafana/grafana/pull/109839), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Provisioning:** Use inline secrets for gitsync [#&#8203;109908](https://redirect.github.com/grafana/grafana/pull/109908), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Stars:** Remove deprecated internal ID apis [#&#8203;110499](https://redirect.github.com/grafana/grafana/pull/110499), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)

##### Plugin development fixes & changes

- **Drawer:** Truncate Drawer title to just one line [#&#8203;109540](https://redirect.github.com/grafana/grafana/pull/109540), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **Modal:** Center modals at smaller screen heights [#&#8203;109256](https://redirect.github.com/grafana/grafana/pull/109256), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **MultiCombobox:** Fix async options to being able to be removed [#&#8203;109473](https://redirect.github.com/grafana/grafana/pull/109473), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **MultiCombobox:** Fix select all when only a single option is available [#&#8203;109910](https://redirect.github.com/grafana/grafana/pull/109910), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)

<!-- 12.2.0 END -->

<!-- 12.1.2 START -->

### [`v12.1.2`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1212-2025-09-23)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.1.1...v12.1.2)

##### Features and enhancements

- **Alerting:** Update alerting module [#&#8203;109999](https://redirect.github.com/grafana/grafana/pull/109999), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Auditing:** Add settings to control recording of datasource query request and response body (Enterprise)
- **Auditing:** Document new options for recording datasource query request/response body [#&#8203;109981](https://redirect.github.com/grafana/grafana/pull/109981), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Chore:** Don't show a "Not found" for public-dashboard fetches if the service is disabled via config [#&#8203;110144](https://redirect.github.com/grafana/grafana/pull/110144), [@&#8203;mmandrus](https://redirect.github.com/mmandrus)
- **CloudWatch:** Use default region when query region is unset [#&#8203;111079](https://redirect.github.com/grafana/grafana/pull/111079), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)

##### Bug fixes

- **Alerting:** Fix bug where rules with identical mute/active intervals produced conflicting routes [#&#8203;110973](https://redirect.github.com/grafana/grafana/pull/110973), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix copying of recording rule fields [#&#8203;110312](https://redirect.github.com/grafana/grafana/pull/110312), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Fix:** Fix redirection after login when Grafana is served from subpath [#&#8203;111097](https://redirect.github.com/grafana/grafana/pull/111097), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)

##### Plugin development fixes & changes

- **Fix:** Prevent Rollup from treeshaking NPM packages [#&#8203;108570](https://redirect.github.com/grafana/grafana/pull/108570), [@&#8203;jackw](https://redirect.github.com/jackw)

<!-- 12.1.2 END -->

<!-- 12.0.5 START -->

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "before 5am" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 24, 2025 at 12:15 AM UTC

Commits missing Jira IDs:
8836a4b92533cfe606416f2d1aafbf3355b1fe1e


### Comment by @codecov-commenter on September 24, 2025 at 12:20 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1857?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.53%. Comparing base ([`0c6f0d7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0c6f0d75676909b09677b3878454498b364858bf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`8836a4b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8836a4b92533cfe606416f2d1aafbf3355b1fe1e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1857   +/-   ##
=======================================
  Coverage   54.53%   54.53%           
=======================================
  Files         138      138           
  Lines       10743    10743           
=======================================
  Hits         5859     5859           
  Misses       4346     4346           
  Partials      538      538           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1857/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1857/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.53% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1857?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1857*
