// Code generated - DO NOT EDIT.

package glib

// Asciitype is a representation of the C type AsciiType.
type Asciitype int

const (
	Asciitype_Alnum  Asciitype = 1
	Asciitype_Alpha  Asciitype = 2
	Asciitype_Cntrl  Asciitype = 4
	Asciitype_Digit  Asciitype = 8
	Asciitype_Graph  Asciitype = 16
	Asciitype_Lower  Asciitype = 32
	Asciitype_Print  Asciitype = 64
	Asciitype_Punct  Asciitype = 128
	Asciitype_Space  Asciitype = 256
	Asciitype_Upper  Asciitype = 512
	Asciitype_Xdigit Asciitype = 1024
)

// Filetest is a representation of the C type FileTest.
type Filetest int

const (
	Filetest_IsRegular    Filetest = 1
	Filetest_IsSymlink    Filetest = 2
	Filetest_IsDir        Filetest = 4
	Filetest_IsExecutable Filetest = 8
	Filetest_Exists       Filetest = 16
)

// Formatsizeflags is a representation of the C type FormatSizeFlags.
type Formatsizeflags int

const (
	Formatsizeflags_Default    Formatsizeflags = 0
	Formatsizeflags_LongFormat Formatsizeflags = 1
	Formatsizeflags_IecUnits   Formatsizeflags = 2
	Formatsizeflags_Bits       Formatsizeflags = 4
)

// Hookflagmask is a representation of the C type HookFlagMask.
type Hookflagmask int

const (
	Hookflagmask_Active Hookflagmask = 1
	Hookflagmask_InCall Hookflagmask = 2
	Hookflagmask_Mask   Hookflagmask = 15
)

// Iocondition is a representation of the C type IOCondition.
type Iocondition int

const (
	Iocondition_In   Iocondition = 1
	Iocondition_Out  Iocondition = 4
	Iocondition_Pri  Iocondition = 2
	Iocondition_Err  Iocondition = 8
	Iocondition_Hup  Iocondition = 16
	Iocondition_Nval Iocondition = 32
)

// Ioflags is a representation of the C type IOFlags.
type Ioflags int

const (
	Ioflags_Append      Ioflags = 1
	Ioflags_Nonblock    Ioflags = 2
	Ioflags_IsReadable  Ioflags = 4
	Ioflags_IsWritable  Ioflags = 8
	Ioflags_IsWriteable Ioflags = 8
	Ioflags_IsSeekable  Ioflags = 16
	Ioflags_Mask        Ioflags = 31
	Ioflags_GetMask     Ioflags = 31
	Ioflags_SetMask     Ioflags = 3
)

// Keyfileflags is a representation of the C type KeyFileFlags.
type Keyfileflags int

const (
	Keyfileflags_None             Keyfileflags = 0
	Keyfileflags_KeepComments     Keyfileflags = 1
	Keyfileflags_KeepTranslations Keyfileflags = 2
)

// Loglevelflags is a representation of the C type LogLevelFlags.
type Loglevelflags int

const (
	Loglevelflags_FlagRecursion Loglevelflags = 1
	Loglevelflags_FlagFatal     Loglevelflags = 2
	Loglevelflags_LevelError    Loglevelflags = 4
	Loglevelflags_LevelCritical Loglevelflags = 8
	Loglevelflags_LevelWarning  Loglevelflags = 16
	Loglevelflags_LevelMessage  Loglevelflags = 32
	Loglevelflags_LevelInfo     Loglevelflags = 64
	Loglevelflags_LevelDebug    Loglevelflags = 128
	Loglevelflags_LevelMask     Loglevelflags = -4
)

// Markupcollecttype is a representation of the C type MarkupCollectType.
type Markupcollecttype int

const (
	Markupcollecttype_Invalid  Markupcollecttype = 0
	Markupcollecttype_String   Markupcollecttype = 1
	Markupcollecttype_Strdup   Markupcollecttype = 2
	Markupcollecttype_Boolean  Markupcollecttype = 3
	Markupcollecttype_Tristate Markupcollecttype = 4
	Markupcollecttype_Optional Markupcollecttype = 65536
)

// Markupparseflags is a representation of the C type MarkupParseFlags.
type Markupparseflags int

const (
	Markupparseflags_DoNotUseThisUnsupportedFlag Markupparseflags = 1
	Markupparseflags_TreatCdataAsText            Markupparseflags = 2
	Markupparseflags_PrefixErrorPosition         Markupparseflags = 4
	Markupparseflags_IgnoreQualified             Markupparseflags = 8
)

// Optionflags is a representation of the C type OptionFlags.
type Optionflags int

const (
	Optionflags_None        Optionflags = 0
	Optionflags_Hidden      Optionflags = 1
	Optionflags_InMain      Optionflags = 2
	Optionflags_Reverse     Optionflags = 4
	Optionflags_NoArg       Optionflags = 8
	Optionflags_Filename    Optionflags = 16
	Optionflags_OptionalArg Optionflags = 32
	Optionflags_Noalias     Optionflags = 64
)

// Regexcompileflags is a representation of the C type RegexCompileFlags.
//
// since 2.14
type Regexcompileflags int

const (
	Regexcompileflags_Caseless         Regexcompileflags = 1
	Regexcompileflags_Multiline        Regexcompileflags = 2
	Regexcompileflags_Dotall           Regexcompileflags = 4
	Regexcompileflags_Extended         Regexcompileflags = 8
	Regexcompileflags_Anchored         Regexcompileflags = 16
	Regexcompileflags_DollarEndonly    Regexcompileflags = 32
	Regexcompileflags_Ungreedy         Regexcompileflags = 512
	Regexcompileflags_Raw              Regexcompileflags = 2048
	Regexcompileflags_NoAutoCapture    Regexcompileflags = 4096
	Regexcompileflags_Optimize         Regexcompileflags = 8192
	Regexcompileflags_Firstline        Regexcompileflags = 262144
	Regexcompileflags_Dupnames         Regexcompileflags = 524288
	Regexcompileflags_NewlineCr        Regexcompileflags = 1048576
	Regexcompileflags_NewlineLf        Regexcompileflags = 2097152
	Regexcompileflags_NewlineCrlf      Regexcompileflags = 3145728
	Regexcompileflags_NewlineAnycrlf   Regexcompileflags = 5242880
	Regexcompileflags_BsrAnycrlf       Regexcompileflags = 8388608
	Regexcompileflags_JavascriptCompat Regexcompileflags = 33554432
)

// Regexmatchflags is a representation of the C type RegexMatchFlags.
//
// since 2.14
type Regexmatchflags int

const (
	Regexmatchflags_Anchored        Regexmatchflags = 16
	Regexmatchflags_Notbol          Regexmatchflags = 128
	Regexmatchflags_Noteol          Regexmatchflags = 256
	Regexmatchflags_Notempty        Regexmatchflags = 1024
	Regexmatchflags_Partial         Regexmatchflags = 32768
	Regexmatchflags_NewlineCr       Regexmatchflags = 1048576
	Regexmatchflags_NewlineLf       Regexmatchflags = 2097152
	Regexmatchflags_NewlineCrlf     Regexmatchflags = 3145728
	Regexmatchflags_NewlineAny      Regexmatchflags = 4194304
	Regexmatchflags_NewlineAnycrlf  Regexmatchflags = 5242880
	Regexmatchflags_BsrAnycrlf      Regexmatchflags = 8388608
	Regexmatchflags_BsrAny          Regexmatchflags = 16777216
	Regexmatchflags_PartialSoft     Regexmatchflags = 32768
	Regexmatchflags_PartialHard     Regexmatchflags = 134217728
	Regexmatchflags_NotemptyAtstart Regexmatchflags = 268435456
)

// Spawnflags is a representation of the C type SpawnFlags.
type Spawnflags int

const (
	Spawnflags_Default              Spawnflags = 0
	Spawnflags_LeaveDescriptorsOpen Spawnflags = 1
	Spawnflags_DoNotReapChild       Spawnflags = 2
	Spawnflags_SearchPath           Spawnflags = 4
	Spawnflags_StdoutToDevNull      Spawnflags = 8
	Spawnflags_StderrToDevNull      Spawnflags = 16
	Spawnflags_ChildInheritsStdin   Spawnflags = 32
	Spawnflags_FileAndArgvZero      Spawnflags = 64
	Spawnflags_SearchPathFromEnvp   Spawnflags = 128
	Spawnflags_CloexecPipes         Spawnflags = 256
)

// Testsubprocessflags is a representation of the C type TestSubprocessFlags.
type Testsubprocessflags int

const (
	Testsubprocessflags_Stdin  Testsubprocessflags = 1
	Testsubprocessflags_Stdout Testsubprocessflags = 2
	Testsubprocessflags_Stderr Testsubprocessflags = 4
)

// Testtrapflags is a representation of the C type TestTrapFlags.
type Testtrapflags int

const (
	Testtrapflags_SilenceStdout Testtrapflags = 128
	Testtrapflags_SilenceStderr Testtrapflags = 256
	Testtrapflags_InheritStdin  Testtrapflags = 512
)

// Traverseflags is a representation of the C type TraverseFlags.
type Traverseflags int

const (
	Traverseflags_Leaves    Traverseflags = 1
	Traverseflags_NonLeaves Traverseflags = 2
	Traverseflags_All       Traverseflags = 3
	Traverseflags_Mask      Traverseflags = 3
	Traverseflags_Leafs     Traverseflags = 1
	Traverseflags_NonLeafs  Traverseflags = 2
)
