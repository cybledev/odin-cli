# Odin CLI

ODIN's primary focus is to equip infosec teams with a precise depiction of the internet, enabling them to strengthen their security defences and proactively detect threats within their attack surface.

The Odin CLI provides a simple way to interact with the [Odin API](https://odin.io/docs/api) and access various services related to cybersecurity, certificates, exposed buckets, files and more.

## Installation

require Go 1.21+

```bash
go install github.com/cybledev/odin-cli/cmd/odin@latest
```

Make sure `$GOPATH` is set in `$PATH` to access the odin binary, else add it via `$ bash export PATH=$PATH:$GOPATH/bin`

## Usage

```bash
$odin
Usage:
  odin [command]

Examples:

Search for hosts and services:
        $ odin hosts search --request.query="services.port:443 AND asn.number:AS13335" --X-API-Key=<API-Key>

Search for certificates:
        $ odin certificate search --request.query="certificate.subject_alt_name.dns_names:'cloudflare.com' AND certificate.validity.not_after:'2024-09-20T18:19:24'"

Search exposed buckets:
        $ odin exposed-buckets search --request.query='provider: aws AND file_cat_count.src: [10 TO *]' --request.sort_by="ins_at"  --request.sort_dir="desc"

Search domain names (and store output to a file):
        $ odin dns search --request.domain="cloudflare.com" > cf_domains.txt


Available Commands:
  hosts           Search the list of hosts and their services
  certificate     Search digital certificates associated with websites or online services
  exposed-buckets Search across millions of buckets exposed over the internet
  exposed-files   Search across all the exposed files over the internet
  dns             Search domains/subdomains and whois details
  credits         Get the account related remaining and used credits details
  fields          Get list of searchable fields
  health          Get the status of the API server
  init            Initialize the config file and api keys
  completion      Generate completion script
  help            Help about any command

Flags:
      --config string      Config file path
      --dry-run            Do not send the request to server
      --raw-output         Output raw json response
      --X-API-Key string
  -h, --help               help for odin

Use "odin [command] --help" for more information about a command.

For more information about our offerings and services, visit https://odin.io
```

### Using raw request body

```bash
# use pagination.last to fetch next page
$ odin hosts search
> ....
      "whois_updated_at":"2024-07-22T01:45:37Z"
    }
  ],
  "pagination": {
    "last": [
      1721808680130,
      2
    ],
    "limit":1,
    "start":null,
    "total":10000
  },
  "success":true
}

$ odin hosts search --query='{
  "query": "*",
  "limit": 10,
  "start": [
    1721808680130,
    2
  ]
}'

# using body from a file
$ odin hosts search --query="$(cat request.json)"
```

[<img src="https://github.com/user-attachments/assets/e5e30aeb-3a7a-4f8d-bdbb-443da32fba52" width="50%">](https://github.com/user-attachments/assets/cbc3f6c2-3cfa-4b1d-8ea1-2fa2e3c7c4e0 "Demo Video")

Detailed command usage can be found in [docs](https://github.com/cybledev/odin-cli/tree/main/docs/odin.md)

Make sure to replace `<APIKey>` with your actual [Odin API key from the odin dashboard](https://odin.io/account/profile/api-keys).

Thank you for using the Odin SDK for Go. If you encounter any issues, find a bug, or want to contribute, feel free to open an issue or submit a pull request. Your feedback and contributions are highly appreciated!

For more information about our other projects and services, visit our website at <https://www.odin.io>.
