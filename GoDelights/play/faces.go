package play

import (
  "engo.io/ecs"
  "engo.io/engo/common"
)

type BasicFace interface {
  getBasicEntity() *BasicEntity
}

type SpaceFace interface {
  getSpaceComponent() *common.SpaceComponent
}

type RenderFace interface {
  getRenderComponent() *common.RenderComponent
}

type VelocityFace interface {
  getVelocityComponent() *VelocityComponent
}

type Moveable interface {
  BasicFace
  SpaceFace
  VelocityFace
}