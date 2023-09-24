package main

import (
  "errors"
  "testing"
)
// Тестирование функции CheckFuelCommand (проверяет уровень топлива)
func TestCheckFuelCommand_Execute(t *testing.T) {
  v := &Vehicle{
    Fuel: 0,
  }
  c := &CheckFuelCommand{Vehicle: v}
  err := c.Execute()
  if err == nil || err.Error() != "not enough fuel" {
    t.Errorf("Expected 'not enough fuel' error, got: %v", err)
  }

  v.Fuel = 100
  err = c.Execute()
  if err != nil {
    t.Errorf("Expected no error, got: %v", err)
  }
}



