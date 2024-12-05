# gmsv_fakeonline
Fixed gmsv_query

All credits go to this guys: https://github.com/DuckyC/gmsv_query
All I did was add a country detection to block A2S requests from non-CIS countries and fix an issue with server instance not being assigned to a pointer properly.
I implemented this feature to reduce ban possibility, obviously.
However, it does not exclude it at all.
Search for a function "ISFromCISRegion" in netfilter.cpp to change the countries you want to accept requests from.

Also, there was an issue with finding an API service that we can use to detect the IP's country, as we get far more requests/minute
than any API can provide without getting us to pay. So, I created my own, extremely simple API. It's in the ip_api directory.
You will need to download a maxmind GeoLite2 Country mmdb database and place it into the ip_api directory. You can get yours here:
https://dev.maxmind.com/geoip/geolite2-free-geolocation-data/
All you need to do is to install GO on your machine, switch into the ip_api directory and run:
```
go mod tidy
go run main.go
```
That's it, your API is up and running on http://localhost:8080
Now, our module will create requests on it to get the IP's country.

To run this, your machine should have a 32-bit version of libcurl4 installed on it.

Also, for the server to display the fake player count properly, it must have at least one real player on it.
This can be achieved with: https://github.com/Leystryku/leysourceengineclient
There is not much to fix, if you are not dumb - you will do it in no time.
If you are dumb, though, please refrain from using this module at all.

And finally, this modules requiers:
https://github.com/danielga/garrysmod_common.
Git clone it with recuirsive modules flag and use a path to it as a "gmcommon" option when running premake.

Please be adviced, that it is distributed for education purposes only.
Have fun breaking GMOD once again!
