package structs

type AuthenticateResult struct {
	Status                bool
	Username              string
	WISPrBandwidthMaxDown uint32
	WISPrBandwidthMaxUp   uint32
	IPAddress             string
	Other                 string
}
