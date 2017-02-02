# regexnogada
Multiple regular expression engaging tool

# concept
This tool deals with 2 input files, and returns 1 text output.

Usage: regnogada sourcefile rulesfile outputfile

The [source file] will be read into memory as one.

The [rules file] must include regular expression (RE2) commands listed line by line.

This tool will engage RE2s to the source text, in sequence line by line.

If [output file] name is specified, then the output file will be made.

If output file name is omitted, then result text will be printed out to STDOUT. So, user can save the result to file using OS pipe function. (for an example: regexnogada sourcefile rulesfile > outputfile )
