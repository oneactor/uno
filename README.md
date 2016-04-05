# uno
Golang cards game interface about uno and other games like uno.
基于go语言的纸牌游戏通用接口,以uno为例。

## V 1.0

## Interface

	context:
		The game context,control the game and turn logic.
		游戏上下文 控制游戏逻辑和回合流转
  
  	user:
  		The users info and status.
  		玩家信息与状态
  
  	desk:
  		All the cards.
  		牌堆
  
  	card:
  		Card type.
  		卡牌
  
## Base Struct

	BaseContext
	
	BaseUser
	
	BaseDesk
	
	BaseCard
	
Base structs implement the base functions of their Interface.We can also write or rewrite our own logic and functions in our own Structs.
基础类实现了以上接口的基础方法，我们也在我们自己的类里新的方法或重写已有方法以实现新的功能与逻辑。
  
## Example

  See the uno example code : uno/example/main.go
  
  以上Uno纸牌游戏例子
  
  See the Black Jack 21 example code : uno/example/main21.go
  
  以上21点游戏例子
  
## 说明

建议通过以下方式获取 

	go get -u github.com/jesusslim/uno
	
	go get -u github.com/jesusslim/uno
		
	go get -u github.com/jesusslim/uno
	
重要的事情说三遍！

因为这个项目我会从github同步到oschina等地方 如果您是从非github看到的 拉取下来可能包路径会有问题 因此强烈建议通过go get获取！

感谢各位支持