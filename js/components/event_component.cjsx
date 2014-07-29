React = require('react/addons')
RefActions = require('../actions/event_actions')

module.exports = React.createClass
  displayName: 'Event'
  propTypes:
    event: React.PropTypes.object.isRequired

  render: ->
    <tr className="events__event">
      <td dangerouslySetInnerHTML={ __html: @props.event.id }/>
      <td dangerouslySetInnerHTML={ __html: @props.event.type }/>
    </tr>
