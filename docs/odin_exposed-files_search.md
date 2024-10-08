## odin exposed-files search

returns the matching records according to provided filters

```
odin exposed-files search [flags]
```

### Examples

```
odin-cli exposed-files search --query '{"bucket":bucket:\"lit-link-prd.appspot.com\"}'
```

### Options

```
  -h, --help                      help for search
      --query string              Optional json string for [query]. Search Query
      --request.limit int         (default 1) (default 1)
      --request.query string       (default "*")
      --request.sort_by string    
      --request.sort_dir string   
```

### Options inherited from parent commands

```
      --X-API-Key string   
      --config string      Config file path
      --dry-run            Do not send the request to server
      --raw-output         Output raw json response
```

### SEE ALSO

* [odin exposed-files](odin_exposed-files.md)	 - Search across all the exposed files over the internet

###### Auto generated by spf13/cobra on 9-Aug-2024
