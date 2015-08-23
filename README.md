# Laura
A secret diary system in Go.

<h2>About</h2>
Laura lets you keep your thoughts private with a password-protected digital diary. Diaries are kept in ~/Documents/laura, but this can be easily configured by changing the string in ```filesys.go```. Laura only lets you manage multiple diaries, each with their own unique name. 

Laura only supports appending text to a diary, because if you wanted to write something down, you should keep it.

<h2>Usage</h2>
```$ laura [command]```

Available Commands:
 * new [name]: Makes a new diary with given name.
 * list: Lists all diaries
 * addto [name] [text]: Adds text to a named diary. Prompts for password.
 * read [name]: Echoes contents of a named diary. Prompts for password.
 * delete [name]: Deletes a named diary.
 * help: Help about any command.
