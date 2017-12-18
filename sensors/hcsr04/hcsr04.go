package HCSR04

import (
  "time"
  "github.com/kranfix/go-rpio"
)

type Hcsr04 struct {
  Trigger rpio.Pin
  Echo    rpio.Pin
  Scale   float64
}

const SoundSpeed = 343

func New(trigger, echo rpio.Pin) *Hcsr04 {
  trigger.Output()
  trigger.Low()
  echo.Input()
  return &Hcsr04{trigger,echo, 1.0}
}

func (hc *Hcsr04) SetScale(scale float64) {
  if(scale > 0){
    hc.Scale = scale
  }
}

func (hc *Hcsr04) ReadTime() time.Duration {
  hc.Trigger.High()
  time.Sleep(10 * time.Microsecond)
  hc.Trigger.Low()

  for hc.Echo.Read() != rpio.High {

  }

  start := time.Now()
  for hc.Echo.Read() == rpio.High {
    now := time.Now()
  }

  elapse := now.Sub(start)
}

// Return the distance in meters
func (hc * Hcsr04) ReadDistance() float64 {
  sound = float64(SoundSpeed) * 1e-9 * hc.Sclae
  return float64(hc.ReadTime()) * sound
}
