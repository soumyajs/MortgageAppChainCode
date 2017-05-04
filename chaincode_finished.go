package main

import (
	"encoding/json"
	"errors"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
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
}

// UserDetails is for storing UserDetails Details

type UserDetails struct{	
	userId string `json:"userId"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}


// GetMile is for storing retreived Get the total Points

type GetMile struct{	
	TotalPoint string `json:"totalPoint"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	



// Init initializes the smart contracts
func (t *MORTGAGE) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("BorrowerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("BorrowerDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "uid", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "loanAmount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "interestRate", Type: shim.ColumnDefinition_STRING, Key: false},		
		&shim.ColumnDefinition{Name: "documentsSubmitted", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "switchUserFlag", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "switchLenderId", Type: shim.ColumnDefinition_STRING, Key: false},		
	})
	if err != nil {
		return nil, errors.New("Failed creating BorrowerDetails.")
	}
	


	// Check if table already exists
	_, err = stub.GetTable("UserDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("UserDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "userId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "password", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "userType", Type: shim.ColumnDefinition_STRING, Key: false},		
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
		
	// setting up the users role
	stub.PutState("user_type1_1", []byte("broker"))
	stub.PutState("user_type1_2", []byte("lender"))
	stub.PutState("user_type1_3", []byte("broker"))
	stub.PutState("user_type1_4", []byte("lender"))	
	
	return nil, nil
}
	

	
//registerBorrower to register a user
func (t *MORTGAGE) registerBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 12 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		uid:=args[0]
		gender:=args[2]
		firstName:=args[3]
		lastName:=args[4]
		dob:=args[5]
		email:=args[6]
		phone:=args[7]
		address:=args[8]
		city:=args[9]
		zip:=args[10]
		lenderId:=args[11]
		lenderName:=args[12]
		productName:=args[13]
		loanAmount:=args[14]
		interestRate:=args[15]
		documentsSubmitted:=args[16]
		switchUserFlag:=args[17]
		switchLenderId	:=args[18]		
		//assignerOrg1, err := stub.GetState(args[19])
		//assignerOrg := string(assignerOrg1)
		
		//createdBy:=assignerOrg
		totalPoint:="0"


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
				&shim.Column{Value: &shim.Column_String_{String_: totalPoint}},
				&shim.Column{Value: &shim.Column_String_{String_: documentsSubmitted}},
				&shim.Column{Value: &shim.Column_String_{String_: switchUserFlag}},
				&shim.Column{Value: &shim.Column_String_{String_: switchLenderId}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}


//get all UserDetails against the uid (depends on org)
func (t *MORTGAGE) getUserDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	uid := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("UserDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*UserDetails{}	
	
	for row := range rows {		
		newApp:= new(UserDetails)
		newApp.userId = row.Columns[0].GetString_()
		newApp.Password = row.Columns[1].GetString_()
		newApp.UserType = row.Columns[2].GetString_()
		
		
		if newApp.userId == uid {
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}







// to get the deatils of a user against uid (for internal testing, irrespective of org)
func (t *MORTGAGE) getBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	uid := args[0]
	

	// Get the row pertaining to this uid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: uid}}
	columns = append(columns, col1)

	row, err := stub.GetRow("BorrowerDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uid + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uid + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	res2E := BorrowerDetails{}
	
	res2E.uid = row.Columns[0].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.Dob = row.Columns[5].GetString_()
	res2E.Email = row.Columns[6].GetString_()
	res2E.Phone = row.Columns[7].GetString_()
	res2E.Address = row.Columns[8].GetString_()
	res2E.City = row.Columns[9].GetString_()
	res2E.Zip = row.Columns[10].GetString_()
	res2E.LenderId = row.Columns[11].GetString_()
	res2E.LenderName = row.Columns[12].GetString_()
	res2E.ProductName = row.Columns[13].GetString_()
	res2E.LoanAmount = row.Columns[14].GetString_()
	res2E.InterestRate = row.Columns[15].GetString_()
	res2E.DocumentsSubmitted = row.Columns[16].GetString_()
	res2E.SwitchUserFlag = row.Columns[17].GetString_()
	res2E.SwitchLenderId = row.Columns[18].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// verify the user is present or not (for internal testing, irrespective of org)
func (t *MORTGAGE) verifyUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	uid := args[0]
	dob := args[1]
	

	// Get the row pertaining to this uid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: uid}}
	columns = append(columns, col1)

	row, err := stub.GetRow("BorrowerDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uid + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uid + "\"}"
		return nil, errors.New(jsonResp)
	}

	userDob := row.Columns[5].GetString_()
	
	res2E := VerifyU{}
	
	if dob == userDob{
		res2E.Result="success"
	}else{
		res2E.Result="failed"
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



// Invoke invokes the chaincode
func (t *MORTGAGE) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerBorrower" {
		t := MORTGAGE{}
		return t.registerBorrower(stub, args)	
	} /*else if function == "addDeleteMile" { 
		t := MORTGAGE{}
		return t.addDeleteMile(stub, args)
	}*/

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *MORTGAGE) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getBorrower" { 
		t := MORTGAGE{}
		return t.getBorrower(stub, args)
	}else if function == "getUserDetails" { 
		t := MORTGAGE{}
		return t.getUserDetails(stub, args)
	} 
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(MORTGAGE))
	if err != nil {
		fmt.Printf("Error starting MORTGAGE: %s", err)
	}
} 