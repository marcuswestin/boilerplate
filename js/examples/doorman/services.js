var grpc = require('grpc')

var proto = grpc.load({ root:__dirname+'/../../../proto', file:'examples/grpc-echo/echo.proto' }, null, { convertFieldsToCamelCase:true })

var echo = proto.echo
var importer = new echo.Echo(':9090', grpc.credentials.createInsecure())

module.exports = {
	setup: setup,
	echo: echo,
	importer: importer
}

function setup(callback) {
	echo.ping(new echo.PingReq({ ping:'ping' }), function(err, res) {
		if (err) { return callback(err) }
		callback()
	})
}
