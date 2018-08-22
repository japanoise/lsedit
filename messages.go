package main

import "fmt"

func helpScreen() {
	fmt.Println(`          MUFedit Help Screen.  Arguments in [] are optional.
    Any line not starting with a '.' is inserted at the current line.
Lines starting with '..', '."' , or '.:' are added with the '.' removed.
-------  st = start line   en = end line   de = destination line  -------
 .end                    Exits the editor with the changes intact.
 .abort                  Aborts the edit.
 .h                      Displays this help screen.
 .i [st]                 Changes the current line for insertion.
 .l [st [en]]            Lists the line(s) given. (if none, lists all.)
 .p [st [en]]            Like .l, except that it prints line numbers too.
 .del [st [en]]          Deletes the given lines, or the current one.
 .copy [st [en]]=de      Copies the given range of lines to the dest.
 .move [st [en]]=de      Moves the given range of lines to the dest.
 .find [st]=text         Searches for the given text starting at line start.
 .repl [st [en]]=/old/new  Replaces old text with new in the given lines.
 .join [st [en]]         Joins together the lines given in the range.
 .split [st]=text        Splits given line into 2 lines.  Splits after text
 .left [st [en]]         Aligns all the text to the left side of the screen.
 .center [st [en]]=cols  Centers the given lines for cols screenwidth.
 .right [st [en]]=col    Right justifies to column col.
 .indent [st [en]]=cols  Indents or undents text by cols characters
 .format [st [en]]=cols  Formats text nicely to cols columns.
---- Example line refs:  $ = last line, . = curr line, ^ = first line. ----`)
}

func welcome() {
	fmt.Println(`<    Welcome to the list editor.  You can get help by entering '.h'     >
< '.end' will exit and save the list.  '.abort' will abort any changes. >
<    To save changes to the list, and continue editing, use '.save'     >`)
}
