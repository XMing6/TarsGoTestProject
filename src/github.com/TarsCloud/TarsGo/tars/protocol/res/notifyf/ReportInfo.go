//Package notifyf comment
// This file war generated by tars2go 1.1
// Generated from NotifyF.tars
package notifyf

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//ReportInfo strcut implement
type ReportInfo struct {
	EType      ReportType  `json:"eType"`
	SApp       string      `json:"sApp"`
	SSet       string      `json:"sSet"`
	SContainer string      `json:"sContainer"`
	SServer    string      `json:"sServer"`
	SMessage   string      `json:"sMessage"`
	SThreadId  string      `json:"sThreadId"`
	ELevel     NOTIFYLEVEL `json:"eLevel"`
}

func (st *ReportInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *ReportInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int32((*int32)(&st.EType), 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SApp, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SSet, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SContainer, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SServer, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SMessage, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SThreadId, 7, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32((*int32)(&st.ELevel), 8, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ReportInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ReportInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ReportInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(int32(st.EType), 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SApp, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SSet, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SContainer, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SServer, 5)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SMessage, 6)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SThreadId, 7)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(st.ELevel), 8)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *ReportInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
