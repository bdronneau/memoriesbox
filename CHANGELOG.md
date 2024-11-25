# Changelog

## [2.0.2](https://github.com/bdronneau/memoriesbox/compare/v2.0.1...v2.0.2) (2024-11-25)


### 🧰 Other

* bump github.com/stretchr/testify in the deps group ([18e08ca](https://github.com/bdronneau/memoriesbox/commit/18e08ca581c53beb04b6155b2630e46ac6cc56f5))

## [2.0.1](https://github.com/bdronneau/memoriesbox/compare/v2.0.0...v2.0.1) (2024-11-18)


### 🧰 Other

* bump github.com/jackc/pgx/v5 in the deps group ([6c9468b](https://github.com/bdronneau/memoriesbox/commit/6c9468b6a11247a5fe6b87161eff110723c6fa5a))
* bump github.com/rs/xid from 1.5.0 to 1.6.0 in the deps group ([f137d9d](https://github.com/bdronneau/memoriesbox/commit/f137d9d3d795e1676c28a18d116f0863a3156c04))
* bump github.com/volatiletech/sqlboiler/v4 in the deps group ([0e5af5b](https://github.com/bdronneau/memoriesbox/commit/0e5af5b7e329a5a8007a5f529b1bbaef86535789))
* bump github.com/volatiletech/sqlboiler/v4 in the deps group ([938dce0](https://github.com/bdronneau/memoriesbox/commit/938dce09ad8985a15d56f1cfb45ec8108a2e0009))
* bump github.com/volatiletech/strmangle in the deps group ([fbbf867](https://github.com/bdronneau/memoriesbox/commit/fbbf867454a0f34ebeee6e1ea88ddf9c048d694d))
* bump github.com/volatiletech/strmangle in the deps group ([91a778c](https://github.com/bdronneau/memoriesbox/commit/91a778c703ffe9b95feb571d46c918a535f2233f))
* bump go.uber.org/mock from 0.4.0 to 0.5.0 in the deps group ([e4a4ea1](https://github.com/bdronneau/memoriesbox/commit/e4a4ea1ac1fd92f249517417916035e6824f9c61))
* bump migrate/migrate from v4.17.1 to v4.18.1 ([bea2bbd](https://github.com/bdronneau/memoriesbox/commit/bea2bbdd2df37d415178a5d29e85b22f461604e7))
* bump the deps group with 2 updates ([b6d752d](https://github.com/bdronneau/memoriesbox/commit/b6d752d6d9e9a58eb87a8c57065004e257f9d173))
* **deps:** bump golang from 1.22 to 1.23 ([2fa0c92](https://github.com/bdronneau/memoriesbox/commit/2fa0c928dd19fa6c3e480b73d8b1465951862758))
* go mod tidy ([2cdb133](https://github.com/bdronneau/memoriesbox/commit/2cdb133d372a319507e3700361426bafb978a3d8))

## [2.0.0](https://github.com/bdronneau/memoriesbox/compare/v1.5.0...v2.0.0) (2024-06-03)


### ⚠ BREAKING CHANGES

* migrate primary key to xid

### 🚀 Features

* add xid column ([8b7e628](https://github.com/bdronneau/memoriesbox/commit/8b7e6283755e645bde5653c6968cc5ca89f35e0c))
* **cmd:** add migrate script to xid ([d85b31e](https://github.com/bdronneau/memoriesbox/commit/d85b31edfe3fbdd53c623d929c6c46116210995a))
* **dependabot:** add docker update ([7092dc2](https://github.com/bdronneau/memoriesbox/commit/7092dc2efe93c6283432ddc74f48a1885a920e46))
* **dependabot:** add group ([2b19952](https://github.com/bdronneau/memoriesbox/commit/2b19952dee1e9be438d6431ba1bca7e0c514b3d4))
* **golang:** bump to 1.21 ([2699033](https://github.com/bdronneau/memoriesbox/commit/26990337a3fc23281d27fbce605fc5ce05dba6b8))
* **Makefile:** add init step ([3eed9be](https://github.com/bdronneau/memoriesbox/commit/3eed9be2726ce5a31832fec11aea757738ea4de8))


### 🐞 Bug Fixes

* **addmemory:** conditionnal logger ([e37d6a6](https://github.com/bdronneau/memoriesbox/commit/e37d6a6eb8c47fa74904a1179a90c727d4515167))
* **dotenv:** wrong example ([6cbc398](https://github.com/bdronneau/memoriesbox/commit/6cbc398f32f198dfce5a4272bdbcb54e8be5e70f))
* **utils:** use contains ([84e2486](https://github.com/bdronneau/memoriesbox/commit/84e2486339aa4299af6eb0e0da833a7c2d3df77b))


### 🛠️ Refactor

* clean old migrate cmd ([01f6e97](https://github.com/bdronneau/memoriesbox/commit/01f6e9778c122b4e6c9d5c269c4e5c59fd000f3c))
* **log:** move to slog ([b27716c](https://github.com/bdronneau/memoriesbox/commit/b27716c6336b762061b8701cc59362c2e2154180))
* migrate primary key to xid ([3203ac6](https://github.com/bdronneau/memoriesbox/commit/3203ac6f9b61265be6f0f31a4ce97a1e63028e29))
* **release-please:** handle new version ([a181611](https://github.com/bdronneau/memoriesbox/commit/a181611559c9613d4fad7bf452a40eaef497cbf9))


### 🧰 Other

* betteralign ([579911d](https://github.com/bdronneau/memoriesbox/commit/579911d6155883f22d4a2a910c51cc19966a2ce2))
* bump github.com/DATA-DOG/go-sqlmock from 1.5.0 to 1.5.1 ([525d84a](https://github.com/bdronneau/memoriesbox/commit/525d84a0aac9644559f35180cb70f3cb5abc968e))
* bump github.com/jackc/pgx/v4 from 4.18.1 to 4.18.2 ([af31a62](https://github.com/bdronneau/memoriesbox/commit/af31a62d29ce2eea12ed8fcfb7afd90d7f72bf19))
* bump github.com/jackc/pgx/v5 from 5.4.3 to 5.5.0 ([6f6a675](https://github.com/bdronneau/memoriesbox/commit/6f6a675e34f6c03a0fe4529de57b0c950c3d4049))
* bump github.com/jackc/pgx/v5 from 5.5.0 to 5.5.1 ([9502208](https://github.com/bdronneau/memoriesbox/commit/95022085476d3962a3e7156d70af7ab582b3a20d))
* bump github.com/jackc/pgx/v5 in the deps group ([6fdd2c8](https://github.com/bdronneau/memoriesbox/commit/6fdd2c8017977cd4b107789a0a591ac0bcdbc560))
* bump github.com/labstack/echo/v4 from 4.11.1 to 4.11.2 ([bcd0c3f](https://github.com/bdronneau/memoriesbox/commit/bcd0c3f0a199ba2eac12eba29ec8eab2cbbc039c))
* bump github.com/labstack/echo/v4 from 4.11.2 to 4.11.3 ([546c1ab](https://github.com/bdronneau/memoriesbox/commit/546c1aba589854b03e2dde13932cec0bdb9f5d01))
* bump github.com/rs/xid from 1.2.1 to 1.5.0 ([3b274bb](https://github.com/bdronneau/memoriesbox/commit/3b274bbd85d26f10dae6700184684963223d6a08))
* bump github.com/spf13/viper from 1.16.0 to 1.17.0 ([78282ef](https://github.com/bdronneau/memoriesbox/commit/78282efb11a7f76fa5d9f24812f0957fee2170a4))
* bump github.com/spf13/viper from 1.17.0 to 1.18.1 ([585ad2d](https://github.com/bdronneau/memoriesbox/commit/585ad2d8a2353ba5539c3d90fdf7efca02421257))
* bump github.com/spf13/viper in the deps group ([382796e](https://github.com/bdronneau/memoriesbox/commit/382796ed0c7e2a5979d149ef6dc6dd1f1d8f839e))
* bump golang from 1.21 to 1.22 ([23f5d12](https://github.com/bdronneau/memoriesbox/commit/23f5d120674a10c375656e71aa1a78e962298bb9))
* bump golang.org/x/crypto from 0.16.0 to 0.17.0 ([4e97018](https://github.com/bdronneau/memoriesbox/commit/4e97018123359e2ca7a9834eee9e98cf60686865))
* bump golang.org/x/net from 0.15.0 to 0.17.0 ([12431fb](https://github.com/bdronneau/memoriesbox/commit/12431fb95a1c00e25faac9cde75cb54cdd508341))
* bump migrate/migrate from v4.15.2 to v4.17.0 ([5f5818a](https://github.com/bdronneau/memoriesbox/commit/5f5818a74bdd03ff4e17478e67a12c1d25451612))
* bump migrate/migrate from v4.17.0 to v4.17.1 ([d1dea75](https://github.com/bdronneau/memoriesbox/commit/d1dea75b2185cfc589367e76f6e47b2ce402044b))
* bump the deps group with 1 update ([3bcac7a](https://github.com/bdronneau/memoriesbox/commit/3bcac7abb96a1c40237dcb025f641512506bf779))
* bump the deps group with 1 update ([5e8cbc0](https://github.com/bdronneau/memoriesbox/commit/5e8cbc0eec363e11804656c74c1039eb5f269c1a))
* bump the deps group with 1 update ([b03e5ed](https://github.com/bdronneau/memoriesbox/commit/b03e5ed425254212b92b16dd731315ea14c3943a))
* bump the deps group with 1 update ([debde3a](https://github.com/bdronneau/memoriesbox/commit/debde3ae90079283a699d0e8299fe3ec9c1834bc))
* bump the deps group with 1 update ([929bc95](https://github.com/bdronneau/memoriesbox/commit/929bc955ba584472aee4be7835ec6449fdfd21ef))
* bump the deps group with 1 update ([ce8eaca](https://github.com/bdronneau/memoriesbox/commit/ce8eacaa88d7ac1b2a7fe2c0db1979b4ad622d8e))
* bump the deps group with 2 updates ([c7eea37](https://github.com/bdronneau/memoriesbox/commit/c7eea37bcd43653c1a10146506cf80b7315dc9a4))
* bump the deps group with 2 updates ([b9146b0](https://github.com/bdronneau/memoriesbox/commit/b9146b031b852221b0a4bf015ecdc625960d816f))
* bump the deps group with 5 updates ([5da5a1c](https://github.com/bdronneau/memoriesbox/commit/5da5a1c2478288ebeb89d53228d5426f06d67957))

## [1.5.0](https://github.com/bdronneau/memoriesbox/compare/v1.4.2...v1.5.0) (2023-09-25)


### 🚀 Features

* allow insert memory ([45caf24](https://github.com/bdronneau/memoriesbox/commit/45caf2482252284642e5e5a297ac088426c93187))
* **Makefile:** add coverage option ([bc1a734](https://github.com/bdronneau/memoriesbox/commit/bc1a734a6aedd0d120869a4b0b0ab101dc1961a6))
* **test:** add functionnal tests ([1475d0a](https://github.com/bdronneau/memoriesbox/commit/1475d0ac9a7667ce2425c74718c3e9c72d2243d3))


### 🐞 Bug Fixes

* **Dockerfile:** wrong executable ([89376cd](https://github.com/bdronneau/memoriesbox/commit/89376cd905489d5918556378d4510fe4ec4c6ef6))
* **Makefile:** protect DSN ([555bc02](https://github.com/bdronneau/memoriesbox/commit/555bc0275369f6f6ece81c70586629d4a6aa6ef1))
* **sonar:** handle alert ([a0ba206](https://github.com/bdronneau/memoriesbox/commit/a0ba206a69b80b731f9f55c166b52fec714abae0))


### ✨ Polish

* add way to disable ADD interface ([f717243](https://github.com/bdronneau/memoriesbox/commit/f717243a972d24710d48ed87c978ad40e8491cf6))
* **DB:** optimize declaration ([539b5e8](https://github.com/bdronneau/memoriesbox/commit/539b5e8de038249f69b62821f229fde4744dd345))
* **github:** optimize golangci-lint step ([c1d6adb](https://github.com/bdronneau/memoriesbox/commit/c1d6adba8ad085e03ca838a36bc0aee112bd7abb))
* **interface:** migrate to bulma ([ca7631c](https://github.com/bdronneau/memoriesbox/commit/ca7631c50f5dd5b1fe568168850d5ba790e39222))
* **memories:** add error handling ([d660b82](https://github.com/bdronneau/memoriesbox/commit/d660b823b26298463ad54c041bccd78990561d97))
* **mocks:** update generate way ([c1c5dc9](https://github.com/bdronneau/memoriesbox/commit/c1c5dc9bc8ee0e3d298a5be68df78cc41f408c39))
* **pgx:** bump version ([8cdae60](https://github.com/bdronneau/memoriesbox/commit/8cdae6074b9efbfe893a4ebf01fd6b0c04d23e53))
* **sonar:** rename ([1c54e1b](https://github.com/bdronneau/memoriesbox/commit/1c54e1bb067df1bda4ff751936577114e53e1d1c))
* **sonar:** skip test quality ([c71c411](https://github.com/bdronneau/memoriesbox/commit/c71c41152d8fb861faba4c0351b51434c3975d8d))
* **sqlboiler:** allow override user ([6f78788](https://github.com/bdronneau/memoriesbox/commit/6f78788e0e7ca7e36cbb8d6779d3fd126c864896))
* **template:** remove duplicate ([8f7394a](https://github.com/bdronneau/memoriesbox/commit/8f7394ab52cddbc8fea6cdc0468f86b561ac2240))


### 🧰 Other

* bump github.com/jackc/pgx/v5 from 5.3.1 to 5.4.3 ([d9bd2d9](https://github.com/bdronneau/memoriesbox/commit/d9bd2d97ea0c2b26d2ca327ecb51e430fd7158fb))
* bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([a10ffc3](https://github.com/bdronneau/memoriesbox/commit/a10ffc3a54c07da870b51b40618ff96a0e1c0488))
* bump github.com/volatiletech/sqlboiler/v4 from 4.14.2 to 4.15.0 ([bac0e98](https://github.com/bdronneau/memoriesbox/commit/bac0e98c62fb71323c285e0cf365ad88b056e21a))
* bump go.uber.org/mock from 0.2.0 to 0.3.0 ([0678b88](https://github.com/bdronneau/memoriesbox/commit/0678b8857fcd280b782e89171e740670061e7b8b))
* bump go.uber.org/zap from 1.24.0 to 1.25.0 ([2e0bd2f](https://github.com/bdronneau/memoriesbox/commit/2e0bd2fe74eb6dc09c885c7bfb19d1de6f4e2c14))
* bump go.uber.org/zap from 1.25.0 to 1.26.0 ([d9ba7be](https://github.com/bdronneau/memoriesbox/commit/d9ba7bea739c8fdbd3793f27738be469f51ab16f))
* **docs:** rewamp documentation ([afe13b6](https://github.com/bdronneau/memoriesbox/commit/afe13b6f9c37629cd50d87fcf33812acbbfe9404))
* **go:** migrate to uber mock ([737c59c](https://github.com/bdronneau/memoriesbox/commit/737c59c2a6c26bf577c7afb4c98453ce600c5beb))
* please the linter ([342af73](https://github.com/bdronneau/memoriesbox/commit/342af732d2e4bca1f65c10f4d0e38c7f4499223c))
* remove duplicate ([3159b24](https://github.com/bdronneau/memoriesbox/commit/3159b242d8825e2ca362ffb07d30d9b8a8135170))
* update go.mod ([04bc159](https://github.com/bdronneau/memoriesbox/commit/04bc159307f3386fdf1581a605697e8a95de95ee))

## [1.4.2](https://github.com/bdronneau/memoriesbox/compare/v1.4.1...v1.4.2) (2023-07-24)


### 🧰 Other

* bump github.com/jackc/pgx/v4 from 4.18.0 to 4.18.1 ([4983027](https://github.com/bdronneau/memoriesbox/commit/4983027485a84c54a272d89518c2838b399825e7))
* bump github.com/labstack/echo/v4 from 4.10.2 to 4.11.1 ([ced59f3](https://github.com/bdronneau/memoriesbox/commit/ced59f3139b3517e872b6e04fbb782b0d4b06df5))
* bump github.com/lib/pq from 1.10.7 to 1.10.8 ([90755f6](https://github.com/bdronneau/memoriesbox/commit/90755f6b89de431b2a36290e12c75124ab186c3e))
* bump github.com/lib/pq from 1.10.8 to 1.10.9 ([ce2a29e](https://github.com/bdronneau/memoriesbox/commit/ce2a29ec2e5e7b142e2c2d71abafa8c600b940f4))
* bump github.com/peterbourgon/ff/v3 from 3.3.0 to 3.3.1 ([a4730c9](https://github.com/bdronneau/memoriesbox/commit/a4730c91d5c9ffa143dad8b6c401ebe0c58e0f93))
* bump github.com/peterbourgon/ff/v3 from 3.3.1 to 3.3.2 ([f397dac](https://github.com/bdronneau/memoriesbox/commit/f397dac9d363883f69bc61a5851c4b0e8d5614d8))
* bump github.com/peterbourgon/ff/v3 from 3.3.2 to 3.4.0 ([b879748](https://github.com/bdronneau/memoriesbox/commit/b879748fef46bcca98ba5db0e0d031e007b9d39b))
* bump github.com/spf13/viper from 1.15.0 to 1.16.0 ([d5d6c76](https://github.com/bdronneau/memoriesbox/commit/d5d6c76ad7868bd92cf0b3d6e1ee6e2b2225ae4a))
* bump github.com/volatiletech/sqlboiler/v4 from 4.14.1 to 4.14.2 ([7dc0a05](https://github.com/bdronneau/memoriesbox/commit/7dc0a0529234f9f0966526b8dc4f3922a7671233))
* bump github.com/volatiletech/strmangle from 0.0.4 to 0.0.5 ([c4a0f50](https://github.com/bdronneau/memoriesbox/commit/c4a0f509799545a5785f0a08a786a514fdd29dcf))
* update license ([b177ba2](https://github.com/bdronneau/memoriesbox/commit/b177ba24fa87fbe1b50ef552c3ae47e6ed68b4ae))

## [1.4.1](https://github.com/bdronneau/memoriesbox/compare/v1.4.0...v1.4.1) (2023-02-27)


### 🐞 Bug Fixes

* **sonarcloud:** wrong directive ([f589de1](https://github.com/bdronneau/memoriesbox/commit/f589de1929691f19564caa0c48bc3b8375730c5e))


### ✨ Polish

* **tests:** use constants ([b7f72f1](https://github.com/bdronneau/memoriesbox/commit/b7f72f1bc1dbe02db906cb67253a973c5ae6efe1))
* **tools:** do not take care of tests files ([306c161](https://github.com/bdronneau/memoriesbox/commit/306c161fbdb6e4ecbb5e017c4b151b5294ab3872))


### 🧰 Other

* bump github.com/DATA-DOG/go-sqlmock from 1.4.1 to 1.5.0 ([cc3c516](https://github.com/bdronneau/memoriesbox/commit/cc3c516b449c4699db44a56a5c4d9d90afa8382e))
* bump github.com/labstack/echo/v4 from 4.10.1 to 4.10.2 ([1672540](https://github.com/bdronneau/memoriesbox/commit/16725401388f3f02fd91c346739f874ede8086dc))
* **dockerfile:** upper case ([3709fc2](https://github.com/bdronneau/memoriesbox/commit/3709fc2cbb04b3231adc7e491ee7ca33dd4d098b))

## [1.4.0](https://github.com/bdronneau/memoriesbox/compare/v1.3.0...v1.4.0) (2023-02-26)


### 📚 Documentation

* **README:** wordin ([cff8fad](https://github.com/bdronneau/memoriesbox/commit/cff8fad2e55da93e1b647196a4ea047b6c074640))


### 🚀 Features

* **repositories:** init tests ([9b64503](https://github.com/bdronneau/memoriesbox/commit/9b64503ed122537b3831b361c8250af35d67a19a))


### 🧰 Other

* clean code ([8e54aa3](https://github.com/bdronneau/memoriesbox/commit/8e54aa36cef80f69bfd95797f41704ad52b3722b))
* **repositories:** clean logger ([cdc8d77](https://github.com/bdronneau/memoriesbox/commit/cdc8d77613f6e0064b8f4fa56e7a3b043f1408cf))
* trim line ([a0612c2](https://github.com/bdronneau/memoriesbox/commit/a0612c23429c56f6305a675519251b9c7e4b2b1a))


### ✨ Polish

* **ci:** add coverage ([b1d262e](https://github.com/bdronneau/memoriesbox/commit/b1d262e93d875913ce89e1b1f96b2069c16ec13d))
* **github:** rename action ([5a4cee4](https://github.com/bdronneau/memoriesbox/commit/5a4cee4937f94c27c136c764a02defcf90041bf4))
* **import:** rename module ([453e390](https://github.com/bdronneau/memoriesbox/commit/453e390d5d035bfc28abde403cd02034bdebe4ed))

## [1.3.0](https://github.com/bdronneau/memoriesbox/compare/v1.2.0...v1.3.0) (2023-02-21)


### 🚀 Features

* **golang:** move to 1.20 ([f85ac5b](https://github.com/bdronneau/memoriesbox/commit/f85ac5b93e472f382ee53b0365433bf56fbd7712))
* **index:** rewamp some informations and add favicon ([38c7b4f](https://github.com/bdronneau/memoriesbox/commit/38c7b4f3b138a009d056db19882e9a1ce1f30f04))


### 🧰 Other

* add some backlinks ([de1f5e2](https://github.com/bdronneau/memoriesbox/commit/de1f5e2db5dee2c2597967d6dcead9de9575d8c2))


### 🐞 Bug Fixes

* missing embed ([acf1586](https://github.com/bdronneau/memoriesbox/commit/acf15867eb5b04b8949dc7ba3f4cfe2112793bb5))

## [1.2.0](https://github.com/bdronneau/memoriesbox/compare/v1.1.1...v1.2.0) (2023-02-21)


### 📚 Documentation

* **README:** add badges ([e8f30a4](https://github.com/bdronneau/memoriesbox/commit/e8f30a453c20b7f03cffd9b794f22d185f8a2c5d))


### ✨ Polish

* **ci:** merge build and ci ([414749e](https://github.com/bdronneau/memoriesbox/commit/414749ec23ed5e4cb79bb82c708d40414ba4317c))
* **indeapi:** change default binding on interfaces ([21029ab](https://github.com/bdronneau/memoriesbox/commit/21029abd06947a74b0b0d78ea70f92268ac0e87c))
* **index:** move script declaration ([da786df](https://github.com/bdronneau/memoriesbox/commit/da786df72f3bf0037f5321bc9c6703badbee399f))
* **release-please:** add changelog types ([e832461](https://github.com/bdronneau/memoriesbox/commit/e832461c03067f2a12b151b873e36bbaff99e7ee))


### 🚀 Features

* **github:** add files required for opensource ([89b5893](https://github.com/bdronneau/memoriesbox/commit/89b58932c1c257009291db0ffb4d3901aca98962))
* **index:** formating of date and author ([e9cbbbc](https://github.com/bdronneau/memoriesbox/commit/e9cbbbc16cb65d6f149cd283714d018a2d3716a1))


### 🧰 Other

* **github:** init file ([69d5bcd](https://github.com/bdronneau/memoriesbox/commit/69d5bcdd9e9bcc19c02924ec98518f6830522db5))
* **github:** wording ([750ea58](https://github.com/bdronneau/memoriesbox/commit/750ea58f8606e93d2122711b0bfa2610174e7457))

## [1.1.1](https://github.com/bdronneau/memoriesbox/compare/v1.1.0...v1.1.1) (2023-02-20)


### Bug Fixes

* **db:** missing schema creation ([cb1f4ec](https://github.com/bdronneau/memoriesbox/commit/cb1f4ecae193bcbcae083b35e71ef0a393e3e074))

## [1.1.0](https://github.com/bdronneau/memoriesbox/compare/v1.0.3...v1.1.0) (2023-02-20)


### Features

* **docker:** include db migrations in container ([363e2e9](https://github.com/bdronneau/memoriesbox/commit/363e2e9f0406be5e45eeeeb80aafb215ddb0a99b))

## [1.0.3](https://github.com/bdronneau/memoriesbox/compare/v1.0.2...v1.0.3) (2023-02-20)


### Bug Fixes

* **docker:** use short ref name for build ([f4d634c](https://github.com/bdronneau/memoriesbox/commit/f4d634c557cacc882e82a00e3329fb30d8d94cec))

## [1.0.2](https://github.com/bdronneau/memoriesbox/compare/v1.0.1...v1.0.2) (2023-02-20)


### Bug Fixes

* **ci:** use personnal token in order to trigger workflow ([205da7b](https://github.com/bdronneau/memoriesbox/commit/205da7b6bf6ff7b6197029b35a19f1780bceccad))

## [1.0.1](https://github.com/bdronneau/memoriesbox/compare/v1.0.0...v1.0.1) (2023-02-20)


### Bug Fixes

* **ci:** wrong tag detection ([e9870e1](https://github.com/bdronneau/memoriesbox/commit/e9870e1688b252d0829e640fa4a7f0c1bdd6aab0))

## 1.0.0 (2023-02-20)


### Features

* **ci:** add step for docker build and push ([f3c4b4c](https://github.com/bdronneau/memoriesbox/commit/f3c4b4c5fa1143f67d17ee003f891d11c603abcd))
* init application ([6c599bd](https://github.com/bdronneau/memoriesbox/commit/6c599bd01c246d4b2548f95ce466f7228e647a84))
* **sonar:** init file ([0aeeb22](https://github.com/bdronneau/memoriesbox/commit/0aeeb228171b0f49084425bfc11576bdf00aab27))


### Bug Fixes

* **dockerfile:** sonarcloud error ([276553c](https://github.com/bdronneau/memoriesbox/commit/276553cb3ea92749a6cbd89323347caffbd7e96e))
* **docker:** use non root user ([26a32a0](https://github.com/bdronneau/memoriesbox/commit/26a32a0fef00dec0eaeef0fb42306e78dfde5dee))
* **html:** sonarcloud error ([d59bb05](https://github.com/bdronneau/memoriesbox/commit/d59bb05557f6d36494e959cc31f5953aadac183c))
* **lint:** delete trick ([3fb5f19](https://github.com/bdronneau/memoriesbox/commit/3fb5f19d84a42c96178ca4e96f006596b42c2e48))
* **lint:** try trick ([78b3f20](https://github.com/bdronneau/memoriesbox/commit/78b3f20077edfc668976e592d16fc633e7d7af98))
* **lint:** use latest ([8951b5f](https://github.com/bdronneau/memoriesbox/commit/8951b5f80856969e63fc9148944a901ef7a6d0d4))
