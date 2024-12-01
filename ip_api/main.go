package main

import (
	"fmt"
	"net/http"
	"net/netip"

	"github.com/oschwald/maxminddb-golang/v2"
)

func GetCountryFromIP(ipaddr string) string {
	db, err := maxminddb.Open("GeoLite2-Country.mmdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	addr := netip.MustParseAddr(ipaddr)

	var record struct {
		Country struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
	}

	err = db.Lookup(addr).Decode(&record)
	if err != nil {
		panic(err)
	}

	countryName := record.Country.Names["en"]
	return countryName
}

func main() {
	http.HandleFunc("/getcountry", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		country := GetCountryFromIP(ip)
		fmt.Fprintf(w, "%s\n", country)
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
