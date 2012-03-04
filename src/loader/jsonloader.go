package loading

import "reflect"

type DataManager interface {
  Index() []string
  Get(identifier string) interface{}
  Store(identifier, interface{}) bool
}

type JSonDataManager struct {
  location string
  kind reflect.Kind

}

