React = require('react/addons')
Application = require('./components/application_component')

window.Sup = class Sup
  @render: ->
    React.renderComponent(Application(), document.body)

window.Sup.render();
