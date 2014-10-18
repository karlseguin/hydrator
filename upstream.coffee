http = require('http')
server = http.createServer (req, res) ->
  res.setHeader('X-Hydrate', '!ref')
  res.setHeader('Content-Type', 'application/json')
  data = {
    more: false,
    results: [
      {'!ref': {id: '233p', type: 'product'}},
      {'!ref': {id: '553p', type: 'product'}},
      {weight: 1.0, product: {'!ref': {id: '553p', type: 'product'}}},
    ]
  }
  res.end(JSON.stringify(data))
server.listen(4005, '127.0.0.1')
console.log('Server running at http://127.0.0.1:4005/')
