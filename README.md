### tallinn-transport [![Build Status](https://github.com/better-than-yours/tallinn-transport/workflows/actions/badge.svg)](https://github.com/better-than-yours/tallinn-transport/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/better-than-yours/tallinn-transport)](https://goreportcard.com/report/github.com/better-than-yours/tallinn-transport)

### deps

```sh 
$ go mod tidy && go get -u
$ npx npm-check-updates -u
```

### secrets

```sh
VAULT_ROLE_ID=
VAULT_SECRET_ID=
```

### vault

```sh
$ vault auth enable approle
$ vault write auth/approle/role/tallinn-transport-role secret_id_ttl=0 token_policies=common-policy
$ vault read auth/approle/role/tallinn-transport-role/role-id
$ vault write -f auth/approle/role/tallinn-transport-role/secret-id
```