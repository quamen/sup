Dispatcher = require('../dispatcher')
constants = require('../stores/event_store').constants

module.exports = class EventActions
  @fetch: ->
    Dispatcher.handleViewAction(actionType: constants.FETCH_EVENTS)
