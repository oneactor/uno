# uno
Golang cards game interface about uno and other games like uno.

## V 1.0

## Interface

	context:The game context,control the game and turn logic.
  
  	user:The users info and status.
  
  	desk:All the cards.
  
  	card:Card type.
  
## Base Struct

	BaseContext
	
	BaseUser
	
	BaseDesk
	
	BaseCard
	
Base structs implement the base functions of their Interface.We can also write or rewrite our own logic and functions in our own Structs.
  
## Example

  See the uno example code : uno/example/main.go
  
  See the Black Jack 21 example code : uno/example/main21.go