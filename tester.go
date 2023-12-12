package main

import (
	"os"
	"time"
)

func GenerateTemplate() (string, error) {
	templateFileName := "input.html"
	templateFilePath := "input.html"

	// -------------------------------------- structure --------------------------------------

	// FreightCounted represents schema for FreightCounted
	type FreightCounted struct {
		ByShipper       bool
		ByDriverPallets bool
		ByDriverPieces  bool
	}

	// TrailerLoaded represents schema for TrailerLoaded
	type TrailerLoaded struct {
		ByShipper bool
		ByDriver  bool
	}

	type Currency struct {
		Code   string
		Symbol string
	}

	// CODAmount represents schema for CODAmount
	type CODAmount struct {
		Currency Currency
		Amount   string
	}

	// FeeTerms represents schema for FeeTerms
	type FeeTerms struct {
		Prepaid                 bool
		Collect                 bool
		CustomerCheckAcceptable bool
	}

	// CarrierData represents schema for Carrier data
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

	// CarrierInfo represents schema for CarrierInfo
	type CarrierInfo struct {
		CarrierData          []CarrierData
		TotalPackages        int
		TotalWeight          string
		TotalHandlingUnitQty int
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
		ClientName            string
		Status                string
		OrderNumber           string
		Pkgs                  int
		Weight                float64
		WeightUnit            string
		PalletOrSlip          bool
		AdditionalShipperInfo string
		ShipTo                *ShipTo
		BillTo                *BillTo
	}

	// CustomerOrderInfo represents schema for CustomerOrderInfo
	type CustomerOrderInfo struct {
		OrderInfo     []OrderInfo
		TotalPackages int
		TotalWeight   string
	}

	// FreightChargeTerms represents schema for FreightChargeTerms
	type FreightChargeTerms struct {
		Prepaid    bool
		Collect    bool
		ThirdParty bool
	}

	type ShipFrom struct {
		Address string
		SID     string
		FOB     bool
	}

	type BillOfLading struct {
		WarehouseLogo      string
		Code               int64
		Url                string
		BOLNumber          string
		ShipFrom           *ShipFrom
		ShipTo             *ShipTo
		BillTo             *BillTo
		CarrierName        string
		TrailerNo          string
		SealNumbers        string
		SCAC               string
		ProNo              string
		FreightChargeTerms *FreightChargeTerms
		MasterBOL          bool
		Instruction        string
		CustomerOrderInfo  *CustomerOrderInfo
		CarrierInfo        *CarrierInfo
		TotalShipmentValue string
		FOB                string
		CODAmount          *CODAmount
		FeeTerms           *FeeTerms
		TrailerLoaded      *TrailerLoaded
		FreightCounted     *FreightCounted
		IsCancelled        bool
		CancelledAt        *time.Time
		CancellationReason string
		GeneratedDate      string

		CreatedAt *time.Time
		UpdatedAt *time.Time
	}

	// --------------------------------------

	// SampleObject represents a sample instance of BillOfLading
	var SampleObject = BillOfLading{
		WarehouseLogo: "sample_logo_url",
		Code:          12345,
		Url:           "sample_url",
		BOLNumber:     "BOL123456",
		ShipFrom: &ShipFrom{
			Address: "Sample Ship From Address",
			SID:     "SID123",
			FOB:     true,
		},
		ShipTo: &ShipTo{
			Address:    "Sample Ship To Address",
			FOB:        false,
			CID:        "CID789",
			LocationNo: "Location123",
		},
		BillTo: &BillTo{
			Address: "Sample Bill To Address",
		},
		CarrierName: "Sample Carrier",
		TrailerNo:   "TR123",
		SealNumbers: "Seal123",
		SCAC:        "SCAC123",
		ProNo:       "PRO456",
		FreightChargeTerms: &FreightChargeTerms{
			Prepaid:    true,
			Collect:    false,
			ThirdParty: true,
		},
		MasterBOL:   true,
		Instruction: "Handle with care",
		CustomerOrderInfo: &CustomerOrderInfo{
			OrderInfo: []OrderInfo{
				{
					ClientName:            "Client1",
					Status:                "Shipped",
					OrderNumber:           "Order123",
					Pkgs:                  5,
					Weight:                150.5,
					WeightUnit:            "lbs",
					PalletOrSlip:          true,
					AdditionalShipperInfo: "Fragile",
					ShipTo: &ShipTo{
						Address:    "Sample Ship To Address",
						FOB:        false,
						CID:        "CID789",
						LocationNo: "Location123",
					},
					BillTo: &BillTo{
						Address: "Sample Bill To Address",
					},
				},
			},
			TotalPackages: 5,
			TotalWeight:   "150.5 lbs",
		},
		CarrierInfo: &CarrierInfo{
			CarrierData: []CarrierData{
				{
					HandlingUnitQty:      2,
					HandlingUnitType:     "Pallet",
					PackageQty:           10,
					PackageType:          "Box",
					Weight:               100.2,
					WeightUnit:           "lbs",
					HazardousMaterial:    false,
					CommodityDescription: "Sample Commodity",
					NMFC:                 "NMFC123",
					Class:                "ClassA",
				},
			},
			TotalPackages:        10,
			TotalWeight:          "200.5 lbs",
			TotalHandlingUnitQty: 2,
		},
		TotalShipmentValue: "5000 USD",
		FOB:                "FOB Sample",
		CODAmount: &CODAmount{
			Currency: Currency{
				Code:   "USD",
				Symbol: "$",
			},
			Amount: "100.00",
		},
		FeeTerms: &FeeTerms{
			Prepaid:                 true,
			Collect:                 false,
			CustomerCheckAcceptable: true,
		},
		TrailerLoaded: &TrailerLoaded{
			ByShipper: true,
			ByDriver:  false,
		},
		FreightCounted: &FreightCounted{
			ByShipper:       true,
			ByDriverPallets: false,
			ByDriverPieces:  true,
		},
		IsCancelled:        false,
		CancelledAt:        nil, // Replace with actual time if canceled
		CancellationReason: "Not applicable",
		GeneratedDate:      "2023-01-01", // Replace with actual generated date
		CreatedAt:          nil,          // Replace with actual time
		UpdatedAt:          nil,          // Replace with actual time
	}

	body, err := ParseTemplateFile(templateFileName, templateFilePath, SampleObject)
	if err != nil {
		return "Failure!", err
	}
	buff, err := GeneratePDF(body)
	if err != nil {
		return "Failure!", err
	}

	pdfFilePath := "output.pdf"

	err = os.WriteFile(pdfFilePath, buff.Bytes(), 0644)
	if err != nil {
		return "Failure!", err
	}

	return "Success!", nil
}
