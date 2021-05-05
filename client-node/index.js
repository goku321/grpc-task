var PROTO_PATH = __dirname + '/../task/task.proto';
console.log(PROTO_PATH)
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// The protoDescriptor object has the full package hierarchy
var task = protoDescriptor.task.Task;

var tc = new task('localhost:59690', grpc.credentials.createInsecure());

var taskRequest = {name: "new important task"}
tc.Create(taskRequest, function(err, feature) {
    if(err) {
        console.log("error")
    } else {
        console.log("task created")
    }
})

taskRequest = {name: "new important task"}
tc.Get(taskRequest, function(err, feature) {
    if(err) {
        console.log("error")
    } else {
        console.log("retrieve task: ", feature.name)
    }
})
