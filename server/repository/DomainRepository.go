package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"

	repository "github.com/kliver98/api_for_domains/server/repository"
	"github.com/openzipkin/zipkin-go"
)

type DomainRepository struct {
	db     *sql.DB
	tracer *zipkin.Tracer
}

// NewRepository creates a crockoach repository with the necessary dependencies
func NewRepository(db *sql.DB, tracer *zipkin.Tracer) model.Repository {
	return gopherRepository{db: db, tracer: tracer}
}

func (r gopherRepository) CreateGopher(ctx context.Context, g *gopher.Gopher) error {
	return errors.New("method not implemented")
}

func (r gopherRepository) FetchGophers(ctx context.Context) ([]gopher.Gopher, error) {
	return nil, errors.New("method not implemented")
}

func (r gopherRepository) DeleteGopher(ctx context.Context, ID string) error {
	return errors.New("method not implemented")
}

func (r gopherRepository) UpdateGopher(ctx context.Context, ID string, g gopher.Gopher) error {
	return errors.New("method not implemented")
}

func (r gopherRepository) FetchGopherByID(ctx context.Context, ID string) (*gopher.Gopher, error) {
	return nil, errors.New("method not implemented")
}

// Domain defines the properties of a domain to be listed
type Domain struct {
	Servers 			[]Server 	`json:"servers,omitempty"`
	ServersChanged 		bool 		`json:"servers_changed,omitempty"`
	SslGrade 			string 		`json:"ssl_grade,omitempty"`
	PreviousSslGrade 	string 		`json:"previous_ssl_grade,omitempty"`
	Logo 				string 		`json:"logo,omitempty"`
	Title 				string 		`json:"title,omitempty"`
	IsDown 				bool 		`json:"is_down,omitempty"`
	CreatedAt 			*time.Time 	`json:"-"`
}

//Creates a new Domain
func New(servers *[]Server, se bool, image string, age int) *Domain {
	return &Domain{
		ID:    ID,
		Name:  name,
		Image: image,
		Age:   age,
	}
}

//Repository provides access to the domain storage
type Repository interface {
	// CreateDomain saves a given domain
	CreateDomain(ctx context.Context, domain *Domain) error
	// FetchDomain return all domains saved in storage
	FetchDomain(ctx context.Context) ([]Domain, error)
	// DeleteDomain remove domains with given name
	DeleteDomain(ctx context.Context, string name) error
	// UpdateDomain modify domain with given name and given new data
	UpdateDomain(ctx context.Context, name string, domain Domain) error
	// FetchDomainByID returns the domain with given name identifier
	//FetchDomainByID(ctx context.Context, name string) (*Domain, error)
}