# IP Information Lookup Tool (iPilot)

This Go application takes an IP address or domain name and gathers the following information:
- Location
- Timezone
- Internet Service Provider
- Autonomous System
- Advertised Prefixes

## Running the script
Here are some ways that you can run the script:
| Command                | Description                              |
| ---------------------- | ---------------------------------------- |
| `./iPilot -i me`         | Run against your own connection          |
| `./iPilot -i 1.1.1.1`    | Run against the IP address `1.1.1.1`     |
| `./iPilot -i google.com` | Run against the domain name `google.com` |
| `./iPilot -i google.com -p` | Run against the domain name `google.com` and lists BGP prefixes |

## Support
For support please open an issue on [GitLab](https://gitlab.com/luketainton/iPilot) or email [ipilot@help.tainton.uk](mailto:ipilot@help.tainton.uk).

## Credits
This script runs thanks to the APIs provided by [IP-API](http://ip-api.com) and [HackerTarget](https://hackertarget.com/as-ip-lookup).
