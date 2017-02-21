package main

import (

    "errors"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
	  "github.com/hyperledger/fabric/core/crypto/primitives"

)

//structure for HealthCareChaincode implementation
type HealthCareChaincode struct {

}

//Structure for Points(RewardPoint)
type RewardPoint struct{

	Points string `json:"Points"`
	Hash string `json:"Hash"`
	SignatureAssigner string `json:"SignatureAssigner"`
  Tx_ID string `json:"Tx_ID_req"`

}

func main()  {
  primitives.SetSecurityLevel("SHA3", 256)
  err := shim.Start(new(HealthCareChaincode))
  if err != nil {
    fmt.Printf("Error starting AssetManagementChaincode: %s", err)
  }
}

//Init function : Table-Creaction
func (t *HealthCareChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error){
  var err error

  if(len(args)!=0){
    return nil , errors.New("Incorrect number of arguments..Want 0 arguments");
  }

//  res := RewardPoint{}

  // err = stub.PutState("UserID",res)
  // if err != nil {
	// 	return nil, err
	// }

  //Table Per User
  err = stub.CreateTable("TransactionTable", []*shim.ColumnDefinition{
    &shim.ColumnDefinition{Name :"Tx_ID",Type :shim.ColumnDefinition_STRING,Key :true},
  //  &shim.ColumnDefinition{Name :"RewardPoint",Type : shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"From",Type :shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"To",Type :shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"Signature_receiver",Type :shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"Signature_assigner",Type :shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"TimeStamp",Type :shim.ColumnDefinition_STRING,Key :false},
    &shim.ColumnDefinition{Name :"Reason",Type :shim.ColumnDefinition_STRING,Key :false},
  })

if err != nil {
  return nil,errors.New("Failed to create table")
}


  //Set adminCert
  adminCert, err := stub.GetCallerMetadata()
  if err!=nil{
    return nil, errors.New("Not getting proper metadata")
  }

  if len(adminCert)== 0{
    return nil,errors.New("Invalid Admin certificate:")
  }

stub.PutState("admin",adminCert)






return nil,nil;
}


func(t *HealthCareChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error) {
return nil,nil;

}
func (t *HealthCareChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error)  {
return nil,nil;
}
