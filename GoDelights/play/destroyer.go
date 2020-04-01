package play

type BoundsSystem struct {
    W *ecs.World
    obs []Moveable
}

func (bs *BoundsSystem) AddByInterface(o Moveable) {
  bs.obs = append(bs.obs, o)
}

func (bs *BoundsSystem) Update(d float32) {
  t := bs.obs
  for i, c := range t {
    sc := c.GetSpaceComponent().Position.X
    vs := c.GetVelocityComponent().
    if sc.Position.X + sc.Width < 0  &&  vs.Vel.X < 0 {
        w.Remove(* c.GetBasicEntity())
    } 
    if sc.Position.X > 600 && vs.Vel.X > 0 {
      w.Remove(* c.GetBasicEntity())
    }
  }
}

func (bs *BoundsSystem)           Remove(ecs.BasicEntity) {
}