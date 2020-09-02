你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# GitHub Contribution Graph Hack

# Usage

## Build

Download from [build.io](http://gobuild.io/download/https:/github.com/taichi/gontributions).

or 

```
go get github.com/taichi/gontributions
go install github.com/taichi/gontributions
```

## Make your ASCII ART
requirements are
* 8 lines
    * last line must be blank
* 50 characters
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

## source.txt
```
#     # #   #     #   #    ####  ### #####   ##* 
 #   #  #   ##   ##   #   #*   #  #  #*   # #* #*
  # #  ###  #*# # #  ###  #*      #  #*   ##*   #
   #   #@#  #* #  #  #@#   ####*  #  ##### #*   #
   #  ##### #*    # #####      #  #  #* #  #*   #
   #  #   # #*    # #   # #*   #  #  #*  #  #* #*
   # ##   ###*    ###   ## ####  ### #*   #  ##* 
```
## rendered graph
![](https://raw.github.com/taichi/gontributions/master/example.png)

# see also
* [Viewing contributions](https://help.github.com/articles/viewing-contributions)
* <https://github.com/gelstudios/gitfiti>
* <https://github.com/jbranchaud/commitart>
* <https://github.com/tricknotes/we-love-ember>
* http://ascii.mastervb.net/
    * you should use **clr8x10.flf** or **banner.flf**
