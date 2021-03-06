// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sayhello

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type HelloAndSum interface {
  // Parameters:
  //  - Msg
  Sayhello(ctx context.Context, msg string) (r string, err error)
  // Parameters:
  //  - A
  //  - B
  Sum(ctx context.Context, a int32, b int32) (r string, err error)
}

type HelloAndSumClient struct {
  c thrift.TClient
}

func NewHelloAndSumClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HelloAndSumClient {
  return &HelloAndSumClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewHelloAndSumClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HelloAndSumClient {
  return &HelloAndSumClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewHelloAndSumClient(c thrift.TClient) *HelloAndSumClient {
  return &HelloAndSumClient{
    c: c,
  }
}

func (p *HelloAndSumClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Msg
func (p *HelloAndSumClient) Sayhello(ctx context.Context, msg string) (r string, err error) {
  var _args0 HelloAndSumSayhelloArgs
  _args0.Msg = msg
  var _result1 HelloAndSumSayhelloResult
  if err = p.Client_().Call(ctx, "sayhello", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - A
//  - B
func (p *HelloAndSumClient) Sum(ctx context.Context, a int32, b int32) (r string, err error) {
  var _args2 HelloAndSumSumArgs
  _args2.A = a
  _args2.B = b
  var _result3 HelloAndSumSumResult
  if err = p.Client_().Call(ctx, "sum", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

type HelloAndSumProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler HelloAndSum
}

func (p *HelloAndSumProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *HelloAndSumProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *HelloAndSumProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewHelloAndSumProcessor(handler HelloAndSum) *HelloAndSumProcessor {

  self4 := &HelloAndSumProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["sayhello"] = &helloAndSumProcessorSayhello{handler:handler}
  self4.processorMap["sum"] = &helloAndSumProcessorSum{handler:handler}
return self4
}

func (p *HelloAndSumProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x5

}

type helloAndSumProcessorSayhello struct {
  handler HelloAndSum
}

func (p *helloAndSumProcessorSayhello) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := HelloAndSumSayhelloArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("sayhello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := HelloAndSumSayhelloResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Sayhello(ctx, args.Msg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing sayhello: " + err2.Error())
    oprot.WriteMessageBegin("sayhello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("sayhello", thrift.REPLY, seqId); err2 != nil {
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

type helloAndSumProcessorSum struct {
  handler HelloAndSum
}

func (p *helloAndSumProcessorSum) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := HelloAndSumSumArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("sum", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := HelloAndSumSumResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Sum(ctx, args.A, args.B); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing sum: " + err2.Error())
    oprot.WriteMessageBegin("sum", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("sum", thrift.REPLY, seqId); err2 != nil {
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


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Msg
type HelloAndSumSayhelloArgs struct {
  Msg string `thrift:"msg,1" db:"msg" json:"msg"`
}

func NewHelloAndSumSayhelloArgs() *HelloAndSumSayhelloArgs {
  return &HelloAndSumSayhelloArgs{}
}


func (p *HelloAndSumSayhelloArgs) GetMsg() string {
  return p.Msg
}
func (p *HelloAndSumSayhelloArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *HelloAndSumSayhelloArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *HelloAndSumSayhelloArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sayhello_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloAndSumSayhelloArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err) }
  return err
}

func (p *HelloAndSumSayhelloArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloAndSumSayhelloArgs(%+v)", *p)
}

// Attributes:
//  - Success
type HelloAndSumSayhelloResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewHelloAndSumSayhelloResult() *HelloAndSumSayhelloResult {
  return &HelloAndSumSayhelloResult{}
}

var HelloAndSumSayhelloResult_Success_DEFAULT string
func (p *HelloAndSumSayhelloResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return HelloAndSumSayhelloResult_Success_DEFAULT
  }
return *p.Success
}
func (p *HelloAndSumSayhelloResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *HelloAndSumSayhelloResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *HelloAndSumSayhelloResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *HelloAndSumSayhelloResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sayhello_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloAndSumSayhelloResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *HelloAndSumSayhelloResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloAndSumSayhelloResult(%+v)", *p)
}

// Attributes:
//  - A
//  - B
type HelloAndSumSumArgs struct {
  A int32 `thrift:"a,1" db:"a" json:"a"`
  B int32 `thrift:"b,2" db:"b" json:"b"`
}

func NewHelloAndSumSumArgs() *HelloAndSumSumArgs {
  return &HelloAndSumSumArgs{}
}


func (p *HelloAndSumSumArgs) GetA() int32 {
  return p.A
}

func (p *HelloAndSumSumArgs) GetB() int32 {
  return p.B
}
func (p *HelloAndSumSumArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *HelloAndSumSumArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.A = v
}
  return nil
}

func (p *HelloAndSumSumArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.B = v
}
  return nil
}

func (p *HelloAndSumSumArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sum_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloAndSumSumArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("a", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:a: ", p), err) }
  if err := oprot.WriteI32(int32(p.A)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.a (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:a: ", p), err) }
  return err
}

func (p *HelloAndSumSumArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("b", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:b: ", p), err) }
  if err := oprot.WriteI32(int32(p.B)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.b (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:b: ", p), err) }
  return err
}

func (p *HelloAndSumSumArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloAndSumSumArgs(%+v)", *p)
}

// Attributes:
//  - Success
type HelloAndSumSumResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewHelloAndSumSumResult() *HelloAndSumSumResult {
  return &HelloAndSumSumResult{}
}

var HelloAndSumSumResult_Success_DEFAULT string
func (p *HelloAndSumSumResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return HelloAndSumSumResult_Success_DEFAULT
  }
return *p.Success
}
func (p *HelloAndSumSumResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *HelloAndSumSumResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *HelloAndSumSumResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *HelloAndSumSumResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sum_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloAndSumSumResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *HelloAndSumSumResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloAndSumSumResult(%+v)", *p)
}


