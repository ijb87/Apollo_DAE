import Physics = require("nativescript-physics-js")
import { Page } from "ui"

const gravity_scale = 0.003
var page: Page
var init = false
exportfunction pageLoaded(args) {
  // ... initialization
  // Add the ball
  addBall(world, 50, 150)
}

function addBall(world, x: number, y: number) {
  var ball = Physics.body('circle', {
    label: "ball",
    x: x,
    y: y,
    radius: 15,
    label: "ball",
    styles: { image:"~/images/ball.png" }
  })
  ball.restitution = 0.3
  
  world.add(ball)
}

function addWall(world, x: number, y: number, width: number, height: number, angle: number = 0) {
  world.addPhysics.body('rectangle', {
    treatment: 'static',
    x: x,
    y: y,
    width: width,
    height: height,
    angle: angle,
    styles: { color: "orange" }
  }))
}

function addTarget(world, x: number, y: number) {
  world.add(Physics.body('circle', {
    label: 'target',
    treatment: 'static',
    x: x,
    y: y,
    radius: 20,
    styles: { image:"~/images/target.png" }
  }))
}
  
  // Get references to container and meta-info views
  var page = args.object
  var container = page.getViewById("container")
  var metaText = page.getViewById("meta")
  
  // Create physics world
  var world = Physics({sleepDisabled:true})
  
  var query = Physics.query({
    $or: [
      { bodyA: { label: 'ball' }, bodyB: { label:'target' } },
      { bodyB: { label: 'target' }, bodyA: { label: 'ball' } }
    ]
  })
  
  world.on('collisions:detected', function(data, e) {
    if (Physics.util.find(data.collisions, query)) {
      world.pause()
      alert("You Winn!!!")
    }
  })
  
  // After physics world initialization
  addWall(world, 0, 150, 20, 300)
  addWall(world, 300, 150, 20, 300)
  addWall(world, 150, 0, 300, 20)
  addWall(world, 150, 300, 300, 20)
  addWall(world, 150, 250, 10, 200)
  
  addTarget(world, 225, 225)
  addBall(world, 50, 250)
  
  // Add {N} renderer
  world.add(Physics.renderer('ns', {
    container: container,
    metaText: metaText,
    meta: true
  }))
  
  // Add behaviors
  world.add([
    Physics.behavior('edge-collision-detection',
    { aabb: Physics.aabb(0, 0, 300, 300) }), // Scene bounds
    Physics.behavior('body-collision-detection'), // Collision related
    Physics.behavior('body-inpulse-response'), // Collision related
    Physics.behavior('sweep-prune'), // Collision related
    Physics.behavior('constant-acceleration') // Gravity
  ])
  
  var gravity = Physics.behavior('constant-acceleration', 
  { acc: { x: 0, y: 0 } })
  world.add([
    // other behaviors
    gravity
  ])
  
  // Start accelerometer events 1 second after the word is created.
  setTimeout(function() {
    accService.startAccelerometerUpdates((data) => {
      var xAcc = -data.x * 0.003
      var yAcc = data.y * 0.003
      gravity.setAcceleration({ x: xAcc, y: yAcc })
    })
  }, 1000)
  
  // Start ticking - render new scene every 20ms
  world.on('step', function () { world.render() })
  setInterval(function () { world.step(Date.now()) }, 20)
}