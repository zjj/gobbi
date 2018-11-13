// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Error codes returned by regular expressions functions.
type RegexError C.GRegexError

const (
	// Compilation of the regular expression failed.
	REGEX_ERROR_COMPILE RegexError = 0

	// Optimization of the regular expression failed.
	REGEX_ERROR_OPTIMIZE RegexError = 1

	// Replacement failed due to an ill-formed replacement
	// string.
	REGEX_ERROR_REPLACE RegexError = 2

	// The match process failed.
	REGEX_ERROR_MATCH RegexError = 3

	// Internal error of the regular expression engine.
	// Since 2.16
	REGEX_ERROR_INTERNAL RegexError = 4

	// "\\" at end of pattern. Since 2.16
	REGEX_ERROR_STRAY_BACKSLASH RegexError = 101

	// "\\c" at end of pattern. Since 2.16
	REGEX_ERROR_MISSING_CONTROL_CHAR RegexError = 102

	// Unrecognized character follows "\\".
	// Since 2.16
	REGEX_ERROR_UNRECOGNIZED_ESCAPE RegexError = 103

	// Numbers out of order in "{}"
	// quantifier. Since 2.16
	REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER RegexError = 104

	// Number too big in "{}" quantifier.
	// Since 2.16
	REGEX_ERROR_QUANTIFIER_TOO_BIG RegexError = 105

	// Missing terminating "]" for
	// character class. Since 2.16
	REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS RegexError = 106

	// Invalid escape sequence
	// in character class. Since 2.16
	REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS RegexError = 107

	// Range out of order in character class.
	// Since 2.16
	REGEX_ERROR_RANGE_OUT_OF_ORDER RegexError = 108

	// Nothing to repeat. Since 2.16
	REGEX_ERROR_NOTHING_TO_REPEAT RegexError = 109

	// Unrecognized character after "(?",
	// "(?<" or "(?P". Since 2.16
	REGEX_ERROR_UNRECOGNIZED_CHARACTER RegexError = 112

	// POSIX named classes are
	// supported only within a class. Since 2.16
	REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS RegexError = 113

	// Missing terminating ")" or ")"
	// without opening "(". Since 2.16
	REGEX_ERROR_UNMATCHED_PARENTHESIS RegexError = 114

	// Reference to non-existent
	// subpattern. Since 2.16
	REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE RegexError = 115

	// Missing terminating ")" after comment.
	// Since 2.16
	REGEX_ERROR_UNTERMINATED_COMMENT RegexError = 118

	// Regular expression too large.
	// Since 2.16
	REGEX_ERROR_EXPRESSION_TOO_LARGE RegexError = 120

	// Failed to get memory. Since 2.16
	REGEX_ERROR_MEMORY_ERROR RegexError = 121

	// Lookbehind assertion is not
	// fixed length. Since 2.16
	REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND RegexError = 125

	// Malformed number or name after "(?(".
	// Since 2.16
	REGEX_ERROR_MALFORMED_CONDITION RegexError = 126

	// Conditional group contains
	// more than two branches. Since 2.16
	REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES RegexError = 127

	// Assertion expected after "(?(".
	// Since 2.16
	REGEX_ERROR_ASSERTION_EXPECTED RegexError = 128

	// Unknown POSIX class name.
	// Since 2.16
	REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME RegexError = 130

	// POSIX collating
	// elements are not supported. Since 2.16
	REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED RegexError = 131

	// Character value in "\\x{...}" sequence
	// is too large. Since 2.16
	REGEX_ERROR_HEX_CODE_TOO_LARGE RegexError = 134

	// Invalid condition "(?(0)". Since 2.16
	REGEX_ERROR_INVALID_CONDITION RegexError = 135

	// \\C not allowed in
	// lookbehind assertion. Since 2.16
	REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND RegexError = 136

	// Recursive call could loop indefinitely.
	// Since 2.16
	REGEX_ERROR_INFINITE_LOOP RegexError = 140

	// Missing terminator
	// in subpattern name. Since 2.16
	REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR RegexError = 142

	// Two named subpatterns have
	// the same name. Since 2.16
	REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME RegexError = 143

	// Malformed "\\P" or "\\p" sequence.
	// Since 2.16
	REGEX_ERROR_MALFORMED_PROPERTY RegexError = 146

	// Unknown property name after "\\P" or
	// "\\p". Since 2.16
	REGEX_ERROR_UNKNOWN_PROPERTY RegexError = 147

	// Subpattern name is too long
	// (maximum 32 characters). Since 2.16
	REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG RegexError = 148

	// Too many named subpatterns (maximum
	// 10,000). Since 2.16
	REGEX_ERROR_TOO_MANY_SUBPATTERNS RegexError = 149

	// Octal value is greater than "\\377".
	// Since 2.16
	REGEX_ERROR_INVALID_OCTAL_VALUE RegexError = 151

	// "DEFINE" group contains more
	// than one branch. Since 2.16
	REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE RegexError = 154

	// Repeating a "DEFINE" group is not allowed.
	// This error is never raised. Since: 2.16 Deprecated: 2.34
	REGEX_ERROR_DEFINE_REPETION RegexError = 155

	// Inconsistent newline options.
	// Since 2.16
	REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS RegexError = 156

	// "\\g" is not followed by a braced,
	// angle-bracketed, or quoted name or number, or by a plain number. Since: 2.16
	REGEX_ERROR_MISSING_BACK_REFERENCE RegexError = 157

	// relative reference must not be zero. Since: 2.34
	REGEX_ERROR_INVALID_RELATIVE_REFERENCE RegexError = 158

	// the backtracing
	// control verb used does not allow an argument. Since: 2.34
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN RegexError = 159

	// unknown backtracing
	// control verb. Since: 2.34
	REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB RegexError = 160

	// number is too big in escape sequence. Since: 2.34
	REGEX_ERROR_NUMBER_TOO_BIG RegexError = 161

	// Missing subpattern name. Since: 2.34
	REGEX_ERROR_MISSING_SUBPATTERN_NAME RegexError = 162

	// Missing digit. Since 2.34
	REGEX_ERROR_MISSING_DIGIT RegexError = 163

	// In JavaScript compatibility mode,
	// "[" is an invalid data character. Since: 2.34
	REGEX_ERROR_INVALID_DATA_CHARACTER RegexError = 164

	// different names for subpatterns of the
	// same number are not allowed. Since: 2.34
	REGEX_ERROR_EXTRA_SUBPATTERN_NAME RegexError = 165

	// the backtracing control
	// verb requires an argument. Since: 2.34
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED RegexError = 166

	// "\\c" must be followed by an ASCII
	// character. Since: 2.34
	REGEX_ERROR_INVALID_CONTROL_CHAR RegexError = 168

	// "\\k" is not followed by a braced, angle-bracketed, or
	// quoted name. Since: 2.34
	REGEX_ERROR_MISSING_NAME RegexError = 169

	// "\\N" is not supported in a class. Since: 2.34
	REGEX_ERROR_NOT_SUPPORTED_IN_CLASS RegexError = 171

	// too many forward references. Since: 2.34
	REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES RegexError = 172

	// the name is too long in "(*MARK)", "(*PRUNE)",
	// "(*SKIP)", or "(*THEN)". Since: 2.34
	REGEX_ERROR_NAME_TOO_LONG RegexError = 175

	// the character value in the \\u sequence is
	// too large. Since: 2.34
	REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE RegexError = 176
)

// These are logical ids for special directories which are defined
// depending on the platform used. You should use g_get_user_special_dir()
// to retrieve the full path associated to the logical id.
//
// The #GUserDirectory enumeration can be extended at later date. Not
// every platform has a directory for every logical id in this
// enumeration.
type UserDirectory C.GUserDirectory

const (
	// the user's Desktop directory
	USER_DIRECTORY_DESKTOP UserDirectory = 0

	// the user's Documents directory
	USER_DIRECTORY_DOCUMENTS UserDirectory = 1

	// the user's Downloads directory
	USER_DIRECTORY_DOWNLOAD UserDirectory = 2

	// the user's Music directory
	USER_DIRECTORY_MUSIC UserDirectory = 3

	// the user's Pictures directory
	USER_DIRECTORY_PICTURES UserDirectory = 4

	// the user's shared directory
	USER_DIRECTORY_PUBLIC_SHARE UserDirectory = 5

	// the user's Templates directory
	USER_DIRECTORY_TEMPLATES UserDirectory = 6

	// the user's Movies directory
	USER_DIRECTORY_VIDEOS UserDirectory = 7

	// the number of enum values
	USER_N_DIRECTORIES UserDirectory = 8
)
