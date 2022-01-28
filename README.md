# Get Ip Region
This module return the region of an ip address. It has a database(redis) with all the ipv4 cidr address stored in sorted manner. If any request comes, it use binary search to efficiently serach the ip.

The ips' stored in database like below:

```
{
  "values": [
    {
      "cidr": "1.0.0.0/24",
      "country": "AU",
      "first_host": "1.0.0.0",
      "last_host": "1.0.0.0"
    },
    {
      "cidr": "1.0.1.0/24",
      "country": "CN",
      "first_host": "1.0.1.0",
      "last_host": "1.0.1.0"
    },
    {
      "cidr": "1.0.2.0/23",
      "country": "CN",
      "first_host": "1.0.2.0",
      "last_host": "1.0.3.255"
    },
    --------------------------
    --------------------------
    --------------------------
  ]
}
```
