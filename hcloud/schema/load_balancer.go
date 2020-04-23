package schema

import "time"

// LoadBalancer represents a Load Balancer in the Hetzner Cloud.
type LoadBalancer struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	PublicNet        LoadBalancerPublicNet    `json:"public_net"`
	PrivateNet       []LoadBalancerPrivateNet `json:"private_net"`
	Location         Location                 `json:"location"`
	LoadBalancerType LoadBalancerType         `json:"load_balancer_type"`
	Protection       LoadBalancerProtection   `json:"protection"`
	Labels           map[string]string        `json:"labels"`
	Created          time.Time                `json:"created"`
	Services         []LoadBalancerService    `json:"services"`
	Targets          []LoadBalancerTarget     `json:"targets"`
	Algorithm        LoadBalancerAlgorithm    `json:"algorithm"`
}

// LoadBalancerPublicNet defines the schema of a Load Balancers
// public network information.
type LoadBalancerPublicNet struct {
	Enabled bool                      `json:"enabled"`
	IPv4    LoadBalancerPublicNetIPv4 `json:"ipv4"`
	IPv6    LoadBalancerPublicNetIPv6 `json:"ipv6"`
}

// LoadBalancerPublicNetIPv4 defines the schema of a Load Balancers public
// network information for an IPv4.
type LoadBalancerPublicNetIPv4 struct {
	IP string `json:"ip"`
}

// LoadBalancerPublicNetIPv6 defines the schema of a Load Balancers public
// network information for an IPv6.
type LoadBalancerPublicNetIPv6 struct {
	IP string `json:"ip"`
}

// LoadBalancerPrivateNet defines the schema of a Load Balancers private network information.
type LoadBalancerPrivateNet struct {
	Network int    `json:"network"`
	IP      string `json:"ip"`
}

// LoadBalancerAlgorithm represents the algorithm of a Load Balancer.
type LoadBalancerAlgorithm struct {
	Type string `json:"type"`
}

// LoadBalancerProtection represents the protection level of a Load Balancer.
type LoadBalancerProtection struct {
	Delete bool `json:"delete"`
}

// LoadBalancerService represents a service of a Load Balancer.
type LoadBalancerService struct {
	Protocol        string                          `json:"protocol"`
	ListenPort      int                             `json:"listen_port,omitempty"`
	DestinationPort int                             `json:"destination_port,omitempty"`
	Proxyprotocol   bool                            `json:"proxyprotocol,omitempty"`
	HTTP            *LoadBalancerServiceHTTP        `json:"http,omitempty"`
	HealthCheck     *LoadBalancerServiceHealthCheck `json:"health_check,omitempty"`
}

// LoadBalancerServiceHTTP represents the http configuration for a LoadBalancerService.
type LoadBalancerServiceHTTP struct {
	CookieName     string `json:"cookie_name,omitempty"`
	CookieLifetime int    `json:"cookie_lifetime,omitempty"`
	Certificates   []int  `json:"certificates,omitempty"`
	RedirectHTTP   bool   `json:"redirect_http,omitempty"`
}

// LoadBalancerServiceHealthCheck represents a service health check configuration.
type LoadBalancerServiceHealthCheck struct {
	Protocol string                              `json:"protocol"`
	Port     int                                 `json:"port"`
	Interval int                                 `json:"interval"`
	Timeout  int                                 `json:"timeout"`
	Retries  int                                 `json:"retries"`
	HTTP     *LoadBalancerServiceHealthCheckHTTP `json:"http,omitempty"`
}

// LoadBalancerServiceHealthCheckHTTP represents a http health check configuration.
type LoadBalancerServiceHealthCheckHTTP struct {
	Domain      string `json:"domain"`
	Path        string `json:"path"`
	Response    string `json:"response"`
	StatusCodes []int  `json:"status_codes"`
	TLS         bool   `json:"tls"`
}

// LoadBalancerTarget represents a target of a Load Balancer.
type LoadBalancerTarget struct {
	Type          string                           `json:"type"`
	Server        *LoadBalancerTargetServer        `json:"server,omitempty"`
	LabelSelector *LoadBalancerTargetLabelSelector `json:"label_selector,omitempty"`
	HealthStatus  []LoadBalancerTargetHealthStatus `json:"health_status,omitempty"`
	Targets       []LoadBalancerTarget             `json:"targets,omitempty"`
}

// LoadBalancerTargetHealthStatus represents a health status of target of a Load Balancer.
type LoadBalancerTargetHealthStatus struct {
	ListenPort int    `json:"listen_port"`
	Status     string `json:"status"`
}

// LoadBalancerTargetServer represents a server target of a Load Balancer.
type LoadBalancerTargetServer struct {
	ID int `json:"id"`
}

// LoadBalancerTargetLabelSelector represents a label selector target of a Load Balancer.
type LoadBalancerTargetLabelSelector struct {
	Selector string `json:"selector"`
}

// LoadBalancerListResponse defines the schema of the response when
// listing Load Balancer.
type LoadBalancerListResponse struct {
	LoadBalancers []LoadBalancer `json:"load_balancers"`
}

// LoadBalancerGetResponse defines the schema of the response when
// retrieving a single Load Balancer.
type LoadBalancerGetResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
}

// LoadBalancerActionTargetRequest defines the schema of the request to
// add or remove a target from a Load Balancer.
type LoadBalancerActionTargetRequest struct {
	Type          string                           `json:"type"`
	Server        *LoadBalancerTargetServer        `json:"server,omitempty"`
	LabelSelector *LoadBalancerTargetLabelSelector `json:"label_selector,omitempty"`
}

// LoadBalancerActionTargetResponse defines the schema of the response when
// adding or removing a target from a Load Balancer.
type LoadBalancerActionTargetResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionAddServiceRequest defines the schema of the request to
// adding a service to a Load Balancer.
type LoadBalancerActionAddServiceRequest struct {
	Protocol        string                          `json:"protocol"`
	ListenPort      int                             `json:"listen_port,omitempty"`
	DestinationPort int                             `json:"destination_port,omitempty"`
	ProxyProtocol   *bool                           `json:"proxyprotocol,omitempty"`
	HTTP            *LoadBalancerServiceHTTP        `json:"http,omitempty"`
	HealthCheck     *LoadBalancerServiceHealthCheck `json:"health_check,omitempty"`
}

// LoadBalancerActionAddServiceResponse defines the schema of the response when
// creating a add service action.
type LoadBalancerActionAddServiceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerDeleteServiceRequest defines the schema of the request to
// delete a service from a Load Balancer.
type LoadBalancerDeleteServiceRequest struct {
	ListenPort int `json:"listen_port"`
}

// LoadBalancerDeleteServiceResponse defines the schema of the response when
// creating a delete_service action.
type LoadBalancerDeleteServiceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerCreateRequest defines the schema of the request to create a LoadBalancer.
type LoadBalancerCreateRequest struct {
	Name             string                 `json:"name"`
	LoadBalancerType interface{}            `json:"load_balancer_type"` // int or string
	Algorithm        *LoadBalancerAlgorithm `json:"algorithm,omitempty"`
	Location         string                 `json:"location,omitempty"`
	NetworkZone      string                 `json:"network_zone,omitempty"`
	Labels           *map[string]string     `json:"labels,omitempty"`
	Targets          []LoadBalancerTarget   `json:"targets,omitempty"`
	Services         []LoadBalancerService  `json:"services,omitempty"`
}

// LoadBalancerCreateResponse defines the schema of the response to
// create a LoadBalancer.
type LoadBalancerCreateResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
	Action       Action       `json:"action"`
}

// LoadBalancerActionChangeProtectionRequest defines the schema of the request to
// change the resource protection of a load balancer.
type LoadBalancerActionChangeProtectionRequest struct {
	Delete *bool `json:"delete,omitempty"`
}

// LoadBalancerActionChangeProtectionResponse defines the schema of the response when
// changing the resource protection of a load balancer.
type LoadBalancerActionChangeProtectionResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerUpdateRequest defines the schema of the request to update a load balancer.
type LoadBalancerUpdateRequest struct {
	Name   string             `json:"name,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// LoadBalancerUpdateResponse defines the schema of the response when updating a load balancer.
type LoadBalancerUpdateResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
}

// LoadBalancerActionChangeAlgorithmRequest defines the schema of the request to
// change the algorithm of a load balancer.
type LoadBalancerActionChangeAlgorithmRequest struct {
	Type string `json:"type"`
}

// LoadBalancerActionChangeAlgorithmResponse defines the schema of the response when
// changing the algorithm of a load balancer.
type LoadBalancerActionChangeAlgorithmResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionUpdateHealthCheckRequest defines the schema of the request to
// updates the health check of a load balancer service.
type LoadBalancerActionUpdateHealthCheckRequest struct {
	ListenPort  int                            `json:"listen_port"`
	HealthCheck LoadBalancerServiceHealthCheck `json:"health_check"`
}

// LoadBalancerActionUpdateHealthCheckResponse defines the schema of the response when
// updating the health check of a load balancer service.
type LoadBalancerActionUpdateHealthCheckResponse struct {
	Action Action `json:"action"`
}
