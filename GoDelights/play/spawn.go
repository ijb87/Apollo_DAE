package play

import "engo.io/ecs"

type SysList struct {
  RenderSys common.RenderSystem
  MovementSys *MovementSystem
}

type SpawnSystem struct {
  since float32
  Systems *SysList
  nCars int
}

func (ss *SpawnSystem) Remove(c ecs.BasicEntity) {
}

func (ss *SpawnSystem) Update(d float32) {
  ss.since = since + d
  if ss.since > 0.2 {
    row := rand.Intn(7)
    vel := engo.Point{float32((row%2) * 60 - 30,0}
    pos := engo.Point{float32(row%2)*-800 + 700, float32(row*50)}
    
    ss.since = 0
    ss.nCars++
    
    c := &Car {
      BasicEntity: ecs.NewBasic(),
      SpaceComponent: common.SpaceComponent{Position:pos, Width: 100, Height: 50},
      RenderComponent: common.RenderComponent {
        Drawable: common.Rectangle,
        Color: color.Black,
      },
      VelocityComponent : VelocityComponent{vel}
    }
    
    ss.Systems.RenderSys.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
    ss.Systems.MovementSys.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
  }
}