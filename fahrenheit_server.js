var PROTO_PATH = __dirname + '/temperature.proto';

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

var COORD_FACTOR = 1e7;

/**
 * Generate a tempature between -20 and 100 degrees Fahrenheit
 */
function GenerateFahrenheitTemp() {
  var min = -20.0;
  var max = 100.0

  return Math.random() * (max - min) + min;
}

/**
 * GetTemp handler. Receives a stream of Celsius temperatures,
 * and responds by sending a new tempature in Fahrenheit.
 * @param {Duplex} call The stream for incoming and outgoing messages
 */
function GetTemp(call) {
  call.on('data', function() {
    var temp = {
      fahrenheit:GenerateFahrenheitTemp(),
      celsius: 0.0
    }
    call.write(temp);
  });
  call.on('end', function() {
    call.end();
  });
}

/**
 * Get a new server with the handler functions in this file bound to the methods
 * it serves.
 * @return {Server} The new server object
 */
function getServer() {
  var server = new grpc.Server();
  server.addService(temperature.FahrenheitService.service, {
    GetTemp: GetTemp
  });
  return server;
}

if (require.main === module) {
  var server = getServer();
  server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
    server.start();
  });
}

exports.getServer = getServer;
