var Couchbase = require("couchbase");
var Express = require("express");
var BodyParser = require("body-parser");
var Request = require("request");
 
var app = Express();
var N1qlQuery = Couchbase.N1qlQuery;
 
app.use(BodyParser.json());
app.use(BodyParser.urlencoded({ extended: true }));
 
var bucket = (new Couchbase.Cluster("couchbase://localhost")).openBucket("example");
 
var makePostRequest = function(url, body, callback) {
    Request.post(url, { json: body },
        function (error, response, body) {
            if (!error &amp;&amp; response.statusCode == 200) {
                callback(null, body);
            } else {
                callback(error, null);
            }
        }
    );
}
 
app.get("/list", function(req, res) {
    bucket.query(N1qlQuery.fromString("SELECT type, firstname, lastname FROM `" + bucket._name + "` WHERE type = 'person'"), function(error, result) {
        if(error) {
            return res.status(400).send(error);
        }
        res.send(result);
    });
});
 
app.post("/create", function(req, res) {
    if(!req.body.firstname) {
        return res.status(400).send({"status": "error", "message": "A firstname is required"});
    } else if(!req.body.lastname) {
        return res.status(400).send({"status": "error", "message": "A lastname is required"});
    }
    makePostRequest("http://localhost:4984/" + bucket._name + "/", {type: "person", firstname: req.body.firstname, lastname: req.body.lastname}, function(error, result) {
        if(error) {
            return res.status(400).send(error);
        }
        res.send(req.body);
    });
});
 
var server = app.listen(3000, function() {
    console.log("Listening on port %s...", server.address().port);
});