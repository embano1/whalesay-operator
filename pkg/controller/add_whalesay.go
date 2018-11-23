package controller

import (
	"github.com/embano1/whalesay-operator/pkg/controller/whalesay"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, whalesay.Add)
}
