#!/bin/bash

SESH="dev"
WIN="main"

tmux select-pane -t "$SESH:$WIN.right"
tmux send-keys -t "$SESH:$WIN.right" "q" Enter
tmux send-keys -t "$SESH:$WIN.right" "go run main.go" Enter
