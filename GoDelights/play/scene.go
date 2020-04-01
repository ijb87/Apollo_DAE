package play

import (
  "engo.io/engo"
  "image/color"
  "engo.io/engo/common"
)

type PlayScene struct {
  
}

func (*PlayScene) Type() string{return "PlayScene"}

func (*PlayScene) Preload() {}

func (*PlayScene) Setup(w *ecs.World) {
  common.SetBackground(color.White)
  
  systems := &SysList{
  RenderSys = &common.RenderSystem{},
  MovementSys: &MovementSystem{},
  }

  spawnSys := SpawnSystem{Systems: systems}  
  fg := NewFrog()
  rs.Add(&fg.BasicEntity, &fg.RenderComponent, &fg.SpaceComponent)
  systems.RenderSys.AddByInterface(fg)
  
  w.AddSystem(systems.RenderSys)
  w.AddSystem(systems.MovementSys)
  w.AddSystem(spawnSys)
  
  w.AddSystem(rs)
}