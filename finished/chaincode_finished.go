/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"

)



// SimpleChaincode example simple Chaincode implementation
type StudentManagementChaincode struct {
}


func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(StudentManagementChaincode))
	if err != nil {
		fmt.Printf("Error starting AssetManagementChaincode: %s", err)
	}
}

// Init resets all the things
func (t *StudentManagementChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// create student imformation table

	err := stub.CreateTable("StudentData", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "RollNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Marks", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	if err != nil{
		return nil, errors.New("Unable to create student Data table")
	}


	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *StudentManagementChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		//Initialize the StudentData
		return t.Init(stub, "init", args)
	} else if function == "insert_student" {
		//insert new entry in StudentData
		return t.insert_student(stub, "insert_student" , args)
	} else if function == "delete_student"{
		return t.delete_student(stub,"delete_student",args) 
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *StudentManagementChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read_data(stub ,"read_data",args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

//insert new student
func (t *StudentManagementChaincode) insert_student(stub shim.ChaincodeStubInterface , function string ,args []string)([]byte,error) {

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	var err error

	RollNumber :=args[0]
	Name := args[1]
	Marks :=args[2]

	_, err = stub.InsertRow(
		"StudentData",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: RollNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: Name}},
				&shim.Column{Value: &shim.Column_String_{String_: Marks}},
			},
	})
		if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

fmt.Println("Succesfully added student");
	
msg := "Succesfully added student"
sendMsg := fmt.Sprintf("%s",msg)
return []byte(sendMsg), nil
	

}

//reading StudentData
func (t *StudentManagementChaincode) read_data(stub shim.ChaincodeStubInterface , function string ,args []string)([]byte,error) {

	if(len(args) != 1){
		return nil, errors.New("Wrong args..plz check")

	}
	

RollNumber := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: RollNumber}}
	columns = append(columns, col1)

var err error
	row, err := stub.GetRow("StudentData", columns)

	if(err != nil){
		return nil, errors.New("Failed inserting row.")
	}
	fmt.Println("Details are",row)

	rowString := fmt.Sprintf("%s", row)
	return []byte(rowString), nil


	


}
// delete particular student Entry from StudentData
func (t *StudentManagementChaincode) delete_student(stub shim.ChaincodeStubInterface , function string ,args []string)([]byte,error) {

	if(len(args)!=1){
				return nil,errors.New("Wrong arguments")
	}
	delete_number := args[0]

	var columns [] shim.Column
	col1 := shim.Column{Value : &shim.Column_String_{String_:delete_number}}
	columns = append(columns,col1)

	var err error

	 row,err := stub.GetRow("StudentData",columns)
	if(err != nil){
			return nil,errors.New("RollNumber doenNot exists")
	}
	
	err = stub.DeleteRow(
		"StudentData",
		[]shim.Column{shim.Column{Value: &shim.Column_String_{String_: delete_number}}},
	)

	if(err !=nil){
		return nil,errors.New("Wrong Details")

	}

	Srow :=fmt.Sprintf("%s",row)


	return []byte(Srow),nil
}