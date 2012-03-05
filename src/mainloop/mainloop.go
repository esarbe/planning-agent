package mainloop

import "bufio"
import "os"

type Runner interface {
  Prepare ()
  Run () bool
  Finalize()
}

func Loop (r Runner) {
  run := true
  for run {
    run = r.Run()
  }
}

type Mainloop struct {
  PrepareFunction func()
  RunFunction func() bool
  FinalizeFunction func()
}

type ConsoleMainloop struct {
  *Mainloop
  *bufio.ReadWriter
}

func NewConsoleMainloop () ConsoleMainloop {

  indevice := bufio.NewReader(os.Stdin)
  outdevice := bufio.NewWriter(os.Stdin)
  inout := bufio.NewReadWriter(indevice, outdevice)

  ml := Mainloop{
    func () {},
    func () bool { return false },
    func () {},
  }


  cm := ConsoleMainloop{&ml, inout}

  return cm
}

func (cm *Mainloop) Prepare () {

}

func (cm *Mainloop) Run () bool {
  return cm.RunFunction()
}

func (cm *Mainloop) Finalize () {

}

