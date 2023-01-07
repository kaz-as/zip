# Zip archives processor

## Overview

This is an API for zipping (with output archive streaming) and unzipping (with chunk uploading).

To read documentation and try requests, run `make run` from `zip-go` container. Docs will be at http://localhost:8080/docs

## Generate code from documentation

Run `make gen` from `zip-gen` container.

Do not create files `restapi/configure_zip.go`, `restapi/server.go` by your own: they are deleted by `make gen`.

## Migrations

Should be run from `zip-go` container.

Up:
```bash
make migrate.up
```

Down:
```bash
make migrate.down
```

Status:
```bash
make migrate.status
```

Create new migration with name `new123`:
```bash
make migrate.new migration_name=new123
```
