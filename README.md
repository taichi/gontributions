# GitHub Contribution Graph Hack

# Usage

## Build
```
go get github.com/taichi/gontributions
go install github.com/taichi/gontributions
```

## Make your ASCII ART
requirements are
* 8 lines
    * last line must be blank
* 50 charactors
* use ' ','-','*','@','#'. right one is more dark green.

## Execute gontributions command
if you don't specify source file, gontributions uses **source.txt**

## Push generated repository to GitHub

## Options

key | description | default value
----- | ------------- | -----------
name | author and commiter name | read from your global config
email | author and commiter email | read from your global config
repo | new repository destination path | ./repo
base | number of commits for graph | 10
msg | commit message | happy hacking !!

### Options example

```
gontributions -name=anonymous -mail=anonymous@example.com yamashiro.txt
```

# Examples

```
#     # #   #     #   #    ####  ### #####   ##* 
 #   #  #   ##   ##   #   #*   #  #  #*   # #* #*
  # #  ###  #*# # #  ###  #*      #  #*   ##*   #
   #   #@#  #* #  #  #@#   ####*  #  ##### #*   #
   #  ##### #*    # #####      #  #  #* #  #*   #
   #  #   # #*    # #   # #*   #  #  #*  #  #* #*
   # ##   ###*    ###   ## ####  ### #*   #  ##* 
```

![](https://raw.github.com/taichi/gontributions/master/example.png)

# see also
* [Viewing contributions](https://help.github.com/articles/viewing-contributions)
* <https://github.com/gelstudios/gitfiti>
* <https://github.com/jbranchaud/commitart>
* <https://github.com/tricknotes/we-love-ember>
* http://ascii.mastervb.net/
    * you should use **clr8x10.flf** or **banner.flf**
