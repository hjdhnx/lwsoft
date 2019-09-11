// +build windows
//运行时请带上参数GOARCH=386

package mylwsoft

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	IID_all = ole.IID_IDispatch
)
type LwSoft struct {
	unknown *ole.IUnknown
	lw      *ole.IDispatch
}
func NewLwSoft() (*LwSoft, error) {
	com := new(LwSoft)
	var err error
	com.unknown, err = oleutil.CreateObject("lw.lwsoft3")
	if err != nil {
		return nil, err
	}
	com.lw, err = com.unknown.QueryInterface(IID_all)
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (com *LwSoft) Ver() string {
	ver, _ := com.lw.CallMethod("ver")
	return ver.ToString()
}