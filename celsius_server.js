var PROTO_PATH = __dirname + '/temperature.proto';

var async = require('async');
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
var client = new temperature.FahrenheitService('localhost:50051',
                                       grpc.credentials.createInsecure());

var COORD_FACTOR = 1e7;

function FahrenheitToCelsius(fahrenheit) {
  return (fahrenheit - 32.0) * 5.0 / 9.0;
}

function runGetTemp(callback) {
  var call = client.GetTemp();
  call.on('data', function(temp) {
    console.log('Received ' + temp.fahrenheit +
            ' degrees Fahreinheit.\nConverting to degrees Celsius.');
    console.log(FahrenheitToCelsius(temp.fahrenheit).toString() + ' degrees Celsius.');
  });

  call.on('end', callback);
  var request = {
    fahrenheit: 0.0,
    celsius: 0.0
  }
  call.write(request);
  call.end();
}

function main() {
  console.log("In main");
  async.series([
    runGetTemp
  ]);
}

if (require.main === module) {
  main();
}

exports.runGetTemp = runGetTemp;