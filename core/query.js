/*
=========================================================================================
File description:
    Filename    : query.js
    Module      : Module to query chaincode function
	Description :
			      This file contains server side code to execute chaincode functions
	Developed By: PLM LOGIX
=========================================================================================
Date         	Developer Name          Description of Change
23-Apr-2018  	Uday Garud              Initial Version
23-Apr-2018  	Uday Garud              imports and global variables
25-Apr-2018  	Uday Garud              Added function to connect to hyperledger chaincode and invoke chaincode functions
======================================================================================= 
*/
'use strict';
var express = require("express");
var router = express.Router();
var Fabric_Client = require("fabric-client");
var path = require("path");
var util = require("util");
var os = require("os");

//
var fabric_client = new Fabric_Client();

// setup the fabric network
var channel = fabric_client.newChannel("mychannel");
var peer = fabric_client.newPeer("grpc://localhost:7051");
channel.addPeer(peer);

//
var member_user = null;
var store_path = path.join(__dirname, "hfc-key-store");
var tx_id = null;

module.exports = {
  // Export query function with custom payload
  query: function(res,payload) {
    // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    Fabric_Client.newDefaultKeyValueStore({ path: store_path })
      .then(state_store => {
        // assign the store to the fabric client
        fabric_client.setStateStore(state_store);
        var crypto_suite = Fabric_Client.newCryptoSuite();
        var crypto_store = Fabric_Client.newCryptoKeyStore({
          path: store_path
        });
        crypto_suite.setCryptoKeyStore(crypto_store);
        fabric_client.setCryptoSuite(crypto_suite);

        // get the enrolled user from persistence
        return fabric_client.getUserContext("user1", true);
      })
      .then(user_from_store => {
        if (user_from_store && user_from_store.isEnrolled()) {
          console.log("Successfully loaded user1 from persistence");
          member_user = user_from_store;
        } else {
          throw new Error("Failed to get user1.... run registerUser.js");
        }
        const request = payload;
        console.log(request);
        // send the query to the peer
        return channel.queryByChaincode(request);
      })
      .then(query_responses => {
        console.log("Query has completed, checking results");
        // query_responses could have more than one  results if there multiple peers were used as targets
        if (query_responses && query_responses.length == 1) {
          if (query_responses[0] instanceof Error) {
            console.error("error from query = ", query_responses[0]);
          } else {
            res.json(query_responses[0].toString());   
          }
        } else {
          console.log("No payloads were returned from query");
        }
      })
      .catch(err => {
        console.error("Failed to query successfully :: " + err);
        res.json({ message: 'Error Occured while fetching record.' });
      })
  }
};
