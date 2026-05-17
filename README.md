# Modern Checkers Web Platform

## 📌 Product Description

**What we built:** A modern, interactive web-platform for playing checkers. Currently, it features a clean, responsive UI built with Next.js and Tailwind CSS, complete with full game state persistence using LocalStorage. The architecture is explicitly designed with scalability in mind, separating the presentation layer from the core game logic.

**Who it is for:** This platform is designed for casual players looking for a quick, seamless local game of checkers (hot-seat style on a single screen). It also serves as a robust foundation for developers looking to see how classic board games can be implemented using modern web frameworks and containerization.

**Why it is valuable:** Rather than being just another generic checkers script, this project focuses on a highly extendable architecture and a frictionless developer/user experience. It is fully Dockerized—meaning anyone can launch the game in seconds without dealing with Node.js environments or dependency hell. It provides a reliable, crash-resilient local game experience out of the box.

---

## 🚀 Current Project Status (Level 2 Completed)

This project currently fulfills the **Level 2 (Medium)** requirements of the task.

**Working Features:**
* Playable two-player game on a single screen.
* Strict state management and UI updates built on React/Next.js.
* Game state persistence (if you accidentally close the tab, your game resumes exactly where you left off via LocalStorage).
* Fully Dockerized frontend environment.

### 🚧 Challenges & Future Roadmap (Level 3+)

While the initial goal was to build a fully distributed system with an AI opponent and online multiplayer, I encountered significant challenges during the research and implementation phases of these advanced features within the given time limit.

Specifically:
* **WebSockets (Online Multiplayer):** Synchronizing real-time game state securely between two different clients without race conditions requires a complex Observer/Pub-Sub pattern on the backend.
* **AI Player:** Implementing a performant AI (e.g., using the Minimax algorithm with Alpha-Beta pruning) required deep domain logic that outscoped the current time constraints.

Because I value clean, working code over broken, half-finished features, I chose to polish the Level 2 local experience and establish a solid architectural foundation (which is ready to connect to a Go-kit backend) rather than rushing the multiplayer and AI features.

---

## 🛠️ Tech Stack
* **Framework:** React / Next.js (TypeScript)
* **Styling:** Tailwind CSS
* **State Management:** Custom React Hooks + LocalStorage
* **Deployment:** Docker (Multi-stage build for a lightweight image)

---

## 🎮 How to Run the Game

You do not need to install JavaScript, Node.js, or NPM on your computer to run this project. You only need Docker.

1. **Clone the repository and navigate into the folder:**
```bash
git clone https://github.com/itelman/checkers-web
cd checkers-web
```

2. **Run the application:**

```bash
chmod +x entrypoint.sh
./entrypoint.sh
```

3. **Play:**

Open your browser and navigate to http://localhost:3000.