var tinytest = require('tinytest')
var grpc = require('grpc')
var fs = require('fs')
var os = require('os')
var csv = require('csvtojson')

tinytest.hijackConsoleLog()
var { deepEqual } = tinytest

setTimeout(function() {
	tinytest.runTests({
		failFast: true
	})
})

test.group('echo', function() {
	var proto
	var addr = 'localhost:9010'
	
	test('load protos', function() {
		proto = grpc.load({ root:__dirname+'/proto', file:'examples/grpc-echo/echo.proto' }, null, { convertFieldsToCamelCase:true })
	})

	test('setup server', function(done) {
		var server = new grpc.Server()
		server.addService(proto.echo.Echo.service, _createEchoServer())
		server.bind(addr, grpc.ServerCredentials.createInsecure())
		server.start()
		done()
	})
	test('connect & ping', function(done) {
		var client = new proto.echo.Echo(addr, grpc.credentials.createInsecure())
		client.ping(new proto.echo.PingReq({ ping:'foo' }), function(err, res) {
			if (err) { return done(err) }
			assert(res.pong == 'pong foo')
			done()
		})
	})
	
	function _createEchoServer() {
		return {
			ping: function(call, callback) {
				callback(null, new proto.echo.PingRes({
					pong: 'pong ' + call.request.ping
				}))
			}
		}
	}
})

test.group('crm', function() {
	var proto
	var client
	var addr = 'localhost:9090'
	var crm
	test('load proto', function() {
		proto = grpc.load({ root:__dirname+'/proto', file:'projectname/apps/crm/crm.proto' }, null, { convertFieldsToCamelCase:true })
		crm = proto.crm
	})
	test('connect & ping', function(done) {
		client = new crm.CRMImporter(addr, grpc.credentials.createInsecure())
		client.ping(new crm.PingReq(), function(err, res) {
			if (err) { return done(err) }
			assert(res.message == 'pong')
			done()
		})
	})
	test('import CSV', function(done) {
		var csvString = fs.readFileSync(__dirname+'/test-data/test-sample.csv').toString()
		
		client.importString(new crm.ImportStringReq(csvString), function(err, res) {
			if (err) { return done(err) }
			csv().fromString(csvString).on('end_parsed', function(expected) {
				// assert(deepEqual(expected, res.results))
				assert(res.newAccounts.length == expected.length)
				done()
			})
		})
	})
})
