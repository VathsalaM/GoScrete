package Interface

import "github.com/VathsalaM/GoScrete/Interface"

type Element interface{
	UpdatePosition(Interface.Position) Element
}
