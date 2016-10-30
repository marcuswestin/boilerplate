var AutoReact = require('autoreact')
var React = require('react')
var Tags = require('tags/react-native')
var ProfileView = require('./ProfileView')

var { NavigatorIOS } = Tags

var Navigator = module.exports = React.createFactory(React.createClass({
	render: function() {
		return NavigatorIOS(Flex(1), {
			ref: 'nav',
			initialRoute: _makeRoute(HomeView)
		})
	}
}))

Navigator.push = function(Component, title) {
	nav.push(_makeRoute(Component, title))
}

// Internal
///////////

var HomeView = AutoReact.createClass({
	statics: {
		title: 'StackAck'
	},
	
	_onBack: function () {
		this.props.navigator.pop()
	},

	render: function () {
		return View(
			Text('Current Scene: ', this.props.title),
			TouchableHighlight(Text("Profile"), OnPress(() => {
				Navigator.push(ProfileView)
			}))
		)
	},
})

var Scene = React.createFactory(React.createClass({
	componentDidMount: function() {
		nav = this.props.navigator
	},
	render: function() {
		return View(Padding(80, 20),
			this.props.Component()
		)
	}
}))

function _makeRoute(Component, title) {
	return {
		component: Scene,
		title: (title || Component.title),
		passProps: { Component:Component },
		// tintColor: "#008888",
		// rightButtonTitle:'Add',
		// onRightButtonPress: () => this._handleNavigationRequest()
	}
}
