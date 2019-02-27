/*
=========================================================================================
File description:
    Filename    : api.js
    Module      : Part Module
	Dependency  : 
			      ../core/query.js
				  ../core/invoke.js
	Description :
			      This file contains REST API to execute chaincode function
	Developed By: PLM LOGIX
=========================================================================================
Date         	Developer Name          Description of Change
19-Apr-2018  	Uday Garud              Initial Version
19-Apr-2018   	Uday Garud 				Added function to start node server
19-Apr-2018   	Uday Garud              Added imports and global variables
23-Apr-2018  	Uday Garud              Added API to execute chaincode function to registered part in blockchain
23-Apr-2018  	Uday Garud              Added API to execute chaincode function to add manufacturing information for related part in blockchain
24-Apr-2018  	Uday Garud              Added API to execute chaincode function to add part history for related part in blockchain
25-Apr-2018  	Uday Garud              Added API to execute chaincode function to add engineering details for related part in blockchain
25-Apr-2018  	Uday Garud              Added API to execute chaincode function to add location details for related part in blockchain
26-Apr-2018  	Uday Garud              Added API to execute chaincode function get all information from blockchain for part using item number
======================================================================================= 
*/
'use strict';
var express = require("express");
var router = express.Router();
var querymethods = require("../core/query");
var invokemethods = require("../core/invoke");

//Get all information form blockchain related to part
router.get("/getpartinfo/:ItemNumber", function(req, res) {
  var ItemNumber = req.params.ItemNumber;
  var payload = {
    chaincodeId: "aras",
    fcn: "getPartInfo",
    args: [ItemNumber]
  };
  querymethods.query(res, payload);
});

// Register part in Blockchain
router.post("/addpartinfo", function(req, res) {
  var ItemNumber = req.body.ItemNumber;
  var payload = {
    chaincodeId: "aras",
    fcn: "addPartInfo",
    args: [ItemNumber],
    chainId: "mychannel",
    txId: ""
  };
  invokemethods.invoke(res, payload);
});


//Add Manufacturing information for specified part in blockchain
router.post("/addmanufacturinginfo", function(req, res) {
  var ItemNumber = req.body.ItemNumber;
  var Owner  = req.body.Owner;
  var ManufacturerName = req.body.ManufacturerName;
  var SerialNumber = req.body.SerialNumber;
  var ManufacturingDate = req.body.ManufacturingDate;
  var ManufacturerItemNumber = req.body.ManufacturerItemNumber;

  var payload = {
    chaincodeId: "aras",
    fcn: "addManufacturingInfo",
    args: [ItemNumber,Owner,ManufacturerName,SerialNumber,ManufacturingDate,ManufacturerItemNumber],
    chainId: "mychannel",
    txId: ""
  };
  invokemethods.invoke(res, payload);
});

//Add part history for specified part in blockchain
router.post("/addparthistoryinfo", function(req, res) {
  var ItemNumber = req.body.ItemNumber;
  var CreatedById  = req.body.CreatedById;
  var ModifiedById = req.body.ModifiedById;
  var ModifiedOn = req.body.ModifiedOn;
  var State = req.body.State;
  var Generation = req.body.Generation;
  var MajorRevision = req.body.MajorRevision;

  var payload = {
    chaincodeId: "aras",
    fcn: "addPartHistoryInfo",
    args: [ItemNumber,CreatedById,ModifiedById,ModifiedOn,State,Generation,
    MajorRevision],
    chainId: "mychannel",
    txId: ""
  };
  invokemethods.invoke(res, payload);
});

//Add engineering details for specified part in blockchain
router.post("/addengineeringchanges", function(req, res) {
  var ItemNumber = req.body.ItemNumber;
  var EcoNumber  = req.body.EcoNumber;
  var InitiatedBy = req.body.InitiatedBy;
  var InitiatedOn = req.body.InitiatedOn;
  var State = req.body.State;
  var ReleaseDate = req.body.ReleaseDate;
  var Version = req.body.Version;

  var payload = {
    chaincodeId: "aras",
    fcn: "addEngineeringChanges",
    args: [ItemNumber,EcoNumber,InitiatedBy,InitiatedOn,State,ReleaseDate,
      Version],
    chainId: "mychannel",
    txId: ""
  };
  invokemethods.invoke(res, payload);
});

//Add location details for specified part in blockchain
router.post("/addlocation", function(req, res) {
  var ItemNumber = req.body.ItemNumber;
  var Latitude  = req.body.Latitude;
  var Longitude = req.body.Longitude;
  var Timestamp = req.body.Timestamp;

  var payload = {
    chaincodeId: "aras",
    fcn: "addLocation",
    args: [ItemNumber,Latitude,Longitude,Timestamp],
    chainId: "mychannel",
    txId: ""
  };
  invokemethods.invoke(res, payload);
});

module.exports.router = router;
