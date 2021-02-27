# Vimprover

Vimprover helps you get rid of bad keyboard habits using your computer in general, and some specifics for using Vim. Vimprover detects poor patterns by monitoring your keyboard inputs and helps you to improve using [Aversion Therapy](https://en.wikipedia.org/wiki/Aversion_therapy). This sounds fancy and all, but in lay mens terms this means it plays an annoying sound if you violate the checks.

## Improvement checks

### Left/Right shift use
It's more efficient, more ergonomical, and overall just better to use both shift keys on a keyboard. Vimprover sees it as a violation if you use the right shift key together with keys on the right half of the keyboard, or use the left shift key with keys on the left half. This will over time help users get adjusted to using both shift keys

### Repeated use of H|J|K|L
The keys H, J, K and L are used to move the cursor one step at a time in vim. Vim however offers tons of better options for moving your cursor, and the repeated use of moving it one step is inefficient. Hence vimprover sees 5 or more consecutive clicks on one of these keys as a violation. This will help you to start using better cursor navigation in vim

## Execution
Due to vimprover needing to run a keylogger on your inputs to detect patterns, there have been no efforts into automating the execution of it. Hence I recommend just cloning this repository and using the `go run` command to start it, and leave the terminal window where it runs open in the background. There's a minimal run script included in the repository, so you can start Vimprover from the vimprover directory.

```$
sudo ./run
```
