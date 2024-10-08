## odin domain

search across domains, subdomains and whois data

### Examples

```

$ odin domain search --request.keyword="cloudflare"
$ odin domain subdomain-search --request.domain="cloudflare.com"
$ odin domain whois --domain-name="cloudflare.com"

```

### Options

```
  -h, --help   help for domain
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
* [odin domain count](odin_domain_count.md)	 - Returns the count of domain records based on query
* [odin domain search](odin_domain_search.md)	 - Search domains and subdomains based on the query
* [odin domain subdomain-count](odin_domain_subdomain-count.md)	 - Returns the count of subdomain records based on domain
* [odin domain subdomain-search](odin_domain_subdomain-search.md)	 - Returns the subdomain records based on query
* [odin domain whois](odin_domain_whois.md)	 - Provides extensive details about the domain whois record like name servers, domain status, registrar, etc.
* [odin domain whois-expired](odin_domain_whois-expired.md)	 - Provide info if the domain is expired
* [odin domain whois-historical](odin_domain_whois-historical.md)	 - Provides historical details about the domain whois record
* [odin domain whois-is-registered](odin_domain_whois-is-registered.md)	 - Provides info if the domain is registered or not

###### Auto generated by spf13/cobra on 9-Aug-2024
