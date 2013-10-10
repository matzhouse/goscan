package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type scanner struct {
	nodes []node
}

type node struct {
	name  string
	scans []scan
}

type scan struct {
	scantype int
	hostname string
	ip       string
	ports    []int
}

func main() {

	web := []int{80}
	//std := []int{21, 23, 53, 80, 110, 443}

	scn := Newscanner()

	n1 := Newnode()

	n1.name = "webscans"

	scan1 := scan{
		scantype: 1,
		hostname: "http://www.google.com",
		ip:       "213.155.151.181",
		ports:    web,
	}

	scan2 := scan{
		scantype: 1,
		hostname: "http://www.facebook.com",
		ip:       "213.155.151.181",
		ports:    web,
	}

	n1.addScan(scan1)
	n1.addScan(scan2)

	scn.Addnode(n1)

	scn.Scan()

	/*
		var keys []string

		for key := range list {
			keys = append(keys, key)
		}

		var wg sync.WaitGroup

		for _, key := range keys {

			fmt.Println(key)

			wg.Add(1)

			// pass the scanrequest element to the go routine to do it concurrently
			go func(req scanrequest) {
				defer wg.Done()

				fmt.Println("Scanning.. ", req.hostname)

				res, err := req.scanhttp()

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(res.StatusCode)

				res.Body.Close()

			}(list[key])
		}

		wg.Wait()
	*/

}

// create new scanner - returns a nice pointer
func Newscanner() scanner {
	return scanner{}
}

// create new node - returns a nice pointer
func Newnode() node {
	return node{}
}

func (s *scanner) Addnode(n node) {
	s.nodes = append(s.nodes, n)
}

func (n *node) addScan(s scan) {
	n.scans = append(n.scans, s)
	return
}

func (s scanner) Scan() {

	var wg sync.WaitGroup

	for _, value := range s.nodes {
		for _, sc := range value.scans {
			wg.Add(1)
			go func(req scan) {

				defer wg.Done()

				switch req.scantype {
				case 1:
					res, err := httpscan(req.hostname)
					if err != nil {
						log.Fatal("scan failed")
					}
					fmt.Println(req.hostname, ": ", res.StatusCode)
					break
				}
			}(sc)

		}

		// waiting for completion
		fmt.Println("Waiting for scans to complete..")
		wg.Wait()
	}

}

func httpscan(url string) (res *http.Response, err error) {
	res, err = http.Get(url)
	return
}
