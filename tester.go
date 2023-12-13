package main

import (
	"os"
	"template-tester/schema"
)

func GenerateTemplate() (string, error) {
	templateFileName := "input.html"
	templateFilePath := "input.html"

	// ---------------------------------- SAMPLE DATA ----------------------------------

	shipFrom := schema.ShipFrom{
		Address: "123 Main St, Cityville, State",
		FOB:     true,
		SID:     "S123",
	}

	freightChargeTerms := schema.FreightChargeTerms{
		Collect:    true,
		Prepaid:    false,
		ThirdParty: true,
	}

	shipTo := schema.ShipTo{
		Address:    "456 Oak St, Townsville, State",
		FOB:        false,
		CID:        "C123",
		LocationNo: "L456",
	}

	billTo := schema.BillTo{
		Address: "789 Pine St, Villageton, State",
	}

	orderInfo := schema.OrderInfo{
		OrderNumber:           "ORD123",
		Pkgs:                  3,
		Weight:                150.5,
		WeightUnit:            "lbs",
		PalletOrSlip:          true,
		AdditionalShipperInfo: "Additional Shipper Info",
	}

	customerOrderInfo := schema.CustomerOrderInfo{
		OrderInfo:     []schema.OrderInfo{orderInfo},
		TotalPackages: 3,
		TotalWeight:   "150.5 lbs",
	}

	carrierData := schema.CarrierData{
		HandlingUnitQty:      2,
		HandlingUnitType:     "Pallet",
		PackageQty:           5,
		PackageType:          "Box",
		Weight:               300.8,
		WeightUnit:           "lbs",
		HazardousMaterial:    false,
		CommodityDescription: "Electronics",
		NMFC:                 "12345",
		Class:                "85",
	}

	carrierInfo := schema.CarrierInfo{
		CarrierData:          []schema.CarrierData{carrierData},
		TotalHandlingUnitQty: 2,
		TotalPackages:        5,
		TotalWeight:          "300.8 lbs",
	}

	currency := schema.Currency{
		Symbol: "$",
		Code:   "USD",
	}

	codAmount := schema.CODAmount{
		Amount:   100.0,
		Currency: currency,
	}

	feeTerms := schema.FeeTerms{
		Collect:                 true,
		Prepaid:                 false,
		CustomerCheckAcceptable: true,
	}

	trailerLoaded := schema.TrailerLoaded{
		ByShipper: true,
		ByDriver:  false,
	}

	freightCounted := schema.FreightCounted{
		ByShipper:       true,
		ByDriverPallets: false,
		ByDriverPieces:  true,
	}

	// Create an instance of BolStructure and fill it with sample data
	bolData := schema.BolStructure{
		GeneratedDate:      "2023-01-01",
		WarehouseLogo:      "https://upload.wikimedia.org/wikipedia/commons/9/97/Document_icon_%28the_Noun_Project_27904%29.svg",
		ShipFrom:           shipFrom,
		BOLNumber:          "BOL123",
		CarrierName:        "ACME Shipping",
		TrailerNo:          "TR123",
		SealNumbers:        "SEAL456",
		SCAC:               "ABCD",
		ProNo:              "PRO456",
		FreightChargeTerms: freightChargeTerms,
		MasterBOL:          true,
		ShipTo:             shipTo,
		BillTo:             billTo,
		CustomerOrderInfo:  customerOrderInfo,
		Instruction:        "Handle with care",
		CarrierInfo:        carrierInfo,
		CODAmount:          codAmount,
		FeeTerms:           feeTerms,
		TrailerLoaded:      trailerLoaded,
		FreightCounted:     freightCounted,
	}

	// ----------------------------------

	body, err := ParseTemplateFile(templateFileName, templateFilePath, bolData)
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
