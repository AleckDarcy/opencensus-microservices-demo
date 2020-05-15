# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: message.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='message.proto',
  package='tracer',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\rmessage.proto\x12\x06tracer\"r\n\x06Record\x12 \n\x04type\x18\x01 \x01(\x0e\x32\x12.tracer.RecordType\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12\x14\n\x0cmessage_name\x18\x03 \x01(\t\x12\x0c\n\x04uuid\x18\x04 \x01(\t\x12\x0f\n\x07service\x18\x05 \x01(\t\"l\n\x05Trace\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1f\n\x07records\x18\x02 \x03(\x0b\x32\x0e.tracer.Record\x12\x1b\n\x05rlfis\x18\x14 \x03(\x0b\x32\x0c.tracer.RLFI\x12\x19\n\x04tfis\x18\x15 \x03(\x0b\x32\x0b.tracer.TFI\"D\n\x04RLFI\x12\x1f\n\x04type\x18\x01 \x01(\x0e\x32\x11.tracer.FaultType\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x64\x65lay\x18\x03 \x01(\x03\"7\n\x07TFIMeta\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05times\x18\x02 \x01(\x03\x12\x0f\n\x07\x61lready\x18\x03 \x01(\x03\"c\n\x03TFI\x12\x1f\n\x04type\x18\x01 \x01(\x0e\x32\x11.tracer.FaultType\x12\x0c\n\x04name\x18\x02 \x03(\t\x12\r\n\x05\x64\x65lay\x18\x03 \x01(\x03\x12\x1e\n\x05\x61\x66ter\x18\x04 \x03(\x0b\x32\x0f.tracer.TFIMeta*F\n\x0bMessageType\x12\x0c\n\x08Message_\x10\x00\x12\x13\n\x0fMessage_Request\x10\x01\x12\x14\n\x10Message_Response\x10\x02*<\n\nRecordType\x12\x0b\n\x07Record_\x10\x00\x12\x0e\n\nRecordSend\x10\x01\x12\x11\n\rRecordReceive\x10\x02*7\n\tFaultType\x12\n\n\x06\x46\x61ult_\x10\x00\x12\x0e\n\nFaultCrash\x10\x01\x12\x0e\n\nFaultDelay\x10\x02\x62\x06proto3'
)

_MESSAGETYPE = _descriptor.EnumDescriptor(
  name='MessageType',
  full_name='tracer.MessageType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Message_', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Message_Request', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Message_Response', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=479,
  serialized_end=549,
)
_sym_db.RegisterEnumDescriptor(_MESSAGETYPE)

MessageType = enum_type_wrapper.EnumTypeWrapper(_MESSAGETYPE)
_RECORDTYPE = _descriptor.EnumDescriptor(
  name='RecordType',
  full_name='tracer.RecordType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Record_', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RecordSend', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RecordReceive', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=551,
  serialized_end=611,
)
_sym_db.RegisterEnumDescriptor(_RECORDTYPE)

RecordType = enum_type_wrapper.EnumTypeWrapper(_RECORDTYPE)
_FAULTTYPE = _descriptor.EnumDescriptor(
  name='FaultType',
  full_name='tracer.FaultType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Fault_', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FaultCrash', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FaultDelay', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=613,
  serialized_end=668,
)
_sym_db.RegisterEnumDescriptor(_FAULTTYPE)

FaultType = enum_type_wrapper.EnumTypeWrapper(_FAULTTYPE)
Message_ = 0
Message_Request = 1
Message_Response = 2
Record_ = 0
RecordSend = 1
RecordReceive = 2
Fault_ = 0
FaultCrash = 1
FaultDelay = 2



_RECORD = _descriptor.Descriptor(
  name='Record',
  full_name='tracer.Record',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='tracer.Record.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='tracer.Record.timestamp', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message_name', full_name='tracer.Record.message_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='tracer.Record.uuid', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='service', full_name='tracer.Record.service', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=139,
)


_TRACE = _descriptor.Descriptor(
  name='Trace',
  full_name='tracer.Trace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='tracer.Trace.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='records', full_name='tracer.Trace.records', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rlfis', full_name='tracer.Trace.rlfis', index=2,
      number=20, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tfis', full_name='tracer.Trace.tfis', index=3,
      number=21, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=141,
  serialized_end=249,
)


_RLFI = _descriptor.Descriptor(
  name='RLFI',
  full_name='tracer.RLFI',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='tracer.RLFI.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='tracer.RLFI.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='delay', full_name='tracer.RLFI.delay', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=251,
  serialized_end=319,
)


_TFIMETA = _descriptor.Descriptor(
  name='TFIMeta',
  full_name='tracer.TFIMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='tracer.TFIMeta.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='times', full_name='tracer.TFIMeta.times', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='already', full_name='tracer.TFIMeta.already', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=321,
  serialized_end=376,
)


_TFI = _descriptor.Descriptor(
  name='TFI',
  full_name='tracer.TFI',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='tracer.TFI.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='tracer.TFI.name', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='delay', full_name='tracer.TFI.delay', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='after', full_name='tracer.TFI.after', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=378,
  serialized_end=477,
)

_RECORD.fields_by_name['type'].enum_type = _RECORDTYPE
_TRACE.fields_by_name['records'].message_type = _RECORD
_TRACE.fields_by_name['rlfis'].message_type = _RLFI
_TRACE.fields_by_name['tfis'].message_type = _TFI
_RLFI.fields_by_name['type'].enum_type = _FAULTTYPE
_TFI.fields_by_name['type'].enum_type = _FAULTTYPE
_TFI.fields_by_name['after'].message_type = _TFIMETA
DESCRIPTOR.message_types_by_name['Record'] = _RECORD
DESCRIPTOR.message_types_by_name['Trace'] = _TRACE
DESCRIPTOR.message_types_by_name['RLFI'] = _RLFI
DESCRIPTOR.message_types_by_name['TFIMeta'] = _TFIMETA
DESCRIPTOR.message_types_by_name['TFI'] = _TFI
DESCRIPTOR.enum_types_by_name['MessageType'] = _MESSAGETYPE
DESCRIPTOR.enum_types_by_name['RecordType'] = _RECORDTYPE
DESCRIPTOR.enum_types_by_name['FaultType'] = _FAULTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Record = _reflection.GeneratedProtocolMessageType('Record', (_message.Message,), {
  'DESCRIPTOR' : _RECORD,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:tracer.Record)
  })
_sym_db.RegisterMessage(Record)

Trace = _reflection.GeneratedProtocolMessageType('Trace', (_message.Message,), {
  'DESCRIPTOR' : _TRACE,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:tracer.Trace)
  })
_sym_db.RegisterMessage(Trace)

RLFI = _reflection.GeneratedProtocolMessageType('RLFI', (_message.Message,), {
  'DESCRIPTOR' : _RLFI,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:tracer.RLFI)
  })
_sym_db.RegisterMessage(RLFI)

TFIMeta = _reflection.GeneratedProtocolMessageType('TFIMeta', (_message.Message,), {
  'DESCRIPTOR' : _TFIMETA,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:tracer.TFIMeta)
  })
_sym_db.RegisterMessage(TFIMeta)

TFI = _reflection.GeneratedProtocolMessageType('TFI', (_message.Message,), {
  'DESCRIPTOR' : _TFI,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:tracer.TFI)
  })
_sym_db.RegisterMessage(TFI)


# @@protoc_insertion_point(module_scope)
