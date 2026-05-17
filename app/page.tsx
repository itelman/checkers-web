"use client";
import { useState } from "react";
import { useLocalStorage } from "../hooks/useLocalStorage";

// 0: Empty, 1: P1(Red), 2: P2(Black)
type BoardState = number[][];

const INITIAL_BOARD: BoardState = Array(8).fill(null).map((_, y) =>
    Array(8).fill(0).map((_, x) => {
        if ((x + y) % 2 !== 0) {
            if (y < 3) return 2;
            if (y > 4) return 1;
        }
        return 0;
    })
);

export default function CheckersGame() {
    const [board, setBoard] = useLocalStorage<BoardState>("checkers_board", INITIAL_BOARD);
    const [turn, setTurn] = useLocalStorage<number>("checkers_turn", 1);
    const [selectedPiece, setSelectedPiece] = useState<{x: number, y: number} | null>(null);


    const handleCellClick = (x: number, y: number) => {
        // 1. If clicking own piece, select it
        if (board[y][x] === turn) {
            setSelectedPiece({ x, y });
            return;
        }

        // 2. If clicking empty square and piece is selected, attempt move
        if (selectedPiece && board[y][x] === 0 && (x + y) % 2 !== 0) {
            // NOTE: In a full app, you would send this to your Go-Kit backend
            // via fetch('/api/game/move') to validate rules before setting state.

            const newBoard = [...board.map(row => [...row])];
            newBoard[y][x] = turn;
            newBoard[selectedPiece.y][selectedPiece.x] = 0;

            setBoard(newBoard);
            setSelectedPiece(null);
            setTurn(turn === 1 ? 2 : 1);
        }
    };


    /*
    const handleCellClick = async (x: number, y: number) => {
        if (board[y][x] === turn) {
            setSelectedPiece({ x, y });
            return;
        }

        if (selectedPiece && board[y][x] === 0 && (x + y) % 2 !== 0) {

            // 1. Define your backend URL (using an environment variable)
            const backendUrl = 'http://localhost:8888';

            try {
                // 2. Send the move to the Go-Kit backend
                const response = await fetch(`${backendUrl}/game/move`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        state: { board, current_turn: turn, winner: 0 },
                        from_x: selectedPiece.x,
                        from_y: selectedPiece.y,
                        to_x: x,
                        to_y: y
                    })
                });

                const data = await response.json();

                if (data.err) {
                    alert("Invalid move: " + data.err);
                    return;
                }

                // 3. Update the frontend with the validated state from the backend
                setBoard(data.state.board);
                setTurn(data.state.current_turn);
                setSelectedPiece(null);

            } catch (error) {
                console.error("Failed to reach the backend:", error);
            }
        }
    };
     */

    const resetGame = () => {
        setBoard(INITIAL_BOARD);
        setTurn(1);
        setSelectedPiece(null);
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
            <h1 className="text-3xl font-bold mb-4">Modern Checkers</h1>
            <p className="mb-4">Current Turn: Player {turn === 1 ? "1 (Red)" : "2 (Black)"}</p>

            <div className="border-4 border-gray-800">
                {board.map((row, y) => (
                    <div key={y} className="flex">
                        {row.map((cell, x) => {
                            const isDark = (x + y) % 2 !== 0;
                            const isSelected = selectedPiece?.x === x && selectedPiece?.y === y;

                            return (
                                <div
                                    key={`${x}-${y}`}
                                    onClick={() => handleCellClick(x, y)}
                                    className={`w-12 h-12 flex items-center justify-center
                    ${isDark ? 'bg-amber-800' : 'bg-amber-200'}
                    ${isSelected ? 'ring-4 ring-yellow-400 z-10' : ''}
                  `}
                                >
                                    {cell === 1 && <div className="w-8 h-8 rounded-full bg-red-600 shadow-md" />}
                                    {cell === 2 && <div className="w-8 h-8 rounded-full bg-gray-900 shadow-md" />}
                                </div>
                            );
                        })}
                    </div>
                ))}
            </div>

            <button onClick={resetGame} className="mt-6 px-4 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-700">
                Restart Game
            </button>
        </div>
    );
}