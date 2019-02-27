/*
=========================================================================================
File description:
    Filename    : aras.go
    Module      : Hyperledger chaincode
	Description :
			      This file contains hyperledger chaincode written in GO language.
	Developed By: PLM LOGIX
================================================================================================
  Date         	Developer Name          Description of Change
19-Apr-2018  	Nisha Mane              Initial Version
19-Apr-2018  	Nisha Mane              Added global variables(struct)
19-Apr-2018  	Nisha Mane              Added function invoke a chaincode function (Invoke)
20-Apr-2018  	Nisha Mane              Added chaincode function to register part and add part in stub (addPartInfo)
20-Apr-2018  	Nisha Mane              Added chaincode function to add manufacturing information(addManufacturingInfo)
20-Apr-2018  	Nisha Mane              Added chaincode function to add part history(addPartHistoryInfo) 
24-Apr-2018  	Nisha Mane              Added chaincode function to add engineering details(addEngineeringChanges)
24-Apr-2018  	Nisha Mane              Added chaincode function to add location details (addLocation) 
24-Apr-2018  	Nisha Mane              Added chaincode function to get all information for part using item number(getPartInfo)
=================================================================================================*/

package main

import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type ArasDappChaincode struct {
}

type LocationDetails struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Timestamp string `json:"timestamp"`
}

type ManufacturingInfo struct {
 Owner                   string `json:"owner"`
 ManufacturerName	     string `json:"manufacturerName"`
 SerialNumber            string `json:"serialNumber"`
 ManufacturingDate       string `json:"manufacturingDate"`
 ManufacturerItemNumber  string `json:"manufacturingItemNumber"`
}

type PartHistoryInfo struct {
 CreatedById   string `json:"createdById"`
 ModifiedById  string `json:"modifiedById"`
 ModifiedOn    string `json:"modifiedOn"`
 State         string `json:"state"`
 Generation    string `json:"generation"`
 MajorRevision string `json:"majorRevision"`
}

type EngineeringChanges struct {
	EcoNumber string `json:"ecoNumber"`
	InitiatedBy string `json:"initiatedBy"`
	InitiatedOn string `json:"initiatedOn"`
	State string `json:"state"`
	ReleaseDate string `json:"releaseDate"`
	Version string `json:"version"`
}

type PartInfo struct {
	ItemNumber           string               `json:"itemNumber"`
	ManufacturingDetails ManufacturingInfo    `json:"manufacturingDetails"`
	ParyHistory          []PartHistoryInfo    `json:"partHistory"`
	EngineeringDetails   []EngineeringChanges `json:"engineeringDetails"`
	Location             []LocationDetails    `json:"location"`
}

func main() {	
	err := shim.Start(new(ArasDappChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *ArasDappChaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is entry point to invoke a chaincode function
func (t *ArasDappChaincode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	if function == "addPartInfo" {
		return t.addPartInfo(APIstub,args)
	} else if function == "addManufacturingInfo" {
		return t.addManufacturingInfo(APIstub,args)	
	} else if function == "addPartHistoryInfo" {
		return t.addPartHistoryInfo(APIstub,args)	
	} else if function == "addEngineeringChanges" {
		return t.addEngineeringChanges(APIstub,args)	
	} else if function == "addLocation" {
		return t.addLocation(APIstub,args)	
	}else if function == "getPartInfo" { 
		return t.getPartInfo(APIstub, args)
	} 
	return shim.Error("Invalid Smart Contract function name.")
}


// Function  to registered part in stub 
func (t *ArasDappChaincode) addPartInfo(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    var err error		
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}   
	partInfo := PartInfo{}
	partInfo.ItemNumber = args[0]
	partInfoAsbytes,_ := json.Marshal(partInfo)
	err = stub.PutState(args[0],partInfoAsbytes)
	if err != nil {
		return shim.Error("Error")
	}
	return shim.Success(partInfoAsbytes)
}

// Function  to add manufacturing information for registered part
func (t *ArasDappChaincode) addManufacturingInfo(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error		
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5.")
	}   
	partInfoAsbytes, err1 := stub.GetState(args[0])
	if err1 != nil {
		return shim.Error("Error")
	}
	manufacturingInfo := ManufacturingInfo{}
	manufacturingInfo.Owner = args[1]
    manufacturingInfo.ManufacturerName = args[2]
	manufacturingInfo.SerialNumber = args[3]
	manufacturingInfo.ManufacturingDate = args[4]
	manufacturingInfo.ManufacturerItemNumber = args[5]
	
	partInfo := PartInfo{}
	json.Unmarshal(partInfoAsbytes,&partInfo)
	partInfo.ManufacturingDetails = manufacturingInfo
	partInfoBytes,_ := json.Marshal(partInfo)
	err = stub.PutState(args[0],partInfoBytes)
	if err != nil {
		return shim.Error("Error")
	}
	return shim.Success(partInfoBytes)
}

// Function  to add part history for registered part
func (t *ArasDappChaincode) addPartHistoryInfo(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    var err error		
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}   
	partInfoAsbytes, err1 := stub.GetState(args[0])
	if err1 != nil {
		return shim.Error("Error")
	}
	partHistoryInfo := PartHistoryInfo{}
	partHistoryInfo.CreatedById = args[1]
    partHistoryInfo.ModifiedById = args[2]
	partHistoryInfo.ModifiedOn = args[3]
	partHistoryInfo.State = args[4]
	partHistoryInfo.Generation = args[5]
	partHistoryInfo.MajorRevision = args[6]
	
	partInfo := PartInfo{}
	json.Unmarshal(partInfoAsbytes,&partInfo)
	partInfo.ParyHistory =  append(partInfo.ParyHistory,partHistoryInfo)
	partInfoBytes,_ := json.Marshal(partInfo)
	err = stub.PutState(args[0],partInfoBytes)
	if err != nil {
		return shim.Error("Error")
	}
	return shim.Success(partInfoBytes)
}


// Function  to engineering details for registered part
func (t *ArasDappChaincode) addEngineeringChanges(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    var err error		
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}   
	partInfoAsbytes, err1 := stub.GetState(args[0])
	if err1 != nil {
		return shim.Error("Error")
	}
	engineeringChanges := EngineeringChanges{}
	engineeringChanges.EcoNumber = args[1]
    engineeringChanges.InitiatedBy = args[2]
	engineeringChanges.InitiatedOn = args[3]
	engineeringChanges.State = args[4]
	engineeringChanges.ReleaseDate = args[5]
	engineeringChanges.Version = args[6]
	
	partInfo := PartInfo{}
	json.Unmarshal(partInfoAsbytes,&partInfo)
	partInfo.EngineeringDetails =  append(partInfo.EngineeringDetails,engineeringChanges)
	partInfoBytes,_ := json.Marshal(partInfo)
	err = stub.PutState(args[0],partInfoBytes)
	if err != nil {
		return shim.Error("Error")
	}
	return shim.Success(partInfoBytes)
}
// Function  to add location details for registered part
func (t *ArasDappChaincode) addLocation(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    var err error		
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4.")
	}   
	partInfoAsbytes, err1 := stub.GetState(args[0])
	if err1 != nil {
		return shim.Error("Error")
	}
	locationDetails := LocationDetails{}
	locationDetails.Latitude = args[1]
    locationDetails.Longitude = args[2]
	locationDetails.Timestamp = args[3]
	
	partInfo := PartInfo{}
	json.Unmarshal(partInfoAsbytes,&partInfo)
	partInfo.Location =  append(partInfo.Location,locationDetails)
	partInfoBytes,_ := json.Marshal(partInfo)
	err = stub.PutState(args[0],partInfoBytes)
	if err != nil {
		return shim.Error("Error")
	}
	return shim.Success(partInfoBytes)
}

// Function to get all part information
func (t *ArasDappChaincode) getPartInfo(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		
		return shim.Error("Incorrect number of arguments. Expecting 1(Item numbersss)")
	}
	valAsbytes, err := stub.GetState(args[0])	
	if err != nil {	
		return shim.Error("Error")
	}
	return shim.Success(valAsbytes)
}