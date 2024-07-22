// Code generated by thriftgo (0.3.14). DO NOT EDIT.

package shop

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// import "github.com/hertz-contrib/sessions"
type Empty struct {
}

func NewEmpty() *Empty {
	return &Empty{}
}

func (p *Empty) InitDefault() {
}

var fieldIDToName_Empty = map[int16]string{}

func (p *Empty) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Empty) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("Empty"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Empty) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Empty(%+v)", *p)

}

type ShopType struct {
	ID   int64  `thrift:"id,1" form:"id" json:"id" query:"id"`
	Name string `thrift:"name,2" form:"name" json:"name" query:"name"`
	Icon string `thrift:"icon,3" form:"icon" json:"icon" query:"icon"`
	Sort int32  `thrift:"sort,4" form:"sort" json:"sort" query:"sort"`
}

func NewShopType() *ShopType {
	return &ShopType{}
}

func (p *ShopType) InitDefault() {
}

func (p *ShopType) GetID() (v int64) {
	return p.ID
}

func (p *ShopType) GetName() (v string) {
	return p.Name
}

func (p *ShopType) GetIcon() (v string) {
	return p.Icon
}

func (p *ShopType) GetSort() (v int32) {
	return p.Sort
}

var fieldIDToName_ShopType = map[int16]string{
	1: "id",
	2: "name",
	3: "icon",
	4: "sort",
}

func (p *ShopType) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ShopType[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ShopType) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ID = _field
	return nil
}
func (p *ShopType) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Name = _field
	return nil
}
func (p *ShopType) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Icon = _field
	return nil
}
func (p *ShopType) ReadField4(iprot thrift.TProtocol) error {

	var _field int32
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Sort = _field
	return nil
}

func (p *ShopType) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ShopType"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ShopType) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ShopType) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ShopType) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("icon", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Icon); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ShopType) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("sort", thrift.I32, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Sort); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ShopType) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopType(%+v)", *p)

}

type ShopService interface {
	ShopList(ctx context.Context, request *Empty) (r []*ShopType, err error)
}

type ShopServiceClient struct {
	c thrift.TClient
}

func NewShopServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ShopServiceClient {
	return &ShopServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewShopServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ShopServiceClient {
	return &ShopServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewShopServiceClient(c thrift.TClient) *ShopServiceClient {
	return &ShopServiceClient{
		c: c,
	}
}

func (p *ShopServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *ShopServiceClient) ShopList(ctx context.Context, request *Empty) (r []*ShopType, err error) {
	var _args ShopServiceShopListArgs
	_args.Request = request
	var _result ShopServiceShopListResult
	if err = p.Client_().Call(ctx, "ShopList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type ShopServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ShopService
}

func (p *ShopServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ShopServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ShopServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewShopServiceProcessor(handler ShopService) *ShopServiceProcessor {
	self := &ShopServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("ShopList", &shopServiceProcessorShopList{handler: handler})
	return self
}
func (p *ShopServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type shopServiceProcessorShopList struct {
	handler ShopService
}

func (p *shopServiceProcessorShopList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ShopServiceShopListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ShopList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ShopServiceShopListResult{}
	var retval []*ShopType
	if retval, err2 = p.handler.ShopList(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ShopList: "+err2.Error())
		oprot.WriteMessageBegin("ShopList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ShopList", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ShopServiceShopListArgs struct {
	Request *Empty `thrift:"request,1"`
}

func NewShopServiceShopListArgs() *ShopServiceShopListArgs {
	return &ShopServiceShopListArgs{}
}

func (p *ShopServiceShopListArgs) InitDefault() {
}

var ShopServiceShopListArgs_Request_DEFAULT *Empty

func (p *ShopServiceShopListArgs) GetRequest() (v *Empty) {
	if !p.IsSetRequest() {
		return ShopServiceShopListArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_ShopServiceShopListArgs = map[int16]string{
	1: "request",
}

func (p *ShopServiceShopListArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *ShopServiceShopListArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ShopServiceShopListArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ShopServiceShopListArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Request = _field
	return nil
}

func (p *ShopServiceShopListArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ShopList_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ShopServiceShopListArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ShopServiceShopListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceShopListArgs(%+v)", *p)

}

type ShopServiceShopListResult struct {
	Success []*ShopType `thrift:"success,0,optional"`
}

func NewShopServiceShopListResult() *ShopServiceShopListResult {
	return &ShopServiceShopListResult{}
}

func (p *ShopServiceShopListResult) InitDefault() {
}

var ShopServiceShopListResult_Success_DEFAULT []*ShopType

func (p *ShopServiceShopListResult) GetSuccess() (v []*ShopType) {
	if !p.IsSetSuccess() {
		return ShopServiceShopListResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_ShopServiceShopListResult = map[int16]string{
	0: "success",
}

func (p *ShopServiceShopListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ShopServiceShopListResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ShopServiceShopListResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ShopServiceShopListResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	_field := make([]*ShopType, 0, size)
	values := make([]ShopType, size)
	for i := 0; i < size; i++ {
		_elem := &values[i]
		_elem.InitDefault()

		if err := _elem.Read(iprot); err != nil {
			return err
		}

		_field = append(_field, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *ShopServiceShopListResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ShopList_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ShopServiceShopListResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return err
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *ShopServiceShopListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceShopListResult(%+v)", *p)

}
