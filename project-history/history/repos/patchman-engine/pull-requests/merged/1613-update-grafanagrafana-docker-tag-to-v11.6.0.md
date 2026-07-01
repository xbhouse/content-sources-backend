---
type: pull_request
number: 1613
title: "Update grafana/grafana Docker tag to v11.6.0"
state: merged
author: red-hat-konflux
created: 2025-03-26T01:41:01Z
updated: 2026-04-01T03:57:46Z
closed: 2025-03-28T07:45:09Z
merged: 2025-03-28T07:45:09Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1613
---

# Pull Request #1613: Update grafana/grafana Docker tag to v11.6.0

**Author**: @red-hat-konflux
**Created**: March 26, 2025 at 01:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | minor | `11.5.2` -> `11.6.0` |

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v11.6.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1160-2025-03-25)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v11.5.3...v11.6.0)

##### Features and enhancements

-   **API keys:** Migrate API keys to service accounts at startup [#&#8203;96924](https://redirect.github.com/grafana/grafana/pull/96924), [@&#8203;dmihai](https://redirect.github.com/dmihai)
-   **AccessControl:** Allow plugin roles to include `plugins:write` [#&#8203;101089](https://redirect.github.com/grafana/grafana/pull/101089), [@&#8203;gamab](https://redirect.github.com/gamab)
-   **Alerting:** Add DAG errors to alert rule creation and view [#&#8203;99423](https://redirect.github.com/grafana/grafana/pull/99423), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Add Jira integration to cloud AMs [#&#8203;100482](https://redirect.github.com/grafana/grafana/pull/100482), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Add alert rule version history - part1 [#&#8203;99490](https://redirect.github.com/grafana/grafana/pull/99490), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Add migration to clean up rule versions table [#&#8203;102562](https://redirect.github.com/grafana/grafana/pull/102562), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Add multiple threshold operators [#&#8203;99516](https://redirect.github.com/grafana/grafana/pull/99516), [@&#8203;paulojmdias](https://redirect.github.com/paulojmdias)
-   **Alerting:** Add tracking for the mode used in query and notifications step when c… [#&#8203;100824](https://redirect.github.com/grafana/grafana/pull/100824), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Adding color option for slack receiver [#&#8203;99615](https://redirect.github.com/grafana/grafana/pull/99615), [@&#8203;wymangr](https://redirect.github.com/wymangr)
-   **Alerting:** Allow selection of recording rule write target on per-rule basis. [#&#8203;101778](https://redirect.github.com/grafana/grafana/pull/101778), [@&#8203;stevesg](https://redirect.github.com/stevesg)
-   **Alerting:** Allow specifying uid for new rules added to groups [#&#8203;99858](https://redirect.github.com/grafana/grafana/pull/99858), [@&#8203;moustafab](https://redirect.github.com/moustafab)
-   **Alerting:** Improve template testing by trying non-root scopes [#&#8203;101471](https://redirect.github.com/grafana/grafana/pull/101471), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **Alerting:** Include time range in template dashboard and panel urls [#&#8203;101095](https://redirect.github.com/grafana/grafana/pull/101095), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **Alerting:** Keep the latest version of deleted rule in version table [#&#8203;101481](https://redirect.github.com/grafana/grafana/pull/101481), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Promote alertingSaveStateCompressed flag to public preview [#&#8203;99935](https://redirect.github.com/grafana/grafana/pull/99935), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
-   **Alerting:** Remove ID and OrgID from hash calculation [#&#8203;100140](https://redirect.github.com/grafana/grafana/pull/100140), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Remove feature toggle alertingNoNormalState [#&#8203;99905](https://redirect.github.com/grafana/grafana/pull/99905), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Remove rule group edit from single rule editor [#&#8203;100191](https://redirect.github.com/grafana/grafana/pull/100191), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Return 404 when /api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname} does not exist [#&#8203;100264](https://redirect.github.com/grafana/grafana/pull/100264), [@&#8203;fayzal-g](https://redirect.github.com/fayzal-g)
-   **Alerting:** Rule history restore feature [#&#8203;100609](https://redirect.github.com/grafana/grafana/pull/100609), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Support Jira Integration [#&#8203;100480](https://redirect.github.com/grafana/grafana/pull/100480), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Track if new gm rules are created with queries and expressions transformable to simple mode [#&#8203;101121](https://redirect.github.com/grafana/grafana/pull/101121), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Update IRM copies in Configuration Tracker [#&#8203;100069](https://redirect.github.com/grafana/grafana/pull/100069), [@&#8203;teodosii](https://redirect.github.com/teodosii)
-   **Alerting:** Update design of rule details tab and add `updated by` [#&#8203;99895](https://redirect.github.com/grafana/grafana/pull/99895), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
-   **Alerting:** Update irm links for incident and oncall in case new irm plugin is present [#&#8203;99952](https://redirect.github.com/grafana/grafana/pull/99952), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Use exponential backoff in the remote Alertmanager readiness check [#&#8203;99756](https://redirect.github.com/grafana/grafana/pull/99756), [@&#8203;santihernandezc](https://redirect.github.com/santihernandezc)
-   **Alerting:** Use uid instead of id in AnnotationsStateHistory [#&#8203;101207](https://redirect.github.com/grafana/grafana/pull/101207), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Auth:** Add IP address login attempt validation [#&#8203;98123](https://redirect.github.com/grafana/grafana/pull/98123), [@&#8203;colin-stuart](https://redirect.github.com/colin-stuart)
-   **Auth:** Add support for the TlsSkipVerify parameter to JWT Auth [#&#8203;91514](https://redirect.github.com/grafana/grafana/pull/91514), [@&#8203;Ret2Me](https://redirect.github.com/Ret2Me)
-   **Auth:** Make ssoSettingsSAML GA and enabled by default [#&#8203;101766](https://redirect.github.com/grafana/grafana/pull/101766), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
-   **Azure Monitor:** Filter namespaces by resource group [#&#8203;100325](https://redirect.github.com/grafana/grafana/pull/100325), [@&#8203;alyssabull](https://redirect.github.com/alyssabull)
-   **Azure:** Resource picker improvements [#&#8203;101462](https://redirect.github.com/grafana/grafana/pull/101462), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **Azure:** Variable editor and resource picker improvements [#&#8203;101695](https://redirect.github.com/grafana/grafana/pull/101695), [@&#8203;alyssabull](https://redirect.github.com/alyssabull)
-   **Badge:** Add darkgrey color [#&#8203;100699](https://redirect.github.com/grafana/grafana/pull/100699), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
-   **Canvas:** One click links and actions [#&#8203;99616](https://redirect.github.com/grafana/grafana/pull/99616), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
-   **Chore:** Bump Go to 1.23.7 [#&#8203;101576](https://redirect.github.com/grafana/grafana/pull/101576), [@&#8203;macabu](https://redirect.github.com/macabu)
-   **Chore:** Bump Go to 1.23.7 (Enterprise)
-   **Chore:** Bump github.com/expr-lang/expr to v1.17.0 to address CVE-2025-29786 [#&#8203;102533](https://redirect.github.com/grafana/grafana/pull/102533), [@&#8203;macabu](https://redirect.github.com/macabu)
-   **Chore:** Remove `sqlQuerybuilderFunctionParameters` feature toggle [#&#8203;100809](https://redirect.github.com/grafana/grafana/pull/100809), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
-   **CloudWatch:** Track Logs Insights query language [#&#8203;100254](https://redirect.github.com/grafana/grafana/pull/100254), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Configuration tracker:** Update copy in IRM and point to new IRM slack integration [#&#8203;100440](https://redirect.github.com/grafana/grafana/pull/100440), [@&#8203;teodosii](https://redirect.github.com/teodosii)
-   **Dashboard:** Folder move unexpected behavior [#&#8203;100394](https://redirect.github.com/grafana/grafana/pull/100394), [@&#8203;yincongcyincong](https://redirect.github.com/yincongcyincong)
-   **Dashboards:** Allow custom quick time ranges specified in dashboard model [#&#8203;93724](https://redirect.github.com/grafana/grafana/pull/93724), [@&#8203;sknaumov](https://redirect.github.com/sknaumov)
-   **Dashboards:** Monitor dashboard loading performance [#&#8203;99629](https://redirect.github.com/grafana/grafana/pull/99629), [@&#8203;dprokop](https://redirect.github.com/dprokop)
-   **Dashboards:** Remove default empty string from variable create view [#&#8203;98922](https://redirect.github.com/grafana/grafana/pull/98922), [@&#8203;yincongcyincong](https://redirect.github.com/yincongcyincong)
-   **Dashboards:** WeekStart is now of type WeekStart | undefined instead of string [#&#8203;101123](https://redirect.github.com/grafana/grafana/pull/101123), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
-   **DesignSystem:** Menu and popover styling update to use new elevated background token [#&#8203;100255](https://redirect.github.com/grafana/grafana/pull/100255), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **Docker:** Use our own glibc 2.40 binaries [#&#8203;99903](https://redirect.github.com/grafana/grafana/pull/99903), [@&#8203;DanCech](https://redirect.github.com/DanCech)
-   **Docs:** Add a note on query caching for Cloudwatch datasource [#&#8203;100180](https://redirect.github.com/grafana/grafana/pull/100180), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Drilldown:** Require `datasources:explore` RBAC action [#&#8203;101366](https://redirect.github.com/grafana/grafana/pull/101366), [@&#8203;svennergr](https://redirect.github.com/svennergr)
-   **Elasticsearch:** Remove frontend testDatasource method [#&#8203;99894](https://redirect.github.com/grafana/grafana/pull/99894), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Elasticsearch:** Replace level in adhoc filters with level field name [#&#8203;100315](https://redirect.github.com/grafana/grafana/pull/100315), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
-   **Elasticsearch:** Replace term size dropdown with text input [#&#8203;99718](https://redirect.github.com/grafana/grafana/pull/99718), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
-   **Explore:** Add `hide_logs_download` and hide button to download logs [#&#8203;99512](https://redirect.github.com/grafana/grafana/pull/99512), [@&#8203;svennergr](https://redirect.github.com/svennergr)
-   **Explore:** Move drilldown apps from Explore to a new navbar item "Drilldown" [#&#8203;100409](https://redirect.github.com/grafana/grafana/pull/100409), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)
-   **ExploreMetrics:** Add toggle to enable routing to externalized Explore Metrics app plugin [#&#8203;99481](https://redirect.github.com/grafana/grafana/pull/99481), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **Feat:** OSS connections page state filter and update all added [#&#8203;100688](https://redirect.github.com/grafana/grafana/pull/100688), [@&#8203;s4kh](https://redirect.github.com/s4kh)
-   **Features:** Remove openSearchBackendFlowEnabled feature toggle [#&#8203;99068](https://redirect.github.com/grafana/grafana/pull/99068), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
-   **Folders:** Add validation that folder is not a parent of itself [#&#8203;101569](https://redirect.github.com/grafana/grafana/pull/101569), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
-   **Geomap:** WebGL for Marker Layer [#&#8203;95457](https://redirect.github.com/grafana/grafana/pull/95457), [@&#8203;drew08t](https://redirect.github.com/drew08t)
-   **Grafana/ui:** Export UsersIndicator [#&#8203;100698](https://redirect.github.com/grafana/grafana/pull/100698), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
-   **Graphite:** Compare query builder query to raw query [#&#8203;101104](https://redirect.github.com/grafana/grafana/pull/101104), [@&#8203;bossinc](https://redirect.github.com/bossinc)
-   **Histogram:** Handle multiple native histograms [#&#8203;98404](https://redirect.github.com/grafana/grafana/pull/98404), [@&#8203;domasx2](https://redirect.github.com/domasx2)
-   **Image Renderer:** Add support for SSL in plugin mode [#&#8203;98009](https://redirect.github.com/grafana/grafana/pull/98009), [@&#8203;nmarrs](https://redirect.github.com/nmarrs)
-   **ImportDashboards:** Use NestedFolderPicker [#&#8203;99696](https://redirect.github.com/grafana/grafana/pull/99696), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
-   **Loki:** Removal of `Resolution` in query editors [#&#8203;101860](https://redirect.github.com/grafana/grafana/pull/101860), [@&#8203;svennergr](https://redirect.github.com/svennergr)
-   **Menu:** Uniform padding to make menu item hover state look better [#&#8203;100275](https://redirect.github.com/grafana/grafana/pull/100275), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **MetricsDrilldown:** Update name of queryless metrics experience [#&#8203;100675](https://redirect.github.com/grafana/grafana/pull/100675), [@&#8203;yangkb09](https://redirect.github.com/yangkb09)
-   **MultiCombobox:** Export from grafana/ui [#&#8203;100368](https://redirect.github.com/grafana/grafana/pull/100368), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
-   **NodeGraph:** Improve view traces for uninstrumented services [#&#8203;98442](https://redirect.github.com/grafana/grafana/pull/98442), [@&#8203;edvard-falkskar](https://redirect.github.com/edvard-falkskar)
-   **PluginExtensions:** Added support for sharing functions [#&#8203;98888](https://redirect.github.com/grafana/grafana/pull/98888), [@&#8203;theSuess](https://redirect.github.com/theSuess)
-   **PluginExtensions:** Added support for sharing functions (Enterprise)
-   **PluginExtensions:** Exposing registry meta for components returned via `usePluginComponents` [#&#8203;100587](https://redirect.github.com/grafana/grafana/pull/100587), [@&#8203;mckn](https://redirect.github.com/mckn)
-   **Plugins:** Improve plugin details UX for core plugins [#&#8203;99830](https://redirect.github.com/grafana/grafana/pull/99830), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Plugins:** Remove managedPluginsInstall feature toggle [#&#8203;100416](https://redirect.github.com/grafana/grafana/pull/100416), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Plugins:** Remove managedPluginsInstall feature toggle (Enterprise)
-   **Plugins:** Remove uninstall plugin step from cli plugins update-all [#&#8203;101632](https://redirect.github.com/grafana/grafana/pull/101632), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Prometheus:** Get the utcOffset value of timezone when it's specified [#&#8203;99910](https://redirect.github.com/grafana/grafana/pull/99910), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
-   **Prometheus:** Remove query assistant and related components [#&#8203;100669](https://redirect.github.com/grafana/grafana/pull/100669), [@&#8203;edwardcqian](https://redirect.github.com/edwardcqian)
-   **QueryOptions:** Handle invalid time shift values [#&#8203;101670](https://redirect.github.com/grafana/grafana/pull/101670), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
-   **RBAC:** Remove accessControlOnCall feature toggle [#&#8203;101222](https://redirect.github.com/grafana/grafana/pull/101222), [@&#8203;gamab](https://redirect.github.com/gamab)
-   **RBAC:** Remove accessControlOnCall feature toggle (Enterprise)
-   **Reporting:** Add email subject support (Enterprise)
-   **Security:** Update to Go 1.23.5 (Enterprise)
-   **Tempo:** Support TraceQL instant metrics queries [#&#8203;99732](https://redirect.github.com/grafana/grafana/pull/99732), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
-   **Tempo:** TraceQL metrics streaming [#&#8203;99037](https://redirect.github.com/grafana/grafana/pull/99037), [@&#8203;adrapereira](https://redirect.github.com/adrapereira)
-   **Time regions:** Add option for cron syntax to support complex schedules [#&#8203;99548](https://redirect.github.com/grafana/grafana/pull/99548), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
-   **TimePicker:** Ability to manually specify quick ranges [#&#8203;101465](https://redirect.github.com/grafana/grafana/pull/101465), [@&#8203;Sergej-Vlasov](https://redirect.github.com/Sergej-Vlasov)
-   **TimeRangePicker:** Options list padding [#&#8203;100343](https://redirect.github.com/grafana/grafana/pull/100343), [@&#8203;torkelo](https://redirect.github.com/torkelo)
-   **TopNav:** Move news into profile menu [#&#8203;99535](https://redirect.github.com/grafana/grafana/pull/99535), [@&#8203;bergquist](https://redirect.github.com/bergquist)
-   **Trace View:** Add link from the Trace View to the Profiles Drilldown [#&#8203;101422](https://redirect.github.com/grafana/grafana/pull/101422), [@&#8203;joey-grafana](https://redirect.github.com/joey-grafana)
-   **Transformation:** Add support for variables to ALL transformations [#&#8203;100225](https://redirect.github.com/grafana/grafana/pull/100225), [@&#8203;dprokop](https://redirect.github.com/dprokop)
-   **Transformations:** Add round() to Unary mode of `Add field from calc` [#&#8203;101295](https://redirect.github.com/grafana/grafana/pull/101295), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
-   **VizActions:** Add confirmation message [#&#8203;100012](https://redirect.github.com/grafana/grafana/pull/100012), [@&#8203;adela-almasan](https://redirect.github.com/adela-almasan)
-   **grafana-ui:** Update InlineField error prop type to React.ReactNode [#&#8203;100347](https://redirect.github.com/grafana/grafana/pull/100347), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)

##### Bug fixes

-   **Alerting:** Add error handling for missing data source [#&#8203;101508](https://redirect.github.com/grafana/grafana/pull/101508), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Call RLock() before reading sendAlertsTo map [#&#8203;99812](https://redirect.github.com/grafana/grafana/pull/99812), [@&#8203;santihernandezc](https://redirect.github.com/santihernandezc)
-   **Alerting:** Disable create rule menu item from panel when unifiedAlerting is disabled [#&#8203;100701](https://redirect.github.com/grafana/grafana/pull/100701), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Fix KeyValueMap input bug [#&#8203;101367](https://redirect.github.com/grafana/grafana/pull/101367), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Fix crash when invalid matcher is used in silence query params [#&#8203;101500](https://redirect.github.com/grafana/grafana/pull/101500), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Fix evaluation of rules with no-op math expressions [#&#8203;101436](https://redirect.github.com/grafana/grafana/pull/101436), [@&#8203;moustafab](https://redirect.github.com/moustafab)
-   **Alerting:** Fix exporting new rule with a new group [#&#8203;101404](https://redirect.github.com/grafana/grafana/pull/101404), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
-   **Alerting:** Fix fieldSelector encoding [#&#8203;99751](https://redirect.github.com/grafana/grafana/pull/99751), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Fix inheritance of the timing options for policy tree [#&#8203;99398](https://redirect.github.com/grafana/grafana/pull/99398), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Fix notification templates layout [#&#8203;101232](https://redirect.github.com/grafana/grafana/pull/101232), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
-   **Alerting:** Fix state reason [#&#8203;101530](https://redirect.github.com/grafana/grafana/pull/101530), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
-   **Alerting:** Fix token-based Slack image upload to work with channel names [#&#8203;100988](https://redirect.github.com/grafana/grafana/pull/100988), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **App Platform:** Pin bleve to fix CVE-2022-31022 [#&#8203;102531](https://redirect.github.com/grafana/grafana/pull/102531), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
-   **App:** Fix web app behaviour on iOS [#&#8203;100382](https://redirect.github.com/grafana/grafana/pull/100382), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
-   **Auth:** Fix AzureAD config UI's ClientAuthentication dropdown [#&#8203;100752](https://redirect.github.com/grafana/grafana/pull/100752), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
-   **Auth:** Fix redirect with JWT auth URL login [#&#8203;100295](https://redirect.github.com/grafana/grafana/pull/100295), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
-   **AuthN:** Refetch user on "ErrUserAlreadyExists" [#&#8203;100346](https://redirect.github.com/grafana/grafana/pull/100346), [@&#8203;kalleep](https://redirect.github.com/kalleep)
-   **Caching:** Fix duplicate metric registration for cache size (Enterprise)
-   **CloudWatch:** Fix condition for running annotation queries to require dimensions [#&#8203;101660](https://redirect.github.com/grafana/grafana/pull/101660), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
-   **Combobox:** Fix list not being virtualized initially in some cases [#&#8203;100188](https://redirect.github.com/grafana/grafana/pull/100188), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
-   **Dashboard:** Fix for overwriting an edited dashboard in the old architecture [#&#8203;100247](https://redirect.github.com/grafana/grafana/pull/100247), [@&#8203;bfmatei](https://redirect.github.com/bfmatei)
-   **Dashboard:** Fix the unintentional time range and variables updates on saving [#&#8203;101475](https://redirect.github.com/grafana/grafana/pull/101475), [@&#8203;harisrozajac](https://redirect.github.com/harisrozajac)
-   **Dashboard:** Playlist - Fix issue with back button [#&#8203;99401](https://redirect.github.com/grafana/grafana/pull/99401), [@&#8203;yincongcyincong](https://redirect.github.com/yincongcyincong)
-   **DashboardList:** Throttle the re-renders [#&#8203;99982](https://redirect.github.com/grafana/grafana/pull/99982), [@&#8203;bfmatei](https://redirect.github.com/bfmatei)
-   **Dashboards:** Bring back scripted dashboards [#&#8203;100575](https://redirect.github.com/grafana/grafana/pull/100575), [@&#8203;dprokop](https://redirect.github.com/dprokop)
-   **Dashboards:** Fix missing `v/e/i` keybindings to return back to dashboard [#&#8203;102364](https://redirect.github.com/grafana/grafana/pull/102364), [@&#8203;mdvictor](https://redirect.github.com/mdvictor)
-   **Explore:** Fix resizing split view with Loki query editor [#&#8203;100257](https://redirect.github.com/grafana/grafana/pull/100257), [@&#8203;ifrost](https://redirect.github.com/ifrost)
-   **ExploreMetrics:** Fix escaping of regex metacharacters in label filters [#&#8203;100513](https://redirect.github.com/grafana/grafana/pull/100513), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **Fix:** Optimise frontend Postgresql plugin cache busting [#&#8203;100406](https://redirect.github.com/grafana/grafana/pull/100406), [@&#8203;jackw](https://redirect.github.com/jackw)
-   **InfluxDB:** Improve handling of template variables contained in regular expressions (InfluxQL) [#&#8203;100762](https://redirect.github.com/grafana/grafana/pull/100762), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **Interval variable:** Fix $\__auto value behavior [#&#8203;100479](https://redirect.github.com/grafana/grafana/pull/100479), [@&#8203;yincongcyincong](https://redirect.github.com/yincongcyincong)
-   **Log Context:** Fix bug where variables are not replaced in dashboards [#&#8203;100433](https://redirect.github.com/grafana/grafana/pull/100433), [@&#8203;svennergr](https://redirect.github.com/svennergr)
-   **OpenTSDB:** Support v2.4 [#&#8203;100673](https://redirect.github.com/grafana/grafana/pull/100673), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **PDF:** Fix repeating panels when there are less items than maxPerRow (Enterprise)
-   **Plugin Metrics:** Eliminate data race in plugin metrics middleware [#&#8203;99396](https://redirect.github.com/grafana/grafana/pull/99396), [@&#8203;clord](https://redirect.github.com/clord)
-   **Plugins:** Fix update button behavior on downgrade [#&#8203;101048](https://redirect.github.com/grafana/grafana/pull/101048), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **Plugins:** Fix version tab breaking for non semantic version [#&#8203;101225](https://redirect.github.com/grafana/grafana/pull/101225), [@&#8203;oshirohugo](https://redirect.github.com/oshirohugo)
-   **PromLib:** Take AdHoc filters into account when requesting suggestions without label [#&#8203;101555](https://redirect.github.com/grafana/grafana/pull/101555), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
-   **Prometheus:** Fix cursor jump in prometheus code editor [#&#8203;100273](https://redirect.github.com/grafana/grafana/pull/100273), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
-   **Prometheus:** Fix operator handling when making label expressions utf-8 friendly [#&#8203;100475](https://redirect.github.com/grafana/grafana/pull/100475), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
-   **Prometheus:** Fix setting utcOffset when absolute time range is used [#&#8203;101065](https://redirect.github.com/grafana/grafana/pull/101065), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)
-   **RBAC:** Don't check folder access if `annotationPermissionUpdate` FT is enabled [#&#8203;99717](https://redirect.github.com/grafana/grafana/pull/99717), [@&#8203;IevaVasiljeva](https://redirect.github.com/IevaVasiljeva)
-   **SSO:** Fix team_ids validation for Generic OAuth [#&#8203;100732](https://redirect.github.com/grafana/grafana/pull/100732), [@&#8203;dmihai](https://redirect.github.com/dmihai)
-   **Service Accounts:** Don't show error pop-ups for Service Account and Renderer UI flows [#&#8203;101776](https://redirect.github.com/grafana/grafana/pull/101776), [@&#8203;IevaVasiljeva](https://redirect.github.com/IevaVasiljeva)
-   **Share:** Fix short links when root_url is different from the browser URL [#&#8203;99950](https://redirect.github.com/grafana/grafana/pull/99950), [@&#8203;AgnesToulet](https://redirect.github.com/AgnesToulet)

##### Breaking changes

-   **Data source:** Change Permissions for query to only have query and not `read OR query` (Enterprise)

##### Plugin development fixes & changes

-   **GrafanaUI:** Deprecate Select in favor of Combobox [#&#8203;100294](https://redirect.github.com/grafana/grafana/pull/100294), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
-   **Multi/Combobox:** Use pointer cursor when not focused [#&#8203;100878](https://redirect.github.com/grafana/grafana/pull/100878), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
-   **Slider:** Fix text input box being too wide [#&#8203;100138](https://redirect.github.com/grafana/grafana/pull/100138), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)

<!-- 11.6.0 END -->

<!-- 11.5.3 START -->

### [`v11.5.3`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1153-2025-03-25)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v11.5.2...v11.5.3)

##### Features and enhancements

-   **Chore:** Bump Go to 1.23.7 [#&#8203;101581](https://redirect.github.com/grafana/grafana/pull/101581), [@&#8203;macabu](https://redirect.github.com/macabu)
-   **Chore:** Bump Go to 1.23.7 (Enterprise)

##### Bug fixes

-   **Alerting:** Fix token-based Slack image upload to work with channel names [#&#8203;101078](https://redirect.github.com/grafana/grafana/pull/101078), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
-   **Auth:** Fix AzureAD config UI's ClientAuthentication dropdown [#&#8203;100869](https://redirect.github.com/grafana/grafana/pull/100869), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
-   **Dashboard:** Fix the unintentional time range and variables updates on saving [#&#8203;101671](https://redirect.github.com/grafana/grafana/pull/101671), [@&#8203;harisrozajac](https://redirect.github.com/harisrozajac)
-   **Dashboards:** Fix missing `v/e/i` keybindings to return back to dashboard [#&#8203;102365](https://redirect.github.com/grafana/grafana/pull/102365), [@&#8203;mdvictor](https://redirect.github.com/mdvictor)
-   **InfluxDB:** Improve handling of template variables contained in regular expressions (InfluxQL) [#&#8203;100977](https://redirect.github.com/grafana/grafana/pull/100977), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
-   **Org redirection:** Fix linking between orgs [#&#8203;102089](https://redirect.github.com/grafana/grafana/pull/102089), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)

<!-- 11.5.3 END -->

<!-- 11.4.3 START -->

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "before 5am" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on March 26, 2025 at 01:41 AM UTC

Commits missing Jira IDs:
38683ef88007d4dc1fcd2871e099bd8ff3fb83be


### Comment by @codecov-commenter on March 26, 2025 at 01:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1613?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.08%. Comparing base ([`0f26381`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0f263817f5f345fb26a13ddc36a1d12c3d3cf868?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`38683ef`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/38683ef88007d4dc1fcd2871e099bd8ff3fb83be?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 924 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1613   +/-   ##
=======================================
  Coverage   58.08%   58.08%           
=======================================
  Files         145      145           
  Lines       11271    11271           
=======================================
  Hits         6547     6547           
  Misses       4139     4139           
  Partials      585      585           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1613/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1613/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.08% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1613?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1613*
