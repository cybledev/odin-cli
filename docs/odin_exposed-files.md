## odin exposed-files

Search across all the exposed files over the internet

### Examples

```
odin exposed-files search --request.query='provider: aws AND name: prod AND ext: ".db"'
```

### Options

```
  -h, --help   help for exposed-files
```

### Options inherited from parent commands

```
      --X-API-Key string   
      --config string      Config file path
      --dry-run            Do not send the request to server
      --raw-output         Output raw json response
```

### SEE ALSO

* [odin](odin.md)	 - 
* [odin exposed-files count](odin_exposed-files_count.md)	 - Returns overall count of exposed bucket files according to provided filters
* [odin exposed-files search](odin_exposed-files_search.md)	 - returns the matching records according to provided filters
* [odin exposed-files summary](odin_exposed-files_summary.md)	 - Returns a summary of exposed bucket files according to provided filters

###### Auto generated by spf13/cobra on 9-Aug-2024
