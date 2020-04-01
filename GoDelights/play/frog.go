package play

import (
  "engo.io/engo/common"
  "engo.io/ecs"
  "image/color"
)

type Frog struct {
  ecs.BasicEntity
  common.RenderComponent
  common.SpaceComponent
}

func NewFrog() *Frog {
  return &Frog {
    BasicEntity: ecs.NewBasic(),
    RenderComponent: common.RenderComponent {
      Drawable: common.Triangle{},
      Color:color.Black,
    },
    SpaceComponent: common.SpaceComponent {
      Position:engo.Point {330, 460},
      Width: 40,
      Height: 40,
    },
  }
}