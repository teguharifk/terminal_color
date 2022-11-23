package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gookit/color"
)

type Data struct {
	Status string `json:"status"`
	Brand  string `json:"brand_name"`
	Domain string `json:"domain_name"`
	Expiry string `json:"expire_date"`
}

func main() {
	resp, err := http.Get("https://api.licensecity.in/api/getinfo?key=cpanel")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var f Data

	err = json.Unmarshal(body, &f)
	if f.Status == "success" {
		sed("Cpanel::Server::Type::get_max_users()", "0", "/usr/local/cpanel/Cpanel/Config/LoadUserDomains.pm")
		sed("Cpanel::Server::Type::get_max_users()", "0", "/usr/local/cpanel/Cpanel/Config/LoadUserDomains/Tiny.pm")
		sed("Cpanel::Server::Type::get_max_users()", "0", "/usr/local/cpanel/Whostmgr/API/1/Accounts.pm")
		sed("Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1 );", "Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1,1 );", "/usr/local/cpanel/Whostmgr/HTMLInterface/Userlist.pm")
		sed("Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1 );", "Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1,1 );", "/usr/local/cpanel/Whostmgr/Accounts/Suspended.pm")
		sed("Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1 );", "Cpanel::Config::LoadUserDomains::loadtrueuserdomains( undef, 1,1 );", "/usr/local/cpanel/Whostmgr/Resellers/Stat.pm")
		sed("Cpanel::Config::LoadUserDomains::loadtrueuserdomains( \\%DOMAINS, 1 );", "Cpanel::Config::LoadUserDomains::loadtrueuserdomains( \\%DOMAINS, 1,1 );", "/usr/local/cpanel/Whostmgr/Resellers/Stat.pm")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/cpsrvd")
		sed("/etc/domainusers", "/etc/d0mainusers", "/usr/local/cpanel/cpsrvd")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/cpsrvd.so")
		sed("/etc/domainusers", "/etc/d0mainusers", "/usr/local/cpanel/cpsrvd.so")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/whostmgr")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/whostmgr2")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/whostmgr3")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr4")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr5")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr6")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr7")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr9")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr10")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr11")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/whostmgr12")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/whostmgr13")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/whostmgr14")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/.cpsrvd")
		sed("/etc/domainusers", "/etc/d0mainusers", "/usr/local/cpanel/.cpsrvd")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.xml-api")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/xml-api2")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/.whostmgr")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/.whostmgr2")
		sed("/etc/trueuserd0mains", "/etc/trueuserdomains", "/usr/local/cpanel/whostmgr/bin/.whostmgr3")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr4")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr5")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr6")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr7")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr9")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr10")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr11")
		sed("/etc/trueuserdomains", "/etc/trueuserd0mains", "/usr/local/cpanel/whostmgr/bin/.whostmgr12")
		sed("'key'                  => '/etc/trueuserdomains'", "'key'                  => '/etc/trueuserd0mains'", "/usr/local/cpanel/Whostmgr/Transfers/Session/Preflight/RemoteRoot/Analyze.pm")

	} else {
		color.Red.Println("You cannot use it without my permission, notification sent to admin")
	}
}
func sed(old string, new string, file string) {
	filePath := file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {

	} else {
		fileString := string(fileData)
		fileString = strings.ReplaceAll(fileString, old, new)
		fileData = []byte(fileString)
		_ = ioutil.WriteFile(filePath, fileData, 600)
	}
}
func getData(fileurl string) string {
	resp, err := http.Get(fileurl)
	if err != nil {
		fmt.Println("Unable to get Data")
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	data := string(html[:])
	data = strings.TrimSpace(data)
	return data
}


