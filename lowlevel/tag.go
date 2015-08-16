package lowlevel

/*
#include <fmod_common.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type TagType C.FMOD_TAGTYPE

const (
	TAGTYPE_UNKNOWN       TagType = C.FMOD_TAGTYPE_UNKNOWN
	TAGTYPE_ID3V1                 = C.FMOD_TAGTYPE_ID3V1
	TAGTYPE_ID3V2                 = C.FMOD_TAGTYPE_ID3V2
	TAGTYPE_VORBISCOMMENT         = C.FMOD_TAGTYPE_VORBISCOMMENT
	TAGTYPE_SHOUTCAST             = C.FMOD_TAGTYPE_SHOUTCAST
	TAGTYPE_ICECAST               = C.FMOD_TAGTYPE_ICECAST
	TAGTYPE_ASF                   = C.FMOD_TAGTYPE_ASF
	TAGTYPE_MIDI                  = C.FMOD_TAGTYPE_MIDI
	TAGTYPE_PLAYLIST              = C.FMOD_TAGTYPE_PLAYLIST
	TAGTYPE_FMOD                  = C.FMOD_TAGTYPE_FMOD
	TAGTYPE_USER                  = C.FMOD_TAGTYPE_USER
	TAGTYPE_MAX                   = C.FMOD_TAGTYPE_MAX
	TAGTYPE_FORCEINT              = C.FMOD_TAGTYPE_FORCEINT
)

type TagDataType C.FMOD_TAGDATATYPE

const (
	TAGDATATYPE_BINARY         TagDataType = C.FMOD_TAGDATATYPE_BINARY
	TAGDATATYPE_INT                        = C.FMOD_TAGDATATYPE_INT
	TAGDATATYPE_FLOAT                      = C.FMOD_TAGDATATYPE_FLOAT
	TAGDATATYPE_STRING                     = C.FMOD_TAGDATATYPE_STRING
	TAGDATATYPE_STRING_UTF16               = C.FMOD_TAGDATATYPE_STRING_UTF16
	TAGDATATYPE_STRING_UTF16BE             = C.FMOD_TAGDATATYPE_STRING_UTF16BE
	TAGDATATYPE_STRING_UTF8                = C.FMOD_TAGDATATYPE_STRING_UTF8
	TAGDATATYPE_CDTOC                      = C.FMOD_TAGDATATYPE_CDTOC
	TAGDATATYPE_MAX                        = C.FMOD_TAGDATATYPE_MAX
	TAGDATATYPE_FORCEINT                   = C.FMOD_TAGDATATYPE_FORCEINT
)

type Tag struct {
	Type     TagType
	DataType TagDataType
	Name     string
	Data     unsafe.Pointer
	DataLen  uint32
	Updated  bool
}

func NewTag() Tag {
	return Tag{}
}

func (t *Tag) fromC(ct C.FMOD_TAG) {
	t.Type = TagType(ct._type)
	t.DataType = TagDataType(ct.datatype)
	t.Name = C.GoString(ct.name)
	t.Data = unsafe.Pointer(ct.data)
	t.DataLen = uint32(ct.datalen)
	t.Updated = setBool(ct.updated)
}

func (t *Tag) toC() C.FMOD_TAG {
	var ctag C.FMOD_TAG
	ctag._type = C.FMOD_TAGTYPE(t.Type)
	ctag.datatype = C.FMOD_TAGDATATYPE(t.DataType)
	name := C.CString(t.Name)
	C.free(unsafe.Pointer(name))
	ctag.name = name
	ctag.data = t.Data
	ctag.datalen = C.uint(t.DataLen)
	ctag.updated = getBool(t.Updated)
	return ctag
}
