# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.5](https://github.com/gosusnp/whackamole/compare/v0.1.4...v0.1.5) (2026-07-28)


### Features

* add whack mcp install-claude-code ([0882c6e](https://github.com/gosusnp/whackamole/commit/0882c6eb1dce8d6fc33ccb77e735c409e713bfc7))

## [0.1.4](https://github.com/gosusnp/whackamole/compare/v0.1.3...v0.1.4) (2026-03-20)


### Features

* enum field updates are less strict for less agent failures ([951702d](https://github.com/gosusnp/whackamole/commit/951702d17f958404221d5f5a10a6cdc5cd6622cf))
* **ui:** make description text area resizable ([2bc363c](https://github.com/gosusnp/whackamole/commit/2bc363cc9fd4a3b17d462dd81b8024b58568782c))

## [0.1.3](https://github.com/gosusnp/whackamole/compare/v0.1.2...v0.1.3) (2026-03-16)


### Features

* add clear finished ([11c4204](https://github.com/gosusnp/whackamole/commit/11c4204ea75f019ca638bdd23739e24e37504147))
* add whack config write-local-mds ([ca5bc51](https://github.com/gosusnp/whackamole/commit/ca5bc51bbc81ab5a9f69858276c8e93a3f8b647a))

## [0.1.2](https://github.com/gosusnp/whackamole/compare/v0.1.1...v0.1.2) (2026-03-13)


### Features

* **ui:** leaderboard ([8890c40](https://github.com/gosusnp/whackamole/commit/8890c404b54c8a6be293025f3f1a4db5ddeadc0c))
* **ui:** whac-a-mole ([c2ac163](https://github.com/gosusnp/whackamole/commit/c2ac163c0b456ec98d41be2deee842198c972f08))


### Bug Fixes

* remove per request store instantiation ([e348cf7](https://github.com/gosusnp/whackamole/commit/e348cf78a7ad482931bf9dfa4752f15a504dfcdb))

## [0.1.1](https://github.com/gosusnp/whackamole/compare/v0.1.0...v0.1.1) (2026-03-13)


### Features

* add whack_history table ([faacd22](https://github.com/gosusnp/whackamole/commit/faacd228a82ad4b3137913e0c9a9a27332ac1a21))
* limit sqlite connections ([8ac7085](https://github.com/gosusnp/whackamole/commit/8ac70850aac246738a1b81f6d0b3de793b6e99e3))
* make large card collapsible ([b405a50](https://github.com/gosusnp/whackamole/commit/b405a5036bb79df8af6a5bcab2e6cd5bce9ff8dd))
* populate history table ([5ecfad2](https://github.com/gosusnp/whackamole/commit/5ecfad2ffbc2269f3523ca4014cc2e03d9082819))
* **ui:** listen to changes and update the content as needed ([3d7fc0c](https://github.com/gosusnp/whackamole/commit/3d7fc0c31a58ef6014a5a368c4ad0bbe8624c501))


### Bug Fixes

* ensure task status consistency ([10d41a3](https://github.com/gosusnp/whackamole/commit/10d41a39aa955483289a35e6115e7b5995853df8))
* improve history polling performance ([eb83e53](https://github.com/gosusnp/whackamole/commit/eb83e535354ca029a4b404fcd4a66a2772317b30))
* **ui:** address deletion progress bar cpu load ([2f3217c](https://github.com/gosusnp/whackamole/commit/2f3217c7d7655e0380c3e4ef4f58f231da4e37ec))
* **ui:** avoid re-render cascades ([0d721d6](https://github.com/gosusnp/whackamole/commit/0d721d6608056b6241eb1f7b14ee24e5f4f5aabc))
* **ui:** creating task do not trigger the refresh button ([c8e1f81](https://github.com/gosusnp/whackamole/commit/c8e1f8183c065cb4ef45ae8e20609010d7a5ce20))
* **ui:** deletion progressbar loop restarts on unrelated action ([cbdbbab](https://github.com/gosusnp/whackamole/commit/cbdbbabc58181bbe1d82b64b7dc3bb9aa0347bc7))
* **ui:** keep track of deletions despite tab changes ([6883b69](https://github.com/gosusnp/whackamole/commit/6883b69cac5782a9b258ea7b62ec74927e6b014a))

## [0.1.0](https://github.com/gosusnp/whackamole/compare/v0.0.8...v0.1.0) (2026-03-10)


### ⚠ BREAKING CHANGES

* merge fix type into bug for a cleare typing

### Features

* add a review status ([375e16b](https://github.com/gosusnp/whackamole/commit/375e16b82983316eafc75c31d6a4160e4c6e3fa4))
* add project key validation ([5f33422](https://github.com/gosusnp/whackamole/commit/5f33422ca6ecbecb446c8fe4d72ecbc21caee597))
* add support for configurable MCP.instructions ([8999156](https://github.com/gosusnp/whackamole/commit/89991562617b34f80159f2c33ceb9c67807e7c72))
* merge fix type into bug for a cleare typing ([e490110](https://github.com/gosusnp/whackamole/commit/e490110cda87d0eebc6a16f50ac6d3bf76399d16))

## [0.0.8](https://github.com/gosusnp/whackamole/compare/v0.0.7...v0.0.8) (2026-03-09)


### Features

* bolder task type color coding in the UI ([f6c0903](https://github.com/gosusnp/whackamole/commit/f6c0903a0adc693da2fe26052088936cfdad1287))
* **ui:** add project add button ([f9012f8](https://github.com/gosusnp/whackamole/commit/f9012f86c94826cc4c0e5d70e40ec66e7319fda0))
* **ui:** enable project delete ([09cf1a3](https://github.com/gosusnp/whackamole/commit/09cf1a3c0ad730e0303e33947fd4c842f97a4868))
* **ui:** style the statuses ([5058254](https://github.com/gosusnp/whackamole/commit/5058254a3bffae48082d6ec928402de15fff5ea8))


### Bug Fixes

* **ui:** fype type picker zIndex and color ([a6febc1](https://github.com/gosusnp/whackamole/commit/a6febc18b999de04df9456e375dce45efe7961fb))

## [0.0.7](https://github.com/gosusnp/whackamole/compare/v0.0.6...v0.0.7) (2026-03-07)


### Features

* task delete from the ui ([b0ac303](https://github.com/gosusnp/whackamole/commit/b0ac30328d7594dec23d5aaaeb885cfa1bea20fe))


### Bug Fixes

* ui can delete multiple tasks in parallel ([8910ab0](https://github.com/gosusnp/whackamole/commit/8910ab0fd93b7400628dd6ad60d440cffe479416))
* undo go version bump for gorelease and remove checks ([860ce70](https://github.com/gosusnp/whackamole/commit/860ce700fdab2553da2d90a18e5129d257bbf1c2))

## [0.0.6](https://github.com/gosusnp/whackamole/compare/v0.0.5...v0.0.6) (2026-03-07)


### Bug Fixes

* goreleaser with tests ([#10](https://github.com/gosusnp/whackamole/issues/10)) ([eade200](https://github.com/gosusnp/whackamole/commit/eade200585240b334cab0a1c5dd33d6370251c57))

## [0.0.5](https://github.com/gosusnp/whackamole/compare/v0.0.4...v0.0.5) (2026-03-07)


### Bug Fixes

* go release syntax ([6bd86dd](https://github.com/gosusnp/whackamole/commit/6bd86dd7360293ec5462eca5387ce0b7647fa0f8))

## [0.0.4](https://github.com/gosusnp/whackamole/compare/v0.0.3...v0.0.4) (2026-03-07)


### Reverts

* feat: trim binary size ([c8a23f6](https://github.com/gosusnp/whackamole/commit/c8a23f6453324e5138bbb46cff1a6a6f6567fe05))

## [0.0.3](https://github.com/gosusnp/whackamole/compare/v0.0.2...v0.0.3) (2026-03-06)


### Features

* ability to create task from the UI ([fb59587](https://github.com/gosusnp/whackamole/commit/fb59587e626d74d35e138934088af3c6f3be3aa1))
* add light and dark mode to the ui ([022a1ce](https://github.com/gosusnp/whackamole/commit/022a1ce4be52102e7e13fa45847ce087d940c83f))
* add whack ui ([#6](https://github.com/gosusnp/whackamole/issues/6)) ([2fe92db](https://github.com/gosusnp/whackamole/commit/2fe92dbc4b53969f53dc0a27c599440241abb51a))
* enable SQLite WAL mode ([7ccfc67](https://github.com/gosusnp/whackamole/commit/7ccfc67328a35ee4af306ede707f9f88a137de1d))
* expose markdown support to agents ([97a131f](https://github.com/gosusnp/whackamole/commit/97a131f528bf268ee4f948c54b1f9895c6a52917))
* show logo ([5f9cd16](https://github.com/gosusnp/whackamole/commit/5f9cd16d3cacac46e0beac55312c39741b863980))
* trim binary size ([7c6007e](https://github.com/gosusnp/whackamole/commit/7c6007e36a1c2478340fa41d5b0d219a72f9df09))
* ui patches tasks instead of overriding the whole object ([857871b](https://github.com/gosusnp/whackamole/commit/857871b6672f173a1816a734060adba4743d1afd))
* ui remembers last viewed project across refreshes ([9f6e84c](https://github.com/gosusnp/whackamole/commit/9f6e84ccd8a7102a6de283fd56dcb7bb96b090d3))

## [0.0.2](https://github.com/gosusnp/whackamole/compare/v0.0.1...v0.0.2) (2026-03-06)


### Bug Fixes

* clean up release-please-config ([33bf51d](https://github.com/gosusnp/whackamole/commit/33bf51d87f7d3deffbe0a9d6b0e1c3b553c53140))
* fix release please workflow ([60b368d](https://github.com/gosusnp/whackamole/commit/60b368d2c20ee04ccc733dc7bd0ab3b012d501a9))

## 0.0.1 (2026-03-05)


### Features

* add config file support ([54b2bff](https://github.com/gosusnp/whackamole/commit/54b2bffde081792488b37dcab7644633b853a98c))
* add project show ([c90a777](https://github.com/gosusnp/whackamole/commit/c90a777415946dc269a24f8c30fbaeee8cddf5cf))
* add stdio mcp server ([32f2b27](https://github.com/gosusnp/whackamole/commit/32f2b274940ae38a75d005adf747d0a3ad66f4f7))
* add task ([6486882](https://github.com/gosusnp/whackamole/commit/64868822c5269a53a21938c0b0e13e6a5b5b8df5))
* add whack config to print config used ([464f661](https://github.com/gosusnp/whackamole/commit/464f661d2e9b7a54ecac0e8a7a9b6af7cbf97d82))
* create db path if not exists ([31770c3](https://github.com/gosusnp/whackamole/commit/31770c356ec6ea20fd80f2feb57c08baf1beff01))
* default db path is in home instead of current dir ([3960062](https://github.com/gosusnp/whackamole/commit/3960062741a4d2c7ae7dfc7c6ee3888e8b38c6b7))
* improve task list output ([26af784](https://github.com/gosusnp/whackamole/commit/26af7846f68bf471ed69eb6045eedf9b151644dc))
* setup sqlite persistence and add whack project ([65aab65](https://github.com/gosusnp/whackamole/commit/65aab65459b36d761e39ee6b7ee56d2619f8d0a2))
* switch to using readable project key instead of id ([0b7d280](https://github.com/gosusnp/whackamole/commit/0b7d2802301776ab604bf7e69f704895f2f90f66))
* task list filters done tasks by default ([7e112e4](https://github.com/gosusnp/whackamole/commit/7e112e49c2fd3e35e0dd003ab95f8ca2f4d3d219))
* use hardened types ([2c08d77](https://github.com/gosusnp/whackamole/commit/2c08d7710cfa45e23932c274ef2a10f65560383a))


### Bug Fixes

* fix install flags ([fbd8adb](https://github.com/gosusnp/whackamole/commit/fbd8adb72ccb3a6c37abc62480831472b3cff68b))
* fix version string ([219e9ca](https://github.com/gosusnp/whackamole/commit/219e9ca01c6d2494ae691454d61dbdd42c6949a7))

## [Unreleased]
