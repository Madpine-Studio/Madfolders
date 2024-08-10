# Madfolders
Madfolders is a CLI tool, for internal usage, intended to create the default project structure and iterate over it.


## Prerequisites
- GO compiler
- Make tool

## How to build
The binary must appear on the binary folder if it builds correctly by using the command below

```bash
  make build
```

## Adding to Path

### Windows

#### Automagically
Try running this first if it doesn't work try the manual way
```bash
  Add-Content -Path $Profile -Value "`$Env:PATH += `"$($PWD.Path)\binary;`""
```
#### Manually
On windows is a bit tricky somethimes buggy, but you have to acess your powershell `$Profile` by typing the
command below

```bash
    notepad $Profile
```

Your profile will be displayed on a notepad and all you have to do is add the string below to the file and save it. 
Please do not forget the semicolon!!!

```bash
  $Env:PATH=$Env:PATH + "C:\path\to\your\binary;"
```


