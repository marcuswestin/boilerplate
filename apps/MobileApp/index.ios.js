require('./globals')

// App UI
/////////

alert(require('foo'))

var MobileApp = AutoReact.createClass({
  _handleNavigationRequest: function () {
    this.refs.nav.push({
      component: MyScene,
      title: 'Genius',
      // passProps: { myProp: 'genius' },
    });
  },

  render: function() {
    return Navigator()
  }
})

Tags.render(MobileApp, 'MobileApp');
