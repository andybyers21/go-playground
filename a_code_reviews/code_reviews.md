# Code Reviews

- Collaborative process
- Dev team colleagues or co-workers can provide critique, potential improvements
  and as questions.
- Add comments to specific lines of code.

- Maintain & normalise best practices (styling, naming conventions, etc.)
- Educate other them members on the code that's going into the code base.
- Reduces bus factor (one team member having too much knowledge on a specific
  area of the code)
- Allows for collaborative discussion.

## Where do they happen

- Git, GitHub etc.
- Usually in teh context of a pull request. 

## How to do one?

1. Figure out why the code was written and what's the purpose behind it.  - Changes in requirements?

- Issues?
- etc.

2. What is this code doing?

- How does this code fit into the overall project?

3. Understand what value can you add to the code review

- What is your experience with this code base?
- How familiar are you with the code/ lang?
- Have you built similar tools in the past?

4. Test the code.

- Does it do what it's supposed to do?
- How long does it take to run?
- Does it work within the non-functional parameters (performance etc.)
- Any errors?
- Do the changes require any modification to the rest of the code base?
- Do the changes require any additional set-up steps?
- Test the UI, can you break it?

5. Inspect the code:

- Naming conventions?
- DRY principals
- styling
- consistancy of the code
- consistancy with the rest of the project

6. Compile your review

- does it need a formal doc?
- comment the code?
- new branch?
- any other?

## Be Aware

- Your job is not to find the problems with the code author but to ensure the
  code functions as expected. 
- Be aware of nit picking, opinion based.
- Be respectful, know the code author
- Discuss!
- Explain your why.
