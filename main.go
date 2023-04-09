package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	data "github.com/zapstar/aks-release-tracker/models/latest"
)

var URLs map[string]string = map[string]string{
	"azure.com":         "https://releases.aks.azure.com/data.json",
	"usgovcloudapi.net": "https://releasetrackerprod.blob.core.usgovcloudapi.net/webpage/data.json",
	"chinacloudapi.cn":  "https://releasetrackerprod.blob.core.chinacloudapi.cn/webpage/data.json",
}

func main() {
	for key, url := range URLs {
		client := http.Client{Timeout: time.Second * 10}
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}
		res, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		var response data.AKSReleaseStatusRemote

		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			panic(err)
		}

		// sort the releases by region alias names
		sort.Slice(response.Regions, func(i, j int) bool {
			return *response.Regions[i].Alias < *response.Regions[j].Alias
		})

		// remove release notes, lots of noise in updating those links
		// we can always find the release notes for a build by going to
		// Azure release page
		for _, region := range response.Regions {
			if region.Current != nil {
				region.Current.ReleaseNote = nil
			}
			for _, recent := range region.Recent {
				recent.ReleaseNote = nil
			}
		}

		fileName := fmt.Sprintf("output/%s.data.json", key)
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		storage := data.AKSReleaseStatusStorage(response.Regions)

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		if err := encoder.Encode(storage); err != nil {
			panic(err)
		}
	}
}
