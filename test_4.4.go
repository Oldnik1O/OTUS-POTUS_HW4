// Тестирование функции ModifyVelocityCommand (изменение направления объекта в зависимости от его текущей скорости и угла поворота)
package main

import (
  "errors"
  "testing"
)
func TestModifyVelocityCommand_Execute(t *testing.T) {
  v := &Vehicle{
    Velocity:  0,
    Direction: 0,
  }
  m := &ModifyVelocityCommand{Vehicle: v, Direction: 45}
  err := m.Execute()
  if err != nil {
    t.Errorf("Expected no error, got: %v", err)
  }
  if v.Direction != 0 {
    t.Errorf("Expected Direction to be 0, got: %f", v.Direction)
  }
}