- [ ] Fix camera : translate > rotate > translate
- [ ] Clean up
- [ ] Shadows
- [ ] Textures
- [ ] Layout design
- [ ] Collision system
- [ ] Interactive viewport

---

## Shadows
- Compute new z-buffer for each point or direction light (compute distance to light instead of distance to camera)
- Store light results in an array
- When drawing each point, check if distance to light is smaller than z-buffer value
- If equal, apply light
- If not, only apply other light sources

<!--->
    Very costly for multiple lights !

