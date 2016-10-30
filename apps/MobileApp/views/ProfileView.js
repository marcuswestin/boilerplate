var AutoReact = require('autoreact')
var Tags = require('tags/react-native')

var state = AutoReact.createState({
	name: String
})
state.name = 'Marcus'

var ProfileView = AutoReact.createClass({
	statics: {
		title: 'Profile'
	},
	render: function() {
		return View(
			Text('Profile'),
			TouchableHighlight(
				OnPress(() => { state.name += ' W' }),
				Text(state.name)
			),
			TouchableHighlight(Text("Next"), OnPress(() => {
				Navigator.push(ProfileView)
			}))

		)
	}
})

module.exports = ProfileView
