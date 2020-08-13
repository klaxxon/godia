# GoDia
GoDia is a simple go program that takes golang files and produces a .dia diagram of the structs and their references.  The diagram below is generated by running the program on itself.<br/>

![Self Diagram](/self.png)
(colors added for demo)<br/>
<br/>
This has some bugs and is a work in progress. It has been working for some large Go projects, but has issues if I include vendor directories.<br/>
<br/>
<b>Features:</b>
<ul>
  <li>Uses DIAs UML diagram</li>
  <li>Uses "implements" connections between the field and associated struct type</li>
  <li>The "stereotype" is set to the package name</li>
  <li>All fields and their types are shown</li>
  <li>If the .dia file already exists when this is run, it will attempt to keep the position, color and size the same.</li>
</ul>
<br/>
<B>Running</b>
<br/>
To simply generate the start of the diagram above....<br/>
From within the godia directory:<br/>
<br/>
<code>
$> go run *.go -i ./ -o self
</code>
<br/>
This will generate a self.dia file you can open in dia. It places all of the structs on top of each other so you will need to move then around.  However,  if you save your changes and re-run, it should create a new self.dia but respect the positions you set on the existing structs (hopefully).<br/>
<br/>
For large projects with packages you do not want (like vendor) simply add a comma delimited line of directories to ignore.<br/>
<br/>
<code>
$> go run *.go -i ./ -o self -ignore vendor,.vscode,ignoredir
</code>
<br/>
<br/>
<b>Issues</b><br/>
Including the vendor directory (4000+ structs!) does create some XML escape issues I have not fixed.</br>





