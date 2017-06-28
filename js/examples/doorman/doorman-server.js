var debug = require('debug')('doorman:server');
var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var lessMiddleware = require('less-middleware');

var port = parseInt(process.env.PORT || '9080')
var app = express()

app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'hbs')

app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')))
app.use(logger('dev'))
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: false }))
app.use(cookieParser())
app.use(lessMiddleware(path.join(__dirname, 'public')))
app.use(express.static(path.join(__dirname, 'public')))

var routes = require('./routes')
var services = require('./services')
routes.setup(app)

services.setup(function(err) {
	if (err) {
		console.log(err.stack || err)
		process.exit(-1)
		return
	}
	app.listen(port, function () {
		console.log('Doorman listening on :'+port)
	})
})
