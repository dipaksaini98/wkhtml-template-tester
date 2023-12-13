package main

import (
	"os"
	"strings"
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

	addendumString := `<div class="ql-container ql-snow"><div class="ql-editor"><h2>Hey this is title one</h2><p><strong style="color: rgb(0, 0, 0);">Lorem Ipsum</strong><span style="color: rgb(0, 0, 0);">&nbsp;is simply dummy text of th</span><strong style="color: rgb(0, 0, 0);"><em>e printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but a</em></strong><span style="color: rgb(0, 0, 0);">lso the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of </span></p><p><br></p><pre class="ql-syntax" spellcheck="false">Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
	</pre><p><br></p><ol><li><span style="color: rgb(0, 0, 0);">Operation one </span></li><li><span style="color: rgb(0, 0, 0);">Operation two.                    </span></li><li><span style="color: rgb(0, 0, 0);">operation three</span></li></ol><p><br></p><p><br></p><h2>Hey this is title one</h2><blockquote><strong style="color: rgb(0, 0, 0);">Lorem Ipsum</strong><span style="color: rgb(0, 0, 0);">&nbsp;is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</span></blockquote><p><br></p><ol><li><span style="color: rgb(0, 0, 0);">Operation one </span></li><li><span style="color: rgb(0, 0, 0);">Operation two</span></li><li><span style="color: rgb(0, 0, 0);">operation three</span></li></ol><p><br></p><p><br></p></div></div>`

	if addendumString != "" {

		templateContent, err := os.ReadFile(templateFileName)

		if err != nil {
			return "Failure!", err
		}

		finalHTML := strings.Replace(string(templateContent), "<!-- addendumDataPlaceholder -->", addendumString, -1)

		tempHTMLFileName := "output.html"
		err = os.WriteFile(tempHTMLFileName, []byte(finalHTML), 0644)
		if err != nil {
			return "Failure!", err
		}

		templateFileName = tempHTMLFileName
		templateFilePath = tempHTMLFileName

	}

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
