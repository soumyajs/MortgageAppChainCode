package main

import (
	"encoding/json"
	"errors"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)



// MORTGAGE is a high level smart contract that MORTGAGEs together business artifact based smart contracts
type MORTGAGE struct {

}

// BorrowerDetails is for storing User Details

type BorrowerDetails struct{	
	uid string `json:"uid"`	
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dob string `json:"dob"`
	Email string `json:"email"`
	Phone string `json:"phone"`	
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductName string `json:"productName"`
	LoanAmount string `json:"loanAmount"`
	InterestRate string `json:"interestRate"`	
	DocumentsSubmitted string `json:"documentsSubmitted"`
	SwitchUserFlag string `json:"switchUserFlag"`
	SwitchLenderId string `json:"switchLenderId"`
	SwitchLenderName string `json:"switchLenderName"`
}

// LenderDetails is for storing Lender Details

type LenderDetails struct{	
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductId string   `json:"productId"`
	ProductName string `json:"productName"`
	ProductType string `json:"productType"`
	Rate string `json:"rate"`
	
	
}

// Transaction is for storing transaction Details


// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	

func createLenderDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "productId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "productType", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "rate", Type: shim.ColumnDefinition_STRING, Key: false}
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("LenderDetails", columnDefsTableOne)
}


func createBorrowerDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "uid", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSevenTableOneDef := shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: false}
	columnEightTableOneDef := shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false}
	columnNineTableOneDef := shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTenTableOneDef := shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false}
	columnElevenTableOneDef := shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwelveTableOneDef := shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThirteenTableOneDef := shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourteenTableOneDef := shim.ColumnDefinition{Name: "loanAmount", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFifteenTableOneDef := shim.ColumnDefinition{Name: "interestRate", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixteenTableOneDef := shim.ColumnDefinition{Name: "documentsSubmitted", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSeventeenTableOneDef := shim.ColumnDefinition{Name: "switchUserFlag", Type: shim.ColumnDefinition_STRING, Key: false}
	columnEighteenTableOneDef := shim.ColumnDefinition{Name: "switchLenderId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnNineteenTableOneDef := shim.ColumnDefinition{Name: "switchLenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	
		
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSevenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnEightTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnNineTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnElevenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwelveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThirteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFifteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSeventeenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnEighteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnNineteenTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("BorrowerDetails", columnDefsTableOne)
}
// Init initializes the smart contracts
func (t *MORTGAGE) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

fmt.Println("creating table !!!!! ")
	// Check if table already exists
	_, err := stub.GetTable("LenderDetails")
	fmt.Println("err  table !!!!! &b",err)
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createLenderDetails(stub)
	
	
	
	
	fmt.Println("creating table created!!!!! %b",err)
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
		
	
	
	
	
	
	// Check if table already exists
	_, err = stub.GetTable("BorrowerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createBorrowerDetails(stub)
	
	if err != nil {
		return nil, errors.New("Failed creating BorrowerDetails.")
	}
	
	return nil, nil
}
	

	
//registerUser to register a user
func (t *MORTGAGE) registerLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("no of argumrents  %d",len(args) )

		
	
		lenderId:=args[0]
		lenderName:=args[1]
		productId:=args[2]
		productName:=args[3]
		productType:=args[4]
		rate:=args[5]
		

		
		
		


		// Insert a row
		ok, err := stub.InsertRow("LenderDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: lenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderName}},
				&shim.Column{Value: &shim.Column_String_{String_: productId}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: productType}},
				&shim.Column{Value: &shim.Column_String_{String_: rate}},
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}
//registerUser to register a user
func (t *MORTGAGE) registerBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("no of argumrents  %d",len(args) )


	uid:=args[0]
	gender:=args[1]
	firstName:=args[2]
	lastName:=args[3]
	dob:=args[4]
	email:=args[5]
	phone:=args[6]
	address:=args[7]
	city:=args[8]
	zip:=args[9]
	lenderId:=args[10]
	lenderName:=args[11]
	productName:=args[12]
	loanAmount:=args[13]
	interestRate:=args[14]
	documentsSubmitted:=args[15]
	switchUserFlag:=args[16]
	switchLenderId:=args[17]
	switchLenderName:=args[18]
		
		


		// Insert a row
		ok, err := stub.InsertRow("BorrowerDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: uid}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderName}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: loanAmount}},
				&shim.Column{Value: &shim.Column_String_{String_: interestRate}},
				&shim.Column{Value: &shim.Column_String_{String_: documentsSubmitted}},
				&shim.Column{Value: &shim.Column_String_{String_: switchUserFlag}},
				&shim.Column{Value: &shim.Column_String_{String_: switchLenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: switchLenderName}},
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}


// to get the deatils of a product against productId 
func (t *MORTGAGE) getProductRatesFromLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	productId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this productId
	var columns []shim.Column
	//col3 := shim.Column{Value: &shim.Column_String_{String_: productId}}
	
	//columns = append(columns, col3)
fmt.Println("Hello, World!!!")
	rows, err := stub.GetRows("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the product with id  " + productId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist	
	res2E :=  []*LenderDetails{}
	for row := range rows { 
		newApp:= new(LenderDetails)
		newApp.LenderId = row.Columns[0].GetString_()
		newApp.LenderName = row.Columns[1].GetString_()
		newApp.ProductId = row.Columns[2].GetString_()
		newApp.ProductName = row.Columns[3].GetString_()
		newApp.ProductType = row.Columns[4].GetString_()
		newApp.Rate = row.Columns[5].GetString_()
		if newApp.ProductId == productId{
			res2E=append(res2E,newApp)		
		}		
	}
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a product against lenderId 
func (t *MORTGAGE) getLenderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	lenderId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this lenderId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: lenderId}}
	
	columns = append(columns, col1)
fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the lender with id  " + lenderId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	

	
	res2E :=  LenderDetails{}
	
	res2E.LenderId = row.Columns[0].GetString_()
	res2E.LenderName = row.Columns[1].GetString_()
	res2E.ProductId = row.Columns[2].GetString_()
	res2E.ProductName = row.Columns[3].GetString_()
	res2E.ProductType = row.Columns[4].GetString_()
	res2E.Rate = row.Columns[5].GetString_()
	
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a borrower against email (for internal testing, irrespective of org)
func (t *MORTGAGE) getBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	Email := args[0]
	

	// Get the row pertaining to this Email
	var columns []shim.Column	
	rows, err := stub.GetRows("BorrowerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	res2E := []*BorrowerDetails{}
	for row := range rows {	
		newApp:= new(BorrowerDetails)
		newApp.uid = row.Columns[0].GetString_()
		newApp.Gender = row.Columns[1].GetString_()
		newApp.FirstName = row.Columns[2].GetString_()
		newApp.LastName = row.Columns[3].GetString_()
		newApp.Dob = row.Columns[4].GetString_()
		newApp.Email = row.Columns[5].GetString_()
		newApp.Phone = row.Columns[6].GetString_()
		newApp.Address = row.Columns[7].GetString_()
		newApp.City = row.Columns[8].GetString_()
		newApp.Zip = row.Columns[9].GetString_()
		newApp.LenderId = row.Columns[10].GetString_()
		newApp.LenderName = row.Columns[11].GetString_()
		newApp.ProductName = row.Columns[12].GetString_()
		newApp.LoanAmount = row.Columns[13].GetString_()
		newApp.InterestRate = row.Columns[14].GetString_()
		newApp.DocumentsSubmitted = row.Columns[15].GetString_()
		newApp.SwitchUserFlag = row.Columns[16].GetString_()
		newApp.SwitchLenderId = row.Columns[17].GetString_()
		
		if newApp.Email == Email{
		res2E=append(res2E,newApp)		
		}	
	}
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a product against borrowerId 
func (t *MORTGAGE) fetchBorrowerDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	borrowerId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this borrowerId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
	
	columns = append(columns, col1)
fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("BorrowerDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the borrower with id  " + borrowerId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	fmt.Println("number of columsn :!!!!!!!!!!!!", len(row.Columns))

	
	res2E :=  BorrowerDetails{}

	
	res2E.uid = row.Columns[0].GetString_()
	res2E.Gender = row.Columns[1].GetString_()
	res2E.FirstName = row.Columns[2].GetString_()
	res2E.LastName = row.Columns[3].GetString_()
	res2E.Dob = row.Columns[4].GetString_()
	res2E.Email = row.Columns[5].GetString_()
	res2E.Phone = row.Columns[6].GetString_()
	res2E.Address = row.Columns[7].GetString_()
	res2E.City = row.Columns[8].GetString_()
	res2E.Zip = row.Columns[9].GetString_()
	res2E.LenderId = row.Columns[10].GetString_()
	res2E.LenderName = row.Columns[11].GetString_()
	res2E.ProductName = row.Columns[12].GetString_()
	res2E.LoanAmount = row.Columns[13].GetString_()
	res2E.InterestRate = row.Columns[14].GetString_()
	res2E.DocumentsSubmitted = row.Columns[15].GetString_()
	res2E.SwitchUserFlag = row.Columns[16].GetString_()
	res2E.SwitchLenderId = row.Columns[17].GetString_()
	res2E.SwitchLenderName = row.Columns[18].GetString_()
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
func (t *MORTGAGE) switchLenders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("Hello, World!")
	

	
	switchLenderId := args[0]
	borrowerId := args[1]
	switchLenderName := args[2]
	fmt.Println("Hello, World!!")

	
var colum []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
	
	colum = append(colum, col1)
fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("BorrowerDetails", colum)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the borrower with id  " + borrowerId + "\"}"
		return nil, errors.New(jsonResp)
	}
		
		

	
	
	
		
	//row, err := stub.GetRow("BorrowerDetails", columns)
	
	var columns []*shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: row.Columns[0].GetString_()}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[1].GetString_()}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[2].GetString_()}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[3].GetString_()}}
		col5 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[4].GetString_()}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[5].GetString_()}}
		col7 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[6].GetString_()}}
		col8 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[7].GetString_()}}
		col9 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[8].GetString_()}}
		col10 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[9].GetString_()}}
		col11 := shim.Column{Value: &shim.Column_String_{String_: switchLenderId}}
		col12 := shim.Column{Value: &shim.Column_String_{String_: switchLenderName}}
		col13 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[12].GetString_()}}
		col14 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[13].GetString_()}}
		col15 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[14].GetString_()}}
		col16 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[15].GetString_()}}
		col17 := shim.Column{Value: &shim.Column_String_{String_: "No"}}
		col18 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		col19 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		
		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		columns = append(columns, &col9)
		columns = append(columns, &col10)
		columns = append(columns, &col11)
		columns = append(columns, &col12)
		columns = append(columns, &col13)
		columns = append(columns, &col14)
		columns = append(columns, &col15)
		columns = append(columns, &col16)
		columns = append(columns, &col17)
		columns = append(columns, &col18)
		columns = append(columns, &col19)
		
		row = shim.Row{Columns: columns}
		
	ok, err := stub.ReplaceRow("BorrowerDetails", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}

return nil, nil
}


// Invoke invokes the chaincode
func (t *MORTGAGE) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerLender" {
		t := MORTGAGE{}
		return t.registerLender(stub, args)	
	}else if function == "registerBorrower" {
		t := MORTGAGE{}
		return t.registerBorrower(stub, args)	
	}else if function == "switchLenders" {
		t := MORTGAGE{}
		return t.switchLenders(stub, args)	
	}	

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *MORTGAGE) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getProductRatesFromLender" { 
		t := MORTGAGE{}
		return t.getProductRatesFromLender(stub, args)
	}else if function == "getLenderDetails" {
		t := MORTGAGE{}
		return t.getLenderDetails(stub, args)	
	}else if function == "fetchBorrowerDetails" {
		t := MORTGAGE{}
		return t.fetchBorrowerDetails(stub, args)	
	}else if function == "getBorrower" {
		t := MORTGAGE{}
		return t.getBorrower(stub, args)	
	}

	
	return nil, nil
}

func main() {
	
	err := shim.Start(new(MORTGAGE))
	if err != nil {
		fmt.Printf("Error starting MORTGAGE: %s", err)
	}
} 