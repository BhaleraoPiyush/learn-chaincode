package main

import (

    "errors"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)


type HealthCareChaincode struct {

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
  if(len(args)!=0){
    return nil , errors.New("Incorrect number of arguments..Want 0 arguments");
  }

return nil,nil;
}


func(t *HealthCareChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error) {
return nil,nil;

}
func (t *HealthCareChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error)  {
return nil,nil;
}
