React = require('react/addons')
RefActions = require('../actions/event_actions')

module.exports = React.createClass
  displayName: 'Event'
  propTypes:
    event: React.PropTypes.object.isRequired

  render: ->
    image_tag = '<img src="' + @props.event.actor.avatar_url + ' class="img-circle img-responsive" alt="" height="80", width="80">'
    <li className="events__event list-group-item">
      <div className="row">
        <div className="col-xs-2 col-md-1" dangerouslySetInnerHTML={ __html: image_tag }>
        </div>
        <div className="col-xs-10 col-md-11">
          <h3>
            { @title() }
          </h3>
        </div>
      </div>
    </li>

  title: ->
    @props.event.payload?.pull_request?.title ||
    @props.event.payload?.issue?.title ||
    @props.event.type
