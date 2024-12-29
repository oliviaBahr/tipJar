#!/bin/bash

SESH="dev"
WIN="main"
BUILD="./tmp/tipJar"

tmux select-pane -t "$SESH:$WIN.right"
tmux send-keys -t "$SESH:$WIN.right" "q" Enter
tmux send-keys -t "$SESH:$WIN.right" "$BUILD log" Enter
