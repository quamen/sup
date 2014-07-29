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
    <div>
      <h1>Events</h1>
      <table className="events">
        { @renderHeader() }
        { @renderBody() }
      </table>
    </div>

  EventCount: ->
    _.size(@state.events)

  renderEvent: (event) ->
    <EventComponent event={ event } />

  renderHeader: ->
    <thead>
      <tr>
        <td className="events__NameCol">Event ID</td>
        <td className="events__DescriptionCol">Event Type</td>
      </tr>
    </thead>

  renderBody: ->
    if @EventCount() isnt 0
      <tbody>
        { _(@state.events).map(@renderEvent) }
      </tbody>
    else
      @renderNoEvents()

  renderNoEvents: ->
    <tbody>
      <tr>
        <td colSpan="2" className="Event__Noevents">Sorry, Friend! No events.</td>
      </tr>
    </tbody>

  onEventChange: ->
    @setState(events: @getEventsState())

  getEventsState: ->
    EventStore.all()
