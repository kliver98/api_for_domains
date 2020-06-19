package service

import (
	"database/sql"
	"net/http"
	"io/ioutil"
	"log"
	"time"
	"strings"
	"strconv"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/badoux/goscraper"
	"github.com/likexian/whois-go"

	model "../model"
	repository "../repository"
)

const SSL_LABS_URL = "https://api.ssllabs.com/api/v3/analyze?host="
const HTTP = "http://"

type IDomainService interface {
	FetchDomain(db *sql.DB, domain string) (*model.Domain, error)
}

type HostAPI struct {
	Host 		string 		`json:"host,omitempty"`
	Endpoints 	[]ServerAPI `json:"endpoints,omitempty"`
}

type ServerAPI struct {
	IpAddress string `json:"ipAddress,omitempty"`
	Grade     string `json:"grade,omitempty"`
}

func FetchDomain(db *sql.DB, domain string) (model.Domain, error) { //Here formatted to just return domain names, not times
	domainRepository := repository.NewDomainRepository(db)
	//Put into history the domain searched. No matter if found coincidences or not. It's a history of searches
	historyRepository := repository.NewHistoryRepository(db)
	historyRepository.CreateHistory(domain)
	//Looks if the domain already exits, if not domainDB will be nil
	domainsDB,_ := domainRepository.FetchDomain(domain) //domainDb its array class of database -> []DomainDB
	var domainModel model.Domain
	servers := getServers(domain)
	sslGrade := getMinorSslGrade(servers)
	isDown := strconv.FormatBool(isDown(domain))
	logo, title := getLogoAndTitle(domain)
	if len(domainsDB)==0 { //No exist domain stored into database, not searched before
		domainModel.ServersChanged = strconv.FormatBool(false)
		domainModel.PreviousSslGrade = sslGrade
		domainRepository.CreateDomain(domain,sslGrade,sslGrade) //Put domain into database for 1 time
	} else { //Exist into database
		c := moreThan(domainsDB[0].SearchedAt,0)
		if c {
			domainModel.ServersChanged = strconv.FormatBool(domainsDB[0].PreviousSslGrade!=sslGrade) //Not sure, assuming
			domainModel.PreviousSslGrade = domainsDB[0].SslGrade
		} else {
			domainModel.ServersChanged = strconv.FormatBool(false) //Not sure, assuming
			domainModel.PreviousSslGrade = sslGrade
		}
		domainRepository.UpdateDomain(domain,sslGrade,domainModel.PreviousSslGrade) //Updates domain into database
	}
	domainModel.Servers = servers
	domainModel.IsDown = isDown
	domainModel.Title = title
	domainModel.Logo = logo
	domainModel.SslGrade = sslGrade
	return domainModel, nil
}

func moreThan(compare *time.Time, hours int) bool {
	loc, _ := time.LoadLocation("UTC")
	today := time.Now().In(loc)
	comparator := compare.Add(time.Duration(hours) * time.Hour)
	if today.Before(comparator) {
		return false
	} else {
		return true
	}
}

func getDataFromHostAPI(domain string) HostAPI {
	response, err := http.Get(SSL_LABS_URL + domain)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var hostAPI HostAPI
	json.Unmarshal([]byte(responseData), &hostAPI)
	return hostAPI
}

func getCountryAndOwner(domain string) (string,string) {
	result, err := whois.Whois(domain)
	if err != nil {
		log.Fatal(err)
	}
	var owner, country string
	dataOwner := (strings.Split(result, "OrgName:"))
	if len(dataOwner)<2 {
		dataOwner = (strings.Split(result, "org-name:"))
	}
	dataCountry := (strings.Split(result, "Country:"))
	if len(dataCountry)<2 {
		dataCountry = (strings.Split(result, "country:"))
	}
	if len(dataOwner)>1 {
		dataOwner = (strings.Split(dataOwner[1], "\n"))
	}
	owner = strings.Trim(dataOwner[0], " ")

	if len(dataCountry)>1 {
		dataCountry = strings.Split(dataCountry[1], "\n")
	}
	country = strings.Trim(dataCountry[0], " ")

	return country,owner
}

func getServers(domain string) []model.Server {
	endpoints := getDataFromHostAPI(domain).Endpoints
	var servers []model.Server
	for _,v := range endpoints {
		var server model.Server
		server.Address = v.IpAddress
		server.SslGrade = v.Grade
		country,owner := getCountryAndOwner(v.IpAddress)
		server.Country = country
		server.Owner = owner
		servers = append(servers, server)
	}
	return servers
}

func getMinorSslGrade(servers []model.Server) string {
	if len(servers)==0 {
		return model.MIN_SSL_GRADE
	}
	found := model.MAX_SSL_GRADE
	i := model.SSL_GRADE[found]
	for _,v := range servers {
		index := model.SSL_GRADE[v.SslGrade]
		if index>i {
			found = v.SslGrade
			i = index
		}
	}
	return found

}

func getLogoAndTitle(domain string) (string,string) {
	s, err := goscraper.Scrape(HTTP+domain, 5)
	if err != nil {
		return "", ""
	}
	return s.Preview.Icon,s.Preview.Title
}

func isDown(domain string) bool {
	_, err := http.Get(HTTP+domain)
	if err != nil {
	    return true
	} else {
	    return false
	}
}