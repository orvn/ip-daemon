package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func getPublicIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func isValidIP(ip string) bool {
	// basic regex to vaidate IP
	re := regexp.MustCompile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
	return re.MatchString(ip)
	// todo: add IPv6
}

// check multiple services
func main() {
	urls := []string{"http://icanhazip.com", "http://ifconfig.me"}

	for {
		for _, url := range urls {
			ip, err := getPublicIP(url)
			if err != nil {
				fmt.Println("Error fetching IP from", url, ":", err)
				continue
			}

			if isValidIP(ip) {
				fmt.Println("Your public IP is:", ip)
				break
			} else {
				fmt.Println("Invalid IP format received from", url)
			}
		}

		time.Sleep(30 * time.Minute) // simple sleep for now
	}
}
