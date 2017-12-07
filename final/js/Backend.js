// Starts a simple express server upon startup. This server then handles all the backend communication b/w the GUI, web APIs, and local database.
const express = require("express");

console.log("server_start")
var backend = express();

var server = backend.listen(8081, function () {
   var host = server.address().address
   var port = server.address().port
   
   console.log("Backend server listening at http://%s:%s", host, port)
})

// Entrez Stuff
// Base URL: https://eutils.ncbi.nlm.nih.gov/entrez/eutils/

