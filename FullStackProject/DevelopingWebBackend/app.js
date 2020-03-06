const Couchbase = require("couchbase")
const Hapi = require("hapi")
const Joi = require("joi")
const UUID = require("uuid")

const server = new Hapi.Server()
const cluster = new Couchbase.Cluster("couchbase://localhost")
cluster.authenticate("demo", "123456")
const bucket = cluster.openBucket("example")

server.connection({ host: "localhost", port: 3000, routes: { cars: true } })

bucket.on("error", error => {
  throw error
})

server.start(error => {
  if(error) {
    throw error
  }
  console.log("Listening at " + server.info.uri)
  
  server.route({
    method: "POST",
    path: "/person",
    config: {
      validate: {
        payload: {
          firstname: Joi.string().required(),
          lastname: Joi.string().required(),
          type: Joi.string().forbidden().default("person"),
          timestamp: Joi.any().forbidden().default((new Date).getTime())
        }
      }
    },
    handler: (request, response) => {
      var id = UUID.v4()
      bucket.insert(id, request.payload, (error, result) => {
        if(error) {
          return response({ code: error.code, message: error.message }).code(500)
        }
        request.payload.id = id
        response(requestd.payload)
      })
    }
  })
  
  server.route({
    method: "GET",
    path: "/address/{addressid}",
    handler: (request, response) => {
      bucket.get(request.params.addressid, (error, result) => {
        if(error) {
          return response({ code: error.code, message: error.message }).code(500)
        }
        response(result.value)
      })
    }
  })