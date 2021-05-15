package game

// Renderer - interface to implement by any entity that can provide rendering
// logic.
type Renderer interface {
	Render(board Board, currentSign Sign)
}
