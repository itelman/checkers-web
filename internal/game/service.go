package game

import "context"

// GameState represents the board and current turn.
type GameState struct {
	Board       [][]int `json:"board"` // 0 = empty, 1 = p1, 2 = p2, 3 = p1_king, 4 = p2_king
	CurrentTurn int     `json:"current_turn"`
	Winner      int     `json:"winner"`
}

// Service defines the game operations.
type Service interface {
	NewGame(ctx context.Context) (GameState, error)
	ValidateMove(ctx context.Context, input ValidateMoveInput) (GameState, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) NewGame(ctx context.Context) (GameState, error) {
	// Standard 8x8 checkers starting positions
	board := make([][]int, 8)
	for i := range board {
		board[i] = make([]int, 8)
		for j := range board[i] {
			if (i+j)%2 != 0 {
				if i < 3 {
					board[i][j] = 2 // Player 2 (Top)
				} else if i > 4 {
					board[i][j] = 1 // Player 1 (Bottom)
				}
			}
		}
	}
	return GameState{Board: board, CurrentTurn: 1, Winner: 0}, nil
}

type ValidateMoveInput struct {
	State GameState `json:"state"`
	FromX int       `json:"from_x"`
	FromY int       `json:"from_y"`
	ToX   int       `json:"to_x"`
	ToY   int       `json:"to_y"`
}

func (s *service) ValidateMove(ctx context.Context, input ValidateMoveInput) (GameState, error) {
	// TODO: Implement Strategy Pattern here for deep validation.
	// 1. Check if it's the player's turn.
	// 2. Check diagonal boundaries and forced captures.
	// 3. Promote to King if reaching the opposite end.
	// For now, this is where your core Checkers rules engine lives.

	// Example state mutation:
	piece := input.State.Board[input.FromY][input.FromX]
	input.State.Board[input.FromY][input.FromX] = 0
	input.State.Board[input.ToY][input.ToX] = piece

	// Switch turn
	if input.State.CurrentTurn == 1 {
		input.State.CurrentTurn = 2
	} else {
		input.State.CurrentTurn = 1
	}

	return input.State, nil
}
