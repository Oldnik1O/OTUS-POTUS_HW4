// Тестирование функции BurnFuelCommand (уменьшает уровень топлива на значение, равное потреблению топлива)
package main

import (
  "errors"
  "testing"
)
func TestBurnFuelCommand_Execute(t *testing.T) {
  v := &Vehicle{
    Fuel:           100,
    FuelConsumption: 1,
  }
  b := &BurnFuelCommand{Vehicle: v}
  err := b.Execute()
  if err != nil {
    t.Errorf("Expected no error, got: %v", err)
  }
  if v.Fuel != 99 {
    t.Errorf("Expected Fuel to be 99, got: %f", v.Fuel)
  }
}