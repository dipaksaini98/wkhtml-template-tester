package schema

type ShipFrom struct {
	Address string
	FOB     bool
	SID     string
}

type FreightChargeTerms struct {
	Collect    bool
	Prepaid    bool
	ThirdParty bool
}

type ShipTo struct {
	Address    string
	FOB        bool
	CID        string
	LocationNo string
}

type BillTo struct {
	Address string
}

type OrderInfo struct {
	OrderNumber           string
	Pkgs                  int
	Weight                float64
	WeightUnit            string
	PalletOrSlip          bool
	AdditionalShipperInfo string
}

type CustomerOrderInfo struct {
	OrderInfo     []OrderInfo
	TotalPackages int
	TotalWeight   string
}

type CarrierData struct {
	HandlingUnitQty      int
	HandlingUnitType     string
	PackageQty           int
	PackageType          string
	Weight               float64
	WeightUnit           string
	HazardousMaterial    bool
	CommodityDescription string
	NMFC                 string
	Class                string
}

type CarrierInfo struct {
	CarrierData          []CarrierData
	TotalHandlingUnitQty int
	TotalPackages        int
	TotalWeight          string
}

type Currency struct {
	Symbol string
	Code   string
}

type CODAmount struct {
	Amount   float64
	Currency Currency
}

type FeeTerms struct {
	Collect                 bool
	Prepaid                 bool
	CustomerCheckAcceptable bool
}

type TrailerLoaded struct {
	ByShipper bool
	ByDriver  bool
}

type FreightCounted struct {
	ByShipper       bool
	ByDriverPallets bool
	ByDriverPieces  bool
}

type BolStructure struct {
	GeneratedDate      string
	WarehouseLogo      string
	ShipFrom           ShipFrom
	BOLNumber          string
	CarrierName        string
	TrailerNo          string
	SealNumbers        string
	SCAC               string
	ProNo              string
	FreightChargeTerms FreightChargeTerms
	MasterBOL          bool
	ShipTo             ShipTo
	BillTo             BillTo
	CustomerOrderInfo  CustomerOrderInfo
	Instruction        string
	CarrierInfo        CarrierInfo
	CODAmount          CODAmount
	FeeTerms           FeeTerms
	TrailerLoaded      TrailerLoaded
	FreightCounted     FreightCounted
}
