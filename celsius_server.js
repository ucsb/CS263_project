var PROTO_PATH = __dirname + '/temperature.proto';

var async = require('async');
var fs = require('fs');
var parseArgs = require('minimist');
var path = require('path');
var _ = require('lodash');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var temperature = grpc.loadPackageDefinition(packageDefinition).temperature;
var client = new routeguide.RouteGuide('localhost:50051',
                                       grpc.credentials.createInsecure());

var COORD_FACTOR = 1e7;

function runSendTempInFahrenheit(callback) {
  var call = client.routeChat();
  call.on('data', function(temp) {
    console.log('Received ' + temp.fahrenheit +
            ' degrees Fahreinheit. Converting to degrees Celsius.');
    call.write();
  });

  call.on('end', callback);
  call.end();
}

function main() {
  async.series([
    runSendTempInFahrenheit
  ]);
}

if (require.main === module) {
  main();
}

exports.runSendTempInFahrenheit = runSendTempInFahrenheit;