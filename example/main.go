package example

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/kenkoii/webcache"
)

var (
	cache = webcache.NewInMemoryCache()
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getFromCache(w, r)
		} else if r.Method == http.MethodPost {
			setCache(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	// for this example, lets use the request URI as our key
	// in a real app, this should be something else
	key := r.RequestURI

	ce := cache.Get(key)
	// available in cache
	if ce != nil {
		now := time.Now()
		// cache is stale, do something
		if now.Before(ce.Expiration) {
			cache.Invalidate(key)
			return
		}

		// just render cached value
		w.WriteHeader(http.StatusOK)
		w.Write(ce.Data)
	}

	// do original logic here
}

func setCache(w http.ResponseWriter, r *http.Request) {
	// again for this example, lets use the request URI as our key
	// in a real app, this should be something else
	key := r.RequestURI
	// pass some info on header about expiry or maxage

	var maxAgeRegexp = regexp.MustCompile(`maxage(\d+)`)
	cacheHeader := r.Header.Get("cache-control")

	matches := maxAgeRegexp.FindStringSubmatch(cacheHeader)
	if len(matches) == 2 {
		dur, _ := strconv.Atoi(matches[1])
		data, _ := ioutil.ReadAll(r.Body)

		cache.Save(key, data, time.Duration(dur))
	}

}
