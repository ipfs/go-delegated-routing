package proto

import(
	pd1 "github.com/ipld/edelweiss/values"
	pd2 "github.com/ipld/go-ipld-prime/datamodel"
	pd3 "fmt"
	pd4 "io"
	pd5 "github.com/ipld/edelweiss/services"
	pd6 "bytes"
	pd7 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd8 "errors"
	pd9 "github.com/ipld/go-ipld-prime"
	pd10 "net/url"
	pd11 "context"
	pd12 "net/http"
	pd13 "github.com/ipfs/go-log"
	pd14 "github.com/ipfs/go-cid"
	pd15 "github.com/ipld/go-ipld-prime/linking/cid"
)


// -- protocol type DelegatedRouting_IdentifyArg --

type DelegatedRouting_IdentifyArg struct {

}

func (x DelegatedRouting_IdentifyArg) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_IdentifyArg) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_IdentifyArg_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyArg
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Done() bool {
	return x.i + 1 >= 0
}

func (x DelegatedRouting_IdentifyArg) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_IdentifyArg) LookupByString(key string) (pd2.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_IdentifyArg_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyArg) ListIterator() pd2.ListIterator {
	return nil
}

func (x DelegatedRouting_IdentifyArg) Length() int64 {
	return 0
}

func (x DelegatedRouting_IdentifyArg) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_IdentifyArg) IsNull() bool {
	return false
}

func (x DelegatedRouting_IdentifyArg) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonList1 --

type AnonList1 []pd1.String

func (v AnonList1) Node() pd2.Node {
	return v
}

func (v *AnonList1) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList1, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList1) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList1) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList1) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList1) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList1) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList1) ListIterator() pd2.ListIterator {
	return &AnonList1_ListIterator{v, 0}
}

func (v AnonList1) Length() int64 {
	return int64(len(v))
}

func (AnonList1) IsAbsent() bool {
	return false
}

func (AnonList1) IsNull() bool {
	return false
}

func (v AnonList1) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonList1) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList1) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList1_ListIterator struct {
	list AnonList1
	at   int64
}

func (iter *AnonList1_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList1_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type DelegatedRouting_IdentifyResult --

type DelegatedRouting_IdentifyResult struct {
		Methods AnonList1

}

func (x DelegatedRouting_IdentifyResult) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_IdentifyResult) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Methods": x.Methods.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Methods":
			if _, notParsed := fieldMap["Methods"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Methods")
			}
			if err := x.Methods.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Methods")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_IdentifyResult_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyResult
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Methods"), x.s.Methods.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x DelegatedRouting_IdentifyResult) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_IdentifyResult) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_IdentifyResult_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyResult) ListIterator() pd2.ListIterator {
	return nil
}

func (x DelegatedRouting_IdentifyResult) Length() int64 {
	return 1
}

func (x DelegatedRouting_IdentifyResult) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_IdentifyResult) IsNull() bool {
	return false
}

func (x DelegatedRouting_IdentifyResult) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type DelegatedRouting_Error --

type DelegatedRouting_Error struct {
		Code pd1.String

}

func (x DelegatedRouting_Error) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_Error) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Code": x.Code.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Code":
			if _, notParsed := fieldMap["Code"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Code")
			}
			if err := x.Code.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Code")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_Error_MapIterator struct {
	i int64
	s *DelegatedRouting_Error
}

func (x *DelegatedRouting_Error_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Code"), x.s.Code.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_Error_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x DelegatedRouting_Error) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_Error) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_Error_MapIterator{-1, &x}
}

func (x DelegatedRouting_Error) ListIterator() pd2.ListIterator {
	return nil
}

func (x DelegatedRouting_Error) Length() int64 {
	return 1
}

func (x DelegatedRouting_Error) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_Error) IsNull() bool {
	return false
}

func (x DelegatedRouting_Error) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_Error) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonInductive4 --

type AnonInductive4 struct {
		Identify *DelegatedRouting_IdentifyArg
		FindProviders *FindProvidersRequest
		GetIPNS *GetIPNSRequest
		PutIPNS *PutIPNSRequest


}

func (x *AnonInductive4) Parse(n pd2.Node) error {
	*x = AnonInductive4{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "IdentifyRequest":
		var y DelegatedRouting_IdentifyArg
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "FindProvidersRequest":
		var y FindProvidersRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.FindProviders = &y
		return nil
	case "GetIPNSRequest":
		var y GetIPNSRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetIPNS = &y
		return nil
	case "PutIPNSRequest":
		var y PutIPNSRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.PutIPNS = &y
		return nil


	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type AnonInductive4_MapIterator struct {
	done bool
	s    *AnonInductive4
}

func (x *AnonInductive4_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Identify != nil:
			return pd1.String("IdentifyRequest"), x.s.Identify.Node(), nil
			case x.s.FindProviders != nil:
			return pd1.String("FindProvidersRequest"), x.s.FindProviders.Node(), nil
			case x.s.GetIPNS != nil:
			return pd1.String("GetIPNSRequest"), x.s.GetIPNS.Node(), nil
			case x.s.PutIPNS != nil:
			return pd1.String("PutIPNSRequest"), x.s.PutIPNS.Node(), nil


		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive4_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive4) Node() pd2.Node {
	return x
}

func (x AnonInductive4) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x AnonInductive4) LookupByString(key string) (pd2.Node, error) {
	switch {
		case x.Identify != nil && key == "IdentifyRequest":
		return x.Identify.Node(), nil
		case x.FindProviders != nil && key == "FindProvidersRequest":
		return x.FindProviders.Node(), nil
		case x.GetIPNS != nil && key == "GetIPNSRequest":
		return x.GetIPNS.Node(), nil
		case x.PutIPNS != nil && key == "PutIPNSRequest":
		return x.PutIPNS.Node(), nil


	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive4) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "IdentifyRequest":
		return x.Identify.Node(), nil
		case "FindProvidersRequest":
		return x.FindProviders.Node(), nil
		case "GetIPNSRequest":
		return x.GetIPNS.Node(), nil
		case "PutIPNSRequest":
		return x.PutIPNS.Node(), nil


	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) MapIterator() pd2.MapIterator {
	return &AnonInductive4_MapIterator{false, &x}
}

func (x AnonInductive4) ListIterator() pd2.ListIterator {
	return nil
}

func (x AnonInductive4) Length() int64 {
	return 1
}

func (x AnonInductive4) IsAbsent() bool {
	return false
}

func (x AnonInductive4) IsNull() bool {
	return false
}

func (x AnonInductive4) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x AnonInductive4) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive4) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) Prototype() pd2.NodePrototype {
	return nil
}
// -- protocol type AnonInductive5 --

type AnonInductive5 struct {
		Identify *DelegatedRouting_IdentifyResult
		FindProviders *FindProvidersResponse
		GetIPNS *GetIPNSResponse
		PutIPNS *PutIPNSResponse
		Error *DelegatedRouting_Error


}

func (x *AnonInductive5) Parse(n pd2.Node) error {
	*x = AnonInductive5{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "IdentifyResponse":
		var y DelegatedRouting_IdentifyResult
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "FindProvidersResponse":
		var y FindProvidersResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.FindProviders = &y
		return nil
	case "GetIPNSResponse":
		var y GetIPNSResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetIPNS = &y
		return nil
	case "PutIPNSResponse":
		var y PutIPNSResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.PutIPNS = &y
		return nil
	case "Error":
		var y DelegatedRouting_Error
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Error = &y
		return nil


	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type AnonInductive5_MapIterator struct {
	done bool
	s    *AnonInductive5
}

func (x *AnonInductive5_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Identify != nil:
			return pd1.String("IdentifyResponse"), x.s.Identify.Node(), nil
			case x.s.FindProviders != nil:
			return pd1.String("FindProvidersResponse"), x.s.FindProviders.Node(), nil
			case x.s.GetIPNS != nil:
			return pd1.String("GetIPNSResponse"), x.s.GetIPNS.Node(), nil
			case x.s.PutIPNS != nil:
			return pd1.String("PutIPNSResponse"), x.s.PutIPNS.Node(), nil
			case x.s.Error != nil:
			return pd1.String("Error"), x.s.Error.Node(), nil


		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive5_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive5) Node() pd2.Node {
	return x
}

func (x AnonInductive5) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x AnonInductive5) LookupByString(key string) (pd2.Node, error) {
	switch {
		case x.Identify != nil && key == "IdentifyResponse":
		return x.Identify.Node(), nil
		case x.FindProviders != nil && key == "FindProvidersResponse":
		return x.FindProviders.Node(), nil
		case x.GetIPNS != nil && key == "GetIPNSResponse":
		return x.GetIPNS.Node(), nil
		case x.PutIPNS != nil && key == "PutIPNSResponse":
		return x.PutIPNS.Node(), nil
		case x.Error != nil && key == "Error":
		return x.Error.Node(), nil


	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive5) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "IdentifyResponse":
		return x.Identify.Node(), nil
		case "FindProvidersResponse":
		return x.FindProviders.Node(), nil
		case "GetIPNSResponse":
		return x.GetIPNS.Node(), nil
		case "PutIPNSResponse":
		return x.PutIPNS.Node(), nil
		case "Error":
		return x.Error.Node(), nil


	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) MapIterator() pd2.MapIterator {
	return &AnonInductive5_MapIterator{false, &x}
}

func (x AnonInductive5) ListIterator() pd2.ListIterator {
	return nil
}

func (x AnonInductive5) Length() int64 {
	return 1
}

func (x AnonInductive5) IsAbsent() bool {
	return false
}

func (x AnonInductive5) IsNull() bool {
	return false
}

func (x AnonInductive5) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x AnonInductive5) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive5) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) Prototype() pd2.NodePrototype {
	return nil
}
var logger_client_DelegatedRouting = pd13.Logger("service/client/DelegatedRouting")

type DelegatedRouting_Client interface {

Identify(ctx pd11.Context, req *DelegatedRouting_IdentifyArg) ([]*DelegatedRouting_IdentifyResult, error)

FindProviders(ctx pd11.Context, req *FindProvidersRequest) ([]*FindProvidersResponse, error)

GetIPNS(ctx pd11.Context, req *GetIPNSRequest) ([]*GetIPNSResponse, error)

PutIPNS(ctx pd11.Context, req *PutIPNSRequest) ([]*PutIPNSResponse, error)


Identify_Async(ctx pd11.Context, req *DelegatedRouting_IdentifyArg) (<-chan DelegatedRouting_Identify_AsyncResult, error)

FindProviders_Async(ctx pd11.Context, req *FindProvidersRequest) (<-chan DelegatedRouting_FindProviders_AsyncResult, error)

GetIPNS_Async(ctx pd11.Context, req *GetIPNSRequest) (<-chan DelegatedRouting_GetIPNS_AsyncResult, error)

PutIPNS_Async(ctx pd11.Context, req *PutIPNSRequest) (<-chan DelegatedRouting_PutIPNS_AsyncResult, error)

}


type DelegatedRouting_Identify_AsyncResult struct {
	Resp *DelegatedRouting_IdentifyResult
	Err  error
}

type DelegatedRouting_FindProviders_AsyncResult struct {
	Resp *FindProvidersResponse
	Err  error
}

type DelegatedRouting_GetIPNS_AsyncResult struct {
	Resp *GetIPNSResponse
	Err  error
}

type DelegatedRouting_PutIPNS_AsyncResult struct {
	Resp *PutIPNSResponse
	Err  error
}


type DelegatedRouting_ClientOption func(*client_DelegatedRouting) error

type client_DelegatedRouting struct {
	httpClient       *pd12.Client
	endpoint     *pd10.URL
}

func DelegatedRouting_Client_WithHTTPClient(hc *pd12.Client) DelegatedRouting_ClientOption {
	return func(c *client_DelegatedRouting) error {
		c.httpClient = hc
		return nil
	}
}

func New_DelegatedRouting_Client(endpoint string, opts ...DelegatedRouting_ClientOption) (*client_DelegatedRouting, error) {
	u, err := pd10.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_DelegatedRouting{endpoint: u, httpClient: pd12.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}



func (c *client_DelegatedRouting) Identify(ctx pd11.Context, req *DelegatedRouting_IdentifyArg) ([]*DelegatedRouting_IdentifyResult, error) {
	ctx, cancel := pd11.WithCancel(ctx)
	defer cancel()
	ch, err := c.Identify_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*DelegatedRouting_IdentifyResult
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) Identify_Async(ctx pd11.Context, req *DelegatedRouting_IdentifyArg) (<-chan DelegatedRouting_Identify_AsyncResult, error) {
	envelope := &AnonInductive4{
		Identify: req,
	}

	buf, err := pd9.Encode(envelope, pd7.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd10.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd6.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.http+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan DelegatedRouting_Identify_AsyncResult, 1)
	go process_DelegatedRouting_Identify_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_Identify_AsyncResult(ctx pd11.Context, ch chan<- DelegatedRouting_Identify_AsyncResult, r pd4.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd5.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd9.DecodeStreaming(r, pd7.Decode)
		if pd8.Is(err, pd4.EOF) || pd8.Is(err, pd4.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd5.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd5.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd5.ErrService{Cause: pd8.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Identify == nil {
			continue
		}
		ch <- DelegatedRouting_Identify_AsyncResult{Resp: env.Identify}
	}
}


func (c *client_DelegatedRouting) FindProviders(ctx pd11.Context, req *FindProvidersRequest) ([]*FindProvidersResponse, error) {
	ctx, cancel := pd11.WithCancel(ctx)
	defer cancel()
	ch, err := c.FindProviders_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*FindProvidersResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) FindProviders_Async(ctx pd11.Context, req *FindProvidersRequest) (<-chan DelegatedRouting_FindProviders_AsyncResult, error) {
	envelope := &AnonInductive4{
		FindProviders: req,
	}

	buf, err := pd9.Encode(envelope, pd7.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd10.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd6.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.http+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan DelegatedRouting_FindProviders_AsyncResult, 1)
	go process_DelegatedRouting_FindProviders_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_FindProviders_AsyncResult(ctx pd11.Context, ch chan<- DelegatedRouting_FindProviders_AsyncResult, r pd4.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_FindProviders_AsyncResult{Err: pd5.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd9.DecodeStreaming(r, pd7.Decode)
		if pd8.Is(err, pd4.EOF) || pd8.Is(err, pd4.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_FindProviders_AsyncResult{Err: pd5.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_FindProviders_AsyncResult{Err: pd5.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_FindProviders_AsyncResult{Err: pd5.ErrService{Cause: pd8.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.FindProviders == nil {
			continue
		}
		ch <- DelegatedRouting_FindProviders_AsyncResult{Resp: env.FindProviders}
	}
}


func (c *client_DelegatedRouting) GetIPNS(ctx pd11.Context, req *GetIPNSRequest) ([]*GetIPNSResponse, error) {
	ctx, cancel := pd11.WithCancel(ctx)
	defer cancel()
	ch, err := c.GetIPNS_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*GetIPNSResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) GetIPNS_Async(ctx pd11.Context, req *GetIPNSRequest) (<-chan DelegatedRouting_GetIPNS_AsyncResult, error) {
	envelope := &AnonInductive4{
		GetIPNS: req,
	}

	buf, err := pd9.Encode(envelope, pd7.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd10.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd6.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.http+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan DelegatedRouting_GetIPNS_AsyncResult, 1)
	go process_DelegatedRouting_GetIPNS_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_GetIPNS_AsyncResult(ctx pd11.Context, ch chan<- DelegatedRouting_GetIPNS_AsyncResult, r pd4.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_GetIPNS_AsyncResult{Err: pd5.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd9.DecodeStreaming(r, pd7.Decode)
		if pd8.Is(err, pd4.EOF) || pd8.Is(err, pd4.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_GetIPNS_AsyncResult{Err: pd5.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_GetIPNS_AsyncResult{Err: pd5.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_GetIPNS_AsyncResult{Err: pd5.ErrService{Cause: pd8.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.GetIPNS == nil {
			continue
		}
		ch <- DelegatedRouting_GetIPNS_AsyncResult{Resp: env.GetIPNS}
	}
}


func (c *client_DelegatedRouting) PutIPNS(ctx pd11.Context, req *PutIPNSRequest) ([]*PutIPNSResponse, error) {
	ctx, cancel := pd11.WithCancel(ctx)
	defer cancel()
	ch, err := c.PutIPNS_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*PutIPNSResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) PutIPNS_Async(ctx pd11.Context, req *PutIPNSRequest) (<-chan DelegatedRouting_PutIPNS_AsyncResult, error) {
	envelope := &AnonInductive4{
		PutIPNS: req,
	}

	buf, err := pd9.Encode(envelope, pd7.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd10.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd6.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.http+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan DelegatedRouting_PutIPNS_AsyncResult, 1)
	go process_DelegatedRouting_PutIPNS_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_PutIPNS_AsyncResult(ctx pd11.Context, ch chan<- DelegatedRouting_PutIPNS_AsyncResult, r pd4.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_PutIPNS_AsyncResult{Err: pd5.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd9.DecodeStreaming(r, pd7.Decode)
		if pd8.Is(err, pd4.EOF) || pd8.Is(err, pd4.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_PutIPNS_AsyncResult{Err: pd5.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_PutIPNS_AsyncResult{Err: pd5.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_PutIPNS_AsyncResult{Err: pd5.ErrService{Cause: pd8.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.PutIPNS == nil {
			continue
		}
		ch <- DelegatedRouting_PutIPNS_AsyncResult{Resp: env.PutIPNS}
	}
}


var logger_server_DelegatedRouting = pd13.Logger("service/server/DelegatedRouting")

type DelegatedRouting_Server interface {

	FindProviders(ctx pd11.Context, req *FindProvidersRequest, respCh chan<- *DelegatedRouting_FindProviders_AsyncResult) error
	GetIPNS(ctx pd11.Context, req *GetIPNSRequest, respCh chan<- *DelegatedRouting_GetIPNS_AsyncResult) error
	PutIPNS(ctx pd11.Context, req *PutIPNSRequest, respCh chan<- *DelegatedRouting_PutIPNS_AsyncResult) error
}

func DelegatedRouting_AsyncHandler(s DelegatedRouting_Server) pd12.HandlerFunc {
	return func(writer pd12.ResponseWriter, request *pd12.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd9.Decode([]byte(msg), pd7.Decode)
		if err != nil {
			logger_server_DelegatedRouting.Errorf("received request not decodeable (%v)", err)
			writer.WriteHeader(400)
			return
		}
		env := &AnonInductive4{}
		if err = env.Parse(n); err != nil {
			logger_server_DelegatedRouting.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		writer.Header()["Content-Type"] = []string{
			"application/vnd.ipfs.http+dag-json; version=1",
		}

		// demultiplex request
		switch {

		case env.FindProviders != nil:
			ch := make(chan *DelegatedRouting_FindProviders_AsyncResult)
			if err = s.FindProviders(pd11.Background(), env.FindProviders, ch); err != nil {
				logger_server_DelegatedRouting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{ Error: &DelegatedRouting_Error{Code: pd1.String(resp.Err.Error())} }
				} else {
					env = &AnonInductive5{ FindProviders: resp.Resp }
				}
				var buf pd6.Buffer
				if err = pd9.EncodeStreaming(&buf, env, pd7.Encode); err != nil {
					logger_server_DelegatedRouting.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
		}

		case env.GetIPNS != nil:
			ch := make(chan *DelegatedRouting_GetIPNS_AsyncResult)
			if err = s.GetIPNS(pd11.Background(), env.GetIPNS, ch); err != nil {
				logger_server_DelegatedRouting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{ Error: &DelegatedRouting_Error{Code: pd1.String(resp.Err.Error())} }
				} else {
					env = &AnonInductive5{ GetIPNS: resp.Resp }
				}
				var buf pd6.Buffer
				if err = pd9.EncodeStreaming(&buf, env, pd7.Encode); err != nil {
					logger_server_DelegatedRouting.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
		}

		case env.PutIPNS != nil:
			ch := make(chan *DelegatedRouting_PutIPNS_AsyncResult)
			if err = s.PutIPNS(pd11.Background(), env.PutIPNS, ch); err != nil {
				logger_server_DelegatedRouting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{ Error: &DelegatedRouting_Error{Code: pd1.String(resp.Err.Error())} }
				} else {
					env = &AnonInductive5{ PutIPNS: resp.Resp }
				}
				var buf pd6.Buffer
				if err = pd9.EncodeStreaming(&buf, env, pd7.Encode); err != nil {
					logger_server_DelegatedRouting.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
		}


		case env.Identify != nil:
			var env *AnonInductive5
			env = &AnonInductive5{
				Identify: &DelegatedRouting_IdentifyResult{
					Methods: []pd1.String{
						"FindProviders",
						"GetIPNS",
						"PutIPNS",

					},
				},
			}
			var buf pd6.Buffer
			if err = pd9.EncodeStreaming(&buf, env, pd7.Encode); err != nil {
				logger_server_DelegatedRouting.Errorf("cannot encode identify response (%v)", err)
				writer.WriteHeader(500)
				return
			}
			buf.WriteByte("\n"[0])
			writer.Write(buf.Bytes())

		default:
			logger_server_DelegatedRouting.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}

// -- protocol type FindProvidersRequest --

type FindProvidersRequest struct {
		Key LinkToAny

}

func (x FindProvidersRequest) Node() pd2.Node {
	return x
}

func (x *FindProvidersRequest) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Key": x.Key.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Key":
			if _, notParsed := fieldMap["Key"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Key")
			}
			if err := x.Key.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Key")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type FindProvidersRequest_MapIterator struct {
	i int64
	s *FindProvidersRequest
}

func (x *FindProvidersRequest_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Key"), x.s.Key.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *FindProvidersRequest_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x FindProvidersRequest) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x FindProvidersRequest) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Key":
		return x.Key.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.Key.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Key":
		return x.Key.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) MapIterator() pd2.MapIterator {
	return &FindProvidersRequest_MapIterator{-1, &x}
}

func (x FindProvidersRequest) ListIterator() pd2.ListIterator {
	return nil
}

func (x FindProvidersRequest) Length() int64 {
	return 1
}

func (x FindProvidersRequest) IsAbsent() bool {
	return false
}

func (x FindProvidersRequest) IsNull() bool {
	return false
}

func (x FindProvidersRequest) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x FindProvidersRequest) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x FindProvidersRequest) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x FindProvidersRequest) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x FindProvidersRequest) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x FindProvidersRequest) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type ProvidersList --

type ProvidersList []Provider

func (v ProvidersList) Node() pd2.Node {
	return v
}

func (v *ProvidersList) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(ProvidersList, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (ProvidersList) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (ProvidersList) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (ProvidersList) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v ProvidersList) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v ProvidersList) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (ProvidersList) MapIterator() pd2.MapIterator {
	return nil
}

func (v ProvidersList) ListIterator() pd2.ListIterator {
	return &ProvidersList_ListIterator{v, 0}
}

func (v ProvidersList) Length() int64 {
	return int64(len(v))
}

func (ProvidersList) IsAbsent() bool {
	return false
}

func (ProvidersList) IsNull() bool {
	return false
}

func (v ProvidersList) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (ProvidersList) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (ProvidersList) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (ProvidersList) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (ProvidersList) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (ProvidersList) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (ProvidersList) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type ProvidersList_ListIterator struct {
	list ProvidersList
	at   int64
}

func (iter *ProvidersList_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *ProvidersList_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type FindProvidersResponse --

type FindProvidersResponse struct {
		Providers ProvidersList

}

func (x FindProvidersResponse) Node() pd2.Node {
	return x
}

func (x *FindProvidersResponse) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Providers": x.Providers.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Providers":
			if _, notParsed := fieldMap["Providers"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Providers")
			}
			if err := x.Providers.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Providers")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type FindProvidersResponse_MapIterator struct {
	i int64
	s *FindProvidersResponse
}

func (x *FindProvidersResponse_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Providers"), x.s.Providers.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *FindProvidersResponse_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x FindProvidersResponse) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x FindProvidersResponse) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Providers":
		return x.Providers.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.Providers.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Providers":
		return x.Providers.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) MapIterator() pd2.MapIterator {
	return &FindProvidersResponse_MapIterator{-1, &x}
}

func (x FindProvidersResponse) ListIterator() pd2.ListIterator {
	return nil
}

func (x FindProvidersResponse) Length() int64 {
	return 1
}

func (x FindProvidersResponse) IsAbsent() bool {
	return false
}

func (x FindProvidersResponse) IsNull() bool {
	return false
}

func (x FindProvidersResponse) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x FindProvidersResponse) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x FindProvidersResponse) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x FindProvidersResponse) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x FindProvidersResponse) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x FindProvidersResponse) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type GetIPNSRequest --

type GetIPNSRequest struct {
		ID pd1.Bytes

}

func (x GetIPNSRequest) Node() pd2.Node {
	return x
}

func (x *GetIPNSRequest) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"ID": x.ID.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "ID":
			if _, notParsed := fieldMap["ID"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "ID")
			}
			if err := x.ID.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "ID")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type GetIPNSRequest_MapIterator struct {
	i int64
	s *GetIPNSRequest
}

func (x *GetIPNSRequest_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("ID"), x.s.ID.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *GetIPNSRequest_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x GetIPNSRequest) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x GetIPNSRequest) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "ID":
		return x.ID.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.ID.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "ID":
		return x.ID.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) MapIterator() pd2.MapIterator {
	return &GetIPNSRequest_MapIterator{-1, &x}
}

func (x GetIPNSRequest) ListIterator() pd2.ListIterator {
	return nil
}

func (x GetIPNSRequest) Length() int64 {
	return 1
}

func (x GetIPNSRequest) IsAbsent() bool {
	return false
}

func (x GetIPNSRequest) IsNull() bool {
	return false
}

func (x GetIPNSRequest) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x GetIPNSRequest) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x GetIPNSRequest) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x GetIPNSRequest) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x GetIPNSRequest) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x GetIPNSRequest) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type GetIPNSResponse --

type GetIPNSResponse struct {
		Record pd1.Bytes

}

func (x GetIPNSResponse) Node() pd2.Node {
	return x
}

func (x *GetIPNSResponse) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Record": x.Record.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Record":
			if _, notParsed := fieldMap["Record"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Record")
			}
			if err := x.Record.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Record")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type GetIPNSResponse_MapIterator struct {
	i int64
	s *GetIPNSResponse
}

func (x *GetIPNSResponse_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Record"), x.s.Record.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *GetIPNSResponse_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x GetIPNSResponse) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x GetIPNSResponse) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Record":
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Record":
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) MapIterator() pd2.MapIterator {
	return &GetIPNSResponse_MapIterator{-1, &x}
}

func (x GetIPNSResponse) ListIterator() pd2.ListIterator {
	return nil
}

func (x GetIPNSResponse) Length() int64 {
	return 1
}

func (x GetIPNSResponse) IsAbsent() bool {
	return false
}

func (x GetIPNSResponse) IsNull() bool {
	return false
}

func (x GetIPNSResponse) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x GetIPNSResponse) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x GetIPNSResponse) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x GetIPNSResponse) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x GetIPNSResponse) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x GetIPNSResponse) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type PutIPNSRequest --

type PutIPNSRequest struct {
		ID pd1.Bytes
		Record pd1.Bytes

}

func (x PutIPNSRequest) Node() pd2.Node {
	return x
}

func (x *PutIPNSRequest) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"ID": x.ID.Parse,
		"Record": x.Record.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "ID":
			if _, notParsed := fieldMap["ID"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "ID")
			}
			if err := x.ID.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "ID")
			case "Record":
			if _, notParsed := fieldMap["Record"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Record")
			}
			if err := x.Record.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Record")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type PutIPNSRequest_MapIterator struct {
	i int64
	s *PutIPNSRequest
}

func (x *PutIPNSRequest_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("ID"), x.s.ID.Node(), nil
			case 1:
			return pd1.String("Record"), x.s.Record.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *PutIPNSRequest_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x PutIPNSRequest) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x PutIPNSRequest) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "ID":
		return x.ID.Node(), nil
		case "Record":
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.ID.Node(), nil
		case 1:
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "ID":
		return x.ID.Node(), nil
		case "1", "Record":
		return x.Record.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) MapIterator() pd2.MapIterator {
	return &PutIPNSRequest_MapIterator{-1, &x}
}

func (x PutIPNSRequest) ListIterator() pd2.ListIterator {
	return nil
}

func (x PutIPNSRequest) Length() int64 {
	return 2
}

func (x PutIPNSRequest) IsAbsent() bool {
	return false
}

func (x PutIPNSRequest) IsNull() bool {
	return false
}

func (x PutIPNSRequest) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x PutIPNSRequest) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x PutIPNSRequest) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x PutIPNSRequest) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x PutIPNSRequest) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x PutIPNSRequest) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type PutIPNSResponse --

type PutIPNSResponse struct {

}

func (x PutIPNSResponse) Node() pd2.Node {
	return x
}

func (x *PutIPNSResponse) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type PutIPNSResponse_MapIterator struct {
	i int64
	s *PutIPNSResponse
}

func (x *PutIPNSResponse_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *PutIPNSResponse_MapIterator) Done() bool {
	return x.i + 1 >= 0
}

func (x PutIPNSResponse) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x PutIPNSResponse) LookupByString(key string) (pd2.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) MapIterator() pd2.MapIterator {
	return &PutIPNSResponse_MapIterator{-1, &x}
}

func (x PutIPNSResponse) ListIterator() pd2.ListIterator {
	return nil
}

func (x PutIPNSResponse) Length() int64 {
	return 0
}

func (x PutIPNSResponse) IsAbsent() bool {
	return false
}

func (x PutIPNSResponse) IsNull() bool {
	return false
}

func (x PutIPNSResponse) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x PutIPNSResponse) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x PutIPNSResponse) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x PutIPNSResponse) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x PutIPNSResponse) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x PutIPNSResponse) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type LinkToAny --

type LinkToAny pd14.Cid

func (v *LinkToAny) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Link {
		return pd1.ErrNA
	} else {
		ipldLink, _ := n.AsLink()
		// TODO: Is there a more general way to convert ipld.Link interface into a concrete user object?
		cidLink, ok := ipldLink.(pd15.Link)
		if !ok {
			return pd3.Errorf("only cid links are supported")
		} else {
			*v = LinkToAny(cidLink.Cid)
			return nil
		}
	}
}

func (v LinkToAny) Node() pd2.Node {
	return v
}

func (LinkToAny) Kind() pd2.Kind {
	return pd2.Kind_Link
}

func (LinkToAny) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (LinkToAny) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (LinkToAny) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (LinkToAny) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (LinkToAny) MapIterator() pd2.MapIterator {
	return nil
}

func (LinkToAny) ListIterator() pd2.ListIterator {
	return nil
}

func (LinkToAny) Length() int64 {
	return -1
}

func (LinkToAny) IsAbsent() bool {
	return false
}

func (LinkToAny) IsNull() bool {
	return false
}

func (LinkToAny) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (v LinkToAny) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (LinkToAny) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (LinkToAny) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (LinkToAny) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (v LinkToAny) AsLink() (pd2.Link, error) {
	return pd15.Link{Cid: pd14.Cid(v)}, nil
}

func (LinkToAny) Prototype() pd2.NodePrototype {
	return nil // not needed
}
// -- protocol type Provider --

type Provider struct {
		ProviderNode Node
		ProviderProto TransferProtocolList

}

func (x Provider) Node() pd2.Node {
	return x
}

func (x *Provider) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"Node": x.ProviderNode.Parse,
		"Proto": x.ProviderProto.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Node":
			if _, notParsed := fieldMap["Node"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Node")
			}
			if err := x.ProviderNode.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Node")
			case "Proto":
			if _, notParsed := fieldMap["Proto"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Proto")
			}
			if err := x.ProviderProto.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Proto")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type Provider_MapIterator struct {
	i int64
	s *Provider
}

func (x *Provider_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("Node"), x.s.ProviderNode.Node(), nil
			case 1:
			return pd1.String("Proto"), x.s.ProviderProto.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Provider_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x Provider) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Provider) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "Node":
		return x.ProviderNode.Node(), nil
		case "Proto":
		return x.ProviderProto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.ProviderNode.Node(), nil
		case 1:
		return x.ProviderProto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "Node":
		return x.ProviderNode.Node(), nil
		case "1", "Proto":
		return x.ProviderProto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) MapIterator() pd2.MapIterator {
	return &Provider_MapIterator{-1, &x}
}

func (x Provider) ListIterator() pd2.ListIterator {
	return nil
}

func (x Provider) Length() int64 {
	return 2
}

func (x Provider) IsAbsent() bool {
	return false
}

func (x Provider) IsNull() bool {
	return false
}

func (x Provider) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Provider) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Provider) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Provider) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Provider) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Provider) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Provider) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type TransferProtocolList --

type TransferProtocolList []TransferProtocol

func (v TransferProtocolList) Node() pd2.Node {
	return v
}

func (v *TransferProtocolList) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(TransferProtocolList, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (TransferProtocolList) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (TransferProtocolList) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (TransferProtocolList) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v TransferProtocolList) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v TransferProtocolList) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (TransferProtocolList) MapIterator() pd2.MapIterator {
	return nil
}

func (v TransferProtocolList) ListIterator() pd2.ListIterator {
	return &TransferProtocolList_ListIterator{v, 0}
}

func (v TransferProtocolList) Length() int64 {
	return int64(len(v))
}

func (TransferProtocolList) IsAbsent() bool {
	return false
}

func (TransferProtocolList) IsNull() bool {
	return false
}

func (v TransferProtocolList) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (TransferProtocolList) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (TransferProtocolList) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (TransferProtocolList) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (TransferProtocolList) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (TransferProtocolList) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (TransferProtocolList) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type TransferProtocolList_ListIterator struct {
	list TransferProtocolList
	at   int64
}

func (iter *TransferProtocolList_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *TransferProtocolList_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type Node --

type Node struct {
		Peer *Peer


		DefaultKey string
		DefaultValue *pd1.Any

}

func (x *Node) Parse(n pd2.Node) error {
	*x = Node{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "peer":
		var y Peer
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Peer = &y
		return nil


	default:
		var y pd1.Any
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.DefaultKey = k
		x.DefaultValue = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type Node_MapIterator struct {
	done bool
	s    *Node
}

func (x *Node_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Peer != nil:
			return pd1.String("peer"), x.s.Peer.Node(), nil


	case x.s.DefaultValue != nil:
		return pd1.String(x.s.DefaultKey), x.s.DefaultValue.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *Node_MapIterator) Done() bool {
	return x.done
}

func (x Node) Node() pd2.Node {
	return x
}

func (x Node) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Node) LookupByString(key string) (pd2.Node, error) {
	switch {
		case x.Peer != nil && key == "peer":
		return x.Peer.Node(), nil


	case x.DefaultValue != nil && key == x.DefaultKey:
		return x.DefaultValue.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Node) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x Node) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x Node) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "peer":
		return x.Peer.Node(), nil


	case x.DefaultKey:
		return x.DefaultValue.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Node) MapIterator() pd2.MapIterator {
	return &Node_MapIterator{false, &x}
}

func (x Node) ListIterator() pd2.ListIterator {
	return nil
}

func (x Node) Length() int64 {
	return 1
}

func (x Node) IsAbsent() bool {
	return false
}

func (x Node) IsNull() bool {
	return false
}

func (x Node) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Node) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Node) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Node) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Node) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Node) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Node) Prototype() pd2.NodePrototype {
	return nil
}
// -- protocol type AnonList18 --

type AnonList18 []pd1.Bytes

func (v AnonList18) Node() pd2.Node {
	return v
}

func (v *AnonList18) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList18, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList18) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList18) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList18) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList18) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList18) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList18) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList18) ListIterator() pd2.ListIterator {
	return &AnonList18_ListIterator{v, 0}
}

func (v AnonList18) Length() int64 {
	return int64(len(v))
}

func (AnonList18) IsAbsent() bool {
	return false
}

func (AnonList18) IsNull() bool {
	return false
}

func (v AnonList18) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (AnonList18) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList18) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList18) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList18) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList18) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList18) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList18_ListIterator struct {
	list AnonList18
	at   int64
}

func (iter *AnonList18_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList18_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type Peer --

type Peer struct {
		ID pd1.Bytes
		Multiaddresses AnonList18

}

func (x Peer) Node() pd2.Node {
	return x
}

func (x *Peer) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"ID": x.ID.Parse,
		"Multiaddresses": x.Multiaddresses.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "ID":
			if _, notParsed := fieldMap["ID"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "ID")
			}
			if err := x.ID.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "ID")
			case "Multiaddresses":
			if _, notParsed := fieldMap["Multiaddresses"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Multiaddresses")
			}
			if err := x.Multiaddresses.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Multiaddresses")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type Peer_MapIterator struct {
	i int64
	s *Peer
}

func (x *Peer_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("ID"), x.s.ID.Node(), nil
			case 1:
			return pd1.String("Multiaddresses"), x.s.Multiaddresses.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Peer_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x Peer) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Peer) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "ID":
		return x.ID.Node(), nil
		case "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.ID.Node(), nil
		case 1:
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "ID":
		return x.ID.Node(), nil
		case "1", "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) MapIterator() pd2.MapIterator {
	return &Peer_MapIterator{-1, &x}
}

func (x Peer) ListIterator() pd2.ListIterator {
	return nil
}

func (x Peer) Length() int64 {
	return 2
}

func (x Peer) IsAbsent() bool {
	return false
}

func (x Peer) IsNull() bool {
	return false
}

func (x Peer) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x Peer) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Peer) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Peer) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Peer) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Peer) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Peer) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type TransferProtocol --

type TransferProtocol struct {
		Bitswap *BitswapProtocol
		GraphSyncFILv1 *GraphSyncFILv1Protocol


		DefaultKey string
		DefaultValue *pd1.Any

}

func (x *TransferProtocol) Parse(n pd2.Node) error {
	*x = TransferProtocol{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "2304":
		var y BitswapProtocol
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Bitswap = &y
		return nil
	case "2320":
		var y GraphSyncFILv1Protocol
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GraphSyncFILv1 = &y
		return nil


	default:
		var y pd1.Any
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.DefaultKey = k
		x.DefaultValue = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type TransferProtocol_MapIterator struct {
	done bool
	s    *TransferProtocol
}

func (x *TransferProtocol_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Bitswap != nil:
			return pd1.String("2304"), x.s.Bitswap.Node(), nil
			case x.s.GraphSyncFILv1 != nil:
			return pd1.String("2320"), x.s.GraphSyncFILv1.Node(), nil


	case x.s.DefaultValue != nil:
		return pd1.String(x.s.DefaultKey), x.s.DefaultValue.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *TransferProtocol_MapIterator) Done() bool {
	return x.done
}

func (x TransferProtocol) Node() pd2.Node {
	return x
}

func (x TransferProtocol) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x TransferProtocol) LookupByString(key string) (pd2.Node, error) {
	switch {
		case x.Bitswap != nil && key == "2304":
		return x.Bitswap.Node(), nil
		case x.GraphSyncFILv1 != nil && key == "2320":
		return x.GraphSyncFILv1.Node(), nil


	case x.DefaultValue != nil && key == x.DefaultKey:
		return x.DefaultValue.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x TransferProtocol) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x TransferProtocol) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x TransferProtocol) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "2304":
		return x.Bitswap.Node(), nil
		case "2320":
		return x.GraphSyncFILv1.Node(), nil


	case x.DefaultKey:
		return x.DefaultValue.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x TransferProtocol) MapIterator() pd2.MapIterator {
	return &TransferProtocol_MapIterator{false, &x}
}

func (x TransferProtocol) ListIterator() pd2.ListIterator {
	return nil
}

func (x TransferProtocol) Length() int64 {
	return 1
}

func (x TransferProtocol) IsAbsent() bool {
	return false
}

func (x TransferProtocol) IsNull() bool {
	return false
}

func (x TransferProtocol) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x TransferProtocol) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x TransferProtocol) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x TransferProtocol) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x TransferProtocol) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x TransferProtocol) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x TransferProtocol) Prototype() pd2.NodePrototype {
	return nil
}
// -- protocol type BitswapProtocol --

type BitswapProtocol struct {

}

func (x BitswapProtocol) Node() pd2.Node {
	return x
}

func (x *BitswapProtocol) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
		
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type BitswapProtocol_MapIterator struct {
	i int64
	s *BitswapProtocol
}

func (x *BitswapProtocol_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *BitswapProtocol_MapIterator) Done() bool {
	return x.i + 1 >= 0
}

func (x BitswapProtocol) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x BitswapProtocol) LookupByString(key string) (pd2.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) MapIterator() pd2.MapIterator {
	return &BitswapProtocol_MapIterator{-1, &x}
}

func (x BitswapProtocol) ListIterator() pd2.ListIterator {
	return nil
}

func (x BitswapProtocol) Length() int64 {
	return 0
}

func (x BitswapProtocol) IsAbsent() bool {
	return false
}

func (x BitswapProtocol) IsNull() bool {
	return false
}

func (x BitswapProtocol) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x BitswapProtocol) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x BitswapProtocol) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x BitswapProtocol) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x BitswapProtocol) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x BitswapProtocol) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type GraphSyncFILv1Protocol --

type GraphSyncFILv1Protocol struct {
		PieceCID LinkToAny
		VerifiedDeal pd1.Bool
		FastRetrieval pd1.Bool

}

func (x GraphSyncFILv1Protocol) Node() pd2.Node {
	return x
}

func (x *GraphSyncFILv1Protocol) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
				"PieceCID": x.PieceCID.Parse,
		"VerifiedDeal": x.VerifiedDeal.Parse,
		"FastRetrieval": x.FastRetrieval.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "PieceCID":
			if _, notParsed := fieldMap["PieceCID"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "PieceCID")
			}
			if err := x.PieceCID.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "PieceCID")
			case "VerifiedDeal":
			if _, notParsed := fieldMap["VerifiedDeal"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "VerifiedDeal")
			}
			if err := x.VerifiedDeal.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "VerifiedDeal")
			case "FastRetrieval":
			if _, notParsed := fieldMap["FastRetrieval"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "FastRetrieval")
			}
			if err := x.FastRetrieval.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "FastRetrieval")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type GraphSyncFILv1Protocol_MapIterator struct {
	i int64
	s *GraphSyncFILv1Protocol
}

func (x *GraphSyncFILv1Protocol_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd1.String("PieceCID"), x.s.PieceCID.Node(), nil
			case 1:
			return pd1.String("VerifiedDeal"), x.s.VerifiedDeal.Node(), nil
			case 2:
			return pd1.String("FastRetrieval"), x.s.FastRetrieval.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *GraphSyncFILv1Protocol_MapIterator) Done() bool {
	return x.i + 1 >= 3
}

func (x GraphSyncFILv1Protocol) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x GraphSyncFILv1Protocol) LookupByString(key string) (pd2.Node, error) {
	switch key {
		case "PieceCID":
		return x.PieceCID.Node(), nil
		case "VerifiedDeal":
		return x.VerifiedDeal.Node(), nil
		case "FastRetrieval":
		return x.FastRetrieval.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
		case 0:
		return x.PieceCID.Node(), nil
		case 1:
		return x.VerifiedDeal.Node(), nil
		case 2:
		return x.FastRetrieval.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
		case "0", "PieceCID":
		return x.PieceCID.Node(), nil
		case "1", "VerifiedDeal":
		return x.VerifiedDeal.Node(), nil
		case "2", "FastRetrieval":
		return x.FastRetrieval.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) MapIterator() pd2.MapIterator {
	return &GraphSyncFILv1Protocol_MapIterator{-1, &x}
}

func (x GraphSyncFILv1Protocol) ListIterator() pd2.ListIterator {
	return nil
}

func (x GraphSyncFILv1Protocol) Length() int64 {
	return 3
}

func (x GraphSyncFILv1Protocol) IsAbsent() bool {
	return false
}

func (x GraphSyncFILv1Protocol) IsNull() bool {
	return false
}

func (x GraphSyncFILv1Protocol) AsBool() (bool, error) {
	return false, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x GraphSyncFILv1Protocol) Prototype() pd2.NodePrototype {
	return nil
}