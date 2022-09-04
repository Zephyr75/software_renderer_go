- [ ] Fix camera : translate > rotate > translate
- [ ] Clean up
- [ ] Shadows
- [ ] Textures : convert png/jpg to bitmap image
- [ ] Layout design
- [ ] Collision system
- [ ] Interactive viewport
- [ ] Anti-aliasing

---

## Shadows
- Compute new z-buffer for each point or direction light (compute distance to light instead of distance to camera)
- Store light results in an array
- When drawing each point, check if distance to light is smaller than z-buffer value
- If equal, apply light
- If not, only apply other light sources

        Very costly for multiple lights ! 


## Textures

- Convert png/jpg to bitmap image from the **image** package
- Apply formula to get x value from vertices
- Apply formula to get y value from vertices
- Fix warp by adding a value of 1 going through perspective by which you divide the resulting x and y