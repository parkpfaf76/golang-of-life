# Plan - Version 1.0.0

1. make mvp of simple conway game of life
2. hardcode 100x100 grid
3. allow user to configure/clear state before playing

# Architecture

## Model

### Requirements

- cells
- grid

### Classes

- Cell
- Grid

## View

### Requirements

- on play updates state 60 frames a second (block user interaction unless paused)

### Classes

- Canvas
- Image

## Controller

### Requirements

- play/pause
- next state calculations for grid
- rules for cells to abide for state update

### Classes

- Controller

Controller manages View/Data