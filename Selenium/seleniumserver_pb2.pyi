from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class HtmlArgs(_message.Message):
    __slots__ = ["Url"]
    URL_FIELD_NUMBER: _ClassVar[int]
    Url: str
    def __init__(self, Url: _Optional[str] = ...) -> None: ...
