package http

// Init will initialize http package by injecting all dependency needed by the package
func Init() *HttpDelivery {
	return &HttpDelivery{}
}
