# Boids (Bird-oid object)

# Content
- [What is SDL?](#what-is-sdl)
- [About this project](#about-this-project)
    - [Compatibility](##compatibility)
    - [Showcase](##showcase)
- [How to?](#how-to)
    - [Prerequisites](##prerequisites)
    - [Build](##build)

# What is SDL?
> Simple DirectMedia Layer is a cross-platform development library 
> designed to provide low level access to audio, keyboard, mouse, joystick, 
> and graphics hardware via OpenGL and Direct3D.

[SDL official page.](https://www.libsdl.org/)

# About this project
## Inspiration
[Video Source](https://www.youtube.com/watch?v=bqtqltqcQhw)
[Modeling Natural Systems Stanford](https://cs.stanford.edu/people/eroberts/courses/soco/projects/2008-09/modeling-natural-systems/boids.html)

> Boids is an artificial life simulation originally developed by Craig Reynolds.
> The aim of the simulation was to replicate the behavior of flocks of birds.
> Instead of controlling the interactions of an entire flock, however, the
> Boids simulation only specifies the behavior of each individual bird.
> With only a few simple rules, the program manages to generate a result that
> is complex and realistic enough to be used as a framework for computer graphics
> applications such as computer generated behavioral animation in motion picture films.

## Compatibility
Currently tested only on MacOS arm64.
Feel free to test and open PRs/issues to include/fix for other OS.

## Showcase
<img src="./assets/boids-showcase.gif" />

# How to?
## Prerequisites
* Go >= 1.23.4
* [SDLv2](https://github.com/veandco/go-sdl2?tab=readme-ov-file#requirements)

## Build and Run
In the cloned directory:
```sh
    go build . && ./boids
```

> Boids are afraid of mouse motions.
