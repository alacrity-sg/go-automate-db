# go-automate-db - Configuration

## Description
go-automate-db contains configurations for different use cases in the tool. 
This document will cover all of these in relevant subsections

## Base Configuration
Currently, the configurations are all mutually exclusive. You will not be able to cross
match configurations(eg. provide yaml config but override some fields with environment variables)

### Input Mode - Yaml
The yaml configuration mode serves to be more than a input but as a record keeper that you can commit 
as a permanent configuration add-on similar to terraform resources. 
```yaml
version: 1
sql: 
    type: postgres
    mode: operations
    host: localhost
    port: 5432
    database: postgres
    username: postgres
    password: P@ssw0rd123 
    #password: 
    #  secret:
    #    type: secretsmanager | kubernetes | vault
    #    name: my-secret-name
    #    key: plaintext | my-secret-key 
databases: 
  - name: "my-first-database"
    username: "my-first-user"
    password: "password" # do not provide to auto generate password



```
### Input Mode - Direct CLI Input

### Input Mode - Base Environment variables
The following environment variables are the base configuration required.


| Variable     | Type    | Required | Description                                                                       |
|--------------|---------| -------- |-----------------------------------------------------------------------------------|
| SQL_HOST     | string  | Yes | SQL Host to connect to (eg. localhost)                                            |
| SQL_USERNAME | string | Yes | SQL Username to connect as. Must have enough permissions to run the requested operations |
| SQL_PASSWORD | string | Yes | SQL Password to connect with. |
| SQL_PORT     | string | No | SQL Port to connect to. Defaults to default port of the specified mode (eg. 5432) |
| SQL_DATABASE | string | No | SQL Database to connect to. Defaults to default db if not provided (eg. postgres) |
| SQL_TYPE | string | No | SQL Type to connect as. Defaults to postgres if not specified |

#### Postgres DB Creation Example
```bash
export SQL_HOST="192.168.1.1"
export SQL_USERNAME="my_privileged_user"
export SQL_PASSWORD="password"
export SQL_PORT="5432" # Not required, defaults to 5432 in postgres type
export SQL_DATABASE="postgres" # Not required, defaults to postgres in postgres type
export SQL_TYPE="postgres" # Not required, defaults to postgres
# Extra required env variables for db creation
export SQL_NEW_USERNAME="my_new_user"
export SQL_NEW_PASSWORD="some-password"
export SQL_NEW_DATABASE="my_new_database"
```

