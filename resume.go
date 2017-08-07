// Package resume is a résumé formatted as a godoc
//
// Patrick Moroney
//
// Senior Go Developer with extensive experience with network, database, and system administration.
// Specialized knowledge of domain registries and registrars, wireless networks, and laser safety.
package resume

const (
	Name     = "Patrick Moroney"
	Email    = "pat@pat.email"
	Phone    = "720.663.0025"
	Location = "Denver, CO or Remote"
)

type company struct{}
type school struct{}

// I am making NameDotCom a function so that it shows up first.
// Otherwise it sorts in alphabetical order.

// Name.com, Sr. Software Engineer: 2011 - Present
//
// Designed and implemented many services, including the order processing, domain lifecycle management, and DNS services.
//
// OrderProcessor
//
// The order processor handles every purchase made at Name.com and processed over $32M in transactions in 2016.
// It integrates with almost every other service and many 3rd party APIs.
//
// DNS
//
// The DNS service hosts over 1 million zones.
// Created a feature to resolve ANAME records into A and AAAA records automatically and keep them updated.
// Designed a system to monitor propagation of updates and verify consistency.
//
// DomainLifecycleManagement
//
// The domain lifecycle management service handles expiration and deletion of domains.
// This is a vital service since domains typically automatically renew at the registry which can be costly if not deleted on time.
// It also handles various reconciliations with the registries to verify accurate inventory.
func NameDotCom() company {}

// Laser Institute of America, IT Manager: 2006 - 2009
//
// Managed employees and contractors to meet the IT needs of a scientific non-profit organization.
// Developed a custom web-based customer management database to replace their previous
// database, iMIS. Developed a scientific calculator application that evaluates the
// hazards presented by lasers, based on the ANSI Z136.1 standard. Created a database
// and website for management of scientific conferences. Wrote a content management
// system for the website, allowing the marketing department to edit pages on the fly.
type LaserInstitute company

// Pure Connection, Network Operations Manager: 2000 - 2004
//
// Designed, built, and maintained a multihomed wireless network to provide Internet service
// to more then half of Orange County Florida. Wrote a system that implements wireless hot-spots,
// allowing users to only connect to the Internet after paying a small fee.
type PureConnection company

// Self Employed, IT Consultant: 2004 - 2013
//
// Clients include Florida Hospital Assoc. and Laser Institute of America
type SelfEmployed company

// B.S. Computational Mathematics, University of Central Florida
type UniversityOfCentralFlorida school

// Thanks for reading my resume's source too!
