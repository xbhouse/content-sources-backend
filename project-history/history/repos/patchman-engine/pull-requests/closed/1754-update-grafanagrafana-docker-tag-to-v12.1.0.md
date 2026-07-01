---
type: pull_request
number: 1754
title: "Update grafana/grafana Docker tag to v12.1.0"
state: closed
author: red-hat-konflux
created: 2025-07-28T01:46:47Z
updated: 2025-07-28T09:32:52Z
closed: 2025-07-28T08:41:39Z
base_branch: security-compliance
head_branch: konflux/mintmaker/security-compliance/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1754
---

# Pull Request #1754: Update grafana/grafana Docker tag to v12.1.0

**Author**: @red-hat-konflux
**Created**: July 28, 2025 at 01:46 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `security-compliance` ← **Head**: `konflux/mintmaker/security-compliance/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | minor | `12.0.2` -> `12.1.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.1.0`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1210-2025-07-23)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.0.3...v12.1.0)

##### Features and enhancements

- **Access:** Disable role none option if advanced access control is not enabled [#&#8203;107378](https://redirect.github.com/grafana/grafana/pull/107378), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **Alerting:** Add OAuth2 Support for Webhook Receiver [#&#8203;106302](https://redirect.github.com/grafana/grafana/pull/106302), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Add ability to import rules to GMA from Prometheus YAML [#&#8203;105807](https://redirect.github.com/grafana/grafana/pull/105807), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add details to the payload when tracking import to GMA [#&#8203;106404](https://redirect.github.com/grafana/grafana/pull/106404), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add export folder action to the new list view [#&#8203;106256](https://redirect.github.com/grafana/grafana/pull/106256), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add filters for health and contact point in Prometheus Rules api [#&#8203;106580](https://redirect.github.com/grafana/grafana/pull/106580), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Add loading spinner for loading groups state [#&#8203;106289](https://redirect.github.com/grafana/grafana/pull/106289), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add need more info for import ui datasource field [#&#8203;106364](https://redirect.github.com/grafana/grafana/pull/106364), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Add provenance to Prometheus API [#&#8203;106596](https://redirect.github.com/grafana/grafana/pull/106596), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Add provenance to remote-ruler extension response (Enterprise)
- **Alerting:** Add simplified routing metadata to the details tab [#&#8203;106403](https://redirect.github.com/grafana/grafana/pull/106403), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Alerting:** Add state history backend to write ALERTS metric [#&#8203;104361](https://redirect.github.com/grafana/grafana/pull/104361), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Add support for Redis Sentinel for Alerting HA [#&#8203;106322](https://redirect.github.com/grafana/grafana/pull/106322), [@&#8203;vstpme](https://redirect.github.com/vstpme)
- **Alerting:** Allow disabling recording rules write for a data source in the UI [#&#8203;106664](https://redirect.github.com/grafana/grafana/pull/106664), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Correctly persist FiredAt in SyncRuleStatePersister [#&#8203;106658](https://redirect.github.com/grafana/grafana/pull/106658), [@&#8203;fayzal-g](https://redirect.github.com/fayzal-g)
- **Alerting:** Ensure errors cleared when Alerting after error [#&#8203;105246](https://redirect.github.com/grafana/grafana/pull/105246), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Evaluate all imported from Prometheus rules sequentially [#&#8203;106295](https://redirect.github.com/grafana/grafana/pull/106295), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Extensible Settings module [#&#8203;107831](https://redirect.github.com/grafana/grafana/pull/107831), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Filter out rules managed by integrations and add an info alert [#&#8203;106602](https://redirect.github.com/grafana/grafana/pull/106602), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Filter out synthetic datasource-managed rules when importing to GMA [#&#8203;106358](https://redirect.github.com/grafana/grafana/pull/106358), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** List V2 - Add labels popup [#&#8203;107193](https://redirect.github.com/grafana/grafana/pull/107193), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** List V2 - Grouped view filters [#&#8203;106400](https://redirect.github.com/grafana/grafana/pull/106400), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** List V2 - Use backend filters for GMA rules [#&#8203;106897](https://redirect.github.com/grafana/grafana/pull/106897), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Make paginated rules endpoint strongly consistent (Enterprise)
- **Alerting:** Optimize out unnecessary permission check for rule groups (Enterprise)
- **Alerting:** Optimize prometheus api permission checks [#&#8203;106299](https://redirect.github.com/grafana/grafana/pull/106299), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Optimize prometheus api permission checks (Enterprise)
- **Alerting:** Persist alert instance FiredAt field [#&#8203;105927](https://redirect.github.com/grafana/grafana/pull/105927), [@&#8203;fayzal-g](https://redirect.github.com/fayzal-g)
- **Alerting:** Remove ruler from alert list view2 [#&#8203;106778](https://redirect.github.com/grafana/grafana/pull/106778), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Resend alerts for states that are missing in the eval results [#&#8203;105965](https://redirect.github.com/grafana/grafana/pull/105965), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Send notifications immediately on Error|NoData -> Normal transitions [#&#8203;106421](https://redirect.github.com/grafana/grafana/pull/106421), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Support PDC in Grafana-managed recording rules [#&#8203;106677](https://redirect.github.com/grafana/grafana/pull/106677), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Use default\_datasource\_uid as the default target for recording rules in UI [#&#8203;106415](https://redirect.github.com/grafana/grafana/pull/106415), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Annotations:** Use dashboard uids instead of dashboard ids [#&#8203;106676](https://redirect.github.com/grafana/grafana/pull/106676), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **App Platform Provisioning:** Add experimental nanogit mode for Git Sync [#&#8203;106763](https://redirect.github.com/grafana/grafana/pull/106763), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Auth:** Add Azure/Entra workload identity support [#&#8203;104807](https://redirect.github.com/grafana/grafana/pull/104807), [@&#8203;mehighlow](https://redirect.github.com/mehighlow)
- **Auth:** Enable improved session handling by default for OAuth and SAML [#&#8203;107442](https://redirect.github.com/grafana/grafana/pull/107442), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
- **Auth:** Enable ssoSettingsLDAP by default [#&#8203;106310](https://redirect.github.com/grafana/grafana/pull/106310), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
- **Auth:** Remove api key endpoints [#&#8203;106019](https://redirect.github.com/grafana/grafana/pull/106019), [@&#8203;dmihai](https://redirect.github.com/dmihai)
- **Auth:** Remove code for authenticating API keys [#&#8203;105998](https://redirect.github.com/grafana/grafana/pull/105998), [@&#8203;dmihai](https://redirect.github.com/dmihai)
- **Azure:** Support scope selection in Resource Graph queries [#&#8203;105835](https://redirect.github.com/grafana/grafana/pull/105835), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Betterer:** Only allow singleton Storage use [#&#8203;105310](https://redirect.github.com/grafana/grafana/pull/105310), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
- **Caching:** Remove memcached reconnect\_interval setting (Enterprise)
- **Chore:** Update k8s.io to v0.33.1 [#&#8203;105307](https://redirect.github.com/grafana/grafana/pull/105307), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Cloud Monitoring:** Add support for service account impersonation [#&#8203;107022](https://redirect.github.com/grafana/grafana/pull/107022), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
- **CloudMigrations:** Add Mute Timings as dependency for Notification Policies [#&#8203;106751](https://redirect.github.com/grafana/grafana/pull/106751), [@&#8203;macabu](https://redirect.github.com/macabu)
- **CloudWatch:** Backport aws-sdk-go-v2 update from external plugin [#&#8203;107136](https://redirect.github.com/grafana/grafana/pull/107136), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **CloudWatch:** Improve instance attribute variable query editor [#&#8203;105206](https://redirect.github.com/grafana/grafana/pull/105206), [@&#8203;iwysiu](https://redirect.github.com/iwysiu)
- **Cloudwatch:** Add missing AWS regions [#&#8203;106304](https://redirect.github.com/grafana/grafana/pull/106304), [@&#8203;chriscerie](https://redirect.github.com/chriscerie)
- **Dashboard Provisioning:** Reduce db load [#&#8203;106114](https://redirect.github.com/grafana/grafana/pull/106114), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Dashboard:** Add Alert icon in library panels [#&#8203;107723](https://redirect.github.com/grafana/grafana/pull/107723), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Dashboard:** Add server-configurable quick ranges for the time picker [#&#8203;102254](https://redirect.github.com/grafana/grafana/pull/102254), [@&#8203;chodges15](https://redirect.github.com/chodges15)
- **Dashboard:** Formatting Currency - add new custom 'financial' currency format without abbreviations [#&#8203;106604](https://redirect.github.com/grafana/grafana/pull/106604), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Dashboard:** Library Panels - Add ability to search by folder name [#&#8203;106997](https://redirect.github.com/grafana/grafana/pull/106997), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Dashboard:** Schema V2 - Auto-transform V2 dashboards in V1Resource export mode [#&#8203;105997](https://redirect.github.com/grafana/grafana/pull/105997), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Datasources:** Migrate to new sigv4 middleware (Enterprise)
- **Datasources:** Update grafana-aws-sdk for new sigv4 middleware and aws-sdk-go v1 removal [#&#8203;107522](https://redirect.github.com/grafana/grafana/pull/107522), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **DatePicker:** Add cursor not-allowed style and hover background color [#&#8203;106451](https://redirect.github.com/grafana/grafana/pull/106451), [@&#8203;ywzheng1](https://redirect.github.com/ywzheng1)
- **Dependencies:** Bump Go to v1.24.4 [#&#8203;106533](https://redirect.github.com/grafana/grafana/pull/106533), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Dependencies:** Bump github.com/go-viper/mapstructure/v2 from 2.2.1 to 2.3.0 [#&#8203;107379](https://redirect.github.com/grafana/grafana/pull/107379), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Dependencies:** Bump github.com/openfga/openfga to v1.8.13 to address CVE-2025-48371 [#&#8203;106064](https://redirect.github.com/grafana/grafana/pull/106064), [@&#8203;macabu](https://redirect.github.com/macabu)
- **ElasticSearch:** Remove frontend response parsing [#&#8203;104148](https://redirect.github.com/grafana/grafana/pull/104148), [@&#8203;nojaf](https://redirect.github.com/nojaf)
- **Geomap:** Add HiDPI support to CARTO basemap ([#&#8203;81195](https://redirect.github.com/grafana/grafana/issues/81195)) [#&#8203;106211](https://redirect.github.com/grafana/grafana/pull/106211), [@&#8203;cledwynl](https://redirect.github.com/cledwynl)
- **Git Sync UI:** Delete Provisioned Dashboard Flow [#&#8203;106593](https://redirect.github.com/grafana/grafana/pull/106593), [@&#8203;ywzheng1](https://redirect.github.com/ywzheng1)
- **Grafana/data:** Extract fuzzy search core [#&#8203;107110](https://redirect.github.com/grafana/grafana/pull/107110), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **I18n:** Update eslint rule to catch some untranslated object properties [#&#8203;105072](https://redirect.github.com/grafana/grafana/pull/105072), [@&#8203;tomratcliffe](https://redirect.github.com/tomratcliffe)
- **InfluxDB:** Add an optional time range filter for tag queries in the query panel autocompleteInflux tag filter [#&#8203;107195](https://redirect.github.com/grafana/grafana/pull/107195), [@&#8203;NikolayTsvetkov](https://redirect.github.com/NikolayTsvetkov)
- **LBAC for data sources:** Adds team filtering for lbac rules (Enterprise)
- **Library Panels:** Mark library panel RBAC as GA & enable by default [#&#8203;106833](https://redirect.github.com/grafana/grafana/pull/106833), [@&#8203;kaydelaney](https://redirect.github.com/kaydelaney)
- **Library Panels:** Modify connection api endpoint to be compatible with unified storage [#&#8203;107088](https://redirect.github.com/grafana/grafana/pull/107088), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Library elements:** Remove ability to set as "library variable" [#&#8203;106594](https://redirect.github.com/grafana/grafana/pull/106594), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Library panels:** Remove `libraryPanelRBAC` feature flag, and enable rbac by default [#&#8203;107222](https://redirect.github.com/grafana/grafana/pull/107222), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Loki:** Remove experimental lokiQuerySplittingConfig [#&#8203;107298](https://redirect.github.com/grafana/grafana/pull/107298), [@&#8203;ivanahuckova](https://redirect.github.com/ivanahuckova)
- **Loki:** Remove experimental predefined operations [#&#8203;107289](https://redirect.github.com/grafana/grafana/pull/107289), [@&#8203;ivanahuckova](https://redirect.github.com/ivanahuckova)
- **OAuth:** Add access token as third source for user info extraction [#&#8203;107636](https://redirect.github.com/grafana/grafana/pull/107636), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **Plugin Extensions:** Expose PluginMeta generic in usePluginContext [#&#8203;107577](https://redirect.github.com/grafana/grafana/pull/107577), [@&#8203;MattIPv4](https://redirect.github.com/MattIPv4)
- **Postgres:** Switch the datasource plugin from lib/pq to pgx [#&#8203;103961](https://redirect.github.com/grafana/grafana/pull/103961), [@&#8203;zoltanbedi](https://redirect.github.com/zoltanbedi)
- **Preferences:** Use dashboard uid for the home dashboard [#&#8203;106666](https://redirect.github.com/grafana/grafana/pull/106666), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Profiles:** Stop passing response headers for Grafana-Pyroscope and parca datasources [#&#8203;106577](https://redirect.github.com/grafana/grafana/pull/106577), [@&#8203;simonswine](https://redirect.github.com/simonswine)
- **Prometheus:** Deprecation message for Azure auth [#&#8203;106490](https://redirect.github.com/grafana/grafana/pull/106490), [@&#8203;bossinc](https://redirect.github.com/bossinc)
- **Prometheus:** Facilitate tree shaking with exports and bundler mode [#&#8203;105575](https://redirect.github.com/grafana/grafana/pull/105575), [@&#8203;NWRichmond](https://redirect.github.com/NWRichmond)
- **Prometheus:** Migrate remaining selectors to data-testid [#&#8203;106564](https://redirect.github.com/grafana/grafana/pull/106564), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **ProvisionedFolder:** Delete folder drawer [#&#8203;107089](https://redirect.github.com/grafana/grafana/pull/107089), [@&#8203;ywzheng1](https://redirect.github.com/ywzheng1)
- **Provisioning:** Add pure git repository type [#&#8203;106815](https://redirect.github.com/grafana/grafana/pull/106815), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Querying:** Pass dashboard and panel title as headers [#&#8203;107032](https://redirect.github.com/grafana/grafana/pull/107032), [@&#8203;ivanahuckova](https://redirect.github.com/ivanahuckova)
- **Remote Alertmanager:** Send SMTP config [#&#8203;106337](https://redirect.github.com/grafana/grafana/pull/106337), [@&#8203;santihernandezc](https://redirect.github.com/santihernandezc)
- **Restore dashboards:** Add filters and search [#&#8203;106994](https://redirect.github.com/grafana/grafana/pull/106994), [@&#8203;Clarity-89](https://redirect.github.com/Clarity-89)
- **SCIM:** Ignore unsupported fields in user PATCH requests (Enterprise)
- **SCIM:** Implement operation for adding an externalId value to a team (Enterprise)
- **SCIM:** Implement the add members operation in group PATCH requests (Enterprise)
- **SCIM:** Implement the remove members operation in group PATCH requests (Enterprise)
- **SCIM:** Update externalId field in group PATCH request (Enterprise)
- **SQL Expressions:** Always convert on type first [#&#8203;106083](https://redirect.github.com/grafana/grafana/pull/106083), [@&#8203;kylebrandt](https://redirect.github.com/kylebrandt)
- **Select:** Set min width for the current selected item when width=auto [#&#8203;106131](https://redirect.github.com/grafana/grafana/pull/106131), [@&#8203;tskarhed](https://redirect.github.com/tskarhed)
- **StateTimeline:** Display false and empty string values [#&#8203;107059](https://redirect.github.com/grafana/grafana/pull/107059), [@&#8203;jesdavpet](https://redirect.github.com/jesdavpet)
- **StateTimeline:** Support `NaN` and `null` value mappings [#&#8203;105638](https://redirect.github.com/grafana/grafana/pull/105638), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **Storage:** Take `migration_locking` setting into account [#&#8203;105938](https://redirect.github.com/grafana/grafana/pull/105938), [@&#8203;JohnnyQQQQ](https://redirect.github.com/JohnnyQQQQ)
- **TableNG:** Refactor to better take advantage of react-data-grid [#&#8203;103755](https://redirect.github.com/grafana/grafana/pull/103755), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)
- **Tables:** Pills for Table Cells [#&#8203;107485](https://redirect.github.com/grafana/grafana/pull/107485), [@&#8203;timlevett](https://redirect.github.com/timlevett)
- **Teams:** Add support for updating externalId field [#&#8203;106406](https://redirect.github.com/grafana/grafana/pull/106406), [@&#8203;dmihai](https://redirect.github.com/dmihai)
- **Tempo:** Enable native histograms for Tempo service graph [#&#8203;105989](https://redirect.github.com/grafana/grafana/pull/105989), [@&#8203;bohandley](https://redirect.github.com/bohandley)
- **TimeRangePicker:** Highlight range on hover [#&#8203;106616](https://redirect.github.com/grafana/grafana/pull/106616), [@&#8203;joshhunt](https://redirect.github.com/joshhunt)
- **TraceView:** Resource attributes links extension point [#&#8203;104680](https://redirect.github.com/grafana/grafana/pull/104680), [@&#8203;edvard-falkskar](https://redirect.github.com/edvard-falkskar)
- **Transformations:** Add "Auto" mode to Organize Fields [#&#8203;103055](https://redirect.github.com/grafana/grafana/pull/103055), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Transformations:** GA the Regression transformation [#&#8203;106074](https://redirect.github.com/grafana/grafana/pull/106074), [@&#8203;gelicia](https://redirect.github.com/gelicia)
- **Unified storage:** Respect GF\_DATABASE\_URL override [#&#8203;105331](https://redirect.github.com/grafana/grafana/pull/105331), [@&#8203;pstibrany](https://redirect.github.com/pstibrany)
- **VQB:** Add selected columns to GROUP BY dropdown ([#&#8203;106349](https://redirect.github.com/grafana/grafana/issues/106349)) [#&#8203;106391](https://redirect.github.com/grafana/grafana/pull/106391), [@&#8203;Shubham19032004](https://redirect.github.com/Shubham19032004)
- **VQB:** Allow custom table names in TableSelector [#&#8203;106420](https://redirect.github.com/grafana/grafana/pull/106420), [@&#8203;Victorthedev](https://redirect.github.com/Victorthedev)
- **XYChart:** Add support for x=time [#&#8203;106459](https://redirect.github.com/grafana/grafana/pull/106459), [@&#8203;leeoniya](https://redirect.github.com/leeoniya)

##### Bug fixes

- **Alerting:** Fix $value type when single data source is queried [#&#8203;106080](https://redirect.github.com/grafana/grafana/pull/106080), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix ImportToGMARules flaky test [#&#8203;106495](https://redirect.github.com/grafana/grafana/pull/106495), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix RefIds not being shown when creating or editing Grafana-managed recording rule [#&#8203;106840](https://redirect.github.com/grafana/grafana/pull/106840), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)
- **Alerting:** Fix contact points tab visibility when user can only create [#&#8203;106735](https://redirect.github.com/grafana/grafana/pull/106735), [@&#8203;JacobsonMT](https://redirect.github.com/JacobsonMT)
- **Alerting:** Fix eval time unit in list view [#&#8203;106488](https://redirect.github.com/grafana/grafana/pull/106488), [@&#8203;ebuildy](https://redirect.github.com/ebuildy)
- **Alerting:** Fix group interval override when adding new rules [#&#8203;107324](https://redirect.github.com/grafana/grafana/pull/107324), [@&#8203;konrad147](https://redirect.github.com/konrad147)
- **Alerting:** Fix group-level labels and query\_offset in the import API [#&#8203;106379](https://redirect.github.com/grafana/grafana/pull/106379), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix notification policy conflicts originating from provenance mismatch [#&#8203;107343](https://redirect.github.com/grafana/grafana/pull/107343), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Alerting:** Fix resolved notifications for same-label Error to Normal transitions [#&#8203;106210](https://redirect.github.com/grafana/grafana/pull/106210), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Hide labels section if we only have private labels [#&#8203;105996](https://redirect.github.com/grafana/grafana/pull/105996), [@&#8203;gillesdemey](https://redirect.github.com/gillesdemey)
- **Annotations:** Remove prometheus from legacy runner [#&#8203;106737](https://redirect.github.com/grafana/grafana/pull/106737), [@&#8203;scottlepp](https://redirect.github.com/scottlepp)
- **Azure:** Fix Application Insights metadata requests [#&#8203;105614](https://redirect.github.com/grafana/grafana/pull/105614), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Fix duplicated trace links [#&#8203;105698](https://redirect.github.com/grafana/grafana/pull/105698), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Fix legend formatting [#&#8203;106504](https://redirect.github.com/grafana/grafana/pull/106504), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Fix resource name determination in template variable queries [#&#8203;105705](https://redirect.github.com/grafana/grafana/pull/105705), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **BarChart/StateTimeline:** Use noValue setting for error message when data is empty [#&#8203;107147](https://redirect.github.com/grafana/grafana/pull/107147), [@&#8203;fastfrwrd](https://redirect.github.com/fastfrwrd)
- **CloudWatch:** Fix http client handling + assume role bug [#&#8203;107893](https://redirect.github.com/grafana/grafana/pull/107893), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **CloudWatch:** Fix proxy transport issue [#&#8203;107807](https://redirect.github.com/grafana/grafana/pull/107807), [@&#8203;njvrzm](https://redirect.github.com/njvrzm)
- **Dashboard:** FF `dashboardNewLayouts` Fix library panels non-editable when multiple added [#&#8203;107052](https://redirect.github.com/grafana/grafana/pull/107052), [@&#8203;axelavargas](https://redirect.github.com/axelavargas)
- **Dashboard:** Fix cache validation to prevent stale cache [#&#8203;105918](https://redirect.github.com/grafana/grafana/pull/105918), [@&#8203;yashschandra](https://redirect.github.com/yashschandra)
- **Dashboard:** Fixes issue with dashboard links that include all variables [#&#8203;106356](https://redirect.github.com/grafana/grafana/pull/106356), [@&#8203;torkelo](https://redirect.github.com/torkelo)
- **Dashboards:** Fix history list for dashboard uids that end in `-` [#&#8203;107073](https://redirect.github.com/grafana/grafana/pull/107073), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Drilldown:** Fix js crash when using http [#&#8203;105646](https://redirect.github.com/grafana/grafana/pull/105646), [@&#8203;chu121su12](https://redirect.github.com/chu121su12)
- **Fix:** Increase login\_attempt.ip\_address column length for IPv6 support [#&#8203;107035](https://redirect.github.com/grafana/grafana/pull/107035), [@&#8203;Jguer](https://redirect.github.com/Jguer)
- **FlameGraph:** Fix bug for function names that conflict with JavaScript object prototype properties [#&#8203;106338](https://redirect.github.com/grafana/grafana/pull/106338), [@&#8203;simonswine](https://redirect.github.com/simonswine)
- **Folders:** Correctly resolve nested folder breadcrumbs [#&#8203;106344](https://redirect.github.com/grafana/grafana/pull/106344), [@&#8203;IevaVasiljeva](https://redirect.github.com/IevaVasiljeva)
- **GrafanaUI:** Fix Combobox ignoring loading prop [#&#8203;105584](https://redirect.github.com/grafana/grafana/pull/105584), [@&#8203;ValeraS](https://redirect.github.com/ValeraS)
- **Graphite:** Fix annotation queries [#&#8203;106553](https://redirect.github.com/grafana/grafana/pull/106553), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Graphite:** Fix date mutation [#&#8203;107414](https://redirect.github.com/grafana/grafana/pull/107414), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Graphite:** Fix nested variable interpolation for repeated rows [#&#8203;106976](https://redirect.github.com/grafana/grafana/pull/106976), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **K8s:** Dashboards /apis: Fix library element connections [#&#8203;106734](https://redirect.github.com/grafana/grafana/pull/106734), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Loki:** Fix health check message [#&#8203;107170](https://redirect.github.com/grafana/grafana/pull/107170), [@&#8203;wooffie](https://redirect.github.com/wooffie)
- **Loki:** Fix issue where step parameter using a template variable was marked as invalid [#&#8203;106541](https://redirect.github.com/grafana/grafana/pull/106541), [@&#8203;ivanahuckova](https://redirect.github.com/ivanahuckova)
- **Loki:** Fix label browser not sorted after selection of a label [#&#8203;107394](https://redirect.github.com/grafana/grafana/pull/107394), [@&#8203;paulojmdias](https://redirect.github.com/paulojmdias)
- **Org:** Fix org deletion [#&#8203;106193](https://redirect.github.com/grafana/grafana/pull/106193), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Plugins:** Fix and encode invalid gRPC header values [#&#8203;107339](https://redirect.github.com/grafana/grafana/pull/107339), [@&#8203;ivanahuckova](https://redirect.github.com/ivanahuckova)
- **PostgreSQL:** Fix error on panel when toggling sqlDatasourceDatabaseSelection feature [#&#8203;106965](https://redirect.github.com/grafana/grafana/pull/106965), [@&#8203;HasithDeAlwis](https://redirect.github.com/HasithDeAlwis)
- **Profiles:** Fix for passing the response headers [#&#8203;106293](https://redirect.github.com/grafana/grafana/pull/106293), [@&#8203;simonswine](https://redirect.github.com/simonswine)
- **Reporting:** Stop sending reports with Never schedule on creation (Enterprise)
- **SCIM:** Fix PUT request for deactivating a user (Enterprise)
- **SCIM:** Fix the removal of all members in group PUT requests (Enterprise)
- **SCIM:** Fix user patch operation (Enterprise)
- **Security:** Add fix for CVE-2025-3580 [#&#8203;105976](https://redirect.github.com/grafana/grafana/pull/105976), [@&#8203;baldm0mma](https://redirect.github.com/baldm0mma)
- **Security:** Fixes for CVE-2025-6197 and CVE-2025-6023 [#&#8203;108333](https://redirect.github.com/grafana/grafana/pull/108333), [@&#8203;mgyongyosi](https://redirect.github.com/mgyongyosi)
- **Settings:** Fix reencryption and rollback of encrypted values in setting table (Enterprise)
- **Tempo:** Fix showing dangling edges in NodeGraph [#&#8203;107245](https://redirect.github.com/grafana/grafana/pull/107245), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **ToolTip:** Fix flexbox bug with tooltip when `maxWidth` is set manually [#&#8203;107145](https://redirect.github.com/grafana/grafana/pull/107145), [@&#8203;jdmarshall](https://redirect.github.com/jdmarshall)
- **URLParams:** Stringify true values as key=true always (fixes issues with variables with true value) [#&#8203;106440](https://redirect.github.com/grafana/grafana/pull/106440), [@&#8203;torkelo](https://redirect.github.com/torkelo)

##### Breaking changes

- **Alerting:** Enable recording rules by default [#&#8203;105603](https://redirect.github.com/grafana/grafana/pull/105603), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)

##### Plugin development fixes & changes

- **Carousel:** Always center image [#&#8203;106468](https://redirect.github.com/grafana/grafana/pull/106468), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)
- **Drawer:** Include divider and close button when passing a custom title element [#&#8203;106896](https://redirect.github.com/grafana/grafana/pull/106896), [@&#8203;ashharrison90](https://redirect.github.com/ashharrison90)

<!-- 12.1.0 END -->

<!-- 12.0.3 START -->

### [`v12.0.3`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1203-2025-07-23)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.0.2...v12.0.3)

##### Bug fixes

- **Security:** Fixes for CVE-2025-6197 and CVE-2025-6023 [#&#8203;108280](https://redirect.github.com/grafana/grafana/pull/108280), [@&#8203;volcanonoodle](https://redirect.github.com/volcanonoodle)

<!-- 12.0.3 END -->

<!-- 11.6.4 START -->

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6InNlY3VyaXR5LWNvbXBsaWFuY2UiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on July 28, 2025 at 01:46 AM UTC

Commits missing Jira IDs:
2deb879a82e0cd336a5c379ecb202bbaeeb92c1e


### Comment by @red-hat-konflux on July 28, 2025 at 09:32 AM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`12.1.0`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1754*
