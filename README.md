# godia
GoDia is a simple go program that takes golang files and produces a .dia diagram of the structs and their references.  The diagram below is generated by running the program on itself.
![Self Diagram](/self.png)
(colors added for demo)<br/>
<br/>
This has some bugs and is a work in progress.  I threw it together this morning so I could better visualize a large go project I am working on.  However, it does work for my purposes.<br/>
<br/>
# Features:<br/>
<ul>
  <li>Uses DIAs UML diagram</li>
  <li>The "stereotype" is set to the package name</li>
  <li>All fields and their types are shown</li>
  <li>If the .dia file already exists when this is run, it will attempt to keep the position, color and size the same.</li>
</ul>
<br/>
# Running<br/>
<br/>
To simply generate the start of the diagram above....<br/>
From within the godia directory:<br/>
<br/>
```bash
$> go run *.go ./ self
```
<br/>
This will generate a self.dia file you can open in dia. It places all of the structs on top of each other so you will need to move then around.  However,  if you save your changes and re-run, it should create a new self.dia but respect the positions you set on the existing structs (hopefully).<br/>
<br/>
For large projects with packages you do not want (like vendors) simply add a comma delimited line of directopries to ignore.<br/>
<br/>
```bash
$> go run *.go self vendors,.vscode,ignodrdir
```



