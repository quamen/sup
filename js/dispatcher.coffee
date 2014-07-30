Promise = require('es6-promise').Promise

module.exports = class Dispatcher
  class PrivateDispatcher
    callbacks: []
    promises: []

    _addPromise: (callback, payload) ->
      @promises.push(new Promise((resolve, reject) ->
        if callback(payload)
          resolve(payload)
        else
          reject(new Error('Dispatcher callback unsuccessful'))
      ))

    _clearPromises: ->
      @promises = []

  # create an internal instance of the private class
  instance = new PrivateDispatcher

  @register: (callback) ->
    instance.callbacks.push(callback)
    instance.callbacks.length - 1

  @dispatch: (payload) ->
    instance._addPromise(callback, payload) for callback in instance.callbacks
    Promise.all(instance.promises).then(instance._clearPromises)

  @handleViewAction: (action) ->
    @dispatch(source: 'VIEW_ACTION', action: action)
