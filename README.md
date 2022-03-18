# wedding-api
This boilerplate support Golang version 1.13.x or latest, For futher information, send a email or chat in  teams:
1. Sofyan Saputra
2. Mohammad Risky

## Getting Started
For local development install golang latest version or min version 1.13.x. after install golang next install [air](https://github.com/cosmtrek/air).
```bash
sudo curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```
Install service dependecy
```
docker-compose up -d
```
Running air for up the server
```
air
```
Start your enhancement

## Project Tree
```
.
├── buildspecs
│   ├── orn-comm-svc-dev.yml
│   └── orn-comm-svc-prd.yml
├── database
│   └── migration
│       └── 20200830072040_example.sql
├── deployment
│   ├── deployment.sh
│   └── dockerfiles
│       ├── Dockerfile-comm-svc-dev
│       └── Dockerfile-comm-svc-prd
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
├── README.md
└── src
    ├── bootstrap
    │   └── logger.go
    ├── cmd
    │   ├── cmd.go
    │   ├── http
    │   │   └── http.go
    │   └── migrate
    │       └── db.go
    ├── config
    │   └── config.go
    ├── internal
    │   ├── appctx
    │   │   ├── appctx.go
    │   │   └── rest
    │   │       └── rest.go
    │   ├── consts
    │   │   ├── algorithm.go
    │   │   ├── config.go
    │   │   ├── date_layout.go
    │   │   ├── environment.go
    │   │   ├── message.go
    │   │   ├── paging.go
    │   │   ├── sql.go
    │   │   └── svc.go
    │   ├── entity
    │   │   └── database.go
    │   ├── library
    │   │   ├── aws.go
    │   │   ├── cache.go
    │   │   ├── cache_test.go
    │   │   ├── contract.go
    │   │   ├── database.go
    │   │   ├── decrypter.go
    │   │   ├── decrypter_test.go
    │   │   ├── encrypter.go
    │   │   ├── encrypter_test.go
    │   │   └── validator.go
    │   ├── presentation
    │   │   ├── presentation.go
    │   │   ├── request.go
    │   │   └── response.go
    │   ├── repositories
    │   │   ├── contract.go
    │   │   └── example.go
    │   ├── routes
    │   │   ├── contract.go
    │   │   ├── loader.go
    │   │   └── routes.go
    │   └── ucase
    │       ├── contract
    │       │   └── contract.go
    │       ├── example
    │       │   └── create.go
    │       └── health
    │           └── health.go
    ├── main.go
    ├── middleware
    │   ├── auth.go
    │   ├── contract.go
    │   └── cors.go
    ├── package
    │   ├── awslib
    │   │   ├── aws.go
    │   │   ├── contract.go
    │   │   ├── credential
    │   │   │   └── credential.go
    │   │   ├── sqs
    │   │   │   └── sqs.go
    │   │   └── storage
    │   │       └── storage.go
    │   ├── cache
    │   │   ├── cache.go
    │   │   ├── cache_test.go
    │   │   └── contract.go
    │   ├── database
    │   │   ├── contract.go
    │   │   ├── create_session.go
    │   │   ├── create_session_test.go
    │   │   ├── maria.go
    │   │   ├── maria_master_slave.go
    │   │   ├── maria_master_slave_test.go
    │   │   ├── maria_test.go
    │   │   ├── migration.go
    │   │   └── migration_test.go
    │   ├── logger
    │   │   ├── cfg.go
    │   │   ├── constant.go
    │   │   ├── environment.go
    │   │   ├── environment_test.go
    │   │   ├── formatter.go
    │   │   ├── formatter_test.go
    │   │   ├── logging.go
    │   │   ├── logging_test.go
    │   │   ├── sentry_hook.go
    │   │   ├── sentry_hook_test.go
    │   │   ├── setup.go
    │   │   └── setup_test.go
    │   ├── requester
    │   │   ├── contract.go
    │   │   └── requester.go
    │   ├── restcache
    │   │   └── restcache.go
    │   ├── sentry
    │   │   ├── contract.go
    │   │   └── sentry.go
    │   ├── util
    │   │   ├── array.go
    │   │   ├── array_test.go
    │   │   ├── dump.go
    │   │   ├── dump_test.go
    │   │   ├── env_translator.go
    │   │   ├── env_translator_test.go
    │   │   ├── helper.go
    │   │   ├── helper_test.go
    │   │   ├── log.go
    │   │   ├── parse_date.go
    │   │   ├── parse_date_test.go
    │   │   ├── random_generator.go
    │   │   ├── random_generator_test.go
    │   │   ├── replacer.go
    │   │   ├── replacer_test.go
    │   │   ├── type.go
    │   │   ├── type_test.go
    │   │   ├── util.go
    │   │   └── util_test.go
    │   └── validator
    │       └── validator.go
    └── swagger
        ├── docs.go
        ├── swagger.json
        └── swagger.yaml

```
