# Zip archives processor

## Overview

This is an API for zipping and unzipping.

To read documentation and try requests, build the [application](cmd/app/main.go) and run it with
`--port <your-port-number>`. Docs are at `localhost:<your-port-number/docs`

## Generate code from swagger

Firstly, install `swagger` from [here](https://github.com/go-swagger/go-swagger/releases/tag/v0.30.3).

Then run the commands below inside the project root.

```bash
swagger generate server -A Zip --spec docs/swagger.yml --exclude-main
rm restapi/configure_zip.go restapi/server.go
```