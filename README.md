# regexnogada
Multiple regular expression engaging tool

# concept
This tool deals with 2 input files, and returns 1 text stream output.

2 input text files are: [source file], [rules file]

The [source file] will be read into memory as one.

The [rules file] must include regular expression (RE2) commands listed line by line.

This tool will engage RE2s to the source text, in sequence line by line.

Result text will be printed out to STDOUT.

User can save the result to file using OS pipe function.

(for an example: regexnogada sourcefile rulesfile > outputfile )
