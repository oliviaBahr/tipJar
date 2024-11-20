#!/bin/bash

if ! command -v tmux &>/dev/null; then
    echo "tmux is not installed. Please install tmux to use auto-reload feature."
    echo "Run: brew install tmux (Mac) or apt-get install tmux (Ubuntu)"
    exit 1
fi

SESSION="dev"
WINDOW="main"

# Only send commands if the tmux session exists
if tmux has-session -t $SESSION 2>/dev/null; then
    # Stop the current processes
    tmux send-keys -t "$SESSION:$WINDOW.2" C-c
    sleep 0.15
    tmux send-keys -t "$SESSION:$WINDOW.2" "clear && go run main.go" Enter
else
    echo "Tmux session '$SESSION' not found. Please run setup_dev.sh first."
    exit 1
fi
