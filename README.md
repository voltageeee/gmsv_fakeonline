# gmsv_fakeonline
Fixed gmsv_query

All credits go to this guys: https://github.com/DuckyC/gmsv_query
All I did was add a country detection to block A2S requests from non-CIS countries.
Search for a function "ISFromCISRegion" in netfilter.cpp to change the countries you want to accept requests from.

Also, there was an issue with finding an API service that we can use to detect the IP's country, as we get far more requests/minute
than any API can provide without getting us to pay. So, I created my own, extremely simple API. It's in the ip_api directory.
All you need to do is to install GO on your machine, switch into the ip_api directory and run:
```
go mod tidy
go run main.go
```
That's it, your API is up and running on http://localhost:8080
Now, our module will create requests on it to get the IP's country.

I implemented this feature to reduce ban possibility, obviously.
However, it does not exclude it at all.

Please be adviced, that it is distributed for education purposes only.
Have fun breaking GMOD once again!
