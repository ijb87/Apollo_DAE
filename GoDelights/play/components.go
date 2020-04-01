package play

type VelocityComponent struct { 
    Vel engo.Point
  }
  
  func (*VelocityComponent) GetVelocityComponent() *VelocityComponent {
    return v
  }
  
  