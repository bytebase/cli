# bb - the CLI to interact with your Bytebase instance

`bb` is the CLI to use [Bytebase API](https://api.bytebase.com/) to interact with your Bytebase instance.

## Sample usage

```shell
bb version
```

```shell
# List all projects
bb projects list --url=https://bytebase.example.com --token=xxxx
```

## Docker image

```shell
docker run --rm --name bb bytebase/bb --version
```
