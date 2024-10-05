package crawler

import (
	"fmt"
	ip2 "github.com/rmv0x11/crawler/internal/storage/ip"
	log "github.com/rmv0x11/crawler/logger"
	"io"
	"math/rand/v2"
	"net"
	"net/http"
	"os"
	"time"
)

//list of ip addresses:
//1.0.0.0/8
//..
//223.0.0.0/8

func Start() {
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				go crawl(GetRandomIP())
			}
		}
	}()
}

func crawl(ip net.IP) {
	l := log.Log

	l.Debug("Start crawl  ip:%v\n", ip.String())
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	//for _, port := range ports {
	//	port := strconv.Itoa(port)
	//	go func(port string) {

	url := "http://" + ip.String() + ":80"
	resp, err := client.Get(url)
	if err != nil {
		l.Warn(err.Error())
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	err = os.Mkdir("ips/"+ip.String(), 0777)
	if err != nil {
		l.Warn(err.Error())
		return
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Warn(err.Error())
		return
	}

	err = os.WriteFile("ips/"+ip.String()+"/80", bytes, 0777)
	if err != nil {
		l.Warn(err.Error())
		return
	}

}

var ports = []int{
	20, 21, 25, 80, 110, 119, 143, 23, 53, 123,
}

func GetRandomIP() net.IP {
	firstOctet := rand.IntN(222) + 1
	secondOctet := rand.IntN(254)
	thirdOctet := rand.IntN(255)
	fourthOctet := rand.IntN(255)

	ip := net.ParseIP(fmt.Sprintf("%d.%d.%d.%d", firstOctet, secondOctet, thirdOctet, fourthOctet))
	if _, ok := ip2.Storage.LookupKey(ip.String()); ok {
		ip := GetRandomIP()
		return ip
	}
	ip2.Storage.Set(ip.String(), nil)

	return ip
}
