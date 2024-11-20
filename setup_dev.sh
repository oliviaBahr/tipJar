#!/bin/bash

if ! command -v tmux &>/dev/null; then
    echo "tmux is not installed. Please install tmux to use auto-reload feature."
    echo "Run: brew install tmux (Mac) or apt-get install tmux (Ubuntu)"
    exit 1
fi

# Define variables first
DIR=$(pwd)
SESSION="dev"
WINDOW="main"

# Kill existing session if it exists
tmux kill-session -t $SESSION 2>/dev/null

# Create session with current directory and shell
tmux new-session -d -s $SESSION -n $WINDOW "cd $DIR; $SHELL"

# Enable mouse support
tmux set -g mouse on

# Name the current pane 'air'
tmux select-pane -t $SESSION:$WINDOW -T air

# Split into two main panes (left and right):
tmux split-window -h -p 50 -t $SESSION:$WINDOW
tmux select-pane -t $SESSION:$WINDOW.0 -T air
tmux select-pane -t $SESSION:$WINDOW.1 -T main

# Split left pane vertically
tmux select-pane -t $SESSION:$WINDOW.0
tmux split-window -v -p 95 -t $SESSION:$WINDOW.0
tmux select-pane -t $SESSION:$WINDOW.2 -T logs

# Send commands to specific panes using their index
# Air in top-left (pane 0)
tmux send-keys -t "$SESSION:$WINDOW.0" "cd $DIR && air" Enter

# Logs in bottom-left (pane 1)
tmux send-keys -t "$SESSION:$WINDOW.1" "cd $DIR && tail -f log.log" Enter

# Main.go in right pane (pane 2)
tmux send-keys -t "$SESSION:$WINDOW.2" "cd $DIR && clear" Enter

# Focus the right pane (logs)
tmux select-pane -t $SESSION:$WINDOW.2

# Attach if we're not already in tmux
if [ -z "$TMUX" ]; then
    tmux attach -t $SESSION
fi
