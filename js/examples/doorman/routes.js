var services = require('./services')

module.exports = {
	setup: setup
}

function setup(app) {
	app.use('/', index)
	app.use('/users', users)
}

var { crm, importer } = services

module.exports = {
	setup: setup
}

function setup(app) {
	app.get('/', (req, res) => {
		res.render('index', { title:'PROJECTNAME' })
	})
	app.get('/welcome/:name', (req, res, next) => {
		res.render('index', { title:req.params.name })
	})
	app.get('/admin', (req, res) => {
		if (!handleAdminAuth(req, res)) { return }
		res.render('admin')
	})

	// Catch-all
	app.use((err, req, res, next) => {
		handle404(err, req, res, next)
	})
	app.use((err, req, res, next) => {
		handleError(err, req, res, next)
	})
}

function handle404(err, req, res, next) {
	if (err) { return next(err) }
	var err = new Error('Not Found')
	err.status = 404
	next(err)
}

function handleError(err, req, res, next) {
	if (err.userError) {
		res.status(400)
		res.render('error', {
			message: err.userError.message
		})
		
	} else {
		var isDev = true || (req.app.get('env') === 'development')
		var stack = isDev ? (err.Stack || err.stack) : ''
		var message = isDev ? (err.Message || err.message) : 'Oops! Something went wrong. Please try again!'
		res.status(err.status || 500)
		res.render('error', {
			message: message,
			stack: stack
		})
	}
}

function submitInvite(req, res, next) {
	var d = req.body
	var csvStr = ''+
		'externalAccountId,firstName,middleNames,lastName,phoneNumber,emailAddress,lastBillDate,lastBillAmountPaid,nextBillDate\n'+
		[d.externalAccountId, d.firstName, '', d.lastName, d.phoneNumber, d.emailAddress, '', '', ''].join(',')
	importer.importString(new crm.ImportStringReq(csvStr), checkErr(next, function(rpcRes) {
		res.render('admin', {
			statusMessage: 'Invite sent!'
		})
	}))
}

function renderEnroll(req, res, next, secret) {
	importer.getEnrollInfo(new crm.GetEnrollInfoReq(secret), checkErr(next, function(rpcRes) {
		res.render('enroll', {
			secret: secret,
			enrollInfo: rpcRes.enrollInfo
		})
	}))
}

function submitEnrollment(req, res, next) {
	var data = req.body
	var enrollReq = new crm.EnrollReq({
		secret: data.secret,
		enrollInfo: new crm.EnrollInfo({
			fullName: data.fullName,
			phoneNumber: data.phoneNumber,
			emailAddress: data.emailAddress,
		}),
	})
	importer.enroll(enrollReq, checkErr(next, function(rpcRes) {
		res.render('enrolled')
	}))
}

function renderAccount(req, res, next, secret) {
	importer.getAccountInfo(new crm.GetAccountInfoReq(secret), checkErr(next, function(rpcRes) {
		res.render('account', {
			statusMessage: 'TODO: Account page CRM integration',
			accountInfo: rpcRes.accountInfo
		})
	}))
}

function percent(part, whole) {
	if (whole == 0) {
		return '0%'
	}
	return (Math.round((part / whole) * 1000) / 10) + '%'
}

// Util
///////

function checkErr(next, handler) {
	if (arguments.length != 2 || typeof(next) != 'function' || typeof(handler) != 'function') {
		console.log('Bad checkErr arguments')
		process.exit(1)
	}
	return function(rpcErr, rpcRes) {
		if (rpcErr) {
			var err
			try {
				err = JSON.parse(rpcErr.message || rpcErr)
			} catch(e) {
				err = rpcErr
			}
			return next(err)

		} else if (rpcRes.userError) {
			return next({ userError:rpcRes.userError })
			
		} else {
			return handler(rpcRes)
		}
	}
}

function handleAdminAuth(req, res) {
	const b64auth = (req.headers.authorization || '').split(' ')[1] || ''
	const [login, password] = new Buffer(b64auth, 'base64').toString().split(':')
	if (login != 'marcus' || password != 'makeithappen') {
		// res.set('WWW-Authenticate', 'Basic realm="nope"') // change this
		res.set('WWW-Authenticate', 'Basic realm="asd78g6asdg617"')
		res.status(401).send('Not authorized.')
		return
	}
}
