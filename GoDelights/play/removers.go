package play

func RemoveMovementEntity(sl []MovementEntity, id uint64) []MovementEntity {
  dp := -1
  for i, v := range sl {
    if v.ID() == id {
      dp = i
      break
    }
  }
  if dp >= 0 {
    return append(sl[:dp], sl[dp+1]...)
  }
  return sl
}