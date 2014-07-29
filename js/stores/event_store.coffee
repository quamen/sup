_ = require('lodash')
keyMirror = require('react/lib/keyMirror')
EventEmitter = require('events').EventEmitter
Promise = require('es6-promise').Promise
Request = require('superagent')
Dispatcher = require('../dispatcher')

constants = keyMirror
  FETCH_EVENTS: null

class EventStore extends EventEmitter
  class InternalStore
    CHANGE_EVENT: 'change'
    events: {}

    reset: ->
      @events = {}

    replace: (events) ->
      events ||= {}
      @reset()
      @create(events) for ref in events

    create: (data) ->
      @events[data.id] = data

    get: (id) ->
      @events[id]

    update: (id, updates) ->
      @events[id] = _.extend(@events[id], updates)

    fetch: ->
      new Promise((resolve, reject) =>
        request = Request.get('/events/')
        request.set('X-Requested-With', 'XMLHttpRequest')
        request.set('Accept', 'application/json')
        request.end (response) =>
          return reject(response) unless response.ok
          @replace(response.body)
          resolve(response)
      )

  # create an internal instance of the private class
  store = new InternalStore

  @all: ->
    store.events

  @find: (name) ->
    store.get(name)

  @emitChange: =>
    @::emit(store.CHANGE_EVENT)

  @addChangeListener: (callback) =>
    @::addListener(store.CHANGE_EVENT, callback)

  @removeChangeListener: (callback) =>
    @::removeListener(store.CHANGE_EVENT, callback)

  # Register callbacks with the dispatcher
  Dispatcher.register (payload) ->
    switch payload.action.actionType
      when constants.FETCH_EVENTS
        store.fetch().then(eventstore.emitChange)
    true

  source = new EventSource('/events')
  source.onmessage = (e) =>
    console.log e
    event = JSON.parse(e.data)
    store.create(event)
    @emitChange()


module.exports = {constants, EventStore}
