// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 02 Jul 2025 23:31:51 UTC.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ffms_bindings

/*
#cgo pkg-config: ffms2
#include "ffms.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// H as defined in include/ffms.h:22

	// VERSION as defined in include/ffms.h:25
	VERSION = ((5 << 24) | (0 << 16) | (0 << 8) | 0)
	// EXTERNC as defined in include/ffms.h:42

	// CC as defined in include/ffms.h:49

	// DEPRECATED as defined in include/ffms.h:55
	DEPRECATED = 0x6468e0
)

// Errors as declared in include/ffms.h:132
type Errors int32

// Errors enumeration from include/ffms.h:132
const (
	ERRORSUCCESS          Errors = iota
	ERRORINDEX            Errors = 1
	ERRORINDEXING         Errors = 2
	ERRORPOSTPROCESSING   Errors = 3
	ERRORSCALING          Errors = 4
	ERRORDECODING         Errors = 5
	ERRORSEEKING          Errors = 6
	ERRORPARSER           Errors = 7
	ERRORTRACK            Errors = 8
	ERRORWAVEWRITER       Errors = 9
	ERRORCANCELLED        Errors = 10
	ERRORRESAMPLING       Errors = 11
	ERRORUNKNOWN          Errors = 20
	ERRORUNSUPPORTED      Errors = 21
	ERRORFILEREAD         Errors = 22
	ERRORFILEWRITE        Errors = 23
	ERRORNOFILE           Errors = 24
	ERRORVERSION          Errors = 25
	ERRORALLOCATIONFAILED Errors = 26
	ERRORINVALIDARGUMENT  Errors = 27
	ERRORCODEC            Errors = 28
	ERRORNOTAVAILABLE     Errors = 29
	ERRORFILEMISMATCH     Errors = 30
	ERRORUSER             Errors = 31
)

// SeekMode as declared in include/ffms.h:140
type SeekMode int32

// SeekMode enumeration from include/ffms.h:140
const (
	SEEKLINEARNORW SeekMode = -1
	SEEKLINEAR     SeekMode = 0
	SEEKNORMAL     SeekMode = 1
	SEEKUNSAFE     SeekMode = 2
	SEEKAGGRESSIVE SeekMode = 3
)

// IndexErrorHandling as declared in include/ffms.h:147
type IndexErrorHandling int32

// IndexErrorHandling enumeration from include/ffms.h:147
const (
	IEHABORT      IndexErrorHandling = iota
	IEHCLEARTRACK IndexErrorHandling = 1
	IEHSTOPTRACK  IndexErrorHandling = 2
	IEHIGNORE     IndexErrorHandling = 3
)

// TrackType as declared in include/ffms.h:156
type TrackType int32

// TrackType enumeration from include/ffms.h:156
const (
	TYPEUNKNOWN    TrackType = -1
	TYPEVIDEO      TrackType = 0
	TYPEAUDIO      TrackType = 1
	TYPEDATA       TrackType = 2
	TYPESUBTITLE   TrackType = 3
	TYPEATTACHMENT TrackType = 4
)

// SampleFormat as declared in include/ffms.h:164
type SampleFormat int32

// SampleFormat enumeration from include/ffms.h:164
const (
	FMTU8  SampleFormat = iota
	FMTS16 SampleFormat = 1
	FMTS32 SampleFormat = 2
	FMTFLT SampleFormat = 3
	FMTDBL SampleFormat = 4
)

// AudioChannel as declared in include/ffms.h:187
type AudioChannel int32

// AudioChannel enumeration from include/ffms.h:187
const (
	CHFRONTLEFT          AudioChannel = 1
	CHFRONTRIGHT         AudioChannel = 2
	CHFRONTCENTER        AudioChannel = 4
	CHLOWFREQUENCY       AudioChannel = 8
	CHBACKLEFT           AudioChannel = 16
	CHBACKRIGHT          AudioChannel = 32
	CHFRONTLEFTOFCENTER  AudioChannel = 64
	CHFRONTRIGHTOFCENTER AudioChannel = 128
	CHBACKCENTER         AudioChannel = 256
	CHSIDELEFT           AudioChannel = 512
	CHSIDERIGHT          AudioChannel = 1024
	CHTOPCENTER          AudioChannel = 2048
	CHTOPFRONTLEFT       AudioChannel = 4096
	CHTOPFRONTCENTER     AudioChannel = 8192
	CHTOPFRONTRIGHT      AudioChannel = 16384
	CHTOPBACKLEFT        AudioChannel = 32768
	CHTOPBACKCENTER      AudioChannel = 65536
	CHTOPBACKRIGHT       AudioChannel = 131072
	CHSTEREOLEFT         AudioChannel = 536870912
	CHSTEREORIGHT        AudioChannel = 1073741824
)

// Resizers as declared in include/ffms.h:201
type Resizers int32

// Resizers enumeration from include/ffms.h:201
const (
	RESIZERFASTBILINEAR Resizers = 1
	RESIZERBILINEAR     Resizers = 2
	RESIZERBICUBIC      Resizers = 4
	RESIZERX            Resizers = 8
	RESIZERPOINT        Resizers = 16
	RESIZERAREA         Resizers = 32
	RESIZERBICUBLIN     Resizers = 64
	RESIZERGAUSS        Resizers = 128
	RESIZERSINC         Resizers = 256
	RESIZERLANCZOS      Resizers = 512
	RESIZERSPLINE       Resizers = 1024
)

// AudioDelayModes as declared in include/ffms.h:207
type AudioDelayModes int32

// AudioDelayModes enumeration from include/ffms.h:207
const (
	DELAYNOSHIFT         AudioDelayModes = -3
	DELAYTIMEZERO        AudioDelayModes = -2
	DELAYFIRSTVIDEOTRACK AudioDelayModes = -1
)

// AudioGapFillModes as declared in include/ffms.h:213
type AudioGapFillModes int32

// AudioGapFillModes enumeration from include/ffms.h:213
const (
	GAPFILLAUTO     AudioGapFillModes = -1
	GAPFILLDISABLED AudioGapFillModes = 0
	GAPFILLENABLED  AudioGapFillModes = -1
)

// ChromaLocations as declared in include/ffms.h:223
type ChromaLocations int32

// ChromaLocations enumeration from include/ffms.h:223
const (
	LOCUNSPECIFIED ChromaLocations = iota
	LOCLEFT        ChromaLocations = 1
	LOCCENTER      ChromaLocations = 2
	LOCTOPLEFT     ChromaLocations = 3
	LOCTOP         ChromaLocations = 4
	LOCBOTTOMLEFT  ChromaLocations = 5
	LOCBOTTOM      ChromaLocations = 6
)

// ColorRanges as declared in include/ffms.h:229
type ColorRanges int32

// ColorRanges enumeration from include/ffms.h:229
const (
	CRUNSPECIFIED ColorRanges = iota
	CRMPEG        ColorRanges = 1
	CRJPEG        ColorRanges = 2
)

// Stereo3DType as declared in include/ffms.h:240
type Stereo3DType int32

// Stereo3DType enumeration from include/ffms.h:240
const (
	S3DTYPE2D                 Stereo3DType = iota
	S3DTYPESIDEBYSIDE         Stereo3DType = 1
	S3DTYPETOPBOTTOM          Stereo3DType = 2
	S3DTYPEFRAMESEQUENCE      Stereo3DType = 3
	S3DTYPECHECKERBOARD       Stereo3DType = 4
	S3DTYPESIDEBYSIDEQUINCUNX Stereo3DType = 5
	S3DTYPELINES              Stereo3DType = 6
	S3DTYPECOLUMNS            Stereo3DType = 7
)

// Stereo3DFlags as declared in include/ffms.h:244
type Stereo3DFlags int32

// Stereo3DFlags enumeration from include/ffms.h:244
const (
	S3DFLAGSINVERT Stereo3DFlags = 1
)

// MixingCoefficientType as declared in include/ffms.h:250
type MixingCoefficientType int32

// MixingCoefficientType enumeration from include/ffms.h:250
const (
	MIXINGCOEFFICIENTQ8  MixingCoefficientType = iota
	MIXINGCOEFFICIENTQ15 MixingCoefficientType = 1
	MIXINGCOEFFICIENTFLT MixingCoefficientType = 2
)

// MatrixEncoding as declared in include/ffms.h:260
type MatrixEncoding int32

// MatrixEncoding enumeration from include/ffms.h:260
const (
	MATRIXENCODINGNONE           MatrixEncoding = iota
	MATRIXENCODINGDOBLY          MatrixEncoding = 1
	MATRIXENCODINGPROLOGICII     MatrixEncoding = 2
	MATRIXENCODINGPROLOGICIIX    MatrixEncoding = 3
	MATRIXENCODINGPROLOGICIIZ    MatrixEncoding = 4
	MATRIXENCODINGDOLBYEX        MatrixEncoding = 5
	MATRIXENCODINGDOLBYHEADPHONE MatrixEncoding = 6
)

// ResampleFilterType as declared in include/ffms.h:266
type ResampleFilterType int32

// ResampleFilterType enumeration from include/ffms.h:266
const (
	RESAMPLEFILTERCUBIC  ResampleFilterType = iota
	RESAMPLEFILTERSINC   ResampleFilterType = 1
	RESAMPLEFILTERKAISER ResampleFilterType = 2
)

// AudioDitherMethod as declared in include/ffms.h:274
type AudioDitherMethod int32

// AudioDitherMethod enumeration from include/ffms.h:274
const (
	RESAMPLEDITHERNONE                   AudioDitherMethod = iota
	RESAMPLEDITHERRECTANGULAR            AudioDitherMethod = 1
	RESAMPLEDITHERTRIANGULAR             AudioDitherMethod = 2
	RESAMPLEDITHERTRIANGULARHIGHPASS     AudioDitherMethod = 3
	RESAMPLEDITHERTRIANGULARNOISESHAPING AudioDitherMethod = 4
)

// LogLevels as declared in include/ffms.h:286
type LogLevels int32

// LogLevels enumeration from include/ffms.h:286
const (
	LOGQUIET   LogLevels = -8
	LOGPANIC   LogLevels = 0
	LOGFATAL   LogLevels = 8
	LOGERROR   LogLevels = 16
	LOGWARNING LogLevels = 24
	LOGINFO    LogLevels = 32
	LOGVERBOSE LogLevels = 40
	LOGDEBUG   LogLevels = 48
	LOGTRACE   LogLevels = 56
)
