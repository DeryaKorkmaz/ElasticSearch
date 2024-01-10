func Test_Products(t *testing.T) {
	csvFile, err := os.Open(".../scripts/Products.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	jsonFile, err := os.Create(".../scripts/Products.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	indexName := "products"

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		products := Products{}
		for i, value := range record {
			switch header[i] {
			case "ProductID":
				products.ProductID, _ = strconv.Atoi(value)
			case "ProductName":
				products.ProductName = value
			case "ProductType":
				products.ProductType = value
			case "StockStatus":
				products.StockStatus = value
			}
		}

		action := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": indexName,
			},
		}

		jsonData, err := json.Marshal(action)
		if err != nil {
			log.Fatal(err)
		}

		jsonFile.Write(jsonData)
		jsonFile.WriteString("\n")

		jsonData, err = json.Marshal(products)
		if err != nil {
			log.Fatal(err)
		}

		jsonFile.Write(jsonData)
		jsonFile.WriteString("\n")
	}

}
