#!/bin/bash

if ! command -v tmux &>/dev/null; then
    echo "tmux is not installed. Please install tmux to use auto-reload feature."
    echo "Run: brew install tmux (Mac) or apt-get install tmux (Ubuntu)"
    exit 1
fi

## general setup

DIR=$(pwd)
SESH="dev"
WIN="main"

# Kill existing session if it exists
tmux kill-session -t $SESH 2>/dev/null

# Create session with current directory and shell
tmux new-session -d -s $SESH -n $WIN "cd $DIR; $SHELL"

# Enable mouse support
tmux set -g mouse on

# Make status bar transparent/minimal
tmux set -g status-style bg=default
tmux set -g status-fg darkgreen
tmux set -g status-right ""

## panes and commands

# Split into two main panes (left and right):
tmux split-window -h -p 50 -t $SESH:$WIN

# Split left pane vertically
tmux select-pane -t $SESH:$WIN.left
tmux split-window -v -p 85 -t $SESH:$WIN.left

# Air logs in top-left (pane 0)
tmux send-keys -t "$SESH:$WIN.top-left" "cd $DIR && air" Enter

# Main logs in bottom-left (pane 1)
tmux send-keys -t "$SESH:$WIN.bottom-left" "cd $DIR && tail -f tmp/log.log" Enter

# Main.go in right pane (pane 2)
tmux select-pane -t $SESH:$WIN.right
# tmux send-keys -t "$SESH:$WIN.right" "cd $DIR && go run main.go" Enter

# Attach if we're not already in tmux
if [ -z "$TMUX" ]; then
    tmux attach -t $SESH
fi
