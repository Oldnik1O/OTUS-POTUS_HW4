// Разработаны  команды  для управления объектами в игре:
// CheckFuelCommand проверяет уровень топлива.
// BurnFuelCommand уменьшает уровень топлива на значение, равное потреблению топлива.
//  MoveCommand представляет собой заглушку для движения объекта.
// MacroCommand позволяет объединить несколько команд в одну и выполнить их по порядку.
// ModifyVelocityCommand изменяет направление объекта в зависимости от его текущей скорости и угла поворота.
// RotateCommand - заглушка для поворота объекта.

// В main две макрокоманды:
// 1. Для движения объекта и расхода топлива.
// 2. Для поворота объекта и изменения его направления.

go
package main

import (
  "errors"
  "fmt"
  "math"
)

type Command interface {
  Execute() error
}

type Vehicle struct {
  Fuel           float64
  FuelConsumption float64
  Velocity       float64
  Direction      float64
}

type CheckFuelCommand struct {
  Vehicle *Vehicle
}

func (c *CheckFuelCommand) Execute() error {
  if c.Vehicle.Fuel <= 0 {
    return errors.New("not enough fuel")
  }
  return nil
}

type BurnFuelCommand struct {
  Vehicle *Vehicle
}

func (b *BurnFuelCommand) Execute() error {
  c.Vehicle.Fuel -= c.Vehicle.FuelConsumption
  return nil
}

type MoveCommand struct {
  Vehicle *Vehicle
}

func (m *MoveCommand) Execute() error {
  // Logic for moving the vehicle goes here...
  return nil
}

type MacroCommand struct {
  Commands []Command
}

func (m *MacroCommand) Execute() error {
  for _, command := range m.Commands {
    if err := command.Execute(); err != nil {
      return err
    }
  }
  return nil
}

type ModifyVelocityCommand struct {
  Vehicle   *Vehicle
  Direction float64
}

func (m *ModifyVelocityCommand) Execute() error {
  if c.Vehicle.Velocity == 0 {
    return nil
  }
  c.Vehicle.Direction = math.Mod(c.Vehicle.Direction+m.Direction, 360)
  return nil
}

type RotateCommand struct {
  Vehicle *Vehicle
}

func (r *RotateCommand) Execute() error {
  // Logic for rotating the vehicle goes here...
  return nil
}

func main() {
  v := &Vehicle{
    Fuel:           100,
    FuelConsumption: 1,
    Velocity:       0,
    Direction:      0,
  }

  moveAndConsumeFuel := &MacroCommand{
    Commands: []Command{
      &CheckFuelCommand{Vehicle: v},
      &MoveCommand{Vehicle: v},
      &BurnFuelCommand{Vehicle: v},
    },
  }

  rotateAndChangeDirection := &MacroCommand{
    Commands: []Command{
      &RotateCommand{Vehicle: v},
      &ModifyVelocityCommand{Vehicle: v, Direction: 45},
    },
  }

  if err := moveAndConsumeFuel.Execute(); err != nil {
    fmt.Println("Error:", err)
  }

  if err := rotateAndChangeDirection.Execute(); err != nil {
    fmt.Println("Error:", err)
  }
}

