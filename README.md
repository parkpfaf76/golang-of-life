# Summary

Game of life implementation in go. Starting simple with desktop app for user to play game of life on 100x100 grid.
Will iternate on design to make game more configurable for different types of experimentation.

## Future plans

Eventually want to develop into a 3d game of life model with variety of features, here's some current ideas:
- seeds
- user generated rules-
- saved configurations
- rotations
- zoom
- multiple game of life instances running in canvas (universes)

Will flesh out features later.

# Version 1.0.0

1. make mvp of simple conway game of life
2. hardcode 100x100 grid
3. allow user to configure/clear state before playing

# Architecture

## Model

### Requirements

- cell, grid, and rules

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