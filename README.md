# regexnogada
Multiple regular expression engaging tool

# concept
This tool deals with source text file and rules text file, and returns 1 text output.
Each lines in rules file will be engaged to source text in sequence.
Result text will be on output.

# Usage
Usage: regnogada sourcefile rulesfile outputfile
 (or:) regnogada sourcefile rulesfile

The [source file] will be read into memory as one.

The [rules file] must include regular expression (RE2) commands listed line by line.

This tool will engage RE2s to the source text, in sequence line by line.

If [output file] name is specified, then the output file will be made.

If output file name is omitted, then result text will be printed out to STDOUT. So, user can save the result to file using OS pipe function. (for an example: regexnogada sourcefile rulesfile > outputfile )

# Rule file
Each line would has one or two regular expressions(RE2).

First RE2 is for matching, and second RE2 is for replacement.

Those two RE2s must be separated by tab character. (remind your editor tab setting)

If second RE2 is omitted, then matching text will be erased.

Inside of rule file, the line starts with # will be ignored (comment line).

Empty line (null or only whitespace character filled) will be ignored.
