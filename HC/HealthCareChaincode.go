package main

import (

    "errors"
    "fmt"
    "encoding/json"
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
fmt.Println("In Init function ")
  var err error

  if(len(args)!=0){
    return nil , errors.New("Incorrect number of arguments..Want 0 arguments");
  }

  //res := RewardPoint{}

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
  fmt.Println("In Invoke func")

  if function =="init" {
    return t.Init(stub,"init",args)
  }else if(function=="AssignPoints"){
    return t.AssignPoints(stub,"AssignPoints",args)
  }else if(function=="RedeemPoints"){
    return t.RedeemPoints(stub,"RedeemPoints",args)
  }


return nil,nil;

}
func (t *HealthCareChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error)  {
  fmt.Println("Query function is running")

  if function == "read"{
    return t.read(stub ,"read",args)
  }
  fmt.Println("Error in Query"+ function)
return nil,nil
}

func (t *HealthCareChaincode) read(stub shim.ChaincodeStubInterface, function string , args []string)([]byte ,error){
  var name ,jsonResp string
  var err error

  if len(args) != 1{
      return nil,errors.New("Incorrect number of arguments")
  }
  name = args[0]
  valAsBytes, err :=stub.GetState(name)

  if err != nil{
    jsonResp ="{\"Error\":\"Failed to get state for "+ name + "\"}"
    return nil,errors.New(jsonResp)
  }

  return valAsBytes , nil

}

func (t *HealthCareChaincode) AssignPoints(stub shim.ChaincodeStubInterface , function string, args []string)([]byte,error)  {

    var err error

    if len(args) !=3 {
          return nil,errors.New("Incorrect numbers of arguments")
    }


   name := args[0]
   value, err := stub.GetState(name)

   if(err !=nil){
      return  t.init_eReward(stub,"eReward",args)
   }else{

     var inputPoints = args[1]
     var inputAssigner = args[2]

     res := RewardPoint{}
     json.Unmarshal(value , &res)


     res.Points = (res.Points + inputPoints)
     res.SignatureAssigner = inputAssigner

     jsonAsBytes, _ := json.Marshal(res)
     err = stub.PutState(name , jsonAsBytes)

     if err != nil {
       return nil, err
     }
   }
   successMsg := fmt.Sprintf("%s",value)
return []byte(successMsg),nil
}

func (t *HealthCareChaincode) RedeemPoints(stub shim.ChaincodeStubInterface , function string, args []string)([]byte,error)  {
return nil,nil
}

func (t *HealthCareChaincode) init_eReward(stub shim.ChaincodeStubInterface,function string, args []string)([]byte,error) {

  if len(args) != 2 {
    return nil,errors.New("Incorrect number of arguments")
  }

  return nil,nil;

}
