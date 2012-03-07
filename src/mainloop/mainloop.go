package mainloop

import "bufio"
import "os"

type Runner interface {
  Prepare ()
  Setup ()
  Run () bool
  Finalize()
}

func Loop (r Runner) {
  run := true
  for run {
    r.Setup()
    run = r.Run()
  }
}

type Mainloop struct {
  PrepareFunction func()
  RunFunction func() bool
  SetupFunction func()
  FinalizeFunction func()
}

type ConsoleMainloop struct {
  *Mainloop
  *bufio.ReadWriter
  prompt string
}

func NewConsoleMainloop (prompt string) ConsoleMainloop {

  indevice := bufio.NewReader(os.Stdin)
  outdevice := bufio.NewWriter(os.Stdin)
  inout := bufio.NewReadWriter(indevice, outdevice)

  ml := Mainloop{
    func () {},
    func () bool { return false },
    func () { inout.WriteString(prompt)},
    func () {},
  }


  cm := ConsoleMainloop{&ml, inout, prompt}

  return cm
}

func (cm *Mainloop) Prepare () {

}

func (cm *Mainloop) Run () bool {
  return cm.RunFunction()
}

func (cm *Mainloop) Finalize () {

}

func (ml *Mainloop) Setup () {
  ml.SetupFunction()
}


