var http = require('http');
var fs = require('fs');

var handleRequest = function(request, response) {
  console.log('Received request for URL: ' + request.url);
  response.writeHead(200);
  response.write('Hello World!\n');
  response.write(new Date().getTime() + '\n')
  try {
    var bytes = fs.readFileSync('/etc/hostname')
    response.write(bytes)
  } catch(err) {
    console.log(err)
    response.write(err.toString())
  }
  response.end('\n\n')
};
var www = http.createServer(handleRequest);
www.listen(30002);
