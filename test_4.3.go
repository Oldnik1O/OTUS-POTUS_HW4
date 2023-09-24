// Тестирование функции MacroCommand (позволяет объединить несколько команд в одну и выполнить их по порядку)
package main

import (
  "errors"
  "testing"
)
func TestMacroCommand_Execute(t *testing.T) {
  v := &Vehicle{
    Fuel:           100,
    FuelConsumption: 1,
  }
  macro := &MacroCommand{
    Commands: []Command{
      &CheckFuelCommand{Vehicle: v},
      &BurnFuelCommand{Vehicle: v},
    },
  }
  err := macro.Execute()
  if err != nil {
    t.Errorf("Expected no error, got: %v", err)
  }
  if v.Fuel != 99 {
    t.Errorf("Expected Fuel to be 99, got: %f", v.Fuel)
  }

  macro.Commands = append(macro.Commands, &CheckFuelCommand{Vehicle: v})
  v.Fuel = 0
  err = macro.Execute()
  if err == nil || !errors.Is(err, errors.New("not enough fuel")) {
    t.Errorf("Expected 'not enough fuel' error, got: %v", err)
  }
}
