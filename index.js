/*
=========================================================================================
File description:
    Filename    : index.js
    Module      : Main Module
	Dependency  : 
			      ./api/api.js
	Description :
			      This file contains Server Initialization code.
	Developed By: PLM LOGIX
=========================================================================================
Date         	Developer Name          Description of Change
19-Apr-2018  	Uday Garud              Initial Version
19-Apr-2018  	Uday Garud 				Added function to start node server
======================================================================================= 
*/
/* jshint node: true */
'use strict';
var express = require('express');
var app = express();
var bodyParser = require('body-parser');
// configure app to use bodyParser()
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

var port = process.env.PORT || 3000;        // set our port

// ROUTES FOR OUR API
var router = express.Router();              // get an instance of the express Router

// middleware to use for all requests
router.use(function(req, res, next) {
    next(); 
});

app.use('/', router);
router.use('/api', require('./api/api').router);

// test route to make sure everything is working
router.get('/', function(req, res) {
    res.json({ message: 'API is working' });   
});

var server = app.listen(port, function () {
   var host = server.address().address
   var port = server.address().port
   console.log("Example app listening at http://%s:%s", host, port)
});