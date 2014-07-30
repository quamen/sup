_ = require('lodash')
React = require('react/addons')
EventStore = require('../stores/event_store').EventStore
EventActions = require('../actions/event_actions')
EventComponent = require('./event_component')

module.exports = React.createClass
  displayName: 'Events'

  componentDidMount: ->
    EventStore.addChangeListener(@onEventChange)
    EventActions.fetch()

  componentWillUnmount: ->
    EventStore.removeChangeListener(@onEventChange)

  getInitialState: ->
    events: @getEventsState()

  render: ->
    <div className="container">
      <div className="row">
        <div className="panel panel-default widget">
          <div className="panel-heading">
            <h3 className="panel-title">Recent Events</h3>
            <span className="label label-info" dangerouslySetInnerHTML={ __html: @EventCount()}/>
          </div>
          <div className="panel-body">
            <ul className="list-group">
              { _(_(@state.events).sortBy((event) -> event.ID)).reverse().map(@renderEvent) }
            </ul>
          </div>
        </div>
      </div>
    </div>

  EventCount: ->
    _.size(@state.events)

  renderEvent: (event) ->
    <EventComponent event={ event } />

  onEventChange: ->
    @setState(events: @getEventsState())

  getEventsState: ->
    EventStore.all()
