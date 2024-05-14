package imports

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"

	"git.woa.com/vasd_masc_ba/YitihuaOteam/base/gofun/register"
)

var _ = reflect.Int

func init() {
	register.AddPackage("github.com/json-iterator/go", "jsoniter",
		register.NewType("API", reflect.TypeOf(func(jsoniter.API) {}).In(0), ""),
		register.NewType("Any", reflect.TypeOf(func(jsoniter.Any) {}).In(0), ""),
		register.NewConst("ArrayValue", jsoniter.ArrayValue, ""),
		register.NewType("Binding", reflect.TypeOf(func(jsoniter.Binding) {}).In(0), ""),
		register.NewConst("BoolValue", jsoniter.BoolValue, ""),
		register.NewFunction("CastJsonNumber", jsoniter.CastJsonNumber, ""),
		register.NewType("Config", reflect.TypeOf(func(jsoniter.Config) {}).In(0), ""),
		register.NewVar("ConfigCompatibleWithStandardLibrary", &jsoniter.ConfigCompatibleWithStandardLibrary, reflect.TypeOf(func(jsoniter.API) {}).In(0), ""),
		register.NewVar("ConfigDefault", &jsoniter.ConfigDefault, reflect.TypeOf(func(jsoniter.API) {}).In(0), ""),
		register.NewVar("ConfigFastest", &jsoniter.ConfigFastest, reflect.TypeOf(func(jsoniter.API) {}).In(0), ""),
		register.NewType("Decoder", reflect.TypeOf(func(jsoniter.Decoder) {}).In(0), ""),
		register.NewType("DecoderExtension", reflect.TypeOf(func(jsoniter.DecoderExtension) {}).In(0), ""),
		register.NewType("DecoderFunc", reflect.TypeOf(func(jsoniter.DecoderFunc) {}).In(0), ""),
		register.NewType("DummyExtension", reflect.TypeOf(func(jsoniter.DummyExtension) {}).In(0), ""),
		register.NewType("Encoder", reflect.TypeOf(func(jsoniter.Encoder) {}).In(0), ""),
		register.NewType("EncoderExtension", reflect.TypeOf(func(jsoniter.EncoderExtension) {}).In(0), ""),
		register.NewType("EncoderFunc", reflect.TypeOf(func(jsoniter.EncoderFunc) {}).In(0), ""),
		register.NewType("Extension", reflect.TypeOf(func(jsoniter.Extension) {}).In(0), ""),
		register.NewFunction("Get", jsoniter.Get, ""),
		register.NewConst("InvalidValue", jsoniter.InvalidValue, ""),
		register.NewType("IsEmbeddedPtrNil", reflect.TypeOf(func(jsoniter.IsEmbeddedPtrNil) {}).In(0), ""),
		register.NewType("Iterator", reflect.TypeOf(func(jsoniter.Iterator) {}).In(0), ""),
		register.NewType("IteratorPool", reflect.TypeOf(func(jsoniter.IteratorPool) {}).In(0), ""),
		register.NewFunction("Marshal", jsoniter.Marshal, ""),
		register.NewFunction("MarshalIndent", jsoniter.MarshalIndent, ""),
		register.NewFunction("MarshalToString", jsoniter.MarshalToString, ""),
		register.NewFunction("NewDecoder", jsoniter.NewDecoder, ""),
		register.NewFunction("NewEncoder", jsoniter.NewEncoder, ""),
		register.NewFunction("NewIterator", jsoniter.NewIterator, ""),
		register.NewFunction("NewStream", jsoniter.NewStream, ""),
		register.NewConst("NilValue", jsoniter.NilValue, ""),
		register.NewType("Number", reflect.TypeOf(func(jsoniter.Number) {}).In(0), ""),
		register.NewConst("NumberValue", jsoniter.NumberValue, ""),
		register.NewConst("ObjectValue", jsoniter.ObjectValue, ""),
		register.NewType("OptionalDecoder", reflect.TypeOf(func(jsoniter.OptionalDecoder) {}).In(0), ""),
		register.NewType("OptionalEncoder", reflect.TypeOf(func(jsoniter.OptionalEncoder) {}).In(0), ""),
		register.NewFunction("Parse", jsoniter.Parse, ""),
		register.NewFunction("ParseBytes", jsoniter.ParseBytes, ""),
		register.NewFunction("ParseString", jsoniter.ParseString, ""),
		register.NewType("RawMessage", reflect.TypeOf(func(jsoniter.RawMessage) {}).In(0), ""),
		register.NewFunction("RegisterExtension", jsoniter.RegisterExtension, ""),
		register.NewFunction("RegisterFieldDecoder", jsoniter.RegisterFieldDecoder, ""),
		register.NewFunction("RegisterFieldDecoderFunc", jsoniter.RegisterFieldDecoderFunc, ""),
		register.NewFunction("RegisterFieldEncoder", jsoniter.RegisterFieldEncoder, ""),
		register.NewFunction("RegisterFieldEncoderFunc", jsoniter.RegisterFieldEncoderFunc, ""),
		register.NewFunction("RegisterTypeDecoder", jsoniter.RegisterTypeDecoder, ""),
		register.NewFunction("RegisterTypeDecoderFunc", jsoniter.RegisterTypeDecoderFunc, ""),
		register.NewFunction("RegisterTypeEncoder", jsoniter.RegisterTypeEncoder, ""),
		register.NewFunction("RegisterTypeEncoderFunc", jsoniter.RegisterTypeEncoderFunc, ""),
		register.NewType("Stream", reflect.TypeOf(func(jsoniter.Stream) {}).In(0), ""),
		register.NewType("StreamPool", reflect.TypeOf(func(jsoniter.StreamPool) {}).In(0), ""),
		register.NewConst("StringValue", jsoniter.StringValue, ""),
		register.NewType("StructDescriptor", reflect.TypeOf(func(jsoniter.StructDescriptor) {}).In(0), ""),
		register.NewFunction("Unmarshal", jsoniter.Unmarshal, ""),
		register.NewFunction("UnmarshalFromString", jsoniter.UnmarshalFromString, ""),
		register.NewType("ValDecoder", reflect.TypeOf(func(jsoniter.ValDecoder) {}).In(0), ""),
		register.NewType("ValEncoder", reflect.TypeOf(func(jsoniter.ValEncoder) {}).In(0), ""),
		register.NewFunction("Valid", jsoniter.Valid, ""),
		register.NewType("ValueType", reflect.TypeOf(func(jsoniter.ValueType) {}).In(0), ""),
		register.NewFunction("Wrap", jsoniter.Wrap, ""),
		register.NewFunction("WrapFloat64", jsoniter.WrapFloat64, ""),
		register.NewFunction("WrapInt32", jsoniter.WrapInt32, ""),
		register.NewFunction("WrapInt64", jsoniter.WrapInt64, ""),
		register.NewFunction("WrapString", jsoniter.WrapString, ""),
		register.NewFunction("WrapUint32", jsoniter.WrapUint32, ""),
		register.NewFunction("WrapUint64", jsoniter.WrapUint64, ""),
	)
}
