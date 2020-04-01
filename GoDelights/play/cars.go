package play

import {
  "engo.io/ecs"
  "engo.io/engo/common"
}

type Car struct {
  ecs.BasicEntity
  engo.SpaceComponent
  engo.RenderComponent
  VelocityComponent
}

type MovementEntity struct {
  *ecs.BasicEntity
  *common.SpaceComponent
  *VelocityComponent
}

type MovementSystem struct {
  obs []MovementEntity
}

func (ms *MovementSystem) Add(ob Moveable) {
  ms.Add(ob.GetBasicEntity(), ob.GetSpaceComponent(), ob.GetVelocityComponent)
}

func (ms *MovementSystem) AddByInterace(ob) {
}

func (ms *MovementSystem) Remove(e ecs.BasicEntity) {
  ms.obs = RemoveMovemenetEntity(ms.obs, ecs.ID())
}
func (ms *MovementSystem) Update(d float32) {
  for _, c := range ms.obs {
    c.Position.X += c.Vel.X * d
    c.Position.Y += c.Vel.Y * d
  }
}