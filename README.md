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
| `./main.py -i me`         | Run against your own connection          |
| `./main.py -i 1.1.1.1`    | Run against the IP address `1.1.1.1`     |
| `./main.py -i google.com` | Run against the domain name `google.com` |
| `./main.py -i google.com -p` | Run against the domain name `google.com` and lists BGP prefixes |

## Credits
This script runs thanks to the APIs provided by [IP-API](http://ip-api.com) and [HackerTarget](https://hackertarget.com/as-ip-lookup).
